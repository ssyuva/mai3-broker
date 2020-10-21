package message

const MatchTypeNewPerpetual = "newPerpetual"
const MatchTypeNewOrder = "newOrder"
const MatchTypeCancelOrder = "cancelOrder"
const MatchTypeCancelAll = "cancelAll"
const MatchTypeReloadOrder = "reloadOrder"

type MatchMessage struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type MatchNewPerpetualPayload struct {
	PerpetualAddress string `json:"perpetualAddress"`
}

type MatchNewOrderPayload struct {
	PerpetualAddress string `json:"perpetualAddress"`
	OrderHash        string `json:"orderHash"`
}

type MatchCancelOrderPayload struct {
	PerpetualAddress string `json:"perpetualAddress"`
	OrderHash        string `json:"orderHash"`
}

type MatchCancelAllPayload struct {
	TraderAddress string `json:"traderAddress"`
}

type MatchReloadOrderPayload struct {
	PerpetualAddress string `json:"perpetualAddress"`
	OrderHash        string `json:"orderHash"`
}
