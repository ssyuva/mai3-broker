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

func (ts LaunchTransactionStatus) String() string {
	switch ts {
	case TxInitial:
		return "INITIAL"
	case TxPending:
		return "PENDING"
	case TxSuccess:
		return "SUCCESS"
	case TxFailed:
		return "FAILED"
	case TxCanceled:
		return "CANCELED"
	}
	return "UNKNOWN"
}

func (ts LaunchTransactionStatus) TransactionStatus() TransactionStatus {
	switch ts {
	case TxInitial:
		return TransactionStatusInit
	case TxPending:
		return TransactionStatusPending
	case TxSuccess:
		return TransactionStatusSuccess
	case TxFailed:
		return TransactionStatusExcFail
	case TxCanceled:
		return TransactionStatusAbort
	default:
		return TransactionStatus("UNKNOWN")
	}
}

type TransactionType int

const (
	TxNormal      TransactionType = 0
	TxCancel      TransactionType = 1
	TxAccelerated TransactionType = 2
)

type LaunchTransaction struct {
	ID              uint64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	TxID            string `gorm:"index"`
	FromAddress     string
	ToAddress       string
	Type            TransactionType
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
	Value           decimal.Decimal
	CommitTime      time.Time
	UpdateTime      time.Time
}

func (LaunchTransaction) TableName() string {
	return "launch_transactions"
}

func Uint64(n uint64) *uint64 {
	ni := new(uint64)
	*ni = n
	return ni
}

func String(s string) *string {
	ns := new(string)
	*ns = s
	return ns
}

func MustString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func MustUint64(n *uint64) uint64 {
	if n != nil {
		return *n
	}
	return 0
}
