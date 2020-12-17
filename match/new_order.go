package match

import (
	"context"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/message"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
)

const MAX_ORDER_NUM = 20

func (m *match) NewOrder(order *model.Order) string {
	m.mu.Lock()
	defer m.mu.Unlock()

	account, err := m.chainCli.GetMarginAccount(m.ctx, m.perpetual.PerpetualAddress, order.TraderAddress)
	if err != nil {
		logger.Errorf("GetMarginAccount err:%s", err)
		return model.MatchInternalErrorID
	}

	if !m.CheckCloseOnly(account, order) {
		return model.MatchCloseOnlyErrorID
	}

	if !m.CheckOrderMargin(account, order) {
		return model.MatchInsufficientBalanceErrorID
	}

	activeOrders, err := m.dao.QueryOrder(order.TraderAddress, order.PerpetualAddress, []model.OrderStatus{model.OrderPending, model.OrderStop}, 0, 0, 0)
	if err != nil {
		logger.Errorf("QueryOrder err:%s", err)
		return model.MatchInternalErrorID
	}

	if len(activeOrders) >= MAX_ORDER_NUM {
		return model.MatchMaxOrderNumReachID
	}

	// check gas
	gasBalance, err := m.chainCli.GetGasBalance(m.ctx, conf.Conf.BrokerAddress, order.TraderAddress)
	if err != nil {
		logger.Errorf("checkUserPendingOrders:%w", err)
		return model.MatchInternalErrorID
	}
	gasReward := m.gasMonitor.GetGasPrice() * 1e9 * conf.Conf.GasStation.GasLimit * uint64(len(activeOrders)+1)
	if utils.ToWad(gasBalance).LessThan(decimal.NewFromInt(int64(gasReward))) {
		return model.MatchGasNotEnoughErrorID
	}

	if order.BrokerFeeLimit.LessThan(decimal.NewFromInt(int64(gasReward))) {
		return model.MatchGasNotEnoughErrorID
	}

	// create order and insert to db and orderbook
	err = m.dao.Transaction(context.Background(), false /* readonly */, func(dao dao.DAO) error {
		dao.ForUpdate()
		if err := dao.CreateOrder(order); err != nil {
			return err
		}

		memoryOrder := m.getMemoryOrder(order)
		if order.Status == model.OrderPending {
			memoryOrder.SortKey = order.Price
			if err := m.orderbook.InsertOrder(memoryOrder); err != nil {
				return err
			}
		} else {
			memoryOrder.SortKey = order.StopPrice
			if err := m.stopbook.InsertOrder(memoryOrder); err != nil {
				return err
			}
		}
		if err := m.setExpirationTimer(order.OrderHash, order.ExpiresAt); err != nil {
			return err
		}
		return nil
	})

	if err == nil {
		// notice websocket for new order
		wsMsg := message.WebSocketMessage{
			ChannelID: message.GetAccountChannelID(order.TraderAddress),
			Payload: message.WebSocketOrderChangePayload{
				Type:  message.WsTypeOrderChange,
				Order: order,
			},
		}
		m.wsChan <- wsMsg
		return model.MatchOK
	}
	return model.MatchInternalErrorID
}
