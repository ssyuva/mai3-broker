package match

import (
	"github.com/mcarloai/mai-v3-broker/common/mai3"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/common/orderbook"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
	"sort"
)

type OrderCancel struct {
	OrderHash string
	Status    model.OrderStatus
	ToCancel  decimal.Decimal
}

func (m *match) CheckOrderMargin(account *model.AccountStorage, order *model.Order) bool {
	g := m.perpetualContext.GovParams
	storage := m.perpetualContext.PerpStorage
	err := mai3.ComputeTradeWithPrice(storage, account, order.Price, order.Amount, g.LpFeeRate.Add(g.VaultFeeRate).Add(g.OperatorFeeRate))
	if err != nil {
		return true
	}
	computedAccount := mai3.ComputeAccount(g, storage, account)
	if !computedAccount.IsSafe {
		return false
	}

	return false
}

func (m *match) CheckCloseOnly(account *model.AccountStorage, order *model.Order) bool {
	if order.IsCloseOnly {
		if account.PositionAmount.IsZero() {
			return false
		} else if !account.PositionAmount.IsZero() && utils.HasTheSameSign(account.PositionAmount, order.Amount) {
			return false
		}
	}
	return true
}

func (m *match) CheckAndModifyCloseOnly(account *model.AccountStorage, activeOrders []*model.Order) []*OrderCancel {
	cancels := make([]*OrderCancel, 0)
	closeOnlyOrders := make([]*model.Order, 0)
	closeOnlyTotalAmount := decimal.Zero
	for _, order := range activeOrders {
		if order.IsCloseOnly {
			if utils.HasTheSameSign(account.PositionAmount, order.AvailableAmount) {
				cancel := &OrderCancel{
					OrderHash: order.OrderHash,
					Status:    order.Status,
					ToCancel:  order.AvailableAmount,
				}
				cancels = append(cancels, cancel)
				continue
			}
			closeOnlyOrders = append(closeOnlyOrders, order)
			closeOnlyTotalAmount = closeOnlyTotalAmount.Add(order.AvailableAmount)
		}
	}

	if closeOnlyTotalAmount.Abs().GreaterThan(account.PositionAmount.Abs()) {
		if account.PositionAmount.IsNegative() {
			// order long
			sort.Slice(closeOnlyOrders, func(i, j int) bool {
				return closeOnlyOrders[i].Price.LessThan(closeOnlyOrders[j].Price)
			})
		} else {
			// order short
			sort.Slice(closeOnlyOrders, func(i, j int) bool {
				return closeOnlyOrders[i].Price.GreaterThan(closeOnlyOrders[j].Price)
			})
		}
		amountToBeCanceled := closeOnlyTotalAmount.Abs().Sub(account.PositionAmount.Abs())
		if closeOnlyTotalAmount.IsPositive() {
			amountToBeCanceled = amountToBeCanceled.Neg()
		}
		for _, order := range closeOnlyOrders {
			if !amountToBeCanceled.IsZero() && amountToBeCanceled.Abs().GreaterThanOrEqual(order.AvailableAmount.Abs()) {
				cancel := &OrderCancel{
					OrderHash: order.OrderHash,
					Status:    order.Status,
					ToCancel:  order.AvailableAmount,
				}
				cancels = append(cancels, cancel)
				amountToBeCanceled = amountToBeCanceled.Sub(order.AvailableAmount)
			} else if !amountToBeCanceled.IsZero() && amountToBeCanceled.Abs().LessThan(order.AvailableAmount.Abs()) {
				cancel := &OrderCancel{
					OrderHash: order.OrderHash,
					Status:    order.Status,
					ToCancel:  amountToBeCanceled,
				}
				cancels = append(cancels, cancel)
				amountToBeCanceled = amountToBeCanceled.Sub(amountToBeCanceled)
				break
			}
		}
	}
	return cancels
}

// func normalizeOrders(orders []*model.Order) (bids, asks, cancels, closeOnly []*model.Orders) {
// 	for _, order := range orders {

// 	}
// }

type MatchItem struct {
	Order              *orderbook.MemoryOrder // NOTE: mutable! should only be modified where execute match
	OrderOriginAmount  decimal.Decimal
	OrderCancelAmounts []decimal.Decimal
	OrderCancelReasons []model.CancelReasonType
	OrderTotalCancel   decimal.Decimal

	MatchedAmount decimal.Decimal
}

func (m *match) MatchOrderSideBySide() []*MatchItem {
	result := make([]*MatchItem, 0)
	var tradePrice decimal.Decimal
	isBuy := true

	for {
		if len(result) == mai3.MaiV3MaxMatchGroup {
			break
		}

		minAsk := m.orderbook.MinAsk()
		maxBid := m.orderbook.MaxBid()

		if maxBid != nil && isBuy {
			tradePrice = *maxBid
			result = m.matchOneSide(tradePrice, isBuy, result)
		} else if minAsk != nil && !isBuy {
			tradePrice = *minAsk
			result = m.matchOneSide(tradePrice, isBuy, result)
		} else {
			break
		}
		isBuy = !isBuy
	}

	return result
}

func (m *match) matchOneSide(tradePrice decimal.Decimal, isBuy bool, result []*MatchItem) []*MatchItem {
	orders := make([]*orderbook.MemoryOrder, 0)
	if isBuy {
		orders = append(orders, m.orderbook.GetBidOrdersByPrice(tradePrice)...)
	} else {
		orders = append(orders, m.orderbook.GetAskOrdersByPrice(tradePrice)...)
	}
	if len(orders) == 0 {
		return result
	}

	maxTradeAmount := mai3.ComputeAMMAmountWithPrice(m.perpetualContext.GovParams, m.perpetualContext.PerpStorage,
		m.perpetualContext.AMM, isBuy, tradePrice)
	if maxTradeAmount.IsZero() || !utils.HasTheSameSign(maxTradeAmount, orders[0].Amount) {
		return result
	}
	for _, order := range orders {
		if len(result) == mai3.MaiV3MaxMatchGroup {
			return result
		}
		account, err := m.chainCli.GetMarginAccount(m.ctx, m.perpetual.PerpetualAddress, order.Trader)
		if err != nil {
			logger.Errorf("matchOneSide: GetMarginAccount fail! err:%s", err.Error())
			return result
		}

		if maxTradeAmount.Abs().GreaterThanOrEqual(order.Amount.Abs()) {
			matchItem := &MatchItem{
				Order:              order,
				OrderCancelAmounts: make([]decimal.Decimal, 0),
				OrderCancelReasons: make([]model.CancelReasonType, 0),
				OrderTotalCancel:   decimal.Zero,
				OrderOriginAmount:  order.Amount,
				MatchedAmount:      order.Amount,
			}
			result = append(result, matchItem)
			_, _, _, _, err = mai3.ComputeAMMTrade(m.perpetualContext.GovParams, m.perpetualContext.PerpStorage, account,
				m.perpetualContext.AMM, order.Amount)
			if err != nil {
				logger.Errorf("matchOneSide: ComputeAMMTrade fail. err:%s", err)
				return result
			}
			maxTradeAmount = maxTradeAmount.Sub(order.Amount)
		} else {
			matchItem := &MatchItem{
				Order:              order,
				OrderCancelAmounts: make([]decimal.Decimal, 0),
				OrderCancelReasons: make([]model.CancelReasonType, 0),
				OrderTotalCancel:   decimal.Zero,
				OrderOriginAmount:  order.Amount,
				MatchedAmount:      maxTradeAmount,
			}
			result = append(result, matchItem)
			_, _, _, _, err = mai3.ComputeAMMTrade(m.perpetualContext.GovParams, m.perpetualContext.PerpStorage, account,
				m.perpetualContext.AMM, maxTradeAmount)
			if err != nil {
				logger.Errorf("matchOneSide: ComputeAMMTrade fail. err:%s", err)
				return result
			}
			if order.Amount.Sub(maxTradeAmount).Abs().LessThan(order.MinTradeAmount.Abs()) {
				matchItem.OrderCancelAmounts = append(matchItem.OrderCancelAmounts, order.Amount.Sub(maxTradeAmount))
				matchItem.OrderCancelReasons = append(matchItem.OrderCancelReasons, model.CancelReasonRemainTooSmall)
				matchItem.OrderTotalCancel = order.Amount.Sub(maxTradeAmount)
			}
			break
		}

	}
	return result
}
