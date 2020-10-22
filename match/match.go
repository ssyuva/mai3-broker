package match

import (
	"context"
	"github.com/mcarloai/mai-v3-broker/common/message"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/common/orderbook"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/micro/go-micro/v2/logger"
	"time"
)

type match struct {
	ctx       context.Context
	wsChan    chan interface{}
	msgChan   chan interface{}
	orderbook *orderbook.Orderbook
	stopbook  *orderbook.Orderbook
	perpetual *model.Perpetual
	dao       dao.DAO
	timers    map[string]*time.Timer
}

func newMatch(ctx context.Context, dao dao.DAO, perpetual *model.Perpetual, wsChan, msgChan chan interface{}) *match {
	return &match{
		ctx:       ctx,
		wsChan:    wsChan,
		msgChan:   msgChan,
		perpetual: perpetual,
		orderbook: orderbook.NewOrderbook(),
		stopbook:  orderbook.NewOrderbook(),
		dao:       dao,
		timers:    make(map[string]*time.Timer),
	}
}

func (m *match) run() error {
	if err := m.initOrderBook(); err != nil {
		return err
	}

	// go recieve match message
	go m.startConsumer()

	// go getOraclePrice
	go m.getOraclePrice()

	// go monitor check user margin gas
	go m.checkOrdersMargin()

	// go check stop order
	go m.checkStopOrders()

	// go check order
	go m.checkOrders()

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
	case message.MatchTypeCancelOrder:
	case message.MatchTypeReloadOrder:
	default:
		logger.Errorf("Match unknown message type:%s, perpetual:%s", msg.Type, msg.PerpetualAddress)
	}
}

func (m *match) getOraclePrice() error {
	return nil
}

func (m *match) checkOrdersMargin() error {
	// TODO
	// check margin
	// check gas
	return nil
}

func (m *match) checkStopOrders() error {
	return nil
}

func (m *match) checkOrders() error {
	return nil
}

func (m *match) initOrderBook() error {
	orders, err := m.dao.QueryOrder("", m.perpetual.PerpetualAddress, []model.OrderStatus{model.OrderPending, model.OrderStop}, 0, 0, 0)
	if err != nil {
		return err
	}
	for _, order := range orders {
		memoryOrder := &orderbook.MemoryOrder{
			ID:               order.OrderHash,
			PerpetualAddress: order.PerpetualAddress,
			Price:            order.Price,
			StopPrice:        order.StopPrice,
			Amount:           order.Amount,
			Type:             order.Type,
			Side:             order.Side,
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
	}
	return nil
}
