package chain

import (
	"context"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/shopspring/decimal"
	"math/big"
)

type ChainClient interface {
	AddAccount(pk string) error
	GetSignAccount() (string, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*model.BlockHeader, error)
	HeaderByHash(ctx context.Context, hash string) (*model.BlockHeader, error)
	PendingNonceAt(ctx context.Context, account string) (uint64, error)
	TransactionByHash(ctx context.Context, txHash string) (bool, error)
	SendTransaction(ctx context.Context, tx *model.LaunchTransaction) (string, error)
	WaitTransactionReceipt(ctx context.Context, txHash string) (*model.Receipt, error)
	GetMarginAccount(ctx context.Context, perpetualAddress, account string) (*model.AccountStorage, error)
	GetPrice(ctx context.Context, oracle string) (decimal.Decimal, error)
	FilterCreatePerpetual(ctx context.Context, factoryAddress string, start, end uint64) ([]*model.PerpetualEvent, error)
	FilterMatch(ctx context.Context, perpetualAddress string, start, end uint64) ([]*model.MatchEvent, error)
	BatchTradeDataPack(orderParams []*model.WalletOrderParam, amounts []decimal.Decimal, gases []*big.Int) ([]byte, error)
	WaitForReceipt(ctx context.Context, transactionHash string) (blockNumber uint64, blockHash string, succ bool, err error)
}
