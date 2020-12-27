package mai3

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSqrt(t *testing.T) {
	n := Sqrt(decimal.NewFromInt(25))
	assert.Equal(t, true, decimal.NewFromInt(5).Equal(n))

	n1 := Sqrt(decimal.NewFromFloat(3.56666))
	Approximate(t, decimal.NewFromFloat(1.888560298216607115), n1)
}
