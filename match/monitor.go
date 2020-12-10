package match

import (
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/message"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
	"time"
)

func (m *match) checkOrdersMargin() {
	for {
		select {
		case <-m.ctx.Done():
			return
		case <-time.After(10 * time.Second):
			users, err := m.dao.GetPendingOrderUsers(m.perpetual.PerpetualAddress, []model.OrderStatus{model.OrderPending, model.OrderStop})
			if err != nil {
				logger.Errorf("checkOrdersMargin: GetPendingOrderUsers %s", err)
				continue
			}
			for _, user := range users {
				m.checkUserPendingOrders(user)
			}
		}
	}
}

func (m *match) checkUserPendingOrders(user string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	cancels := make([]*OrderCancel, 0)
	account, err := m.chainCli.GetMarginAccount(m.ctx, m.perpetual.PerpetualAddress, user)
	if err != nil {
		return
	}
	gasBalance, err := m.chainCli.GetBalance(m.ctx, user)
	if err != nil {
		logger.Errorf("checkUserPendingOrders:%w", err)
		return
	}
	// check order margin and close Only order
	orders, err := m.dao.QueryOrder(user, m.perpetual.PerpetualAddress, []model.OrderStatus{model.OrderPending, model.OrderStop}, 0, 0, 0)
	if err != nil {
		logger.Errorf("checkUserPendingOrders:%w", err)
		return
	}

	// gas check
	gasReward := m.priceMonitor.GetGasPrice() * 1e9 * conf.Conf.GasStation.GasLimit
	maxOrderNum := utils.ToWad(gasBalance).Div(decimal.NewFromInt(int64(gasReward))).IntPart()
	num := 0
	if len(orders) > int(maxOrderNum) {
		num := len(orders) - int(maxOrderNum)
		for _, order := range orders[0:num] {
			cancel := &OrderCancel{
				OrderHash: order.OrderHash,
				Status:    order.Status,
				ToCancel:  order.AvailableAmount,
			}
			cancels = append(cancels, cancel)
		}
	}

	closeCancles := m.CheckAndModifyCloseOnly(account, orders[num:])
	if len(closeCancles) > 0 {
		cancels = append(cancels, closeCancles...)
	}
	computeCancels, err := m.ComputeOrderMargin(account, orders[num:])
	if err != nil {
		logger.Errorf("checkUserPendingOrders:%w", err)
		return
	} else if len(computeCancels) > 0 {
		cancels = append(cancels, computeCancels...)
	}

	if len(cancels) == 0 {
		return
	}

	ordersForNotify := make([]*model.Order, 0)

	err = m.dao.Transaction(func(dao dao.DAO) error {
		for _, cancel := range cancels {
			order, err := dao.GetOrder(cancel.OrderHash)
			if err != nil {
				return err
			}
			order.AvailableAmount = order.AvailableAmount.Sub(cancel.ToCancel)
			order.CanceledAmount = order.CanceledAmount.Add(cancel.ToCancel)
			if err = dao.UpdateOrder(order); err != nil {
				return err
			}
			ordersForNotify = append(ordersForNotify, order)
		}
		return nil
	})

	if err != nil {
		// notice websocket for new order
		for _, order := range orders {
			wsMsg := message.WebSocketMessage{
				ChannelID: message.GetAccountChannelID(order.TraderAddress),
				Payload: message.WebSocketOrderChangePayload{
					Type:  message.WsTypeOrderChange,
					Order: order,
				},
			}
			m.wsChan <- wsMsg
		}
	}

}
