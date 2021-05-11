package mai3

import (
	"testing"

	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

var defaultPool1 = &model.LiquidityPoolStorage{
	VaultFeeRate:    decimal.NewFromFloat(0.001),
	PoolCashBalance: decimal.NewFromFloat(1000), // set me later

	Perpetuals: make(map[int64]*model.PerpetualStorage), // set me later
}

var perpetual3 = &model.PerpetualStorage{
	IsNormal: true,

	MarkPrice:               decimal.NewFromFloat(1000),
	IndexPrice:              decimal.NewFromFloat(1000),
	UnitAccumulativeFunding: decimal.NewFromFloat(0),

	InitialMarginRate:      decimal.NewFromFloat(0.01),
	MaintenanceMarginRate:  decimal.NewFromFloat(0.005),
	OperatorFeeRate:        decimal.NewFromFloat(0.001),
	LpFeeRate:              decimal.NewFromFloat(0.001),
	ReferrerRebateRate:     decimal.NewFromFloat(0.2),
	LiquidationPenaltyRate: decimal.NewFromFloat(0.002),
	KeeperGasReward:        decimal.NewFromFloat(0.5),
	InsuranceFundRate:      decimal.NewFromFloat(0.5),
	OpenInterest:           decimal.NewFromFloat(0),
	MaxOpenInterestRate:    decimal.NewFromFloat(4),

	HalfSpread:            decimal.NewFromFloat(0.01),
	OpenSlippageFactor:    decimal.NewFromFloat(0.1),
	CloseSlippageFactor:   decimal.NewFromFloat(0.06),
	FundingRateLimit:      decimal.NewFromFloat(0),
	MaxLeverage:           decimal.NewFromFloat(5),
	MaxClosePriceDiscount: decimal.NewFromFloat(0.05),

	AmmCashBalance:    _0, // assign me later
	AmmPositionAmount: _0, // assign me later
}

var perpetual4 = &model.PerpetualStorage{
	IsNormal: true,

	MarkPrice:               decimal.NewFromFloat(1000),
	IndexPrice:              decimal.NewFromFloat(1000),
	UnitAccumulativeFunding: decimal.NewFromFloat(0),

	InitialMarginRate:      decimal.NewFromFloat(0.01),
	MaintenanceMarginRate:  decimal.NewFromFloat(0.005),
	OperatorFeeRate:        decimal.NewFromFloat(0.001),
	LpFeeRate:              decimal.NewFromFloat(0.001),
	ReferrerRebateRate:     decimal.NewFromFloat(0.2),
	LiquidationPenaltyRate: decimal.NewFromFloat(0.002),
	KeeperGasReward:        decimal.NewFromFloat(0.5),
	InsuranceFundRate:      decimal.NewFromFloat(0.5),
	OpenInterest:           decimal.NewFromFloat(0),
	MaxOpenInterestRate:    decimal.NewFromFloat(4),

	HalfSpread:            decimal.NewFromFloat(0.01),
	OpenSlippageFactor:    decimal.NewFromFloat(0.1),
	CloseSlippageFactor:   decimal.NewFromFloat(0.06),
	FundingRateLimit:      decimal.NewFromFloat(0),
	MaxLeverage:           decimal.NewFromFloat(5),
	MaxClosePriceDiscount: decimal.NewFromFloat(0.05),

	AmmCashBalance:    _0, // assign me later
	AmmPositionAmount: _0, // assign me later
}

var accountStorage = &model.AccountStorage{
	CashBalance:    decimal.NewFromFloat(0),
	PositionAmount: decimal.NewFromFloat(0),
	TargetLeverage: _2,
}

//addLiq + tradeWithLev long 3, short 2, short 2, long 1
func TestComputeAMMTrade1(t *testing.T) {
	poolStorage := defaultPool1
	poolStorage.Perpetuals[0] = perpetual3
	poolStorage.Perpetuals[1] = perpetual4
	poolStorage = CopyLiquidityPoolStorage(poolStorage)
	account := CopyAccountStorage(accountStorage)
	account.WalletBalance = decimal.NewFromFloat(1960.35) // for ajust margin
	// long 3 (open)
	afterTrade, tradeIsSafe, tradingPrice, _ := ComputeAMMTrade(poolStorage, 0, account, decimal.NewFromFloat(3))
	// amm deltaCash = 3450
	// margin = cash + positionValue = | positionValue | / 2xLev. so cash = -1500
	// cash = deposit - 3450 - 3450 * 0.003(fee). so deposit = 1960.35
	Approximate(t, decimal.NewFromFloat(1150), tradingPrice)
	assert.Equal(t, true, afterTrade.IsMMSafe)
	assert.Equal(t, true, tradeIsSafe)
	Approximate(t, decimal.NewFromFloat(0), account.WalletBalance)
	Approximate(t, decimal.NewFromFloat(-1500), account.CashBalance)
	Approximate(t, decimal.NewFromFloat(3), account.PositionAmount)
	Approximate(t, decimal.NewFromFloat(1500), afterTrade.MarginBalance)
	Approximate(t, decimal.NewFromFloat(4453.45), poolStorage.PoolCashBalance)
	Approximate(t, decimal.NewFromFloat(-3), poolStorage.Perpetuals[0].AmmPositionAmount)

	// short 2 (partial close)
	afterTrade, tradeIsSafe, tradingPrice, _ = ComputeAMMTrade(poolStorage, 0, account, decimal.NewFromFloat(-2))
	// amm deltaCash = -2100
	// margin = cash + positionValue = | positionValue | / 2xLev. so cash = -500
	// newCash = oldCash - withdraw + 2100 - 2100 * 0.003(fee). so withdraw = 1093.7
	Approximate(t, decimal.NewFromFloat(1050), tradingPrice)
	assert.Equal(t, true, afterTrade.IsMMSafe)
	assert.Equal(t, true, tradeIsSafe)
	Approximate(t, decimal.NewFromFloat(1093.7), account.WalletBalance)
	Approximate(t, decimal.NewFromFloat(-500), account.CashBalance)
	Approximate(t, decimal.NewFromFloat(1), account.PositionAmount)
	// AMM rebalance. margin = 1000 * 1 * 1% = 10
	// amm cash + mark pos. so cash = 10 + 1000 * 1
	// final transferFee, cash += 2100 * 0.001(fee)
	Approximate(t, decimal.NewFromFloat(2355.55), poolStorage.PoolCashBalance)
	Approximate(t, decimal.NewFromFloat(-1), poolStorage.Perpetuals[0].AmmPositionAmount)

	// short 2 (close all + open)
	afterTrade, tradeIsSafe, tradingPrice, _ = ComputeAMMTrade(poolStorage, 0, account, decimal.NewFromFloat(-2))
	// amm deltaCash = -1984.996757074682502
	// margin = cash + positionValue = | positionValue | / 2xLev. so cash = 1500
	// idealMargin = oldCash + deltaCash + deposit - fee + mark newPos.
	// so deposit = 500 - (-500) - (1984...) + 1984... * 0.003 - (-1000) = 20.958233196541545506
	Approximate(t, decimal.NewFromFloat(992.498378537341251), tradingPrice)
	assert.Equal(t, true, afterTrade.IsMMSafe)
	assert.Equal(t, true, tradeIsSafe)
	Approximate(t, decimal.NewFromFloat(1093.7-20.958233196541545506), account.WalletBalance)
	Approximate(t, decimal.NewFromFloat(1500), account.CashBalance)
	Approximate(t, decimal.NewFromFloat(-1), account.PositionAmount)
	Approximate(t, decimal.NewFromFloat(372.53823968239218050), poolStorage.PoolCashBalance)
	Approximate(t, decimal.NewFromFloat(1), poolStorage.Perpetuals[0].AmmPositionAmount)

	// long 1 (close all)
	afterTrade, tradeIsSafe, tradingPrice, _ = ComputeAMMTrade(poolStorage, 0, account, decimal.NewFromFloat(1))
	// amm deltaCash = 977.783065493367778
	Approximate(t, decimal.NewFromFloat(977.783065493367778), tradingPrice)
	assert.Equal(t, true, afterTrade.IsMMSafe)
	assert.Equal(t, true, tradeIsSafe)
	Approximate(t, decimal.NewFromFloat(1093.7-20.958233196541545506+519.2835853101521187), account.WalletBalance)
	Approximate(t, decimal.NewFromFloat(0), account.CashBalance)
	Approximate(t, decimal.NewFromFloat(0), account.PositionAmount)
	// AMM rebalance. margin = 1000 * 1 * 1% = 10
	// amm cash + mark pos. so cash = 10 + 1000 * 1
	// final transferFee, cash += 2100 * 0.001(fee)
	Approximate(t, decimal.NewFromFloat(1351.29908824125332628), poolStorage.PoolCashBalance)
	Approximate(t, decimal.NewFromFloat(0), poolStorage.Perpetuals[0].AmmPositionAmount)
}

// deposit + long 3(auto deposit on demand)
func TestComputeAMMTrade2(t *testing.T) {
	poolStorage := defaultPool1
	poolStorage.Perpetuals[0] = perpetual3
	poolStorage.Perpetuals[1] = perpetual4
	poolStorage = CopyLiquidityPoolStorage(poolStorage)
	// deposit
	account := CopyAccountStorage(accountStorage)
	account.CashBalance = decimal.NewFromFloat(500)
	account.WalletBalance = decimal.NewFromFloat(1460.35)

	// long 3 (open)
	afterTrade, tradeIsSafe, tradingPrice, _ := ComputeAMMTrade(poolStorage, 0, account, decimal.NewFromFloat(3))
	// amm deltaCash = 3450
	// margin = cash + positionValue = | positionValue | / 2xLev. so cash = -1500
	// newCash = oldCash + deposit - 3450 - 3450 * 0.003(fee). so deposit = 1460.35
	Approximate(t, decimal.NewFromFloat(1150), tradingPrice)
	assert.Equal(t, true, afterTrade.IsMMSafe)
	assert.Equal(t, true, tradeIsSafe)
	Approximate(t, decimal.NewFromFloat(0), account.WalletBalance)
	Approximate(t, decimal.NewFromFloat(-1500), account.CashBalance)
	Approximate(t, decimal.NewFromFloat(3), account.PositionAmount)
	Approximate(t, decimal.NewFromFloat(1500), afterTrade.MarginBalance)

	Approximate(t, decimal.NewFromFloat(4453.45), poolStorage.PoolCashBalance)
	Approximate(t, decimal.NewFromFloat(-3), poolStorage.Perpetuals[0].AmmPositionAmount)
}

// short 1 when MM < margin < IM, normal fees
func TestComputeAMMTrade3(t *testing.T) {
	poolStorage := defaultPool1
	poolStorage.Perpetuals[0] = perpetual3
	poolStorage.Perpetuals[1] = perpetual4
	poolStorage = CopyLiquidityPoolStorage(poolStorage)
	// deposit
	account := CopyAccountStorage(accountStorage)
	account.WalletBalance = decimal.NewFromFloat(1460.35)
	// long 3 (open)
	afterTrade, _, _, _ := ComputeAMMTrade(poolStorage, 0, account, decimal.NewFromFloat(3))
	// close when MM < margin < IM, normal fees
	Approximate(t, decimal.NewFromFloat(505), poolStorage.Perpetuals[0].MarkPrice)
	Approximate(t, decimal.NewFromFloat(505), poolStorage.Perpetuals[0].IndexPrice)
	assert.Equal(t, false, afterTrade.IsMMSafe)

}
