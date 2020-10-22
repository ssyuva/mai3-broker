package watcher

import (
	"context"
	"github.com/mcarloai/mai-v3-broker/dao"
)

type Server struct {
	ctx            context.Context
	factoryAddress string
	wsChan         chan interface{}
	matchChan      chan interface{}
	dao            dao.DAO
}

func New(ctx context.Context, dao dao.DAO, factoryAddress string, wsChan, matchChan chan interface{}) *Server {
	return &Server{
		ctx:            ctx,
		factoryAddress: factoryAddress,
		wsChan:         wsChan,
		matchChan:      matchChan,
		dao:            dao,
	}
}

func (s *Server) Start() error {
	// TODO
	// 1.sync block
	// 2.check match transactions
	// 3.check createPerpetual event, add new perpetual to database, update perpetuals map
	return nil
}
