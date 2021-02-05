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
	VaultFeeRate:    decimal.NewFromFloat(0.0002),
	PoolCashBalance: _0, // set me later

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
	ReferrerRebateRate:     decimal.NewFromFloat(0.0000),
	LiquidationPenaltyRate: decimal.NewFromFloat(0.005),
	KeeperGasReward:        decimal.NewFromFloat(1),
	InsuranceFundRate:      decimal.NewFromFloat(0.0001),
	InsuranceFundCap:       decimal.NewFromFloat(10000),
	InsuranceFund:          _0,
	DonatedInsuranceFund:   _0,

	HalfSpread:            decimal.NewFromFloat(0.001),
	OpenSlippageFactor:    OpenSlippageFactor,
	CloseSlippageFactor:   CloseSlippageFactor,
	FundingRateLimit:      decimal.NewFromFloat(0.005),
	MaxLeverage:           decimal.NewFromFloat(5),
	MaxClosePriceDiscount: decimal.NewFromFloat(0.05),

	AmmCashBalance:    _0, // assign me later
	AmmPositionAmount: _0, // assign me later
}

var perpetual0 = &model.PerpetualStorage{
	IsNormal: true,

	MarkPrice:               decimal.NewFromFloat(95),
	IndexPrice:              decimal.NewFromFloat(100),
	UnitAccumulativeFunding: decimal.NewFromFloat(1.9),

	InitialMarginRate:      decimal.NewFromFloat(0.1),
	MaintenanceMarginRate:  decimal.NewFromFloat(0.05),
	OperatorFeeRate:        decimal.NewFromFloat(0.0001),
	LpFeeRate:              decimal.NewFromFloat(0.0008),
	ReferrerRebateRate:     decimal.NewFromFloat(0.0000),
	LiquidationPenaltyRate: decimal.NewFromFloat(0.005),
	KeeperGasReward:        decimal.NewFromFloat(2),
	InsuranceFundRate:      decimal.NewFromFloat(0.0001),
	InsuranceFundCap:       decimal.NewFromFloat(10000),
	InsuranceFund:          _0,
	DonatedInsuranceFund:   _0,

	HalfSpread:            decimal.NewFromFloat(0.001),
	OpenSlippageFactor:    decimal.NewFromFloat(1),
	CloseSlippageFactor:   decimal.NewFromFloat(0.9),
	FundingRateLimit:      decimal.NewFromFloat(0.005),
	MaxLeverage:           decimal.NewFromFloat(3),
	MaxClosePriceDiscount: decimal.NewFromFloat(0.2),

	AmmCashBalance:    _0, // assign me later
	AmmPositionAmount: _0, // assign me later
}

var perpetual2 = &model.PerpetualStorage{
	IsNormal: true,

	MarkPrice:               decimal.NewFromFloat(95),
	IndexPrice:              decimal.NewFromFloat(100),
	UnitAccumulativeFunding: decimal.NewFromFloat(1.9),

	InitialMarginRate:      decimal.NewFromFloat(0.1),
	MaintenanceMarginRate:  decimal.NewFromFloat(0.05),
	OperatorFeeRate:        decimal.NewFromFloat(0.0001),
	LpFeeRate:              decimal.NewFromFloat(0.0008),
	ReferrerRebateRate:     decimal.NewFromFloat(0.0000),
	LiquidationPenaltyRate: decimal.NewFromFloat(0.005),
	KeeperGasReward:        decimal.NewFromFloat(2),
	InsuranceFundRate:      decimal.NewFromFloat(0.0001),
	InsuranceFundCap:       decimal.NewFromFloat(10000),
	InsuranceFund:          _0,
	DonatedInsuranceFund:   _0,

	HalfSpread:            decimal.NewFromFloat(0.001),
	OpenSlippageFactor:    decimal.NewFromFloat(1),
	CloseSlippageFactor:   decimal.NewFromFloat(0.9),
	FundingRateLimit:      decimal.NewFromFloat(0.005),
	MaxLeverage:           decimal.NewFromFloat(3),
	MaxClosePriceDiscount: decimal.NewFromFloat(0.2),

	AmmCashBalance:    _0, // assign me later
	AmmPositionAmount: _0, // assign me later
}

var accountStorage1 = &model.AccountStorage{
	CashBalance:    decimal.NewFromFloat(7698.86),
	PositionAmount: decimal.NewFromFloat(2.3),
}

const TEST_PERPETUAL_INDEX0 = 0
const TEST_PERPETUAL_INDEX1 = 1

var poolStorage0 = &model.LiquidityPoolStorage{
	VaultFeeRate:    decimal.NewFromFloat(0.0002),
	PoolCashBalance: decimal.NewFromFloat(100000),

	Perpetuals: make(map[int64]*model.PerpetualStorage), // set me later
}

// amm holds short, trader buys
// case 1: amm unsafe
func TestComputeAMMAmountWithPrice1(t *testing.T) {
	limitPrice := decimal.NewFromFloat(100000)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(17096.21634375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(-2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, true, limitPrice)
	assert.Equal(t, _0, amount)
}

// amm holds short, trader buys
// case 2: lower than spread
func TestComputeAMMAmountWithPrice2(t *testing.T) {
	limitPrice := decimal.NewFromFloat(7023.1160999)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(116095.73134375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(-2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, true, limitPrice)
	assert.Equal(t, _0, amount)
}

// amm holds short, trader buys
// case 3: exactly the best ask/bid price
func TestComputeAMMAmountWithPrice3(t *testing.T) {
	limitPrice := decimal.NewFromFloat(7023.1161)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(116095.73134375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(-2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, true, limitPrice)
	Approximate(t, decimal.NewFromFloat(2.0046), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	assert.Equal(t, tradingPrice.LessThanOrEqual(limitPrice), true)
}

// amm holds short, trader buys
// case 4: limitPrice is far away
func TestComputeAMMAmountWithPrice4(t *testing.T) {
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
// case 5: normal
func TestComputeAMMAmountWithPrice5(t *testing.T) {
	limitPrice := decimal.NewFromFloat(7200)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(116095.73134375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(-2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, true, limitPrice)
	Approximate(t, decimal.NewFromFloat(52.542857142857142857), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	Approximate(t, tradingPrice, limitPrice)
}

// amm holds short, trader sells
// case 6: higher than spread
func TestComputeAMMAmountWithPrice6(t *testing.T) {
	limitPrice := decimal.NewFromFloat(7000.001)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(17096.21634375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(-2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, false, limitPrice)
	assert.Equal(t, _0, amount)
}

// amm holds short, trader sells
// case 7:amm unsafe - exactly the best ask/bid price - close + open
func TestComputeAMMAmountWithPrice7(t *testing.T) {
	limitPrice := decimal.NewFromFloat(7000)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(17096.21634375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(-2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, false, limitPrice)
	Approximate(t, decimal.NewFromFloat(-2.3), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	assert.Equal(t, tradingPrice.LessThanOrEqual(limitPrice), true)
}

// amm holds short, trader sells
// case 8:amm unsafe - largest amount
func TestComputeAMMAmountWithPrice8(t *testing.T) {
	limitPrice := decimal.NewFromFloat(0)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(17096.21634375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(-2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, false, limitPrice)
	Approximate(t, decimal.NewFromFloat(-3.248643177964958208), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	assert.Equal(t, tradingPrice.GreaterThanOrEqual(limitPrice), true)
}

// amm holds short, trader sells
// case 9:amm unsafe close + open
func TestComputeAMMAmountWithPrice9(t *testing.T) {
	limitPrice := decimal.NewFromFloat(6992)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(17096.21634375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(-2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, false, limitPrice)
	Approximate(t, decimal.NewFromFloat(-2.54339106672133532007243536012), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	Approximate(t, tradingPrice, limitPrice)
}

// amm holds short, trader sells
// case 10: higher than spread
func TestComputeAMMAmountWithPrice10(t *testing.T) {
	limitPrice := decimal.NewFromFloat(7007.476)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(116095.73134375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(-2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, false, limitPrice)
	assert.Equal(t, _0, amount)
}

// amm holds short, trader sells
// case 11:safe - exactly the best ask/bid price - close + open
func TestComputeAMMAmountWithPrice11(t *testing.T) {
	limitPrice := decimal.NewFromFloat(7007.4752419462290525818804101137)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(116095.73134375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(-2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, false, limitPrice)
	Approximate(t, decimal.NewFromFloat(-2.226863373523786822), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	Approximate(t, tradingPrice, limitPrice)
}

// amm holds short, trader sells
// case 12:close only
func TestComputeAMMAmountWithPrice12(t *testing.T) {
	limitPrice := decimal.NewFromFloat(7007.4)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(116095.73134375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(-2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, false, limitPrice)
	Approximate(t, decimal.NewFromFloat(-2.250750147989139645), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	Approximate(t, tradingPrice, limitPrice)
}

// amm holds short, trader sells
// case 13:close + open
func TestComputeAMMAmountWithPrice13(t *testing.T) {
	limitPrice := decimal.NewFromFloat(7006)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(116095.73134375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(-2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, false, limitPrice)
	Approximate(t, decimal.NewFromFloat(-2.688951590780905289), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	Approximate(t, tradingPrice, limitPrice)
}

// amm holds long, trader buys
// case 14: amm unsafe - lower than spread
func TestComputeAMMAmountWithPrice14(t *testing.T) {
	limitPrice := decimal.NewFromFloat(6999.999)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(-13677.21634375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, true, limitPrice)
	assert.Equal(t, _0, amount)
}

// amm holds long, trader buys
// case 15: amm unsafe - exactly the best ask/bid price - close + open
func TestComputeAMMAmountWithPrice15(t *testing.T) {
	limitPrice := decimal.NewFromFloat(7000)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(-13677.21634375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, true, limitPrice)
	Approximate(t, decimal.NewFromFloat(2.3), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	Approximate(t, tradingPrice, limitPrice)
}

// amm holds long, trader buys
// case 16: amm unsafe - largest amount
func TestComputeAMMAmountWithPrice16(t *testing.T) {
	limitPrice := decimal.NewFromFloat(100000)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(-13677.21634375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, true, limitPrice)
	Approximate(t, decimal.NewFromFloat(4.534292077640725907), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	assert.Equal(t, tradingPrice.LessThanOrEqual(limitPrice), true)
}

// amm holds long, trader buys
// case 17: amm unsafe - close + open
func TestComputeAMMAmountWithPrice17(t *testing.T) {
	limitPrice := decimal.NewFromFloat(7008)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(-13677.21634375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, true, limitPrice)
	Approximate(t, decimal.NewFromFloat(2.68369217482083603940884140606), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	Approximate(t, tradingPrice, limitPrice)
}

// amm holds long, trader buys
// case 18: lower than spread
func TestComputeAMMAmountWithPrice18(t *testing.T) {
	limitPrice := decimal.NewFromFloat(6992.495)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(83941.29865625)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, true, limitPrice)
	assert.Equal(t, _0, amount)
}

// amm holds long, trader buys
// case 19: exactly the best ask/bid price - close + open
func TestComputeAMMAmountWithPrice19(t *testing.T) {
	limitPrice, _ := decimal.NewFromString("6992.4957785904151334990367462224")
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(83941.29865625)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, true, limitPrice)
	Approximate(t, decimal.NewFromFloat(2.217663373523786822), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	Approximate(t, tradingPrice, limitPrice)
}

// amm holds long, trader buys
// case 20: close only
func TestComputeAMMAmountWithPrice20(t *testing.T) {
	limitPrice := decimal.NewFromFloat(6992.7)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(83941.29865625)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, true, limitPrice)
	Approximate(t, decimal.NewFromFloat(2.282496767610908028), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	Approximate(t, tradingPrice, limitPrice)
}

// amm holds long, trader buys
// case 21: close + open
func TestComputeAMMAmountWithPrice21(t *testing.T) {
	limitPrice := decimal.NewFromFloat(6994)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(83941.29865625)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, true, limitPrice)
	Approximate(t, decimal.NewFromFloat(2.688951590780905289), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	Approximate(t, tradingPrice, limitPrice)
}

// amm holds long, trader sells
// case 22: amm unsafe
func TestComputeAMMAmountWithPrice22(t *testing.T) {
	limitPrice := decimal.NewFromFloat(0)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(-13677.21634375)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, false, limitPrice)
	assert.Equal(t, _0, amount)
}

// amm holds long, trader sells
// case 23: higher than index
func TestComputeAMMAmountWithPrice23(t *testing.T) {
	limitPrice := decimal.NewFromFloat(6976.9161001)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(83941.29865625)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, false, limitPrice)
	assert.Equal(t, _0, amount)
}

// amm holds long, trader sells
// case 24: exactly the best ask/bid price
func TestComputeAMMAmountWithPrice24(t *testing.T) {
	limitPrice := decimal.NewFromFloat(6976.9161)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(83941.29865625)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, false, limitPrice)
	Approximate(t, decimal.NewFromFloat(-1.9954), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	assert.Equal(t, tradingPrice.LessThanOrEqual(limitPrice), true)
}

// amm holds long, trader sells
// case 25: limitPrice is far away
func TestComputeAMMAmountWithPrice25(t *testing.T) {
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

// amm holds long, trader sells
// case 26: normal
func TestComputeAMMAmountWithPrice26(t *testing.T) {
	limitPrice := decimal.NewFromFloat(6800)
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(83941.29865625)
	perpetual1.AmmPositionAmount = decimal.NewFromFloat(2.3)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual1
	amount := ComputeAMMAmountWithPrice(poolStorage, TEST_PERPETUAL_INDEX0, false, limitPrice)
	Approximate(t, decimal.NewFromFloat(-52.542857142857142857), amount)
	tradingPrice, _ := ComputeAMMTrade(poolStorage, TEST_PERPETUAL_INDEX0, accountStorage1, amount)
	Approximate(t, tradingPrice, limitPrice)
}

// open 0 -> -x
func TestComputeBestAskBidPrice0(t *testing.T) {
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(10000)
	perpetual0.AmmPositionAmount = decimal.NewFromFloat(0)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual0
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX1] = perpetual0

	bestPrice := ComputeBestAskBidPrice(poolStorage, TEST_PERPETUAL_INDEX0, false)
	Approximate(t, bestPrice, decimal.NewFromFloat(100.1))
}

// open -10
func TestComputeBestAskBidPrice1(t *testing.T) {
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(10100)
	perpetual0.AmmPositionAmount = decimal.NewFromFloat(-10)
	perpetual2.AmmPositionAmount = decimal.NewFromFloat(10)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual0
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX1] = perpetual2

	bestPrice := ComputeBestAskBidPrice(poolStorage, TEST_PERPETUAL_INDEX0, false)
	Approximate(t, bestPrice, decimal.NewFromFloat(110.11))
}

// open 0 -> +x
func TestComputeBestAskBidPrice2(t *testing.T) {
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(10000)
	perpetual0.AmmPositionAmount = decimal.NewFromFloat(0)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual0
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX1] = perpetual0

	bestPrice := ComputeBestAskBidPrice(poolStorage, TEST_PERPETUAL_INDEX0, true)
	Approximate(t, bestPrice, decimal.NewFromFloat(99.9))
}

// open 10
func TestComputeBestAskBidPrice3(t *testing.T) {
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(8138)
	perpetual0.AmmPositionAmount = decimal.NewFromFloat(10)
	perpetual2.AmmPositionAmount = decimal.NewFromFloat(10)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual0
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX1] = perpetual2

	bestPrice := ComputeBestAskBidPrice(poolStorage, TEST_PERPETUAL_INDEX0, true)
	Approximate(t, bestPrice, decimal.NewFromFloat(89.91))
}

// close -10
func TestComputeBestAskBidPrice4(t *testing.T) {
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(10100)
	perpetual0.AmmPositionAmount = decimal.NewFromFloat(-10)
	perpetual2.AmmPositionAmount = decimal.NewFromFloat(10)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual0
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX1] = perpetual2

	bestPrice := ComputeBestAskBidPrice(poolStorage, TEST_PERPETUAL_INDEX0, true)
	Approximate(t, bestPrice, decimal.NewFromFloat(108.88646369499801395463383186703))
}

// close 10
func TestComputeBestAskBidPrice5(t *testing.T) {
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(8138)
	perpetual0.AmmPositionAmount = decimal.NewFromFloat(10)
	perpetual2.AmmPositionAmount = decimal.NewFromFloat(10)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual0
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX1] = perpetual2

	bestPrice := ComputeBestAskBidPrice(poolStorage, TEST_PERPETUAL_INDEX0, false)
	Approximate(t, bestPrice, decimal.NewFromFloat(91.09554538669368171312465896007))
}

// close unsafe -10
func TestComputeBestAskBidPrice6(t *testing.T) {
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(17692)
	perpetual0.AmmPositionAmount = decimal.NewFromFloat(-80)
	perpetual2.AmmPositionAmount = decimal.NewFromFloat(10)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual0
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX1] = perpetual2

	bestPrice := ComputeBestAskBidPrice(poolStorage, TEST_PERPETUAL_INDEX0, true)
	Approximate(t, bestPrice, decimal.NewFromFloat(100))
}

// close unsafe 10
func TestComputeBestAskBidPrice7(t *testing.T) {
	poolStorage := defaultPool
	poolStorage.PoolCashBalance = decimal.NewFromFloat(1996)
	perpetual0.AmmPositionAmount = decimal.NewFromFloat(80)
	perpetual2.AmmPositionAmount = decimal.NewFromFloat(10)
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX0] = perpetual0
	poolStorage.Perpetuals[TEST_PERPETUAL_INDEX1] = perpetual2

	bestPrice := ComputeBestAskBidPrice(poolStorage, TEST_PERPETUAL_INDEX0, false)
	Approximate(t, bestPrice, decimal.NewFromFloat(100))
}
