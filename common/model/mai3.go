package model

import (
	"github.com/shopspring/decimal"
)

type GovParams struct {
	// perpetual
	InitialMarginRate     decimal.Decimal
	MaintenanceMarginRate decimal.Decimal
	OperatorFeeRate       decimal.Decimal
	VaultFeeRate          decimal.Decimal
	LpFeeRate             decimal.Decimal
	ReferrerRebateRate    decimal.Decimal
	LiquidatorPenaltyRate decimal.Decimal
	KeeperGasReward       decimal.Decimal
	// amm
	HalfSpreadRate         decimal.Decimal
	Beta1                  decimal.Decimal
	Beta2                  decimal.Decimal
	FundingRateCoefficient decimal.Decimal
	TargetLeverage         decimal.Decimal
}

type PerpetualStorage struct {
	IsEmergency                   bool
	IsGlobalSettled               bool
	InsuranceFund1                decimal.Decimal
	InsuranceFund2                decimal.Decimal
	MarkPrice                     decimal.Decimal
	IndexPrice                    decimal.Decimal
	AccumulatedFundingPerContract decimal.Decimal
	FundingTime                   int64
}

type AccountStorage struct {
	CashBalance      decimal.Decimal
	PositionAmount   decimal.Decimal
	EntryFundingLoss decimal.Decimal
	Gas              decimal.Decimal
}

type AccountComputed struct {
	PositionValue        decimal.Decimal
	PositionMargin       decimal.Decimal
	Leverage             decimal.Decimal
	MaintenanceMargin    decimal.Decimal
	FundingLoss          decimal.Decimal
	AvailableCashBalance decimal.Decimal
	MarginBalance        decimal.Decimal
	AvailableMargin      decimal.Decimal
	MaxWithdrawable      decimal.Decimal
	WithdrawableBalance  decimal.Decimal
	IsSafe               bool
}

type TradeCost struct {
	AccountStorage  AccountStorage
	AccountComputed AccountComputed
	MarginCost      decimal.Decimal
	Fee             decimal.Decimal
}

type AMMTradingContext struct {
	Index         decimal.Decimal
	Lev           decimal.Decimal
	Cash          decimal.Decimal
	Pos1          decimal.Decimal
	IsSafe        bool
	M0            decimal.Decimal
	Mv            decimal.Decimal
	Ma1           decimal.Decimal
	DeltaMargin   decimal.Decimal
	DeltaPosition decimal.Decimal
}

type TradingContext struct {
	TakerAccount AccountStorage
	MakerAccount AccountStorage
	LpFee        decimal.Decimal
	VaultFee     decimal.Decimal
	OperatorFee  decimal.Decimal
	TradingPrice decimal.Decimal
}
