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
	CancelReasonTooManyMatches               CancelReasonType = "TOO_MANY_MATCHES"
	CancelReasonInternalError                CancelReasonType = "INTERNAL_ERROR"
	CancelReasonInsufficientFunds            CancelReasonType = "INSUFFICIENT_FUNDS"
	CancelReasonLongPriceTooHighAfterExpired CancelReasonType = "LONG_PRICE_TOO_HIGH_AFTER_EXPIRED"
	CancelReasonShortPriceTooLowAfetrExpired CancelReasonType = "SHORT_PRICE_TOO_LOW_AFTER_EXPIRED"
	CancelReasonContractSettled              CancelReasonType = "CONTRACT_SETTLED"
)

type OrderType string

const (
	LimitOrder      OrderType = "limit"
	MarketOrder     OrderType = "market"
	StopLimitOrder  OrderType = "stop-limit"
	StopMarketOrder OrderType = "stop-market"
)

type OrderStatus string

const (
	OrderCanceled      OrderStatus = "canceled"
	OrderPending       OrderStatus = "pending"
	OrderPartialFilled OrderStatus = "partial_filled"
	OrderFullFilled    OrderStatus = "full_filled"
	OrderStop          OrderStatus = "stop"
)

type OrderSide string

const (
	SideBuy  OrderSide = "buy"
	SideSell OrderSide = "sell"
)

type OrderParam struct {
	TraderAddress string          `json:"traderAddress" db:"trader_address"`
	Type          OrderType       `json:"type" db:"type"`
	Side          OrderSide       `json:"side" db:"side"`
	Price         decimal.Decimal `json:"price" db:"price"`
	Amount        decimal.Decimal `json:"amount" db:"amount"`
	StopPrice     decimal.Decimal `json:"stopPrice" db:"stopPrice"`
	ExpiresAt     time.Time       `json:"expiresAt" db:"expires_at"`
	Salt          int64           `json:"salt" db:"salt"`
	ChainID       int64           `json:"chainID" db:"chain_id"`
}

type Order struct {
	ID int64 `json:"id"               db:"id"               primaryKey:"true" gorm:"primary_key"`
	OrderParam
	OrderHash        string               `json:"orderHash" db:"order_hash"`
	PerpetualAddress string               `json:"perpetualAddress" db:"perpetual_address"`
	OldStatus        OrderStatus          `json:"oldStatus" sql:"-"`
	Status           OrderStatus          `json:"status" db:"status"`
	AvailableAmount  decimal.Decimal      `json:"availableAmount" db:"available_amount"`
	ConfirmedAmount  decimal.Decimal      `json:"confirmedAmount" db:"confirmed_amount"`
	ConfirmedVolume  decimal.Decimal      `json:"confirmedVolume" db:"confirmed_volume"`
	FilledPrice      decimal.Decimal      `json:"filledPrice" db:"filled_price"`
	CanceledAmount   decimal.Decimal      `json:"canceledAmount" db:"canceled_amount"`
	PendingAmount    decimal.Decimal      `json:"pendingAmount" db:"pending_amount"`
	PendingVolume    decimal.Decimal      `json:"pendingVolume" db:"pending_volume"`
	GasFeeAmount     decimal.Decimal      `json:"gasFeeAmount" db:"gas_fee_amount"`
	CreatedAt        time.Time            `json:"createdAt" db:"created_at"`
	UpdatedAt        time.Time            `json:"updatedAt" db:"updated_at"`
	Signature        string               `json:"-" db:"signature"`
	CancelReasons    []*OrderCancelReason `json:"cancelReasons" sql:"-"`
}

//OrderCancelReason records of cancel
type OrderCancelReason struct {
	Reason          CancelReasonType `json:"reason"`
	Amount          decimal.Decimal  `json:"amount"`
	CanceledAt      time.Time        `json:"canceledAt"`
	TransactionHash string           `json:"transactionHash,omitempty"`
}
