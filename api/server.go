package api

import (
	"context"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mcarloai/mai-v3-broker/common/chain"
	redis "github.com/mcarloai/mai-v3-broker/common/redis"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/dao"
	logger "github.com/sirupsen/logrus"
	limiter_redis "github.com/ulule/limiter/v3/drivers/store/redis"
	"net/http"
	"time"
)

type Server struct {
	ctx       context.Context
	e         *echo.Echo
	wsChan    chan interface{}
	matchChan chan interface{}
	chainCli  chain.ChainClient
	dao       dao.DAO
}

func New(ctx context.Context, cli chain.ChainClient, dao dao.DAO, wsChan, matchChan chan interface{}) (*Server, error) {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339}","remote_ip":"${remote_ip}","method":"${method}","uri":"${uri}","status":${status},"error":"${error}"}` + "\n",
	}))

	var err error
	RateLimitCacheService, err = limiter_redis.NewStore(redis.RedisClient)
	if err != nil {
		return nil, err
	}
	e.Use(IPRatelimiter())
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

	s := &Server{
		ctx:       ctx,
		e:         e,
		wsChan:    wsChan,
		matchChan: matchChan,
		chainCli:  cli,
		dao:       dao,
	}
	s.initRouter()
	return s, nil
}

func (s *Server) Start() error {
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

func (s *Server) initRouter() {
	eg := s.e.Group("/orders")
	addGroupRoute(eg, "GET", "", &QueryOrderReq{}, s.GetOrders)
	addGroupRoute(eg, "GET", "/:orderHash", &QuerySingleOrderReq{}, s.GetOrderByOrderHash)
	addGroupRoute(eg, "POST", "/byhashs", &QueryOrdersByOrderHashsReq{}, s.GetOrdersByOrderHashs)
	addGroupRoute(eg, "POST", "", &PlaceOrderReq{}, s.PlaceOrder)
	addGroupRoute(eg, "DELETE", "/:orderHash", &CancelOrderReq{}, s.CancelOrder, MaiAuthMiddleware, JwtAuthMiddleware, CheckAuthMiddleware)
	addGroupRoute(eg, "DELETE", "", &CancelAllOrdersReq{}, s.CancelAllOrders, MaiAuthMiddleware, JwtAuthMiddleware, CheckAuthMiddleware)

	addRoute(s.e, "GET", "/jwt", &BaseReq{}, GetJwtAuth, MaiAuthMiddleware, CheckAuthMiddleware)
	addRoute(s.e, "GET", "/perpetuals/:perpetual", &GetPerpetualReq{}, s.GetPerpetual)
	addRoute(s.e, "GET", "/brokerRelay", &GetBrokerRelayReq{}, s.GetBrokerRelay)
}

func addGroupRoute(eg *echo.Group, method, url string, param Param, handler func(p Param) (interface{}, error), middlewares ...echo.MiddlewareFunc) {
	eg.Add(method, url, CommonHandler(param, handler), middlewares...)
}

func addRoute(e *echo.Echo, method, url string, param Param, handler func(p Param) (interface{}, error), middlewares ...echo.MiddlewareFunc) {
	e.Add(method, url, CommonHandler(param, handler), middlewares...)
}