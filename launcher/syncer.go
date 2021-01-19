package launcher

import (
	"context"
	"github.com/mcarloai/mai-v3-broker/common/chain"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/mcarloai/mai-v3-broker/match"
	"github.com/mcarloai/mai-v3-broker/runnable"
	"github.com/pkg/errors"
	logger "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type Syncer struct {
	ctx      context.Context
	dao      dao.DAO
	runner   *runnable.Timed
	chainCli chain.ChainClient
	match    *match.Server
}

func NewSyncer(ctx context.Context, dao dao.DAO, chainCli chain.ChainClient, match *match.Server) *Syncer {
	return &Syncer{
		ctx:      ctx,
		dao:      dao,
		runner:   runnable.NewTimed(ChannelHWM),
		chainCli: chainCli,
		match:    match,
	}
}

func (s *Syncer) Run() error {
	// sync pending transaction
	return s.runner.Run(s.ctx, time.Second, s.syncTransaction)
}

func (s *Syncer) syncTransaction() {
	users, err := s.dao.GetUsersWithStatus(model.TxPending)
	if err != nil {
		logger.Warnf("find user with pending status failed: %s", err)
		return
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
	return
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
			logger.Errorf("syncer: transaction hash is nill txID:%s", tx.TxID)
			return
		}
		receipt, err := s.chainCli.WaitTransactionReceipt(ctx, *tx.TransactionHash)
		if err != nil {
			logger.Errorf("syncer: WaitTransactionReceipt error: %s", err)
			continue
		}
		err = s.dao.Transaction(context.Background(), false /* readonly */, func(dao dao.DAO) error {
			tx.BlockNumber = &receipt.BlockNumber
			tx.BlockHash = &receipt.BlockHash
			tx.BlockTime = &receipt.BlockTime
			tx.Status = receipt.Status
			tx.GasUsed = &receipt.GasUsed
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
