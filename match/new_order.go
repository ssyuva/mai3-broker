package match

import (
	"context"

	"github.com/mcarloai/mai-v3-broker/common/mai3"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/message"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
)

func (m *match) NewOrder(order *model.Order) string {
	m.mu.Lock()
	defer m.mu.Unlock()

	// get all pending orders (multiple perpetuals maybe)
	activeOrders, err := m.dao.QueryOrder(order.TraderAddress, "", 0, []model.OrderStatus{model.OrderPending}, 0, 0, 0)
	if err != nil {
		logger.Errorf("new order: QueryOrder err:%s", err)
		return model.MatchInternalErrorID
	}

	if len(activeOrders) >= conf.Conf.MaxOrderNum {
		return model.MatchMaxOrderNumReachID
	}

	account, err := m.chainCli.GetAccountStorage(m.ctx, conf.Conf.ReaderAddress, m.perpetual.PerpetualIndex, m.perpetual.LiquidityPoolAddress, order.TraderAddress)
	if account == nil || err != nil {
		logger.Errorf("new order:GetAccountStorage err:%v", err)
		return model.MatchInternalErrorID
	}

	balance, err := m.chainCli.BalanceOf(m.ctx, m.perpetual.CollateralAddress, order.TraderAddress, m.perpetual.CollateralDecimals)
	if err != nil {
		logger.Errorf("new order:BalanceOf err:%v", err)
		return model.MatchInternalErrorID
	}
	allowance, err := m.chainCli.Allowance(m.ctx, m.perpetual.CollateralAddress, order.TraderAddress, m.perpetual.LiquidityPoolAddress, m.perpetual.CollateralDecimals)
	if err != nil {
		logger.Errorf("new order:Allowance err:%v", err)
		return model.MatchInternalErrorID
	}
	walletBalance := decimal.Min(balance, allowance)

	if !CheckCloseOnly(account, order).Equal(_0) {
		return model.MatchCloseOnlyErrorID
	}

	// get from cache
	poolStorage := m.poolSyncer.GetPoolStorage(m.perpetual.LiquidityPoolAddress)
	if poolStorage == nil {
		logger.Errorf("new order: GetLiquidityPoolStorage fail! err:%v", err)
		return model.MatchInternalErrorID
	}

	perpetual, ok := poolStorage.Perpetuals[m.perpetual.PerpetualIndex]
	if !ok || !perpetual.IsNormal {
		logger.Errorf("new order: perpetual status is not normal!")
		return model.MatchInternalErrorID
	}

	// check gas
	order.GasFeeLimit = mai3.GetGasFeeLimit(len(poolStorage.Perpetuals))
	if conf.Conf.GasEnable {
		gasBalance, err := m.chainCli.GetGasBalance(m.ctx, conf.Conf.BrokerAddress, order.TraderAddress)
		if err != nil {
			logger.Errorf("new order: checkUserPendingOrders:%v", err)
			return model.MatchInternalErrorID
		}
		gasReward := m.gasMonitor.GasPriceGwei().Mul(decimal.NewFromInt(order.GasFeeLimit))
		ordersGasReword := gasReward.Mul(decimal.NewFromInt(int64(len(activeOrders) + 1)))
		if gasBalance.LessThan(ordersGasReword) {
			return model.MatchGasNotEnoughErrorID
		}

		if decimal.NewFromInt(order.BrokerFeeLimit).LessThan(utils.ToGwei(gasReward)) {
			return model.MatchGasNotEnoughErrorID
		}
	}

	activeOrders = append(activeOrders, order)
	orderMaps, err := m.splitActiveOrdersInMultiPerpetuals(activeOrders)
	if err != nil {
		return model.MatchInternalErrorID
	}

	for _, v := range orderMaps {
		var a *model.AccountStorage
		// get account storage
		if v.LiquidityPoolAddress == m.perpetual.LiquidityPoolAddress && v.PerpetualIndex == m.perpetual.PerpetualIndex {
			a = account
		} else {
			a, err = m.chainCli.GetAccountStorage(m.ctx, conf.Conf.ReaderAddress, v.PerpetualIndex, v.LiquidityPoolAddress, order.TraderAddress)
			if account == nil || err != nil {
				logger.Errorf("new order:GetAccountStorage err:%v", err)
				return model.MatchInternalErrorID
			}
		}
		a.WalletBalance = walletBalance
		// TODO: consider cancel orders
		_, available := ComputeOrderAvailable(v.PoolStorage, v.PerpetualIndex, a, v.Orders)
		// available balance less than 0, InsufficientBalance
		if available.LessThan(_0) {
			return model.MatchInsufficientBalanceErrorID
		}
		walletBalance = available
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
