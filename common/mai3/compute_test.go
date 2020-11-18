package mai3

import (
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/shopspring/decimal"
	// "github.com/stretchr/testify/assert"
	"testing"
)

var storage = &model.PerpetualStorage{
	Leverage:        decimal.NewFromFloat(5),
	Cash:            decimal.NewFromFloat(1000000),
	Position:        decimal.Zero,
	IndexPrice:      decimal.NewFromFloat(200),
	Beta:            decimal.NewFromFloat(0.1),
	Beta2:           decimal.NewFromFloat(0.06),
	Fee:             decimal.NewFromFloat(0.001),
	Spread:          decimal.NewFromFloat(0.001),
	Gama:            decimal.NewFromFloat(0.05),
	LastFundingTime: -1,
	FundingRate:     decimal.Zero,
	IndexBuffer:     decimal.NewFromFloat(0.05),
}

func TestAMM1(t *testing.T) {
	storage.Leverage = decimal.NewFromFloat(2)
	storage.Beta = decimal.NewFromFloat(0.1)
	Print(storage)
	Buy(storage, decimal.NewFromFloat(10), -1)
	Print(storage)
	storage.IndexPrice = decimal.NewFromFloat(101)
	Print(storage)
	Sell(storage, decimal.NewFromFloat(10), -1)
	Print(storage)
}

func TestAMM2(t *testing.T) {
	storage.Leverage = decimal.NewFromFloat(2)
	storage.Beta = decimal.NewFromFloat(0.1)
	Print(storage)
	Buy(storage, decimal.NewFromFloat(139), -1)
	Print(storage)
	storage.IndexPrice = decimal.NewFromFloat(99)
	Print(storage)
	Sell(storage, decimal.NewFromFloat(139), -1)
	Print(storage)
}
