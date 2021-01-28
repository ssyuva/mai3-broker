package rpc

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mcarloai/mai-v3-broker/api"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/l2/relayer"

	logger "github.com/sirupsen/logrus"
)

type L2RelayerServer struct {
	ctx context.Context
	e   *echo.Echo
	r   *relayer.Relayer
}

func NewL2RelayerServer(ctx context.Context, r *relayer.Relayer) (*L2RelayerServer, error) {
	e := echo.New()
	e.HideBanner = true

	e.Use(api.RecoverHandler)
	e.HTTPErrorHandler = errorHandler

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
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
		Addr:         conf.L2RelayerConf.L2RelayerHost,
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

	addGroupRoute(eg, "POST", "/call", &CallL2FunctionReq{}, s.CallL2Function)
}

func (s *L2RelayerServer) CallL2Function(p Param) (interface{}, error) {
	params, ok := p.(*CallL2FunctionReq)
	if !ok {
		return nil, errors.New("Unknown error, bad params types")
	}

	gasLimit, err := strconv.ParseUint(params.GasLimit, 10, 64)
	if err != nil {
		return nil, relayer.NewInvalidRequestError("invalid gas limit")
	}

	ctx, cancel := context.WithTimeout(s.ctx, conf.L2RelayerConf.L2Timeout)
	defer cancel()
<<<<<<< HEAD
	tx, err := s.r.CallFunction(ctx, params.From, params.To, params.FunctionSignature, params.CallData, params.Nonce, params.Expiration, params.GasLimit, params.Signature)
=======

	tx, err := s.r.CallFunction(ctx, params.From, params.To, params.Method, params.CallData, params.Nonce, params.Expiration, gasLimit, params.Signature)
>>>>>>> update relayer
	if err != nil {
		return nil, err
	}

	res := &CallL2FunctionResp{
		TransactionHash: tx,
	}

	return res, nil
}
