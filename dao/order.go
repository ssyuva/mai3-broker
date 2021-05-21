package dao

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mcdexio/mai3-broker/common/model"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
)

type OrderDAO interface {
	CreateOrder(order *model.Order) error
	GetOrder(orderHash string) (*model.Order, error)
	GetPendingOrderUsers(poolAddress string, perpIndex int64, status []model.OrderStatus) ([]string, error)
	QueryOrder(traderAddress string, poolAddress string, perpIndex int64, status []model.OrderStatus, beforeOrderID, afterOrderID int64, limit int) ([]*model.Order, error)
	QueryOrderWithCreateTime(traderAddress string, poolAddress string, perpIndex int64,
		status []model.OrderStatus, beforeOrderID, afterOrderID int64, beginTime, endTime time.Time, limit int) ([]*model.Order, error)
	GetOrderByHashs(hashs []string) ([]*model.Order, error)
	UpdateOrder(order *model.Order) error
	LoadMatchOrders(matchItems []*model.MatchItem) error
}
type dbOrder struct {
	model.Order
	CancelsJson string `json:"-" db:"cancels_json"`
}

func (dbOrder) TableName() string {
	return "orders"
}

func (o *dbOrder) setStatusByAmounts() {
	if o.ConfirmedAmount.Equal(o.Amount) {
		o.Status = model.OrderFullFilled
	} else if o.CanceledAmount.Equal(o.Amount) {
		o.Status = model.OrderCanceled
	} else if o.AvailableAmount.Add(o.PendingAmount).Abs().GreaterThan(decimal.Zero) {
		o.Status = model.OrderPending
	} else {
		o.Status = model.OrderPartialFilled
	}
}

func (o *dbOrder) unmarshalCancelReason() error {
	return json.Unmarshal([]byte(o.CancelsJson), &o.CancelReasons)
}

func (o *dbOrder) marshalCancelReason() error {
	sort.Slice(o.CancelReasons, func(i, j int) bool {
		return o.CancelReasons[i].CanceledAt.Before(o.CancelReasons[j].CanceledAt)
	})

	json, err := json.Marshal(o.CancelReasons)
	if err != nil {
		return fmt.Errorf("marshalCancelReason:%w", err)
	}
	o.CancelsJson = string(json)
	return nil
}

type orderDAO struct {
	db *gorm.DB
}

func NewOrderDAO(db *gorm.DB) OrderDAO {
	return &orderDAO{db: db}
}

func (o *orderDAO) GetOrder(orderHash string) (*model.Order, error) {
	var order dbOrder
	if err := o.db.Where("order_hash = ?", orderHash).First(&order).Error; err != nil {
		return nil, fmt.Errorf("GetOrder:%w", err)
	}
	if err := order.unmarshalCancelReason(); err != nil {
		logger.Warnf("load order cancel reason error:%v", err)
	}
	return &order.Order, nil
}

func (o *orderDAO) GetPendingOrderUsers(poolAddress string, perpIndex int64, status []model.OrderStatus) ([]string, error) {

	var (
		orders []*dbOrder
		users  []string
	)
	where := o.db.Table("orders")

	if poolAddress != "" {
		where = where.Where("liquidity_pool_address = ? AND perpetual_index = ?", poolAddress, perpIndex)
	}

	if len(status) > 0 {
		where = where.Where("status in (?)", status)
	}
	err := where.Select("DISTINCT(trader_address)").Find(&orders).Error
	if err != nil {
		return nil, fmt.Errorf("fail to get pending users: %w", err)
	}
	for _, order := range orders {
		users = append(users, order.TraderAddress)
	}
	return users, nil
}

func (o *orderDAO) QueryOrder(traderAddress string, poolAddress string, perpIndex int64, status []model.OrderStatus, beforeOrderID, afterOrderID int64, limit int) (orders []*model.Order, err error) {
	var dbOrders []*dbOrder
	where := o.db.Table("orders")

	if traderAddress != "" {
		where = where.Where("trader_address=?", traderAddress)
	}

	if poolAddress != "" {
		where = where.Where("liquidity_pool_address = ? AND perpetual_index = ?", poolAddress, perpIndex)
	}

	if len(status) > 0 {
		where = where.Where("status in (?)", status)
	}

	if beforeOrderID > 0 {
		where = where.Where("id < ?", beforeOrderID)
	}

	if afterOrderID > 0 {
		where = where.Where("id > ?", afterOrderID)
	}

	if limit > 0 {
		where = where.Order("id desc")
		where = where.Limit(limit)
	}

	if err = where.Find(&dbOrders).Error; err != nil {
		err = fmt.Errorf("QueryOrder:%w", err)
		return
	}

	for _, o := range dbOrders {
		if loadErr := o.unmarshalCancelReason(); loadErr != nil {
			logger.Warnf("load order cancel reason error:%v", loadErr)
		}
		orders = append(orders, &o.Order)
	}
	return
}

func (o *orderDAO) QueryOrderWithCreateTime(traderAddress string, poolAddress string, perpIndex int64,
	status []model.OrderStatus, beforeOrderID, afterOrderID int64, beginTime, endTime time.Time, limit int) (orders []*model.Order, err error) {
	var dbOrders []*dbOrder
	where := o.db.Table("orders")

	if traderAddress != "" {
		where = where.Where("trader_address=?", traderAddress)
	}

	if poolAddress != "" {
		where = where.Where("liquidity_pool_address = ? AND perpetual_index = ?", poolAddress, perpIndex)
	}

	if len(status) > 0 {
		where = where.Where("status in (?)", status)
	}

	if beforeOrderID > 0 {
		where = where.Where("id < ?", beforeOrderID)
	}

	if afterOrderID > 0 {
		where = where.Where("id > ?", afterOrderID)
	}

	if !beginTime.IsZero() {
		where = where.Where("created_at > ?", beginTime)
	}
	if !endTime.IsZero() {
		where = where.Where("created_at < ?", endTime)
	}

	if limit > 0 {
		where = where.Order("id desc")
		where = where.Limit(limit)
	}

	if err = where.Find(&dbOrders).Error; err != nil {
		err = fmt.Errorf("QueryOrder:%w", err)
		return
	}

	for _, o := range dbOrders {
		if loadErr := o.unmarshalCancelReason(); loadErr != nil {
			logger.Warnf("load order cancel reason error:%v", loadErr)
		}
		orders = append(orders, &o.Order)
	}
	return
}

func (o *orderDAO) GetOrderByHashs(hashs []string) (orders []*model.Order, err error) {
	var dbOrders []*dbOrder

	if err = o.db.Where("order_hash in (?)", hashs).Find(&dbOrders).Error; err != nil {
		err = fmt.Errorf("QueryOrder:%w", err)
		return
	}

	for _, o := range dbOrders {
		if loadErr := o.unmarshalCancelReason(); loadErr != nil {
			logger.Warnf("load order cancel reason error:%v", loadErr)
		}
		orders = append(orders, &o.Order)
	}
	return
}

func (o *orderDAO) CreateOrder(order *model.Order) error {
	var t = &dbOrder{Order: *order}
	if err := t.marshalCancelReason(); err != nil {
		return fmt.Errorf("CreateOrder:%w", err)
	}

	t.CreatedAt = time.Now().UTC()
	t.setStatusByAmounts()
	order.OldStatus = order.Status
	order.Status = t.Status
	if err := o.db.Save(&t).Error; err != nil {
		return fmt.Errorf("CreateOrder:%w", err)
	}
	return nil
}

func (o *orderDAO) UpdateOrder(order *model.Order) error {
	var t = &dbOrder{Order: *order}
	if err := t.marshalCancelReason(); err != nil {
		return fmt.Errorf("UpdateOrder:%w", err)
	}

	t.UpdatedAt = time.Now().UTC()
	t.setStatusByAmounts()
	order.OldStatus = order.Status
	order.Status = t.Status
	if err := o.db.Save(&t).Error; err != nil {
		return fmt.Errorf("UpdateOrder:%w", err)
	}
	return nil
}

func (o *orderDAO) LoadMatchOrders(matchItems []*model.MatchItem) error {
	var err error
	for _, item := range matchItems {
		if item.Order, err = o.GetOrder(item.OrderHash); err != nil {
			return fmt.Errorf("LoadMatchOrders:%w", err)
		}
	}
	return nil
}
