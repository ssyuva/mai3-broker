package api

import (
	"github.com/mcarloai/mai-v3-broker/common/model"
)

type (
	// BaseReq represents a basic request data containing wallet addresses
	// from which the request is sent
	BaseReq struct {
		Address string `json:"address"`
	}

	// BaseResp represents a basic response data containing status and description to indicates
	// the executing result of an api request
	BaseResp struct {
		Status int    `json:"status"`
		Desc   string `json:"desc"`
	}

	PlaceOrderReq struct {
		BaseReq
		PerpetualAddress string `json:"perpetualAddress"  validate:"required"`
		Side             string `json:"side"      validate:"required,oneof=buy sell"`
		OrderType        string `json:"orderType" validate:"required,oneof=limit stop-limit"`
		Price            string `json:"price"     validate:"required"`
		StopPrice        string `json:"stopPrice"`
		Amount           string `json:"amount"    validate:"required"`
		Expires          int64  `json:"expires"`
		Salt             int    `json:"salt" validate:"required"`
		Signature        string `json:"Signature" validate:"required"`
	}

	PlaceOrderResp struct {
		ID string `json:"orderID"   validate:"required,len=66"`
	}

	CancelOrderReq struct {
		BaseReq
		ID string `json:"id" param:"orderID" validate:"required,len=66"`
	}

	CancelAllOrdersReq struct {
		BaseReq
		MarketID string `json:"marketID"`
	}

	QuerySingleOrderReq struct {
		BaseReq
		OrderID string `json:"orderID" param:"orderID" validate:"required"`
	}

	QuerySingleOrderResp struct {
		Order *model.Order `json:"order"`
	}

	QueryOrdersByIDsReq struct {
		BaseReq
		OrderIDs []string `json:"orderIDs" validate:"required"`
	}

	QueryOrderReq struct {
		BaseReq
		PerpetualAddress string `json:"perpetualAddress" query:"perpetualAddress"`
		Status           string `json:"status"   query:"status"`
		BeforeOrderHash  string `json:"beforeOrderHash"     query:"beforeOrderHash"`
		AfterOrderHash   string `json:"afterOrderHash"  query:"afterOrderHash"`
		Limit            int    `json:"limit" query:"limit"`
	}

	QueryOrdersResp struct {
		Orders []*model.Order `json:"orders"`
	}
)

func (b *BaseReq) GetAddress() string {
	return b.Address
}

func (b *BaseReq) SetAddress(address string) {
	b.Address = address
}

type Param interface {
	GetAddress() string
	SetAddress(address string)
}
