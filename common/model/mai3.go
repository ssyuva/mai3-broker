package model

import (
	"github.com/shopspring/decimal"
)

type LiquidityPoolStorage struct {
	VaultFeeRate    decimal.Decimal
	PoolCashBalance decimal.Decimal

	Perpetuals map[int64]*PerpetualStorage
}

const PerpetualNormal = 2

type PerpetualStorage struct {
	IsNormal bool

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
	OpenInterest           decimal.Decimal
	MaxOpenInterestRate    decimal.Decimal // openInterest <= poolMargin * maxOpenInterestRate / indexPrice

	HalfSpread            decimal.Decimal // α
	OpenSlippageFactor    decimal.Decimal // β1
	CloseSlippageFactor   decimal.Decimal // β2
	FundingRateFactor     decimal.Decimal // γ
	FundingRateLimit      decimal.Decimal // Γ
	MaxLeverage           decimal.Decimal // λ
	MaxClosePriceDiscount decimal.Decimal // δ

	AmmCashBalance    decimal.Decimal
	AmmPositionAmount decimal.Decimal
}

type AccountStorage struct {
	TargetLeverage decimal.Decimal
	WalletBalance  decimal.Decimal
	CashBalance    decimal.Decimal
	PositionAmount decimal.Decimal
	// EntryFundingLoss decimal.Decimal
}

type AccountComputed struct {
	PositionValue        decimal.Decimal
	PositionMargin       decimal.Decimal
	MaintenanceMargin    decimal.Decimal
	AvailableCashBalance decimal.Decimal
	MarginBalance        decimal.Decimal
	AvailableMargin      decimal.Decimal
	WithdrawableBalance  decimal.Decimal
	IsMMSafe             bool // use this if check liquidation
	IsIMSafe             bool // use this if open positions
	IsMarginSafe         bool // use this if close positions. also known as bankrupt
}

type AMMTradingContext struct {
	// current trading perpetual
	Index                 decimal.Decimal
	Position1             decimal.Decimal
	HalfSpread            decimal.Decimal
	OpenSlippageFactor    decimal.Decimal
	CloseSlippageFactor   decimal.Decimal
	FundingRateLimit      decimal.Decimal
	MaxClosePriceDiscount decimal.Decimal // δ_m
	MaxLeverage           decimal.Decimal

	// other perpetuals
	OtherIndex              []decimal.Decimal
	OtherPosition           []decimal.Decimal
	OtherOpenSlippageFactor []decimal.Decimal
	OtherMaxLeverage        []decimal.Decimal

	// total
	Cash       decimal.Decimal
	PoolMargin decimal.Decimal

	// trading result
	DeltaMargin     decimal.Decimal
	DeltaPosition   decimal.Decimal
	BestAskBidPrice decimal.Decimal

	// eager evaluation
	ValueWithoutCurrent          decimal.Decimal
	SquareValueWithoutCurrent    decimal.Decimal
	PositionMarginWithoutCurrent decimal.Decimal
}
