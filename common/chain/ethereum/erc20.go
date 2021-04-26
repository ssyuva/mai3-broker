package ethereum

import (
	"context"
	"fmt"

	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/mcarloai/mai-v3-broker/common/chain/ethereum/erc20"

	"github.com/shopspring/decimal"
)

func (c *Client) BalanceOf(ctx context.Context, token, owner string, decimals int32) (decimal.Decimal, error) {
	opts := &ethBind.CallOpts{
		Context: ctx,
	}

	tokenAddr, err := HexToAddress(token)
	if err != nil {
		return decimal.Zero, fmt.Errorf("invalid token address:%s err:%w", token, err)
	}

	ownerAddr, err := HexToAddress(owner)
	if err != nil {
		return decimal.Zero, fmt.Errorf("invalid owner address:%s err:%w", owner, err)
	}

	contract, err := erc20.NewERC20(tokenAddr, c.GetEthClient())
	if err != nil {
		return decimal.Zero, fmt.Errorf("init reader contract failed:%w", err)
	}

	res, err := contract.BalanceOf(opts, ownerAddr)
	if err != nil {
		return decimal.Zero, fmt.Errorf("get balance failed:%w", err)
	}

	rsp := decimal.NewFromBigInt(res, -decimals)
	return rsp, nil
}

func (c *Client) Allowance(ctx context.Context, token, owner, spender string, decimals int32) (decimal.Decimal, error) {
	opts := &ethBind.CallOpts{
		Context: ctx,
	}

	tokenAddr, err := HexToAddress(token)
	if err != nil {
		return decimal.Zero, fmt.Errorf("invalid token address:%s err:%w", token, err)
	}

	ownerAddr, err := HexToAddress(owner)
	if err != nil {
		return decimal.Zero, fmt.Errorf("invalid owner address:%s err:%w", owner, err)
	}

	spenderAddr, err := HexToAddress(spender)
	if err != nil {
		return decimal.Zero, fmt.Errorf("invalid spender address:%s err:%w", spender, err)
	}

	contract, err := erc20.NewERC20(tokenAddr, c.GetEthClient())
	if err != nil {
		return decimal.Zero, fmt.Errorf("init reader contract failed:%w", err)
	}

	res, err := contract.Allowance(opts, ownerAddr, spenderAddr)
	if err != nil {
		return decimal.Zero, fmt.Errorf("get allowance failed:%w", err)
	}

	rsp := decimal.NewFromBigInt(res, -decimals)
	return rsp, nil
}
