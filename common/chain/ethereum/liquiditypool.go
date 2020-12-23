package ethereum

import (
	"context"
	"fmt"
	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/shopspring/decimal"
	"math/big"

	"github.com/mcarloai/mai-v3-broker/common/chain/ethereum/liquiditypool"
	"github.com/mcarloai/mai-v3-broker/common/mai3"
	"github.com/mcarloai/mai-v3-broker/common/model"
)

func (c *Client) GetMarginAccount(ctx context.Context, perpetualIndex int64, poolAddress, trader string) (*model.AccountStorage, error) {
	var opts *ethBind.CallOpts

	address, err := HexToAddress(poolAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid liquidity pool address:%w", err)
	}

	traderAddress, err := HexToAddress(trader)
	if err != nil {
		return nil, fmt.Errorf("invalid trader address:%w", err)
	}

	contract, err := liquiditypool.NewLiquidityPool(address, c.ethCli)
	if err != nil {
		return nil, fmt.Errorf("init liquidity pool contract failed:%w", err)
	}

	res, err := contract.MarginAccount(opts, big.NewInt(perpetualIndex), traderAddress)
	if err != nil {
		return nil, fmt.Errorf("get margin account failed:%w", err)
	}

	rsp := &model.AccountStorage{}
	rsp.CashBalance = decimal.NewFromBigInt(res.CashBalance, -mai3.DECIMALS)
	rsp.PositionAmount = decimal.NewFromBigInt(res.PositionAmount, -mai3.DECIMALS)
	return rsp, nil
}
