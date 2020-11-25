package perp

import (
	"github.com/mcarloai/mai-v3-broker/common/model"
)

type OrderContext struct {
	Account   *model.AccountStorage
	Perpetual *model.PerpetualStorage
}

func CheckActiveOrders(orderCtx *OrderContext, activeOrders []*model.Order) (cancels, modify []*model.Orders, err error) {

}

func CheckAndModifyCloseOnly(activeOrders []*model.Order) ([]*model.Orders, bool) {

}

func normalizeOrders(orders []*model.Order) (bids, asks, cancels, closeOnly []*model.Orders) {
	for _, order := range orders {

	}
}
