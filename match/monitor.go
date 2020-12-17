package match

import (
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/conf"
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
			m.checkPerpUserOrders()
		}
	}
}

func (m *match) checkPerpUserOrders() {
	users, err := m.dao.GetPendingOrderUsers(m.perpetual.PerpetualAddress, []model.OrderStatus{model.OrderPending, model.OrderStop})
	if err != nil {
		logger.Errorf("checkOrdersMargin: GetPendingOrderUsers %s", err)
		return
	}
	for _, user := range users {
		cancels := m.checkUserPendingOrders(user)
		for _, cancel := range cancels {
			err := m.CancelOrder(cancel.OrderHash, model.CancelReasonAdminCancel, true, cancel.ToCancel)
			if err != nil {
				logger.Errorf("cancel Order fail! err:%s", err)
			}
		}
	}
}

func (m *match) checkUserPendingOrders(user string) []*OrderCancel {
	cancels := make([]*OrderCancel, 0)
	account, err := m.chainCli.GetMarginAccount(m.ctx, m.perpetual.PerpetualAddress, user)
	if err != nil {
		return cancels
	}
	gasBalance, err := m.chainCli.GetGasBalance(m.ctx, conf.Conf.BrokerAddress, user)
	if err != nil {
		logger.Errorf("checkUserPendingOrders:%w", err)
		return cancels
	}
	// check order margin and close Only order
	orders, err := m.dao.QueryOrder(user, m.perpetual.PerpetualAddress, []model.OrderStatus{model.OrderPending, model.OrderStop}, 0, 0, 0)
	if err != nil {
		logger.Errorf("checkUserPendingOrders:%w", err)
		return cancels
	}

	// gas check
	gasReward := m.gasMonitor.GetGasPrice() * 1e9 * conf.Conf.GasStation.GasLimit
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

	for _, order := range orders[num:] {
		if !m.CheckCloseOnly(account, order) {
			cancel := &OrderCancel{
				OrderHash: order.OrderHash,
				Status:    order.Status,
				ToCancel:  order.AvailableAmount,
			}
			cancels = append(cancels, cancel)
			continue
		}

		if !m.CheckOrderMargin(account, order) {
			cancel := &OrderCancel{
				OrderHash: order.OrderHash,
				Status:    order.Status,
				ToCancel:  order.AvailableAmount,
			}
			cancels = append(cancels, cancel)
		}

		if order.BrokerFeeLimit.LessThan(decimal.NewFromInt(int64(gasReward))) {
			cancel := &OrderCancel{
				OrderHash: order.OrderHash,
				Status:    order.Status,
				ToCancel:  order.AvailableAmount,
			}
			cancels = append(cancels, cancel)
		}
	}

	return cancels
}
