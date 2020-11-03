package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type AccountStorage struct {
	CashBalance      decimal.Decimal
	Position         decimal.Decimal
	EntrySocialLoss  decimal.Decimal
	EntryFundingLoss decimal.Decimal
}

type BlockHeader struct {
	BlockNumber int
	BlockHash   string
	ParentHash  string
	BlockTime   time.Time
}

type MatchEvent struct {
	PerpetualAddress string
	BlockNumber      int
	TransactionSeq   int
	TransactionHash  string
	TraderAddress    string
	OrderHash        string
	Amount           decimal.Decimal
	Gas              decimal.Decimal
}

type PerpetualEvent struct {
	FactoryAddress   string
	BlockNumber      int
	TransactionSeq   int
	TransactionHash  string
	PerpetualAddress string
	OracleAddress    string
	OperatorAddress  string
}
