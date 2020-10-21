package websocket

type IChannel interface {
	GetID() string

	// Thread safe calls
	AddSubscriber(*Client)
	RemoveSubscriber(string)
	AddMessage(message *WebSocketMessage)

	UnsubscribeChan() chan string
	SubScribeChan() chan *Client
	MessagesChan() chan *WebSocketMessage

	handleMessage(*WebSocketMessage)
	handleSubscriber(*Client)
	handleUnsubscriber(string)
}
