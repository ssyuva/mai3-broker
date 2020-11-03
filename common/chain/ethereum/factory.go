package ethereum

 import (
 	"context"
// 	"fmt"
// 	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	gethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/mcarloai/mai-v3-broker/common/model"

// 	"github.com/mcarloai/mai-v3-broker/common/chain/ethereum/abis/factory"
 )

 func (c *Client) FilterCreatePerpetual(ctx context.Context, factoryAddress string, start, end uint64) ([]*model.PerpetualEvent, error) {
// 	opts := &ethBind.FilterOpts{
// 		Start:   start,
// 		End:     &end,
// 		context: ctx,
// 	}

	rsp := make([]*model.PerpetualEvent, 0)

// 	addresss, err := HexToAddress(factoryAddress)
// 	if err != nil {
// 		return rsp, fmt.Errorf("invalid factory address:%w", err)
// 	}

// 	contract, err := factory.NewFactory(address, c.ethCli)
// 	if err != nil {
// 		return rsp, fmt.Errorf("init factory contract failed:%w", err)
// 	}

// 	iter, err := contract.FilterCreatePerpetual(opts, []gethCommon.Address{})
// 	if err != nil {
// 		return rsp, fmt.Errorf("filter create perpetual failed:%w", err)
// 	}

// 	if iter.Next() {
// 		perpetual := &model.PerpetualEvent{
// 			FactoryAddress:   strings.ToLower(iter.Event.Raw.Address.Hex()),
// 			TransactionSeq:   int(iter.Event.Raw.TxIndex),
// 			TransactionHash:  strings.ToLower(iter.Event.Raw.TxHash.Hex()),
// 			BlockNumber:      int(iter.Event.Raw.BlockNumber),
// 			PerpetualAddress: strings.ToLower(iter.Event.Proxy.Hex()),
// 			OperatorAddress:  strings.ToLower(iter.Event.Operator.Hex()),
// 			OracleAddress:    strings.ToLower(iter.Event.Oracle.Hex()),
// 		}

// 		rsp := append(rsp, perpetual)
// 	}

 	return rsp, nil
}
