package syncer

import (
	"context"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/dao"
)

type BlockSyncer interface {
	Rollback(*SyncBlockContext) error
	Forward(*SyncBlockContext) error
}

type SyncBlockContext struct {
	ForceRollback int // -1 means not set
	Context       context.Context
	// EthProxy            *ethproxy.EthProxy
	Dao                 dao.DAO
	WatcherW            *model.Watcher
	LatestBlockNumber   int                        // include
	RollbackBeginHeight int                        // include
	RollbackEndHeight   int                        // exclude
	Blocks              map[int]*model.SyncedBlock // every headers
	// Signer              *gethTypes.EIP155Signer

}

func NewSyncBlockContext() *SyncBlockContext {
	return &SyncBlockContext{
		Blocks: make(map[int]*model.SyncedBlock),
	}
}

func (c *SyncBlockContext) NeedRollback() bool {
	return c.RollbackBeginHeight < c.RollbackEndHeight
}

func (c *SyncBlockContext) NeedForward() bool {
	return c.RollbackBeginHeight <= c.LatestBlockNumber
}
