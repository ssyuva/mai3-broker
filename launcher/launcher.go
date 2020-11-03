package launcher

import (
	"context"
	// "fmt"
	"github.com/mcarloai/mai-v3-broker/common/chain"
	// "github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/dao"
	logger "github.com/sirupsen/logrus"
	"time"
)

type Launch struct {
	ctx       context.Context
	dao       dao.DAO
	chainCli  chain.ChainClient
	wsChan    chan interface{}
	matchChan chan interface{}
}

func NewLaunch(ctx context.Context, dao dao.DAO, chainCli chain.ChainClient, wsChan, matchChan chan interface{}) *Launch {
	return &Launch{
		ctx:       ctx,
		dao:       dao,
		chainCli:  chainCli,
		wsChan:    wsChan,
		matchChan: matchChan,
	}
}

func (l *Launch) Start() error {
	// start syncer for sync pending transactions
	syncer := NewSyncer(l.ctx, l.dao, l.chainCli)
	go syncer.Run()

	for {
		select {
		case <-l.ctx.Done():
			logger.Infof("Launcher stop")
			return nil
		case <-time.After(5 * time.Second):
			err := l.excuteMatchTransaction()
			if err != nil {
				logger.Errorf("excuteMatchTransaction failed! err:%v", err.Error())
			}
		}
	}
}

func (l *Launch) excuteMatchTransaction() error {
	// transactions, err := l.dao.QueryMatchTransaction("", []model.TransactionStatus{model.TransactionStatusInit})
	// if err != nil {
	// 	return fmt.Errorf("QueryUnconfirmedTransactions failed error:%w", err)
	// }
	// for _, transaction := range transactions {

	// }
	return nil
}
