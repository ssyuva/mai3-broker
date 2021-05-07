package mai3

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestSqrt(t *testing.T) {
	n := Sqrt(decimal.NewFromInt(25))
	assert.Equal(t, true, decimal.NewFromInt(5).Equal(n))

	n1 := Sqrt(decimal.NewFromFloat(3.56666))
	Approximate(t, decimal.NewFromFloat(1.888560298216607115), n1)
}

func TestGss(t *testing.T) {
	f := func(x float64) float64 {
		tmp := x - 2
		return tmp * tmp
	}
	res := Gss(f, 1, 5, 1e-8, 15)
	assert.Less(t, 1.9987067953999975, res)
	assert.Greater(t, 2.001639345143427, res)
}
