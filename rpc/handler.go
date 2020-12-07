package rpc

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/launcher"
	"github.com/mcarloai/mai-v3-broker/match"
)

type RPCHandler struct {
	match    *match.Server
	launcher *launcher.Launcher
}

func NewRPCHandler(match *match.Server, l *launcher.Launcher) *RPCHandler {
	return &RPCHandler{
		match:    match,
		launcher: l,
	}
}

func (r *RPCHandler) Bind(e *echo.Echo) {
	e.Add("DELETE", "/perpetuals/:perpetual_address", r.ClosePerpetual)
	e.Add("DELETE", "/orders/:order_hash", r.CancelOrder)
	e.Add("POST", "/orders", r.NewOrder)
	e.Add("DELETE", "/orders", r.CancelOrders)
	e.Add("POST", "/batch_trade", r.BatchTradeOrders)
}

func (r *RPCHandler) NewOrder(e echo.Context) (err error) {
	var order model.Order

	err = e.Bind(&order)
	if err != nil {
		return fmt.Errorf("NewOrder:%w", err)
	}

	err = r.match.NewOrder(&order)
	if err != nil {
		return fmt.Errorf("NewOrder:%w", err)
	}

	return CmdResponse(e, nil)
}

func (r *RPCHandler) CancelOrder(e echo.Context) error {
	orderHash := e.Param("order_hash")
	if orderHash == "" {
		return fmt.Errorf("CancelOrder: orderHash is blank, check param")
	}

	var req struct {
		PerpetualAddress string `json:"perpetual_address" query:"perpetual_address"`
	}
	err := e.Bind(&req)
	if err != nil {
		return fmt.Errorf("CancelOrder: %w", err)
	}

	err = r.match.CancelOrder(req.PerpetualAddress, orderHash)
	if err != nil {
		return fmt.Errorf("CancelOrder: %w", err)
	}

	return CmdResponse(e, nil)
}

func (r *RPCHandler) CancelOrders(e echo.Context) error {
	var req struct {
		PerpetualAddress string `json:"perpetual_address" query:"perpetual_address" validate:"required"`
		Trader           string `json:"trader" query:"trader" validate:"required"`
	}
	err := e.Bind(&req)
	if err != nil {
		return fmt.Errorf("CancelOrders: %w", err)
	}

	err = r.match.CancelAllOrders(req.PerpetualAddress, req.Trader)
	if err != nil {
		return fmt.Errorf("CancelOrders: %w", err)
	}

	return CmdResponse(e, nil)
}

func (r *RPCHandler) ClosePerpetual(e echo.Context) error {
	perpetualAddress := e.Param("perpetual_address")
	if perpetualAddress == "" {
		return fmt.Errorf("ClosePerpetual: PerpetualAddress is blank, check param")
	}

	err := r.match.ClosePerpetual(perpetualAddress)
	if err != nil {
		return fmt.Errorf("ClosePerpetual: %w", err)
	}

	return CmdResponse(e, nil)
}

func (r *RPCHandler) BatchTradeOrders(e echo.Context) error {
	var req struct {
		TxID            string `json:"tx_id" query:"tx_id" validate:"required"`
		TransactionHash string `json:"transactionHash" query:"transactionHash" validate:"required"`
		BlockNumber     uint64 `json:"blockNumber" query:"blockNumber" validate:"required"`
		BlockHash       string `json:"blockHash" query:"blockHash" validate:"required"`
		BlockTime       uint64 `json:"blockTime" query:"blockTime" validate:"required"`
		Status          string `json:"status" query:"status" validate:"required"`
	}
	err := e.Bind(&req)
	if err != nil {
		return fmt.Errorf("CancelOrders: %w", err)
	}
	err = r.match.BatchTradeOrders(req.TxID, req.Status, req.TransactionHash, req.BlockHash, req.BlockNumber, req.BlockTime)
	if err != nil {
		return fmt.Errorf("BatchTradeOrder:%w", err)
	}
	return CmdResponse(e, nil)
}
