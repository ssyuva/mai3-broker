package match

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mcarloai/mai-v3-broker/common/chain"
	"github.com/mcarloai/mai-v3-broker/common/message"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/common/orderbook"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type match struct {
	ctx       context.Context
	mu        sync.Mutex
	wsChan    chan interface{}
	msgChan   chan interface{}
	orderbook *orderbook.Orderbook
	stopbook  *orderbook.Orderbook
	perpetual *model.Perpetual
	chainCli  chain.ChainClient
	dao       dao.DAO
	timers    map[string]*time.Timer
}

func newMatch(ctx context.Context, cli chain.ChainClient, dao dao.DAO, perpetual *model.Perpetual, wsChan, msgChan chan interface{}) *match {
	return &match{
		ctx:       ctx,
		wsChan:    wsChan,
		msgChan:   msgChan,
		perpetual: perpetual,
		orderbook: orderbook.NewOrderbook(),
		stopbook:  orderbook.NewOrderbook(),
		chainCli:  cli,
		dao:       dao,
		timers:    make(map[string]*time.Timer),
	}
}

func (m *match) run() error {
	if err := m.reloadOrderBook(); err != nil {
		return err
	}

	// go recieve match message
	go m.startConsumer()

	// go monitor check user margin gas
	go m.checkOrdersMargin()

	// go match order
	go m.runMatch()

	<-m.ctx.Done()
	return nil
}

func (m *match) startConsumer() {
	for {
		select {
		case <-m.ctx.Done():
			logger.Infof("Match Consumer Exit")
			return
		case msg, ok := <-m.msgChan:
			if !ok {
				return
			}
			m.parseMessage(msg.(message.MatchMessage))
		}
	}
}

func (m *match) parseMessage(msg message.MatchMessage) {
	switch msg.Type {
	case message.MatchTypeNewOrder:
		payload := msg.Payload.(message.MatchNewOrderPayload)
		order, err := m.dao.GetOrder(payload.OrderHash)
		if err != nil {
			logger.Errorf("Match New Order error perpetual:%s, orderHash:%s", msg.PerpetualAddress, payload.OrderHash)
			return
		}
		if err = m.insertNewOrder(order); err != nil {
			logger.Errorf("Insert New Order error perpetual:%s, orderHash:%s", msg.PerpetualAddress, payload.OrderHash)
		}
	case message.MatchTypeCancelOrder:
		payload := msg.Payload.(message.MatchCancelOrderPayload)
		if err := m.cancelOrder(payload.OrderHash, model.CancelReasonUserCancel, true, decimal.Zero); err != nil {
			logger.Errorf("Cancel Order error perpetual:%s, orderHash:%s", msg.PerpetualAddress, payload.OrderHash)
		}
	case message.MatchTypeChangeOrder:
		payload := msg.Payload.(message.MatchChangeOrderPayload)
		if err := m.changeOrder(payload.OrderHash, payload.Amount); err != nil {
			logger.Errorf("Match Change Order error perpetual:%s", msg.PerpetualAddress)
		}
	default:
		logger.Errorf("Match unknown message type:%s, perpetual:%s", msg.Type, msg.PerpetualAddress)
	}
}

func (m *match) checkOrdersMargin() {
	for {
		select {
		case <-m.ctx.Done():
			return
		case <-time.After(10 * time.Second):
			// TODO
			// check margin
			// check gas
		}
	}
}

func (m *match) matchStopOrders(indexPrice decimal.Decimal) {
	m.mu.Lock()
	defer m.mu.Unlock()
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for {
			minBidPrice := m.stopbook.MinBid()
			if minBidPrice != nil && minBidPrice.LessThanOrEqual(indexPrice) {
				orders := m.stopbook.GetBidOrdersByPrice(*minBidPrice)
				for _, order := range orders {
					m.changeStopOrder(order)
				}
			} else {
				break
			}
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for {
			MaxAskPrice := m.stopbook.MaxAsk()
			if MaxAskPrice != nil && MaxAskPrice.GreaterThanOrEqual(indexPrice) {
				orders := m.stopbook.GetAskOrdersByPrice(*MaxAskPrice)
				for _, order := range orders {
					m.changeStopOrder(order)
				}
			} else {
				break
			}
		}
		wg.Done()
	}()

	wg.Wait()
	return
}

func (m *match) changeStopOrder(memoryOrder *orderbook.MemoryOrder) {
	err := m.dao.Transaction(func(dao dao.DAO) error {
		dao.ForUpdate()
		order, err := dao.GetOrder(memoryOrder.ID)
		if err != nil {
			return err
		}
		order.Status = model.OrderPending
		if err = dao.UpdateOrder(order); err != nil {
			return err
		}

		if err := m.stopbook.RemoveOrder(memoryOrder); err != nil {
			return err
		}

		memoryOrder.ComparePrice = memoryOrder.Price
		if err := m.orderbook.InsertOrder(memoryOrder); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logger.Errorf("change stop order to pending order fail! err:%s", err)
	}
}

func (m *match) matchOrders(indexPrice decimal.Decimal) {
	m.mu.Lock()
	defer m.mu.Unlock()
	transactions, err := m.dao.QueryUnconfirmedTransactionsByContract(m.perpetual.PerpetualAddress)
	if err != nil {
		logger.Errorf("Match: QueryUnconfirmedTransactionsByContract failed perpetual:%s error:%s", m.perpetual.PerpetualAddress, err.Error())
		return
	}
	if len(transactions) > 0 {
		logger.Errorf("Match: unconfirmed transaction exists. wait for it to be confirmed perpetual:%s", m.perpetual.PerpetualAddress)
		return
	}
	//TODO
	// 1. compute match orders by index price
	matchItems := m.orderbook.MatchOrder(indexPrice)
	if len(matchItems) == 0 {
		return
	}

	u, err := uuid.NewRandom()
	if err != nil {
		logger.Errorf("generate transaction uuid error:%s", err.Error())
	}
	matchTransaction := &model.MatchTransaction{
		ID:               u.String(),
		Status:           model.TransactionStatusInit,
		PerpetualAddress: m.perpetual.PerpetualAddress,
	}
	err = m.dao.Transaction(func(dao dao.DAO) error {
		dao.ForUpdate()
		for _, item := range matchItems {
			order, err := dao.GetOrder(item.Order.ID)
			if err != nil {
				return fmt.Errorf("Match: Get Order[%s] failed, error:%w", item.Order.ID, err)
			}
			newAmount := order.AvailableAmount.Sub(item.MatchedAmount)
			if newAmount.IsNegative() {
				return fmt.Errorf("Match: order[%s] avaliable is negative after match", order.OrderHash)
			}
			order.AvailableAmount = newAmount
			order.PendingAmount = order.PendingAmount.Add(item.MatchedAmount)
			matchTransaction.MatchResult.MatchItems = append(matchTransaction.MatchResult.MatchItems, &model.MatchItem{
				OrderHash: order.OrderHash,
				Amount:    item.MatchedAmount,
			})
			if err := dao.UpdateOrder(order); err != nil {
				return fmt.Errorf("Match: order[%s] update failed error:%w", order.OrderHash, err)
			}
			if err := m.orderbook.ChangeOrder(item.Order, item.MatchedAmount.Neg()); err != nil {
				return fmt.Errorf("Match: order[%s] orderbook ChangeOrder failed error:%w", order.OrderHash, err)
			}
		}
		if err := dao.CreateMatchTransaction(matchTransaction); err != nil {
			return fmt.Errorf("Match: matchTransaction create failed error:%w", err)
		}
		return nil
	})

	if err == nil {
		// 4. send ws msg
	}

	return
}

func (m *match) runMatch() {
	for {
		select {
		case <-m.ctx.Done():
			return
		case <-time.After(time.Second):
			ctxTimeout, ctxTimeoutCancel := context.WithTimeout(m.ctx, conf.Conf.BlockChain.Timeout.Duration)
			defer ctxTimeoutCancel()
			indexPrice, err := m.chainCli.GetPrice(ctxTimeout, m.perpetual.OracleAddress)
			if err != nil {
				logger.Errorf("get index price fail! err:%s", err.Error())
				continue
			}
			m.matchStopOrders(indexPrice)
			m.matchOrders(indexPrice)
		}
	}
}

func (m *match) reloadOrderBook() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	orders, err := m.dao.QueryOrder("", m.perpetual.PerpetualAddress, []model.OrderStatus{model.OrderPending, model.OrderStop}, 0, 0, 0)
	if err != nil {
		return err
	}
	for _, order := range orders {
		if err := m.insertNewOrder(order); err != nil {
			return err
		}
	}
	return nil
}

func (m *match) insertNewOrder(order *model.Order) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	memoryOrder := &orderbook.MemoryOrder{
		ID:               order.OrderHash,
		PerpetualAddress: order.PerpetualAddress,
		Price:            order.Price,
		StopPrice:        order.StopPrice,
		Amount:           order.Amount,
		Type:             order.Type,
		Trader:           order.TraderAddress,
		GasFeeAmount:     order.GasFeeAmount,
	}
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
}

func (m *match) cancelOrder(orderID string, reason model.CancelReasonType, cancelAll bool, cancelAmount decimal.Decimal) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	var order *model.Order
	cancelBookAmount := cancelAmount
	cancelDBAmount := cancelAmount
	err := m.dao.Transaction(func(dao dao.DAO) error {
		var err error
		dao.ForUpdate()
		order, err = dao.GetOrder(orderID)
		if err != nil {
			return err
		}
		bookOrder, ok := m.orderbook.GetOrder(orderID, order.Amount.IsNegative(), order.Price)
		if !ok {
			if !order.AvailableAmount.IsPositive() {
				logger.Warnf("cancel order:order[%s] is closed.", orderID)
				order = nil
				return nil
			} else {
				logger.Errorf("cancel order: order[%s] not exists in book, cancel all in db!", orderID)
				cancelDBAmount = order.AvailableAmount
				cancelBookAmount = decimal.Zero
			}
		} else {
			if !order.AvailableAmount.Equal(bookOrder.Amount) {
				logger.Errorf("cancel order: order[%s] amount mismatch between db[%s] and book[%s] cancel all!", orderID, order.AvailableAmount, bookOrder.Amount)
				cancelAll = true
			}
			if cancelAll {
				cancelDBAmount = order.AvailableAmount
				cancelBookAmount = bookOrder.Amount
			}
			if order.AvailableAmount.LessThan(cancelAmount) {
				logger.Warnf("cancel amount[%s] larger than db available amount[%s]", cancelAmount, order.AvailableAmount)
				cancelDBAmount = order.AvailableAmount
			}
			if bookOrder.Amount.LessThan(cancelAmount) {
				logger.Warnf("cancel amount[%s] larger than book available amount[%s]", cancelAmount, bookOrder.Amount)
				cancelBookAmount = order.AvailableAmount
			}
		}

		if cancelBookAmount.IsPositive() {
			if err := m.orderbook.ChangeOrder(bookOrder, cancelBookAmount.Neg()); err != nil {
				return err
			}
		}

		if cancelDBAmount.IsPositive() {
			if err = model.CancelOrder(order, reason, cancelDBAmount); err != nil {
				return err
			}

			if err = dao.UpdateOrder(order); err != nil {
				return err
			}
		}

		m.deleteOrderTimer(orderID)
		return nil
	})
	if err == nil && order != nil && cancelDBAmount.IsPositive() {
		// notice websocket for cancel order
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

func (m *match) changeOrder(orderID string, changeAmount decimal.Decimal) error {
	order, err := m.dao.GetOrder(orderID)
	if err != nil {
		return err
	}
	bookOrder, ok := m.orderbook.GetOrder(orderID, order.Amount.IsNegative(), order.Price)
	if ok {
		if err = m.orderbook.ChangeOrder(bookOrder, changeAmount); err != nil {
			return err
		}
	} else {
		if err = m.orderbook.InsertOrder(bookOrder); err != nil {
			return err
		}
	}
	return nil
}
