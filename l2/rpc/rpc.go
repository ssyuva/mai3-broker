package rpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/labstack/echo"
	"github.com/mcarloai/mai-v3-broker/api"
	"github.com/mcarloai/mai-v3-broker/l2/relayer"
	v "gopkg.in/go-playground/validator.v9"

	logger "github.com/sirupsen/logrus"
)

type (
	Param interface {
	}

	CallL2FunctionReq struct {
		From       string `json:"from"`
		To         string `json:"to"`
		Method     string `json:"method"`
		CallData   string `json:"callData"`
		Nonce      uint32 `json:"nonce"`
		Expiration uint32 `json:"expiration"`
		GasLimit   string `json:"gasLimit"`
		Signature  string `json:"signature"`
	}

	CallL2FunctionResp struct {
		TransactionHash string `json:"transactionHash"`
	}
)

var validate = v.New()

func bindAndValidParams(c echo.Context, params Param) (err error) {
	if err := c.Bind(params); err != nil {
		return relayer.NewInvalidRequestError(fmt.Sprintf("bind params error:%s", err.Error()))
	}

	bindUrlParam(c, params)
	if err := validate.Struct(params); err != nil {
		return relayer.NewInvalidRequestError(fmt.Sprintf("validate params error:%s", err.Error()))
	}
	return nil
}

func bindUrlParam(c echo.Context, ptr interface{}) {
	typ := reflect.TypeOf(ptr).Elem()
	val := reflect.ValueOf(ptr).Elem()

	for i := 0; i < typ.NumField(); i++ {
		typeField := typ.Field(i)
		structField := val.Field(i)
		if !structField.CanSet() {
			continue
		}
		inputFieldName := typeField.Tag.Get("param")
		if inputFieldName == "" {
			continue
		}

		structField.SetString(c.Param(inputFieldName))
		continue
	}
}

func logRequest(level logger.Level, req, resp interface{}) {
	var reqStr, respStr string
	l := logger.StandardLogger()
	if reqJson, err := json.Marshal(req); err == nil {
		reqStr = string(reqJson)
	} else {
		reqStr = fmt.Sprintf("%v", req)
	}

	if respJson, err := json.Marshal(resp); err == nil {
		respStr = string(respJson)
	} else {
		respStr = fmt.Sprintf("%v", resp)
	}

	l.Logf(level, "req:%s resp:%s", reqStr, respStr)
}

func newHandler(params Param, fn func(Param) (interface{}, error)) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		// Params is shared among all request
		// We create a new one for each request
		var reqParam Param

		if params != nil {
			reqParam = reflect.New(reflect.TypeOf(params).Elem()).Interface().(Param)
			err = bindAndValidParams(c, reqParam)
		}

		url := c.Request().URL

		if err != nil {
			logger.Warnf("%s %s", url, err.Error())
			return
		}

		resp, err := fn(reqParam)

		if err != nil {
			logRequest(logger.WarnLevel, reqParam, err)
			return
		}

		response := &api.Response{
			Status: 0,
			Desc:   "success",
			Data:   resp,
		}

		logRequest(logger.InfoLevel, reqParam, resp)

		return c.JSON(http.StatusOK, response)
	}
}

func addGroupRoute(eg *echo.Group, method, url string, param Param, handler func(p Param) (interface{}, error), middlewares ...echo.MiddlewareFunc) {
	eg.Add(method, url, newHandler(param, handler), middlewares...)
}

func errorHandler(err error, c echo.Context) {
	var relayerError *relayer.Error

	if ok := errors.As(err, &relayerError); !ok {
		relayerError = relayer.NewInternalServerError()
	}

	// Send response
	if !c.Response().Committed {

		response := &api.Response{
			Status: relayerError.Code,
			Desc:   relayerError.Message,
			Data:   relayerError.ChainError,
		}

		err = c.JSON(http.StatusOK, response)
		if err != nil {
			logger.Warnf("send error response fail:%s", err.Error())
		}
	}
}
