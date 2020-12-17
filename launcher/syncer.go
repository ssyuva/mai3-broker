package launcher

import (
	"context"
	"github.com/mcarloai/mai-v3-broker/common/chain"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/mcarloai/mai-v3-broker/match"
	"github.com/pkg/errors"
	logger "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type Syncer struct {
	ctx      context.Context
	dao      dao.DAO
	chainCli chain.ChainClient
	match    *match.Server
	syncChan chan interface{}
}

const MATURE_BLOCKNUM = 60

func NewSyncer(ctx context.Context, dao dao.DAO, chainCli chain.ChainClient, syncChan chan interface{}, match *match.Server) *Syncer {
	return &Syncer{
		ctx:      ctx,
		dao:      dao,
		chainCli: chainCli,
		syncChan: syncChan,
		match:    match,
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
		// check unmature confirmed transaction for rollback
		if err := s.syncUnmatureTransaction(); err != nil {
			logger.Errorf("syncUnMatureTransaction error:%s", err)
		}

		if err := s.syncTransaction(); err != nil {
			logger.Errorf("syncTransaction error:%s", err)
		}
	}
}

func (s *Syncer) syncUnmatureTransaction() error {
	ctx, done := context.WithTimeout(s.ctx, conf.Conf.BlockChain.Timeout.Duration)
	defer done()
	blockNumber, err := s.chainCli.GetLatestBlockNumber(ctx)
	if err != nil {
		return err
	}
	begin := blockNumber - MATURE_BLOCKNUM
	users, err := s.dao.GetUsersWithBlockNumber(begin, model.TxSuccess, model.TxFailed)
	if err != nil {
		logger.Warnf("find user with pending status failed: %s", err)
		return nil
	}
	wg := &sync.WaitGroup{}
	for _, user := range users {
		wg.Add(1)
		go func(user string) {
			defer wg.Done()
			s.updateUnmatureStatusByUser(user, &begin)
		}(user)
	}
	wg.Wait()
	return nil
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
	txs, err := s.dao.GetTxsByUser(user, nil, model.TxPending)
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
			logger.Errorf("transaction hash is nill txID:%s", tx.TxID)
			return
		}
		receipt, err := s.chainCli.WaitTransactionReceipt(ctx, *tx.TransactionHash)
		if err != nil {
			logger.Errorf("WaitTransactionReceipt error: %s", err)
			continue
		}
		err = s.dao.Transaction(context.Background(), false /* readonly */, func(dao dao.DAO) error {
			dao.ForUpdate()
			tx.BlockNumber = &receipt.BlockNumber
			tx.BlockHash = &receipt.BlockHash
			tx.BlockTime = &receipt.BlockTime
			tx.Status = receipt.Status
			tx.GasUsed = &receipt.GasUsed
			// bh, err := dao.FindBlock(conf.Conf.WatcherID, int64(*tx.BlockNumber))
			// if err != nil {
			// 	return errors.Wrap(err, "get block header fail")
			// }
			// if bh.BlockHash != *tx.BlockHash {
			// 	return fmt.Errorf("block hash check failed, blocknumber:%d", *tx.BlockNumber)
			// }
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

			err = s.match.UpdateOrdersStatus(tx.TxID, tx.Status.TransactionStatus(), *tx.TransactionHash, *tx.BlockHash, *tx.BlockNumber, *tx.BlockTime)
			return err
		})
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

func (s *Syncer) updateUnmatureStatusByUser(user string, blockNumber *uint64) {
	// get unmature confirmed transaction
	logger.Debugf("check unmature transaction for %s", user)
	txs, err := s.dao.GetTxsByUser(user, blockNumber, model.TxSuccess, model.TxFailed)
	if dao.IsRecordNotFound(err) || len(txs) == 0 {
		return
	}
	if err != nil {
		logger.Infof("fail to get pending transaction err:%s", err)
		return
	}

	ctx, done := context.WithTimeout(s.ctx, conf.Conf.BlockChain.Timeout.Duration)
	defer done()

	for _, tx := range txs {
		if tx.TransactionHash == nil || tx.BlockNumber == nil {
			logger.Errorf("transaction hash or blockNumber is nill txID:%s", tx.TxID)
			return
		}
		receipt, err := s.chainCli.WaitTransactionReceipt(ctx, *tx.TransactionHash)
		if err != nil {
			logger.Errorf("WaitTransactionReceipt error: %s", err)
			continue
		}

		// check blockNumber blockHash ?
		if tx.Status == receipt.Status {
			continue
		}
		err = s.dao.Transaction(context.Background(), false /* readonly */, func(dao dao.DAO) error {
			dao.ForUpdate()
			tx.BlockNumber = &receipt.BlockNumber
			tx.BlockHash = &receipt.BlockHash
			tx.BlockTime = &receipt.BlockTime
			tx.Status = receipt.Status
			tx.GasUsed = &receipt.GasUsed
			if err = dao.UpdateTx(tx); err != nil {
				return errors.Wrap(err, "fail to update transaction status")
			}

			err = s.match.RollbackOrdersStatus(tx.TxID, tx.Status.TransactionStatus(), *tx.TransactionHash, *tx.BlockHash, *tx.BlockNumber, *tx.BlockTime)
			return err
		})
		if err != nil {
			logger.Warnf("fail to check status: %s", err)
			return
		}
	}
}
