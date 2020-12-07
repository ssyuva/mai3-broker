package match

import (
	"context"
	"github.com/mcarloai/mai-v3-broker/common/message"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/shopspring/decimal"
	"gopkg.in/guregu/null.v3"
	"time"

	logger "github.com/sirupsen/logrus"
)

func (m *match) BatchTradeOrders(txID, status, transactionHash, blockHash string, blockNumber, blockTime uint64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	ordersForNotity := make([]*model.Order, 0)
	err := m.dao.Transaction(func(dao dao.DAO) error {
		dao.ForUpdate()
		// update match_transaction
		matchTx, err := dao.GetMatchTransaction(txID)
		if err != nil {
			return err
		}
		matchTx.BlockConfirmed = true
		matchTx.Status = model.TransactionStatus(status)
		matchTx.TransactionHash = null.StringFrom(transactionHash)
		if (matchTx.Status == model.TransactionStatusExcFail || matchTx.Status == model.TransactionStatusSuccess) && blockNumber > 0 {
			matchTx.BlockHash = null.StringFrom(blockHash)
			matchTx.BlockNumber = null.IntFrom(int64(blockNumber))
			matchTx.ExecutedAt = null.TimeFrom(time.Unix(int64(blockTime), 0).UTC())
		}

		// update orders
		ctxTimeout, ctxTimeoutCancel := context.WithTimeout(m.ctx, conf.Conf.BlockChain.Timeout.Duration)
		defer ctxTimeoutCancel()
		matchEvents, err := m.chainCli.FilterTradeSuccess(ctxTimeout, matchTx.BrokerAddress, blockNumber, blockNumber)
		if err != nil {
			return err
		}
		orderSuccMap := make(map[string]decimal.Decimal)
		orderFailMap := make(map[string]decimal.Decimal)
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
			if _, ok := orderSuccMap[item.OrderHash]; ok {
				continue
			}
			// order failed
			orderFailMap[item.OrderHash] = item.Amount
			orderHashes = append(orderHashes, item.OrderHash)
		}

		orders, err := dao.GetOrderByHashs(orderHashes)
		if err != nil {
			logger.Errorf("BatchTradeOrders:%w", err)
			return err
		}
		for _, order := range orders {
			if amount, ok := orderSuccMap[order.OrderHash]; ok {
				// order success
				order.PendingAmount = order.PendingAmount.Sub(amount)
				order.ConfirmedAmount = order.ConfirmedAmount.Add(amount)
				if err := dao.UpdateOrder(order); err != nil {
					logger.Errorf("BatchTradeOrders:%w", err)
					return err
				}
			} else if amount, ok := orderFailMap[order.OrderHash]; ok {
				// order failed, reload order in orderbook
				availableAmount := order.AvailableAmount
				order.PendingAmount = order.PendingAmount.Sub(amount)
				order.AvailableAmount = order.AvailableAmount.Add(amount)
				if err := dao.UpdateOrder(order); err != nil {
					logger.Errorf("BatchTradeOrders:%w", err)
					return err
				}
				if availableAmount.IsZero() {
					memoryOrder := m.getMemoryOrder(order)
					if err = m.orderbook.InsertOrder(memoryOrder); err != nil {
						logger.Errorf("BatchTradeOrders:%w", err)
						return err
					}
				} else {
					bookOrder, ok := m.orderbook.GetOrder(order.OrderHash, order.Amount.IsNegative(), order.Price)
					if ok {
						if err = m.orderbook.ChangeOrder(bookOrder, amount); err != nil {
							logger.Errorf("BatchTradeOrders:%w", err)
							return err
						}
					}
				}

			}

			ordersForNotity = append(ordersForNotity, order)
		}

		if err = dao.UpdateMatchTransaction(matchTx); err != nil {
			return err
		}

		return nil
	})

	if err == nil {
		for _, order := range ordersForNotity {
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
