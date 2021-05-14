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

const POOL_STORAGE_USE_DURATION = 60

type poolInfo struct {
	storage   *model.LiquidityPoolStorage
	latestGet time.Time
	isDirty   bool
}

type poolSyncer struct {
	ctx         context.Context
	mu          sync.RWMutex
	pools       []string
	poolStorage map[string]*poolInfo
	chainCli    chain.ChainClient
}

func newPoolSyncer(ctx context.Context, cli chain.ChainClient) *poolSyncer {
	return &poolSyncer{
		ctx:         ctx,
		chainCli:    cli,
		pools:       make([]string, 0),
		poolStorage: make(map[string]*poolInfo),
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
	for _, pool := range p.pools {
		p.syncPool(pool)
	}
	return nil
}

func (p *poolSyncer) syncPool(pool string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	info, ok := p.poolStorage[pool]
	// pool storage not be used in duration, do not refresh it until someone used it again
	if ok && time.Since(info.latestGet).Seconds() > POOL_STORAGE_USE_DURATION {
		if !info.isDirty {
			info.isDirty = true
		}
		return
	}

	poolStorage, err := p.chainCli.GetLiquidityPoolStorage(p.ctx, conf.Conf.ReaderAddress, pool)
	if poolStorage == nil || err != nil {
		logger.Errorf("Pool Syncer: GetLiquidityPoolStorage fail! pool:%s err:%v", pool, err)
		p.poolStorage[pool] = nil
	}
	info.storage = poolStorage
	if info.isDirty {
		info.isDirty = false
	}
}

func (p *poolSyncer) AddPool(pool string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if _, ok := p.poolStorage[pool]; ok {
		return
	}
	p.pools = append(p.pools, pool)
	poolStorage, err := p.chainCli.GetLiquidityPoolStorage(p.ctx, conf.Conf.ReaderAddress, pool)
	if poolStorage == nil || err != nil {
		logger.Errorf("Pool Syncer: GetLiquidityPoolStorage fail! pool:%s err:%v", pool, err)
		return
	}
	p.poolStorage[pool] = &poolInfo{
		latestGet: time.Now(),
		isDirty:   false,
		storage:   poolStorage,
	}
}

func (p *poolSyncer) GetPoolStorage(pool string) *model.LiquidityPoolStorage {
	p.mu.RLock()
	defer p.mu.RUnlock()
	info, ok := p.poolStorage[pool]
	if info == nil && !ok {
		return nil
	}
	if info.isDirty {
		// update latestGet, syncer will update PoolStorage in next round
		info.latestGet = time.Now()

		// get from chain
		poolStorage, err := p.chainCli.GetLiquidityPoolStorage(p.ctx, conf.Conf.ReaderAddress, pool)
		if poolStorage == nil || err != nil {
			return nil
		}
		return poolStorage
	}
	return info.storage
}
