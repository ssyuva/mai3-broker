package match

import (
	"github.com/mcarloai/mai-v3-broker/common/mai3"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/common/orderbook"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
	"sort"
)

type OrderCancel struct {
	OrderHash string
	Status    model.OrderStatus
	ToCancel  decimal.Decimal
}

func (m *match) CheckOrderMargin(poolStorage *model.LiquidityPoolStorage, account *model.AccountStorage, order *model.Order) bool {
	perpetualStorage, ok := poolStorage.Perpetuals[m.perpetual.PerpetualIndex]
	if !ok {
		return false
	}
	err := mai3.ComputeTradeWithPrice(poolStorage, m.perpetual.PerpetualIndex, account,
		order.Price, order.Amount,
		poolStorage.VaultFeeRate.Add(perpetualStorage.LpFeeRate).Add(perpetualStorage.OperatorFeeRate))
	if err != nil {
		return false
	}
	computedAccount, err := mai3.ComputeAccount(poolStorage, m.perpetual.PerpetualIndex, account)
	if err != nil || !computedAccount.IsSafe {
		return false
	}

	return true
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

type MatchItem struct {
	Order              *orderbook.MemoryOrder // NOTE: mutable! should only be modified where execute match
	OrderCancelAmounts []decimal.Decimal
	OrderCancelReasons []model.CancelReasonType
	OrderTotalCancel   decimal.Decimal

	MatchedAmount decimal.Decimal
}

func (m *match) MatchOrderSideBySide() []*MatchItem {
	result := make([]*MatchItem, 0)
	bidPrices := m.orderbook.GetBidPricesDesc()
	askPrices := m.orderbook.GetAskPricesAsc()
	bidIdx := 0
	askIdx := 0
	bidContinue := true
	askContinue := true

	if len(bidPrices) == 0 && len(askPrices) == 0 {
		return result
	}
	// compute match orders
	poolStorage, err := m.chainCli.GetLiquidityPoolStorage(m.ctx, conf.Conf.ReaderAddress, m.perpetual.LiquidityPoolAddress)
	if err != nil {
		logger.Errorf("MatchOrderSideBySide: GetLiquidityPoolStorage fail! err:%s", err.Error())
		return result
	}
	for {
		if len(bidPrices) > bidIdx {
			result, bidContinue = m.matchOneSide(poolStorage, bidPrices[bidIdx], true, result)
			bidIdx++
		}

		if len(askPrices) > askIdx {
			result, askContinue = m.matchOneSide(poolStorage, askPrices[askIdx], false, result)
			askIdx++
		}

		if (!bidContinue && !askContinue) || ((len(bidPrices) <= bidIdx) && len(askPrices) <= askIdx) {
			break
		}
	}

	return result
}

func (m *match) matchOneSide(poolStorage *model.LiquidityPoolStorage, tradePrice decimal.Decimal, isBuy bool, result []*MatchItem) ([]*MatchItem, bool) {
	orders := make([]*orderbook.MemoryOrder, 0)
	if isBuy {
		orders = append(orders, m.orderbook.GetBidOrdersByPrice(tradePrice)...)
	} else {
		orders = append(orders, m.orderbook.GetAskOrdersByPrice(tradePrice)...)
	}
	if len(orders) == 0 {
		return result, true
	}

	_, ok := poolStorage.Perpetuals[m.perpetual.PerpetualIndex]
	if !ok {
		return result, false
	}

	maxTradeAmount := mai3.ComputeAMMAmountWithPrice(poolStorage, m.perpetual.PerpetualIndex, isBuy, tradePrice)
	if maxTradeAmount.IsZero() || !utils.HasTheSameSign(maxTradeAmount, orders[0].Amount) {
		return result, false
	}

	bestPrice := decimal.Zero
	for _, order := range orders {
		if len(result) == mai3.MaiV3MaxMatchGroup {
			return result, false
		}

		// check stop order
		if order.Type == model.StopLimitOrder || order.Type == model.TakeProfitOrder {
			bestPrice = mai3.ComputeBestAskBidPrice(poolStorage, m.perpetual.PerpetualIndex, order.Amount.IsNegative())
			// 0 after compute
			if bestPrice.IsZero() {
				continue
			}

			if order.Type == model.StopLimitOrder {
				// When amount > 0, if stop loss order: index price must >= trigger price,
				// When amount < 0, if stop loss order: index price must <= trigger price,
				if order.Amount.IsPositive() && bestPrice.LessThan(order.TriggerPrice) {
					continue
				} else if order.Amount.IsNegative() && bestPrice.GreaterThan(order.TriggerPrice) {
					continue
				}
			} else {
				// When amount > 0, if take profit order: index price must <= trigger price,
				// When amount < 0, if take profit order: index price must >= trigger price,
				if order.Amount.IsPositive() && bestPrice.GreaterThan(order.TriggerPrice) {
					continue
				} else if order.Amount.IsNegative() && bestPrice.LessThan(order.TriggerPrice) {
					continue
				}
			}
		}

		account, err := m.chainCli.GetAccountStorage(m.ctx, conf.Conf.ReaderAddress, m.perpetual.PerpetualIndex, m.perpetual.LiquidityPoolAddress, order.Trader)
		if err != nil {
			logger.Errorf("matchOneSide: GetAccountStorage fail! err:%s", err.Error())
			return result, false
		}

		if maxTradeAmount.Abs().GreaterThanOrEqual(order.Amount.Abs()) {
			matchItem := &MatchItem{
				Order:              order,
				OrderCancelAmounts: make([]decimal.Decimal, 0),
				OrderCancelReasons: make([]model.CancelReasonType, 0),
				OrderTotalCancel:   decimal.Zero,
				MatchedAmount:      order.Amount,
			}
			_, err = mai3.ComputeAMMTrade(poolStorage, m.perpetual.PerpetualIndex, account, order.Amount)
			if err != nil {
				logger.Errorf("matchOneSide: ComputeAMMTrade fail. err:%s", err)
				return result, false
			}
			result = append(result, matchItem)
			maxTradeAmount = maxTradeAmount.Sub(order.Amount)
		} else {
			matchItem := &MatchItem{
				Order:              order,
				OrderCancelAmounts: make([]decimal.Decimal, 0),
				OrderCancelReasons: make([]model.CancelReasonType, 0),
				OrderTotalCancel:   decimal.Zero,
				MatchedAmount:      maxTradeAmount,
			}
			_, err = mai3.ComputeAMMTrade(poolStorage, m.perpetual.PerpetualIndex, account, maxTradeAmount)
			if err != nil {
				logger.Errorf("matchOneSide: ComputeAMMTrade fail. err:%s", err)
				return result, false
			}
			result = append(result, matchItem)
			if order.Amount.Sub(maxTradeAmount).Abs().LessThan(order.MinTradeAmount.Abs()) {
				matchItem.OrderCancelAmounts = append(matchItem.OrderCancelAmounts, order.Amount.Sub(maxTradeAmount))
				matchItem.OrderCancelReasons = append(matchItem.OrderCancelReasons, model.CancelReasonRemainTooSmall)
				matchItem.OrderTotalCancel = order.Amount.Sub(maxTradeAmount)
			}
			break
		}
	}

	return result, true
}
