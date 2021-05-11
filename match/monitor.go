package match

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
)

var OrderAmountRelaxFactor = decimal.NewFromFloat(0.1)

type OrderCancel struct {
	OrderHash string
	Status    model.OrderStatus
	ToCancel  decimal.Decimal
	Reason    model.CancelReasonType
}

func (m *match) checkActiveOrders(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			logger.Infof("match perpetual:%s-%d monitor end", m.perpetual.LiquidityPoolAddress, m.perpetual.PerpetualIndex)
			return nil
		case <-time.After(conf.Conf.MatchMonitorInterval):
			m.checkPerpUserOrders()
		}
	}
}

func (m *match) checkPerpUserOrders() {
	orders, err := m.dao.QueryOrder("", m.perpetual.LiquidityPoolAddress, m.perpetual.PerpetualIndex, []model.OrderStatus{model.OrderPending}, 0, 0, 0)
	if err != nil {
		return
	}

	// update active orders count in metrics
	activeOrderCount.WithLabelValues(fmt.Sprintf("%s-%d", m.perpetual.LiquidityPoolAddress, m.perpetual.PerpetualIndex)).Set(float64(len(orders)))
	if len(orders) == 0 {
		return
	}

	users, err := m.dao.GetPendingOrderUsers(m.perpetual.LiquidityPoolAddress, m.perpetual.PerpetualIndex, []model.OrderStatus{model.OrderPending})
	if err != nil {
		logger.Errorf("monitor: GetPendingOrderUsers %s", err)
		return
	}
	if len(users) == 0 {
		return
	}
	poolStorage := m.poolSyncer.GetPoolStorage(m.perpetual.LiquidityPoolAddress)
	if poolStorage == nil {
		logger.Errorf("monitor: GetLiquidityPoolStorage fail!")
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

	// check order margin and close Only order
	orders, err := m.dao.QueryOrder(user, m.perpetual.LiquidityPoolAddress, m.perpetual.PerpetualIndex, []model.OrderStatus{model.OrderPending}, 0, 0, 0)
	if err != nil {
		logger.Errorf("checkUserPendingOrders:%v", err)
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
		return cancels
	}

	account, err := m.chainCli.GetAccountStorage(m.ctx, conf.Conf.ReaderAddress, m.perpetual.PerpetualIndex, m.perpetual.LiquidityPoolAddress, user)
	if account == nil || err != nil {
		return cancels
	}
	balance, err := m.chainCli.BalanceOf(m.ctx, m.perpetual.CollateralAddress, user, m.perpetual.CollateralDecimals)
	if err != nil {
		return cancels
	}
	allowance, err := m.chainCli.Allowance(m.ctx, m.perpetual.CollateralAddress, user, m.perpetual.LiquidityPoolAddress, m.perpetual.CollateralDecimals)
	if err != nil {
		return cancels
	}
	account.WalletBalance = decimal.Min(balance, allowance)
	gasBalance, err := m.chainCli.GetGasBalance(m.ctx, conf.Conf.BrokerAddress, user)
	if err != nil {
		logger.Errorf("checkUserPendingOrders:%v", err)
		return cancels
	}

	gasPrice := m.gasMonitor.GasPriceGwei()
	gasReward := decimal.Zero
	remainOrders := make([]*model.Order, 0)
	for _, order := range orders {
		if conf.Conf.GasEnable {
			// gas check
			orderGasReward := gasPrice.Mul(decimal.NewFromInt(order.GasFeeLimit))
			if decimal.NewFromInt(order.BrokerFeeLimit).LessThan(utils.ToGwei(gasReward)) {
				cancel := &OrderCancel{
					OrderHash: order.OrderHash,
					Status:    order.Status,
					ToCancel:  order.AvailableAmount,
					Reason:    model.CancelReasonGasNotEnough,
				}
				cancels = append(cancels, cancel)
			}

			gasReward = gasReward.Add(orderGasReward)
			if gasBalance.LessThan(gasReward) {
				cancel := &OrderCancel{
					OrderHash: order.OrderHash,
					Status:    order.Status,
					ToCancel:  order.AvailableAmount,
					Reason:    model.CancelReasonGasNotEnough,
				}
				cancels = append(cancels, cancel)
				continue
			}
		}

		// because broker address was signed, if it is changed, orders need be canceled
		if order.OrderParam.BrokerAddress != strings.ToLower(conf.Conf.BrokerAddress) {
			cancel := &OrderCancel{
				OrderHash: order.OrderHash,
				Status:    order.Status,
				ToCancel:  order.AvailableAmount,
				Reason:    model.CancelReasonTransactionFail,
			}
			cancels = append(cancels, cancel)
			continue
		}

		cancelAmount := m.CheckCloseOnly(account, order)
		if !cancelAmount.Equal(_0) {
			cancel := &OrderCancel{
				OrderHash: order.OrderHash,
				Status:    order.Status,
				ToCancel:  cancelAmount,
				Reason:    model.CancelReasonCloseOnly,
			}
			cancels = append(cancels, cancel)
			continue
		}

		remainOrders = append(remainOrders, order)

	}

	// check remain orders available margin
	cancelsInsufficientFunds, _ := ComputeOrderAvailable(poolStorage, m.perpetual.PerpetualIndex, account, remainOrders)
	cancels = append(cancels, cancelsInsufficientFunds...)
	return cancels
}
