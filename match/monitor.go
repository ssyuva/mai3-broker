package match

import (
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
	"strings"
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
	users, err := m.dao.GetPendingOrderUsers(m.perpetual.LiquidityPoolAddress, m.perpetual.PerpetualIndex, []model.OrderStatus{model.OrderPending})
	if err != nil {
		logger.Errorf("checkOrdersMargin: GetPendingOrderUsers %s", err)
		return
	}
	poolStorage, err := m.chainCli.GetLiquidityPoolStorage(m.ctx, conf.Conf.ReaderAddress, m.perpetual.LiquidityPoolAddress)
	if err != nil {
		logger.Errorf("matchOrders: GetLiquidityPoolStorage fail! err:%s", err.Error())
		return
	}
	for _, user := range users {
		cancels := m.checkUserPendingOrders(poolStorage, user)
		for _, cancel := range cancels {
			err := m.CancelOrder(cancel.OrderHash, model.CancelReasonAdminCancel, true, cancel.ToCancel)
			if err != nil {
				logger.Errorf("cancel Order fail! err:%s", err)
			}
		}
	}
}

func (m *match) checkUserPendingOrders(poolStorage *model.LiquidityPoolStorage, user string) []*OrderCancel {
	cancels := make([]*OrderCancel, 0)
	account, err := m.chainCli.GetAccountStorage(m.ctx, conf.Conf.ReaderAddress, m.perpetual.PerpetualIndex, m.perpetual.LiquidityPoolAddress, user)
	if err != nil {
		return cancels
	}
	gasBalance, err := m.chainCli.GetGasBalance(m.ctx, conf.Conf.BrokerAddress, user)
	if err != nil {
		logger.Errorf("checkUserPendingOrders:%w", err)
		return cancels
	}
	// check order margin and close Only order
	orders, err := m.dao.QueryOrder(user, m.perpetual.LiquidityPoolAddress, m.perpetual.PerpetualIndex, []model.OrderStatus{model.OrderPending}, 0, 0, 0)
	if err != nil {
		logger.Errorf("checkUserPendingOrders:%w", err)
		return cancels
	}

	// gas check
	num := 0
	gasReward := m.gasMonitor.GetGasPrice() * conf.Conf.GasStation.GasLimit
	if gasReward > 0 {
		maxOrderNum := utils.ToGwei(gasBalance).IntPart() / int64(gasReward)
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
	}

	for _, order := range orders[num:] {
		if order.OrderParam.BrokerAddress != strings.ToLower(conf.Conf.BrokerAddress) {
			cancel := &OrderCancel{
				OrderHash: order.OrderHash,
				Status:    order.Status,
				ToCancel:  order.AvailableAmount,
			}
			cancels = append(cancels, cancel)
			continue
		}
		if !m.CheckCloseOnly(account, order) {
			cancel := &OrderCancel{
				OrderHash: order.OrderHash,
				Status:    order.Status,
				ToCancel:  order.AvailableAmount,
			}
			cancels = append(cancels, cancel)
			continue
		}

		if !m.CheckOrderMargin(poolStorage, account, order) {
			cancel := &OrderCancel{
				OrderHash: order.OrderHash,
				Status:    order.Status,
				ToCancel:  order.AvailableAmount,
			}
			cancels = append(cancels, cancel)
		}

		if order.BrokerFeeLimit < int64(gasReward) {
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
