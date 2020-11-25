package mai3

import (
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/shopspring/decimal"
)

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
