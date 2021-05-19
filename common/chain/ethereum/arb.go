package ethereum

import (
	"context"
	"fmt"

	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/mcarloai/mai-v3-broker/common/chain/ethereum/arboscontracts"
	"github.com/mcarloai/mai-v3-broker/common/mai3"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
)

func (c *Client) GetGasPrice(ctx context.Context, gasAddress string) ([]decimal.Decimal, error) {
	res := make([]decimal.Decimal, 0)
	defer func() {
		if r := recover(); r != nil {
			_, ok := r.(error)
			if !ok {
				err := fmt.Errorf("%v", r)
				logger.Warningf("GetGasPrice failed. err:%s", err)
			}
		}
	}()

	opts := &ethBind.CallOpts{
		Context: ctx,
	}
	address, err := HexToAddress(gasAddress)
	if err != nil {
		return res, fmt.Errorf("invalid gas address:%w", err)
	}
	contract, err := arboscontracts.NewArbGasInfo(address, c.GetEthClient())
	if err != nil {
		return res, fmt.Errorf("init reader contract failed:%w", err)
	}

	out0, out1, out2, out3, out4, out5, err := contract.GetPricesInWei(opts)
	if err != nil {
		return res, fmt.Errorf("get margin account failed:%w", err)
	}

	res = append(res, decimal.NewFromBigInt(out0, -mai3.DECIMALS))
	res = append(res, decimal.NewFromBigInt(out1, -mai3.DECIMALS))
	res = append(res, decimal.NewFromBigInt(out2, -mai3.DECIMALS))
	res = append(res, decimal.NewFromBigInt(out3, -mai3.DECIMALS))
	res = append(res, decimal.NewFromBigInt(out4, -mai3.DECIMALS))
	res = append(res, decimal.NewFromBigInt(out5, -mai3.DECIMALS))
	return res, nil
}
