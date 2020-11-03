package syncer

import (
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/chain"
	"github.com/mcarloai/mai-v3-broker/common/message"
	"github.com/shopspring/decimal"
)

type PerpetualMatchSyncer struct {
	chainCli  chain.ChainClient
	matchChan chan interface{}
}

func NewPerpetualMatchSyncer(chainCli chain.ChainClient, matchChan chan interface{}) *PerpetualMatchSyncer {
	return &PerpetualMatchSyncer{
		chainCli:  chainCli,
		matchChan: matchChan,
	}
}

func (p *PerpetualMatchSyncer) Rollback(syncCtx *SyncBlockContext) error {
	matchItems, err := syncCtx.Dao.RollbackTransactions(syncCtx.RollbackBeginHeight, syncCtx.LatestBlockNumber)
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

	for _, order := range orders {
		matchMsg := message.MatchMessage{
			Type:             message.MatchTypeChangeOrder,
			PerpetualAddress: order.PerpetualAddress,
			Payload: message.MatchChangeOrderPayload{
				OrderHash: order.OrderHash,
				Amount:    rollbackAmounts[order.OrderHash],
			},
		}

		p.matchChan <- matchMsg
	}
	return nil
}

func (p *PerpetualMatchSyncer) Forward(syncCtx *SyncBlockContext) error {
	// transactions, err := syncCtx.Dao.QueryUnconfirmedTransactions()
	// if err != nil {
	// 	return fmt.Errorf("watchPerpetualMatch.forward QueryUnconfirmedTransactions failed:%w", err)
	// }

	// for _, transaction := range transactions {
	// 	matchEvents, err := p.chainCli.FilterMatch(syncCtx.Context, transaction.PerpetualAddress, uint64(syncCtx.RollbackBeginHeight), uint64(syncCtx.LatestBlockNumber))
	// 	if err != nil {
	// 		return fmt.Errorf("watchPerpetualMatch.forward FilterMatch failed:%w", err)
	// 	}
	// }
	// return nil

	// only rollback for match transactions, filter match event when waitTransactionReceipt success in match's launch
	return nil
}
