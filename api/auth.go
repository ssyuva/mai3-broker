package api

import (
	"net/url"
	"strings"

	"github.com/labstack/echo"
	"github.com/mcarloai/mai-v3-broker/common/auth"
)

type authenticator interface {
	// IsAvail test if request contains specified header field
	IsAvail(c *MaiApiContext) bool
	// Valid test if the auth field is valid
	Valid(c *MaiApiContext) (string, error)
}

type maiAuthenticator struct {
	AuthHeader string
}

// IsAvail test if a request contains mai auth header
func (ma *maiAuthenticator) IsAvail(c *MaiApiContext) bool {
	return c.Request().Header.Get(ma.AuthHeader) != ""
}

// Valid validate the token
func (ma *maiAuthenticator) Valid(c *MaiApiContext) (string, error) {
	token := c.Request().Header.Get("Mai-Authentication")
	token, err := url.QueryUnescape(token)
	if err != nil {
		return "", AuthError("Mai-Authentication requires uri-encoded", err)
	}
	address, err := auth.ValidateMaiAuth(token)
	if err != nil {
		return "", AuthError("Mai-Authentication fail:", err)
	}
	return address, nil
}

type jwtAuthenticator struct {
	AuthHeader string
	AuthScheme string
}

// jwtFromHeader returns a `jwtExtractor` that extracts token from the request header.
func (ja *jwtAuthenticator) jwtFromHeader(c echo.Context) string {
	auth := c.Request().Header.Get(ja.AuthHeader)
	l := len(ja.AuthScheme)
	if len(auth) > l+1 && auth[:l] == ja.AuthScheme {
		return auth[l+1:]
	}
	return ""
}

func (ja *jwtAuthenticator) IsAvail(c *MaiApiContext) bool {
	return c.Request().Header.Get(ja.AuthHeader) != ""
}

func (ja *jwtAuthenticator) Valid(c *MaiApiContext) (string, error) {
	token := ja.jwtFromHeader(c)
	address, err := auth.ValidateJwt(token)
	if err != nil {
		return "", AuthError("JWT-Authentication fail:", err)
	}
	return address, nil
}

var (
	maiAuth = &maiAuthenticator{
		AuthHeader: "Mai-Authentication",
	}
	jwtAuth = &jwtAuthenticator{
		AuthHeader: "Authentication",
		AuthScheme: "Bearer",
	}
)

func JwtAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*MaiApiContext)
		cc.Response().Header().Set(echo.HeaderServer, "Echo/3.0")

		if cc.Address != "" {
			return next(cc)
		}
		if !jwtAuth.IsAvail(cc) {
			return next(cc)
		}
		address, err := jwtAuth.Valid(cc)
		if err != nil {
			return err
		}
		cc.Address = strings.ToLower(address)
		return next(cc)
	}
}

func MaiAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*MaiApiContext)
		cc.Response().Header().Set(echo.HeaderServer, "Echo/3.0")

		if cc.Address != "" {
			return next(cc)
		}
		if !maiAuth.IsAvail(cc) {
			return next(cc)
		}
		address, err := maiAuth.Valid(cc)
		if err != nil {
			return err
		}

		cc.Address = strings.ToLower(address)
		return next(cc)
	}
}

func CheckAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*MaiApiContext)
		if cc.Address == "" {
			return AuthError("Authentication required.", nil)
		}

		return next(cc)
	}
}
