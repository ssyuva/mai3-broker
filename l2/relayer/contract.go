package relayer

import (
	"bytes"
	"context"
	"encoding/binary"
	"math/big"
	"strings"

	"github.com/mcdexio/mai3-broker/common/chain/ethereum/broker"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const PackedGasFeeLimitShift uint64 = 10 ^ 11

type mai3SignedCallMessage struct {
	From       common.Address
	To         common.Address
	Method     string
	CallData   []byte
	Nonce      uint32
	Expiration uint32
	GasLimit   uint64
	Signaure   []byte
	Fee        uint64

	UserData1 [32]byte // useraddress[20] nonce[4] expiration[4] gasfeelimit[4]
	UserData2 [32]byte // to[20] fee[4]
}

func packUserData(from, to common.Address, Nonce, Expiration uint32, GasFeeLimit, fee uint64) ([32]byte, [32]byte) {
	buf := new(bytes.Buffer)
	buf.Write(from.Bytes())
	_ = binary.Write(buf, binary.BigEndian, Nonce)
	_ = binary.Write(buf, binary.BigEndian, Expiration)
	u32GasFeeLimit := uint32(GasFeeLimit / PackedGasFeeLimitShift)
	_ = binary.Write(buf, binary.BigEndian, u32GasFeeLimit)
	if buf.Len() > 32 {
		panic("bad buffer size")
	}
	var user1, user2 [32]byte
	copy(user1[:], buf.Bytes()[:32])

	buf.Reset()
	buf.Write(to.Bytes())
	u32Fee := uint32(fee / PackedGasFeeLimitShift)
	_ = binary.Write(buf, binary.BigEndian, u32Fee)
	if buf.Len() > 32 {
		panic("bad buffer size")
	}
	copy(user2[:], buf.Bytes()[:32])

	return user1, user2
}

func (m *mai3SignedCallMessage) setUserDataFee(fee uint64) {
	buf := bytes.NewBuffer(m.UserData2[20:24])
	_ = binary.Write(buf, binary.BigEndian, uint32(fee/PackedGasFeeLimitShift))
}

func newMai3SignedCallMessage(from, to, method, callData string, nonce, expiration uint32, gasFeeLimit uint64, signature string, fee uint64) (*mai3SignedCallMessage, error) {
	if !common.IsHexAddress(from) {
		return nil, NewInvalidRequestError("bad from address")
	}
	fromAddress := common.HexToAddress(from)

	if !common.IsHexAddress(to) {
		return nil, NewInvalidRequestError("bad to address")
	}
	toAddress := common.HexToAddress(to)

	userData1, userData2 := packUserData(fromAddress, toAddress, nonce, expiration, gasFeeLimit, fee)

	return &mai3SignedCallMessage{
		From:       fromAddress,
		To:         toAddress,
		Method:     method,
		CallData:   common.FromHex(callData),
		Nonce:      nonce,
		Expiration: expiration,
		GasLimit:   gasFeeLimit,
		Signaure:   common.FromHex(signature),
		UserData1:  userData1,
		UserData2:  userData2,
	}, nil

}

func (msg *mai3SignedCallMessage) FunctionCallParams() []interface{} {
	return []interface{}{msg.UserData1, msg.UserData2, msg.Method, msg.CallData, msg.Signaure}
}

type brokerContract struct {
	address common.Address
	abi     abi.ABI
	backend bind.ContractBackend
	broker  *broker.Broker
}

func newBrokerContract(address common.Address, backend bind.ContractBackend) (*brokerContract, error) {
	parsed, err := abi.JSON(strings.NewReader(broker.BrokerABI))
	if err != nil {
		return nil, err
	}
	broker, err := broker.NewBroker(address, backend)
	if err != nil {
		return nil, err
	}
	return &brokerContract{
		address: address,
		abi:     parsed,
		backend: backend,
		broker:  broker,
	}, nil
}

func (c *brokerContract) callFunction(opts *bind.TransactOpts, msg *mai3SignedCallMessage) (*types.Transaction, error) {
	return c.broker.CallFunction(opts, msg.UserData1, msg.UserData2, msg.Method, msg.CallData, msg.Signaure)
}

func (c *brokerContract) estimateFunctionGas(opts *bind.TransactOpts, msg *mai3SignedCallMessage) (uint64, error) {
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
		gasPrice = big.NewInt(0)
	}

	txMsg := ethereum.CallMsg{From: opts.From, To: &c.address, GasPrice: gasPrice, Value: value, Data: input}
	gasLimit, err := c.backend.EstimateGas(ensureContext(opts.Context), txMsg)
	if err != nil {
		return 0, err
	}
	return gasLimit, nil
}

func (c *brokerContract) balanceOf(opts *bind.CallOpts, userAddress common.Address) (*big.Int, error) {
	return c.broker.BalanceOf(opts, userAddress)
}

// ensureContext is a helper method to ensure a context is not nil, even if the
// user specified it as such.
func ensureContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.TODO()
	}
	return ctx
}
