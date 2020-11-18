package mai3

import (
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/shopspring/decimal"
)

var UnSafeM0Buy = decimal.New(1, -7)
var UnSafeM0Sell = decimal.New(1, -3)

func CalculateM0(p *model.PerpetualStorage, cash, pos1, beta decimal.Decimal) decimal.Decimal {
	var v decimal.Decimal
	if pos1.Equal(decimal.Zero) {
		v = p.Cash.Mul(p.Leverage.Sub(_1))
	} else if pos1.LessThan(decimal.Zero) {
		a := _2.Mul(p.IndexPrice).Mul(pos1)
		b := p.Leverage.Mul(cash).Add(p.IndexPrice.Mul(pos1).Mul(p.Leverage.Add(_1)))
		beforeSqrt := b.Pow(_2).Sub(beta.Mul(p.Leverage).Mul(a.Pow(_2)))
		v = p.Leverage.Sub(_1).Div(_2).Div(p.Leverage).Mul(b.Sub(a).Add(Sqrt(beforeSqrt)))
	} else {
		b := p.Leverage.Mul(cash).Add(p.IndexPrice.Mul(pos1).Mul(p.Leverage.Sub(_1)))
		beforeSqrt := b.Pow(_2).Add(_4.Mul(beta).Mul(p.IndexPrice).Mul(p.Leverage).Mul(cash).Mul(pos1))
		v = p.Leverage.Sub(_1).Div(_2).Div(p.Leverage.Add(beta).Sub(_1)).
			Mul(b.Sub(_2.Mul(_1.Sub(beta)).Mul(cash)).Add(Sqrt(beforeSqrt)))
	}
	m0 := v.Mul(p.Leverage).Div(p.Leverage.Sub(_1))
	return m0
}

func calculateShortMa1MinusMa2(p *model.PerpetualStorage, m0, pos1, pos2, beta decimal.Decimal) decimal.Decimal {
	return p.IndexPrice.Mul(pos2.Sub(pos1)).
		Mul(_1.Sub(beta).Add(beta.Mul(m0.Pow(_2)).Div(m0.Add(pos1.Mul(p.IndexPrice))).Div(m0.Add(pos2.Mul(p.IndexPrice)))))
}

func calculateLongMa2(p *model.PerpetualStorage, m0, ma1, pos1, pos2, beta decimal.Decimal) (decimal.Decimal, error) {
	var res decimal.Decimal
	a := _2.Mul(_1.Sub(beta)).Mul(ma1)
	if a.Equal(decimal.Zero) {
		return res, fmt.Errorf("calculate long ma2 error")
	}

	b := beta.Neg().Mul(m0.Pow(_2)).Add(ma1.Mul(ma1.Mul(_1.Sub(beta).Sub(p.IndexPrice.Mul(pos2.Sub(pos1))))))
	beforeSqrt := b.Pow(_2).Add(_2.Mul(a).Mul(ma1).Mul(m0.Pow(_2)).Mul(beta))
	if beforeSqrt.LessThan(decimal.Zero) {
		return res, fmt.Errorf("calculate long ma2 error")
	}

	res = b.Add(Sqrt(beforeSqrt)).Div(a)
	return res, nil
}

func shortSafe(p *model.PerpetualStorage, index, cash, pos, beta decimal.Decimal) bool {
	newIndex := p.Leverage.Mul(cash).Div(pos.Neg()).Div(p.Leverage.Add(_1).Add(_2.Mul(Sqrt(beta.Mul(p.Leverage)))))
	return index.LessThanOrEqual(newIndex)
}

func longSafe(p *model.PerpetualStorage, index, cash, pos, beta decimal.Decimal) bool {
	if cash.GreaterThanOrEqual(decimal.Zero) {
		return true
	}
	newIndex := p.Leverage.Mul(cash.Neg()).
		Mul(p.Leverage.Sub(_1).Add(_2.Mul(beta)).Add(_2.Mul(Sqrt(beta.Mul(p.Leverage.Sub(_1).Add(beta)))))).
		Div(pos.Mul(Sqrt(p.Leverage.Sub(_1))))
	return index.GreaterThanOrEqual(newIndex)
}

func ComputeBuy(p *model.PerpetualStorage, amount decimal.Decimal) (decimal.Decimal, error) {
	if p.Position.GreaterThan(decimal.Zero) {
		if amount.LessThanOrEqual(decimal.Zero) {
			return decimal.Zero, fmt.Errorf("invalid amount")
		}
		if longSafe(p, p.IndexPrice, p.Cash, p.Position, p.Beta2) {
			m0 := CalculateM0(p, p.Cash, p.Position, p.Beta2)
			ma1 := p.Cash.Add(m0.Mul(p.Leverage.Sub(_1)).Div(p.Leverage))
			pos2 := p.Position.Sub(amount)
			ma2, err := calculateLongMa2(p, m0, ma1, p.Position, pos2, p.Beta2)
			if err != nil {
				return decimal.Zero, fmt.Errorf("calculate long ma2 error:%w", err)
			}
			if ma1.GreaterThan(ma2) {
				return decimal.Zero, fmt.Errorf("ma1 > ma2")
			}
			return ma2.Sub(ma1), nil
		}
		return p.IndexPrice.Mul(amount), nil
	}

	if p.Position.Equal(decimal.Zero) || shortSafe(p, p.IndexPrice, p.Cash, p.Position, p.Beta) {
		m0 := CalculateM0(p, p.Cash, p.Position, p.Beta)
		pos2 := p.Position.Sub(amount)
		if m0.Add(p.IndexPrice.Mul(pos2)).LessThanOrEqual(decimal.Zero) {
			return decimal.Zero, fmt.Errorf("pos2 unsafe")
		}
		ma1MinusMa2 := calculateShortMa1MinusMa2(p, m0, p.Position, pos2, p.Beta)
		if ma1MinusMa2.GreaterThan(decimal.Zero) {
			return decimal.Zero, fmt.Errorf("ma1MinusMa2 > 0")
		}
		newM0 := CalculateM0(p, p.Cash.Sub(ma1MinusMa2), pos2, p.Beta)
		if newM0.Sub(m0).Abs().GreaterThan(UnSafeM0Buy) {
			return decimal.Zero, fmt.Errorf("after buy unsafe(m0 change) old:%s new:%s", m0, newM0)
		}
		if !shortSafe(p, p.IndexPrice.Mul(_1.Add(p.IndexBuffer)), p.Cash.Sub(ma1MinusMa2), pos2, p.Beta) {
			return decimal.Zero, fmt.Errorf("unsafe after buy")
		}
		return ma1MinusMa2.Neg(), nil
	}

	return decimal.Zero, fmt.Errorf("unsafe before buy")

}

func ComputeSell(p *model.PerpetualStorage, amount decimal.Decimal) (decimal.Decimal, error) {
	if p.Position.LessThan(decimal.Zero) {
		if amount.GreaterThan(p.Position.Neg()) {
			return decimal.Zero, fmt.Errorf("invalid amount")
		}
		if shortSafe(p, p.IndexPrice, p.Cash, p.Position, p.Beta2) {
			m0 := CalculateM0(p, p.Cash, p.Position, p.Beta2)
			pos2 := p.Position.Add(amount)
			ma1MinusMa2 := calculateShortMa1MinusMa2(p, m0, p.Position, pos2, p.Beta2)
			if ma1MinusMa2.LessThan(decimal.Zero) {
				return decimal.Zero, fmt.Errorf("ma1 - ma2 < 0")
			}
			return ma1MinusMa2, nil
		}
		return p.IndexPrice.Mul(amount), nil
	}
	if p.Position.Equal(decimal.Zero) || shortSafe(p, p.IndexPrice, p.Cash, p.Position, p.Beta2) {
		m0 := CalculateM0(p, p.Cash, p.Position, p.Beta2)
		pos2 := p.Position.Add(amount)
		ma1 := p.Cash.Add(m0.Mul(p.Leverage.Sub(_1)).Div(p.Leverage))
		ma2, err := calculateLongMa2(p, m0, ma1, p.Position, pos2, p.Beta)
		if err != nil {
			return decimal.Zero, fmt.Errorf("calculate long ma2 error:%w", err)
		}
		if ma1.GreaterThan(ma2) {
			return decimal.Zero, fmt.Errorf("ma1 > ma2")
		}
		newM0 := CalculateM0(p, p.Cash.Sub(ma1.Sub(ma2)), pos2, p.Beta)
		if newM0.Sub(m0).Abs().GreaterThan(UnSafeM0Sell) {
			return decimal.Zero, fmt.Errorf("after buy unsafe(m0 change) old:%s new:%s", m0, newM0)
		}
		if !longSafe(p, p.IndexPrice.Mul(_1.Sub(p.IndexBuffer)), p.Cash.Sub(ma1.Sub(ma2)), pos2, p.Beta) {
			return decimal.Zero, fmt.Errorf("unsafe after sell")
		}
		return ma1.Sub(ma2), nil
	}
	return decimal.Zero, fmt.Errorf("unsafe before sell")
}

func Funding(p *model.PerpetualStorage, timestamp int64) {
	if p.LastFundingTime == -1 || timestamp == -1 {
		p.LastFundingTime = timestamp
		return
	}
	fundingRate := decimal.Zero
	m0 := CalculateM0(p, p.Cash, p.Position, p.Beta)
	ma1 := p.Cash.Add(m0.Mul(p.Leverage.Sub(_1)).Div(p.Leverage))
	if p.Position.GreaterThan(decimal.Zero) {
		fundingRate = p.Gama.Mul(_1.Sub(ma1.Div(m0)))
		p.FundingRate = fundingRate.Neg()
	} else if p.Position.LessThan(decimal.Zero) {
		fundingRate = p.Gama.Neg().Mul(p.IndexPrice.Mul(p.Position).Div(m0))
		p.FundingRate = fundingRate
	} else {
		p.FundingRate = fundingRate
	}
	fundingPayment := fundingRate.Mul(p.IndexPrice).Mul(p.Position.Abs()).Mul(decimal.NewFromInt(timestamp - p.LastFundingTime)).
		Mul(decimal.NewFromInt(60)).Div(decimal.NewFromInt(8)).Div(decimal.NewFromInt(3600))
	p.Cash = p.Cash.Add(fundingPayment)
	p.LastFundingTime = timestamp
}

func Buy(p *model.PerpetualStorage, amount decimal.Decimal, timestamp int64) decimal.Decimal {
	Funding(p, timestamp)
	if p.Position.GreaterThan(decimal.Zero) && amount.GreaterThan(p.Position) {
		leftAmount := amount.Sub(p.Position)
		payQuote1 := innerBuy(p, p.Position)
		payQuote2 := innerBuy(p, leftAmount)
		if payQuote1.IsZero() || payQuote2.IsZero() {
			return decimal.Zero
		}
		return payQuote1.Add(payQuote2)
	}
	return innerBuy(p, amount)
}

func innerBuy(p *model.PerpetualStorage, amount decimal.Decimal) decimal.Decimal {
	payQuote, err := ComputeBuy(p, amount)
	if err != nil {
		return decimal.Zero
	}
	p.Cash = p.Cash.Add(payQuote.Mul(p.Fee).Mul(decimal.NewFromFloat(0.8)))
	p.Cash = p.Cash.Add(payQuote.Mul(_1.Add(p.Spread)))
	p.Position = p.Position.Sub(amount)
	p.Position = p.Position.Round(4)
	return payQuote.Mul(_1.Add(p.Spread).Add(p.Fee))
}

func Sell(p *model.PerpetualStorage, amount decimal.Decimal, timestamp int64) decimal.Decimal {
	Funding(p, timestamp)
	if p.Position.LessThan(decimal.Zero) && amount.GreaterThan(p.Position.Neg()) {
		leftAmount := amount.Add(p.Position)
		payQuote1 := innerSell(p, p.Position.Neg())
		payQuote2 := innerSell(p, leftAmount)
		if payQuote1.IsZero() || payQuote2.IsZero() {
			return decimal.Zero
		}
		return payQuote1.Add(payQuote2)
	}
	return innerSell(p, amount)
}

func innerSell(p *model.PerpetualStorage, amount decimal.Decimal) decimal.Decimal {
	payQuote, err := ComputeSell(p, amount)
	if err != nil {
		return decimal.Zero
	}
	p.Cash = p.Cash.Add(payQuote.Mul(p.Fee).Mul(decimal.NewFromFloat(0.8)))
	p.Cash = p.Cash.Sub(payQuote.Mul(_1.Sub(p.Spread)))
	p.Position = p.Position.Add(amount)
	p.Position = p.Position.Round(4)
	return payQuote.Mul(_1.Sub(p.Spread).Sub(p.Fee))
}

func Print(p *model.PerpetualStorage) {
	m0 := CalculateM0(p, p.Cash, p.Position, p.Beta)
	v := m0.Mul(p.Leverage.Sub(_1)).Div(p.Leverage)
	ma1 := p.Cash.Add(v)
	fmt.Printf("cash=%s, pos=%s, m0=%s, v=%s, ma1=%s\n", p.Cash.Round(10), p.Position.Round(10), m0.Round(10), v.Round(10), ma1.Round(10))
}
