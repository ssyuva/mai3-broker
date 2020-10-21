package message

const WsTypeOrderChange = "orderChange"

type WebSocketMessage struct {
	ChannelID string      `json:"channel_id"`
	Payload   interface{} `json:"payload"`
}

type WebSocketOrderChangePayload struct {
	Type  string      `json:"type"`
	Order interface{} `json:"order"`
}
