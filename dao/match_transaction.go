package dao

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mcarloai/mai-v3-broker/common/model"
)

type MatchTransactionDAO interface {
	CreateMatchTransaction(*model.MatchTransaction) error
	QueryMatchTransaction(marketID string, status []model.TransactionStatus) ([]*model.MatchTransaction, error)
	QueryUnconfirmedTransactions() ([]*model.MatchTransaction, error)
	QueryUnconfirmedTransactionsByContract(address string) (transactions []*model.MatchTransaction, err error)
	GetMatchTransaction(ID string, forUpdate bool) (*model.MatchTransaction, error)
	UpdateMatchTransaction(transaction *model.MatchTransaction) error
	RollbackTransactions(beginRollbackHeight int, endRollbackHeight int) (transactions []*model.MatchItem, err error)
}

type matchTransactionDAO struct {
	db *gorm.DB
}

func (t *matchTransactionDAO) CreateMatchTransaction(transaction *model.MatchTransaction) error {
	jsonData, err := json.Marshal(transaction.MatchItems)
	if err != nil {
		return fmt.Errorf("CreateMatchTransaction:%w", err)
	}
	transaction.MatchJson = string(jsonData)
	transaction.CreatedAt = time.Now().UTC()
	if err := t.db.Create(transaction).Error; err != nil {
		return fmt.Errorf("CreateMatchTransaction:%w", err)
	}
	return nil
}

func (t *matchTransactionDAO) GetMatchTransaction(ID string, forUpdate bool) (*model.MatchTransaction, error) {
	var transaction model.MatchTransaction

	db := t.db
	if forUpdate {
		db = db.Set("gorm:query_option", "FOR UPDATE")
	}
	if err := db.Where("id = ?", ID).First(&transaction).Error; err != nil {
		return nil, fmt.Errorf("GetMatchTransaction:%w", err)
	}

	if err := json.Unmarshal([]byte(transaction.MatchJson), &transaction.MatchItems); err != nil {
		return nil, fmt.Errorf("GetMatchTransaction:%w", err)
	}

	return &transaction, nil
}

func (t *matchTransactionDAO) QueryMatchTransaction(address string, status []model.TransactionStatus) (transactions []*model.MatchTransaction, err error) {
	db := t.db
	if address != "" {
		db = db.Where("perpetual_address = ?", address)
	}
	if len(status) != 0 {
		db = db.Where("status in (?)", status)
	}
	if err = db.Order("created_at").Find(&transactions).Error; err != nil {
		err = fmt.Errorf("QueryMatchTransaction:%w", err)
		return
	}

	for _, transaction := range transactions {
		if err := json.Unmarshal([]byte(transaction.MatchJson), &transaction.MatchItems); err != nil {
			return nil, fmt.Errorf("QueryMatchTransaction:%w", err)
		}
	}
	return
}

func (t *matchTransactionDAO) UpdateMatchTransaction(transaction *model.MatchTransaction) error {
	jsonData, err := json.Marshal(transaction.MatchItems)
	if err != nil {
		return fmt.Errorf("CreateMatchTransaction:%w", err)
	}
	transaction.MatchJson = string(jsonData)

	if err = t.db.Save(transaction).Error; err != nil {
		return fmt.Errorf("UpdateMatchTransaction:%w", err)
	}
	return nil
}

func (t *matchTransactionDAO) QueryUnconfirmedTransactions() (transactions []*model.MatchTransaction, err error) {
	if err = t.db.Where("block_confirmed = ?", false).Find(&transactions).Error; err != nil {
		err = fmt.Errorf("QueryUnstableTransactions:%w", err)
		return
	}
	return
}

func (t *matchTransactionDAO) QueryUnconfirmedTransactionsByContract(address string) (transactions []*model.MatchTransaction, err error) {
	if err = t.db.Where("perpetual_address = ?", address).Where("block_confirmed = ?", false).Find(&transactions).Error; err != nil {
		err = fmt.Errorf("QueryUnstableTransactions:%w", err)
		return
	}
	return
}

func (t *matchTransactionDAO) RollbackTransactions(beginRollbackHeight int, endRollbackHeight int) (items []*model.MatchItem, err error) {
	transactions := make([]*model.MatchTransaction, 0)
	if err = t.db.Where("block_confirmed = ?", true).Where("block_number >= ? AND block_number < ?", beginRollbackHeight, endRollbackHeight).Find(&transactions).Error; err != nil {
		err = fmt.Errorf("QueryTransactions:%w", err)
		return
	}
	for _, transaction := range transactions {
		if err := json.Unmarshal([]byte(transaction.MatchJson), &transaction.MatchItems); err != nil {
			return items, fmt.Errorf("QueryMatchTransaction:%w", err)
		}

		transaction.Status = model.TransactionStatusPending
		if err = t.db.Save(transaction).Error; err != nil {
			return items, fmt.Errorf("UpdateMatchTransaction status:%w", err)
		}
		items = append(items, transaction.MatchItems...)
	}
	return
}
