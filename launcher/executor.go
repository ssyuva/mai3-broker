package launcher

import (
	"context"
	"github.com/mcarloai/mai-v3-broker/common/chain"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/pkg/errors"
	logger "github.com/sirupsen/logrus"
	"time"
)

type Executor struct {
	ctx      context.Context
	dao      dao.DAO
	chainCli chain.ChainClient
	execChan chan interface{}
}

func NewExecutor(ctx context.Context, dao dao.DAO, chainCli chain.ChainClient, execChan chan interface{}) *Executor {
	return &Executor{
		ctx:      ctx,
		dao:      dao,
		chainCli: chainCli,
		execChan: execChan,
	}
}

func (s *Executor) Run() {
	for {
		select {
		case <-s.ctx.Done():
			logger.Infof("Executor stop")
			return
		case <-s.execChan:
		case <-time.After(5 * time.Second):
		}
		if err := s.executeTransaction(); err != nil {
			logger.Errorf("executeTransaction error:%s", err.Error())
		}
	}
}

func (s *Executor) executeTransaction() error {
	ctx, done := context.WithTimeout(s.ctx, conf.Conf.BlockChain.Timeout.Duration)
	defer done()
	users, err := s.dao.GetUsersWithStatus(model.TxInitial)
	if err != nil {
		logger.Warnf("find user with status failed: %s", err)
		return errors.Wrap(err, "fail to get users has initial transactions")
	}
	for _, user := range users {
		tx, err := s.dao.FirstTxByUser(user, model.TxInitial)
		if err != nil {
			logger.Warningf("retreive new transaction failed for user %v", user)
			continue
		}
		logger.Infof("found initial transaction %s from user %s", tx.TxID, user)
		err = s.sendTransaction(ctx, tx)
		if err != nil {
			logger.Warningf("commit new transaction failed for %v: %v", user, err)
		}
	}
	return nil
}

func (s *Executor) sendTransaction(ctx context.Context, tx *model.LaunchTransaction) error {
	logger.Infof("send transaction for user %v", tx.FromAddress)
	expNonce, err := s.chainCli.PendingNonceAt(ctx, tx.FromAddress)
	if err != nil {
		return errors.Wrap(err, "access node rpc failed")
	}
	// has nonce, but external transaction
	if tx.Nonce != nil && expNonce > *tx.Nonce {
		logger.Warnf("found external transactions, try reset current transaction (CANCEL)")
		if err := s.reset(ctx, tx); err != nil {
			return errors.Wrap(err, "reset transaction failed")
		}
		return nil
	}
	minNonce, ok := s.dao.GetNextNonce(tx.FromAddress)
	if (tx.Nonce != nil && expNonce < *tx.Nonce) || (tx.Nonce == nil && ok && expNonce < minNonce) {
		var targetNonce uint64
		if tx.Nonce != nil {
			targetNonce = *tx.Nonce
		} else {
			targetNonce = minNonce
		}
		logger.Warnf("found expect nonce lower than assigned, try fast-forward [%v, %v]", expNonce, targetNonce)
		if err = s.fastForward(ctx, tx.FromAddress, expNonce, targetNonce); err != nil {
			logger.Warnf("fast-forward failed, will try to send transation: %s", err)
		}
		logger.Infof("try fast-forward done")
		return nil
	}
	// allocate nonce
	if err := s.dao.Transaction(func(dao dao.DAO) error {
		tx.Nonce = model.Uint64(expNonce)
		if err := dao.UpdateTx(tx); err != nil {
			return errors.Wrap(err, "save nonce failed")
		}
		if err := dao.UpdateNonce(tx.FromAddress, model.Uint64(*tx.Nonce+1)); err != nil {
			return errors.Wrap(err, "update nonce failed")
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "allocate nonce failed")
	}
	logger.Infof("using nonce %v", *tx.Nonce)
	// send
	if err := s.send(ctx, tx); err != nil {
		return errors.Wrap(err, "send transaction failed")
	}
	return nil
}

func (s *Executor) reset(ctx context.Context, tx *model.LaunchTransaction) error {
	if tx.TransactionHash == nil {
		tx.Status = model.TxCanceled
	} else {
		_, err := s.chainCli.TransactionByHash(ctx, *tx.TransactionHash)
		if err != nil {
			tx.Status = model.TxCanceled
		} else {
			tx.Status = model.TxPending
		}
	}
	if err := s.dao.UpdateTx(tx); err != nil {
		return errors.Wrap(err, "update transaction failed")
	}
	return nil
}

func (s *Executor) fastForward(ctx context.Context, addr string, start uint64, end uint64) error {
	for i := start; i < end; i++ {
		logger.Infof("resend transaction by nonce %v", i)
		prevs, err := s.dao.GetTxsByNonce(addr, model.Uint64(i))
		if err != nil {
			return errors.Wrap(err, "find previous transaction failed")
		}
		if len(prevs) == 0 {
			continue
		}
		for _, tx := range prevs {
			if err := s.send(ctx, tx); err != nil {
				logger.Warnf("(may no be a issue) try resending previous transaction %v failed, err %v", tx.TxID, err)
			}
		}
	}
	return nil
}

// prepare transaction price and gas limit
func (s *Executor) prepare(ctx context.Context, tx *model.LaunchTransaction) error {
	if tx.Nonce == nil {
		return errors.New("missing nonce")
	}

	limit := uint64(5000000)
	tx.GasLimit = &limit
	price := uint64(20) * 1e9
	tx.GasPrice = &price

	//TODO gasprice gas limit
	// if tx.GasLimit == nil || *tx.GasLimit == 0 {
	// 	l := s.limiter.GasLimit(ctx, tx)
	// 	tx.GasLimit = &l
	// }
	// if tx.GasPrice == nil || *tx.GasPrice == 0 {
	// 	if t.pricer == nil {
	// 		return errors.New("gas pricer not ready")
	// 	}
	// 	p := s.pricer.GasPrice(tx.Type)
	// 	tx.GasPrice = &p
	// }

	return nil
}

func (s *Executor) send(ctx context.Context, tx *model.LaunchTransaction) error {
	return s.dao.Transaction(func(dao dao.DAO) error {
		prevHash := tx.TransactionHash
		err := s.prepare(ctx, tx)
		if err != nil {
			return errors.Wrap(err, "prepare transaction failed")
		}
		currHash, err := s.chainCli.SendTransaction(ctx, tx)
		if err != nil {
			return errors.Wrap(err, "send transaction failed")
		}
		tx.Status = model.TxPending
		if prevHash != nil && model.MustString(prevHash) != currHash {
			logger.Infof("resend tx, former tx: %s => %s", model.MustString(prevHash), currHash)
			tx.ID = 0
			if err := dao.CreateTx(tx); err != nil {
				return errors.Wrap(err, "create resend transaction failed")
			}
		} else {
			logger.Infof("update tx, tx: %s", currHash)
			if err := dao.UpdateTx(tx); err != nil {
				return errors.Wrap(err, "update transaction failed")
			}
		}
		return nil
	})
}
