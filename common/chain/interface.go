package chain

import (
	"context"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/shopspring/decimal"
	"math/big"
)

type ChainClient interface {
	HeaderByNumber(ctx context.Context, number *big.Int) (*model.BlockHeader, error)
	HeaderByHash(ctx context.Context, hash string) (*model.BlockHeader, error)
	GetMarginAccount(ctx context.Context, perpetualAddress, account string) (*model.AccountStorage, error)
	GetPrice(ctx context.Context, oracle string) (decimal.Decimal, error)
	FilterCreatePerpetual(ctx context.Context, factoryAddress string, start, end uint64) ([]*model.PerpetualEvent, error)
	FilterMatch(ctx context.Context, perpetualAddress string, start, end uint64) ([]*model.MatchEvent, error)
	WaitForReceipt(ctx context.Context, transactionHash string) (blockNumber uint64, blockHash string, succ bool, err error)
}
