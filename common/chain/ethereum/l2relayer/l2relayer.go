// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2relayer

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// L2RelayerABI is the input ABI used to generate the binding from.
const L2RelayerABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"functionSignature\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"userData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"callFunction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// L2Relayer is an auto generated Go binding around an Ethereum contract.
type L2Relayer struct {
	L2RelayerCaller     // Read-only binding to the contract
	L2RelayerTransactor // Write-only binding to the contract
	L2RelayerFilterer   // Log filterer for contract events
}

// L2RelayerCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2RelayerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2RelayerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2RelayerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2RelayerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2RelayerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2RelayerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2RelayerSession struct {
	Contract     *L2Relayer        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2RelayerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2RelayerCallerSession struct {
	Contract *L2RelayerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// L2RelayerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2RelayerTransactorSession struct {
	Contract     *L2RelayerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// L2RelayerRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2RelayerRaw struct {
	Contract *L2Relayer // Generic contract binding to access the raw methods on
}

// L2RelayerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2RelayerCallerRaw struct {
	Contract *L2RelayerCaller // Generic read-only contract binding to access the raw methods on
}

// L2RelayerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2RelayerTransactorRaw struct {
	Contract *L2RelayerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2Relayer creates a new instance of L2Relayer, bound to a specific deployed contract.
func NewL2Relayer(address common.Address, backend bind.ContractBackend) (*L2Relayer, error) {
	contract, err := bindL2Relayer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2Relayer{L2RelayerCaller: L2RelayerCaller{contract: contract}, L2RelayerTransactor: L2RelayerTransactor{contract: contract}, L2RelayerFilterer: L2RelayerFilterer{contract: contract}}, nil
}

// NewL2RelayerCaller creates a new read-only instance of L2Relayer, bound to a specific deployed contract.
func NewL2RelayerCaller(address common.Address, caller bind.ContractCaller) (*L2RelayerCaller, error) {
	contract, err := bindL2Relayer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2RelayerCaller{contract: contract}, nil
}

// NewL2RelayerTransactor creates a new write-only instance of L2Relayer, bound to a specific deployed contract.
func NewL2RelayerTransactor(address common.Address, transactor bind.ContractTransactor) (*L2RelayerTransactor, error) {
	contract, err := bindL2Relayer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2RelayerTransactor{contract: contract}, nil
}

// NewL2RelayerFilterer creates a new log filterer instance of L2Relayer, bound to a specific deployed contract.
func NewL2RelayerFilterer(address common.Address, filterer bind.ContractFilterer) (*L2RelayerFilterer, error) {
	contract, err := bindL2Relayer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2RelayerFilterer{contract: contract}, nil
}

// bindL2Relayer binds a generic wrapper to an already deployed contract.
func bindL2Relayer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2RelayerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2Relayer *L2RelayerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2Relayer.Contract.L2RelayerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2Relayer *L2RelayerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2Relayer.Contract.L2RelayerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2Relayer *L2RelayerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2Relayer.Contract.L2RelayerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2Relayer *L2RelayerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2Relayer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2Relayer *L2RelayerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2Relayer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2Relayer *L2RelayerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2Relayer.Contract.contract.Transact(opts, method, params...)
}

// CallFunction is a paid mutator transaction binding the contract method 0x940f3445.
//
// Solidity: function callFunction(string functionSignature, bytes callData, bytes32 userData, bytes signature) returns()
func (_L2Relayer *L2RelayerTransactor) CallFunction(opts *bind.TransactOpts, functionSignature string, callData []byte, userData [32]byte, signature []byte) (*types.Transaction, error) {
	return _L2Relayer.contract.Transact(opts, "callFunction", functionSignature, callData, userData, signature)
}

// CallFunction is a paid mutator transaction binding the contract method 0x940f3445.
//
// Solidity: function callFunction(string functionSignature, bytes callData, bytes32 userData, bytes signature) returns()
func (_L2Relayer *L2RelayerSession) CallFunction(functionSignature string, callData []byte, userData [32]byte, signature []byte) (*types.Transaction, error) {
	return _L2Relayer.Contract.CallFunction(&_L2Relayer.TransactOpts, functionSignature, callData, userData, signature)
}

// CallFunction is a paid mutator transaction binding the contract method 0x940f3445.
//
// Solidity: function callFunction(string functionSignature, bytes callData, bytes32 userData, bytes signature) returns()
func (_L2Relayer *L2RelayerTransactorSession) CallFunction(functionSignature string, callData []byte, userData [32]byte, signature []byte) (*types.Transaction, error) {
	return _L2Relayer.Contract.CallFunction(&_L2Relayer.TransactOpts, functionSignature, callData, userData, signature)
}
