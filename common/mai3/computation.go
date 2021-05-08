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
	leverage := _0
	if positionValue.GreaterThan(_0) && marginBalance.GreaterThan(_0) {
		leverage = positionValue.Div(marginBalance)
	}

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
		Leverage:             leverage,
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

	// trader
	afterTrade, tradeIsSafe, totalFee, err := ComputeTradeWithPrice(p, perpetualIndex, trader, tradingPrice, deltaAMMAmount.Neg(),
		perpetual.LpFeeRate.Add(p.VaultFeeRate).Add(perpetual.OperatorFeeRate))
	if err != nil {
		logger.Errorf("ComputeTradeWithPrice trader err:%s", err)
		return nil, false, _0, err
	}

	// lp fee
	lpFee := totalFee.Mul(perpetual.LpFeeRate).Div(perpetual.LpFeeRate.Add(p.VaultFeeRate).Add(perpetual.OperatorFeeRate))

	newOpenInterest = computeOpenInterest(newOpenInterest, trader.PositionAmount, deltaAMMAmount.Neg())

	// new AMM
	fakeAMMAccount := &model.AccountStorage{
		CashBalance:    p.PoolCashBalance,
		WalletBalance:  decimal.Zero,
		PositionAmount: perpetual.AmmPositionAmount,
	}
	_, _, _, err = ComputeTradeWithPrice(p, perpetualIndex, fakeAMMAccount, tradingPrice, deltaAMMAmount, _0)
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

func ComputeTradeWithPrice(p *model.LiquidityPoolStorage, perpetualIndex int64, a *model.AccountStorage, price, amount, feeRate decimal.Decimal) (*model.AccountComputed, bool, decimal.Decimal, error) {
	if price.LessThanOrEqual(_0) || amount.IsZero() {
		return nil, false, _0, fmt.Errorf("bad price %s or amount %s", price, amount)
	}

	close, open := utils.SplitAmount(a.PositionAmount, amount)
	if !close.IsZero() {
		if err := ComputeDecreasePosition(p, perpetualIndex, a, price, close); err != nil {
			return nil, false, _0, err
		}
	}

	if !open.IsZero() {
		if err := ComputeIncreasePosition(p, perpetualIndex, a, price, open); err != nil {
			return nil, false, _0, err
		}
	}

	afterTrade, err := ComputeAccount(p, perpetualIndex, a)
	if err != nil {
		return nil, false, _0, err
	}
	// TODO: consider order referrer fee rate
	fee, err := ComputeFee(!open.IsZero(), price, amount, feeRate, afterTrade)
	if err != nil {
		return nil, false, _0, err
	}

	// ajust margin
	adjustMargin, err := adjustMarginLeverage(p, perpetualIndex, afterTrade, a, price, close, open, fee)
	if err != nil {
		return nil, false, _0, err
	}

	a.CashBalance = a.CashBalance.Add(adjustMargin).Sub(fee)

	// open position requires margin > IM. close position requires !bankrupt
	afterTrade, err = ComputeAccount(p, perpetualIndex, a)
	if err != nil {
		return nil, false, _0, err
	}

	tradeSafe := afterTrade.IsMarginSafe
	if !open.IsZero() {
		tradeSafe = afterTrade.IsIMSafe
	}

	return afterTrade, tradeSafe, fee, nil
}

func adjustMarginLeverage(p *model.LiquidityPoolStorage, perpetualIndex int64, a *model.AccountComputed, trader *model.AccountStorage, price, close, open, totalFee decimal.Decimal) (decimal.Decimal, error) {
	perpetual, ok := p.Perpetuals[perpetualIndex]
	if !ok {
		return _0, fmt.Errorf("perpetual %d not found in the pool", perpetualIndex)
	}
	adjustCollateral := _0
	deltaPosition := close.Add(open)
	deltaCash := deltaPosition.Mul(price).Neg()
	position2 := trader.PositionAmount
	if !close.IsZero() && open.IsZero() {
		// close only
		// when close, keep the effective leverage
		// -withdraw == (availableCash2 * close - deltaCash * position2) / position1 + fee
		adjustCollateral = a.AvailableCashBalance.Mul(close)
		adjustCollateral = adjustCollateral.Sub(deltaCash.Mul(position2))
		adjustCollateral = adjustCollateral.Div(position2.Sub(close))
		adjustCollateral = adjustCollateral.Add(totalFee)
		// withdraw only when IM is satisfied
		limit := totalFee.Sub(a.AvailableMargin)
		adjustCollateral = decimal.Max(adjustCollateral, limit)
		// never deposit when close positions
		adjustCollateral = decimal.Min(adjustCollateral, _0)
		return adjustCollateral, nil

	} else {
		// open only or close + open
		// when open, deposit mark * | openPosition | / lev
		if trader.TargetLeverage.LessThanOrEqual(_0) {
			return _0, fmt.Errorf("target leverage <= 0")
		}
		openPositionMargin := open.Abs().Mul(perpetual.MarkPrice).Div(trader.TargetLeverage).Add(totalFee)
		if position2.Sub(deltaPosition).IsZero() || !close.IsZero() {
			// strategy: let new margin balance = openPositionMargin
			adjustCollateral = openPositionMargin.Sub(a.MarginBalance)
			return adjustCollateral, nil
		} else {
			// strategy: always append positionMargin of openPosition
			return openPositionMargin, nil
		}
	}
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

func ComputeFee(hasOpened bool, price, amount, feeRate decimal.Decimal, afterTrade *model.AccountComputed) (decimal.Decimal, error) {
	if price.LessThanOrEqual(_0) || amount.IsZero() {
		return _0, fmt.Errorf("invalid price or admount. price:%s amount: %s", price, amount)
	}
	totalFee := price.Mul(amount.Abs()).Mul(feeRate)
	if !hasOpened {
		availableMargin := afterTrade.AvailableMargin
		if availableMargin.LessThanOrEqual(_0) {
			totalFee = _0
		} else if totalFee.GreaterThan(availableMargin) {
			// make sure the sum of fees < available margin
			totalFee = availableMargin
		}
	}
	return totalFee, nil
}
