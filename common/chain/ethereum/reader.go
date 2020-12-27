package ethereum

import (
	"context"
	"fmt"
	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/shopspring/decimal"
	"math/big"

	"github.com/mcarloai/mai-v3-broker/common/chain/ethereum/reader"
	"github.com/mcarloai/mai-v3-broker/common/mai3"
	"github.com/mcarloai/mai-v3-broker/common/model"
)

func (c *Client) GetAccountStorage(ctx context.Context, readerAddress string, perpetualIndex int64, poolAddress, trader string) (*model.AccountStorage, error) {
	var opts *ethBind.CallOpts

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

	contract, err := reader.NewReader(address, c.ethCli)
	if err != nil {
		return nil, fmt.Errorf("init reader contract failed:%w", err)
	}

	res, err := contract.GetAccountStorage(opts, pool, big.NewInt(perpetualIndex), traderAddress)
	if err != nil {
		return nil, fmt.Errorf("get margin account failed:%w", err)
	}

	rsp := &model.AccountStorage{}
	rsp.CashBalance = decimal.NewFromBigInt(res.CashBalance, -mai3.DECIMALS)
	rsp.PositionAmount = decimal.NewFromBigInt(res.PositionAmount, -mai3.DECIMALS)
	return rsp, nil
}

func (c *Client) GetLiquidityPoolStorage(ctx context.Context, readerAddress, poolAddress string) (*model.LiquidityPoolStorage, error) {
	var opts *ethBind.CallOpts

	address, err := HexToAddress(readerAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid reader address:%w", err)
	}

	liquidityPool, err := HexToAddress(poolAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid liquidity pool address:%w", err)
	}

	contract, err := reader.NewReader(address, c.ethCli)
	if err != nil {
		return nil, fmt.Errorf("init reader contract failed:%w", err)
	}

	res, err := contract.GetLiquidityPoolStorage(opts, liquidityPool)
	if err != nil {
		return nil, fmt.Errorf("GetLiquidityPoolStorage failed:%w", err)
	}
	rsp := &model.LiquidityPoolStorage{}
	rsp.VaultFeeRate = decimal.NewFromBigInt(res.VaultFeeRate, -mai3.DECIMALS)
	rsp.InsuranceFundCap = decimal.NewFromBigInt(res.InsuranceFundCap, -mai3.DECIMALS)
	rsp.InsuranceFund = decimal.NewFromBigInt(res.InsuranceFund, -mai3.DECIMALS)
	rsp.DonatedInsuranceFund = decimal.NewFromBigInt(res.DonatedInsuranceFund, -mai3.DECIMALS)
	rsp.TotalClaimableFee = decimal.NewFromBigInt(res.TotalClaimableFee, -mai3.DECIMALS)
	rsp.PoolCashBalance = decimal.NewFromBigInt(res.PoolCashBalance, -mai3.DECIMALS)
	rsp.FundingTime = res.FundingTime.Int64()
	rsp.Perpetuals = make(map[int64]*model.PerpetualStorage)

	for i, perpetual := range res.PerpetualStorages {
		storage := &model.PerpetualStorage{
			MarkPrice:               decimal.NewFromBigInt(perpetual.MarkPrice, -mai3.DECIMALS),
			IndexPrice:              decimal.NewFromBigInt(perpetual.IndexPrice, -mai3.DECIMALS),
			UnitAccumulativeFunding: decimal.NewFromBigInt(perpetual.UnitAccumulativeFunding, -mai3.DECIMALS),
			InitialMarginRate:       decimal.NewFromBigInt(perpetual.InitialMarginRate, -mai3.DECIMALS),
			MaintenanceMarginRate:   decimal.NewFromBigInt(perpetual.MaintenanceMarginRate, -mai3.DECIMALS),
			OperatorFeeRate:         decimal.NewFromBigInt(perpetual.OperatorFeeRate, -mai3.DECIMALS),
			ReferrerRebateRate:      decimal.NewFromBigInt(perpetual.ReferrerRebateRate, -mai3.DECIMALS),
			LiquidationPenaltyRate:  decimal.NewFromBigInt(perpetual.LiquidationPenaltyRate, -mai3.DECIMALS),
			KeeperGasReward:         decimal.NewFromBigInt(perpetual.KeeperGasReward, -mai3.DECIMALS),
			InsuranceFundRate:       decimal.NewFromBigInt(perpetual.InsuranceFundRate, -mai3.DECIMALS),
			HalfSpread:              decimal.NewFromBigInt(perpetual.HalfSpread, -mai3.DECIMALS),
			OpenSlippageFactor:      decimal.NewFromBigInt(perpetual.OpenSlippageFactor, -mai3.DECIMALS),
			CloseSlippageFactor:     decimal.NewFromBigInt(perpetual.CloseSlippageFactor, -mai3.DECIMALS),
			FundingRateLimit:        decimal.NewFromBigInt(perpetual.FundingRateLimit, -mai3.DECIMALS),
			MaxLeverage:             decimal.NewFromBigInt(perpetual.MaxLeverage, -mai3.DECIMALS),
			AmmPositionAmount:       decimal.NewFromBigInt(perpetual.AmmPositionAmount, -mai3.DECIMALS),
		}
		rsp.Perpetuals[int64(i)] = storage
	}

	return rsp, nil
}
