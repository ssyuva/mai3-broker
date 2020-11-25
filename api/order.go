package api

import (
	// "context"
	"errors"
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/mai3"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/message"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/dao"
	// "github.com/mcarloai/mai-v3-broker/perp"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

const TIMESTAMP_RANGE = 5 * 60

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

	order := &model.Order{}
	order.OrderHash = params.OrderHash
	order.OrderParam.TraderAddress = strings.ToLower(params.Address)
	if params.OrderType == int(model.LimitOrder) {
		order.OrderParam.Type = model.LimitOrder
		order.Status = model.OrderPending
	} else {
		order.OrderParam.Type = model.StopLimitOrder
		order.Status = model.OrderStop
	}

	amount, _ := decimal.NewFromString(params.Amount)
	if amount.LessThan(decimal.Zero) {
		order.OrderParam.Amount = amount.Neg()
		order.OrderParam.Side = model.SideSell
	} else {
		order.OrderParam.Amount = amount
		order.OrderParam.Side = model.SideBuy
	}

	order.Version = int32(mai3.ProtocolV3)
	order.OrderParam.ChainID = params.ChainID
	order.OrderParam.Price, _ = decimal.NewFromString(params.Price)
	order.OrderParam.StopPrice, _ = decimal.NewFromString(params.StopPrice)
	order.OrderParam.IsCloseOnly = params.IsCloseOnly
	expiresAt := getExpiresAt(params.Expires)
	order.OrderParam.ExpiresAt = time.Unix(expiresAt, 0).UTC()
	order.OrderParam.Salt = params.Salt
	order.PerpetualAddress = strings.ToLower(params.PerpetualAddress)
	order.BrokerAddress = strings.ToLower(params.BrokerAddress)
	order.RelayerAddress = strings.ToLower(params.RelayerAddress)
	if params.ReferrerAddress == "" {
		order.ReferrerAddress = ADDRESS_ZERO
	}
	order.AvailableAmount = order.OrderParam.Amount
	order.ConfirmedAmount = decimal.Zero
	order.CanceledAmount = decimal.Zero
	order.PendingAmount = decimal.Zero
	now := time.Now().UTC()
	order.CreatedAt = now
	order.UpdatedAt = now

	// check orderhash
	orderHash, err := mai3.GetOrderHash(order.TraderAddress, order.BrokerAddress, order.RelayerAddress, order.PerpetualAddress, order.ReferrerAddress,
		amount, order.Price, expiresAt, order.Version, int8(order.Type), order.IsCloseOnly, order.OrderParam.Salt, order.OrderParam.ChainID)
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

	_, err = s.dao.GetPerpetualByAddress(strings.ToLower(params.PerpetualAddress))
	if err != nil {
		if dao.IsRecordNotFound(err) {
			return nil, PerpetualNotFoundError(params.PerpetualAddress)
		}
		return nil, InternalError(err)
	}

	_, err = s.dao.GetOrder(params.OrderHash)
	if err == nil {
		return nil, OrderHashExistError(params.OrderHash)
	} else if !dao.IsRecordNotFound(err) {
		return nil, InternalError(err)
	}

	// get accountStorage
	// ctx, cancel := context.WithTimeout(s.ctx, conf.Conf.BlockChain.Timeout.Duration)
	// defer cancel()
	// accountContext, err := s.chainCli.GetMarginAccount(ctx, order.PerpetualAddress, order.TraderAddress)
	// if err != nil {
	// 	return nil, InternalError(err)
	// }

	// // TODO perp storage
	// orderContext := &perp.OrderContext{Account: accountContext}

	// check margin balance
	// actives, err := s.dao.QueryOrder(order.TraderAddress, order.PerpetualAddress, []model.OrderStatus{model.OrderPending, model.OrderStop}, 0, 0, 0)
	// if err != nil {
	// 	return nil, InternalError(fmt.Errorf("QueryOrder:%w", err))
	// }
	//3. close only

	if err := s.dao.CreateOrder(order); err != nil {
		return nil, InternalError(err)
	}

	// notice websocket for new order
	wsMsg := message.WebSocketMessage{
		ChannelID: message.GetAccountChannelID(order.TraderAddress),
		Payload: message.WebSocketOrderChangePayload{
			Type:  message.WsTypeOrderChange,
			Order: order,
		},
	}
	s.wsChan <- wsMsg
	// notice match for new order
	message := message.MatchMessage{
		PerpetualAddress: order.PerpetualAddress,
		Type:             message.MatchTypeNewOrder,
		Payload: message.MatchNewOrderPayload{
			OrderHash: order.OrderHash,
		},
	}
	s.matchChan <- message
	return PlaceOrderResp{OrderHash: order.OrderHash}, nil
}

func (s *Server) checkActiveOrders(orders []*model.Order) error {
	return nil
}

func getExpiresAt(expiresInSeconds int64) int64 {
	now := time.Now().UTC()
	expire := time.Second * time.Duration(expiresInSeconds)

	return now.Add(expire).Unix()
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

	// order sign timestamp
	now := time.Now().UTC().Unix()
	if (now-TIMESTAMP_RANGE) > req.Timestamp || (now+TIMESTAMP_RANGE) < req.Timestamp {
		return InternalError(errors.New("timestamp must in 5 min"))
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
		return nil, InternalError(err)
	}

	// notice match for cancel order
	message := message.MatchMessage{
		PerpetualAddress: order.PerpetualAddress,
		Type:             message.MatchTypeCancelOrder,
		Payload: message.MatchCancelOrderPayload{
			OrderHash: order.OrderHash,
		},
	}
	s.matchChan <- message
	return nil, nil
}

func (s *Server) CancelAllOrders(p Param) (interface{}, error) {
	params := p.(*CancelAllOrdersReq)
	orders, err := s.dao.QueryOrder(params.Address, params.PerpetualAddress, []model.OrderStatus{model.OrderPending, model.OrderStop}, 0, 0, 0)
	if err != nil {
		return nil, InternalError(err)
	}

	for _, order := range orders {
		// notice match for cancel order
		message := message.MatchMessage{
			PerpetualAddress: order.PerpetualAddress,
			Type:             message.MatchTypeCancelOrder,
			Payload: message.MatchCancelOrderPayload{
				OrderHash: order.OrderHash,
			},
		}
		s.matchChan <- message
	}
	return nil, nil
}
