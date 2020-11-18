package model

import (
	"github.com/shopspring/decimal"
	"math/big"
	"time"
)

type AccountStorage struct {
	CashBalance      decimal.Decimal
	Position         decimal.Decimal
	EntrySocialLoss  decimal.Decimal
	EntryFundingLoss decimal.Decimal
}

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

type MatchEvent struct {
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
	FactoryAddress   string
	BlockNumber      int64
	TransactionSeq   int
	TransactionHash  string
	PerpetualAddress string
	OracleAddress    string
	OperatorAddress  string
}

type PerpetualStorage struct {
	Leverage        decimal.Decimal
	Cash            decimal.Decimal
	Position        decimal.Decimal
	IndexPrice      decimal.Decimal
	Beta            decimal.Decimal
	Beta2           decimal.Decimal
	Fee             decimal.Decimal
	Spread          decimal.Decimal
	Gama            decimal.Decimal
	LastFundingTime int64
	FundingRate     decimal.Decimal
	IndexBuffer     decimal.Decimal
}

type WalletSignature struct {
	Config [32]byte
	R      [32]byte
	S      [32]byte
}

type WalletOrderParam struct {
	Trader    string
	Broker    string
	Perpetual string
	Price     decimal.Decimal
	Amount    decimal.Decimal
	ExpiredAt uint64
	Version   uint32
	Category  uint8
	CloseOnly bool
	Salt      uint64
	ChainID   uint64
	Gas       *big.Int
	Signature WalletSignature
}
