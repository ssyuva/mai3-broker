package match

import (
	"github.com/mcarloai/mai-v3-broker/common/mai3"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/common/orderbook"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
)

var TradeAmountRelaxFactor = decimal.NewFromFloat(0.99)

func (m *match) CheckOrderMargin(poolStorage *model.LiquidityPoolStorage, account *model.AccountStorage, order *model.Order, amount decimal.Decimal) bool {
	perpetual, ok := poolStorage.Perpetuals[m.perpetual.PerpetualIndex]
	if !ok {
		return false
	}
	computedAccount, err := mai3.ComputeAccount(poolStorage, m.perpetual.PerpetualIndex, account)
	if err != nil || !computedAccount.IsSafe {
		return false
	}

	// position after trade
	position := account.PositionAmount.Add(amount)

	// mark price * Abs(position + amount) / (currentMarginBalance - mark price * Abs(amount) * fee)
	feeRate := perpetual.LpFeeRate.Add(poolStorage.VaultFeeRate).Add(perpetual.OperatorFeeRate)
	tradeFee := perpetual.MarkPrice.Mul(amount.Abs()).Mul(feeRate)
	if computedAccount.MarginBalance.Sub(tradeFee).LessThanOrEqual(decimal.Zero) {
		return false
	}
	lev := position.Abs().Mul(perpetual.MarkPrice).Div(computedAccount.MarginBalance.Sub(tradeFee))
	maxLev := decimal.NewFromInt(1).Div(perpetual.InitialMarginRate)
	if lev.LessThan(decimal.Zero) || lev.GreaterThan(maxLev) {
		logger.Warnf("trader: %s amount: %s lev: %s greater than MaxLeverage: %s", order.TraderAddress, amount, lev, maxLev)
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
	if poolStorage == nil || err != nil {
		logger.Errorf("MatchOrderSideBySide: GetLiquidityPoolStorage fail! err:%s", err.Error())
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
			logger.Infof("matchedAmount: %s orderAmount:%s", order.Amount, order.Amount)
			_, err = mai3.ComputeAMMTrade(poolStorage, m.perpetual.PerpetualIndex, account, order.Amount)
			if err != nil {
				logger.Errorf("matchOneSide: ComputeAMMTrade fail. err:%s", err)
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
			_, err = mai3.ComputeAMMTrade(poolStorage, m.perpetual.PerpetualIndex, account, matchedAmount)
			if err != nil {
				logger.Errorf("matchOneSide: ComputeAMMTrade fail. err:%s", err)
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
