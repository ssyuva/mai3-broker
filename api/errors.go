package api

import (
	"bytes"
	"fmt"
	"time"

	"github.com/labstack/echo"
	"github.com/shopspring/decimal"
	"gopkg.in/go-playground/validator.v9"
)

// ApiError represents error occurs in api requests, containing extra error infomation.
type ApiError struct {
	Code          int
	Desc          string
	Params        []string
	internalError error
}

const (
	// MinOrderExpiration and MaxOrderExpiration represents the valid range of order expiration time;
	MinOrderExpiration = time.Hour * 24
	MaxOrderExpiration = time.Hour * 24 * 90
)

// Unwrap extract internal error in ApiError
func (e *ApiError) Unwrap() error {
	return e.internalError
}

func (e *ApiError) Error() string {
	if e.internalError != nil {
		return fmt.Sprintf("API Error[%d]:%s:%s", e.Code, e.Desc, e.internalError.Error())
	}
	return fmt.Sprintf("API Error[%d]:%s", e.Code, e.Desc)
}

// InternalError represents error caused by internal server error
func InternalError(err error) *ApiError {
	return &ApiError{Code: -1, Desc: "Internal Server Error", internalError: err}
}

// ValidationError represents error caused by validation failure
func ValidationError(message string) *ApiError {
	return &ApiError{Code: -2, Desc: message}
}

// BindError represents ...
func BindError(err error) *ApiError {
	return &ApiError{Code: -3, Desc: "bind error", internalError: err}
}

// PerpetualNotFoundError represents error caused by not existe market contract
func PerpetualNotFoundError(perpetual string, index int64) *ApiError {
	return &ApiError{Code: -4, Desc: fmt.Sprintf("not support perpetual address: %s index:%d", perpetual, index)}
}

func InvalidPriceAmountError(desc string) *ApiError {
	return &ApiError{Code: -5, Desc: desc}
}

func OrderIDNotExistError(id string) *ApiError {
	return &ApiError{Code: -6, Desc: fmt.Sprintf("order id[%s] do not exist", id),
		Params: []string{id}}
}

func BadSignatureError() *ApiError {
	return &ApiError{Code: -7, Desc: "bad signature"}
}

func InsufficientBalanceError() *ApiError {
	return &ApiError{Code: -8, Desc: "insufficient balance"}
}

func InsufficientAllowanceError(token string, need, has decimal.Decimal) *ApiError {
	return &ApiError{Code: -9, Desc: fmt.Sprintf("insufficient allowance of %s, need %s, has %s", token, need, has)}
}

func ValidatorError(errors validator.ValidationErrors) *ApiError {
	buff := bytes.Buffer{}

	for _, err := range errors {
		buff.WriteString(buildSingleError(err))
		buff.WriteString(";")
	}

	return &ApiError{Code: -10, Desc: fmt.Sprintf("validation fail: %s", buff.String()), internalError: errors}
}

func buildSingleError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	default:
		return fmt.Sprintf("Key: '%s' Error:Field validation for '%s' failed on the '%s' tag", err.Namespace(), err.Field(), err.Tag())
	}
}

func AuthError(desc string, err error) *ApiError {
	return &ApiError{
		Code:          -11,
		Desc:          fmt.Sprintf("Authentication failed. %s", desc),
		internalError: err,
	}
}

func EchoHTTPError(err *echo.HTTPError) *ApiError {
	return &ApiError{Code: -12, Desc: fmt.Sprintf("%v", err.Message), internalError: err.Internal}
}

func OrderAuthError(id string) *ApiError {
	return &ApiError{Code: -13, Desc: fmt.Sprintf("order hash[%s] do not exist or is not owned by current address", id),
		Params: []string{id}}
}

func RateLimitError(desc string, err error) *ApiError {
	return &ApiError{
		Code:          -14,
		Desc:          fmt.Sprintf("Api rate-limit error. %s", desc),
		internalError: err,
	}
}

func InvalidExpiresError() *ApiError {
	return &ApiError{
		Code:   -15,
		Desc:   fmt.Sprintf("expires must in range [%s, %s]", MinOrderExpiration, MaxOrderExpiration),
		Params: []string{MinOrderExpiration.String(), MaxOrderExpiration.String()}}
}

func ContractSettledError() *ApiError {
	return &ApiError{
		Code: -16,
		Desc: fmt.Sprint("the contract is settled, forbidden to place orders"),
	}
}

func MaxOrderNumReachError() *ApiError {
	return &ApiError{Code: -17, Desc: "cannot create more order"}
}

func OrderHashExistError(orderHash string) *ApiError {
	return &ApiError{Code: -18, Desc: fmt.Sprintf("order hash[%s] exist", orderHash),
		Params: []string{orderHash}}
}

func BrokerAddressError(address string) *ApiError {
	return &ApiError{Code: -19, Desc: fmt.Sprintf("Invalid broker address [%s]", address)}
}

func ChainIDError(chainID int64) *ApiError {
	return &ApiError{Code: -20, Desc: fmt.Sprintf("not in same chain. chainID [%d]", chainID)}
}

func OrderExpired() *ApiError {
	return &ApiError{Code: -21, Desc: "order expired"}
}

func GasBalanceError() *ApiError {
	return &ApiError{Code: -22, Desc: "gas balance not enough for trade"}
}

func CloseOnlyError() *ApiError {
	return &ApiError{Code: -23, Desc: "amount is not right for close only order"}
}

func AddressNotInWhiteListError() *ApiError {
	return &ApiError{Code: -24, Desc: "address is not in whitelist"}
}

func StopOrderError() *ApiError {
	return &ApiError{Code: -25, Desc: "stop order amount not available"}
}
