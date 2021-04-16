package chain

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/shopspring/decimal"
)

type ChainClient interface {
	// account
	EncryptKey(pk *ecdsa.PrivateKey, password string) ([]byte, error)
	DecryptKey(cipher []byte, password string) (*ecdsa.PrivateKey, error)
	HexToPrivate(pk string) (*ecdsa.PrivateKey, string, error)
	AddAccount(private *ecdsa.PrivateKey) error
	GetSignAccount() (string, error)
	GetRelayerAccounts() []string

	IsNotFoundError(err error) bool
	GetGasBalance(ctx context.Context, brokerAddress string, address string) (decimal.Decimal, error)
	GetChainID() (*big.Int, error)
	GetLatestBlockNumber() (uint64, error)
	PendingNonceAt(account string) (uint64, error)
	TransactionByHash(txHash string) (bool, error)
	SendTransaction(tx *model.LaunchTransaction) (string, error)
	WaitTransactionReceipt(txHash string) (*model.Receipt, error)
	GetAccountStorage(ctx context.Context, readerAddress string, perpetualIndex int64, poolAddress, account string) (*model.AccountStorage, error)
	GetLiquidityPoolStorage(ctx context.Context, readerAddress, poolAddress string) (*model.LiquidityPoolStorage, error)
	FilterTradeSuccess(ctx context.Context, poolAddress string, start, end uint64) ([]*model.TradeSuccessEvent, error)
	FilterTradeFailed(ctx context.Context, poolAddress string, start, end uint64) ([]*model.TradeFailedEvent, error)
	BatchTradeDataPack(compressedOrders [][]byte, matchAmounts []decimal.Decimal, gasRewards []*big.Int) ([]byte, error)
}
