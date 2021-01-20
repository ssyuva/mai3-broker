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
	"time"
)

type Monitor struct {
	ctx      context.Context
	dao      dao.DAO
	runner   *runnable.Timed
	chainCli chain.ChainClient
	match    *match.Server
}

const MATURE_BLOCKNUM = 60

func NewMonitor(ctx context.Context, dao dao.DAO, chainCli chain.ChainClient, match *match.Server) *Monitor {
	return &Monitor{
		ctx:      ctx,
		dao:      dao,
		runner:   runnable.NewTimed(ChannelHWM),
		chainCli: chainCli,
		match:    match,
	}
}

func (s *Monitor) Run() error {
	// check unmature confirmed transaction for rollback
	return s.runner.Run(s.ctx, time.Second*10, s.syncUnmatureTransaction)
}

func (s *Monitor) syncUnmatureTransaction() {
	ctx, done := context.WithTimeout(s.ctx, conf.Conf.Timeout)
	defer done()
	blockNumber, err := s.chainCli.GetLatestBlockNumber(ctx)
	if err != nil {
		logger.Infof("GetLatestBlockNumber error:%s", err)
		return
	}
	begin := blockNumber - MATURE_BLOCKNUM
	s.updateUnmatureTransactionStatus(&begin)
	return
}

func (s *Monitor) updateUnmatureTransactionStatus(blockNumber *uint64) {
	// get unmature confirmed transaction
	logger.Debugf("check unmature transaction blockNumber > %d", *blockNumber)
	txs, err := s.dao.GetTxsByBlock(blockNumber, nil, model.TxSuccess, model.TxFailed)
	if dao.IsRecordNotFound(err) || len(txs) == 0 {
		return
	}
	if err != nil {
		logger.Infof("fail to get pending transaction err:%s", err)
		return
	}

	ctx, done := context.WithTimeout(s.ctx, conf.Conf.Timeout)
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
