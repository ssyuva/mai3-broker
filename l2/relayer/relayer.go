package relayer

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const GasLimitRelax = 120

type Relayer struct {
	cli          *ethclient.Client
	address      common.Address
	transactOpts *bind.TransactOpts
	contract     *brokerContract
	mu           sync.Mutex
	feePercent   uint32
}

func NewRelayer(rpc string, chainId *big.Int, key string, contract string, gasPrice uint64, feePercent uint32) (*Relayer, error) {
	cli, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, fmt.Errorf("NewRelayer:%w", err)
	}

	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return nil, fmt.Errorf("NewRelayer:%w", err)
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, fmt.Errorf("NewRelayer:%w", err)
	}

	transactOpts.GasPrice = big.NewInt(int64(gasPrice))

	if !common.IsHexAddress(contract) {
		return nil, fmt.Errorf("NewRelayer:bad broker address:%s", contract)
	}

	relayerContract, err := newBrokerContract(common.HexToAddress(contract), cli)
	if err != nil {
		return nil, fmt.Errorf("NewRelayer:%w", err)
	}

	return &Relayer{
		cli:          cli,
		address:      address,
		transactOpts: transactOpts,
		contract:     relayerContract,
		feePercent:   feePercent,
	}, nil
}

func (r *Relayer) Address() string {
	return r.address.Hex()
}

func (r *Relayer) CallFunction(ctx context.Context, from, to, method, callData string, nonce uint32, expiration uint32, gasLimit uint64, signature string) (tx string, err error) {
	msg, err := newMai3SignedCallMessage(from, to, method, callData, nonce, expiration, gasLimit, signature, 0)
	if err != nil {
		return "", err
	}

	var opts = *r.transactOpts
	opts.Context = ctx

	gas, err := r.contract.estimateFunctionGas(&opts, msg)
	if err != nil {
		return "", NewEstimateGasError(err)
	}

	fee := gas * uint64(r.feePercent) / 100
	if fee > gasLimit {
		return "", NewInsufficentGasError(fee)
	}

	err = r.checkFeeBalance(ctx, fee, msg.From)
	if err != nil {
		return "", err
	}

	msg.setUserDataFee(fee)

	opts.GasLimit = 100000000 //gas * GasLimitRelax / 100

	r.mu.Lock()
	defer r.mu.Unlock()

	transaction, err := r.contract.callFunction(&opts, msg)
	if err != nil {
		return "", NewSendTransactionError(err)
	}

	return transaction.Hash().Hex(), nil
}

func (r *Relayer) checkFeeBalance(ctx context.Context, fee uint64, trader common.Address) error {
	if fee > 0 {
		callOpts := &bind.CallOpts{
			Pending: true,
			From:    r.address,
			Context: ctx,
		}
		balance, err := r.contract.balanceOf(callOpts, trader)
		if err != nil {
			return fmt.Errorf("BalanceOf:%w", err)
		}
		if balance.Cmp(big.NewInt(int64(fee))) > 0 {
			return NewInsufficentGasError(fee)
		}
	}
	return nil
}
