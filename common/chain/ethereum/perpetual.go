package ethereum

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethCommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"math/big"
	"strings"

	"github.com/mcarloai/mai-v3-broker/common/chain/ethereum/perpetual"
	"github.com/mcarloai/mai-v3-broker/common/mai3"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
)

func (c *Client) GetMarginAccount(ctx context.Context, perpetualAddress, account string) (*model.AccountStorage, error) {
	var opts *ethBind.CallOpts

	address, err := HexToAddress(perpetualAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid perpetual address:%w", err)
	}

	accountAddress, err := HexToAddress(account)
	if err != nil {
		return nil, fmt.Errorf("invalid account address:%w", err)
	}

	contract, err := perpetual.NewPerpetual(address, c.ethCli)
	if err != nil {
		return nil, fmt.Errorf("init perpetual contract failed:%w", err)
	}

	cash, pos, funding, err := contract.MarginAccount(opts, accountAddress)
	if err != nil {
		return nil, fmt.Errorf("get margin account failed:%w", err)
	}

	rsp := &model.AccountStorage{}
	rsp.CashBalance = decimal.NewFromBigInt(cash, -mai3.DECIMALS)
	rsp.Position = decimal.NewFromBigInt(pos, -mai3.DECIMALS)
	// rsp.EntrySocialLoss = decimal.NewFromBigInt(storage.EntrySocialLoss, -mai3.DECIMALS)
	rsp.EntryFundingLoss = decimal.NewFromBigInt(funding, -mai3.DECIMALS)
	return rsp, nil
}

// TODO move to oracle contract
func (c *Client) GetPrice(ctx context.Context, oracle string) (decimal.Decimal, error) {
	var opts *ethBind.CallOpts
	var res decimal.Decimal

	oracleAddress, err := HexToAddress(oracle)
	if err != nil {
		return res, fmt.Errorf("invalid oracle address:%w", err)
	}

	contract, err := perpetual.NewPerpetual(oracleAddress, c.ethCli)
	if err != nil {
		return res, fmt.Errorf("init oracle contract failed:%w", err)
	}

	_, _, fast, _, err := contract.Price(opts)
	if err != nil {
		return res, fmt.Errorf("get oracle price failed:%w", err)
	}

	return decimal.NewFromBigInt(fast, -mai3.DECIMALS), nil
}