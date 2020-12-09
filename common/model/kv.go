package model

import "time"

type KVStore struct {
	Key        string `gorm:"primary_key"`
	Category   string `gorm:"primary_key"`
	Value      []byte
	UpdateTime time.Time
}
