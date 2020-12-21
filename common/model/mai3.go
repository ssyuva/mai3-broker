package model

import (
	"github.com/shopspring/decimal"
)

type LiquidityPoolStorage struct {
	VaultFeeRate         decimal.Decimal
	InsuranceFundCap     decimal.Decimal
	InsuranceFund        decimal.Decimal
	DonatedInsuranceFund decimal.Decimal
	TotalClaimableFee    decimal.Decimal
	PoolCashBalance      decimal.Decimal
	FundingTime          int64

	Perpetuals map[int64]*PerpetualStorage
}

type PerpetualStorage struct {
	IsEmergency     bool
	IsGlobalSettled bool

	MarkPrice               decimal.Decimal
	IndexPrice              decimal.Decimal
	UnitAccumulativeFunding decimal.Decimal

	InitialMarginRate      decimal.Decimal
	MaintenanceMarginRate  decimal.Decimal
	OperatorFeeRate        decimal.Decimal
	LpFeeRate              decimal.Decimal
	ReferrerRebateRate     decimal.Decimal
	LiquidationPenaltyRate decimal.Decimal
	KeeperGasReward        decimal.Decimal
	InsuranceFundRate      decimal.Decimal

	HalfSpread          decimal.Decimal // α
	OpenSlippageFactor  decimal.Decimal // β1
	CloseSlippageFactor decimal.Decimal // β2
	FundingRateLimit    decimal.Decimal // γ
	MaxLeverage         decimal.Decimal // λ

	AmmPositionAmount decimal.Decimal
}

type AccountStorage struct {
	CashBalance    decimal.Decimal
	PositionAmount decimal.Decimal
	// EntryFundingLoss decimal.Decimal
}

// type AccountComputed struct {
// 	PositionValue        decimal.Decimal
// 	PositionMargin       decimal.Decimal
// 	MaintenanceMargin    decimal.Decimal
// 	AvailableCashBalance decimal.Decimal
// 	MarginBalance        decimal.Decimal
// 	AvailableMargin      decimal.Decimal
// 	MaxWithdrawable      decimal.Decimal
// 	WithdrawableBalance  decimal.Decimal
// 	IsSafe               bool
// 	Leverage             decimal.Decimal
// }

type AMMTradingContext struct {
	// current trading perpetual
	Index               decimal.Decimal
	Position1           decimal.Decimal
	HalfSpread          decimal.Decimal
	OpenSlippageFactor  decimal.Decimal
	CloseSlippageFactor decimal.Decimal
	FundingRateLimit    decimal.Decimal
	MaxLeverage         decimal.Decimal

	// other perpetuals
	OtherIndex                  []decimal.Decimal
	OtherPosition               []decimal.Decimal
	OtherHalfSpread             []decimal.Decimal
	OtherOpenSlippageFactor     []decimal.Decimal
	OtherCloseSlippageFactor    []decimal.Decimal
	OtherFundingRateCoefficient []decimal.Decimal
	OtherMaxLeverage            []decimal.Decimal

	// total
	Cash       decimal.Decimal
	PoolMargin decimal.Decimal

	// trading result
	DeltaMargin   decimal.Decimal
	DeltaPosition decimal.Decimal

	// eager evaluation
	ValueWithoutCurrent          decimal.Decimal
	SquareValueWithoutCurrent    decimal.Decimal
	PositionMarginWithoutCurrent decimal.Decimal
}
