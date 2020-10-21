package websocket

import (
	"errors"
	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/v2/logger"
	"github.com/satori/go.uuid"
	"net"
	// "sync"
	"strings"
	"time"
)

var pingPeriod = 5 * time.Second
var pongPeriod = 15 * time.Second

// For Mock Test
type clientConn interface {
	Close() error
	WriteJSON(interface{}) error
	ReadJSON(interface{}) error
	RemoteAddr() net.Addr
	SetWriteDeadline(t time.Time) error
	SetPongHandler(h func(appData string) error)
	WriteControl(messageType int, data []byte, deadline time.Time) error
}

type Client struct {
	ID             string
	Conn           clientConn
	msgChan        chan interface{}
	timer          *time.Timer
	LoginAddresses map[string]struct{}
	// mu       sync.Mutex
	// for orderbook channel
	Group    int
	Depth    int
	Interval int
}

func (c *Client) AddLoginAddress(address string) {
	c.LoginAddresses[strings.ToLower(address)] = struct{}{}
}

func (c *Client) CheckLogin(address string) bool {
	_, ok := c.LoginAddresses[strings.ToLower(address)]
	return ok
}

func (c *Client) sendData(data interface{}) error {
	// c.mu.Lock()
	// defer c.mu.Unlock()

	c.Conn.SetWriteDeadline(time.Now().Add(time.Second))
	err := c.Conn.WriteJSON(data)
	return err
}

func (c *Client) clientRecover() {
	if r := recover(); r != nil {
		logger.Errorf("client recover! ID: %s, recover:%v", c.ID, r)
	}
}

func (c *Client) SendMsg(msg interface{}) {
	defer c.clientRecover()
	c.msgChan <- msg
}

func (c *Client) runSendMsg() {
	defer c.clientRecover()
	ticker := time.NewTicker(pingPeriod)

	var dummyCloseError *websocket.CloseError
	var dummyOpError *net.OpError
	for {
		select {
		case <-ticker.C:
			if err := c.Conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(time.Second)); err != nil {
				logger.Warnf("Ping Error:%s", err.Error())
				if errors.As(err, &dummyCloseError) || errors.As(err, &dummyOpError) {
					return
				}
			} else {
				if c.timer != nil {
					c.timer.Stop()
				}
				c.timer = time.AfterFunc(pongPeriod, func() {
					logger.Warnf("not receive pong message in %d seconds! clientID:%s close client", pongPeriod/time.Second, c.ID)
					c.Conn.Close()
				})
			}
		case msg, ok := <-c.msgChan:
			if !ok {
				return
			}
			if err := c.sendData(msg); err != nil {
				logger.Errorf("send data error! ID: %s, err:%s", c.ID, err.Error())
				if errors.As(err, &dummyCloseError) || errors.As(err, &dummyOpError) {
					return
				}
			}
		}
	}
}

func (c *Client) Close() {
	defer c.clientRecover()
	if c.timer != nil {
		c.timer.Stop()
	}
	close(c.msgChan)
	c.Conn.Close()
}

func (c *Client) setPongHandler() {
	c.Conn.SetPongHandler(func(message string) error {
		if c.timer != nil {
			c.timer.Stop()
		}
		return nil
	})
}

func NewClient(conn clientConn) *Client {
	client := &Client{
		ID:             uuid.NewV4().String(),
		Conn:           conn,
		msgChan:        make(chan interface{}, 100),
		LoginAddresses: make(map[string]struct{}),
	}
	client.setPongHandler()

	go client.runSendMsg()
	return client
}
