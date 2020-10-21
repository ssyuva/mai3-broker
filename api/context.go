package api

import (
	"github.com/labstack/echo"
)

// MaiApiContext holds user information for API requests
type MaiApiContext struct {
	echo.Context
	// If address is not empty means this user is authenticated.
	Address string
}

func InitMaiApiContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &MaiApiContext{Context: c}
		return next(cc)
	}
}
