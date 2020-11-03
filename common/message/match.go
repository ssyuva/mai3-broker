package message

import (
	"github.com/shopspring/decimal"
)

const MatchTypeNewPerpetual = "newPerpetual"
const MatchTypeNewOrder = "newOrder"
const MatchTypeCancelOrder = "cancelOrder"
const MatchTypeChangeOrder = "reloadOrder"
const MatchTypePerpetualRollBack = "perpetualRollBack"

type MatchMessage struct {
	PerpetualAddress string      `json:"perpetualAddress"`
	Type             string      `json:"type"`
	Payload          interface{} `json:"payload"`
}

type MatchNewOrderPayload struct {
	OrderHash string `json:"orderHash"`
}

type MatchCancelOrderPayload struct {
	OrderHash string `json:"orderHash"`
}

type MatchChangeOrderPayload struct {
	OrderHash string          `json:"orderHash"`
	Amount    decimal.Decimal `json:"amount"`
}
