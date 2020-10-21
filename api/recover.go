package api

import (
	"fmt"
	"runtime"

	"github.com/labstack/echo"
	"github.com/micro/go-micro/v2/logger"
)

func RecoverHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				stack := make([]byte, 2048)
				length := runtime.Stack(stack, false)
				logger.Errorf("unhandled error: %v %s", err, stack[:length])
				c.Error(err)
			}
		}()
		return next(c)
	}
}
