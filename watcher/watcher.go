package watcher

import (
	"context"
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/chain"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/mcarloai/mai-v3-broker/syncer"
	logger "github.com/sirupsen/logrus"
	"math/big"
	"time"
)

// if the latestBlock.blockNumber - block.blockNumber > MatureBlocks, we tend to believe the block will never change
const MatureBlocks = 1000

type Watcher struct {
	ctx            context.Context
	factoryAddress string
	wsChan         chan interface{}
	matchChan      chan interface{}
	chainCli       chain.ChainClient
	dao            dao.DAO
	blockSyncers   []syncer.BlockSyncer
}

func New(ctx context.Context, cli chain.ChainClient, dao dao.DAO, factoryAddress string, wsChan, matchChan chan interface{}) *Watcher {
	watcher := &Watcher{
		ctx:            ctx,
		factoryAddress: factoryAddress,
		wsChan:         wsChan,
		matchChan:      matchChan,
		chainCli:       cli,
		dao:            dao,
		blockSyncers:   make([]syncer.BlockSyncer, 0),
	}

	watcher.AddBlockSyncer(syncer.NewCreatePerpetualSyncer(factoryAddress, cli, matchChan))
	watcher.AddBlockSyncer(syncer.NewPerpetualMatchSyncer(cli, matchChan))

	return watcher
}

func (w *Watcher) AddBlockSyncer(syncer syncer.BlockSyncer) {
	w.blockSyncers = append(w.blockSyncers, syncer)
}

func (w *Watcher) Start() error {
	logger.Infof("Watcher start")
	if err := w.EnsureWatcherRecord(); err != nil {
		return fmt.Errorf("init watcher db failed:%w", err)
	}

	for {
		select {
		case <-w.ctx.Done():
			logger.Info("Watcher receive conext done")
			return nil
		case <-time.After(conf.Conf.BlockChain.Interval.Duration):
			if err := w.Sync(); err != nil {
				logger.Warnf("sync failed:%v", err.Error())
			}
		}
	}
}

// EnsureWatcherRecord ensure that there is a record in the "watchers" table
func (w *Watcher) EnsureWatcherRecord() error {
	return w.dao.Transaction(func(wDao dao.DAO) error {
		wDao.ForUpdate()
		mW, err := wDao.FindWatcherByID(conf.Conf.WatcherID)
		if err == nil {
			logger.Infof("initDB:start[%d] synced[%d]", mW.InitialBlockNumber, mW.SyncedBlockNumber)
			return nil
		}
		logger.Infof("initDB:this is the 1st time you run the watcher")

		ctxTimeout, ctxTimeoutCancel := context.WithTimeout(w.ctx, conf.Conf.BlockChain.Timeout.Duration)
		defer ctxTimeoutCancel()
		header, err := w.chainCli.HeaderByNumber(ctxTimeout, nil)
		if err != nil {
			return fmt.Errorf("initDB:get latest block failed:%w", err)
		}

		watcher := &model.Watcher{
			ID:                 conf.Conf.WatcherID,
			InitialBlockNumber: header.BlockNumber,
			SyncedBlockNumber:  header.BlockNumber,
		}
		if err = wDao.SaveWatcher(watcher); err != nil {
			return fmt.Errorf("initDB:%w", err)
		}

		return nil
	})
}

// sync data between db and blockchain
func (w *Watcher) Sync() error {
	syncedCount := 0
	begin := time.Now()
	syncedCount, err := w.syncFrom(-1)
	dur := time.Since(begin)
	if err != nil {
		logger.Warnf("sync error in %v: %s", dur, err.Error())
		return err
	} else if syncedCount > 0 {
		logger.Infof("sync %d block(s) in %v", syncedCount, dur)
	}
	return nil
}

func (w *Watcher) syncFrom(forceRollback int64) (syncedCount int, err error) {
	// parameters
	syncCtx := syncer.NewSyncBlockContext()
	syncCtx.ForceRollback = forceRollback
	syncCtx.Context = w.ctx

	// the new end
	ctxTimeout, ctxTimeoutCancel := context.WithTimeout(w.ctx, conf.Conf.BlockChain.Timeout.Duration)
	defer ctxTimeoutCancel()
	header, err := w.chainCli.HeaderByNumber(ctxTimeout, nil)
	if err != nil {
		return 0, fmt.Errorf("get latest block failed:%w", err)
	}

	// syncing range
	syncCtx.LatestBlockNumber = header.BlockNumber
	w.saveBlockHeader(syncCtx, header)

	// sync begins
	err = w.dao.Transaction(func(wDao dao.DAO) error {
		wDao.ForUpdate()
		return w.innerSync(syncCtx)
	})
	if err != nil {
		err = fmt.Errorf("innerSync failed:%w", err)
		return
	}

	// end of syncing
	syncedCount = len(syncCtx.Blocks) - 1 /* ignore the latest block. it's always in this array */
	return
}

func (w *Watcher) saveBlockHeader(syncCtx *syncer.SyncBlockContext, header *model.BlockHeader) {
	syncCtx.Blocks[header.BlockNumber] = &model.SyncedBlock{
		WatcherID:   conf.Conf.WatcherID,
		BlockNumber: header.BlockNumber,
		BlockHash:   header.BlockHash,
		ParentHash:  header.ParentHash,
		BlockTime:   header.BlockTime,
	}
}

func (w *Watcher) innerSync(syncCtx *syncer.SyncBlockContext) (err error) {
	// the previous end
	syncCtx.WatcherW, err = w.dao.FindWatcherByID(conf.Conf.WatcherID)
	if err != nil {
		return fmt.Errorf("get watcher failed:%w", err)
	}
	if syncCtx.ForceRollback >= 0 {
		logger.Warnf("force rollback to [%d], old synced[%d], old initial[%d]",
			syncCtx.ForceRollback, syncCtx.WatcherW.SyncedBlockNumber, syncCtx.WatcherW.InitialBlockNumber)
		if syncCtx.WatcherW.InitialBlockNumber >= syncCtx.ForceRollback {
			syncCtx.WatcherW.InitialBlockNumber = syncCtx.ForceRollback - 1
		}
	} else {
		if syncCtx.WatcherW.SyncedBlockNumber > syncCtx.LatestBlockNumber {
			return fmt.Errorf("block[%d] is greater than latest on chain[%d]",
				syncCtx.WatcherW.SyncedBlockNumber, syncCtx.LatestBlockNumber)
		}
	}

	// limit the syncing range
	var partialSync int64
	partialSync = -1
	if syncCtx.ForceRollback >= 0 {
		if syncCtx.LatestBlockNumber > syncCtx.ForceRollback+MatureBlocks {
			partialSync = syncCtx.ForceRollback + int64(MatureBlocks)
			logger.Warnf("the sync range is too wide, let us partially sync [force %v, %v]",
				syncCtx.ForceRollback, partialSync)
		}
	} else {
		if syncCtx.LatestBlockNumber > syncCtx.WatcherW.SyncedBlockNumber+MatureBlocks {
			partialSync = syncCtx.WatcherW.SyncedBlockNumber + int64(MatureBlocks)
			logger.Warnf("the sync range is too wide, let us partially sync [auto %v, %v]",
				syncCtx.WatcherW.SyncedBlockNumber, partialSync)
		}
	}
	if partialSync >= 0 {
		bigPartialSync := big.NewInt(int64(partialSync))
		ctxTimeout, ctxTimeoutCancel := context.WithTimeout(w.ctx, conf.Conf.BlockChain.Timeout.Duration)
		defer ctxTimeoutCancel()
		header, err := w.chainCli.HeaderByNumber(ctxTimeout, bigPartialSync)
		if err != nil {
			return fmt.Errorf("get block by number(%v) failed:%w", partialSync, err)
		}
		syncCtx.LatestBlockNumber = header.BlockNumber
		w.saveBlockHeader(syncCtx, header)
	}

	// find our common ancestor
	err = w.innerTraceAncestor(syncCtx)
	if err != nil {
		return fmt.Errorf("innerTraceAncestor:%w", err)
	}

	// rollback, forward
	err = w.innerRollback(syncCtx)
	if err != nil {
		return fmt.Errorf("innerRollback [%v, %v) faiked:%w", syncCtx.RollbackBeginHeight, syncCtx.RollbackEndHeight, err)
	}
	err = w.innerForward(syncCtx)
	if err != nil {
		return fmt.Errorf("innerForward [%v, %v] failed:%w", syncCtx.RollbackBeginHeight, syncCtx.LatestBlockNumber, err)
	}

	// update the watcher
	if err = w.dao.SaveWatcher(syncCtx.WatcherW); err != nil {
		return fmt.Errorf("update watch db failed:%w", err)
	}

	return nil
}

func (w *Watcher) innerTraceAncestor(syncCtx *syncer.SyncBlockContext) (err error) {
	var i int64
	for i = syncCtx.LatestBlockNumber; i > syncCtx.WatcherW.InitialBlockNumber; i-- {
		if i%1000 == 0 {
			logger.Infof("checking the latest block[%v]", i)
		}
		var blockHash string
		if i == syncCtx.LatestBlockNumber {
			// the latest block has already been saved
			if latestBlock, ok := syncCtx.Blocks[i]; !ok {
				return fmt.Errorf("[bug]missing block[%d]", i)
			} else {
				blockHash = latestBlock.BlockHash
			}
		} else {
			// trace by parent hash
			if nextBlock, ok := syncCtx.Blocks[i+1]; !ok {
				return fmt.Errorf("[bug]missing block[%d]", i+1)
			} else {
				blockHash = nextBlock.ParentHash
			}

			// read the block chain
			ctxTimeout, ctxTimeoutCancel := context.WithTimeout(w.ctx, conf.Conf.BlockChain.Timeout.Duration)
			defer ctxTimeoutCancel()
			block, err := w.chainCli.HeaderByHash(ctxTimeout, blockHash)
			if err != nil {
				return fmt.Errorf("get block failed:%w", err)
			} else if block.BlockHash != blockHash {
				return fmt.Errorf("[bug]block hash mismatched:%s vs %s", block.BlockHash, blockHash)
			}
			w.saveBlockHeader(syncCtx, block)
		}

		// [common parent ... synced] is in the DB. (synced ... new end) is not in the DB
		if i <= syncCtx.WatcherW.SyncedBlockNumber {
			if syncCtx.ForceRollback >= 0 {
				if i < syncCtx.ForceRollback {
					break
				}
				// force rollback block i
			} else {
				// check the db
				var dbBlock *model.SyncedBlock
				dbBlock, err = w.dao.FindBlock(conf.Conf.WatcherID, i)
				if err != nil {
					return fmt.Errorf("find synced block failed:%w", err)
				}
				if blockHash == dbBlock.BlockHash {
					break
				}
				// rollback block i, dbBlock
			}
		}
	}
	syncCtx.RollbackBeginHeight = i + 1
	syncCtx.RollbackEndHeight = syncCtx.WatcherW.SyncedBlockNumber + 1

	return nil
}

func (w *Watcher) innerRollback(syncCtx *syncer.SyncBlockContext) (err error) {
	if !syncCtx.NeedRollback() {
		return nil
	}
	logger.Infof("innerRollback:rollback between [%v, %v)", syncCtx.RollbackBeginHeight, syncCtx.RollbackEndHeight)

	// sub-syncers
	for _, syncer := range w.blockSyncers {
		err = syncer.Rollback(syncCtx)
		if err != nil {
			return fmt.Errorf("sub-syncers failed:%w", err)
		}
	}

	// final remove
	err = w.dao.RollbackBlock(conf.Conf.WatcherID, syncCtx.RollbackBeginHeight, syncCtx.RollbackEndHeight)
	if err != nil {
		return fmt.Errorf("remove blocks failed:%w", err)
	}
	return nil
}

func (w *Watcher) innerForward(syncCtx *syncer.SyncBlockContext) (err error) {
	if syncCtx.NeedForward() {
		logger.Infof("innerForward:forwarding between [%v, %v]", syncCtx.RollbackBeginHeight, syncCtx.LatestBlockNumber)
	} else {
		// up-to-date
		return nil
	}

	// update blocks
	for i := syncCtx.RollbackBeginHeight; i <= syncCtx.LatestBlockNumber; i++ {
		block, ok := syncCtx.Blocks[i]
		if !ok {
			return fmt.Errorf("[bug]missing block[%d]", i)
		}
		if i%1000 == 0 {
			logger.Infof("sync block[%d][%s]", i, block.BlockHash)
		}
		if err = w.dao.SaveBlock(block); err != nil {
			return fmt.Errorf("save block failed:%w", err)
		}
	}
	syncCtx.WatcherW.SyncedBlockNumber = syncCtx.LatestBlockNumber
	if syncCtx.WatcherW.SyncedBlockNumber <= syncCtx.WatcherW.InitialBlockNumber {
		return fmt.Errorf("bad sync point:%d <= %d", syncCtx.WatcherW.SyncedBlockNumber, syncCtx.WatcherW.InitialBlockNumber)
	}

	// sub-syncers
	for _, syncer := range w.blockSyncers {
		err = syncer.Forward(syncCtx)
		if err != nil {
			return fmt.Errorf("sub-syncers failed:%w", err)
		}
	}

	return nil
}
