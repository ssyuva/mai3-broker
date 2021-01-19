package ethereum

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/shopspring/decimal"
	"math/big"
	"strings"

	"github.com/mcarloai/mai-v3-broker/common/chain/ethereum/broker"
	"github.com/mcarloai/mai-v3-broker/common/mai3"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
)

func (c *Client) BatchTradeDataPack(compressedOrders [][]byte, matchAmounts []decimal.Decimal, gasRewards []*big.Int) ([]byte, error) {
	parsed, err := abi.JSON(strings.NewReader(broker.BrokerABI))
	if err != nil {
		return nil, err
	}
	amounts := make([]*big.Int, 0)
	for _, amount := range matchAmounts {
		amounts = append(amounts, utils.MustDecimalToBigInt(utils.ToWad(amount)))
	}
	inputs, err := parsed.Pack("batchTrade", compressedOrders, amounts, gasRewards)
	return inputs, err
}

func (c *Client) FilterTradeSuccess(ctx context.Context, brokerAddress string, start, end uint64) ([]*model.TradeSuccessEvent, error) {
	opts := &ethBind.FilterOpts{
		Start:   start,
		End:     &end,
		Context: ctx,
	}

	rsp := make([]*model.TradeSuccessEvent, 0)

	address, err := HexToAddress(brokerAddress)
	if err != nil {
		return rsp, fmt.Errorf("invalid broker address:%w", err)
	}

	contract, err := broker.NewBroker(address, c.ethCli)
	if err != nil {
		return rsp, fmt.Errorf("init broker contract failed:%w", err)
	}

	iter, err := contract.FilterTradeSuccess(opts)
	if err != nil {
		return rsp, fmt.Errorf("filter trade event failed:%w", err)
	}

	if iter.Next() {
		match := &model.TradeSuccessEvent{
			PerpetualAddress: strings.ToLower(iter.Event.Raw.Address.Hex()),
			TransactionSeq:   int(iter.Event.Raw.TxIndex),
			TransactionHash:  strings.ToLower(iter.Event.Raw.TxHash.Hex()),
			BlockNumber:      int64(iter.Event.Raw.BlockNumber),
			TraderAddress:    strings.ToLower(iter.Event.Order.Trader.Hex()),
			OrderHash:        utils.Bytes2HexP(iter.Event.OrderHash[:]),
			Amount:           decimal.NewFromBigInt(iter.Event.Amount, -mai3.DECIMALS),
			Gas:              decimal.NewFromBigInt(iter.Event.GasReward, -mai3.DECIMALS),
		}

		rsp = append(rsp, match)
	}

	return rsp, nil
}

func (c *Client) GetGasBalance(ctx context.Context, brokerAddress string, address string) (decimal.Decimal, error) {
	opts := &ethBind.CallOpts{
		Context: ctx,
	}

	var rsp decimal.Decimal

	account, err := HexToAddress(address)
	if err != nil {
		return rsp, fmt.Errorf("invalid user address:%w", err)
	}

	brokerAddr, err := HexToAddress(brokerAddress)
	if err != nil {
		return rsp, fmt.Errorf("invalid broker address:%w", err)
	}

	contract, err := broker.NewBroker(brokerAddr, c.ethCli)
	if err != nil {
		return rsp, fmt.Errorf("init broker contract failed:%w", err)
	}

	b, err := contract.BalanceOf(opts, account)
	if err != nil {
		return rsp, fmt.Errorf("read broker deposit gas balance failed:%w", err)
	}

	rsp = decimal.NewFromBigInt(b, -mai3.DECIMALS)
	return rsp, nil
}
