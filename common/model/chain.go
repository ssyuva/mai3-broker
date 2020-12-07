package model

import (
	"github.com/shopspring/decimal"
	"math/big"
	"time"
)

type BlockHeader struct {
	BlockNumber int64
	BlockHash   string
	ParentHash  string
	BlockTime   time.Time
}

type Receipt struct {
	BlockNumber uint64
	BlockHash   string
	GasUsed     uint64
	Status      LaunchTransactionStatus
	BlockTime   uint64
}

type TradeSuccessEvent struct {
	PerpetualAddress string
	BlockNumber      int64
	TransactionSeq   int
	TransactionHash  string
	TraderAddress    string
	OrderHash        string
	Amount           decimal.Decimal
	Gas              decimal.Decimal
}

type PerpetualEvent struct {
	FactoryAddress    string
	BlockNumber       int64
	TransactionSeq    int
	TransactionHash   string
	PerpetualAddress  string
	GovernorAddress   string
	ShareToken        string
	OracleAddress     string
	CollateralAddress string
	OperatorAddress   string

	InitialMarginRate     decimal.Decimal
	MaintenanceMarginRate decimal.Decimal
	OperatorFeeRate       decimal.Decimal
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

type WalletOrderParam struct {
	Trader    string
	Broker    string
	Relayer   string
	Perpetual string
	Referrer  string
	Amount    decimal.Decimal
	Price     decimal.Decimal
	OrderData [32]byte
	ChainID   uint64
	Gas       *big.Int
	Signature []byte
}
