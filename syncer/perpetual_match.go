package syncer

import (
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/chain"
	"github.com/mcarloai/mai-v3-broker/match"
	"github.com/shopspring/decimal"
)

type PerpetualMatchSyncer struct {
	chainCli chain.ChainClient
	match    *match.Server
}

func NewPerpetualMatchSyncer(chainCli chain.ChainClient, match *match.Server) *PerpetualMatchSyncer {
	return &PerpetualMatchSyncer{
		chainCli: chainCli,
		match:    match,
	}
}

func (p *PerpetualMatchSyncer) Rollback(syncCtx *SyncBlockContext) error {
	err := syncCtx.Dao.RollbackLaunchTransactions(syncCtx.RollbackBeginHeight, syncCtx.LatestBlockNumber)
	matchItems, err := syncCtx.Dao.RollbackMatchTransactions(syncCtx.RollbackBeginHeight, syncCtx.LatestBlockNumber)
	if err != nil {
		return fmt.Errorf("watchPerpetualMatch.rollback failed:%w", err)
	}

	orderHashs := make([]string, 0)
	rollbackAmounts := make(map[string]decimal.Decimal)
	for _, item := range matchItems {
		orderHashs = append(orderHashs, item.OrderHash)
		if amount, ok := rollbackAmounts[item.OrderHash]; ok {
			rollbackAmounts[item.OrderHash] = amount.Add(item.Amount)
		} else {
			rollbackAmounts[item.OrderHash] = item.Amount
		}
	}
	orders, err := syncCtx.Dao.GetOrderByHashs(orderHashs)
	if err != nil {
		return fmt.Errorf("watchPerpetualMatch.rollback GetOrderByHashs failed:%w", err)
	}
	for _, order := range orders {
		order.PendingAmount = order.PendingAmount.Add(rollbackAmounts[order.OrderHash])
		order.ConfirmedAmount = order.ConfirmedAmount.Sub(rollbackAmounts[order.OrderHash])
		err := syncCtx.Dao.UpdateOrder(order)
		if err != nil {
			return fmt.Errorf("watchPerpetualMatch.rollback update order failed:%w", err)
		}
	}

	return nil
}

func (p *PerpetualMatchSyncer) Forward(syncCtx *SyncBlockContext) error {
	// only rollback for match transactions, filter match event when waitTransactionReceipt success in match's launch
	return nil
}
