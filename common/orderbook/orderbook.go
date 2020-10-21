package orderbook

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/cevaris/ordered_map"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/micro/go-micro/v2/logger"
	"github.com/petar/GoLLRB/llrb"
	"github.com/shopspring/decimal"
)

type (
	MatchResult struct {
		TakerOrder              *MemoryOrder // NOTE: mutable! should only be modified where execute match
		TakerOrderOriginAmount  decimal.Decimal
		TakerOrderCancelAmounts []decimal.Decimal
		TakerOrderCancelReasons []model.CancelReasonType
		TakerOrderTotalCancel   decimal.Decimal
		TotalMatched            decimal.Decimal

		MatchItems []*MatchItem
	}

	MatchItem struct {
		MakerOrder              *MemoryOrder // NOTE: mutable! should only be modified where execute match
		MakerOrderOriginAmount  decimal.Decimal
		MakerOrderCancelAmounts []decimal.Decimal
		MakerOrderCancelReasons []model.CancelReasonType
		MakerOrderTotalCancel   decimal.Decimal

		MatchedAmount        decimal.Decimal
		IsMakerForceCanceled bool
	}

	OrderContext interface{}

	MemoryOrder struct {
		ID               string          `json:"id"`
		perpetualAddress string          `json:"perpetualAddress"`
		Price            decimal.Decimal `json:"price"`
		StopPrice        decimal.Decimal `json:"stopPrice"`
		Amount           decimal.Decimal `json:"amount"`
		Side             model.OrderSide `json:"side"`
		Type             model.OrderType `json:"type"`
		Trader           string          `json:"trader"`
		GasFeeAmount     decimal.Decimal `json:"gasFeeAmount"`
	}
)

func (m *MatchItem) MakerOrderConsumeAmount() decimal.Decimal {
	return m.MakerOrderTotalCancel.Add(m.MatchedAmount)
}

func (m *MatchItem) IsMakerOrderDone() bool {
	consume := m.MakerOrderConsumeAmount()
	if consume.GreaterThan(m.MakerOrderOriginAmount) {
		logger.Errorf("TakerOrderIsDone: BUG!")
	}
	return consume.GreaterThanOrEqual(m.MakerOrderOriginAmount)
}

func (m *MatchItem) CancelMakerOrder(amount decimal.Decimal, reason model.CancelReasonType) {
	if amount.IsPositive() {
		m.MakerOrderCancelAmounts = append(m.MakerOrderCancelAmounts, amount)
		m.MakerOrderCancelReasons = append(m.MakerOrderCancelReasons, reason)
		m.MakerOrderTotalCancel = m.MakerOrderTotalCancel.Add(amount)
	}
}

func (m *MatchResult) HasMatch() bool {
	return len(m.MatchItems) > 0
}

func (m *MatchResult) TakerOrderConsumeAmount() decimal.Decimal {
	return m.TakerOrderTotalCancel.Add(m.TotalMatched)
}

func (m *MatchResult) TakerOrderRemainAmount() decimal.Decimal {
	reamin := m.TakerOrderOriginAmount.Sub(m.TakerOrderConsumeAmount())
	if reamin.IsNegative() {
		logger.Errorf("BUG: TakerOrderRemainAmount[%s] is negative. cancel[%s] match[%s]", reamin, m.TakerOrderTotalCancel, m.TotalMatched)
		return decimal.Zero
	}
	return reamin
}

func (m *MatchResult) IsTakerOrderDone() bool {
	return m.TakerOrderRemainAmount().IsZero()
}

func (m *MatchResult) CancelTakerOrder(amount decimal.Decimal, reason model.CancelReasonType) {
	if amount.IsPositive() {
		m.TakerOrderCancelAmounts = append(m.TakerOrderCancelAmounts, amount)
		m.TakerOrderCancelReasons = append(m.TakerOrderCancelReasons, reason)
		m.TakerOrderTotalCancel = m.TakerOrderTotalCancel.Add(amount)
	}
}

func (matchResult *MatchResult) AddMatch(makerOrder *MemoryOrder, amount decimal.Decimal) *MatchItem {
	m := &MatchItem{
		MakerOrder:             makerOrder,
		MatchedAmount:          amount,
		MakerOrderOriginAmount: makerOrder.Amount,
	}
	matchResult.MatchItems = append(matchResult.MatchItems, m)
	matchResult.TotalMatched = matchResult.TotalMatched.Add(amount)
	return m
}

func (matchResult *MatchResult) CancelMatch(m *MatchItem, reason model.CancelReasonType) {
	m.CancelMakerOrder(m.MatchedAmount, reason)
	matchResult.CancelTakerOrder(m.MatchedAmount, reason)
	matchResult.TotalMatched = matchResult.TotalMatched.Sub(m.MatchedAmount)
	m.MatchedAmount = decimal.Zero
}

// ForceCancelMaker is used to force the maker order to be canceled from the order book
func (matchResult *MatchResult) ForceCancelMaker(m *MatchItem) {
	matchResult.CancelTakerOrder(m.MatchedAmount, model.CancelReasonAdminCancel)
	matchResult.TotalMatched = matchResult.TotalMatched.Sub(m.MatchedAmount)
	m.MatchedAmount = decimal.Zero

	m.MakerOrderCancelAmounts = []decimal.Decimal{m.MakerOrderOriginAmount}
	m.MakerOrderCancelReasons = []model.CancelReasonType{model.CancelReasonAdminCancel}
	m.MakerOrderTotalCancel = m.MakerOrderOriginAmount

	m.IsMakerForceCanceled = true
}

func NewMatchResult(takerOrder *MemoryOrder) *MatchResult {
	return &MatchResult{
		TakerOrder:             takerOrder,
		TakerOrderOriginAmount: takerOrder.Amount,
		MatchItems:             make([]*MatchItem, 0),
	}
}

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

type OrderMatcher func(*MemoryOrder) bool

type OrderBookMarket interface {
	ID() string
	LotSize() decimal.Decimal
	Symbol() string
	GetMarketOrderMatcher(takerOrder *MemoryOrder, orderContext OrderContext, result *MatchResult) (OrderMatcher, error)
}

// Orderbook ...
type Orderbook struct {
	market   OrderBookMarket
	bidsTree *llrb.LLRB
	asksTree *llrb.LLRB

	lock sync.RWMutex

	Sequence  uint64
	UpdatedAt time.Time
}

// NewOrderbook return a new book
func NewOrderbook(market OrderBookMarket) *Orderbook {
	book := &Orderbook{
		market:   market,
		bidsTree: llrb.New(),
		asksTree: llrb.New(),
	}

	return book
}

func (book *Orderbook) UpdateMarket(market OrderBookMarket) {
	book.market = market
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

	price := tree.Get(newPriceLevel(order.Price))

	if price == nil {
		price = newPriceLevel(order.Price)
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

	plItem := tree.Get(newPriceLevel(order.Price))
	if plItem == nil {
		return fmt.Errorf("remove order: find price level fail, price=%s:%w", order.Price, OrderNotFoundError)
	}

	price := plItem.(*priceLevel)

	if price == nil {
		return fmt.Errorf("price is nil when RemoveOrder, book: %s, order: %+v", book.market.Symbol(), order)
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

	plItem := tree.Get(newPriceLevel(order.Price))

	if plItem == nil {
		return fmt.Errorf("can't change order which is not in this orderbook. book: %s, order: %+v:%w", book.market.Symbol(), order, OrderNotFoundError)
	}

	price := plItem.(*priceLevel)
	if price == nil {
		return fmt.Errorf("pl is nil when ChangeOrder, book: %s, order: %+v", book.market.Symbol(), order)
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
	if order.Type == "market" {
		return true
	}
	if order.Side == model.SideBuy {
		minItem := book.asksTree.Min()
		if minItem == nil {
			return false
		}

		if order.Price.GreaterThanOrEqual(minItem.(*priceLevel).price) {
			return true
		}

		return false
	} else {
		maxItem := book.bidsTree.Max()
		if maxItem == nil {
			return false
		}

		if order.Price.LessThanOrEqual(maxItem.(*priceLevel).price) {
			return true
		}

		return false
	}
}

func (book *Orderbook) getLimitOrderIterator(result *MatchResult) func(i llrb.Item) bool {
	// This function will be called multi times
	// Return false to break the loop
	takerOrder := result.TakerOrder
	leftAmount := result.TakerOrderRemainAmount()
	lotSize := book.market.LotSize()
	return func(i llrb.Item) bool {
		pl := i.(*priceLevel)

		if takerOrder.Side == "buy" && pl.price.GreaterThan(takerOrder.Price) {
			return false
		} else if takerOrder.Side == "sell" && pl.price.LessThan(takerOrder.Price) {
			return false
		}

		iter := pl.orderMap.IterFunc()
		for kv, ok := iter(); ok; kv, ok = iter() {
			if leftAmount.LessThanOrEqual(decimal.Zero) {
				break
			}

			bookOrder := kv.Value.(*MemoryOrder)
			matchedAmount := decimal.Min(leftAmount, bookOrder.Amount).Div(lotSize).Truncate(0).Mul(lotSize)
			matchItem := result.AddMatch(bookOrder, matchedAmount)
			makerRemain := bookOrder.Amount.Sub(matchedAmount)
			if makerRemain.LessThan(lotSize) {
				matchItem.CancelMakerOrder(makerRemain, model.CancelReasonRemainTooSmall)
			}
			leftAmount = leftAmount.Sub(matchedAmount)
		}

		return leftAmount.IsPositive()
	}
}
