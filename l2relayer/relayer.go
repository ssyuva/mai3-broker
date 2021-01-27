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
	feePercent   uint32
}

func NewL2Relayer(rpc string, chainId *big.Int, key string, contract string, gasPrice uint64, feePercent uint32) (*L2Relayer, error) {
	cli, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, fmt.Errorf("NewL2Relayer:%w", err)
	}

	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return nil, fmt.Errorf("NewL2Relayer:%w", err)
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
		return nil, fmt.Errorf("NewL2Relayer:%w", err)
	}

	return &L2Relayer{
		cli:          cli,
		address:      address,
		transactOpts: transactOpts,
		contract:     relayerContract,
		feePercent:   feePercent,
	}, nil
}

func (r *L2Relayer) Address() string {
	return r.address.Hex()
}

func (r *L2Relayer) CallFunction(ctx context.Context, from, to, method, callData string, nonce uint32, expiration uint32, gasLimit uint64, signature string) (tx string, err error) {
	msg, err := NewMai3SignedCallMessage(from, to, method, callData, nonce, expiration, gasLimit, signature, 0)
	if err != nil {
		return "", err
	}

	var opts = *r.transactOpts
	opts.Context = ctx

	fee := uint64(0)

	//TODO
	/*
			gas, err := r.contract.EstimateFunctionGas(&opts, msg)
			if err != nil {
				return "", fmt.Errorf("CallFunction:EstimateFunctionGas:%e:%w", err, TransactionRevertError)
			}



		fee := gas * uint64(r.feePercent) / 100
		if fee > gasLimit {
			return "", fmt.Errorf("CallFunction:%w", GasLimitTooSmallError)
		}

		err = r.checkFeeBalance(ctx, fee, msg.From)
		if err != nil {
			return "", fmt.Errorf("CallFunction:%w", err)
		}
	*/

	msg.SetUserDataFee(fee)

	opts.GasLimit = 100000000 //gas * GasLimitRelax / 100

	r.mu.Lock()
	defer r.mu.Unlock()

	transaction, err := r.contract.CallFunction(&opts, msg)
	if err != nil {
		return "", fmt.Errorf("CallFunction:%w", err)
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
			return fmt.Errorf("BalanceOf:%w", err)
		}
		if balance.Cmp(big.NewInt(int64(fee))) > 0 {
			return fmt.Errorf("checkFeeBalance:%w", InsufficientGasError)
		}
	}
	return nil
}
