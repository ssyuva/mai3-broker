package mai3

import (
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
)

func ComputeAMMTrade(p *model.LiquidityPoolStorage, perpetualIndex int64, trader *model.AccountStorage, amount decimal.Decimal) (decimal.Decimal, error) {
	if amount.IsZero() {
		return _0, fmt.Errorf("bad amount")
	}
	perpetual, ok := p.Perpetuals[perpetualIndex]
	if !ok {
		return _0, fmt.Errorf("perpetual %d not found in the pool", perpetualIndex)
	}

	// AMM
	_, deltaAMMAmount, tradingPrice, err := ComputeAMMPrice(p, perpetualIndex, amount)
	if err != nil {
		logger.Errorf("ComputeAMMPrice error:%s", err)
		return _0, err
	}
	if !deltaAMMAmount.Neg().Equal(amount) {
		logger.Errorf("trading amount mismatched %s != %s", deltaAMMAmount, amount)
		return _0, fmt.Errorf("trading amount mismatched")
	}

	// fee
	lpFee, err := ComputeFee(tradingPrice, deltaAMMAmount, perpetual.LpFeeRate)
	if err != nil {
		return _0, err
	}

	// trader
	if err = ComputeTradeWithPrice(p, perpetualIndex, trader, tradingPrice, deltaAMMAmount.Neg(),
		perpetual.LpFeeRate.Add(p.VaultFeeRate).Add(perpetual.OperatorFeeRate)); err != nil {
		logger.Errorf("ComputeTradeWithPrice trader err:%s", err)
		return _0, err
	}

	// new AMM
	fakeAMMAccount := &model.AccountStorage{
		CashBalance:    p.PoolCashBalance,
		PositionAmount: perpetual.AmmPositionAmount,
	}
	if err = ComputeTradeWithPrice(p, perpetualIndex, fakeAMMAccount, tradingPrice, deltaAMMAmount, _0); err != nil {
		logger.Errorf("ComputeTradeWithPrice fakeAMMAccount err:%s", err)
		return _0, err
	}
	fakeAMMAccount.CashBalance = fakeAMMAccount.CashBalance.Add(lpFee)
	p.PoolCashBalance = fakeAMMAccount.CashBalance
	perpetual.AmmPositionAmount = fakeAMMAccount.PositionAmount
	return tradingPrice, nil
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

func computeAMMInternalTrade(p *model.LiquidityPoolStorage, perpetualIndex int64, amount decimal.Decimal) (*model.AMMTradingContext, error) {
	context := initAMMTradingContext(p, perpetualIndex)
	close, open := utils.SplitAmount(context.Position1, amount)
	if close.IsZero() && open.IsZero() {
		return nil, fmt.Errorf("AMM trade: trading amount = 0")
	}
	var err error
	if !close.IsZero() {
		if context, err = computeAMMInternalClose(context, close); err != nil {
			return nil, err
		}
	}
	if !open.IsZero() {
		if context, err = computeAMMInternalOpen(context, open); err != nil {
			return nil, err
		}
	}
	if amount.LessThan(_0) {
		context.DeltaMargin = context.DeltaMargin.Mul(_1.Add(context.HalfSpread))
	} else {
		context.DeltaMargin = context.DeltaMargin.Mul(_1.Sub(context.HalfSpread))
	}
	return context, nil
}

func ComputeTradeWithPrice(p *model.LiquidityPoolStorage, perpetualIndex int64, a *model.AccountStorage, price, amount, feeRate decimal.Decimal) error {
	if price.LessThanOrEqual(_0) || amount.IsZero() {
		return fmt.Errorf("bad price %s or amount %s", price, amount)
	}

	close, open := utils.SplitAmount(a.PositionAmount, amount)
	if !close.IsZero() {
		if err := ComputeDecreasePosition(p, perpetualIndex, a, price, close); err != nil {
			return err
		}
	}

	if !open.IsZero() {
		if err := ComputeIncreasePosition(p, perpetualIndex, a, price, open); err != nil {
			return err
		}
	}

	fee, err := ComputeFee(price, amount, feeRate)
	if err != nil {
		return err
	}
	a.CashBalance = a.CashBalance.Sub(fee)
	return nil
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
