package utils

import (
	"github.com/shopspring/decimal"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString2XX(t *testing.T) {
	n1, err := String2BigInt("-1")
	assert.Nil(t, err)
	assert.Equal(t, "-1", n1.String())

	n2, err := StringToDecimal("-1")
	assert.Nil(t, err)
	assert.Equal(t, "-1", n2.String())
}

func TestHex(t *testing.T) {
	_, ok := new(big.Int).SetString("10000000000000000000000000000000000000000000000000000000000000000", 16)
	assert.True(t, ok)
}

func TestToWad(t *testing.T) {
	n := decimal.NewFromFloat(1.2)
	n1, _ := decimal.NewFromString("1000000000000000000")
	assert.Equal(t, n1, ToWad(n))
}

func TestMustDecimalToBigInt(t *testing.T) {
	assert.Equal(t, big.NewInt(1*1e18), MustDecimalToBigInt(ToWad(decimal.NewFromInt(1))))
}
