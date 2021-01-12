package match

import (
	"context"
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/message"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/shopspring/decimal"
	"gopkg.in/guregu/null.v3"
	"time"

	logger "github.com/sirupsen/logrus"
)

func (m *match) UpdateOrdersStatus(txID string, status model.TransactionStatus, transactionHash, blockHash string, blockNumber, blockTime uint64) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	ordersToNotify := make([]*model.Order, 0)
	err := m.dao.Transaction(context.Background(), false /* readonly */, func(dao dao.DAO) error {
		dao.ForUpdate()
		// update match_transaction
		matchTx, err := dao.GetMatchTransaction(txID)
		if err != nil {
			return err
		}
		matchTx.BlockConfirmed = true
		matchTx.Status = status
		matchTx.TransactionHash = null.StringFrom(transactionHash)
		if (matchTx.Status == model.TransactionStatusFail || matchTx.Status == model.TransactionStatusSuccess) && blockNumber > 0 {
			matchTx.BlockHash = null.StringFrom(blockHash)
			matchTx.BlockNumber = null.IntFrom(int64(blockNumber))
			matchTx.ExecutedAt = null.TimeFrom(time.Unix(int64(blockTime), 0).UTC())
		}

		// update orders
		orders, err := m.updateOrdersByTradeEvent(dao, matchTx, blockNumber)
		if err != nil {
			return err
		}
		if err = dao.UpdateMatchTransaction(matchTx); err != nil {
			return err
		}

		ordersToNotify = append(ordersToNotify, orders...)
		return nil
	})

	if err == nil {
		for _, order := range ordersToNotify {
			// notice websocket for order change
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
	return err
}

func (m *match) updateOrdersByTradeEvent(dao dao.DAO, matchTx *model.MatchTransaction, blockNumber uint64) ([]*model.Order, error) {
	ordersToNotify := make([]*model.Order, 0)
	ctxTimeout, ctxTimeoutCancel := context.WithTimeout(m.ctx, conf.Conf.BlockChain.Timeout.Duration)
	defer ctxTimeoutCancel()
	matchEvents, err := m.chainCli.FilterTradeSuccess(ctxTimeout, matchTx.BrokerAddress, blockNumber, blockNumber)
	if err != nil {
		return ordersToNotify, err
	}
	orderSuccMap := make(map[string]decimal.Decimal)
	orderFailMap := make(map[string]decimal.Decimal)
	orderMatchMap := make(map[string]decimal.Decimal)
	orderHashes := make([]string, 0)
	for _, event := range matchEvents {
		matchInfo := &model.MatchItem{
			OrderHash: event.OrderHash,
			Amount:    event.Amount,
		}
		matchTx.MatchResult.ReceiptItems = append(matchTx.MatchResult.ReceiptItems, matchInfo)
		orderSuccMap[event.OrderHash] = event.Amount
		orderHashes = append(orderHashes, event.OrderHash)
	}

	for _, item := range matchTx.MatchResult.MatchItems {
		orderMatchMap[item.OrderHash] = item.Amount
		if _, ok := orderSuccMap[item.OrderHash]; ok {
			continue
		}
		// order failed
		orderFailMap[item.OrderHash] = item.Amount
		orderHashes = append(orderHashes, item.OrderHash)
	}

	orders, err := dao.GetOrderByHashs(orderHashes)
	if err != nil {
		logger.Errorf("UpdateOrdersStatus:%s", err)
		return ordersToNotify, err
	}
	for _, order := range orders {
		// order success
		if amount, ok := orderSuccMap[order.OrderHash]; ok {
			// partial success
			matchAmount := orderMatchMap[order.OrderHash]
			if !amount.Equal(matchAmount) {
				oldAmount := order.AvailableAmount
				order.AvailableAmount = order.AvailableAmount.Add(matchAmount.Sub(amount))
				if err := m.rollbackOrderbook(oldAmount, order.PendingAmount.Sub(amount), order); err != nil {
					logger.Errorf("UpdateOrdersStatus:%s", err)
					return ordersToNotify, err
				}
			}
			order.PendingAmount = order.PendingAmount.Sub(amount)
			order.ConfirmedAmount = order.ConfirmedAmount.Add(amount)
			if err := dao.UpdateOrder(order); err != nil {
				logger.Errorf("UpdateOrdersStatus:%s", err)
				return ordersToNotify, err
			}
		} else if amount, ok := orderFailMap[order.OrderHash]; ok {
			// order failed, reload order in orderbook
			oldAmount := order.AvailableAmount
			order.PendingAmount = order.PendingAmount.Sub(amount)
			order.AvailableAmount = order.AvailableAmount.Add(amount)
			if err := dao.UpdateOrder(order); err != nil {
				logger.Errorf("UpdateOrdersStatus:%s", err)
				return ordersToNotify, err
			}

			if err := m.rollbackOrderbook(oldAmount, amount, order); err != nil {
				logger.Errorf("UpdateOrdersStatus:%s", err)
				return ordersToNotify, err
			}
		}
		ordersToNotify = append(ordersToNotify, order)
	}
	return ordersToNotify, nil
}

func (m *match) rollbackOrderbook(oldAmount, delta decimal.Decimal, order *model.Order) error {
	if oldAmount.IsZero() {
		memoryOrder := m.getMemoryOrder(order)
		if err := m.orderbook.InsertOrder(memoryOrder); err != nil {
			logger.Errorf("UpdateOrdersStatus:%w", err)
			return err
		}
		return nil
	}

	bookOrder, ok := m.orderbook.GetOrder(order.OrderHash, order.Amount.IsNegative(), order.Price)
	if ok {
		if err := m.orderbook.ChangeOrder(bookOrder, delta); err != nil {
			logger.Errorf("UpdateOrdersStatus:%w", err)
			return err
		}
		return nil
	}
	return fmt.Errorf("order %s not fund in orderbook", order.OrderHash)
}

func (m *match) RollbackOrdersStatus(txID string, status model.TransactionStatus, transactionHash, blockHash string, blockNumber, blockTime uint64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	ordersToNotify := make([]*model.Order, 0)
	err := m.dao.Transaction(context.Background(), false /* readonly */, func(dao dao.DAO) error {
		dao.ForUpdate()
		// update match_transaction
		matchTx, err := dao.GetMatchTransaction(txID)
		if err != nil {
			return err
		}
		matchTx.TransactionHash = null.StringFrom(transactionHash)
		matchTx.BlockHash = null.StringFrom(blockHash)
		matchTx.BlockNumber = null.IntFrom(int64(blockNumber))
		matchTx.ExecutedAt = null.TimeFrom(time.Unix(int64(blockTime), 0).UTC())

		if matchTx.Status == model.TransactionStatusSuccess && status == model.TransactionStatusFail {
			orders, err := m.rollbackOrdersOnTransactionFail(dao, matchTx)
			if err != nil {
				return err
			}
			ordersToNotify = append(ordersToNotify, orders...)
		} else if matchTx.Status == model.TransactionStatusFail && status == model.TransactionStatusSuccess {
			orders, err := m.updateOrdersByTradeEvent(dao, matchTx, blockNumber)
			if err != nil {
				return err
			}
			ordersToNotify = append(ordersToNotify, orders...)
		}
		matchTx.Status = status
		if err = dao.UpdateMatchTransaction(matchTx); err != nil {
			return err
		}

		return nil
	})

	if err == nil {
		for _, order := range ordersToNotify {
			// notice websocket for order change
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
	return err
}

func (m *match) rollbackOrdersOnTransactionFail(dao dao.DAO, matchTx *model.MatchTransaction) ([]*model.Order, error) {
	ordersToNotify := make([]*model.Order, 0)
	orderRollbackMap := make(map[string]decimal.Decimal)
	orderHashes := make([]string, 0)
	// success orders
	for _, item := range matchTx.MatchResult.ReceiptItems {
		orderRollbackMap[item.OrderHash] = item.Amount
		orderHashes = append(orderHashes, item.OrderHash)
	}

	orders, err := dao.GetOrderByHashs(orderHashes)
	if err != nil {
		logger.Errorf("rollbackOrdersOnTransactionFail:%w", err)
		return ordersToNotify, err
	}
	for _, order := range orders {
		amount := orderRollbackMap[order.OrderHash]
		oldAmount := order.AvailableAmount
		order.PendingAmount = order.ConfirmedAmount.Sub(amount)
		order.AvailableAmount = order.AvailableAmount.Add(amount)
		if err := dao.UpdateOrder(order); err != nil {
			logger.Errorf("rollbackOrdersOnTransactionFail:%w", err)
			return ordersToNotify, err
		}

		if err := m.rollbackOrderbook(oldAmount, amount, order); err != nil {
			logger.Errorf("rollbackOrdersOnTransactionFail:%w", err)
			return ordersToNotify, err
		}
		ordersToNotify = append(ordersToNotify, order)
	}

	return ordersToNotify, nil
}
