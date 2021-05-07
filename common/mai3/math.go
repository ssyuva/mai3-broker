package mai3

import (
	"math"

	"github.com/shopspring/decimal"
)

const DECIMALS = 18

var _0 = decimal.Zero
var _1 = decimal.NewFromInt(1)
var _2 = decimal.NewFromInt(2)
var _4 = decimal.NewFromInt(4)

func init() {
	decimal.DivisionPrecision = DECIMALS
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
func Gss(f func(float64) float64, a, b, tol float64, maxIter int) float64 {
	return gss(f, a, b, tol, nan, nan, nan, nan, maxIter)
}

func gss(f func(float64) float64, a, b, tol, c, d, fc, fd float64, maxIter int) float64 {
	if a > b {
		a, b = b, a
	}
	h := b - a
	for it := 0; it < maxIter; it++ {
		if h < tol {
			return (a + b) * 0.5
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
	return (a + b) * 0.5
}
