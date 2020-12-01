package model

import (
	"github.com/shopspring/decimal"
	"time"
)

const QueryParamSeperator = ","

type CancelReasonType string

const (
	CancelReasonExpired                      CancelReasonType = "EXPIRED"
	CancelReasonAdminCancel                  CancelReasonType = "CANCELED_BY_ADMIN"
	CancelReasonUserCancel                   CancelReasonType = "CANCELED_BY_USER"
	CancelReasonTransactionFail              CancelReasonType = "TRANSACTION_FAIL"
	CancelReasonRemainTooSmall               CancelReasonType = "REMAIN_TOO_SMALL"
	CancelReasonMatchTooSmall                CancelReasonType = "MATCH_TOO_SMALL"
	CancelReasonInternalError                CancelReasonType = "INTERNAL_ERROR"
	CancelReasonInsufficientFunds            CancelReasonType = "INSUFFICIENT_FUNDS"
	CancelReasonLongPriceTooHighAfterExpired CancelReasonType = "LONG_PRICE_TOO_HIGH_AFTER_EXPIRED"
	CancelReasonShortPriceTooLowAfetrExpired CancelReasonType = "SHORT_PRICE_TOO_LOW_AFTER_EXPIRED"
	CancelReasonContractSettled              CancelReasonType = "CONTRACT_SETTLED"
)

type OrderType int

const (
	LimitOrder     OrderType = 1
	StopLimitOrder OrderType = 2
)

type OrderStatus string

const (
	OrderCanceled      OrderStatus = "canceled"
	OrderPending       OrderStatus = "pending"
	OrderPartialFilled OrderStatus = "partial_filled"
	OrderFullFilled    OrderStatus = "full_filled"
	OrderStop          OrderStatus = "stop"
)

type OrderParam struct {
	TraderAddress    string          `json:"traderAddress" db:"trader_address"`
	PerpetualAddress string          `json:"perpetualAddress" db:"perpetual_address"`
	RelayerAddress   string          `json:"relayerAddress" db:"relayer_address"`
	BrokerAddress    string          `json:"brokerAddress" db:"broker_address"`
	ReferrerAddress  string          `json:"referrerAddress" db:"referrer_address"`
	Type             OrderType       `json:"type" db:"type"`
	Price            decimal.Decimal `json:"price" db:"price"`
	Amount           decimal.Decimal `json:"amount" db:"amount"`
	StopPrice        decimal.Decimal `json:"stopPrice" db:"stopPrice"`
	ExpiresAt        time.Time       `json:"expiresAt" db:"expires_at"`
	Version          int32           `json:"version" db:"version"`
	Salt             int64           `json:"salt" db:"salt"`
	IsCloseOnly      bool            `json:"isCloseOnly" db:"is_close_only"`
	ChainID          int64           `json:"chainID" db:"chain_id"`
}

type Order struct {
	OrderParam
	ID              int64                `json:"-" db:"id" primaryKey:"true" gorm:"primary_key"`
	OrderHash       string               `json:"orderHash" db:"order_hash"`
	OldStatus       OrderStatus          `json:"oldStatus" sql:"-"`
	Status          OrderStatus          `json:"status" db:"status"`
	AvailableAmount decimal.Decimal      `json:"availableAmount" db:"available_amount"`
	ConfirmedAmount decimal.Decimal      `json:"confirmedAmount" db:"confirmed_amount"`
	FilledPrice     decimal.Decimal      `json:"filledPrice" db:"filled_price"`
	CanceledAmount  decimal.Decimal      `json:"canceledAmount" db:"canceled_amount"`
	PendingAmount   decimal.Decimal      `json:"pendingAmount" db:"pending_amount"`
	GasFeeAmount    decimal.Decimal      `json:"gasFeeAmount" db:"gas_fee_amount"`
	CreatedAt       time.Time            `json:"createdAt" db:"created_at"`
	UpdatedAt       time.Time            `json:"updatedAt" db:"updated_at"`
	Signature       string               `json:"-" db:"signature"`
	CancelReasons   []*OrderCancelReason `json:"cancelReasons" sql:"-"`
}

//OrderCancelReason records of cancel
type OrderCancelReason struct {
	Reason          CancelReasonType `json:"reason"`
	Amount          decimal.Decimal  `json:"amount"`
	CanceledAt      time.Time        `json:"canceledAt"`
	TransactionHash string           `json:"transactionHash,omitempty"`
}
