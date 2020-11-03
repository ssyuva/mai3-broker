package model

import (
	"time"
)

type Watcher struct {
	ID                 int `json:"id"  db:"id"   primaryKey:"true" gorm:"primary_key"`
	InitialBlockNumber int `json:"initialBlockNumber" db:"initial_block_number"`
	SyncedBlockNumber  int `json:"syncedBlockNumber"  db:"synced_block_number"`
}

type SyncedBlock struct {
	WatcherID   int       `json:"watcherID"   db:"watcher_id"   primaryKey:"true" gorm:"primary_key"`
	BlockNumber int       `json:"blockNumber" db:"block_number" primaryKey:"true" gorm:"primary_key"`
	BlockHash   string    `json:"blockHash"   db:"block_hash"`
	ParentHash  string    `json:"parentHash"  db:"parent_hash"`
	BlockTime   time.Time `json:"blockTime"   db:"block_time"`
}
