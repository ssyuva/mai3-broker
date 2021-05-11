package match

import (
	"fmt"
	"testing"

	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

var defaultPool = &model.LiquidityPoolStorage{
	VaultFeeRate:    decimal.NewFromFloat(0.0002),
	PoolCashBalance: decimal.NewFromInt(0), // set me later

	Perpetuals: make(map[int64]*model.PerpetualStorage), // set me later
}

var OpenSlippageFactor, _ = decimal.NewFromString("0.0142857142857142857142857142857")
var CloseSlippageFactor, _ = decimal.NewFromString("0.0128571428571428571428571428571")

var perpetual1 = &model.PerpetualStorage{
	IsNormal: true,

	MarkPrice:               decimal.NewFromFloat(6965),
	IndexPrice:              decimal.NewFromFloat(7000),
	UnitAccumulativeFunding: decimal.NewFromFloat(9.9059375),

	InitialMarginRate:      decimal.NewFromFloat(0.1),
	MaintenanceMarginRate:  decimal.NewFromFloat(0.05),
	OperatorFeeRate:        decimal.NewFromFloat(0.0001),
	LpFeeRate:              decimal.NewFromFloat(0.0007),
	ReferrerRebateRate:     decimal.NewFromFloat(0),
	LiquidationPenaltyRate: decimal.NewFromFloat(0.005),
	KeeperGasReward:        decimal.NewFromFloat(1),
	InsuranceFundRate:      decimal.NewFromFloat(0.0001),
	OpenInterest:           decimal.NewFromFloat(10),
	MaxOpenInterestRate:    decimal.NewFromFloat(100),

	HalfSpread:            decimal.NewFromFloat(0.001),
	OpenSlippageFactor:    OpenSlippageFactor,
	CloseSlippageFactor:   CloseSlippageFactor,
	FundingRateLimit:      decimal.NewFromFloat(0.005),
	MaxLeverage:           decimal.NewFromFloat(5),
	MaxClosePriceDiscount: decimal.NewFromFloat(0.05),

	AmmCashBalance:    decimal.NewFromInt(0), // assign me later
	AmmPositionAmount: decimal.NewFromInt(0), // assign me later
}

var accountStorage1 = &model.AccountStorage{
	CashBalance:    decimal.NewFromFloat(7698.86),
	PositionAmount: decimal.NewFromFloat(2.3),
	TargetLeverage: decimal.NewFromFloat(2),
}

var accountStorage0 = &model.AccountStorage{
	CashBalance:    decimal.NewFromFloat(1000),
	PositionAmount: decimal.NewFromFloat(0),
	TargetLeverage: decimal.NewFromFloat(2),
}

const TEST_MARKET_INDEX0 = 0

func Approximate(t *testing.T, expect, actual decimal.Decimal, msgAndArgs ...interface{}) {
	if expect.Sub(actual).Abs().GreaterThan(decimal.New(1, -12)) {
		assert.Fail(t, fmt.Sprintf("approximate: expect[%s] actual[%s]", expect, actual), msgAndArgs...)
	}
}

// empty order book. close only
func TestComputeOrderAvailable1(t *testing.T) {
	order := &model.Order{
		AvailableAmount: decimal.NewFromFloat(-1),
	}
	order.Price = decimal.NewFromFloat(6900)
	poolStorage := defaultPool
	perpetual1.AmmCashBalance = decimal.NewFromFloat(2.3)
	poolStorage.PoolCashBalance = decimal.NewFromFloat(83941.29865625)
	poolStorage.Perpetuals[TEST_MARKET_INDEX0] = perpetual1

	_, available := ComputeOrderAvailable(poolStorage, TEST_MARKET_INDEX0, accountStorage1, []*model.Order{order})
	fmt.Println(available)
	Approximate(t, decimal.NewFromFloat(0), available)
}

// empty order book. close + open. withdraw covers deposit
func TestComputeOrderAvailable2(t *testing.T) {
	order := &model.Order{
		AvailableAmount: decimal.NewFromFloat(-3.3),
	}
	order.Price = decimal.NewFromFloat(6900)
	poolStorage := defaultPool
	perpetual1.AmmCashBalance = decimal.NewFromFloat(2.3)
	poolStorage.PoolCashBalance = decimal.NewFromFloat(83941.29865625)
	poolStorage.Perpetuals[TEST_MARKET_INDEX0] = perpetual1

	_, available := ComputeOrderAvailable(poolStorage, TEST_MARKET_INDEX0, accountStorage1, []*model.Order{order})
	Approximate(t, decimal.NewFromFloat(0), available)
}

// empty order book. close + open
func TestComputeOrderAvailable3(t *testing.T) {
	order := &model.Order{
		AvailableAmount: decimal.NewFromFloat(-10),
	}
	order.Price = decimal.NewFromFloat(6900)

	poolStorage := defaultPool
	perpetual1.AmmCashBalance = decimal.NewFromFloat(2.3)
	poolStorage.PoolCashBalance = decimal.NewFromFloat(83941.29865625)
	poolStorage.Perpetuals[TEST_MARKET_INDEX0] = perpetual1

	// marginBalance = 23695.57634375
	// withdraw = 23695.57634375 - (6965 - 6900)*2.3 - 6900*2.3*0.001 = 23530.20634375
	// deposit = (10 - 2.3)*6900*(1/2 + 0.001) + (6965 - 6900)*(10 - 2.3) = 27118.6
	// cost = deposit - withdraw = 3588.42
	_, available := ComputeOrderAvailable(poolStorage, TEST_MARKET_INDEX0, accountStorage1, []*model.Order{order})
	Approximate(t, decimal.NewFromFloat(-3588.42365625), available)
}

// empty order book. pos = 0 but cash > 0. open
func TestComputeOrderAvailable4(t *testing.T) {
	order := &model.Order{
		AvailableAmount: decimal.NewFromFloat(-10),
	}
	order.Price = decimal.NewFromFloat(6900)

	poolStorage := defaultPool
	perpetual1.AmmCashBalance = decimal.NewFromFloat(2.3)
	poolStorage.PoolCashBalance = decimal.NewFromFloat(83941.29865625)
	poolStorage.Perpetuals[TEST_MARKET_INDEX0] = perpetual1

	// deposit = 10*6900*(1/2 + 0.001) + (6965 - 6900)*10 = 35219
	// cost = deposit - 1000 = 34219
	_, available := ComputeOrderAvailable(poolStorage, TEST_MARKET_INDEX0, accountStorage0, []*model.Order{order})
	Approximate(t, decimal.NewFromFloat(-34219), available)
}
