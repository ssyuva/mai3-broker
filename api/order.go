package api

import (
	"errors"
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

const TIMESTAMP_RANGE = 5 * 60

func (s *Server) GetOrders(p Param) (interface{}, error) {
	params := p.(*QueryOrderReq)
	var beforeOrderID, afterOrderID int64
	if params.BeforeOrderHash != "" {
		beforeOrder, err := s.dao.GetOrder(params.BeforeOrderHash, false)
		if err != nil {
			if dao.IsRecordNotFound(err) {
				return nil, OrderIDNotExistError(params.BeforeOrderHash)
			}
			return nil, InternalError(err)
		}

		beforeOrderID = beforeOrder.ID
	}

	if params.AfterOrderHash != "" {
		afterOrder, err := s.dao.GetOrder(params.AfterOrderHash, false)
		if err != nil {
			if dao.IsRecordNotFound(err) {
				return nil, OrderIDNotExistError(params.AfterOrderHash)
			}
			return nil, InternalError(err)
		}
		afterOrderID = afterOrder.ID
	}

	if params.PerpetualAddress != "" {
		_, err := s.dao.GetPerpetualByAddress(params.PerpetualAddress)
		if err != nil {
			if dao.IsRecordNotFound(err) {
				return nil, PerpetualNotFoundError(params.PerpetualAddress)
			}
			return nil, InternalError(err)
		}
	}

	queryStatus := make([]model.OrderStatus, 0)
	if params.Status != "all" {
		filterStatus := strings.Split(params.Status, model.QueryParamSeperator)
		for _, status := range filterStatus {
			queryStatus = append(queryStatus, model.OrderStatus(status))
		}
	}

	limit := 20
	if params.Limit > 0 {
		limit = params.Limit
	}
	orders, err := s.dao.QueryOrder(params.Address, params.PerpetualAddress, queryStatus, beforeOrderID, afterOrderID, limit)
	if err != nil {
		return nil, InternalError(err)
	}

	res := &QueryOrdersResp{
		Orders: orders,
	}
	return res, nil
}

func (s *Server) GetOrderByID(p Param) (interface{}, error) {
	params := p.(*QuerySingleOrderReq)
	order, err := s.dao.GetOrder(params.OrderID, false)
	if err != nil {
		if dao.IsRecordNotFound(err) {
			return nil, OrderIDNotExistError(params.OrderID)
		}
		return nil, InternalError(err)
	}
	res := &QuerySingleOrderResp{
		Order: order,
	}
	return res, nil
}

func (s *Server) GetOrdersByIDs(p Param) (interface{}, error) {
	params := p.(*QueryOrdersByIDsReq)
	orders, err := s.dao.GetOrderByIDs(params.OrderIDs)

	if err != nil {
		return nil, InternalError(err)
	}

	res := &QueryOrdersResp{
		Orders: orders,
	}
	return res, nil
}

func (s *Server) PlaceOrder(p Param) (interface{}, error) {
	params := p.(*PlaceOrderReq)
	err := validatePlaceOrder(params)
	if err != nil {
		return nil, err
	}

	if params.PerpetualAddress != "" {
		_, err := s.dao.GetPerpetualByAddress(params.PerpetualAddress)
		if err != nil {
			if dao.IsRecordNotFound(err) {
				return nil, PerpetualNotFoundError(params.PerpetualAddress)
			}
			return nil, InternalError(err)
		}
	}

	_, err = s.dao.GetOrderBySignature(params.Signature)
	if err == nil {
		return nil, BadSignatureError()
	} else if !dao.IsRecordNotFound(err) {
		return nil, InternalError(err)
	}

	//TODO Place Order
	//1. check signature
	//2. check margin balance
	//3. insert into order tables
	//4. insert into orderbook
	return nil, nil
}

func validatePlaceOrder(req *PlaceOrderReq) error {
	expiration := time.Second * time.Duration(req.Expires)
	if expiration < MinOrderExpiration || expiration > MaxOrderExpiration {
		return InvalidExpiresError()
	}

	// Amount
	amount, err := decimal.NewFromString(req.Amount)
	if err != nil {
		return InvalidPriceAmountError(fmt.Sprintf("parse amount[%s] error", req.Amount))
	}
	if amount.LessThanOrEqual(decimal.Zero) {
		return InvalidPriceAmountError("amount <= 0")
	}

	price, err := decimal.NewFromString(req.Price)
	if err != nil {
		return InvalidPriceAmountError(fmt.Sprintf("parse price[%s] error", req.Price))
	}

	if price.LessThanOrEqual(decimal.Zero) {
		return InvalidPriceAmountError("price <= 0")
	}

	// order sign timestamp
	now := time.Now().UTC().Unix()
	if (now-TIMESTAMP_RANGE) > req.Timestamp || (now+TIMESTAMP_RANGE) < req.Timestamp {
		return InternalError(errors.New("timestamp over time"))
	}

	// Price OrderType
	if req.OrderType == string(model.StopLimitOrder) {
		stopPrice, err := decimal.NewFromString(req.StopPrice)
		if err != nil {
			return InvalidPriceAmountError(fmt.Sprintf("parse price[%s] error", req.StopPrice))
		}

		if stopPrice.LessThanOrEqual(decimal.Zero) {
			return InvalidPriceAmountError("stop price <= 0")
		}
	} else if req.OrderType != string(model.LimitOrder) {
		return InternalError(errors.New("order type must be limit/stop-limit"))
	}

	return nil
}

func (s *Server) CancelOrder(p Param) (interface{}, error) {
	// params := p.(*CancelOrderReq)
	return nil, nil
}

func (s *Server) CancelAllOrders(p Param) (interface{}, error) {
	// params := p.(*CancelAllOrdersReq)
	return nil, nil
}
