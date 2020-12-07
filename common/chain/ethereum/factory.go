package ethereum

import (
	"context"
	"fmt"
	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/mcarloai/mai-v3-broker/common/mai3"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/shopspring/decimal"
	"strings"

	"github.com/mcarloai/mai-v3-broker/common/chain/ethereum/factory"
)

func (c *Client) FilterCreatePerpetual(ctx context.Context, factoryAddress string, start, end uint64) ([]*model.PerpetualEvent, error) {
	opts := &ethBind.FilterOpts{
		Start:   start,
		End:     &end,
		Context: ctx,
	}

	rsp := make([]*model.PerpetualEvent, 0)

	address, err := HexToAddress(factoryAddress)
	if err != nil {
		return rsp, fmt.Errorf("invalid factory address:%w", err)
	}

	contract, err := factory.NewFactory(address, c.ethCli)
	if err != nil {
		return rsp, fmt.Errorf("init factory contract failed:%w", err)
	}

	iter, err := contract.FilterCreatePerpetual(opts)
	if err != nil {
		return rsp, fmt.Errorf("filter create perpetual failed:%w", err)
	}

	if iter.Next() {
		perpetual := &model.PerpetualEvent{
			FactoryAddress:         strings.ToLower(iter.Event.Raw.Address.Hex()),
			TransactionSeq:         int(iter.Event.Raw.TxIndex),
			TransactionHash:        strings.ToLower(iter.Event.Raw.TxHash.Hex()),
			BlockNumber:            int64(iter.Event.Raw.BlockNumber),
			PerpetualAddress:       strings.ToLower(iter.Event.Perpetual.Hex()),
			GovernorAddress:        strings.ToLower(iter.Event.Governor.Hex()),
			ShareToken:             strings.ToLower(iter.Event.ShareToken.Hex()),
			OperatorAddress:        strings.ToLower(iter.Event.Operator.Hex()),
			OracleAddress:          strings.ToLower(iter.Event.Oracle.Hex()),
			CollateralAddress:      strings.ToLower(iter.Event.Collateral.Hex()),
			InitialMarginRate:      decimal.NewFromBigInt(iter.Event.CoreParams[0], -mai3.DECIMALS),
			MaintenanceMarginRate:  decimal.NewFromBigInt(iter.Event.CoreParams[1], -mai3.DECIMALS),
			OperatorFeeRate:        decimal.NewFromBigInt(iter.Event.CoreParams[2], -mai3.DECIMALS),
			LpFeeRate:              decimal.NewFromBigInt(iter.Event.CoreParams[3], -mai3.DECIMALS),
			ReferrerRebateRate:     decimal.NewFromBigInt(iter.Event.CoreParams[4], -mai3.DECIMALS),
			LiquidatorPenaltyRate:  decimal.NewFromBigInt(iter.Event.CoreParams[5], -mai3.DECIMALS),
			KeeperGasReward:        decimal.NewFromBigInt(iter.Event.CoreParams[6], -mai3.DECIMALS),
			HalfSpreadRate:         decimal.NewFromBigInt(iter.Event.RiskParams[0], -mai3.DECIMALS),
			Beta1:                  decimal.NewFromBigInt(iter.Event.RiskParams[1], -mai3.DECIMALS),
			Beta2:                  decimal.NewFromBigInt(iter.Event.RiskParams[2], -mai3.DECIMALS),
			FundingRateCoefficient: decimal.NewFromBigInt(iter.Event.RiskParams[3], -mai3.DECIMALS),
			TargetLeverage:         decimal.NewFromBigInt(iter.Event.RiskParams[4], -mai3.DECIMALS),
		}

		rsp = append(rsp, perpetual)
	}

	return rsp, nil
}
