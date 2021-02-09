package match

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mcarloai/mai-v3-broker/common/chain"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/message"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/common/orderbook"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/mcarloai/mai-v3-broker/gasmonitor"
	logger "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"sync"
	"time"
)

type match struct {
	ctx        context.Context
	cancel     context.CancelFunc
	mu         sync.Mutex
	wsChan     chan interface{}
	orderbook  *orderbook.Orderbook
	perpetual  *model.Perpetual
	chainCli   chain.ChainClient
	gasMonitor *gasmonitor.GasMonitor
	dao        dao.DAO
	timers     map[string]*time.Timer
}

func newMatch(ctx context.Context, cli chain.ChainClient, dao dao.DAO, perpetual *model.Perpetual, wsChan chan interface{}, gm *gasmonitor.GasMonitor) *match {
	ctx, cancel := context.WithCancel(ctx)
	return &match{
		ctx:        ctx,
		cancel:     cancel,
		wsChan:     wsChan,
		perpetual:  perpetual,
		orderbook:  orderbook.NewOrderbook(),
		chainCli:   cli,
		gasMonitor: gm,
		dao:        dao,
		timers:     make(map[string]*time.Timer),
	}
}

func (m *match) Run() error {
	logger.Infof("Match Perpetual:%s-%d start", m.perpetual.LiquidityPoolAddress, m.perpetual.PerpetualIndex)
	if err := m.reloadActiveOrders(); err != nil {
		return err
	}

	group, ctx := errgroup.WithContext(m.ctx)
	// go monitor check user margin gas
	group.Go(func() error {
		return m.checkOrdersMargin(ctx)
	})

	// go match order
	group.Go(func() error {
		return m.runMatch(ctx)
	})

	err := group.Wait()
	logger.Infof("Match Perpetual:%s-%d end", m.perpetual.LiquidityPoolAddress, m.perpetual.PerpetualIndex)
	return err
}

func (m *match) Close() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.stopTimers()
	m.cancel()
}

func (m *match) matchOrders() {
	m.mu.Lock()
	defer m.mu.Unlock()
	transactions, err := m.dao.QueryUnconfirmedTransactionsByContract(m.perpetual.LiquidityPoolAddress, m.perpetual.PerpetualIndex)
	if err != nil {
		logger.Errorf("Match: QueryUnconfirmedTransactionsByContract failed perpetual:%s-%d error:%s", m.perpetual.LiquidityPoolAddress, m.perpetual.PerpetualIndex, err.Error())
		return
	}
	if len(transactions) > 0 {
		logger.Infof("Match: unconfirmed transaction exists. wait for it to be confirmed perpetual:%s-%d", m.perpetual.LiquidityPoolAddress, m.perpetual.PerpetualIndex)
		return
	}

	matchItems := m.MatchOrderSideBySide()
	if len(matchItems) == 0 {
		return
	}

	u, err := uuid.NewRandom()
	if err != nil {
		logger.Errorf("generate transaction uuid error:%s", err.Error())
	}
	matchTransaction := &model.MatchTransaction{
		ID:                   u.String(),
		Status:               model.TransactionStatusInit,
		LiquidityPoolAddress: m.perpetual.LiquidityPoolAddress,
		PerpetualIndex:       m.perpetual.PerpetualIndex,
		BrokerAddress:        conf.Conf.BrokerAddress,
	}
	orders := make([]*model.Order, 0)
	err = m.dao.Transaction(context.Background(), false /* readonly */, func(dao dao.DAO) error {
		for _, item := range matchItems {
			order, err := dao.GetOrder(item.Order.ID)
			if err != nil {
				return fmt.Errorf("Match: Get Order[%s] failed, error:%w", item.Order.ID, err)
			}
			newAmount := order.AvailableAmount.Sub(item.MatchedAmount)
			if !utils.HasTheSameSign(newAmount, order.AvailableAmount) {
				return fmt.Errorf("Match: order[%s] avaliable is negative after match", order.OrderHash)
			}
			order.AvailableAmount = newAmount
			order.PendingAmount = order.PendingAmount.Add(item.MatchedAmount)
			if !item.OrderTotalCancel.IsZero() {
				order.CanceledAmount = order.CanceledAmount.Add(item.OrderTotalCancel)
				order.AvailableAmount = order.AvailableAmount.Sub(item.OrderTotalCancel)
				for i := 0; i < len(item.OrderCancelAmounts); i++ {
					order.CancelReasons = append(order.CancelReasons, &model.OrderCancelReason{
						Reason:     item.OrderCancelReasons[i],
						Amount:     item.OrderCancelAmounts[i],
						CanceledAt: time.Now().UTC(),
					})
				}
			}
			matchTransaction.MatchResult.MatchItems = append(matchTransaction.MatchResult.MatchItems, &model.MatchItem{
				OrderHash: order.OrderHash,
				Amount:    item.MatchedAmount,
			})
			if err := dao.UpdateOrder(order); err != nil {
				return fmt.Errorf("Match: order[%s] update failed error:%w", order.OrderHash, err)
			}
			if err := m.orderbook.ChangeOrder(item.Order, item.MatchedAmount.Add(item.OrderTotalCancel).Neg()); err != nil {
				return fmt.Errorf("Match: order[%s] orderbook ChangeOrder failed error:%w", order.OrderHash, err)
			}
			orders = append(orders, order)
		}
		if err := dao.CreateMatchTransaction(matchTransaction); err != nil {
			return fmt.Errorf("Match: matchTransaction create failed error:%w", err)
		}
		return nil
	})

	if err == nil {
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
	} else {
		logger.Errorf("match orders fail. error:%s", err)
	}

	return
}

func (m *match) runMatch(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			logger.Infof("match perpetual:%s-%d matchOrders end", m.perpetual.LiquidityPoolAddress, m.perpetual.PerpetualIndex)
			return nil
		case <-time.After(conf.Conf.MatchInterval):
			m.matchOrders()
		}
	}
}

func (m *match) reloadActiveOrders() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	orders, err := m.dao.QueryOrder("", m.perpetual.LiquidityPoolAddress, m.perpetual.PerpetualIndex, []model.OrderStatus{model.OrderPending}, 0, 0, 0)
	if err != nil {
		return err
	}
	for _, order := range orders {
		if !order.AvailableAmount.IsZero() {
			memoryOrder := m.getMemoryOrder(order)
			if err := m.orderbook.InsertOrder(memoryOrder); err != nil {
				return fmt.Errorf("reloadActiveOrders:%w", err)
			}

			if err := m.setExpirationTimer(order.OrderHash, order.ExpiresAt); err != nil {
				return fmt.Errorf("reloadActiveOrders:%w", err)
			}
		}
	}
	return nil
}

func (m *match) getMemoryOrder(order *model.Order) *orderbook.MemoryOrder {
	return &orderbook.MemoryOrder{
		ID:                   order.OrderHash,
		LiquidityPoolAddress: order.LiquidityPoolAddress,
		PerpetualIndex:       order.PerpetualIndex,
		Price:                order.Price,
		TriggerPrice:         order.TriggerPrice,
		Amount:               order.AvailableAmount,
		MinTradeAmount:       order.MinTradeAmount,
		Type:                 order.Type,
		Trader:               order.TraderAddress,
	}
}
