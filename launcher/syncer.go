package launcher

import (
	"context"
	"github.com/mcarloai/mai-v3-broker/common/chain"
	"github.com/mcarloai/mai-v3-broker/dao"
	logger "github.com/sirupsen/logrus"
	"time"
)

type Syncer struct {
	ctx      context.Context
	dao      dao.DAO
	chainCli chain.ChainClient
}

func NewSyncer(ctx context.Context, dao dao.DAO, chainCli chain.ChainClient) *Syncer {
	return &Syncer{
		ctx:      ctx,
		dao:      dao,
		chainCli: chainCli,
	}
}

func (s *Syncer) Run() {
	for {
		select {
		case <-s.ctx.Done():
			logger.Infof("Syncer stop")
			return
		case <-time.After(5 * time.Second):
			if err := s.syncTransaction(); err != nil {
				logger.Errorf("syncTransaction error:%s", err.Error())
			}
		}
	}
}

func (s *Syncer) syncTransaction() error {
	return nil
}
