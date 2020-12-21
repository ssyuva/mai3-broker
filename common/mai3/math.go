package mai3

import (
	"github.com/shopspring/decimal"
	"math"
)

const DECIMALS = 18

type BigNumber = decimal.Decimal

func toDecimal(v string) decimal.Decimal {
	r, err := decimal.NewFromString(v)
	if err != nil {
		panic(err)
	}
	return r
}

var _0 = decimal.Zero
var _1 = decimal.NewFromInt(1)
var _2 = decimal.NewFromInt(2)
var _4 = decimal.NewFromInt(4)
var _10 = decimal.NewFromInt(10)
var _0_1 = toDecimal("0.1")
var _E = toDecimal("2.718281828459045235")
var _1_5 = toDecimal("1.5")
var _LN_1_5 = toDecimal("0.405465108108164381978013115464349137")
var _LN_10 = toDecimal("2.302585092994045684017991454684364208")
var _MAX_LN = toDecimal("10000000000000000000000000000000000000000")

func init() {
	decimal.DivisionPrecision = DECIMALS
}

func bigLn(v BigNumber) BigNumber {
	if v.IsNegative() {
		panic("logE of negative number")
	}
	if v.GreaterThan(_MAX_LN) {
		panic("logE only accepts v <= 1e22 * 1e18")
	}

	x := v
	r := _0

	for x.LessThanOrEqual(_0_1) {
		x = x.Mul(_10)
		r = r.Sub(_LN_10)
	}
	for x.GreaterThanOrEqual(_10) {
		x = x.Div(_10)
		r = r.Add(_LN_10)
	}
	for x.LessThan(_1) {
		x = x.Mul(_E)
		r = r.Sub(_1)
	}
	for x.GreaterThan(_E) {
		x = x.Div(_E)
		r = r.Add(_1)
	}
	if x.Equal(_1) {
		return r.Truncate(DECIMALS)
	}
	if x.Equal(_E) {
		return _1.Add(r.Truncate(DECIMALS))
	}

	//                    2    x           2    x          2    x
	// Ln(a+x) = Ln(a) + ---(------)^1  + ---(------)^3 + ---(------)^5 + ...
	//                    1   2a+x         3   2a+x        5   2a+x
	// Let x = v - a
	//                  2   v-a         2   v-a        2   v-a
	// Ln(v) = Ln(a) + ---(-----)^1  + ---(-----)^3 + ---(-----)^5 + ...
	//                  1   v+a         3   v+a        5   v+a
	r = r.Add(_LN_1_5)
	a1_5 := _1_5
	m := _1.Mul(x.Sub(a1_5).Div(x.Add(a1_5)))
	r = r.Add(m.Mul(_2))
	m2 := m.Mul(m)
	var i int = 3
	for {
		m = m.Mul(m2)
		r = r.Add(m.Mul(_2).Div(decimal.NewFromInt(int64(i))))
		i += 2
		if i >= 3+2*DECIMALS {
			break
		}
	}
	return r.Truncate(DECIMALS)
}

func bigLog(base, x BigNumber) BigNumber {
	return bigLn(x).Div(bigLn(base))
}

// x ^ n
func wpowi(x decimal.Decimal, n int64) decimal.Decimal {
	z := x
	if n%2 == 0 {
		z = decimal.NewFromInt(1)
	}
	for n /= 2; n != 0; n /= 2 {
		x = x.Mul(x)
		if n%2 != 0 {
			z = z.Mul(x)
		}
		x = x.Truncate(DECIMALS)
		z = z.Truncate(DECIMALS)
	}
	return z
}

const SqrtMaxIter = 100000

// Sqrt returns the square root of d, accurate to DivisionPrecision decimal places.
func Sqrt(d decimal.Decimal) decimal.Decimal {
	s, _ := SqrtRound(d, int32(decimal.DivisionPrecision))
	return s
}

// SqrtRound returns the square root of d, accurate to precision decimal places.
// The bool precise returns whether the precision was reached.
func SqrtRound(d decimal.Decimal, precision int32) (decimal.Decimal, bool) {
	maxError := decimal.New(1, -precision)
	one := decimal.NewFromFloat(1)
	var lo decimal.Decimal
	var hi decimal.Decimal
	// Handle cases where d < 0, d = 0, 0 < d < 1, and d > 1
	if d.GreaterThanOrEqual(one) {
		lo = decimal.Zero
		hi = d
	} else if d.Equal(one) {
		return one, true
	} else if d.LessThan(decimal.Zero) {
		return decimal.NewFromFloat(-1), false // call this an error , cannot take sqrt of neg w/o imaginaries
	} else if d.Equal(decimal.Zero) {
		return decimal.Zero, true
	} else {
		// d is between 0 and 1. Therefore, 0 < d < Sqrt(d) < 1.
		lo = d
		hi = one
	}
	var mid decimal.Decimal
	for i := 0; i < SqrtMaxIter; i++ {
		mid = lo.Add(hi).Div(decimal.New(2, 0)) //mid = (lo+hi)/2;
		if mid.Mul(mid).Sub(d).Abs().LessThan(maxError) {
			return mid, true
		}
		if mid.Mul(mid).GreaterThan(d) {
			hi = mid
		} else {
			lo = mid
		}
	}
	return mid, false
}

var (
	sqrt5   = math.Sqrt(5)
	invphi  = (sqrt5 - 1) / 2 //# 1/phi
	invphi2 = (3 - sqrt5) / 2 //# 1/phi^2
	nan     = math.NaN()
)

// Gss golden section search (recursive version)
// https://en.wikipedia.org/wiki/Golden-section_search
// '''
// Golden section search, recursive.
// Given a function f with a single local minimum in
// the interval [a,b], gss returns a subset interval
// [c,d] that contains the minimum with d-c <= tol.
//
//
// example:
// >>> f = lambda x: (x-2)**2
// >>> a = 1
// >>> b = 5
// >>> tol = 1e-5
// >>> (c,d) = gssrec(f, a, b, tol)
// >>> print (c,d)
// (1.9999959837979107, 2.0000050911830893)
// '''
func Gss(f func(float64) float64, a, b, tol float64, maxIter int) (float64, float64) {
	return gss(f, a, b, tol, nan, nan, nan, nan, nan, maxIter)
}

func gss(f func(float64) float64, a, b, tol, h, c, d, fc, fd float64, maxIter int) (float64, float64) {
	if a > b {
		a, b = b, a
	}
	h = b - a
	for it := 0; it < maxIter; it++ {
		if h < tol {
			return a, b
		}
		if a > b {
			a, b = b, a
		}
		if math.IsNaN(c) {
			c = a + invphi2*h
			fc = f(c)
		}
		if math.IsNaN(d) {
			d = a + invphi*h
			fd = f(d)
		}
		if fc < fd {
			b, h, c, fc, d, fd = d, h*invphi, nan, nan, c, fc
		} else {
			a, h, c, fc, d, fd = c, h*invphi, d, fd, nan, nan
		}
	}
	return a, b
}
