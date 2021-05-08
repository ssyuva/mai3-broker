package match

import (
	"sort"

	"github.com/mcarloai/mai-v3-broker/common/mai3"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/common/orderbook"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
)

var TradeAmountRelaxFactor = decimal.NewFromFloat(0.99)
var _0 = decimal.Zero
var _1 = decimal.NewFromInt(1)

const ADDRESS_ZERO = "0x0000000000000000000000000000000000000000"

func splitActiveOrders(orders []*model.Order) (buys, sells []*model.Order) {
	for _, order := range orders {
		amount := order.AvailableAmount.Add(order.PendingAmount)
		if amount.LessThan(decimal.Zero) {
			// sell
			sells = append(sells, order)
		} else if amount.GreaterThan(decimal.Zero) {
			// buy
			buys = append(buys, order)
		}
	}
	sort.Slice(buys, func(i, j int) bool {
		return buys[i].Price.GreaterThan(buys[j].Price)
	})
	sort.Slice(sells, func(i, j int) bool {
		return sells[i].Price.LessThan(sells[j].Price)
	})
	return
}

func (m *match) openOrderCost(pool *model.LiquidityPoolStorage, order *model.Order, leverage decimal.Decimal) decimal.Decimal {
	perp, ok := pool.Perpetuals[m.perpetual.PerpetualIndex]
	if !ok {
		return _0
	}
	amount := order.AvailableAmount.Add(order.PendingAmount)
	feeRate := pool.VaultFeeRate.Add(perp.LpFeeRate).Add(perp.OperatorFeeRate)
	if order.ReferrerAddress != ADDRESS_ZERO {
		feeRate = feeRate.Add(perp.ReferrerRebateRate)
	}
	potentialLoss := _0
	if amount.GreaterThan(_0) && perp.MarkPrice.LessThan(order.Price) {
		potentialLoss = order.Price.Sub(perp.MarkPrice).Mul(amount)
	} else if amount.LessThan(_0) && perp.MarkPrice.GreaterThan(order.Price) {
		potentialLoss = perp.MarkPrice.Sub(order.Price).Mul(amount.Abs())
	}
	return order.Price.Mul(amount).Mul(_1.Div(leverage).Add(feeRate)).Add(potentialLoss)
}

func (m *match) sideAvailable(pool *model.LiquidityPoolStorage, account *model.AccountStorage, orders []*model.Order) (cancels []*OrderCancel, available decimal.Decimal) {
	cancels = make([]*OrderCancel, 0)
	available = account.WalletBalance
	if len(orders) == 0 {
		return
	}
	perpetual, ok := pool.Perpetuals[m.perpetual.PerpetualIndex]
	if !ok {
		return cancels, _0
	}
	computedAccount, err := mai3.ComputeAccount(pool, m.perpetual.PerpetualIndex, account)
	if err != nil {
		return cancels, _0
	}
	feeRate := pool.VaultFeeRate.Add(perpetual.LpFeeRate).Add(perpetual.OperatorFeeRate)
	remainPosition := account.PositionAmount
	remainMargin := computedAccount.MarginBalance
	remainOrders := make([]*model.Order, 0)
	for _, order := range orders {
		amount := order.AvailableAmount.Add(order.PendingAmount)
		close, _ := utils.SplitAmount(remainPosition, amount)
		if !close.IsZero() {
			newPosition := remainPosition.Add(close)
			newPositionMargin := perpetual.MarkPrice.Mul(newPosition.Abs()).Mul(perpetual.InitialMarginRate)
			pnl := perpetual.MarkPrice.Sub(order.Price).Mul(close)
			afterMargin := remainMargin.Add(pnl)
			fee := decimal.Min(decimal.Max(afterMargin.Sub(newPositionMargin), _0), order.Price.Mul(close.Abs()).Mul(feeRate))
			afterMargin = afterMargin.Sub(fee)
			if afterMargin.LessThan(_0) {
				// bankrupt when close. pretend all orders as open orders
				remainPosition = _0
				remainMargin = _0
				remainOrders = append(remainOrders, order)
			} else {
				// withdraw only if marginBalance >= IM
				withdraw := _0
				if afterMargin.GreaterThanOrEqual(newPositionMargin) {
					// withdraw = afterMargin - remainMargin * (1 - | close / remainPosition |)
					withdraw = close.Div(remainPosition).Abs()
					withdraw = _1.Mul(withdraw).Mul(remainMargin)
					withdraw = afterMargin.Sub(withdraw)
					withdraw = decimal.Max(_0, withdraw)
				}
				remainMargin = afterMargin.Sub(withdraw)
				available = available.Add(withdraw)
				remainPosition = remainPosition.Add(close)
				newOrderAmount := amount.Sub(close)
				// TODO: consider min trade amount
				if !newOrderAmount.IsZero() {
					// update order amount just for checking below, can not save order in db
					order.AvailableAmount = newOrderAmount.Sub(order.PendingAmount)
					remainOrders = append(remainOrders, order)
				}
			}
		} else {
			remainOrders = append(remainOrders, order)
		}
	}

	// if close = 0 && position = 0 && margin > 0
	if remainPosition.IsZero() {
		available = available.Add(remainMargin)
	}

	// open position
	for _, order := range remainOrders {
		cost := m.openOrderCost(pool, order, account.TargetLeverage)
		available = available.Sub(cost)
		if available.LessThan(_0) {
			cancel := &OrderCancel{
				OrderHash: order.OrderHash,
				Status:    order.Status,
				ToCancel:  order.AvailableAmount,
				Reason:    model.CancelReasonInsufficientFunds,
			}
			cancels = append(cancels, cancel)
		}
	}

	return
}

func (m *match) ComputeOrderAvailable(poolStorage *model.LiquidityPoolStorage, account *model.AccountStorage, orders []*model.Order) ([]*OrderCancel, decimal.Decimal) {
	cancels := make([]*OrderCancel, 0)
	buyOrders, sellOrders := splitActiveOrders(orders)
	buyCancels, buySideAvailable := m.sideAvailable(poolStorage, account, buyOrders)
	cancels = append(cancels, buyCancels...)
	sellCancels, sellSideAvailable := m.sideAvailable(poolStorage, account, sellOrders)
	cancels = append(cancels, sellCancels...)

	return cancels, decimal.Min(buySideAvailable, sellSideAvailable)
}

func (m *match) ComputeOrderCost(poolStorage *model.LiquidityPoolStorage, account *model.AccountStorage, order *model.Order, activeOrders []*model.Order) decimal.Decimal {
	_, oldAvailable := m.ComputeOrderAvailable(poolStorage, account, activeOrders)
	newOrders := append(activeOrders, order)
	_, newAvailable := m.ComputeOrderAvailable(poolStorage, account, newOrders)
	if newAvailable.LessThan(oldAvailable) {
		return oldAvailable.Sub(newAvailable)
	}
	return _0
}

func (m *match) CheckCloseOnly(account *model.AccountStorage, order *model.Order) decimal.Decimal {
	if order.IsCloseOnly {
		if account.PositionAmount.IsZero() {
			// position is 0, cancel all
			return order.AvailableAmount
		} else {
			// open amount has same sign with position, cancel all
			if utils.HasTheSameSign(account.PositionAmount, order.Amount) {
				return order.AvailableAmount
			}

			// closed amount greater than position, cancel some
			if order.AvailableAmount.Abs().GreaterThan(account.PositionAmount.Abs()) {
				return order.AvailableAmount.Add(account.PositionAmount)
			}
		}
	}
	return _0
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
	poolStorage := m.poolSyncer.GetPoolStorage(m.perpetual.LiquidityPoolAddress)
	if poolStorage == nil {
		logger.Errorf("MatchOrderSideBySide: GetLiquidityPoolStorage fail!")
		return result
	}

	orderGasLimit := mai3.GetGasFeeLimit(len(poolStorage.Perpetuals))
	maiV3MaxMatchGroup := conf.Conf.GasLimit / uint64(orderGasLimit)

	for {
		if len(bidPrices) > bidIdx {
			result, bidContinue = m.matchOneSide(poolStorage, bidPrices[bidIdx], true, result, maiV3MaxMatchGroup)
			bidIdx++
		} else {
			bidContinue = false
		}

		if len(askPrices) > askIdx {
			result, askContinue = m.matchOneSide(poolStorage, askPrices[askIdx], false, result, maiV3MaxMatchGroup)
			askIdx++
		} else {
			askContinue = false
		}

		if !bidContinue && !askContinue {
			break
		}
	}

	return result
}

func (m *match) matchOneSide(poolStorage *model.LiquidityPoolStorage, tradePrice decimal.Decimal, isBuy bool, result []*MatchItem, maiV3MaxMatchGroup uint64) ([]*MatchItem, bool) {
	orders := make([]*orderbook.MemoryOrder, 0)
	if isBuy {
		orders = append(orders, m.orderbook.GetBidOrdersByPrice(tradePrice)...)
	} else {
		orders = append(orders, m.orderbook.GetAskOrdersByPrice(tradePrice)...)
	}
	if len(orders) == 0 {
		return result, true
	}

	perpetual, ok := poolStorage.Perpetuals[m.perpetual.PerpetualIndex]
	if !ok || !perpetual.IsNormal {
		return result, false
	}

	maxTradeAmount := mai3.ComputeAMMAmountWithPrice(poolStorage, m.perpetual.PerpetualIndex, isBuy, tradePrice)
	logger.Infof("maxAmount:%s, isBuy:%v, tradePrice:%s perpetual:%s-%d ", maxTradeAmount, isBuy, tradePrice, m.perpetual.LiquidityPoolAddress, m.perpetual.PerpetualIndex)
	if maxTradeAmount.IsZero() || !utils.HasTheSameSign(maxTradeAmount, orders[0].Amount) {
		return result, false
	}

	for _, order := range orders {
		if uint64(len(result)) == maiV3MaxMatchGroup {
			return result, false
		}

		logger.Infof("memoryOrder:%+v", order)

		if maxTradeAmount.Abs().LessThan(order.MinTradeAmount.Abs()) {
			continue
		}

		// check stop order
		if order.Type == model.StopLimitOrder || order.Type == model.TakeProfitOrder {
			logger.Infof("indexPrice:%s", perpetual.IndexPrice)
			if perpetual.IndexPrice.IsZero() {
				continue
			}

			if order.Type == model.StopLimitOrder {
				// When amount > 0, if stop loss order: index price must >= trigger price,
				// When amount < 0, if stop loss order: index price must <= trigger price,
				if order.Amount.IsPositive() && perpetual.IndexPrice.LessThan(order.TriggerPrice) {
					continue
				} else if order.Amount.IsNegative() && perpetual.IndexPrice.GreaterThan(order.TriggerPrice) {
					continue
				}
			} else {
				// When amount > 0, if take profit order: index price must <= trigger price,
				// When amount < 0, if take profit order: index price must >= trigger price,
				if order.Amount.IsPositive() && perpetual.IndexPrice.GreaterThan(order.TriggerPrice) {
					continue
				} else if order.Amount.IsNegative() && perpetual.IndexPrice.LessThan(order.TriggerPrice) {
					continue
				}
			}
		}

		account, err := m.chainCli.GetAccountStorage(m.ctx, conf.Conf.ReaderAddress, m.perpetual.PerpetualIndex, m.perpetual.LiquidityPoolAddress, order.Trader)
		if account == nil || err != nil {
			logger.Errorf("matchOneSide: GetAccountStorage fail! err:%v", err)
			return result, false
		}

		balance, err := m.chainCli.BalanceOf(m.ctx, m.perpetual.CollateralAddress, order.Trader, m.perpetual.CollateralDecimals)
		if err != nil {
			logger.Errorf("matchOneSide: BalanceOf fail! err:%v", err)
			return result, false
		}
		allowance, err := m.chainCli.Allowance(m.ctx, m.perpetual.CollateralAddress, order.Trader, m.perpetual.LiquidityPoolAddress, m.perpetual.CollateralDecimals)
		if err != nil {
			logger.Errorf("matchOneSide: Allowance fail! err:%v", err)
			return result, false
		}
		account.WalletBalance = decimal.Min(balance, allowance)

		if maxTradeAmount.Abs().GreaterThanOrEqual(order.Amount.Abs()) {
			matchItem := &MatchItem{
				Order:              order,
				OrderCancelAmounts: make([]decimal.Decimal, 0),
				OrderCancelReasons: make([]model.CancelReasonType, 0),
				OrderTotalCancel:   decimal.Zero,
				MatchedAmount:      order.Amount,
			}
			logger.Infof("matchedAmount: %s orderAmount:%s", order.Amount, order.Amount)
			_, tradeIsSafe, _, err := mai3.ComputeAMMTrade(poolStorage, m.perpetual.PerpetualIndex, account, order.Amount)
			if err != nil || !tradeIsSafe {
				logger.Errorf("matchOneSide: ComputeAMMTrade fail or unsafe after trade. err:%s", err)
				return result, true
			}

			result = append(result, matchItem)
			maxTradeAmount = maxTradeAmount.Sub(order.Amount)
		} else {
			matchedAmount := maxTradeAmount.Mul(TradeAmountRelaxFactor).Round(mai3.DECIMALS)
			matchItem := &MatchItem{
				Order:              order,
				OrderCancelAmounts: make([]decimal.Decimal, 0),
				OrderCancelReasons: make([]model.CancelReasonType, 0),
				OrderTotalCancel:   decimal.Zero,
				MatchedAmount:      matchedAmount,
			}
			logger.Infof("matchedAmount: %s orderAmount:%s", matchedAmount, order.Amount)
			_, tradeIsSafe, _, err := mai3.ComputeAMMTrade(poolStorage, m.perpetual.PerpetualIndex, account, matchedAmount)
			if err != nil || !tradeIsSafe {
				logger.Errorf("matchOneSide: ComputeAMMTrade fail or unsafe after trade. err:%v", err)
				return result, true
			}
			if order.Amount.Sub(matchedAmount).Abs().LessThan(order.MinTradeAmount.Abs()) {
				logger.Infof("OrderCancelAmount: %s", order.Amount.Sub(matchedAmount))
				matchItem.OrderCancelAmounts = append(matchItem.OrderCancelAmounts, order.Amount.Sub(matchedAmount))
				matchItem.OrderCancelReasons = append(matchItem.OrderCancelReasons, model.CancelReasonRemainTooSmall)
				matchItem.OrderTotalCancel = order.Amount.Sub(matchedAmount)
			}
			result = append(result, matchItem)
			break
		}
	}

	return result, true
}
