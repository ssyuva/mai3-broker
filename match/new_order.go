package match

import (
	"errors"
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/message"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
)

func (m *match) NewOrder(order *model.Order) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// check order margin and close Only order
	activeOrders, err := m.dao.QueryOrder(order.TraderAddress, order.PerpetualAddress, []model.OrderStatus{model.OrderPending, model.OrderStop}, 0, 0, 0)
	if err != nil {
		return fmt.Errorf("NewOrder:%w", err)
	}
	// check gas
	gasBalance, err := m.chainCli.GetGasBalance(m.ctx, conf.Conf.BrokerAddress, order.TraderAddress)
	if err != nil {
		logger.Errorf("checkUserPendingOrders:%w", err)
		return err
	}
	gasReward := m.priceMonitor.GetGasPrice() * 1e9 * conf.Conf.GasStation.GasLimit * uint64(len(activeOrders)+1)
	if utils.ToWad(gasBalance).LessThan(decimal.NewFromInt(int64(gasReward))) {
		return errors.New("Gas not enough")
	}

	_, err = m.CheckNewOrder(order, activeOrders)
	// update close only order and insert new order and orderbook
	err = m.dao.Transaction(func(dao dao.DAO) error {
		if err := dao.CreateOrder(order); err != nil {
			return err
		}

		memoryOrder := m.getMemoryOrder(order)
		if order.Status == model.OrderPending {
			memoryOrder.ComparePrice = order.Price
			if err := m.orderbook.InsertOrder(memoryOrder); err != nil {
				return err
			}
		} else {
			memoryOrder.ComparePrice = order.StopPrice
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
	}
	return err
}
