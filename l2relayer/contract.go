package l2relayer

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"strings"

	contract "github.com/mcarloai/mai-v3-broker/common/chain/ethereum/l2relayer"
	"github.com/shopspring/decimal"

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
	Signaure          string
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
		Signaure:          signature,
		UserData:          userData,
	}, nil

}

func (tx *Mai3SignedCallMessage) FunctionCallParams() []interface{} {
	return []interface{}{tx.FunctionSignature, tx.CallData, tx.UserData, tx.Signaure}
}

type L2RelayerContract struct {
	address    common.Address
	abi        abi.ABI
	transactor bind.ContractTransactor
}

func NewRelayerContract(address common.Address, transactor bind.ContractTransactor) (*L2RelayerContract, error) {
	parsed, err := abi.JSON(strings.NewReader(contract.L2RelayerABI))
	if err != nil {
		return nil, err
	}
	return &L2RelayerContract{
		address:    address,
		abi:        parsed,
		transactor: transactor,
	}, nil
}

func (c *L2RelayerContract) CallFunction(opts *bind.TransactOpts, msg *Mai3SignedCallMessage) (*types.Transaction, error) {
	input, err := c.abi.Pack("callFunction", msg.FunctionCallParams()...)
	if err != nil {
		return nil, err
	}
	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}

	var nonce uint64
	if opts.Nonce == nil {
		nonce, err = c.transactor.PendingNonceAt(ensureContext(opts.Context), opts.From)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve account nonce: %v", err)
		}
	} else {
		nonce = opts.Nonce.Uint64()
	}

	gasPrice := opts.GasPrice
	if gasPrice == nil {
		return nil, errors.New("invalid gas price")
	}

	gasLimit := opts.GasLimit
	if gasLimit == 0 {
		return nil, errors.New("invlid gas limit")
	}

	rawTx := types.NewTransaction(nonce, c.address, value, gasLimit, gasPrice, input)
	signedTx, err := opts.Signer(opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	if err := c.transactor.SendTransaction(ensureContext(opts.Context), signedTx); err != nil {
		return nil, err
	}
	return signedTx, nil
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
	gasLimit, err := c.transactor.EstimateGas(ensureContext(opts.Context), txMsg)
	if err != nil {
		return 0, fmt.Errorf("failed to estimate gas needed: %w", err)
	}
	return gasLimit, nil
}

func (c *L2RelayerContract) ReservedGas(userAddress common.Address) (decimal.Decimal, error) {
	return decimal.Zero, nil
}

// ensureContext is a helper method to ensure a context is not nil, even if the
// user specified it as such.
func ensureContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.TODO()
	}
	return ctx
}
