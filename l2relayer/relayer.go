package l2relayer

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/launcher"
)

const GasLimitRelax = 120

var InvalidParamError = errors.New("Invalid Parameter")
var TransactionRevertError = errors.New("TransactionRevertError")
var GasLimitTooSmallError = errors.New("GasLimitTooSmall")
var BlockchainError = errors.New("BlockChainError")
var InsufficientGasError = errors.New("InsufficientGasError")

type L2Relayer struct {
	cli                    *ethclient.Client
	address                common.Address
	transactOpts           *bind.TransactOpts
	contract               *L2RelayerContract
	mu                     sync.Mutex
	callFunctionFeePercent uint32
	tradeFee               int64
}

func NewL2Relayer(rpc string, chainId *big.Int, key string, contract string, gasPrice uint64, callFunctionFeePercent uint32, tradeFee int64) (*L2Relayer, error) {
	cli, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, fmt.Errorf("NewL2Relayer dail rpc:%w", err)
	}

	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return nil, fmt.Errorf("NewL2Relayer parse key:%w", err)
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, fmt.Errorf("NewL2Relayer:%w", err)
	}

	transactOpts.GasPrice = big.NewInt(int64(gasPrice))

	if !common.IsHexAddress(contract) {
		return nil, fmt.Errorf("NewL2Relayer:bad contract address")
	}

	relayerContract, err := NewRelayerContract(common.HexToAddress(contract), cli)
	if err != nil {
		return nil, fmt.Errorf("NewL2Relayer:bind contract error:%w", err)
	}

	return &L2Relayer{
		cli:                    cli,
		address:                address,
		transactOpts:           transactOpts,
		contract:               relayerContract,
		callFunctionFeePercent: callFunctionFeePercent,
		tradeFee:               tradeFee,
	}, nil
}

func (r *L2Relayer) Address() string {
	return r.address.Hex()
}

func (r *L2Relayer) CallFunction(ctx context.Context, from, to, functionSignature, callData string, nonce uint32, expiration uint32, gasLimit uint64, signature string) (tx string, err error) {
	msg, err := NewMai3SignedCallMessage(from, to, functionSignature, callData, nonce, expiration, gasLimit, signature, 0)
	if err != nil {
		return "", fmt.Errorf("%e:%w", err, InvalidParamError)
	}

	var opts = *r.transactOpts
	opts.Context = ctx
	gas, err := r.contract.EstimateFunctionGas(&opts, msg)
	if err != nil {
		return "", fmt.Errorf("CallFunction:EstimateFunctionGas:%e:%w", err, TransactionRevertError)
	}

	fee := gas * uint64(r.callFunctionFeePercent) / 100
	if fee > gasLimit {
		return "", fmt.Errorf("CallFunction:%w", GasLimitTooSmallError)
	}

	err = r.checkFeeBalance(ctx, fee, msg.From)
	if err != nil {
		return "", fmt.Errorf("CallFunction:%w", err)
	}

	msg.SetUserDataFee(fee)

	opts.GasLimit = gas * GasLimitRelax / 100

	r.mu.Lock()
	defer r.mu.Unlock()

	transaction, err := r.contract.CallFunction(&opts, msg)
	if err != nil {
		return "", fmt.Errorf("CallFunction:%e:%w", err, BlockchainError)
	}

	return transaction.Hash().Hex(), nil
}

func (r *L2Relayer) Trade(ctx context.Context, order *model.Order) (tx string, err error) {
	compressedOrder, err := launcher.GetCompressOrderData(order)
	if err != nil {
		return "", fmt.Errorf("GetCompressOrderData:%w", err)
	}
	var opts = *r.transactOpts
	opts.Context = ctx

	err = r.checkFeeBalance(ctx, uint64(r.tradeFee), common.HexToAddress(order.TraderAddress))
	if err != nil {
		return "", fmt.Errorf("Trade:%w", err)
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	transaction, err := r.contract.Trade(&opts, compressedOrder, order.Amount.BigInt(), big.NewInt(r.tradeFee))
	if err != nil {
		return "", fmt.Errorf("Trade:%e:%w", err, BlockchainError)
	}
	return transaction.Hash().Hex(), nil
}

func (r *L2Relayer) checkFeeBalance(ctx context.Context, fee uint64, trader common.Address) error {
	if fee > 0 {
		callOpts := &bind.CallOpts{
			Pending: true,
			From:    r.address,
			Context: ctx,
		}
		balance, err := r.contract.BalanceOf(callOpts, trader)
		if err != nil {
			return fmt.Errorf("BalanceOf:%e:%w", err, BlockchainError)
		}
		if balance.Cmp(big.NewInt(int64(fee))) > 0 {
			return fmt.Errorf("checkFeeBalance:%w", InsufficientGasError)
		}
	}
	return nil
}
