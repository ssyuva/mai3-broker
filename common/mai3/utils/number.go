package utils

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/shopspring/decimal"
)

var (
	wad, _  = decimal.NewFromString("1000000000000000000")
	rate, _ = decimal.NewFromString("10000")
)

func ToWad(n decimal.Decimal) decimal.Decimal {
	return n.Mul(wad)
}

func ToRate(n decimal.Decimal) decimal.Decimal {
	return n.Mul(rate)
}

func String2BigInt(str string) (*big.Int, error) {
	n := new(big.Int)
	if _, ok := n.SetString(str, 0); !ok {
		return nil, fmt.Errorf("String2BigInt:parse string[%s] error", str)
	}
	return n, nil
}

func StringToDecimal(str string) (decimal.Decimal, error) {
	if strings.HasPrefix(str, "0x") || strings.HasPrefix(str, "0X") {
		b := new(big.Int)
		if _, ok := b.SetString(str[2:], 16); !ok {
			return decimal.Zero, fmt.Errorf("StringToDecimal:parse[%s] error", str)
		}
		d := decimal.NewFromBigInt(b, 0)
		return d, nil
	} else {
		v, err := decimal.NewFromString(str)
		if err != nil {
			return v, fmt.Errorf("StringToDecimal:%w", err)
		}
		return v, nil
	}
}

func MustDecimalToBigInt(d decimal.Decimal) *big.Int {
	n := new(big.Int)
	n, ok := n.SetString(d.Floor().String(), 10)
	if !ok {
		panic(fmt.Errorf("decimalToBigInt error %+v", d))
	}
	return n
}

func HasTheSameSign(x, y decimal.Decimal) bool {
	if x.IsZero() || y.IsZero() {
		return true
	}
	return (x.Sign() ^ y.Sign()) == 0
}

func SplitAmount(position, amount decimal.Decimal) (close, open decimal.Decimal) {
	if HasTheSameSign(position, amount) {
		close = decimal.Zero
		open = amount
	} else if position.Abs().GreaterThanOrEqual(amount.Abs()) {
		close = amount
		open = decimal.Zero
	} else {
		close = position.Neg()
		open = position.Add(amount)
	}
	return
}
