package mai3

import (
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
)

func ComputeAMMAmountWithPrice(g *model.GovParams, p *model.PerpetualStorage, amm *model.AccountStorage, isBuy bool, limitPrice decimal.Decimal) decimal.Decimal {
	// shift by spread
	if isBuy {
		limitPrice = limitPrice.Div(_1.Add(g.HalfSpreadRate))
	} else {
		limitPrice = limitPrice.Div(_1.Sub(g.HalfSpreadRate))
	}

	ammComputed := ComputeAccount(g, p, amm)
	ammContext := initAMMTradingContext(g, p, ammComputed, amm)
	if ammContext.Pos1.LessThanOrEqual(_0) && isBuy {
		return ComputeAMMAmountShortOpen(ammContext, limitPrice, g).Neg()
	} else if ammContext.Pos1.LessThan(_0) && !isBuy {
		return ComputeAMMAmountShortClose(ammContext, limitPrice, g).Neg()
	} else if ammContext.Pos1.GreaterThanOrEqual(_0) && !isBuy {
		return ComputeAMMAmountLongOpen(ammContext, limitPrice, g).Neg()
	} else if ammContext.Pos1.GreaterThan(_0) && isBuy {
		return ComputeAMMAmountLongClose(ammContext, limitPrice, g).Neg()
	}

	logger.Errorf("bug: unknown trading direction")
	return _0
}

func copyAMMTradingContext(ammContext *model.AMMTradingContext) *model.AMMTradingContext {
	return &model.AMMTradingContext{
		Index:         ammContext.Index,
		Lev:           ammContext.Lev,
		Cash:          ammContext.Cash,
		Pos1:          ammContext.Pos1,
		IsSafe:        ammContext.IsSafe,
		M0:            ammContext.M0,
		Mv:            ammContext.Mv,
		Ma1:           ammContext.Ma1,
		DeltaMargin:   ammContext.DeltaMargin,
		DeltaPosition: ammContext.DeltaPosition,
	}
}

func ComputeAMMAmountShortOpen(ammContext *model.AMMTradingContext, limitPrice decimal.Decimal, g *model.GovParams) decimal.Decimal {
	if !isAmmSafe(ammContext, g.Beta1) {
		return _0
	}
	computeM0(ammContext, g.Beta1)
	context := copyAMMTradingContext(ammContext)
	safePos2 := computeAMMSafeShortPositionAmount(ammContext, g.Beta1)
	if safePos2.GreaterThan(_0) || safePos2.GreaterThan(ammContext.Pos1) {
		return _0
	}
	maxAmount := safePos2.Sub(ammContext.Pos1)
	if maxAmount.GreaterThanOrEqual(_0) {
		logger.Errorf("warn: short open, but pos1 %s < safePos2 %s", ammContext.Pos1, safePos2)
		return _0
	}
	computeAMMInternalOpen(ammContext, maxAmount, g)
	if !maxAmount.Equal(ammContext.DeltaPosition.Sub(context.DeltaPosition)) {
		logger.Errorf("open position failed")
		return _0
	}

	safePos2Price := ammContext.DeltaMargin.Div(context.DeltaPosition).Abs()
	if safePos2Price.LessThanOrEqual(limitPrice) {
		return maxAmount
	}

	amount := ComputeAMMShortInverseVWAP(context, limitPrice, g.Beta1, false)
	if amount.GreaterThanOrEqual(_0) {
		return _0
	}
	return amount
}

func ComputeAMMAmountShortClose(ammContext *model.AMMTradingContext, limitPrice decimal.Decimal, g *model.GovParams) decimal.Decimal {
	if !ammContext.DeltaMargin.IsZero() || !ammContext.DeltaPosition.IsZero() {
		return _0
	}
	zeroContext := copyAMMTradingContext(ammContext)
	if !ammContext.Pos1.IsZero() {
		computeAMMInternalClose(zeroContext, ammContext.Pos1.Neg(), g)
		if zeroContext.DeltaPosition.IsZero() {
			logger.Errorf("close to zero failed")
			return _0
		}
		zeroPrice := zeroContext.DeltaMargin.Div(zeroContext.DeltaPosition).Abs()
		if zeroPrice.GreaterThanOrEqual(limitPrice) {
			ammContext = zeroContext
		} else if !isAmmSafe(ammContext, g.Beta2) {
			return _0
		} else {
			computeM0(ammContext, g.Beta2)
			amount := ComputeAMMShortInverseVWAP(ammContext, limitPrice, g.Beta2, true)
			if amount.GreaterThan(_0) {
				computeAMMInternalClose(ammContext, amount, g)
			}
		}
	}
	if ammContext.Pos1.GreaterThanOrEqual(_0) {
		openAmount := ComputeAMMAmountLongOpen(ammContext, limitPrice, g)
		return ammContext.DeltaPosition.Add(openAmount)
	}
	return ammContext.DeltaPosition
}

func ComputeAMMAmountLongOpen(ammContext *model.AMMTradingContext, limitPrice decimal.Decimal, g *model.GovParams) decimal.Decimal {
	if !isAmmSafe(ammContext, g.Beta1) {
		return _0
	}
	computeM0(ammContext, g.Beta1)
	safePos2 := computeAMMSafeLongPositionAmount(ammContext, g.Beta1)
	if safePos2.LessThan(_0) || safePos2.LessThan(ammContext.Pos1) {
		return _0
	}
	maxAmount := safePos2.Sub(ammContext.Pos1)
	if maxAmount.LessThanOrEqual(_0) {
		return _0
	}
	context := copyAMMTradingContext(ammContext)
	computeAMMInternalOpen(ammContext, maxAmount, g)
	if !maxAmount.Equal(ammContext.DeltaPosition.Sub(context.DeltaPosition)) {
		return _0
	}
	safePos2Price := ammContext.DeltaMargin.Div(ammContext.DeltaPosition).Abs()
	if safePos2Price.GreaterThanOrEqual(limitPrice) {
		return maxAmount
	}
	amount := ComputeAMMLongInverseVWAP(context, limitPrice, g.Beta1, false)
	if amount.LessThanOrEqual(_0) {
		return _0
	}
	return amount
}

func ComputeAMMAmountLongClose(ammContext *model.AMMTradingContext, limitPrice decimal.Decimal, g *model.GovParams) decimal.Decimal {
	if !ammContext.DeltaMargin.IsZero() || !ammContext.DeltaPosition.IsZero() {
		return _0
	}

	zeroContext := copyAMMTradingContext(ammContext)
	if !ammContext.Pos1.IsZero() {
		computeAMMInternalClose(zeroContext, zeroContext.Pos1.Neg(), g)
		if zeroContext.DeltaPosition.IsZero() {
			return _0
		}
		zeroPrice := zeroContext.DeltaMargin.Div(zeroContext.DeltaPosition).Abs()
		if zeroPrice.LessThanOrEqual(limitPrice) {
			ammContext = zeroContext
		} else if !isAmmSafe(ammContext, g.Beta2) {
			return _0
		} else {
			computeM0(ammContext, g.Beta2)
			amount := ComputeAMMLongInverseVWAP(ammContext, limitPrice, g.Beta2, true)
			if amount.LessThan(_0) {
				computeAMMInternalClose(ammContext, amount, g)
			}
		}
	}
	if ammContext.Pos1.LessThanOrEqual(_0) {
		openAmount := ComputeAMMAmountShortOpen(ammContext, limitPrice, g)
		return ammContext.DeltaPosition.Add(openAmount)
	}
	return ammContext.DeltaPosition
}

func ComputeAMMShortInverseVWAP(ammContext *model.AMMTradingContext, price, beta decimal.Decimal, isClosing bool) decimal.Decimal {
	if !ammContext.IsSafe {
		logger.Errorf("bug: do not call computeAMMShortInverseVWAP when unsafe")
		return _0
	}
	index := ammContext.Index
	pos1 := ammContext.Pos1
	m0 := ammContext.M0
	previousMa1MinusMa2 := ammContext.DeltaMargin.Neg()
	previousAmount := ammContext.DeltaPosition
	/*
	  D = previousMa1MinusMa2 - previousAmount price;
	  E = beta - 1;
	  F = m0 + i pos1;
	  A = i F (i E + price);
	  B = i^3 E pos1^2 + m0^2 price +
	    i^2 pos1 (-D + 2 E m0 + pos1 price) -
	    i m0 (D + m0 - 2 pos1 price);
	  C = F^2 D;
	  sols = 1/(2 A) (-B - sqrt(B^2 + 4 A C));
	*/
	d := previousMa1MinusMa2.Sub(previousAmount.Mul(price))
	e := beta.Sub(_1)
	f := index.Mul(pos1).Add(m0)
	a := index.Mul(e).Add(price).Mul(f).Mul(index)
	denominator := a.Mul(_2)
	if denominator.IsZero() {
		g := index.Mul(e).Mul(previousAmount).Add(previousMa1MinusMa2)
		denominator2 := beta.Mul(m0).Mul(m0).Add(f.Mul(g))
		if denominator2.IsZero() {
			return _0
		}
		amount := f.Mul(f).Mul(g).Div(index).Div(denominator2)
		return amount
	}
	b := index.Mul(index).Mul(index).Mul(e).Mul(pos1).Mul(pos1)
	b = b.Add(m0.Mul(m0).Mul(price))
	b = b.Add(pos1.Mul(price).Add(e.Mul(m0).Mul(_2)).Sub(d).Mul(pos1).Mul(index).Mul(index))
	b = b.Sub(d.Add(m0).Sub(pos1.Mul(price).Mul(_2)).Mul(m0).Mul(index))
	c := f.Mul(f).Mul(d)
	beforeSqrt := a.Mul(c).Mul(_4).Add(b.Mul(b))
	if beforeSqrt.LessThan(_0) {
		logger.Warnf("computeAMMShortInverseVWAP Δ < 0 ")
		return _0
	}
	numerator := Sqrt(beforeSqrt)
	if isClosing {
		numerator = numerator.Neg()
	}
	numerator = numerator.Add(b).Neg()
	amount := numerator.Div(denominator)
	return amount
}

// the inverse function of VWAP when AMM holds long
// call computeM0 before this function
// the returned amount(= pos2 - pos1) is the amm's perspective
func ComputeAMMLongInverseVWAP(ammContext *model.AMMTradingContext, price, beta decimal.Decimal, isClosing bool) decimal.Decimal {
	if !ammContext.IsSafe {
		logger.Errorf("bug: do not call ComputeAMMLongInverseVWAP when unsafe")
		return _0
	}
	index := ammContext.Index
	ma1 := ammContext.Ma1
	m0 := ammContext.M0
	previousMa1MinusMa2 := ammContext.DeltaMargin.Neg()
	previousAmount := ammContext.DeltaPosition
	/*
	  D = previousMa1MinusMa2 - previousAmount price;
	  A = ma1 price (i + (-1 + beta) price);
	  B = i ma1 (D + ma1) - beta m0^2 price + (-1 + beta) ma1 (2 D + ma1) price;
	  C = D (ma1 (ma1 + D) + beta (m0^2 - ma1 (ma1 + D)));
	  sols = 1/(2 A) (B + sqrt(B^2 + 4 A C));
	*/
	d := previousMa1MinusMa2.Sub(previousAmount.Mul(price))
	a := beta.Sub(_1).Mul(price).Add(index).Mul(ma1).Mul(price)
	b := ma1.Add(d).Mul(index).Mul(ma1)
	b = b.Sub(beta.Mul(m0).Mul(m0).Mul(price))
	b = b.Add(beta.Sub(_1).Mul(ma1).Mul(_2.Mul(d).Add(ma1)).Mul(price))
	c := ma1.Add(d).Mul(ma1).Neg().Add(m0.Mul(m0)).Mul(beta)
	c = c.Add(ma1.Add(d).Mul(ma1))
	c = c.Mul(d)
	denominator := a.Mul(_2)
	if denominator.IsZero() {
		// G = i previousAmount + previousMa1MinusMa2 (beta - 1);
		// H = beta m0^2 - ma1 G;
		// sols = kappaG (-H + ma1^2 (k - 1))/i/H
		g := index.Mul(previousAmount).Add(beta.Sub(_1).Mul(previousMa1MinusMa2))
		h := beta.Mul(m0).Mul(m0).Sub(ma1.Mul(g))
		if h.IsZero() {
			logger.Warnf("warn: computeAMMLongInverseVWAP denominator2 = 0")
			return _0
		}
		amount := beta.Sub(_1).Mul(ma1).Mul(ma1).Sub(h).Mul(g).Div(index).Div(h)
		return amount
	}
	beforeSqrt := a.Mul(c).Mul(_4).Add(b.Mul(b))
	if beforeSqrt.LessThan(_0) {
		logger.Warnf("warn: computeAMMLongInverseVWAP Δ < 0")
		return _0
	}
	numerator := Sqrt(beforeSqrt)
	if isClosing {
		numerator = numerator.Neg()
	}
	numerator = numerator.Add(b)
	return numerator.Div(denominator)
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
		return _0
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
		return _0
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
		return _0
	}
	// safePosition = -m0 / i / (1 + sqrt(beta * lev))
	beforeSqrt := beta.Mul(context.Lev)
	if beforeSqrt.LessThan(_0) {
		return _0
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
		return _0
	}
	mv := context.Cash.Mul(context.Lev.Sub(_1))
	return mv
}

func computeM0Short(context *model.AMMTradingContext, beta decimal.Decimal) decimal.Decimal {
	if context.Pos1.GreaterThan(_0) {
		return _0
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
		return _0
	}
	afterSqrt := Sqrt(beforeSqrt)
	mv := context.Lev.Sub(_1).Div(_2).Div(context.Lev)
	mv = mv.Mul(b.Sub(a).Add(afterSqrt))
	return mv
}

func computeM0Long(context *model.AMMTradingContext, beta decimal.Decimal) decimal.Decimal {
	if context.Pos1.LessThan(_0) {
		return _0
	}
	// b = lev * cash + index * pos1 * (lev - 1)
	// before_sqrt = b ** 2 + 4 * beta * index * lev * cash * pos1
	// v = (lev - 1) / 2 / (lev + beta - 1) * (b - 2 * (1 - beta) * cash + math.sqrt(before_sqrt))
	b := context.Lev.Mul(context.Cash).Add(context.Index.Mul(context.Pos1).Mul(context.Lev.Sub(_1)))
	beforeSqrt := b.Mul(b)
	beforeSqrt = beforeSqrt.Add(_4.Mul(beta).Mul(context.Index).Mul(context.Lev).Mul(context.Cash).Mul(context.Pos1))
	if beforeSqrt.LessThan(_0) {
		return _0
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
		return _0
	}
}

func computeDeltaMarginLong(context *model.AMMTradingContext, beta, pos2 decimal.Decimal) decimal.Decimal {
	if context.Pos1.LessThan(_0) {
		return _0
	}
	if pos2.LessThan(_0) {
		return _0
	}
	if context.M0.LessThanOrEqual(_0) {
		return _0
	}
	if context.Ma1.LessThanOrEqual(_0) {
		return _0
	}
	// a = 2 * (1 - beta) * ma1
	// assert a != 0
	// b = -beta * m0 ** 2 + ma1 * (ma1 * (1 - beta) - index * (pos2 - pos1))
	// before_sqrt = b**2 + 2 * a * ma1 * m0 ** 2 * beta
	// assert before_sqrt >= 0
	// ma2 = (b + math.sqrt(before_sqrt)) / a
	a := _1.Sub(beta).Mul(context.Ma1).Mul(_2)
	if a.IsZero() {
		return _0
	}
	b := pos2.Sub(context.Pos1).Mul(context.Index)
	b = a.Div(_2).Sub(b).Mul(context.Ma1)
	b = b.Sub(beta.Mul(context.M0).Mul(context.M0))
	beforeSqrt := beta.Mul(a).Mul(context.Ma1).Mul(context.M0).Mul(context.M0).Mul(_2)
	beforeSqrt = beforeSqrt.Add(b.Mul(b))
	if beforeSqrt.LessThan(_0) {
		return _0
	}
	ma2 := Sqrt(beforeSqrt).Add(b).Div(a)
	return ma2.Sub(context.Ma1)
}

func computeDeltaMarginShort(context *model.AMMTradingContext, beta, pos2 decimal.Decimal) decimal.Decimal {
	if context.Pos1.GreaterThan(_0) {
		return _0
	}
	if pos2.GreaterThan(_0) {
		return _0
	}
	if context.M0.LessThanOrEqual(_0) {
		return _0
	}
	// ma2 - ma1 = index * (pos1 - pos2) * (1 - beta + beta * m0**2 / (m0 + pos1 * index) / (m0 + pos2 * index))
	deltaMargin := beta.Mul(context.M0).Mul(context.M0)
	deltaMargin = deltaMargin.Div(context.Pos1.Mul(context.Index).Mul(context.M0))
	deltaMargin = deltaMargin.Div(pos2.Mul(context.Index).Mul(context.M0))
	deltaMargin = deltaMargin.Add(_1).Sub(beta)
	deltaMargin = deltaMargin.Mul(context.Index).Mul(context.Pos1.Sub(pos2))
	return deltaMargin
}
