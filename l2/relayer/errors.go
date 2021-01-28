package relayer

import (
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
)

const (
	RelayerOK            = 0
	InternalServerError  = 1
	InvalidRequestError  = 2
	InsufficentGasError  = 3
	EstimateGasError     = 4
	SendTransactionError = 5
)

type ChainError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Error struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	ChainError *ChainError `json:"chainError,omitempty"`
}

func (r *Error) Error() string {
	return fmt.Sprintf("%d:%s", r.Code, r.Message)
}

func getChainError(err error) *ChainError {
	if r, ok := err.(rpc.Error); ok {
		return &ChainError{
			Code:    r.ErrorCode(),
			Message: r.Error(),
		}
	}
	return nil
}

func NewEstimateGasError(rpcError error) *Error {
	result := &Error{
		Code:       EstimateGasError,
		Message:    "always failing transaction or unpredicatable gas",
		ChainError: getChainError(rpcError),
	}
	return result
}

func NewSendTransactionError(rpcError error) *Error {
	result := &Error{
		Code:       SendTransactionError,
		Message:    "send transaction to block chain error",
		ChainError: getChainError(rpcError),
	}

	return result
}

func NewInsufficentGasError(need uint64) *Error {
	return &Error{
		Code:    InsufficentGasError,
		Message: fmt.Sprintf("need at least %d broker gas", need),
	}
}

func NewInvalidRequestError(msg string) *Error {
	return &Error{
		Code:    InvalidRequestError,
		Message: fmt.Sprintf("bad paramter:%s", msg),
	}
}

func NewInternalServerError() *Error {
	return &Error{
		Code:    InternalServerError,
		Message: "internal server error",
	}
}
