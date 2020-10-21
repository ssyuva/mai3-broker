package dao

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/micro/go-micro/v2/logger"
	"github.com/shopspring/decimal"
)

const COUNT_ORDER_LIMIT = 5000 // never scan more rows when counting the orders
const ordersTableName = "orders"

type OrderDAO interface {
	CreateOrder(order *model.Order) error
	CreateStopOrder(order *model.Order) error
	GetOrder(orderID string, forUpdate bool) (*model.Order, error)
	QueryOrder(traderAddress string, perpetualAddress string, status []model.OrderStatus, beforeOrderID, afterOrderID int64, limit int) ([]*model.Order, error)
	GetOrderByIDs(ids []string) ([]*model.Order, error)
	GetOrderBySignature(signature string) (*model.Order, error)
	UpdateOrder(order *model.Order) error
}
type dbOrder struct {
	model.Order
	CancelsJson string `json:"-" db:"cancels_json"`
}

func (dbOrder) TableName() string {
	return ordersTableName
}

func (o *dbOrder) setStatusByAmounts() {
	if o.ConfirmedAmount.Equal(o.Amount) {
		o.Status = model.OrderFullFilled
	} else if o.CanceledAmount.Equal(o.Amount) {
		o.Status = model.OrderCanceled
	} else if o.AvailableAmount.Add(o.PendingAmount).GreaterThan(decimal.Zero) {
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

func (o *orderDAO) GetOrder(orderID string, forUpdate bool) (*model.Order, error) {
	var order dbOrder
	db := o.db
	if forUpdate {
		db = db.Set("gorm:query_option", "FOR UPDATE")
	}
	if err := db.Where("order_hash = ?", orderID).First(&order).Error; err != nil {
		return nil, fmt.Errorf("GetOrder:%w", err)
	}
	if err := order.unmarshalCancelReason(); err != nil {
		logger.Warnf("load order cancel reason error:%v", err)
	}
	return &order.Order, nil
}

func (o *orderDAO) GetOrderBySignature(signature string) (*model.Order, error) {
	var order dbOrder
	db := o.db
	if err := db.Where("signature = ?", signature).First(&order).Error; err != nil {
		return nil, fmt.Errorf("GetOrder:%w", err)
	}
	return &order.Order, nil
}

func (o *orderDAO) QueryOrder(traderAddress string, perpetualAddress string, status []model.OrderStatus, beforeOrderID, afterOrderID int64, limit int) (orders []*model.Order, err error) {
	var dbOrders []*dbOrder
	where := o.db.Table(ordersTableName)

	if traderAddress != "" {
		where = where.Where("trader_address=?", traderAddress)
	}

	if perpetualAddress != "" {
		where = where.Where("perpetual_address = ?", perpetualAddress)
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

	where = where.Order("id desc")
	if limit > 0 {
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

func (o *orderDAO) GetOrderByIDs(ids []string) (orders []*model.Order, err error) {
	var dbOrders []*dbOrder

	if err = o.db.Where("order_hash in (?)", ids).Find(&dbOrders).Error; err != nil {
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

func (o *orderDAO) CreateStopOrder(order *model.Order) error {
	var t = &dbOrder{Order: *order}
	if err := t.marshalCancelReason(); err != nil {
		return fmt.Errorf("CreateOrder:%w", err)
	}

	t.CreatedAt = time.Now().UTC()
	t.Status = model.OrderStop
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
