package l2relayer

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/mcarloai/mai-v3-broker/common/chain/ethereum/broker"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const PackedGasFeeLimitShift uint64 = 10 ^ 11

type Mai3SignedCallMessage struct {
	FunctionSignature string
	CallData          []byte
	UserAddress       common.Address
	Nonce             uint32
	Expiration        uint32
	GasFeeLimit       uint64
	Signaure          []byte
	UserData          [32]byte
}

func PackUserData(userAddress common.Address, Nonce, Expiration uint32, GasFeeLimit uint64) [32]byte {
	buf := new(bytes.Buffer)
	buf.Write(userAddress.Bytes())
	_ = binary.Write(buf, binary.BigEndian, Nonce)
	_ = binary.Write(buf, binary.BigEndian, Expiration)
	u32GasFeeLimit := uint32(GasFeeLimit / PackedGasFeeLimitShift)
	_ = binary.Write(buf, binary.BigEndian, u32GasFeeLimit)
	if buf.Len() > 32 {
		panic("bad buffer size")
	}
	var result [32]byte
	copy(result[:], buf.Bytes()[:32])
	return result
}

func NewMai3SignedCallMessage(functionSignature string, callData string, userAddress string, nonce uint32, expiration uint32, gasFeeLimit uint64, signature string) (*Mai3SignedCallMessage, error) {
	if !common.IsHexAddress(userAddress) {
		return nil, errors.New("invalid user address")
	}
	user := common.HexToAddress(userAddress)

	userData := PackUserData(user, nonce, expiration, gasFeeLimit)

	return &Mai3SignedCallMessage{
		FunctionSignature: functionSignature,
		CallData:          common.Hex2Bytes(callData),
		UserAddress:       user,
		Nonce:             nonce,
		Expiration:        expiration,
		GasFeeLimit:       gasFeeLimit,
		Signaure:          common.Hex2Bytes(signature),
		UserData:          userData,
	}, nil

}

func (tx *Mai3SignedCallMessage) FunctionCallParams() []interface{} {
	return []interface{}{tx.FunctionSignature, tx.CallData, tx.UserData, tx.Signaure}
}

type L2RelayerContract struct {
	address common.Address
	abi     abi.ABI
	backend bind.ContractBackend
	broker  *broker.Broker
}

func NewRelayerContract(address common.Address, backend bind.ContractBackend) (*L2RelayerContract, error) {
	parsed, err := abi.JSON(strings.NewReader(broker.BrokerABI))
	if err != nil {
		return nil, err
	}
	broker, err := broker.NewBroker(address, backend)
	if err != nil {
		return nil, err
	}
	return &L2RelayerContract{
		address: address,
		abi:     parsed,
		backend: backend,
		broker:  broker,
	}, nil
}

func (c *L2RelayerContract) CallFunction(opts *bind.TransactOpts, msg *Mai3SignedCallMessage) (*types.Transaction, error) {
	return c.broker.CallFunction(opts, msg.FunctionSignature, msg.CallData, msg.UserData, msg.Signaure)
}

func (c *L2RelayerContract) EstimateFunctionGas(opts *bind.TransactOpts, msg *Mai3SignedCallMessage) (uint64, error) {
	input, err := c.abi.Pack("callFunction", msg.FunctionCallParams()...)
	if err != nil {
		return 0, err
	}
	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}

	gasPrice := opts.GasPrice
	if gasPrice == nil {
		return 0, errors.New("invalid gas price")
	}

	txMsg := ethereum.CallMsg{From: opts.From, To: &c.address, GasPrice: gasPrice, Value: value, Data: input}
	gasLimit, err := c.backend.EstimateGas(ensureContext(opts.Context), txMsg)
	if err != nil {
		return 0, fmt.Errorf("failed to estimate gas needed: %w", err)
	}
	return gasLimit, nil
}

func (c *L2RelayerContract) BalanceOf(opts *bind.CallOpts, userAddress common.Address) (*big.Int, error) {
	return c.broker.BalanceOf(opts, userAddress)
}

func (c *L2RelayerContract) Trade(opts *bind.TransactOpts, compressedOrder []byte, amount *big.Int, gasReward *big.Int) (*types.Transaction, error) {
	orders := [][]byte{compressedOrder}
	amounts := []*big.Int{amount}
	gasRewards := []*big.Int{gasReward}
	return c.broker.BatchTrade(opts, orders, amounts, gasRewards)
}

// ensureContext is a helper method to ensure a context is not nil, even if the
// user specified it as such.
func ensureContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.TODO()
	}
	return ctx
}
