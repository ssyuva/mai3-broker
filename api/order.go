package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/mai3"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

const MAX_ORDER_NUM = 20
const ADDRESS_ZERO = "0x0000000000000000000000000000000000000000"

func (s *Server) GetOrders(p Param) (interface{}, error) {
	params := p.(*QueryOrderReq)
	var beforeOrderID, afterOrderID int64
	if params.BeforeOrderHash != "" {
		beforeOrder, err := s.dao.GetOrder(params.BeforeOrderHash)
		if err != nil {
			if dao.IsRecordNotFound(err) {
				return nil, OrderIDNotExistError(params.BeforeOrderHash)
			}
			return nil, InternalError(err)
		}

		beforeOrderID = beforeOrder.ID
	}

	if params.AfterOrderHash != "" {
		afterOrder, err := s.dao.GetOrder(params.AfterOrderHash)
		if err != nil {
			if dao.IsRecordNotFound(err) {
				return nil, OrderIDNotExistError(params.AfterOrderHash)
			}
			return nil, InternalError(err)
		}
		afterOrderID = afterOrder.ID
	}

	if params.PerpetualAddress != "" {
		_, err := s.dao.GetPerpetualByAddress(params.PerpetualAddress, true)
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

func (s *Server) GetOrderByOrderHash(p Param) (interface{}, error) {
	params := p.(*QuerySingleOrderReq)
	order, err := s.dao.GetOrder(params.OrderHash)
	if err != nil {
		if dao.IsRecordNotFound(err) {
			return nil, OrderIDNotExistError(params.OrderHash)
		}
		return nil, InternalError(err)
	}
	res := &QuerySingleOrderResp{
		Order: order,
	}
	return res, nil
}

func (s *Server) GetOrdersByOrderHashs(p Param) (interface{}, error) {
	params := p.(*QueryOrdersByOrderHashsReq)
	orders, err := s.dao.GetOrderByHashs(params.OrderHashs)

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

	order := model.Order{}
	order.OrderHash = params.OrderHash
	order.OrderParam.TraderAddress = strings.ToLower(params.Address)
	if params.OrderType == int(model.LimitOrder) {
		order.OrderParam.Type = model.LimitOrder
		order.Status = model.OrderPending
	} else {
		order.OrderParam.Type = model.StopLimitOrder
		order.Status = model.OrderStop
	}

	order.OrderParam.Amount, _ = decimal.NewFromString(params.Amount)
	order.OrderParam.Version = int32(mai3.ProtocolV3)
	order.OrderParam.ChainID = params.ChainID
	order.OrderParam.Price, _ = decimal.NewFromString(params.Price)
	order.OrderParam.StopPrice, _ = decimal.NewFromString(params.StopPrice)
	order.OrderParam.IsCloseOnly = params.IsCloseOnly
	order.OrderParam.ExpiresAt = time.Unix(params.ExpiresAt, 0).UTC()
	order.OrderParam.Salt = params.Salt
	order.OrderParam.Signature = params.Signature
	order.PerpetualAddress = strings.ToLower(params.PerpetualAddress)
	order.BrokerAddress = strings.ToLower(params.BrokerAddress)
	order.RelayerAddress = strings.ToLower(params.RelayerAddress)
	if params.ReferrerAddress == "" {
		order.OrderParam.ReferrerAddress = ADDRESS_ZERO
	}
	order.AvailableAmount = order.OrderParam.Amount
	order.ConfirmedAmount = decimal.Zero
	order.CanceledAmount = decimal.Zero
	order.PendingAmount = decimal.Zero
	now := time.Now().UTC()
	order.CreatedAt = now
	order.UpdatedAt = now

	// check orderhash
	orderData := mai3.GenerateOrderData(params.ExpiresAt, order.Version, int8(order.Type), order.IsCloseOnly, order.OrderParam.Salt)
	orderHash, err := mai3.GetOrderHash(order.TraderAddress, order.BrokerAddress, order.RelayerAddress, order.PerpetualAddress, order.ReferrerAddress,
		orderData, order.Amount, order.Price, order.ChainID)
	if err != nil {
		return nil, InternalError(fmt.Errorf("get order hash fail err:%s", err))
	}

	if utils.Bytes2HexP(orderHash) != params.OrderHash {
		return nil, InternalError(fmt.Errorf("order hash not match"))
	}

	// check signature
	valid, err := mai3.IsValidOrderSignature(params.Address, params.OrderHash, params.Signature)
	if err != nil {
		return nil, InternalError(fmt.Errorf("check signature fail"))
	}
	if !valid {
		return nil, BadSignatureError()
	}

	_, err = s.dao.GetPerpetualByAddress(strings.ToLower(params.PerpetualAddress), true)
	if err != nil {
		if dao.IsRecordNotFound(err) {
			return nil, PerpetualNotFoundError(params.PerpetualAddress)
		}
		return nil, InternalError(err)
	}

	// check ChainID
	ctx, cancel := context.WithTimeout(s.ctx, conf.Conf.BlockChain.Timeout.Duration)
	defer cancel()
	chainID, err := s.chainCli.GetChainID(ctx)
	if chainID.Int64() != params.ChainID {
		return nil, ChainIDError(params.ChainID)
	}

	_, err = s.dao.GetOrder(params.OrderHash)
	if err == nil {
		return nil, OrderHashExistError(params.OrderHash)
	} else if !dao.IsRecordNotFound(err) {
		return nil, InternalError(errors.New("get order fail"))
	}

	errID := s.match.NewOrder(&order)
	switch errID {
	case model.MatchOK:
		return nil, nil
	case model.MatchInternalErrorID:
		return nil, InternalError(err)
	case model.MatchMaxOrderNumReachID:
		return nil, MaxOrderNumReachError()
	case model.MatchInsufficientBalanceErrorID:
		return nil, InsufficientBalanceError()
	case model.MatchGasNotEnoughErrorID:
		return nil, GasBalanceError()
	case model.MatchCloseOnlyErrorID:
		return nil, CloseOnlyError()
	default:
		return nil, InternalError(errors.New("unknown match error"))
	}
}

func validatePlaceOrder(req *PlaceOrderReq) error {
	// Amount
	amount, err := decimal.NewFromString(req.Amount)
	if err != nil {
		return InvalidPriceAmountError(fmt.Sprintf("parse amount[%s] error", req.Amount))
	}
	if amount.Equal(decimal.Zero) {
		return InvalidPriceAmountError("amount = 0")
	}

	price, err := decimal.NewFromString(req.Price)
	if err != nil {
		return InvalidPriceAmountError(fmt.Sprintf("parse price[%s] error", req.Price))
	}

	if price.LessThanOrEqual(decimal.Zero) {
		return InvalidPriceAmountError("price <= 0")
	}

	// order dealine
	expiresAt := time.Unix(req.ExpiresAt, 0).UTC()
	now := time.Now().UTC()
	if now.After(expiresAt) {
		return OrderExpired()
	}

	// Price OrderType
	if req.OrderType == int(model.StopLimitOrder) {
		stopPrice, err := decimal.NewFromString(req.StopPrice)
		if err != nil {
			return InvalidPriceAmountError(fmt.Sprintf("parse price[%s] error", req.StopPrice))
		}

		if stopPrice.LessThanOrEqual(decimal.Zero) {
			return InvalidPriceAmountError("stop price <= 0")
		}
	} else if req.OrderType != int(model.LimitOrder) {
		return InternalError(errors.New("order type must be limit/stop-limit"))
	}

	// broker contract address
	if strings.ToLower(req.BrokerAddress) != strings.ToLower(conf.Conf.BrokerAddress) {
		return BrokerAddressError(req.BrokerAddress)
	}

	return nil
}

func (s *Server) CancelOrder(p Param) (interface{}, error) {
	params := p.(*CancelOrderReq)
	order, err := s.dao.GetOrder(params.OrderHash)
	if err != nil {
		if dao.IsRecordNotFound(err) {
			return nil, OrderIDNotExistError(params.OrderHash)
		}
		return nil, InternalError(err)
	}

	if order.TraderAddress != strings.ToLower(params.Address) {
		return nil, OrderAuthError(params.OrderHash)
	}

	if err = s.match.CancelOrder(order.PerpetualAddress, order.OrderHash); err != nil {
		return nil, InternalError(err)
	}
	return nil, nil
}

func (s *Server) CancelAllOrders(p Param) (interface{}, error) {
	params := p.(*CancelAllOrdersReq)
	if err := s.match.CancelAllOrders(params.PerpetualAddress, params.Address); err != nil {
		return nil, InternalError(err)
	}
	return nil, nil
}
