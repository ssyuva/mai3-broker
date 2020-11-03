package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type LaunchTransactionStatus int

const (
	TxInitial  LaunchTransactionStatus = 0
	TxPending  LaunchTransactionStatus = 1
	TxSuccess  LaunchTransactionStatus = 2
	TxFailed   LaunchTransactionStatus = 3
	TxCanceled LaunchTransactionStatus = 4
)

type TransactionType int

const (
	TxNormal      TransactionType = 0
	TxCancel      TransactionType = 1
	TxAccelerated TransactionType = 2
)

type Transaction struct {
	ID              uint64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	TxID            string `gorm:"index"`
	FromAddress     string
	ToAddress       string
	Type            TransactionType
	Method          string
	Inputs          []byte
	BlockNumber     *uint64
	BlockHash       *string
	BlockTime       *uint64
	TransactionHash *string
	Nonce           *uint64
	GasPrice        *uint64
	GasLimit        *uint64
	GasUsed         *uint64
	Status          LaunchTransactionStatus
	Value           decimal.Decimal `sql:"type:numeric(78,18)"`
	CommitTime      time.Time
	UpdateTime      time.Time
}
