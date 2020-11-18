package ethereum

import (
	"context"
	ethCommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/pkg/errors"
	"math/big"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type Client struct {
	ctx      context.Context
	ethCli   *ethclient.Client
	accounts map[ethCommon.Address]*Account //accounts for sign transaction
	aliases  []string
	mu       sync.RWMutex
}

func NewClient(ctx context.Context, provider string) (*Client, error) {
	ethCli, err := ethclient.Dial(provider)
	if err != nil {
		return nil, err
	}

	return &Client{
		ctx:      ctx,
		ethCli:   ethCli,
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

func (c *Client) AddAccount(pk string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	private, err := HexToPrivate(pk)
	if err != nil {
		return err
	}
	account := PrivateToAccount(private)
	c.accounts[account.Address()] = account
	c.aliases = append(c.aliases, account.String())
	return nil
}

func (c *Client) GetAccount(account string) (*Account, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	address := ethCommon.HexToAddress(account)
	account, ok := accounts[address]
	if !ok {
		return nil, errors.New("account not exists")
	}
	return account, nil
}

func (c *Client) HeaderByNumber(ctx context.Context, number *big.Int) (*model.BlockHeader, error) {
	header, err := c.ethCli.HeaderByNumber(ctx, number)
	if err != nil {
		return nil, err
	}
	block := &model.BlockHeader{
		BlockNumber: int(header.Number.Int64()),
		BlockHash:   strings.ToLower(header.Hash().Hex()),
		ParentHash:  strings.ToLower(header.ParentHash.Hex()),
		BlockTime:   time.Unix(int64(header.Time), 0).UTC(),
	}
	return block, nil
}

func (c *Client) HeaderByHash(ctx context.Context, hash string) (*model.BlockHeader, error) {
	header, err := c.ethCli.HeaderByHash(ctx, ethCommon.HexToHash(hash))
	if err != nil {
		return nil, err
	}
	block := &model.BlockHeader{
		BlockNumber: header.Number.Int64(),
		BlockHash:   strings.ToLower(header.Hash().Hex()),
		ParentHash:  strings.ToLower(header.ParentHash.Hex()),
		BlockTime:   time.Unix(int64(header.Time), 0).UTC(),
	}
	return block, nil
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
		ethcommon.HexToAddress(tx.ToAddress),
		utils.MustDecimalToBigInt(tx.Value),
		*tx.GasLimit,
		big.NewInt(int64(*tx.GasPrice)),
		tx.Inputs)
	chainID, err := c.ethCli.NetworkID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get network id failed")
	}
	acc, err := c.GetAccount(tx.FromAddress)
	if err != nil {
		return nil, errors.Wrap(err, "account not availables")
	}
	signedTx, err := acc.Signer().Signer(ethtypes.NewEIP155Signer(chainID), acc.Address(), rawTx)
	if err != nil {
		return nil, errors.Wrap(err, "sign transaction failed")
	}
	if err = c.ethCli.SendTransaction(ctx, signedTx); err != nil {
		return errors.Wrap(err, "send transaction failed")
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
