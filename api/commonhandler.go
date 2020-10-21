package api

import (
	"errors"
	"github.com/labstack/echo"
	v "gopkg.in/go-playground/validator.v9"
	"net/http"
	"reflect"
)

var validate = v.New()

type Response struct {
	Status int         `json:"status"`
	Desc   string      `json:"desc"`
	Data   interface{} `json:"data,omitempty"`
}

func bindAndValidParams(c echo.Context, params Param) (err error) {
	if err := c.Bind(params); err != nil {
		return BindError(err)
	}

	bindUrlParam(c, params)
	return validate.Struct(params)
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

func CommonHandler(params Param, fn func(Param) (interface{}, error)) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		// Params is shared among all request
		// We create a new one for each request
		var reqParam Param
		if params != nil {
			reqParam = reflect.New(reflect.TypeOf(params).Elem()).Interface().(Param)
			err = bindAndValidParams(c, reqParam)
		}

		if err != nil {
			return
		}

		if reqParam != nil {
			cc := c.(*MaiApiContext)
			reqParam.SetAddress(cc.Address)
		}
		resp, err := fn(reqParam)

		if err != nil {
			return
		}

		return commonResponse(c, resp)
	}
}

func commonResponse(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, Response{
		Status: 0,
		Desc:   "success",
		Data:   data,
	})
}

func ErrorHandler(err error, c echo.Context) {
	e := c.Echo()

	var apiError *ApiError
	var validationErrors v.ValidationErrors
	var echoHTTPError *echo.HTTPError

	if ok := errors.As(err, &validationErrors); ok {
		apiError = ValidatorError(validationErrors)
	} else if ok := errors.As(err, &echoHTTPError); ok {
		apiError = EchoHTTPError(echoHTTPError)
	} else {
		if ok := errors.As(err, &apiError); !ok {
			apiError = InternalError(err)
		}
	}

	// Send response
	if !c.Response().Committed {
		err = c.JSON(http.StatusOK, Response{
			Status: apiError.Code,
			Desc:   apiError.Desc,
			Data:   apiError.Params,
		})

		if err != nil {
			e.Logger.Error(err)
		}
	}
}
