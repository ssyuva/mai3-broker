package launcher

import (
	"context"
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/chain"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/common/utils"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/pkg/errors"
	logger "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type Syncer struct {
	ctx       context.Context
	dao       dao.DAO
	chainCli  chain.ChainClient
	rpcClient *utils.HttpClient
	syncChan  chan interface{}
}

func NewSyncer(ctx context.Context, dao dao.DAO, chainCli chain.ChainClient, syncChan chan interface{}, rpcClient *utils.HttpClient) *Syncer {
	return &Syncer{
		ctx:       ctx,
		dao:       dao,
		chainCli:  chainCli,
		syncChan:  syncChan,
		rpcClient: rpcClient,
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
			logger.Errorf("transaction hash is nill txID:%s", tx.TxID)
			return
		}
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

			var matchReq struct {
				TxID            string `json:"tx_id"`
				TransactionHash string `json:"transactionHash"`
				BlockNumber     uint64 `json:"blockNumber"`
				BlockHash       string `json:"blockHash"`
				BlockTime       uint64 `json:"blockTime"`
				Status          string `json:"status"`
			}
			matchReq.TxID = tx.TxID
			matchReq.TransactionHash = *tx.TransactionHash
			matchReq.BlockNumber = *tx.BlockNumber
			matchReq.BlockHash = *tx.BlockHash
			matchReq.BlockTime = *tx.BlockTime
			matchReq.Status = string(tx.Status.TransactionStatus())

			err, code, _ := s.rpcClient.Post(fmt.Sprintf("http://%s/batch_trade/", conf.Conf.RPCHost), nil, &matchReq, nil)
			if code != 200 || err != nil {
				return fmt.Errorf("updateStatusByUser rpc batch_trade code:%d err:%w ", code, err)
			}

			return nil
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
