package mai3

import (
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
)

func ComputeAMMTrade(g *model.GovParams, p *model.PerpetualStorage, trader, amm *model.AccountStorage, amount decimal.Decimal) (lpFee, vaultFee, operatorFee, tradingPirce decimal.Decimal, err error) {
	if amount.IsZero() {
		return
	}

	// amm
	deltaAMMAmount, _, tradingPrice, err := ComputeAMMPrice(g, p, amm, amount)
	if err != nil {
		logger.Errorf("ComputeAMMPrice error:%s", err)
		return
	}
	if !deltaAMMAmount.Neg().Equal(amount) {
		logger.Errorf("trading amount mismatched %s != %s", deltaAMMAmount, amount)
		return
	}
	lpFee, err = ComputeFee(tradingPrice, deltaAMMAmount, g.LpFeeRate)
	if err != nil {
		return
	}
	vaultFee, err = ComputeFee(tradingPrice, deltaAMMAmount, g.VaultFeeRate)
	if err != nil {
		return
	}
	operatorFee, err = ComputeFee(tradingPrice, deltaAMMAmount, g.OperatorFeeRate)
	if err != nil {
		return
	}
	amm.CashBalance = amm.CashBalance.Add(lpFee)

	// trader
	err = ComputeTradeWithPrice(p, trader, tradingPrice, deltaAMMAmount.Neg(), g.LpFeeRate.Add(g.VaultFeeRate).Add(g.OperatorFeeRate))
	return
}

func ComputeAMMPrice(g *model.GovParams, p *model.PerpetualStorage, amm *model.AccountStorage, amount decimal.Decimal) (deltaAMMAmount, deltaAMMMargin, tradingPrice decimal.Decimal, err error) {
	if amount.IsZero() {
		return
	}
	ammComputed := ComputeAccount(g, p, amm)
	ammTrading, err := ComputeAMMInternalTrade(g, p, ammComputed, amm, amount.Neg())
	if err != nil {
		return _0, _0, _0, err
	}
	deltaAMMMargin = ammTrading.DeltaMargin
	deltaAMMAmount = ammTrading.DeltaPosition
	tradingPrice = deltaAMMMargin.Div(deltaAMMAmount).Abs()
	ComputeTradeWithPrice(p, amm, tradingPrice, amount.Neg(), _0)
	return
}

func ComputeTradeWithPrice(p *model.PerpetualStorage, a *model.AccountStorage, price, amount, feeRate decimal.Decimal) error {
	close, open := utils.SplitAmount(a.PositionAmount, amount)
	if !close.IsZero() {
		if err := ComputeDecreasePosition(p, a, price, amount); err != nil {
			return err
		}
	}

	if !open.IsZero() {
		if err := ComputeIncreasePosition(p, a, price, amount); err != nil {
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

func ComputeDecreasePosition(p *model.PerpetualStorage, a *model.AccountStorage, price, amount decimal.Decimal) error {
	if a.PositionAmount.IsZero() || amount.IsZero() || utils.HasTheSameSign(a.PositionAmount, amount) {
		return fmt.Errorf("invalid amount or position, position:%s, amount:%s", a.PositionAmount, amount)
	}

	if price.LessThanOrEqual(_0) {
		return fmt.Errorf("invalid price %s", price)
	}

	if a.PositionAmount.Abs().LessThan(amount.Abs()) {
		return fmt.Errorf("position size is less than amount. position:%s, amount:%s", a.PositionAmount, amount)
	}
	oldPosition := a.PositionAmount
	fundingLoss := p.AccumulatedFundingPerContract.Sub(a.EntryFundingLoss.Div(a.PositionAmount)).Mul(amount.Neg())
	a.CashBalance = a.CashBalance.Sub(price.Mul(amount)).Sub(fundingLoss)
	a.PositionAmount = a.PositionAmount.Add(amount)
	a.EntryFundingLoss = a.EntryFundingLoss.Mul(a.PositionAmount).Div(oldPosition)
	return nil
}

func ComputeIncreasePosition(p *model.PerpetualStorage, a *model.AccountStorage, price, amount decimal.Decimal) error {
	if price.LessThanOrEqual(_0) {
		return fmt.Errorf("invalid price %s", price)
	}
	if amount.IsZero() {
		return fmt.Errorf("invalid amount %s", amount)
	}
	if !a.PositionAmount.IsZero() && !utils.HasTheSameSign(a.PositionAmount, amount) {
		return fmt.Errorf("invalid amount or position, position:%s, amount:%s", a.PositionAmount, amount)
	}
	a.CashBalance = a.CashBalance.Sub(price.Mul(amount))
	a.PositionAmount = a.PositionAmount.Add(amount)
	a.EntryFundingLoss = a.EntryFundingLoss.Add(p.AccumulatedFundingPerContract.Mul(amount))
	return nil
}

func ComputeFee(price, amount, feeRate decimal.Decimal) (decimal.Decimal, error) {
	if price.LessThanOrEqual(_0) || amount.IsZero() {
		return _0, fmt.Errorf("invalid price or admount. price:%s amount: %s", price, amount)
	}
	return price.Mul(amount.Abs()).Mul(feeRate), nil
}
