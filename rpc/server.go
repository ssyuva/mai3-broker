package rpc

import (
	"context"
	"github.com/mcarloai/mai-v3-broker/conf"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	logger "github.com/sirupsen/logrus"
)

type Handler interface {
	Bind(e *echo.Echo)
}

func StartServer(ctx context.Context, handlers ...Handler) error {
	e := NewCmdServer(false)

	for _, h := range handlers {
		h.Bind(e)
	}

	s := &http.Server{
		Addr:         conf.Conf.RPCHost,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	stop := make(chan error, 1)

	go func() {
		err := e.StartServer(s)
		stop <- err
	}()

	select {
	case err := <-stop:
		logger.Warnf("RPC Server exit:%s", err.Error())
		return err
	case <-ctx.Done():
		logger.Warnf("RPC Server receive context done")
		grace, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		e.Listener.Close()
		if err := e.Shutdown(grace); err != nil {
			e.Logger.Fatal(err)
		}
	}
	return nil
}

type Response struct {
	Status int         `json:"status"`
	Desc   string      `json:"desc"`
	Data   interface{} `json:"data,omitempty"`
}

func errorHandler(err error, c echo.Context) {
	e := c.Echo()

	// Send response
	if !c.Response().Committed {
		err = c.JSON(http.StatusOK, Response{
			Status: -1,
			Desc:   err.Error(),
		})

		if err != nil {
			e.Logger.Error(err)
		}
	}
}

func NewCmdServer(isLogHTTP bool) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HTTPErrorHandler = errorHandler

	if isLogHTTP {
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: `{"time":"${time_rfc3339}","remote_ip":"${remote_ip}","method":"${method}","uri":"${uri}","status":${status},"error":"${error}"}` + "\n",
		}))
	}

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	return e
}

func CmdResponse(e echo.Context, data interface{}) error {
	var response Response
	response.Status = 0
	response.Desc = "success"
	response.Data = data

	return e.JSONPretty(http.StatusOK, response, "  ")
}
