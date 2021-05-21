package websocket

import "github.com/mcdexio/mai3-broker/common/message"

type IChannel interface {
	GetID() string

	// Thread safe calls
	AddSubscriber(*Client)
	RemoveSubscriber(string)
	AddMessage(message *message.WebSocketMessage)

	UnsubscribeChan() chan string
	SubScribeChan() chan *Client
	MessagesChan() chan *message.WebSocketMessage

	handleMessage(*message.WebSocketMessage)
	handleSubscriber(*Client)
	handleUnsubscriber(string)
}
