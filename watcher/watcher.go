package watcher

import (
	"context"
	"github.com/mcarloai/mai-v3-broker/dao"
	"sync"
)

type Server struct {
	ctx            context.Context
	factoryAddress string
	perpetuals     *sync.Map
	dao            dao.DAO
}

func New(ctx context.Context, factoryAddress string, perpetuals *sync.Map) *Server {
	return &Server{
		ctx:            ctx,
		factoryAddress: factoryAddress,
		perpetuals:     perpetuals,
		dao:            dao.New(),
	}
}

func Start() error {
	// sync block
	// check match transactions
	// check createPerpetual event, add new perpetual to database, update perpetuals map
	return nil
}
