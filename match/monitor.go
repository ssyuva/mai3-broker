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

type OrderCancel struct {
	OrderHash string
	Status    model.OrderStatus
	ToCancel  decimal.Decimal
	Reason    model.CancelReasonType
}

func (m *match) checkOrdersMargin() {
	for {
		select {
		case <-m.ctx.Done():
			return
		case <-time.After(conf.Conf.MatchMonitorInterval):
			m.checkPerpUserOrders()
		}
	}
}

func (m *match) checkPerpUserOrders() {
	users, err := m.dao.GetPendingOrderUsers(m.perpetual.LiquidityPoolAddress, m.perpetual.PerpetualIndex, []model.OrderStatus{model.OrderPending})
	if err != nil {
		logger.Errorf("monitor: GetPendingOrderUsers %s", err)
		return
	}
	if len(users) == 0 {
		return
	}
	poolStorage, err := m.chainCli.GetLiquidityPoolStorage(m.ctx, conf.Conf.ReaderAddress, m.perpetual.LiquidityPoolAddress)
	if err != nil {
		logger.Errorf("monitor: GetLiquidityPoolStorage fail! err:%s", err.Error())
		return
	}

	// close perpetual if perpetual status is not normal
	perpetual, ok := poolStorage.Perpetuals[m.perpetual.PerpetualIndex]
	if !ok || !perpetual.IsNormal {
		m.perpetual.IsPublished = false
		if err := m.dao.UpdatePerpetual(m.perpetual); err != nil {
			logger.Errorf("closePerpetual error:%s", err)
		}
	}

	for _, user := range users {
		cancels := m.checkUserPendingOrders(poolStorage, user)
		for _, cancel := range cancels {
			err := m.CancelOrder(cancel.OrderHash, cancel.Reason, true, cancel.ToCancel)
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

	// close orders if perpetual status is not normal
	perpetual, ok := poolStorage.Perpetuals[m.perpetual.PerpetualIndex]
	if !ok || !perpetual.IsNormal {
		for _, order := range orders {
			cancel := &OrderCancel{
				OrderHash: order.OrderHash,
				Status:    order.Status,
				ToCancel:  order.AvailableAmount,
				Reason:    model.CancelReasonContractSettled,
			}
			cancels = append(cancels, cancel)
		}
	}

	// gas check
	num := 0
	gasReward := m.gasMonitor.GetGasPrice() * conf.Conf.GasLimit
	if gasReward > 0 {
		maxOrderNum := utils.ToGwei(gasBalance).IntPart() / int64(gasReward)
		if len(orders) > int(maxOrderNum) {
			num := len(orders) - int(maxOrderNum)
			for _, order := range orders[0:num] {
				cancel := &OrderCancel{
					OrderHash: order.OrderHash,
					Status:    order.Status,
					ToCancel:  order.AvailableAmount,
					Reason:    model.CancelReasonGasNotEnough,
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
				Reason:    model.CancelReasonAdminCancel,
			}
			cancels = append(cancels, cancel)
			continue
		}

		if !m.CheckCloseOnly(account, order) {
			cancel := &OrderCancel{
				OrderHash: order.OrderHash,
				Status:    order.Status,
				ToCancel:  order.AvailableAmount,
				Reason:    model.CancelReasonAdminCancel,
			}
			cancels = append(cancels, cancel)
			continue
		}

		if !m.CheckOrderMargin(poolStorage, account, order) {
			cancel := &OrderCancel{
				OrderHash: order.OrderHash,
				Status:    order.Status,
				ToCancel:  order.AvailableAmount,
				Reason:    model.CancelReasonInsufficientFunds,
			}
			cancels = append(cancels, cancel)
		}

		if order.BrokerFeeLimit < int64(gasReward) {
			cancel := &OrderCancel{
				OrderHash: order.OrderHash,
				Status:    order.Status,
				ToCancel:  order.AvailableAmount,
				Reason:    model.CancelReasonGasNotEnough,
			}
			cancels = append(cancels, cancel)
		}
	}

	return cancels
}
