package mai3

import (
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Approximate(t *testing.T, expect, actual decimal.Decimal, msgAndArgs ...interface{}) {
	if expect.Sub(actual).Abs().GreaterThan(decimal.New(1, -12)) {
		assert.Fail(t, fmt.Sprintf("approximate: expect[%s] actual[%s]", expect, actual), msgAndArgs...)
	}
}

var defaultPool = &model.LiquidityPoolStorage{
	VaultFeeRate:         decimal.NewFromFloat(0.0002),
	InsuranceFundCap:     decimal.NewFromFloat(10000),
	InsuranceFund:        _0,
	DonatedInsuranceFund: _0,
	TotalClaimableFee:    _0,
	PoolCashBalance:      _0, // set me later
	FundingTime:          1579601290,

	Perpetuals: make(map[int64]*model.PerpetualStorage), // set me later
}

var perpetual1 = &model.PerpetualStorage{
	MarkPrice:               decimal.NewFromFloat(6965),
	IndexPrice:              decimal.NewFromFloat(7000),
	UnitAccumulativeFunding: decimal.NewFromFloat(9.9059375),

	InitialMarginRate:      decimal.NewFromFloat(0.1),
	MaintenanceMarginRate:  decimal.NewFromFloat(0.05),
	OperatorFeeRate:        decimal.NewFromFloat(0.0001),
	LpFeeRate:              decimal.NewFromFloat(0.0007),
	ReferrerRebateRate:     decimal.NewFromFloat(0.0000),
	LiquidationPenaltyRate: decimal.NewFromFloat(0.005),
	KeeperGasReward:        decimal.NewFromFloat(1),
	InsuranceFundRate:      decimal.NewFromFloat(0.0001),

	HalfSpread:          decimal.NewFromFloat(0.001),
	OpenSlippageFactor:  decimal.NewFromFloat(100),
	CloseSlippageFactor: decimal.NewFromFloat(90),
	FundingRateLimit:    decimal.NewFromFloat(0.005),
	MaxLeverage:         decimal.NewFromFloat(5),

	AmmPositionAmount: _0, // assign me later
}

var accountStorage1 = &model.AccountStorage{
	CashBalance:    decimal.NewFromFloat(7698.86),
	PositionAmount: decimal.NewFromFloat(2.3),
}

const TEST_PERPETUAL_INDEX0 = 0

var poolStorage0 = &model.LiquidityPoolStorage{
	VaultFeeRate:         decimal.NewFromFloat(0.0002),
	InsuranceFundCap:     decimal.NewFromFloat(10000),
	InsuranceFund:        _0,
	DonatedInsuranceFund: _0,
	TotalClaimableFee:    _0,
	PoolCashBalance:      decimal.NewFromFloat(100000),
	FundingTime:          1579601290,

	Perpetuals: make(map[int64]*model.PerpetualStorage), // set me later
}

// amm holds short, trader buys
// case 1: amm unsafe
func TestComputeAMMAmountWithPrice1(t *testing.T) {
	limitPrice := decimal.NewFromFloat(100000)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(16096.21634375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(-2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, true, limitPrice)
	assert.Equal(t, _0, amount)
}

// amm holds short, trader buys
// case 2: lower than index
func TestComputeAMMAmountWithPrice2(t *testing.T) {
	limitPrice := decimal.NewFromFloat(7000)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(16096.21634375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(-2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, true, limitPrice)
	assert.Equal(t, _0, amount)
}

// amm holds short, trader buys
// case 3: limitPrice is far away
func TestComputeAMMAmountWithPrice3(t *testing.T) {
	limitPrice := decimal.NewFromFloat(100000)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(116095.73134375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(-2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, true, limitPrice)
	Approximate(t, decimal.NewFromFloat(90.795503235030246126178607648), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	assert.Equal(t, tradingPrice.LessThan(limitPrice), true)
}

// amm holds short, trader buys
// case 4: normal
func TestComputeAMMAmountWithPrice4(t *testing.T) {
	limitPrice := decimal.NewFromFloat(7200).Mul(_1.Add(perpetual1.HalfSpread))
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(116095.73134375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(-2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, true, limitPrice)
	Approximate(t, decimal.NewFromFloat(52.542857142857142857), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	Approximate(t, tradingPrice, limitPrice)
}

// amm holds short, trader buys
// case 5: no solution
func TestComputeAMMAmountWithPrice5(t *testing.T) {
	limitPrice := decimal.NewFromFloat(7010).Mul(_1.Add(perpetual1.HalfSpread))
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(16096.21634375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(-2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, true, limitPrice)
	assert.Equal(t, _0, amount)
}

// amm holds short, trader sells
// case 6:amm unsafe - impossible price
func TestComputeAMMAmountWithPrice6(t *testing.T) {
	limitPrice := decimal.NewFromFloat(100000)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(16096.21634375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(-2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, false, limitPrice)
	assert.Equal(t, _0, amount)
}

// amm holds short, trader sells
// case 7:amm unsafe - largest amount
func TestComputeAMMAmountWithPrice7(t *testing.T) {
	limitPrice := decimal.NewFromFloat(0)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(16096.21634375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(-2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, false, limitPrice)
	Approximate(t, decimal.NewFromFloat(-2.317688145614655747), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	assert.Equal(t, tradingPrice.GreaterThanOrEqual(limitPrice), true)
}

// amm holds short, trader sells
// case 8:amm unsafe close + open
func TestComputeAMMAmountWithPrice8(t *testing.T) {
	limitPrice := decimal.NewFromFloat(6998).Mul(_1.Sub(perpetual1.HalfSpread))
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(16096.21634375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(-2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, false, limitPrice)
	Approximate(t, decimal.NewFromFloat(-2.315856729596580035), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	Approximate(t, tradingPrice, limitPrice)
}

// amm holds short, trader sells
// case 9:close only
func TestComputeAMMAmountWithPrice9(t *testing.T) {
	limitPrice := decimal.NewFromFloat(7014.174737510995938681826).Mul(_1.Sub(perpetual1.HalfSpread))
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(116095.73134375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(-2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, false, limitPrice)
	Approximate(t, decimal.NewFromFloat(-0.1), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	Approximate(t, tradingPrice, limitPrice)
}

// amm holds short, trader sells
// case 10:close + open
func TestComputeAMMAmountWithPrice10(t *testing.T) {
	limitPrice := decimal.NewFromFloat(7003.9888655277856883937409).Mul(_1.Sub(perpetual1.HalfSpread))
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(116095.73134375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(-2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, false, limitPrice)
	Approximate(t, decimal.NewFromFloat(-3.3), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	Approximate(t, tradingPrice, limitPrice)
}

// amm holds long, trader buys
// case 11: amm unsafe - impossible price
func TestComputeAMMAmountWithPrice11(t *testing.T) {
	limitPrice := decimal.NewFromFloat(0)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(-13977.21634375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, true, limitPrice)
	assert.Equal(t, _0, amount)
}

// amm holds long, trader buys
// case 12: amm unsafe - largest amount
func TestComputeAMMAmountWithPrice12(t *testing.T) {
	limitPrice := decimal.NewFromFloat(100000)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(-13977.21634375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, true, limitPrice)
	Approximate(t, decimal.NewFromFloat(4.255005567935635169), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	assert.Equal(t, tradingPrice.LessThanOrEqual(limitPrice), true)
}

// // amm holds long, trader buys
// // case 13: amm unsafe - close + open
func TestComputeAMMAmountWithPrice13(t *testing.T) {
	limitPrice := decimal.NewFromFloat(7002).Mul(_1.Add(perpetual1.HalfSpread))
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(-13977.21634375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, true, limitPrice)
	Approximate(t, decimal.NewFromFloat(2.472240789218530841), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	Approximate(t, tradingPrice, limitPrice)
}

// // amm holds long, trader buys
// // case 14: close only
func TestComputeAMMAmountWithPrice14(t *testing.T) {
	limitPrice := decimal.NewFromFloat(6985.825262489004061318174).Mul(_1.Add(perpetual1.HalfSpread))
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(83941.29865625)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, true, limitPrice)
	Approximate(t, decimal.NewFromFloat(0.1), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	Approximate(t, tradingPrice, limitPrice)
}

// // amm holds long, trader buys
// // case 15: close + open
func TestComputeAMMAmountWithPrice15(t *testing.T) {
	limitPrice := decimal.NewFromFloat(6996.0111344722143116062591).Mul(_1.Add(perpetual1.HalfSpread))
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(83941.29865625)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, true, limitPrice)
	Approximate(t, decimal.NewFromFloat(3.3), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	Approximate(t, tradingPrice, limitPrice)
}

// // amm holds long, trader sells
// // case 16: amm unsafe
func TestComputeAMMAmountWithPrice16(t *testing.T) {
	limitPrice := decimal.NewFromFloat(0)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(-13977.21634375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, false, limitPrice)
	assert.Equal(t, _0, amount)
}

// // amm holds long, trader sells
// // case 17: higher than index
func TestComputeAMMAmountWithPrice17(t *testing.T) {
	limitPrice := decimal.NewFromFloat(7000)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(83941.29865625)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, false, limitPrice)
	assert.Equal(t, _0, amount)
}

// // amm holds long, trader sells
// // case 18: limitPrice is far away
func TestComputeAMMAmountWithPrice18(t *testing.T) {
	limitPrice := decimal.NewFromFloat(0)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(83941.29865625)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, false, limitPrice)
	Approximate(t, decimal.NewFromFloat(-90.795503235030246126178607648), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	assert.Equal(t, tradingPrice.GreaterThanOrEqual(limitPrice), true)
}

// // amm holds long, trader sells
// // case 19: normal
func TestComputeAMMAmountWithPrice19(t *testing.T) {
	limitPrice := decimal.NewFromFloat(6800).Mul(_1.Sub(perpetual1.HalfSpread))
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(83941.29865625)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, false, limitPrice)
	Approximate(t, decimal.NewFromFloat(-52.542857142857142857), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	Approximate(t, tradingPrice, limitPrice)
}

// // amm holds long, trader sells
// // case 20: no solution
func TestComputeAMMAmountWithPrice20(t *testing.T) {
	limitPrice := decimal.NewFromFloat(6990).Mul(_1.Sub(perpetual1.HalfSpread))
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(83941.29865625)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, false, limitPrice)
	assert.Equal(t, _0, amount)
}
