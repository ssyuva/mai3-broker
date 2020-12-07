package mai3

import (
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
	"math"
)

func ComputeMaxTradeAmountWithPrice(g *model.GovParams, p *model.PerpetualStorage, a *model.AccountStorage,
	price, leverage, feeRate decimal.Decimal, isBuy bool) decimal.Decimal {
	closeAmount := a.PositionAmount.Neg()
	if !closeAmount.IsZero() {
		ComputeTradeWithPrice(p, a, price, closeAmount, feeRate)
	}
	accountComputed := ComputeAccount(g, p, a)
	cash := accountComputed.AvailableMargin
	denominator := leverage.Mul(feeRate).Add(_1).Mul(price)
	if denominator.IsZero() {
		return _0
	}
	openAmount := cash.Mul(leverage).Div(denominator)
	if !isBuy {
		openAmount = openAmount.Neg()
	}
	res := closeAmount.Add(openAmount)
	if isBuy && res.LessThanOrEqual(_0) {
		return _0
	} else if !isBuy && res.GreaterThan(_0) {
		return _0
	}
	return res
}

func ComputeAMMMaxTradeAmount(g *model.GovParams, p *model.PerpetualStorage, trader, amm *model.AccountStorage, leverage decimal.Decimal, isBuy bool) decimal.Decimal {
	ammComputed := ComputeAccount(g, p, amm)
	ammContext := initAMMTradingContext(g, p, ammComputed, amm)
	if !isAmmSafe(ammContext, g.Beta1) {
		if isBuy && amm.PositionAmount.LessThan(_0) {
			return _0
		}
		if !isBuy && amm.PositionAmount.GreaterThan(_0) {
			return _0
		}
	}
	traderComputed := ComputeAccount(g, p, trader)
	guess := traderComputed.MarginBalance.Mul(leverage).Div(p.IndexPrice)
	checkTrading := func(a float64) float64 {
		if a == 0 {
			return 0
		}
		_, _, _, _, err := ComputeAMMTrade(g, p, trader, amm, decimal.NewFromFloat(a))
		if err != nil {
			return math.Abs(a)
		}
		traderComputed := ComputeAccount(g, p, trader)
		if traderComputed.Leverage.GreaterThan(leverage) {
			return math.Abs(a)
		}
		return -math.Abs(a)
	}
	var bound0, bound1 float64
	if isBuy {
		bound1, _ = guess.Float64()
	} else {
		bound0, _ = guess.Neg().Float64()
	}
	min, max := Gss(checkTrading, bound0, bound1, 1e-8, 15)
	amount := decimal.NewFromInt(int64((min + max) / 2))
	return amount
}

func initAMMTradingContext(g *model.GovParams, p *model.PerpetualStorage, ammComputed *model.AccountComputed, amm *model.AccountStorage) *model.AMMTradingContext {
	pos1 := amm.PositionAmount
	cash := ammComputed.AvailableCashBalance
	index := p.IndexPrice
	lev := g.TargetLeverage
	return &model.AMMTradingContext{
		Index:         index,
		Lev:           lev,
		Cash:          cash,
		Pos1:          pos1,
		IsSafe:        true,
		M0:            _0,
		Mv:            _0,
		Ma1:           _0,
		DeltaMargin:   _0,
		DeltaPosition: _0,
	}
}

func isAmmSafe(context *model.AMMTradingContext, beta decimal.Decimal) bool {
	if context.Pos1.IsZero() {
		return true
	} else if context.Pos1.LessThan(_0) {
		return isAMMSafeShort(context, beta)
	}

	return isAMMSafeLong(context, beta)
}

func isAMMSafeShort(context *model.AMMTradingContext, beta decimal.Decimal) bool {
	if context.Pos1.GreaterThan(_0) {
		logger.Errorf("pos1 %s > 0 ", context.Pos1)
		return false
	}
	beforeSqrt := beta.Mul(context.Lev)
	if beforeSqrt.LessThan(_0) {
		logger.Errorf("ammSafe sqrt < 0")
		return false
	}
	denominator := context.Lev.Add(_1).Add(_2.Mul(Sqrt(beforeSqrt)))
	safeIndex := context.Lev.Neg().Mul(context.Cash).Div(context.Pos1).Div(denominator)
	return context.Index.LessThanOrEqual(safeIndex)
}

func isAMMSafeLong(context *model.AMMTradingContext, beta decimal.Decimal) bool {
	if context.Pos1.LessThan(_0) {
		logger.Errorf("pos1 %s < 0 ", context.Pos1)
		return false
	}
	if context.Cash.GreaterThanOrEqual(_0) {
		return true
	}
	levMinus1 := context.Lev.Sub(_1)
	beforeSqrt := beta.Mul(levMinus1.Add(beta))
	if beforeSqrt.LessThan(_0) {
		logger.Errorf("ammSafe sqrt < 0")
		return false
	}
	safeIndex := context.Lev.Neg().Mul(context.Cash)
	safeIndex = safeIndex.Mul(levMinus1.Add(_2.Mul(beta.Add(Sqrt(beforeSqrt)))))
	safeIndex = safeIndex.Div(context.Pos1).Div(levMinus1).Div(levMinus1)
	return context.Index.GreaterThanOrEqual(safeIndex)
}

func ComputeAccount(g *model.GovParams, p *model.PerpetualStorage, a *model.AccountStorage) *model.AccountComputed {
	return ComputeAccountWithMarkPrice(g, p, a, p.MarkPrice)
}

func ComputeAccountWithMarkPrice(g *model.GovParams, p *model.PerpetualStorage, a *model.AccountStorage, markPrice decimal.Decimal) *model.AccountComputed {
	positionValue := markPrice.Mul(a.PositionAmount.Abs())
	positionMargin := markPrice.Mul(a.PositionAmount.Abs()).Mul(g.InitialMarginRate)
	maintenanceMargin := markPrice.Mul(a.PositionAmount.Abs()).Mul(g.MaintenanceMarginRate)
	fundingLoss := p.AccumulatedFundingPerContract.Mul(a.PositionAmount).Sub(a.EntryFundingLoss)
	reservedCash := decimal.Zero
	if !a.PositionAmount.IsZero() {
		reservedCash = g.KeeperGasReward
	}
	availableCashBalance := a.CashBalance.Sub(fundingLoss)
	marginBalance := availableCashBalance.Add(markPrice.Mul(a.PositionAmount))
	maxWithdrawable := decimal.Max(decimal.Zero, marginBalance.Sub(positionMargin).Sub(reservedCash))
	availableMargin := decimal.Max(decimal.Zero, maxWithdrawable)
	withdrawableBalance := maxWithdrawable
	isSafe := maintenanceMargin.LessThanOrEqual(marginBalance)
	leverage := decimal.Zero
	if marginBalance.GreaterThan(decimal.Zero) {
		leverage = positionValue.Div(marginBalance)
	}
	return &model.AccountComputed{
		PositionValue:        positionValue,
		PositionMargin:       positionMargin,
		Leverage:             leverage,
		MaintenanceMargin:    maintenanceMargin,
		FundingLoss:          fundingLoss,
		AvailableCashBalance: availableCashBalance,
		AvailableMargin:      availableMargin,
		MarginBalance:        marginBalance,
		MaxWithdrawable:      maxWithdrawable,
		WithdrawableBalance:  withdrawableBalance,
		IsSafe:               isSafe,
	}
}

func ComputeAMMInternalTrade(g *model.GovParams, p *model.PerpetualStorage, a *model.AccountComputed, amm *model.AccountStorage, amount decimal.Decimal) (*model.AMMTradingContext, error) {
	context := initAMMTradingContext(g, p, a, amm)
	close, open := utils.SplitAmount(context.Pos1, amount)
	if close.IsZero() && open.IsZero() {
		return nil, fmt.Errorf("amm trade: trading amount = 0")
	}
	if !close.IsZero() {
		computeAMMInternalClose(context, amount, g)
	}
	if !open.IsZero() {
		computeAMMInternalOpen(context, amount, g)
	}
	if amount.LessThan(_0) {
		context.DeltaMargin = context.DeltaMargin.Mul(_1.Add(g.HalfSpreadRate))
	} else {
		context.DeltaMargin = context.DeltaMargin.Mul(_1.Sub(g.HalfSpreadRate))
	}
	return context, nil
}

func computeAMMInternalClose(context *model.AMMTradingContext, amount decimal.Decimal, g *model.GovParams) {
	beta := g.Beta2
	pos2 := context.Pos1.Add(amount)
	deltaMargin := _0
	if isAmmSafe(context, beta) {
		computeM0(context, beta)
		deltaMargin = computeDeltaMargin(context, beta, pos2)
	} else {
		deltaMargin = context.Index.Mul(amount).Neg()
	}
	context.DeltaMargin = context.DeltaMargin.Add(deltaMargin)
	context.DeltaPosition = context.DeltaPosition.Add(amount)
	context.Cash = context.Cash.Add(deltaMargin)
	context.Pos1 = pos2
}

func computeAMMInternalOpen(context *model.AMMTradingContext, amount decimal.Decimal, g *model.GovParams) {
	beta := g.Beta1
	pos2 := context.Pos1.Add(amount)
	deltaMargin := _0
	if !isAmmSafe(context, beta) {
		panic("amm can not open position anymore: unsafe before trade")
	}
	computeM0(context, beta)
	if amount.GreaterThan(_0) {
		safePos2 := computeAMMSafeLongPositionAmount(context, g.Beta1)
		if pos2.GreaterThan(safePos2) {
			panic("amm can not open position anymore: position too large after trade")
		}
	} else {
		safePos2 := computeAMMSafeShortPositionAmount(context, g.Beta1)
		if pos2.LessThan(safePos2) {
			panic("amm can not open position anymore: position too large after trade")
		}
	}
	deltaMargin = computeDeltaMargin(context, beta, pos2)
	context.DeltaMargin = context.DeltaMargin.Add(deltaMargin)
	context.DeltaPosition = context.DeltaPosition.Add(amount)
	context.Cash = context.Cash.Add(deltaMargin)
	context.Pos1 = pos2
}

func computeAMMSafeLongPositionAmount(context *model.AMMTradingContext, beta decimal.Decimal) decimal.Decimal {
	if !context.IsSafe {
		panic("bug: do not call shortPosition when unsafe")
	}
	safePosition := _0
	edge1 := beta.Mul(context.Lev).Add(beta).Sub(_1)
	// if -1 + beta + beta lev = 0
	// safePosition = m0 / 2 / i / (1 - 2 beta)
	if edge1.IsZero() {
		safePosition = context.M0.Div(_2).Div(context.Index).Div(_1.Sub(_2.Mul(beta)))
		return safePosition
	}
	// a = (lev + beta - 1)
	//                    (2 * beta - 1) * a + sqrt(beta * a) * (lev + 2beta - 2)
	// safePosition = m0 -------------------------------------------------------------
	//                           i * (lev - 1) * (beta * lev + beta - 1)
	a := context.Lev.Add(beta).Sub(_1)
	beforeSqrt := beta.Mul(a)
	if beforeSqrt.LessThan(_0) {
		panic("bug: ammSafe sqrt < 0")
	}
	denominator := edge1.Mul(context.Lev.Sub(_1)).Mul(context.Index)
	safePosition = _2.Mul(beta).Sub(_2).Add(context.Lev)
	safePosition = safePosition.Mul(Sqrt(beforeSqrt))
	safePosition = safePosition.Add(_2.Mul(beta).Sub(_1).Mul(a))
	safePosition = safePosition.Mul(context.M0).Div(denominator)
	return safePosition
}

func computeAMMSafeShortPositionAmount(context *model.AMMTradingContext, beta decimal.Decimal) decimal.Decimal {
	if !context.IsSafe {
		panic("bug: do not call shortPosition when unsafe")
	}
	// safePosition = -m0 / i / (1 + sqrt(beta * lev))
	beforeSqrt := beta.Mul(context.Lev)
	if beforeSqrt.LessThan(_0) {
		panic("bug: ammSafe sqrt < 0")
	}
	safePosition := context.M0.Neg().Div(context.Index).Div(Sqrt(beforeSqrt).Add(_1))
	return safePosition
}

func computeM0(context *model.AMMTradingContext, beta decimal.Decimal) {
	if !isAmmSafe(context, beta) {
		context.IsSafe = false
		return
	}
	mv := _0
	if context.Pos1.IsZero() {
		mv = computeM0Flat(context)
	} else if context.Pos1.GreaterThan(_0) {
		mv = computeM0Short(context, beta)
	} else {
		mv = computeM0Short(context, beta)
	}
	m0 := mv.Mul(context.Lev).Div(context.Lev.Sub(_1))
	ma1 := context.Cash.Add(mv)
	context.IsSafe = true
	context.Mv = mv
	context.M0 = m0
	context.Ma1 = ma1
}

func computeM0Flat(context *model.AMMTradingContext) decimal.Decimal {
	if !context.Pos1.IsZero() {
		panic("pos1 != 0")
	}
	mv := context.Cash.Mul(context.Lev.Sub(_1))
	return mv
}

func computeM0Short(context *model.AMMTradingContext, beta decimal.Decimal) decimal.Decimal {
	if context.Pos1.GreaterThan(_0) {
		panic("pos1 > 0")
	}
	// a = 2 * index * pos1
	// b = lev * cash + index * pos1 * (lev + 1)
	// before_sqrt = b ** 2 - beta * lev * a ** 2
	// v = (lev - 1) / 2 / lev * (b - a + math.sqrt(before_sqrt))
	a := _2.Mul(context.Index).Mul(context.Pos1)
	b := context.Lev.Mul(context.Cash)
	b = b.Add(context.Index.Mul(context.Pos1).Mul(context.Lev.Add(_1)))
	beforeSqrt := b.Mul(b)
	beforeSqrt = beforeSqrt.Sub(beta.Mul(context.Lev).Mul(a).Mul(a))
	if beforeSqrt.LessThan(_0) {
		panic("edge case: short m0 sqrt < 0")
	}
	afterSqrt := Sqrt(beforeSqrt)
	mv := context.Lev.Sub(_1).Div(_2).Div(context.Lev)
	mv = mv.Mul(b.Sub(a).Add(afterSqrt))
	return mv
}

func computeM0Long(context *model.AMMTradingContext, beta decimal.Decimal) decimal.Decimal {
	if context.Pos1.LessThan(_0) {
		panic("pos1 < 0")
	}
	// b = lev * cash + index * pos1 * (lev - 1)
	// before_sqrt = b ** 2 + 4 * beta * index * lev * cash * pos1
	// v = (lev - 1) / 2 / (lev + beta - 1) * (b - 2 * (1 - beta) * cash + math.sqrt(before_sqrt))
	b := context.Lev.Mul(context.Cash).Add(context.Index.Mul(context.Pos1).Mul(context.Lev.Sub(_1)))
	beforeSqrt := b.Mul(b)
	beforeSqrt = beforeSqrt.Add(_4.Mul(beta).Mul(context.Index).Mul(context.Lev).Mul(context.Cash).Mul(context.Pos1))
	if beforeSqrt.LessThan(_0) {
		panic("edge case: short m0 sqrt < 0")
	}
	afterSqrt := Sqrt(beforeSqrt)
	mv := context.Lev.Sub(_1).Div(_2).Div(context.Lev.Add(beta).Sub(_1))
	mv = mv.Mul(b.Sub(_2.Mul(_1.Sub(beta).Mul(context.Cash))).Add(afterSqrt))
	return mv
}

func computeDeltaMargin(context *model.AMMTradingContext, beta, pos2 decimal.Decimal) decimal.Decimal {
	if context.Pos1.GreaterThanOrEqual(_0) && pos2.GreaterThanOrEqual(_0) {
		return computeDeltaMarginLong(context, beta, pos2)
	} else if context.Pos1.LessThanOrEqual(_0) && pos2.LessThanOrEqual(_0) {
		return computeDeltaMarginShort(context, beta, pos2)
	} else {
		panic("bug: cross direction is not supported")
	}
}

func computeDeltaMarginLong(context *model.AMMTradingContext, beta, pos2 decimal.Decimal) decimal.Decimal {
	if context.Pos1.LessThan(_0) {
		panic("pos1 < 0")
	}
	if pos2.LessThan(_0) {
		panic("pos2 < 0")
	}
	if context.M0.LessThanOrEqual(_0) {
		panic("m0 <= 0")
	}
	if context.Ma1.LessThanOrEqual(_0) {
		panic("ma1 <= 0")
	}
	// a = 2 * (1 - beta) * ma1
	// assert a != 0
	// b = -beta * m0 ** 2 + ma1 * (ma1 * (1 - beta) - index * (pos2 - pos1))
	// before_sqrt = b**2 + 2 * a * ma1 * m0 ** 2 * beta
	// assert before_sqrt >= 0
	// ma2 = (b + math.sqrt(before_sqrt)) / a
	a := _1.Sub(beta).Mul(context.Ma1).Mul(_2)
	if a.IsZero() {
		panic("edge case: deltaMarginLong.a = 0")
	}
	b := pos2.Sub(context.Pos1).Mul(context.Index)
	b = a.Div(_2).Sub(b).Mul(context.Ma1)
	b = b.Sub(beta.Mul(context.M0).Mul(context.M0))
	beforeSqrt := beta.Mul(a).Mul(context.Ma1).Mul(context.M0).Mul(context.M0).Mul(_2)
	beforeSqrt = beforeSqrt.Add(b.Mul(b))
	if beforeSqrt.LessThan(_0) {
		panic("edge case: deltaMarginLong.sqrt < 0")
	}
	ma2 := Sqrt(beforeSqrt).Add(b).Div(a)
	return ma2.Sub(context.Ma1)
}

func computeDeltaMarginShort(context *model.AMMTradingContext, beta, pos2 decimal.Decimal) decimal.Decimal {
	if context.Pos1.GreaterThan(_0) {
		panic("pos1 > 0")
	}
	if pos2.GreaterThan(_0) {
		panic("pos2 > 0")
	}
	if context.M0.LessThanOrEqual(_0) {
		panic("m0 <= 0")
	}
	// ma2 - ma1 = index * (pos1 - pos2) * (1 - beta + beta * m0**2 / (m0 + pos1 * index) / (m0 + pos2 * index))
	deltaMargin := beta.Mul(context.M0).Mul(context.M0)
	deltaMargin = deltaMargin.Div(context.Pos1.Mul(context.Index).Mul(context.M0))
	deltaMargin = deltaMargin.Div(pos2.Mul(context.Index).Mul(context.M0))
	deltaMargin = deltaMargin.Add(_1).Sub(beta)
	deltaMargin = deltaMargin.Mul(context.Index).Mul(context.Pos1.Sub(pos2))
	return deltaMargin
}
