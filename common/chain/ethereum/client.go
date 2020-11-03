package ethereum

import (
	"context"
	gethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"math/big"
	"strings"
	"time"
)

type Client struct {
	ctx    context.Context
	ethCli *ethclient.Client
}

func NewClient(ctx context.Context, provider string) (*Client, error) {
	ethCli, err := ethclient.Dial(provider)
	if err != nil {
		return nil, err
	}

	return &Client{
		ctx:    ctx,
		ethCli: ethCli,
	}, nil
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
	header, err := c.ethCli.HeaderByHash(ctx, gethCommon.HexToHash(hash))
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

func (c *Client) WaitTransactionReceipt(ctx context.Context, txHash string) {

}
