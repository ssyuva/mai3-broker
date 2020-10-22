package message

import (
	"fmt"
)

const WsTypeOrderChange = "orderChange"

type WebSocketMessage struct {
	ChannelID string      `json:"channel_id"`
	Payload   interface{} `json:"payload"`
}

type WebSocketOrderChangePayload struct {
	Type  string      `json:"type"`
	Order interface{} `json:"order"`
}

const AccountChannelPrefix = "TraderAddress"

func GetAccountChannelID(address string) string {
	return fmt.Sprintf("%s#%s", AccountChannelPrefix, address)
}
