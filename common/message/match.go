package message

const MatchTypeNewPerpetual = "newPerpetual"
const MatchTypeNewOrder = "newOrder"
const MatchTypeCancelOrder = "cancelOrder"
const MatchTypeReloadOrder = "reloadOrder"

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

type MatchReloadOrderPayload struct {
	OrderHash string `json:"orderHash"`
}
