package match

import (
	"github.com/mcarloai/mai-v3-broker/common/message"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/dao"
)

func (m *match) NewOrder(order *model.Order) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	err := m.dao.Transaction(func(dao dao.DAO) error {
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