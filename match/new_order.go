package match

import (
	"context"
	"time"

	"github.com/mcarloai/mai-v3-broker/common/mai3"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/message"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
)

const MAX_ORDER_NUM = 10

func (m *match) NewOrder(order *model.Order) string {
	m.mu.Lock()
	defer m.mu.Unlock()

	now := time.Now()
	account, err := m.chainCli.GetAccountStorage(m.ctx, conf.Conf.ReaderAddress, m.perpetual.PerpetualIndex, m.perpetual.LiquidityPoolAddress, order.TraderAddress)
	if err != nil {
		logger.Errorf("new order:GetAccountStorage err:%s", err)
		return model.MatchInternalErrorID
	}

	logger.Infof("GetAccountStorage 1111111: used:%d", time.Since(now).Milliseconds())

	if !m.CheckCloseOnly(account, order) {
		return model.MatchCloseOnlyErrorID
	}

	now = time.Now()
	poolStorage, err := m.chainCli.GetLiquidityPoolStorage(m.ctx, conf.Conf.ReaderAddress, m.perpetual.LiquidityPoolAddress)
	if err != nil {
		logger.Errorf("new order: GetLiquidityPoolStorage fail! err:%s", err.Error())
		return model.MatchInternalErrorID
	}
	logger.Infof("GetLiquidityPoolStorage 1111111: used:%d", time.Since(now).Milliseconds())

	perpetual, ok := poolStorage.Perpetuals[m.perpetual.PerpetualIndex]
	if !ok || !perpetual.IsNormal {
		logger.Errorf("new order: perpetual status is not normal!")
		return model.MatchInternalErrorID
	}

	now = time.Now()
	if !m.CheckOrderMargin(poolStorage, account, order, order.AvailableAmount) {
		return model.MatchInsufficientBalanceErrorID
	}
	logger.Infof("CheckOrderMargin 1111111: used:%d", time.Since(now).Milliseconds())

	activeOrders, err := m.dao.QueryOrder(order.TraderAddress, order.LiquidityPoolAddress, order.PerpetualIndex, []model.OrderStatus{model.OrderPending}, 0, 0, 0)
	if err != nil {
		logger.Errorf("new order: QueryOrder err:%s", err)
		return model.MatchInternalErrorID
	}

	if len(activeOrders) >= MAX_ORDER_NUM {
		return model.MatchMaxOrderNumReachID
	}

	// check gas
	order.GasFeeLimit = mai3.GetGasFeeLimit(len(poolStorage.Perpetuals))
	if conf.Conf.GasEnable {
		gasBalance, err := m.chainCli.GetGasBalance(m.ctx, conf.Conf.BrokerAddress, order.TraderAddress)
		if err != nil {
			logger.Errorf("new order: checkUserPendingOrders:%w", err)
			return model.MatchInternalErrorID
		}
		gasReward := m.gasMonitor.GetGasPriceDecimal().Mul(decimal.NewFromInt(order.GasFeeLimit))
		ordersGasReword := gasReward.Mul(decimal.NewFromInt(int64(len(activeOrders) + 1)))
		if gasBalance.LessThan(ordersGasReword) {
			return model.MatchGasNotEnoughErrorID
		}

		if decimal.NewFromInt(order.BrokerFeeLimit).LessThan(utils.ToGwei(gasReward)) {
			return model.MatchGasNotEnoughErrorID
		}
	}

	// create order and insert to db and orderbook
	err = m.dao.Transaction(context.Background(), false /* readonly */, func(dao dao.DAO) error {
		if err := dao.CreateOrder(order); err != nil {
			return err
		}

		memoryOrder := m.getMemoryOrder(order)
		if err := m.orderbook.InsertOrder(memoryOrder); err != nil {
			return err
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
	logger.Errorf("new order: %s", err)
	return model.MatchInternalErrorID
}
