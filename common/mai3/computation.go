package mai3

import (
	"fmt"

	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
)

func ComputeAccount(p *model.LiquidityPoolStorage, perpetualIndex int64, a *model.AccountStorage) (*model.AccountComputed, error) {
	perpetual, ok := p.Perpetuals[perpetualIndex]
	if !ok {
		return nil, fmt.Errorf("perpetual %d not found in the pool", perpetualIndex)
	}
	positionValue := perpetual.MarkPrice.Mul(a.PositionAmount.Abs())
	positionMargin := positionValue.Mul(perpetual.InitialMarginRate)
	maintenanceMargin := positionValue.Mul(perpetual.MaintenanceMarginRate)
	reservedCash := _0
	if !a.PositionAmount.IsZero() {
		reservedCash = perpetual.KeeperGasReward
	}
	availableCashBalance := a.CashBalance.Sub(a.PositionAmount.Mul(perpetual.UnitAccumulativeFunding))
	marginBalance := availableCashBalance.Add(perpetual.MarkPrice.Mul(a.PositionAmount))
	availableMargin := marginBalance.Sub(decimal.Max(reservedCash, positionMargin))
	withdrawableBalance := decimal.Max(_0, availableMargin)
	isMMSafe := marginBalance.GreaterThanOrEqual(decimal.Max(reservedCash, maintenanceMargin))
	isIMSafe := marginBalance.GreaterThanOrEqual(decimal.Max(reservedCash, positionMargin))
	isMarginSafe := marginBalance.GreaterThanOrEqual(reservedCash)

	return &model.AccountComputed{
		PositionValue:        positionValue,
		PositionMargin:       positionMargin,
		MaintenanceMargin:    maintenanceMargin,
		AvailableCashBalance: availableCashBalance,
		MarginBalance:        marginBalance,
		AvailableMargin:      availableMargin,
		WithdrawableBalance:  withdrawableBalance,
		IsMMSafe:             isMMSafe,
		IsIMSafe:             isIMSafe,
		IsMarginSafe:         isMarginSafe,
	}, nil
}

func ComputeAMMTrade(p *model.LiquidityPoolStorage, perpetualIndex int64, trader *model.AccountStorage, amount decimal.Decimal) (*model.AccountComputed, bool, decimal.Decimal, error) {
	if amount.IsZero() {
		return nil, false, _0, fmt.Errorf("bad amount")
	}
	perpetual, ok := p.Perpetuals[perpetualIndex]
	if !ok {
		return nil, false, _0, fmt.Errorf("perpetual %d not found in the pool", perpetualIndex)
	}

	oldOpenInterest := perpetual.OpenInterest
	newOpenInterest := oldOpenInterest

	// AMM
	_, deltaAMMAmount, tradingPrice, err := ComputeAMMPrice(p, perpetualIndex, amount)
	if err != nil {
		logger.Errorf("ComputeAMMPrice error:%s", err)
		return nil, false, _0, err
	}
	if !deltaAMMAmount.Neg().Equal(amount) {
		logger.Errorf("trading amount mismatched %s != %s", deltaAMMAmount, amount)
		return nil, false, _0, fmt.Errorf("trading amount mismatched")
	}

	// fee
	lpFee, err := ComputeFee(tradingPrice, deltaAMMAmount, perpetual.LpFeeRate)
	if err != nil {
		return nil, false, _0, err
	}

	// trader
	afterTrade, tradeIsSafe, err := ComputeTradeWithPrice(p, perpetualIndex, trader, tradingPrice, deltaAMMAmount.Neg(),
		perpetual.LpFeeRate.Add(p.VaultFeeRate).Add(perpetual.OperatorFeeRate))
	if err != nil {
		logger.Errorf("ComputeTradeWithPrice trader err:%s", err)
		return nil, false, _0, err
	}

	newOpenInterest = computeOpenInterest(newOpenInterest, trader.PositionAmount, deltaAMMAmount.Neg())

	// new AMM
	fakeAMMAccount := &model.AccountStorage{
		CashBalance:    p.PoolCashBalance,
		WalletBalance:  decimal.Zero,
		PositionAmount: perpetual.AmmPositionAmount,
	}
	_, _, err = ComputeTradeWithPrice(p, perpetualIndex, fakeAMMAccount, tradingPrice, deltaAMMAmount, _0)
	if err != nil {
		logger.Errorf("ComputeAMMTrade fakeAMMAccount err:%s", err)
		return nil, false, _0, err
	}
	fakeAMMAccount.CashBalance = fakeAMMAccount.CashBalance.Add(lpFee)
	newOpenInterest = computeOpenInterest(newOpenInterest, perpetual.AmmPositionAmount, deltaAMMAmount)

	p.PoolCashBalance = fakeAMMAccount.CashBalance
	perpetual.AmmPositionAmount = fakeAMMAccount.PositionAmount
	perpetual.OpenInterest = newOpenInterest

	// check open interest limit
	if newOpenInterest.GreaterThan(oldOpenInterest) {
		context := initAMMTradingContext(p, perpetualIndex)
		err = computeAMMPoolMargin(context, context.OpenSlippageFactor, true)
		if err != nil {
			logger.Errorf("ComputeAMMTrade computeAMMPoolMargin err:%s", err)
			return nil, false, _0, err
		}
		limit := context.PoolMargin.Mul(perpetual.MaxOpenInterestRate).Div(perpetual.IndexPrice)
		if newOpenInterest.GreaterThan(limit) {
			return nil, false, _0, fmt.Errorf("ComputeAMMTrade open interest exceeds limit: %s > %s", newOpenInterest, limit)
		}
	}

	return afterTrade, tradeIsSafe, tradingPrice, nil
}

// > 0 if more collateral required
func computeOpenInterest(oldOpenInterest, oldPosition, tradeAmount decimal.Decimal) decimal.Decimal {
	newOpenInterest := oldOpenInterest
	newPosition := oldPosition.Add(tradeAmount)
	if oldPosition.GreaterThan(_0) {
		newOpenInterest = newOpenInterest.Sub(oldPosition)
	}
	if newPosition.GreaterThan(_0) {
		newOpenInterest = newOpenInterest.Add(newPosition)
	}
	return newOpenInterest
}

func ComputeAMMPrice(p *model.LiquidityPoolStorage, perpetualIndex int64, amount decimal.Decimal) (deltaAMMMargin, deltaAMMAmount, tradingPrice decimal.Decimal, err error) {
	if amount.IsZero() {
		err = fmt.Errorf("bad amount")
		return
	}
	ammTrading, err := computeAMMInternalTrade(p, perpetualIndex, amount.Neg())
	if err != nil {
		return
	}
	deltaAMMMargin = ammTrading.DeltaMargin
	deltaAMMAmount = ammTrading.DeltaPosition
	tradingPrice = deltaAMMMargin.Div(deltaAMMAmount).Abs()
	return
}

func ComputeTradeWithPrice(p *model.LiquidityPoolStorage, perpetualIndex int64, a *model.AccountStorage, price, amount, feeRate decimal.Decimal) (*model.AccountComputed, bool, error) {
	if price.LessThanOrEqual(_0) || amount.IsZero() {
		return nil, false, fmt.Errorf("bad price %s or amount %s", price, amount)
	}

	close, open := utils.SplitAmount(a.PositionAmount, amount)
	if !close.IsZero() {
		if err := ComputeDecreasePosition(p, perpetualIndex, a, price, close); err != nil {
			return nil, false, err
		}
	}

	if !open.IsZero() {
		if err := ComputeIncreasePosition(p, perpetualIndex, a, price, open); err != nil {
			return nil, false, err
		}
	}

	// TODO: consider order referrer fee rate
	fee, err := ComputeFee(price, amount, feeRate)
	if err != nil {
		return nil, false, err
	}

	adjustMargin = adjustMarginLeverage(pe)

	a.CashBalance = a.CashBalance.Sub(fee)

	// open position requires margin > IM. close position requires !bankrupt
	afterTrade, err := ComputeAccount(p, perpetualIndex, a)
	if err != nil {
		return nil, false, err
	}

	tradeSafe := afterTrade.IsMarginSafe
	if !open.IsZero() {
		tradeSafe = afterTrade.IsIMSafe
	}

	return afterTrade, tradeSafe, nil
}

func adjustMarginLeverage(p *model.LiquidityPoolStorage, perpetualIndex int64, trader *model.AccountComputed, deltaPosition, deltaCash, totalFee decimal.Decimal) (decimal.Decimal, error) {
	perpetual, ok := p.Perpetuals[perpetualIndex]
	if !ok {
		return _0, fmt.Errorf("perpetual %d not found in the pool", perpetualIndex)
	}
	adjustCollateral := _0
	closed, opened := utils.SplitAmount(trader.PositionAmount.Sub(deltaPosition), deltaPosition)
	if !closed.IsZero() && opened.IsZero() {
		adjustCollateral = adjustClosedMargin(perpetual, trader, closed, deltaCash, totalFee)
	} else if !open.IsZero() {
		adjustCollateral := adjustOpenedMargin(perpetual, trader, closed, deltaCash, totalFee)
	}
}

func adjustClosedMargin(perp *model.PerpetualStorage, trader *model.AccountStorage, closed, deltaCash, totalFee decimal.Decimal) decimal.Decimal {
	adjustCollateral := _0
	// when close, keep the effective leverage
	// -withdraw == (availableCash2 * close - deltaCash * position2) / position1 + fee
	adjustCollateral = trader.CashBalance.Sub(trader.PositionAmount.Mul(perp.UnitAccumulativeFunding)).Mul(closed)
	adjustCollateral = adjustCollateral.Sub(deltaCash.Mul(trader.PositionAmount))
	adjustCollateral = adjustCollateral.Div(trader.PositionAmount.Sub(closed))
	adjustCollateral = adjustCollateral.Add(totalFee)
	// withdraw only when IM is satisfied
	// limit := totalFee.Sub(perp.)
	return adjustCollateral
}

func adjustOpenedMargin(perp *model.PerpetualStorage, trader *model.AccountStorage, deltaPosition, closed, opened, totalFee decimal.Decimal) decimal.Decimal {
	adjustCollateral := _0
	return adjustCollateral
}

func ComputeDecreasePosition(p *model.LiquidityPoolStorage, perpetualIndex int64, a *model.AccountStorage, price, amount decimal.Decimal) error {
	perpetual, ok := p.Perpetuals[perpetualIndex]
	if !ok {
		return fmt.Errorf("perpetual %d not found in the pool", perpetualIndex)
	}
	oldAmount := a.PositionAmount
	if oldAmount.IsZero() || amount.IsZero() || utils.HasTheSameSign(oldAmount, amount) {
		return fmt.Errorf("invalid amount or position, position:%s, amount:%s", oldAmount, amount)
	}

	if price.LessThanOrEqual(_0) {
		return fmt.Errorf("invalid price %s", price)
	}

	if oldAmount.Abs().LessThan(amount.Abs()) {
		return fmt.Errorf("position size is less than amount. position:%s, amount:%s", oldAmount, amount)
	}
	a.CashBalance = a.CashBalance.Sub(price.Mul(amount)).Add(perpetual.UnitAccumulativeFunding.Mul(amount))
	a.PositionAmount = a.PositionAmount.Add(amount)
	return nil
}

func ComputeIncreasePosition(p *model.LiquidityPoolStorage, perpetualIndex int64, a *model.AccountStorage, price, amount decimal.Decimal) error {
	perpetual, ok := p.Perpetuals[perpetualIndex]
	if !ok {
		return fmt.Errorf("perpetual %d not found in the pool", perpetualIndex)
	}
	oldAmount := a.PositionAmount
	if price.LessThanOrEqual(_0) {
		return fmt.Errorf("invalid price %s", price)
	}
	if amount.IsZero() {
		return fmt.Errorf("invalid amount %s", amount)
	}
	if !oldAmount.IsZero() && !utils.HasTheSameSign(oldAmount, amount) {
		return fmt.Errorf("invalid amount or position, position:%s, amount:%s", oldAmount, amount)
	}
	a.CashBalance = a.CashBalance.Sub(price.Mul(amount)).Add(perpetual.UnitAccumulativeFunding.Mul(amount))
	a.PositionAmount = a.PositionAmount.Add(amount)
	return nil
}

func ComputeFee(price, amount, feeRate decimal.Decimal) (decimal.Decimal, error) {
	if price.LessThanOrEqual(_0) || amount.IsZero() {
		return _0, fmt.Errorf("invalid price or admount. price:%s amount: %s", price, amount)
	}
	return price.Mul(amount.Abs()).Mul(feeRate), nil
}
