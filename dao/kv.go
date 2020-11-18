package dao

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type KVStore struct {
	Key        string `gorm:"primary_key"`
	Category   string `gorm:"primary_key"`
	Value      []byte
	UpdateTime time.Time
}

type KVStoreDAO interface {
	Get(db *gorm.DB, key, category string) (*KVStore, error)
	Put(db *gorm.DB, value *KVStore) error
	Del(db *gorm.DB, key string) error
	List(db *gorm.DB, category ...string) ([]*KVStore, error)
}

type kvstoreDAO struct {
}

func NewKVStoreDAO() KVStoreDAO {
	return &kvstoreDAO{}
}

func (n *kvstoreDAO) Get(db *gorm.DB, key, category string) (*KVStore, error) {
	kv := new(KVStore)
	if err := db.Where("key = ? and category = ?", key, category).First(kv).Error; err != nil {
		return nil, errors.Wrap(err, "get kvstore entry failed")
	}
	return kv, nil
}

func (n *kvstoreDAO) Put(db *gorm.DB, value *KVStore) error {
	value.UpdateTime = time.Now()
	// return db.Where("key = ? and category = ?", value.Key, value.Category).Assign(value).FirstOrCreate(&KVStore{}).Error
	return db.Save(value).Error
}

func (n *kvstoreDAO) Del(db *gorm.DB, key string) error {
	return db.Delete(&KVStore{}, "key = ?", key).Error
}

func (n *kvstoreDAO) List(db *gorm.DB, category ...string) ([]*KVStore, error) {
	var kvs []*KVStore
	if err := db.Where("category in (?)", category).Find(&kvs).Error; err != nil {
		return nil, errors.Wrap(err, "list kvs failed")
	}
	return kvs, nil
}
