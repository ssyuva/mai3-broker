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
	FactoryAddress   string
	BlockNumber      int64
	TransactionSeq   int
	TransactionHash  string
	PerpetualAddress string
	OracleAddress    string
	OperatorAddress  string
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
