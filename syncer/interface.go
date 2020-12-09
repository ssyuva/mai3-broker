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
	ForceRollback       int64 // -1 means not set
	Context             context.Context
	Dao                 dao.DAO
	WatcherW            *model.Watcher
	LatestBlockNumber   int64                        // include
	RollbackBeginHeight int64                        // include
	RollbackEndHeight   int64                        // exclude
	Blocks              map[int64]*model.SyncedBlock // every headers
}

func NewSyncBlockContext() *SyncBlockContext {
	return &SyncBlockContext{
		Blocks: make(map[int64]*model.SyncedBlock),
	}
}

func (c *SyncBlockContext) NeedRollback() bool {
	return c.RollbackBeginHeight < c.RollbackEndHeight
}

func (c *SyncBlockContext) NeedForward() bool {
	return c.RollbackBeginHeight <= c.LatestBlockNumber
}
