package ethereum

import (
	"context"
	"fmt"
	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/mcarloai/mai-v3-broker/common/model"
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
			FactoryAddress:    strings.ToLower(iter.Event.Raw.Address.Hex()),
			TransactionSeq:    int(iter.Event.Raw.TxIndex),
			TransactionHash:   strings.ToLower(iter.Event.Raw.TxHash.Hex()),
			BlockNumber:       int64(iter.Event.Raw.BlockNumber),
			PerpetualAddress:  strings.ToLower(iter.Event.Perpetual.Hex()),
			OperatorAddress:   strings.ToLower(iter.Event.Operator.Hex()),
			OracleAddress:     strings.ToLower(iter.Event.Oracle.Hex()),
			CollateralAddress: strings.ToLower(iter.Event.Collateral.Hex()),
		}

		rsp = append(rsp, perpetual)
	}

	return rsp, nil
}
