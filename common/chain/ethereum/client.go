package ethereum

import (
	"context"
	"crypto/ecdsa"
	ethCommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/pkg/errors"
	"math/big"
	"math/rand"
	"sync"
)

type Client struct {
	ctx      context.Context
	ethCli   *ethclient.Client
	accounts map[ethCommon.Address]*Account //accounts for sign transaction
	aliases  []string
	mu       sync.RWMutex
}

func NewClient(ctx context.Context, provider string, headers map[string]string) (*Client, error) {
	rpcClient, err := rpc.Dial(provider)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		rpcClient.SetHeader(key, value)
	}

	return &Client{
		ctx:      ctx,
		ethCli:   ethclient.NewClient(rpcClient),
		accounts: make(map[ethCommon.Address]*Account),
	}, nil
}

func (c *Client) GetSignAccount() (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	accLen := len(c.aliases)
	if accLen == 0 {
		return "", errors.New("no account added")
	}

	idx := rand.Intn(accLen)
	return c.aliases[idx], nil
}

func (c *Client) AddAccount(private *ecdsa.PrivateKey) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	account := PrivateToAccount(private)
	if _, ok := c.accounts[account.Address()]; ok {
		return errors.New("account duplicated")
	}
	c.accounts[account.Address()] = account
	c.aliases = append(c.aliases, account.String())
	return nil
}

func (c *Client) GetAccount(account string) (*Account, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	address := ethCommon.HexToAddress(account)
	acc, ok := c.accounts[address]
	if !ok {
		return nil, errors.New("account not exists")
	}
	return acc, nil
}

func (c *Client) GetChainID(ctx context.Context) (*big.Int, error) {
	return c.ethCli.ChainID(ctx)
}

func (c *Client) GetLatestBlockNumber(ctx context.Context) (uint64, error) {
	block, err := c.ethCli.BlockByNumber(ctx, nil)
	if err != nil {
		return 0, errors.Wrap(err, "fail to get latest block number")
	}

	return block.NumberU64(), nil
}

func (c *Client) PendingNonceAt(ctx context.Context, account string) (uint64, error) {
	address := ethCommon.HexToAddress(account)
	return c.ethCli.PendingNonceAt(ctx, address)
}

func (c *Client) TransactionByHash(ctx context.Context, txHash string) (bool, error) {
	transactionHash := ethCommon.HexToHash(txHash)
	_, isPending, err := c.ethCli.TransactionByHash(ctx, transactionHash)
	return isPending, err
}

func (c *Client) SendTransaction(ctx context.Context, tx *model.LaunchTransaction) (string, error) {
	rawTx := ethtypes.NewTransaction(
		*tx.Nonce,
		ethCommon.HexToAddress(tx.ToAddress),
		utils.MustDecimalToBigInt(tx.Value),
		*tx.GasLimit,
		big.NewInt(int64(*tx.GasPrice)),
		tx.Inputs)
	acc, err := c.GetAccount(tx.FromAddress)
	if err != nil {
		return "", errors.Wrap(err, "account not availables")
	}
	signedTx, err := acc.Signer().Signer(acc.Address(), rawTx)
	if err != nil {
		return "", errors.Wrap(err, "sign transaction failed")
	}
	if err = c.ethCli.SendTransaction(ctx, signedTx); err != nil {
		return "", errors.Wrap(err, "send transaction failed")
	}
	return signedTx.Hash().Hex(), nil
}

func (c *Client) WaitTransactionReceipt(ctx context.Context, txHash string) (*model.Receipt, error) {
	rcpt, err := c.ethCli.TransactionReceipt(ctx, ethCommon.HexToHash(txHash))
	if err != nil {
		return nil, err
	}

	if rcpt.BlockNumber == nil {
		return nil, errors.New("empty block number")
	}

	block, err := c.ethCli.BlockByNumber(ctx, rcpt.BlockNumber)
	if err != nil {
		return nil, errors.Wrap(err, "fail to retrieve block for transaction")
	}

	receipt := &model.Receipt{
		BlockNumber: rcpt.BlockNumber.Uint64(),
		BlockHash:   rcpt.BlockHash.Hex(),
		GasUsed:     rcpt.GasUsed,
		BlockTime:   block.Time(),
	}
	if rcpt.Status == ethtypes.ReceiptStatusSuccessful {
		receipt.Status = model.TxSuccess
	} else {
		receipt.Status = model.TxFailed
	}
	return receipt, nil
}
