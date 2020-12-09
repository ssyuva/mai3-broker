package match

import (
	"errors"
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/mai3"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/common/orderbook"
	"github.com/shopspring/decimal"
	"sort"
	"time"

	logger "github.com/sirupsen/logrus"
)

func (m *match) CheckNewOrder(order *model.Order, activeOrders []*model.Order) (modifys []*model.Order, err error) {
	account, err := m.chainCli.GetMarginAccount(m.ctx, m.perpetual.PerpetualAddress, order.TraderAddress)
	if err != nil {
		return
	}
	if order.IsCloseOnly && !account.PositionAmount.IsZero() && utils.HasTheSameSign(account.PositionAmount, order.Amount) {
		err = fmt.Errorf("order amount hash same side with account position")
		return
	}

	activeOrders = append(activeOrders, order)
	modifys = m.CheckAndModifyCloseOnly(account, activeOrders)
	cancels, err := m.ComputeOrderBalance(account, activeOrders)
	if err != nil {
		return
	}
	if len(cancels) > 0 {
		err = errors.New("insufficient balance")
		return
	}
	return
}

func (m *match) ComputeOrderBalance(account *model.AccountStorage, orders []*model.Order) (cancels []*model.Order, err error) {
	return
}

func (m *match) CheckAndModifyCloseOnly(account *model.AccountStorage, activeOrders []*model.Order) []*model.Order {
	modifys := make([]*model.Order, 0)
	closeOnlyOrders := make([]*model.Order, 0)
	closeOnlyTotalAmount := decimal.Zero
	for _, order := range activeOrders {
		if order.IsCloseOnly {
			if utils.HasTheSameSign(account.PositionAmount, order.AvailableAmount) {
				cancelAmount := order.AvailableAmount
				order.CanceledAmount = order.CanceledAmount.Add(cancelAmount)
				order.AvailableAmount = decimal.Zero
				order.CancelReasons = append(order.CancelReasons, &model.OrderCancelReason{
					Reason:          model.CancelReasonAdminCancel,
					Amount:          cancelAmount,
					CanceledAt:      time.Now().UTC(),
					TransactionHash: "",
				})
				modifys = append(modifys, order)
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
				cancelAmount := order.AvailableAmount
				order.CanceledAmount = order.CanceledAmount.Add(cancelAmount)
				order.AvailableAmount = decimal.Zero
				order.CancelReasons = append(order.CancelReasons, &model.OrderCancelReason{
					Reason:          model.CancelReasonAdminCancel,
					Amount:          cancelAmount,
					CanceledAt:      time.Now().UTC(),
					TransactionHash: "",
				})
				modifys = append(modifys, order)
				amountToBeCanceled = amountToBeCanceled.Sub(order.AvailableAmount)
			} else if !amountToBeCanceled.IsZero() && amountToBeCanceled.Abs().LessThan(order.AvailableAmount.Abs()) {
				cancelAmount := amountToBeCanceled
				order.CanceledAmount = order.CanceledAmount.Add(cancelAmount)
				order.AvailableAmount = order.AvailableAmount.Sub(cancelAmount)
				order.CancelReasons = append(order.CancelReasons, &model.OrderCancelReason{
					Reason:          model.CancelReasonAdminCancel,
					Amount:          cancelAmount,
					CanceledAt:      time.Now().UTC(),
					TransactionHash: "",
				})
				modifys = append(modifys, order)
				amountToBeCanceled = amountToBeCanceled.Sub(cancelAmount)
			}
		}
	}
	return modifys
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
	var isBuy bool
	leverage := m.perpetualContext.GovParams.TargetLeverage
	feeRate := m.perpetualContext.GovParams.VaultFeeRate.Add(m.perpetualContext.GovParams.LpFeeRate).Add(m.perpetualContext.GovParams.OperatorFeeRate)

	markPrice := m.perpetualContext.PerpStorage.MarkPrice
	minAsk := m.orderbook.MinAsk()
	maxBid := m.orderbook.MaxBid()

	if maxBid == nil && minAsk == nil {
		return result
	}
	if minAsk != nil {
		tradePrice = *minAsk
		isBuy = false
	}
	if maxBid != nil && (markPrice.Sub(*maxBid).LessThan((*minAsk).Sub(markPrice))) {
		tradePrice = *maxBid
		isBuy = true
	}

	result = m.matchOneSide(tradePrice, leverage, feeRate, isBuy, result)
	if len(result) < mai3.MaiV3MaxMatchGroup {
		isBuy = !isBuy
		if isBuy && maxBid != nil {
			tradePrice = *maxBid
		} else if !isBuy && minAsk != nil {
			tradePrice = *minAsk
		} else {
			tradePrice = decimal.Zero
		}
		if !tradePrice.IsZero() {
			result = m.matchOneSide(tradePrice, leverage, feeRate, isBuy, result)
		}
	}

	return result
}

func (m *match) matchOneSide(tradePrice, leverage, feeRate decimal.Decimal, isBuy bool, result []*MatchItem) []*MatchItem {
	orders := make([]*orderbook.MemoryOrder, 0)
	if isBuy {
		orders = append(orders, m.orderbook.GetBidOrdersByPrice(tradePrice)...)
	} else {
		orders = append(orders, m.orderbook.GetAskOrdersByPrice(tradePrice)...)
	}

	for _, order := range orders {
		if len(result) == mai3.MaiV3MaxMatchGroup {
			return result
		}
		account, err := m.chainCli.GetMarginAccount(m.ctx, m.perpetual.PerpetualAddress, order.Trader)
		if err != nil {
			logger.Errorf("GetMarginAccount fail! err:%s", err.Error())
			continue
		}

		amount := mai3.ComputeMaxTradeAmountWithPrice(m.perpetualContext.GovParams, m.perpetualContext.PerpStorage,
			account, tradePrice, leverage, feeRate, isBuy)
		amount = decimal.Min(amount.Abs(), order.Amount.Abs())
		if !isBuy {
			amount = amount.Neg()
		}
		if !amount.IsZero() && amount.Abs().GreaterThanOrEqual(m.minTradeAmount) {
			matchItem := &MatchItem{
				Order:              order,
				OrderCancelAmounts: make([]decimal.Decimal, 0),
				OrderCancelReasons: make([]model.CancelReasonType, 0),
				OrderTotalCancel:   decimal.Zero,
				OrderOriginAmount:  order.Amount,
				MatchedAmount:      amount,
			}
			result = append(result, matchItem)
			// full match
			if order.Amount.Abs().GreaterThanOrEqual(amount.Abs()) {
				if order.Amount.Sub(amount).Abs().LessThan(m.minTradeAmount) {
					matchItem.OrderCancelAmounts = append(matchItem.OrderCancelAmounts, order.Amount.Sub(amount))
					matchItem.OrderCancelReasons = append(matchItem.OrderCancelReasons, model.CancelReasonRemainTooSmall)
					matchItem.OrderTotalCancel = order.Amount.Sub(amount)
				}
				break
			}
		}

	}
	return result
}
