package ethereum

import (
	"context"
	"fmt"
	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethCommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"strings"

	"github.com/mcarloai/mai-v3-broker/common/chain/ethereum/abis/perpetual"
	"github.com/mcarloai/mai-v3-broker/common/mai3"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
)

var TradeFailureOption = 1 // 0: revert 1: continue for batchTrade

func (c *Client) FilterMatch(ctx context.Context, perpetualAddress string, start, end uint64) ([]*model.MatchEvent, error) {
	opts := &ethBind.FilterOpts{
		Start:   start,
		End:     &end,
		Context: ctx,
	}

	rsp := make([]*model.MatchEvent, 0)

	address, err := HexToAddress(perpetualAddress)
	if err != nil {
		return rsp, fmt.Errorf("invalid perpetual address:%w", err)
	}

	contract, err := perpetual.NewPerpetual(address, c.ethCli)
	if err != nil {
		return rsp, fmt.Errorf("init perpetual contract failed:%w", err)
	}

	iter, err := contract.FilterMatch(opts)
	if err != nil {
		return rsp, fmt.Errorf("filter trade event failed:%w", err)
	}

	if iter.Next() {
		match := &model.MatchEvent{
			PerpetualAddress: strings.ToLower(iter.Event.Raw.Address.Hex()),
			TransactionSeq:   int(iter.Event.Raw.TxIndex),
			TransactionHash:  strings.ToLower(iter.Event.Raw.TxHash.Hex()),
			BlockNumber:      int(iter.Event.Raw.BlockNumber),
			TraderAddress:    strings.ToLower(iter.Event.Arg0.Trader.Hex()),
			Amount:           decimal.NewFromBigInt(iter.Event.Arg1, -mai3.DECIMALS),
			Gas:              decimal.NewFromBigInt(iter.Event.Arg2, -mai3.DECIMALS),
		}

		rsp = append(rsp, match)
	}

	return rsp, nil
}

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

func (c *Client) BatchTradeDataPack(orderParams []*model.WalletOrderParam, matchAmounts []decimal.Decimal, gases []*big.Int) ([]byte, error) {
	parsed, err := abi.JSON(strings.NewReader(perpetual.PerpetualABI))
	if err != nil {
		return nil, err
	}
	orders := make([]perpetual.PerpetualOrder, len(orderParams))
	amounts := make([]*big.Int, len(orderParams))
	for _, param := range orderParams {
		perpetualOrder := perpetual.PerpetualOrder{
			Trader:    gethCommon.HexToAddress(param.Trader),
			Broker:    gethCommon.HexToAddress(param.Broker),
			Perpetual: gethCommon.HexToAddress(param.Perpetual),
			Price:     utils.MustDecimalToBigInt(utils.ToWad(param.Price)),
			Amount:    utils.MustDecimalToBigInt(utils.ToWad(param.Amount)),
			ExpiredAt: param.ExpiredAt,
			Version:   param.Version,
			Category:  param.Category,
			CloseOnly: param.CloseOnly,
			Salt:      param.Salt,
			ChainId:   param.ChainId,
			Signature: perpetual.Signature{
				Config: param.Signature.Config,
				R:      param.Signature.R,
				S:      param.Signature.S,
			},
		}
		orders = append(orders, perpetualOrder)
		amounts = append(amounts, perpetualOrder.Amount)
	}

	for _, amount := range matchAmounts {
		amounts = append(amounts, utils.MustDecimalToBigInt(utils.ToWad(amount)))
	}
	inputs, err := parsed.Pack("batchTrade", orders, amounts, gases, TradeFailureOption)
	return inputs, err
}

func (c *Client) SendBatchTrade(ctx context.Context) {
	return
}

func (c *Client) WaitForReceipt(ctx context.Context, transactionHash string) (blockNumber uint64, blockHash string, succ bool, err error) {
	receipt, err := c.ethCli.TransactionReceipt(ctx, gethCommon.HexToHash(transactionHash))
	if err != nil {
		err = fmt.Errorf("fail to retrieve transaction receipt error:%w", err)
		return
	}

	blockNumber = receipt.BlockNumber.Uint64()
	blockHash = receipt.BlockHash.Hex()
	if receipt.Status == ethtypes.ReceiptStatusSuccessful {
		succ = true
	} else {
		succ = false
	}
	return
}
