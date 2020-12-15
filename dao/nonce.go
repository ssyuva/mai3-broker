package dao

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Nonce struct {
	Address    string `gorm:"primary_key"`
	Nonce      *uint64
	UpdateTime time.Time
}

func (Nonce) TableName() string {
	return "broker_nonces"
}

type NonceDAO interface {
	GetNextNonce(addr string) (uint64, bool)
	UpdateNonce(addr string, nonce *uint64) error
}

type nonceDAO struct {
	db *gorm.DB
}

func NewNonceDAO(db *gorm.DB) NonceDAO {
	return &nonceDAO{db: db}
}

func (n *nonceDAO) GetNextNonce(addr string) (uint64, bool) {
	nonce := new(Nonce)
	if err := n.db.Where("address = ? ", addr).Find(&nonce).Error; err != nil {
		return 0, false
	}
	if nonce.Nonce == nil {
		return 0, false
	}
	return *nonce.Nonce, true
}

func (n *nonceDAO) UpdateNonce(addr string, nonce *uint64) error {
	toUpdate := &Nonce{
		Address:    addr,
		Nonce:      nonce,
		UpdateTime: time.Now(),
	}
	return n.db.Save(toUpdate).Error
}
