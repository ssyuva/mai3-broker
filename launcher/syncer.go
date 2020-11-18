package launcher

import (
	"context"
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/chain"
	"github.com/mcarloai/mai-v3-broker/common/message"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/pkg/errors"
	logger "github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v3"
	"sync"
	"time"
)

type Syncer struct {
	ctx       context.Context
	dao       dao.DAO
	chainCli  chain.ChainClient
	matchChan chan interface{}
	syncChan  chan interface{}
}

func NewSyncer(ctx context.Context, dao dao.DAO, chainCli chain.ChainClient, matchChan, syncChan chan interface{}) *Syncer {
	return &Syncer{
		ctx:       ctx,
		dao:       dao,
		chainCli:  chainCli,
		matchChan: matchChan,
		syncChan:  syncChan,
	}
}

func (s *Syncer) Run() {
	for {
		select {
		case <-s.ctx.Done():
			logger.Infof("Syncer stop")
			return
		case <-s.syncChan:
		case <-time.After(5 * time.Second):
		}
		if err := s.syncTransaction(); err != nil {
			logger.Errorf("syncTransaction error:%s", err)
		}
	}
}

func (s *Syncer) syncTransaction() error {
	users, err := s.dao.GetUsersWithStatus(model.TxPending)
	if err != nil {
		logger.Warnf("find user with pending status failed: %s", err)
		return nil
	}
	wg := &sync.WaitGroup{}
	for _, user := range users {
		wg.Add(1)
		go func(user string) {
			defer wg.Done()
			s.updateStatusByUser(user)
		}(user)
	}
	wg.Wait()
	return nil
}

func (s *Syncer) updateStatusByUser(user string) {
	logger.Debugf("check pending transaction for %s", user)
	txs, err := s.dao.GetTxsByUser(user, model.TxPending)
	if dao.IsRecordNotFound(err) || len(txs) == 0 {
		return
	}
	if err != nil {
		logger.Infof("fail to get pending transaction err:%s", err)
		return
	}

	ctx, done := context.WithTimeout(s.ctx, conf.Conf.BlockChain.Timeout.Duration)
	defer done()

	for i, tx := range txs {
		if tx.TransactionHash == nil {
			logger.Errorf("transaction hash is nill txID:%s", tx.ID)
			return
		}
		undoOrderMsgs := make([]message.MatchMessage, 0)
		err = s.dao.Transaction(func(dao dao.DAO) error {
			dao.ForUpdate()
			receipt, err := s.chainCli.WaitTransactionReceipt(ctx, *tx.TransactionHash)
			if err != nil {
				return err
			}
			tx.BlockNumber = &receipt.BlockNumber
			tx.BlockHash = &receipt.BlockHash
			tx.BlockTime = &receipt.BlockTime
			tx.Status = receipt.Status
			tx.GasUsed = &receipt.GasUsed
			// check block header
			bh, err := dao.FindBlock(conf.Conf.WatcherID, int64(*tx.BlockNumber))
			if err != nil {
				return errors.Wrap(err, "get block header fail")
			}
			if bh.BlockHash != *tx.BlockHash {
				return fmt.Errorf("block hash check failed, blocknumber:%d", *tx.BlockNumber)
			}
			if err = dao.UpdateTx(tx); err != nil {
				return errors.Wrap(err, "fail to update transaction status")
			}
			// handle tx with same nonce
			candidates, err := dao.GetTxsByNonce(tx.FromAddress, tx.Nonce)
			if err != nil {
				return errors.Wrap(err, "fail to find transaction by nonce")
			}
			for _, candidate := range candidates {
				if *candidate.TransactionHash != *tx.TransactionHash {
					candidate.Status = model.TxCanceled
					dao.UpdateTx(candidate)
				}
			}
			// update match_transaction
			matchTx, err := dao.GetMatchTransaction(tx.TxID)
			if err != nil {
				return err
			}
			matchTx.BlockConfirmed = true
			matchTx.BlockHash = null.StringFrom(*tx.BlockHash)
			matchTx.BlockNumber = null.IntFrom(int64(*tx.BlockNumber))
			matchTx.ExecutedAt = null.TimeFrom(time.Unix(int64(*tx.BlockTime), 0).UTC())
			matchTx.Status = tx.Status.TransactionStatus()

			// update orders
			matchEvents, err := s.chainCli.FilterMatch(ctx, matchTx.PerpetualAddress, *tx.BlockNumber, *tx.BlockNumber)
			if err != nil {
				return err
			}
			orderSucc := make(map[string]bool)
			for _, event := range matchEvents {
				matchInfo := &model.MatchItem{
					OrderHash: event.OrderHash,
					Amount:    event.Amount,
				}
				matchTx.MatchResult.ReceiptItems = append(matchTx.MatchResult.ReceiptItems, matchInfo)
				order, err := dao.GetOrder(event.OrderHash)
				if err != nil {
					return err
				}
				order.PendingAmount = order.PendingAmount.Sub(event.Amount)
				order.ConfirmedAmount = order.ConfirmedAmount.Add(event.Amount)
				if err := dao.UpdateOrder(order); err != nil {
					return err
				}
				orderSucc[event.OrderHash] = true
			}

			for _, item := range matchTx.MatchResult.MatchItems {
				if _, ok := orderSucc[item.OrderHash]; ok {
					continue
				}
				// order failed
				order, err := dao.GetOrder(item.OrderHash)
				if err != nil {
					return err
				}
				order.PendingAmount = order.PendingAmount.Sub(item.Amount)
				order.AvailableAmount = order.AvailableAmount.Add(item.Amount)
				if err := dao.UpdateOrder(order); err != nil {
					return err
				}
				matchMsg := message.MatchMessage{
					Type:             message.MatchTypeChangeOrder,
					PerpetualAddress: order.PerpetualAddress,
					Payload: message.MatchChangeOrderPayload{
						OrderHash: order.OrderHash,
						Amount:    item.Amount,
					},
				}
				undoOrderMsgs = append(undoOrderMsgs, matchMsg)
			}

			if err = dao.UpdateMatchTransaction(matchTx); err != nil {
				return err
			}

			return nil
		})
		for _, msg := range undoOrderMsgs {
			s.matchChan <- msg
		}
		// this case is to handle accelarate
		if next := i + 1; next < len(txs) && *tx.Nonce == *txs[next].Nonce {
			continue
		}
		if err != nil {
			logger.Warnf("fail to check status: %s", err)
			return
		}
	}
}
