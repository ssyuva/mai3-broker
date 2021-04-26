package match

import (
	"context"
	"time"

	"github.com/mcarloai/mai-v3-broker/common/chain"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/conf"
	logger "github.com/sirupsen/logrus"

	"sync"
)

type poolSyncer struct {
	ctx         context.Context
	mu          sync.RWMutex
	pools       []string
	poolStorage map[string]*model.LiquidityPoolStorage
	chainCli    chain.ChainClient
}

func newPoolSyncer(ctx context.Context, cli chain.ChainClient) *poolSyncer {
	return &poolSyncer{
		ctx:      ctx,
		chainCli: cli,
	}
}

func (p *poolSyncer) Run() error {
	for {
		select {
		case <-p.ctx.Done():
			logger.Infof("pool syncer end")
			return nil
		case <-time.After(conf.Conf.PoolSyncerInterval):
			p.runSyncer()
		}
	}
}

func (p *poolSyncer) runSyncer() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	for _, pool := range p.pools {
		poolStorage, err := p.chainCli.GetLiquidityPoolStorage(p.ctx, conf.Conf.ReaderAddress, pool)
		if poolStorage == nil || err != nil {
			logger.Errorf("Pool Syncer: GetLiquidityPoolStorage fail! err:%v", err)
			p.poolStorage[pool] = nil
			continue
		}
		p.poolStorage[pool] = poolStorage
	}
	return nil
}

func (p *poolSyncer) AddPool(pool string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if _, ok := p.poolStorage[pool]; ok {
		return
	}
	p.pools = append(p.pools, pool)
	p.poolStorage[pool] = nil
}

func (p *poolSyncer) GetPoolStorage(pool string) *model.LiquidityPoolStorage {
	p.mu.RLock()
	defer p.mu.RUnlock()
	storage, ok := p.poolStorage[pool]
	if !ok {
		return nil
	}
	return storage
}
