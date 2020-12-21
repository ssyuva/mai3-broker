package ethereum

import (
	"context"
	"fmt"
	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/shopspring/decimal"

	"github.com/mcarloai/mai-v3-broker/common/chain/ethereum/perpetual"
	"github.com/mcarloai/mai-v3-broker/common/mai3"
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

	res, err := contract.MarginAccount(opts, accountAddress)
	if err != nil {
		return nil, fmt.Errorf("get margin account failed:%w", err)
	}

	rsp := &model.AccountStorage{}
	rsp.CashBalance = decimal.NewFromBigInt(res.CashBalance, -mai3.DECIMALS)
	rsp.PositionAmount = decimal.NewFromBigInt(res.PositionAmount, -mai3.DECIMALS)
	rsp.EntryFundingLoss = decimal.NewFromBigInt(res.EntryFundingLoss, -mai3.DECIMALS)
	return rsp, nil
}

func (c *Client) GetPerpetualGovParams(ctx context.Context, perpetualAddress string) (*model.GovParams, error) {
	var opts *ethBind.CallOpts

	address, err := HexToAddress(perpetualAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid perpetual address:%w", err)
	}

	contract, err := perpetual.NewPerpetual(address, c.ethCli)
	if err != nil {
		return nil, fmt.Errorf("init perpetual contract failed:%w", err)
	}

	res, err := contract.Information(opts)
	if err != nil {
		return nil, fmt.Errorf("get perpetual Information failed:%w", err)
	}
	rsp := &model.GovParams{
		InitialMarginRate:     decimal.NewFromBigInt(res.CoreParameter[0], -mai3.DECIMALS),
		MaintenanceMarginRate: decimal.NewFromBigInt(res.CoreParameter[1], -mai3.DECIMALS),
		OperatorFeeRate:       decimal.NewFromBigInt(res.CoreParameter[2], -mai3.DECIMALS),
		VaultFeeRate:          decimal.NewFromBigInt(res.CoreParameter[3], -mai3.DECIMALS),
		LpFeeRate:             decimal.NewFromBigInt(res.CoreParameter[4], -mai3.DECIMALS),
		ReferrerRebateRate:    decimal.NewFromBigInt(res.CoreParameter[5], -mai3.DECIMALS),
		LiquidatorPenaltyRate: decimal.NewFromBigInt(res.CoreParameter[6], -mai3.DECIMALS),
		KeeperGasReward:       decimal.NewFromBigInt(res.CoreParameter[7], -mai3.DECIMALS),
		// amm
		HalfSpreadRate:         decimal.NewFromBigInt(res.RiskParameter[0], -mai3.DECIMALS),
		Beta1:                  decimal.NewFromBigInt(res.RiskParameter[1], -mai3.DECIMALS),
		Beta2:                  decimal.NewFromBigInt(res.RiskParameter[2], -mai3.DECIMALS),
		FundingRateCoefficient: decimal.NewFromBigInt(res.RiskParameter[3], -mai3.DECIMALS),
		TargetLeverage:         decimal.NewFromBigInt(res.RiskParameter[4], -mai3.DECIMALS),
	}
	return rsp, nil
}

func (c *Client) GetPerpetualStorage(ctx context.Context, perpetualAddress string) (*model.PerpetualStorage, error) {
	var opts *ethBind.CallOpts

	address, err := HexToAddress(perpetualAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid perpetual address:%w", err)
	}

	contract, err := perpetual.NewPerpetual(address, c.ethCli)
	if err != nil {
		return nil, fmt.Errorf("init perpetual contract failed:%w", err)
	}

	res, err := contract.State(opts)
	if err != nil {
		return nil, fmt.Errorf("get perpetual state failed:%w", err)
	}

	fundingState, err := contract.FundingState(opts)
	if err != nil {
		return nil, fmt.Errorf("get perpetual FundingState failed:%w", err)
	}

	rsp := &model.PerpetualStorage{}
	rsp.IsEmergency = res.IsEmergency
	rsp.IsGlobalSettled = res.IsShuttingdown
	rsp.InsuranceFund1 = decimal.NewFromBigInt(res.InsuranceFund, -mai3.DECIMALS)
	rsp.InsuranceFund2 = decimal.NewFromBigInt(res.DonatedInsuranceFund, -mai3.DECIMALS)
	rsp.MarkPrice = decimal.NewFromBigInt(res.MarkPrice, -mai3.DECIMALS)
	rsp.IndexPrice = decimal.NewFromBigInt(res.IndexPrice, -mai3.DECIMALS)
	rsp.AccumulatedFundingPerContract = decimal.NewFromBigInt(fundingState.UnitAccumulativeFunding, -mai3.DECIMALS)
	rsp.FundingTime = fundingState.FundingTime.Int64()
	return rsp, nil
}
