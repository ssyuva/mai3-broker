package orderbook

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/cevaris/ordered_map"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/petar/GoLLRB/llrb"
	"github.com/shopspring/decimal"
)

type (
	MatchItem struct {
		MakerOrder              *MemoryOrder // NOTE: mutable! should only be modified where execute match
		MakerOrderOriginAmount  decimal.Decimal
		MakerOrderCancelAmounts []decimal.Decimal
		MakerOrderCancelReasons []model.CancelReasonType
		MakerOrderTotalCancel   decimal.Decimal

		MatchedAmount        decimal.Decimal
		IsMakerForceCanceled bool
	}

	MemoryOrder struct {
		ID               string          `json:"id"`
		PerpetualAddress string          `json:"perpetualAddress"`
		ComparePrice     decimal.Decimal `json:"-"`
		Price            decimal.Decimal `json:"price"`
		StopPrice        decimal.Decimal `json:"stopPrice"`
		Amount           decimal.Decimal `json:"amount"`
		Side             model.OrderSide `json:"side"`
		Type             model.OrderType `json:"type"`
		Trader           string          `json:"trader"`
		GasFeeAmount     decimal.Decimal `json:"gasFeeAmount"`
	}
)

type priceLevel struct {
	price       decimal.Decimal
	totalAmount decimal.Decimal
	orderMap    *ordered_map.OrderedMap
}

func newPriceLevel(price decimal.Decimal) *priceLevel {
	return &priceLevel{
		price:       price,
		totalAmount: decimal.Zero,
		orderMap:    ordered_map.NewOrderedMap(),
	}
}

func (p *priceLevel) Len() int {
	return p.orderMap.Len()
}

func (p *priceLevel) InsertOrder(order *MemoryOrder) error {
	//logger.Warnf("InsertOrder:%s", order.ID)

	if _, ok := p.orderMap.Get(order.ID); ok {
		return fmt.Errorf("priceLevel can't add order which is already in this priceLevel. priceLevel: %s, orderID: %s", p.price.String(), order.ID)
	}

	p.orderMap.Set(order.ID, order)
	p.totalAmount = p.totalAmount.Add(order.Amount)
	return nil
}

func (p *priceLevel) RemoveOrder(orderID string) (*MemoryOrder, error) {
	orderItem, ok := p.orderMap.Get(orderID)

	if !ok {
		return nil, fmt.Errorf("priceLevel can't remove order which is not in this priceLevel. priceLevel: %s:%w", p.price.String(), OrderNotFoundError)
	}

	order := orderItem.(*MemoryOrder)
	p.orderMap.Delete(order.ID)
	p.totalAmount = p.totalAmount.Sub(order.Amount)
	return order, nil
}

func (p *priceLevel) GetOrder(id string) (order *MemoryOrder, exist bool) {
	orderItem, exist := p.orderMap.Get(id)
	if !exist {
		return nil, exist
	}

	return orderItem.(*MemoryOrder), exist
}

func (p *priceLevel) ChangeOrder(orderID string, changeAmount decimal.Decimal) error {
	orderItem, ok := p.orderMap.Get(orderID)

	if !ok {
		return fmt.Errorf("can't change order which is not in this priceLevel. priceLevel: %s:%w", p.price.String(), OrderNotFoundError)
	}

	order := orderItem.(*MemoryOrder)
	newAmount := order.Amount.Add(changeAmount)
	if newAmount.IsNegative() {
		return fmt.Errorf("can't change order[%s], after change the amount is negative. old amount=%s change amount=%s",
			order.ID, order.Amount, changeAmount)
	}

	order.Amount = newAmount
	if newAmount.IsZero() {
		p.orderMap.Delete(orderID)
	}

	p.totalAmount = p.totalAmount.Add(changeAmount)
	return nil
}

func (p *priceLevel) Less(item llrb.Item) bool {
	another := item.(*priceLevel)
	return p.price.LessThan(another.price)
}

// Orderbook ...
type Orderbook struct {
	bidsTree *llrb.LLRB
	asksTree *llrb.LLRB

	lock sync.RWMutex

	Sequence  uint64
	UpdatedAt time.Time
}

// NewOrderbook return a new book
func NewOrderbook() *Orderbook {
	book := &Orderbook{
		bidsTree: llrb.New(),
		asksTree: llrb.New(),
	}

	return book
}

func (book *Orderbook) InsertOrder(order *MemoryOrder) error {
	//startTime := time.Now().UTC()
	book.lock.Lock()
	defer book.lock.Unlock()

	//log.Debug("cost in lock, InsertOrder :", order.ID, float64(time.Since(startTime))/1000000)

	var tree *llrb.LLRB
	if order.Side == "sell" {
		tree = book.asksTree
	} else {
		tree = book.bidsTree
	}

	price := tree.Get(newPriceLevel(order.ComparePrice))

	if price == nil {
		price = newPriceLevel(order.ComparePrice)
		tree.InsertNoReplace(price)
	}

	err := price.(*priceLevel).InsertOrder(order)
	if err != nil {
		return fmt.Errorf("InsertOrder:%w", err)
	}

	book.Sequence++
	book.UpdatedAt = time.Now().UTC()

	return nil
}

var OrderNotFoundError = errors.New("order not found")

func (book *Orderbook) RemoveOrder(order *MemoryOrder) error {
	book.lock.Lock()
	defer book.lock.Unlock()

	var tree *llrb.LLRB
	if order.Side == "sell" {
		tree = book.asksTree
	} else {
		tree = book.bidsTree
	}

	plItem := tree.Get(newPriceLevel(order.ComparePrice))
	if plItem == nil {
		return fmt.Errorf("remove order: find price level fail, price=%s:%w", order.ComparePrice, OrderNotFoundError)
	}

	price := plItem.(*priceLevel)

	if price == nil {
		return fmt.Errorf("price is nil when RemoveOrder, order: %+v", order)
	}

	_, err := price.RemoveOrder(order.ID)
	if err != nil {
		return fmt.Errorf("remove order fom price level fail:%w", err)
	}

	if price.Len() <= 0 {
		tree.Delete(price)
	}

	book.Sequence++
	book.UpdatedAt = time.Now().UTC()

	return nil
}

func (book *Orderbook) ChangeOrder(order *MemoryOrder, changeAmount decimal.Decimal) error {
	book.lock.Lock()
	defer book.lock.Unlock()

	var tree *llrb.LLRB
	if order.Side == "sell" {
		tree = book.asksTree
	} else {
		tree = book.bidsTree
	}

	plItem := tree.Get(newPriceLevel(order.ComparePrice))

	if plItem == nil {
		return fmt.Errorf("can't change order which is not in this orderbook. order: %+v:%w", order, OrderNotFoundError)
	}

	price := plItem.(*priceLevel)
	if price == nil {
		return fmt.Errorf("pl is nil when ChangeOrder, order: %+v", order)
	}

	if err := price.ChangeOrder(order.ID, changeAmount); err != nil {
		return fmt.Errorf("change order fail:%w", err)
	}
	if price.Len() <= 0 {
		tree.Delete(price)
	}

	book.Sequence++
	book.UpdatedAt = time.Now().UTC()

	return nil
}

func (book *Orderbook) GetOrder(id string, side model.OrderSide, price decimal.Decimal) (*MemoryOrder, bool) {
	book.lock.Lock()
	defer book.lock.Unlock()

	var tree *llrb.LLRB
	if side == model.SideSell {
		tree = book.asksTree
	} else {
		tree = book.bidsTree
	}

	pl := tree.Get(newPriceLevel(price))

	if pl == nil {
		return nil, false
	}

	return pl.(*priceLevel).GetOrder(id)
}

// MaxBid ...
func (book *Orderbook) MaxBid() *decimal.Decimal {
	book.lock.Lock()
	defer book.lock.Unlock()

	maxItem := book.bidsTree.Max()
	if maxItem != nil {
		return &maxItem.(*priceLevel).price
	}
	return nil
}

// MinAsk ...
func (book *Orderbook) MinAsk() *decimal.Decimal {
	book.lock.Lock()
	defer book.lock.Unlock()

	minItem := book.asksTree.Min()

	if minItem != nil {
		return &minItem.(*priceLevel).price
	}

	return nil
}

func (book *Orderbook) CanMatch(order *MemoryOrder) bool {
	if order.Side == model.SideBuy {
		minItem := book.asksTree.Min()
		if minItem == nil {
			return false
		}

		if order.ComparePrice.GreaterThanOrEqual(minItem.(*priceLevel).price) {
			return true
		}

		return false
	} else {
		maxItem := book.bidsTree.Max()
		if maxItem == nil {
			return false
		}

		if order.ComparePrice.LessThanOrEqual(maxItem.(*priceLevel).price) {
			return true
		}

		return false
	}
}
