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
	"github.com/shopspring/decimal"
)

const GasLimitRelax = 120

var InvalidParamError = errors.New("Invalid Parameter")
var TransactionRevertError = errors.New("TransactionRevertError")
var GasLimitTooSmallError = errors.New("GasLimitTooSmall")
var BlockchainError = errors.New("BlockChainError")
var InsufficientGasError = errors.New("InsufficientGasError")

type L2Relayer struct {
	cli          *ethclient.Client
	address      common.Address
	transactOpts *bind.TransactOpts
	contract     *L2RelayerContract
	mu           sync.Mutex
	gasFeeRate   uint32
}

func NewL2Relayer(rpc string, chainId *big.Int, key string, contract string, gasPrice uint64, gasFeeRate uint32) (*L2Relayer, error) {
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
		cli:          cli,
		address:      address,
		transactOpts: transactOpts,
		contract:     relayerContract,
	}, nil
}

func (r *L2Relayer) CallFunction(ctx context.Context, functionSignature string, callData string, userAddress string, nonce uint32, expiration uint32, gasFeeLimit uint64, signature string) (tx string, err error) {
	msg, err := NewMai3SignedCallMessage(functionSignature, callData, userAddress, nonce, expiration, gasFeeLimit, signature)
	if err != nil {
		return "", fmt.Errorf("%e:%w", err, InvalidParamError)
	}

	var opts = *r.transactOpts
	opts.Context = ctx
	gas, err := r.contract.EstimateFunctionGas(&opts, msg)
	if err != nil {
		return "", fmt.Errorf("CallFunction:EstimateFunctionGas:%e:%w", err, TransactionRevertError)
	}

	fee := gas * uint64(r.gasFeeRate) / 100
	if fee > gasFeeLimit {
		return "", fmt.Errorf("CallFunction:%w", GasLimitTooSmallError)
	}

	if fee > 0 {
		reserved, err := r.contract.ReservedGas(msg.UserAddress)
		if err != nil {
			return "", fmt.Errorf("CallFunction:%e:%w", err, BlockchainError)
		}
		decimalFee, err := decimal.NewFromString(fmt.Sprintf("%d", fee))
		if err != nil {
			return "", fmt.Errorf("bad fee")
		}
		if reserved.LessThan(decimalFee) {
			return "", fmt.Errorf("CallFunction:%w", InsufficientGasError)
		}
	}

	opts.GasLimit = gas * GasLimitRelax / 100

	r.mu.Lock()
	defer r.mu.Unlock()

	transaction, err := r.contract.CallFunction(&opts, msg)
	if err != nil {
		return "", fmt.Errorf("CallFunction:%e:%w", err, BlockchainError)
	}

	return transaction.Hash().Hex(), nil
}
