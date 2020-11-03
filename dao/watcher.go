package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mcarloai/mai-v3-broker/common/model"
)

type WatcherDAO interface {
	CreateWatcher(*model.Watcher) error
	FindWatcherByID(id int) (*model.Watcher, error)
	SaveWatcher(*model.Watcher) error
	SaveBlock(b *model.SyncedBlock) error
	FindBlock(watcherID int, blockNumber int) (*model.SyncedBlock, error)
	FindBlockByHash(watcherID int, blockHash string) (*model.SyncedBlock, error)
	FindBlocksBetween(watcherID int, beginBlockNumber int, endBlockNumber int) ([]*model.SyncedBlock, error)
	RollbackBlock(watcherID int, beginRollbackHeight int, endRollbackHeight int) error
}

type dbWatcher struct {
	model.Watcher
}

func (dbWatcher) TableName() string {
	return "watchers"
}

type dbSyncedBlock struct {
	model.SyncedBlock
}

func (dbSyncedBlock) TableName() string {
	return "synced_blocks"
}

type watcherDAO struct {
	db *gorm.DB
}

func NewWatcherDAO(db *gorm.DB) WatcherDAO {
	return &watcherDAO{db: db}
}

func (l *watcherDAO) CreateWatcher(watcher *model.Watcher) error {
	w := &dbWatcher{Watcher: *watcher}
	if err := l.db.Create(&w).Error; err != nil {
		return fmt.Errorf("CreateWatcher:%w", err)
	}
	return nil
}

func (l *watcherDAO) FindWatcherByID(id int) (*model.Watcher, error) {
	var watcher dbWatcher
	err := l.db.Where("id = ?", id).First(&watcher).Error
	if err != nil {
		return nil, fmt.Errorf("find Watchers by id failed:%w", err)
	}
	return &watcher.Watcher, nil
}

func (l *watcherDAO) SaveWatcher(watcher *model.Watcher) error {
	w := &dbWatcher{Watcher: *watcher}
	if err := l.db.Save(&w).Error; err != nil {
		return fmt.Errorf("UpdateWatcher:%w", err)
	}
	return nil
}

func (l *watcherDAO) SaveBlock(b *model.SyncedBlock) error {
	block := &dbSyncedBlock{SyncedBlock: *b}
	if err := l.db.Save(block).Error; err != nil {
		return fmt.Errorf("save block failed:%w", err)
	}
	return nil
}

func (l *watcherDAO) FindBlock(watcherID int, blockNumber int) (*model.SyncedBlock, error) {
	var b dbSyncedBlock
	err := l.db.Where("watcher_id = ? AND block_number = ?", watcherID, blockNumber).First(&b).Error
	if err != nil {
		return nil, err
	}
	return &b.SyncedBlock, nil
}

func (l *watcherDAO) FindBlockByHash(watcherID int, blockHash string) (*model.SyncedBlock, error) {
	var b dbSyncedBlock
	err := l.db.Where("watcher_id = ? AND block_hash = ?", watcherID, blockHash).First(&b).Error
	if err != nil {
		return nil, err
	}
	return &b.SyncedBlock, nil
}

func (l *watcherDAO) FindBlocksBetween(watcherID int, beginBlockNumber int, endBlockNumber int) ([]*model.SyncedBlock, error) {
	var blocks []*model.SyncedBlock
	err := l.db.Table("synced_blocks").Where("watcher_id = ? AND block_number between ? AND ?", watcherID, beginBlockNumber, endBlockNumber).Find(&blocks).Error
	if err != nil {
		return nil, fmt.Errorf("find synced block failed:%w", err)
	}
	return blocks, nil
}

func (l *watcherDAO) RollbackBlock(watcherID int, beginRollbackHeight int, endRollbackHeight int) error {
	var b dbSyncedBlock
	if err := l.db.Delete(b,
		"watcher_id = ? AND block_number >= ? AND block_number < ?", watcherID, beginRollbackHeight, endRollbackHeight).Error; err != nil {
		return fmt.Errorf("delete block failed:%w", err)
	}
	return nil
}
