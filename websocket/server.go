package websocket

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mcarloai/mai-v3-broker/common/auth"
	"github.com/mcarloai/mai-v3-broker/common/message"
	"github.com/mcarloai/mai-v3-broker/conf"
	logger "github.com/sirupsen/logrus"
)

type channelCreator = func(channelID string) (IChannel, error)

var channelCreators = make(map[string]channelCreator)

func RegisterChannelCreator(prefix string, fn channelCreator) {
	channelCreators[prefix] = fn
}

type ClientRequest struct {
	Type     string   `json:"type"`
	Jwt      string   `json:"jwt"`
	MaiToken string   `json:"maiAuth"`
	Channels []string `json:"channels"`
}

var (
	pongResp = struct {
		Type string `json:"type"`
	}{
		"pong",
	}
)

type loginResponse struct {
	Type string `json:"type"`
	Code int    `json:"code"`
	Desc string `json:"desc"`
}

type subScribeResponse struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Code    int    `json:"code"`
	Desc    string `json:"desc"`
}

func loginAuth(jwt, maiToken string) (trader string, err error) {
	if jwt != "" {
		address, err := auth.ValidateJwt(jwt)
		if err != nil {
			return "", fmt.Errorf("JWT-Authentication fail:%w", err)
		}
		return address, nil
	}

	if maiToken != "" {
		address, err := auth.ValidateMaiAuth(maiToken)
		if err != nil {
			return "", fmt.Errorf("MAI-Authentication fail:%w", err)
		}
		return address, nil
	}

	return "", errors.New("jwt or mai token needed")
}

func handleClientRequest(client *Client) {
	logger.Infof("New Client(%s) IP:(%s) Connect", client.ID, client.Conn.RemoteAddr())
	defer func() {
		client.Close()
		RemoveSubscriberAllChannels(client)
		logger.Infof("Client(%s) IP:(%s) Disconnect", client.ID, client.Conn.RemoteAddr())
	}()

	for {
		var req ClientRequest

		err := client.Conn.ReadJSON(&req)

		var dummyJSONError *json.SyntaxError
		var dummyCloseError *websocket.CloseError
		var dummyOpError *net.OpError
		if err == nil {
		} else if errors.As(err, &dummyJSONError) {
			logger.Debugf("handleClientRequest:user request is not json")
			continue
		} else if errors.As(err, &dummyCloseError) {
			logger.Debugf("handleClientRequest:closed by peer")
			return
		} else if errors.As(err, &dummyOpError) {
			logger.Debugf("handleClientRequest:closed by server")
			return
		} else {
			logger.Errorf("FIXME:handleClientRequest:unkown err:%s", err.Error())
			return
		}

		switch req.Type {
		case "ping":
			client.SendMsg(pongResp)
		case "login":
			loginResp := loginResponse{Type: "login", Code: 0, Desc: "login success"}
			trader, err := loginAuth(req.Jwt, req.MaiToken)
			if err != nil {
				logger.Errorf("login auth err:%s", err.Error())
				loginResp.Code = -1
				loginResp.Desc = err.Error()
			} else {
				client.AddLoginAddress(trader)
			}
			client.SendMsg(loginResp)

		case "subscribe":
			for _, id := range req.Channels {
				// TraderAddress check login
				if strings.HasPrefix(id, message.AccountChannelPrefix) {
					parts := strings.Split(id, "#")
					if len(parts) != 2 || !client.CheckLogin(parts[1]) {
						client.SendMsg(subScribeResponse{Type: "subscribeError", Channel: id, Code: -1, Desc: "need login"})
						continue
					}
				}
				channel := findChannel(id)
				if channel == nil {
					// There is a risk to let user create channel freely.
					channel, err = createChannelByID(id)
					if err != nil {
						logger.Errorf("handleClientRequest:create channel fail:%s", err.Error())
					}
				}
				if channel != nil {
					channel.AddSubscriber(client)
				}
			}
		case "unsubscribe":
			for _, id := range req.Channels {
				channel := findChannel(id)
				if channel == nil {
					continue
				}
				channel.RemoveSubscriber(client.ID)
			}
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:    256,  // bytes
	WriteBufferSize:   1024, // bytes
	WriteBufferPool:   &sync.Pool{},
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func connectHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return
	}
	defer c.Close()

	client := NewClient(c)
	handleClientRequest(client)
}

type Server struct {
	ctx      context.Context
	msgChan  chan interface{}
	wsServer *http.Server
}

func New(ctx context.Context, msgChan chan interface{}) *Server {
	return &Server{
		ctx:     ctx,
		msgChan: msgChan,
	}
}

func (s *Server) startSocketServer() error {
	handler := http.HandlerFunc(connectHandler)
	srv := &http.Server{
		Addr:         conf.Conf.WebsocketHost,
		Handler:      handler,
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 3,
	}
	s.wsServer = srv
	return srv.ListenAndServe()
}

func (s *Server) startConsumer() {
	for {
		select {
		case <-s.ctx.Done():
			logger.Infof("Websocket Consumer Exit")
			return
		case msg, ok := <-s.msgChan:
			if !ok {
				return
			}
			wsMsg := msg.(message.WebSocketMessage)
			channel := findChannel(wsMsg.ChannelID)
			if channel == nil {
				var err error
				channel, err = createChannelByID(wsMsg.ChannelID)
				if err != nil {
					continue
				}
			}
			channel.AddMessage(&wsMsg)
		}
	}
}

func (s *Server) Start() error {
	logger.Infof("websocket start. host: %s", conf.Conf.WebsocketHost)

	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
			stack := make([]byte, 2048)
			length := runtime.Stack(stack, false)
			logger.Errorf("unhandled error: %v %s", err, stack[:length])
		}
	}()

	go func() {
		s.startConsumer()
	}()

	wsServerErrChan := make(chan error, 1)
	go func() {
		if err := s.startSocketServer(); err != nil {
			wsServerErrChan <- err
		}
	}()

	select {
	case <-s.ctx.Done():
		logger.Infof("websocket shutdown")

		// now close the server gracefully ("shutdown")
		graceTime := 10 * time.Second
		timeoutCtx, cancel := context.WithTimeout(context.Background(), graceTime)
		if err := s.wsServer.Shutdown(timeoutCtx); err != nil {
			logger.Errorf("shutdown server error:%s", err.Error())
			return err
		}
		cancel()
	case err := <-wsServerErrChan:
		return err
	}
	return nil
}
