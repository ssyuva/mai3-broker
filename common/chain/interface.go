package chain

import (
	"context"
	"crypto/ecdsa"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/shopspring/decimal"
	"math/big"
)

type ChainClient interface {
	// account
	EncryptKey(pk *ecdsa.PrivateKey, password string) ([]byte, error)
	DecryptKey(cipher []byte, password string) (*ecdsa.PrivateKey, error)
	HexToPrivate(pk string) (*ecdsa.PrivateKey, string, error)
	AddAccount(private *ecdsa.PrivateKey) error
	GetSignAccount() (string, error)

	GetGasBalance(ctx context.Context, brokerAddress string, address string) (decimal.Decimal, error)
	GetChainID(ctx context.Context) (*big.Int, error)
	GetLatestBlockNumber(ctx context.Context) (uint64, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*model.BlockHeader, error)
	HeaderByHash(ctx context.Context, hash string) (*model.BlockHeader, error)
	PendingNonceAt(ctx context.Context, account string) (uint64, error)
	TransactionByHash(ctx context.Context, txHash string) (bool, error)
	SendTransaction(ctx context.Context, tx *model.LaunchTransaction) (string, error)
	WaitTransactionReceipt(ctx context.Context, txHash string) (*model.Receipt, error)
	GetMarginAccount(ctx context.Context, poolAddress, account string) (*model.AccountStorage, error)
	GetPerpetualStorage(ctx context.Context, poolAddress string) (*model.PerpetualStorage, error)
	FilterTradeSuccess(ctx context.Context, poolAddress string, start, end uint64) ([]*model.TradeSuccessEvent, error)
	BatchTradeDataPack(orderParams []*model.WalletOrderParam, matchAmounts []decimal.Decimal, gasRewards []*big.Int) ([]byte, error)
	GetPerpetualGovParams(ctx context.Context, poolAddress string) (*model.GovParams, error)
}
