package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mcdexio/mai3-broker/common/model"
)

type PerpetualDAO interface {
	CreatePerpetual(*model.Perpetual) error
	GetPerpetualSyncedBlockNumber() (int64, error)
	// Query all the perpetuals
	// if isPublished is true, return published perpetuals only otherwise return all the perpetuals
	QueryPerpetuals(publishedOnly bool) ([]*model.Perpetual, error)
	GetPerpetualByPoolAddressAndIndex(address string, index int64, publishedOnly bool) (*model.Perpetual, error)
	UpdatePerpetual(*model.Perpetual) error
}

type dbPerpetual struct {
	model.Perpetual
}

func (dbPerpetual) TableName() string {
	return "perpetuals"
}

type perpetualDAO struct {
	db *gorm.DB
}

func NewPerpetualDAO(db *gorm.DB) PerpetualDAO {
	return &perpetualDAO{db: db}
}

func (m *perpetualDAO) GetPerpetualByPoolAddressAndIndex(address string, index int64, publishedOnly bool) (*model.Perpetual, error) {
	var perpetual dbPerpetual
	db := m.db
	if publishedOnly {
		db = m.db.Where("is_published = ?", publishedOnly)
	}
	if err := db.Where("liquidity_pool_address = ? AND perpetual_index = ?", address, index).First(&perpetual).Error; err != nil {
		return nil, fmt.Errorf("GetPerpetualByAddress:%w", err)
	}
	return &perpetual.Perpetual, nil
}

func (m *perpetualDAO) QueryPerpetuals(publishedOnly bool) ([]*model.Perpetual, error) {
	result := make([]*model.Perpetual, 0)
	s := make([]*dbPerpetual, 0)
	db := m.db
	if publishedOnly {
		db = m.db.Where("is_published = ?", publishedOnly)
	}
	if err := db.Find(&s).Error; err != nil {
		return nil, fmt.Errorf("QueryPerpetuals:%w", err)
	}
	for _, perpetual := range s {
		result = append(result, &perpetual.Perpetual)
	}
	return result, nil
}

func (m *perpetualDAO) UpdatePerpetual(perpetual *model.Perpetual) error {
	err := m.db.Save(&dbPerpetual{Perpetual: *perpetual}).Error
	if err != nil {
		return fmt.Errorf("UpdatePerpetual:%w", err)
	}
	return err
}

func (m *perpetualDAO) CreatePerpetual(perpetual *model.Perpetual) error {
	err := m.db.Create(&dbPerpetual{Perpetual: *perpetual}).Error
	if err != nil {
		return fmt.Errorf("CreatePerpetual:%w", err)
	}
	return err
}

func (m *perpetualDAO) GetPerpetualSyncedBlockNumber() (int64, error) {
	var perpetual dbPerpetual
	if err := m.db.Order("block_number desc").First(&perpetual).Error; err != nil {
		return 0, err
	}
	return perpetual.BlockNumber, nil
}
