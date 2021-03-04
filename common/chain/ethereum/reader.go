package ethereum

import (
	"context"
	"fmt"
	"math/big"

	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/shopspring/decimal"

	"github.com/mcarloai/mai-v3-broker/common/chain/ethereum/reader"
	"github.com/mcarloai/mai-v3-broker/common/mai3"
	"github.com/mcarloai/mai-v3-broker/common/model"
)

func (c *Client) GetAccountStorage(ctx context.Context, readerAddress string, perpetualIndex int64, poolAddress, trader string) (*model.AccountStorage, error) {
	opts := &ethBind.CallOpts{
		Context: ctx,
	}

	address, err := HexToAddress(readerAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid reader address:%w", err)
	}
	pool, err := HexToAddress(poolAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid liquidity pool address:%w", err)
	}

	traderAddress, err := HexToAddress(trader)
	if err != nil {
		return nil, fmt.Errorf("invalid trader address:%w", err)
	}

	contract, err := reader.NewReader(address, c.GetEthClient())
	if err != nil {
		return nil, fmt.Errorf("init reader contract failed:%w", err)
	}

	res, err := contract.GetAccountStorage(opts, pool, big.NewInt(perpetualIndex), traderAddress)
	if err != nil {
		return nil, fmt.Errorf("get margin account failed:%w", err)
	}

	if !res.IsSynced {
		return nil, fmt.Errorf("perpetual sync error")
	}

	rsp := &model.AccountStorage{}
	rsp.CashBalance = decimal.NewFromBigInt(res.AccountStorage.Cash, -mai3.DECIMALS)
	rsp.PositionAmount = decimal.NewFromBigInt(res.AccountStorage.Position, -mai3.DECIMALS)
	return rsp, nil
}

func (c *Client) GetLiquidityPoolStorage(ctx context.Context, readerAddress, poolAddress string) (*model.LiquidityPoolStorage, error) {
	opts := &ethBind.CallOpts{
		Context: ctx,
	}

	address, err := HexToAddress(readerAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid reader address:%w", err)
	}

	liquidityPool, err := HexToAddress(poolAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid liquidity pool address:%w", err)
	}

	contract, err := reader.NewReader(address, c.GetEthClient())
	if err != nil {
		return nil, fmt.Errorf("init reader contract failed:%w", err)
	}

	res, err := contract.GetLiquidityPoolStorage(opts, liquidityPool)
	if err != nil {
		return nil, fmt.Errorf("GetLiquidityPoolStorage failed:%w", err)
	}

	if !res.IsSynced {
		return nil, fmt.Errorf("perpetual sync error")
	}
	rsp := &model.LiquidityPoolStorage{}
	rsp.VaultFeeRate = decimal.NewFromBigInt(res.Pool.VaultFeeRate, -mai3.DECIMALS)
	rsp.PoolCashBalance = decimal.NewFromBigInt(res.Pool.PoolCash, -mai3.DECIMALS)
	rsp.Perpetuals = make(map[int64]*model.PerpetualStorage)

	for i, perpetual := range res.Pool.Perpetuals {
		storage := &model.PerpetualStorage{
			MarkPrice:               decimal.NewFromBigInt(perpetual.Nums[1], -mai3.DECIMALS),
			IndexPrice:              decimal.NewFromBigInt(perpetual.Nums[2], -mai3.DECIMALS),
			UnitAccumulativeFunding: decimal.NewFromBigInt(perpetual.Nums[4], -mai3.DECIMALS),
			InitialMarginRate:       decimal.NewFromBigInt(perpetual.Nums[5], -mai3.DECIMALS),
			MaintenanceMarginRate:   decimal.NewFromBigInt(perpetual.Nums[6], -mai3.DECIMALS),
			OperatorFeeRate:         decimal.NewFromBigInt(perpetual.Nums[7], -mai3.DECIMALS),
			LpFeeRate:               decimal.NewFromBigInt(perpetual.Nums[8], -mai3.DECIMALS),
			ReferrerRebateRate:      decimal.NewFromBigInt(perpetual.Nums[9], -mai3.DECIMALS),
			LiquidationPenaltyRate:  decimal.NewFromBigInt(perpetual.Nums[10], -mai3.DECIMALS),
			KeeperGasReward:         decimal.NewFromBigInt(perpetual.Nums[11], -mai3.DECIMALS),
			InsuranceFundRate:       decimal.NewFromBigInt(perpetual.Nums[12], -mai3.DECIMALS),
			InsuranceFundCap:        decimal.NewFromBigInt(perpetual.Nums[13], -mai3.DECIMALS),
			InsuranceFund:           decimal.NewFromBigInt(perpetual.Nums[14], -mai3.DECIMALS),
			DonatedInsuranceFund:    decimal.NewFromBigInt(perpetual.Nums[15], -mai3.DECIMALS),
			HalfSpread:              decimal.NewFromBigInt(perpetual.Nums[16], -mai3.DECIMALS),
			OpenSlippageFactor:      decimal.NewFromBigInt(perpetual.Nums[19], -mai3.DECIMALS),
			CloseSlippageFactor:     decimal.NewFromBigInt(perpetual.Nums[22], -mai3.DECIMALS),
			FundingRateLimit:        decimal.NewFromBigInt(perpetual.Nums[25], -mai3.DECIMALS),
			MaxLeverage:             decimal.NewFromBigInt(perpetual.Nums[28], -mai3.DECIMALS),
			MaxClosePriceDiscount:   decimal.NewFromBigInt(perpetual.Nums[31], -mai3.DECIMALS),
			OpenInterest:            decimal.NewFromBigInt(perpetual.Nums[34], -mai3.DECIMALS),
			MaxOpenInterestRate:     decimal.NewFromBigInt(perpetual.Nums[35], -mai3.DECIMALS),
			FundingRateFactor:       decimal.NewFromBigInt(perpetual.Nums[36], -mai3.DECIMALS),
			AmmCashBalance:          decimal.NewFromBigInt(perpetual.AmmCashBalance, -mai3.DECIMALS),
			AmmPositionAmount:       decimal.NewFromBigInt(perpetual.AmmPositionAmount, -mai3.DECIMALS),
		}
		if perpetual.State == model.PerpetualNormal {
			storage.IsNormal = true
		}
		rsp.Perpetuals[int64(i)] = storage
	}

	return rsp, nil
}
