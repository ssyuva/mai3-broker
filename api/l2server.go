package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/l2relayer"
	logger "github.com/sirupsen/logrus"
)

type L2RelayerServer struct {
	ctx context.Context
	e   *echo.Echo
	r   *l2relayer.L2Relayer
}

func NewL2RelayerServer(ctx context.Context, r *l2relayer.L2Relayer) (*L2RelayerServer, error) {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339}","remote_ip":"${remote_ip}","method":"${method}","uri":"${uri}","status":${status},"error":"${error}"}` + "\n",
	}))

	e.Use(RecoverHandler)
	e.HTTPErrorHandler = ErrorHandler
	e.Use(InitMaiApiContext)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			"Authentication",
			"Mai-Authentication",
		},
	}))

	s := &L2RelayerServer{
		ctx: ctx,
		e:   e,
		r:   r,
	}
	s.initRouter()
	return s, nil
}

func (s *L2RelayerServer) Start() error {
	srv := &http.Server{
		Addr:         conf.Conf.APIHost,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	srvFail := make(chan error, 1)
	go func() {
		if err := s.e.StartServer(srv); err != nil {
			srvFail <- err
		}
	}()

	select {
	case <-s.ctx.Done():
		s.e.Listener.Close()
		// now close the server gracefully ("shutdown")
		graceTime := 10 * time.Second
		timeoutCtx, cancel := context.WithTimeout(context.Background(), graceTime)
		if err := s.e.Shutdown(timeoutCtx); err != nil {
			logger.Errorf("shutdown server error:%s", err.Error())
		}
		cancel()
	case err := <-srvFail:
		return err
	}
	return nil
}

func (s *L2RelayerServer) initRouter() {
	eg := s.e.Group("/l2relayer")
	addGroupRoute(eg, "GET", "/address", &GetL2RelayerAddressReq{}, s.GetL2RelayerAddress)
	addGroupRoute(eg, "POST", "/call", &CallL2FunctionReq{}, s.CallL2Function)
	addGroupRoute(eg, "POST", "/trade", &PlaceOrderReq{}, s.Trade)
}

func (s *L2RelayerServer) GetL2RelayerAddress(p Param) (interface{}, error) {
	res := &GetL2RelayerAddressResp{
		L2RelayerAddress: s.r.Address(),
	}

	return res, nil
}

func (s *L2RelayerServer) CallL2Function(p Param) (interface{}, error) {
	params := p.(*CallL2FunctionReq)

	ctx, cancel := context.WithTimeout(s.ctx, conf.Conf.L2Relayer.L2Timeout.Duration)
	defer cancel()
	tx, err := s.r.CallFunction(ctx, params.FunctionSignature, params.CallData, params.Address, params.Nonce, params.Expiration, params.GasFeeLimit, params.Signature)
	if err != nil {
		// TODO
		return nil, InternalError(err)
	}

	res := &CallL2FunctionResp{
		TransactionHash: tx,
	}

	return res, nil
}

func (s *L2RelayerServer) Trade(p Param) (interface{}, error) {
	params := p.(*PlaceOrderReq)

	order, err := GetOrderFromPalceOrderReq(params)
	if err != nil {
		return nil, err
	}

	// Limit Order Only
	if order.Type != model.LimitOrder {
		return nil, InternalError(errors.New("order type must be limit"))
	}

	// OnlyAllow 60s ExpiresAt
	if order.CreatedAt.Add(conf.Conf.L2Relayer.MaxTradeExpiration.Duration).After(order.ExpiresAt) {
		return nil, InternalError(fmt.Errorf("too large expiration, max=%s", conf.Conf.L2Relayer.MaxTradeExpiration.String()))
	}

	ctx, cancel := context.WithTimeout(s.ctx, conf.Conf.L2Relayer.L2Timeout.Duration)
	defer cancel()
	tx, err := s.r.Trade(ctx, order)
	if err != nil {
		//TODO
		return nil, InternalError(err)
	}
	res := &L2TradeResp{
		TransactionHash: tx,
	}
	return res, nil
}
