// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package broker

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

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Trader         common.Address
	Broker         common.Address
	Relayer        common.Address
	Referrer       common.Address
	LiquidityPool  common.Address
	MinTradeAmount *big.Int
	Amount         *big.Int
	LimitPrice     *big.Int
	TriggerPrice   *big.Int
	ChainID        *big.Int
	ExpiredAt      uint64
	PerpetualIndex uint32
	BrokerFeeLimit uint32
	Flags          uint32
	Salt           uint32
}

// BrokerABI is the input ABI used to generate the binding from.
const BrokerABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"userData1\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"userData2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"functionSignature\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"CallFunction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"CancelOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fillAmount\",\"type\":\"int256\"}],\"name\":\"FillOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"broker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"minTradeAmount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"limitPrice\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"triggerPrice\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expiredAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"perpetualIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"brokerFeeLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"flags\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"salt\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"TradeFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"broker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"minTradeAmount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"limitPrice\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"triggerPrice\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expiredAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"perpetualIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"brokerFeeLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"flags\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"salt\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasReward\",\"type\":\"uint256\"}],\"name\":\"TradeSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"compressedOrders\",\"type\":\"bytes[]\"},{\"internalType\":\"int256[]\",\"name\":\"amounts\",\"type\":\"int256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasRewards\",\"type\":\"uint256[]\"}],\"name\":\"batchTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"broker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"minTradeAmount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"limitPrice\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"triggerPrice\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expiredAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"perpetualIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"brokerFeeLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"flags\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"salt\",\"type\":\"uint32\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"broker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"minTradeAmount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"limitPrice\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"triggerPrice\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expiredAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"perpetualIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"brokerFeeLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"flags\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"salt\",\"type\":\"uint32\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getOrderFilledAmount\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"filledAmount\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"broker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"minTradeAmount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"limitPrice\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"triggerPrice\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expiredAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"perpetualIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"brokerFeeLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"flags\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"salt\",\"type\":\"uint32\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"isOrderCanceled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Broker is an auto generated Go binding around an Ethereum contract.
type Broker struct {
	BrokerCaller     // Read-only binding to the contract
	BrokerTransactor // Write-only binding to the contract
	BrokerFilterer   // Log filterer for contract events
}

// BrokerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BrokerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BrokerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BrokerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BrokerSession struct {
	Contract     *Broker           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrokerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BrokerCallerSession struct {
	Contract *BrokerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BrokerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BrokerTransactorSession struct {
	Contract     *BrokerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrokerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BrokerRaw struct {
	Contract *Broker // Generic contract binding to access the raw methods on
}

// BrokerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BrokerCallerRaw struct {
	Contract *BrokerCaller // Generic read-only contract binding to access the raw methods on
}

// BrokerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BrokerTransactorRaw struct {
	Contract *BrokerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBroker creates a new instance of Broker, bound to a specific deployed contract.
func NewBroker(address common.Address, backend bind.ContractBackend) (*Broker, error) {
	contract, err := bindBroker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Broker{BrokerCaller: BrokerCaller{contract: contract}, BrokerTransactor: BrokerTransactor{contract: contract}, BrokerFilterer: BrokerFilterer{contract: contract}}, nil
}

// NewBrokerCaller creates a new read-only instance of Broker, bound to a specific deployed contract.
func NewBrokerCaller(address common.Address, caller bind.ContractCaller) (*BrokerCaller, error) {
	contract, err := bindBroker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BrokerCaller{contract: contract}, nil
}

// NewBrokerTransactor creates a new write-only instance of Broker, bound to a specific deployed contract.
func NewBrokerTransactor(address common.Address, transactor bind.ContractTransactor) (*BrokerTransactor, error) {
	contract, err := bindBroker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BrokerTransactor{contract: contract}, nil
}

// NewBrokerFilterer creates a new log filterer instance of Broker, bound to a specific deployed contract.
func NewBrokerFilterer(address common.Address, filterer bind.ContractFilterer) (*BrokerFilterer, error) {
	contract, err := bindBroker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BrokerFilterer{contract: contract}, nil
}

// bindBroker binds a generic wrapper to an already deployed contract.
func bindBroker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BrokerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Broker *BrokerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Broker.Contract.BrokerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Broker *BrokerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Broker.Contract.BrokerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Broker *BrokerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Broker.Contract.BrokerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Broker *BrokerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Broker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Broker *BrokerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Broker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Broker *BrokerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Broker.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address trader) view returns(uint256)
func (_Broker *BrokerCaller) BalanceOf(opts *bind.CallOpts, trader common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "balanceOf", trader)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address trader) view returns(uint256)
func (_Broker *BrokerSession) BalanceOf(trader common.Address) (*big.Int, error) {
	return _Broker.Contract.BalanceOf(&_Broker.CallOpts, trader)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address trader) view returns(uint256)
func (_Broker *BrokerCallerSession) BalanceOf(trader common.Address) (*big.Int, error) {
	return _Broker.Contract.BalanceOf(&_Broker.CallOpts, trader)
}

// GetOrderFilledAmount is a free data retrieval call binding the contract method 0x4e199146.
//
// Solidity: function getOrderFilledAmount((address,address,address,address,address,int256,int256,int256,int256,uint256,uint64,uint32,uint32,uint32,uint32) order) view returns(int256 filledAmount)
func (_Broker *BrokerCaller) GetOrderFilledAmount(opts *bind.CallOpts, order Order) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getOrderFilledAmount", order)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrderFilledAmount is a free data retrieval call binding the contract method 0x4e199146.
//
// Solidity: function getOrderFilledAmount((address,address,address,address,address,int256,int256,int256,int256,uint256,uint64,uint32,uint32,uint32,uint32) order) view returns(int256 filledAmount)
func (_Broker *BrokerSession) GetOrderFilledAmount(order Order) (*big.Int, error) {
	return _Broker.Contract.GetOrderFilledAmount(&_Broker.CallOpts, order)
}

// GetOrderFilledAmount is a free data retrieval call binding the contract method 0x4e199146.
//
// Solidity: function getOrderFilledAmount((address,address,address,address,address,int256,int256,int256,int256,uint256,uint64,uint32,uint32,uint32,uint32) order) view returns(int256 filledAmount)
func (_Broker *BrokerCallerSession) GetOrderFilledAmount(order Order) (*big.Int, error) {
	return _Broker.Contract.GetOrderFilledAmount(&_Broker.CallOpts, order)
}

// IsOrderCanceled is a free data retrieval call binding the contract method 0x89d1e304.
//
// Solidity: function isOrderCanceled((address,address,address,address,address,int256,int256,int256,int256,uint256,uint64,uint32,uint32,uint32,uint32) order) view returns(bool)
func (_Broker *BrokerCaller) IsOrderCanceled(opts *bind.CallOpts, order Order) (bool, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "isOrderCanceled", order)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOrderCanceled is a free data retrieval call binding the contract method 0x89d1e304.
//
// Solidity: function isOrderCanceled((address,address,address,address,address,int256,int256,int256,int256,uint256,uint64,uint32,uint32,uint32,uint32) order) view returns(bool)
func (_Broker *BrokerSession) IsOrderCanceled(order Order) (bool, error) {
	return _Broker.Contract.IsOrderCanceled(&_Broker.CallOpts, order)
}

// IsOrderCanceled is a free data retrieval call binding the contract method 0x89d1e304.
//
// Solidity: function isOrderCanceled((address,address,address,address,address,int256,int256,int256,int256,uint256,uint64,uint32,uint32,uint32,uint32) order) view returns(bool)
func (_Broker *BrokerCallerSession) IsOrderCanceled(order Order) (bool, error) {
	return _Broker.Contract.IsOrderCanceled(&_Broker.CallOpts, order)
}

// BatchTrade is a paid mutator transaction binding the contract method 0xabee6e5c.
//
// Solidity: function batchTrade(bytes[] compressedOrders, int256[] amounts, uint256[] gasRewards) returns()
func (_Broker *BrokerTransactor) BatchTrade(opts *bind.TransactOpts, compressedOrders [][]byte, amounts []*big.Int, gasRewards []*big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "batchTrade", compressedOrders, amounts, gasRewards)
}

// BatchTrade is a paid mutator transaction binding the contract method 0xabee6e5c.
//
// Solidity: function batchTrade(bytes[] compressedOrders, int256[] amounts, uint256[] gasRewards) returns()
func (_Broker *BrokerSession) BatchTrade(compressedOrders [][]byte, amounts []*big.Int, gasRewards []*big.Int) (*types.Transaction, error) {
	return _Broker.Contract.BatchTrade(&_Broker.TransactOpts, compressedOrders, amounts, gasRewards)
}

// BatchTrade is a paid mutator transaction binding the contract method 0xabee6e5c.
//
// Solidity: function batchTrade(bytes[] compressedOrders, int256[] amounts, uint256[] gasRewards) returns()
func (_Broker *BrokerTransactorSession) BatchTrade(compressedOrders [][]byte, amounts []*big.Int, gasRewards []*big.Int) (*types.Transaction, error) {
	return _Broker.Contract.BatchTrade(&_Broker.TransactOpts, compressedOrders, amounts, gasRewards)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xce4d643b.
//
// Solidity: function cancelOrder((address,address,address,address,address,int256,int256,int256,int256,uint256,uint64,uint32,uint32,uint32,uint32) order) returns()
func (_Broker *BrokerTransactor) CancelOrder(opts *bind.TransactOpts, order Order) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xce4d643b.
//
// Solidity: function cancelOrder((address,address,address,address,address,int256,int256,int256,int256,uint256,uint64,uint32,uint32,uint32,uint32) order) returns()
func (_Broker *BrokerSession) CancelOrder(order Order) (*types.Transaction, error) {
	return _Broker.Contract.CancelOrder(&_Broker.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xce4d643b.
//
// Solidity: function cancelOrder((address,address,address,address,address,int256,int256,int256,int256,uint256,uint64,uint32,uint32,uint32,uint32) order) returns()
func (_Broker *BrokerTransactorSession) CancelOrder(order Order) (*types.Transaction, error) {
	return _Broker.Contract.CancelOrder(&_Broker.TransactOpts, order)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Broker *BrokerTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Broker *BrokerSession) Deposit() (*types.Transaction, error) {
	return _Broker.Contract.Deposit(&_Broker.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Broker *BrokerTransactorSession) Deposit() (*types.Transaction, error) {
	return _Broker.Contract.Deposit(&_Broker.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Broker *BrokerTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Broker *BrokerSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.Withdraw(&_Broker.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Broker *BrokerTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.Withdraw(&_Broker.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Broker *BrokerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Broker.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Broker *BrokerSession) Receive() (*types.Transaction, error) {
	return _Broker.Contract.Receive(&_Broker.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Broker *BrokerTransactorSession) Receive() (*types.Transaction, error) {
	return _Broker.Contract.Receive(&_Broker.TransactOpts)
}

// BrokerCallFunctionIterator is returned from FilterCallFunction and is used to iterate over the raw logs and unpacked data for CallFunction events raised by the Broker contract.
type BrokerCallFunctionIterator struct {
	Event *BrokerCallFunction // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BrokerCallFunctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerCallFunction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BrokerCallFunction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BrokerCallFunctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerCallFunctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerCallFunction represents a CallFunction event raised by the Broker contract.
type BrokerCallFunction struct {
	UserData1         [32]byte
	UserData2         [32]byte
	FunctionSignature string
	CallData          []byte
	Signature         []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallFunction is a free log retrieval operation binding the contract event 0xf7a099b6e4317688ffa8d752134614f9ec1394f25cbfc1646032ddead07a1997.
//
// Solidity: event CallFunction(bytes32 userData1, bytes32 userData2, string functionSignature, bytes callData, bytes signature)
func (_Broker *BrokerFilterer) FilterCallFunction(opts *bind.FilterOpts) (*BrokerCallFunctionIterator, error) {

	logs, sub, err := _Broker.contract.FilterLogs(opts, "CallFunction")
	if err != nil {
		return nil, err
	}
	return &BrokerCallFunctionIterator{contract: _Broker.contract, event: "CallFunction", logs: logs, sub: sub}, nil
}

// WatchCallFunction is a free log subscription operation binding the contract event 0xf7a099b6e4317688ffa8d752134614f9ec1394f25cbfc1646032ddead07a1997.
//
// Solidity: event CallFunction(bytes32 userData1, bytes32 userData2, string functionSignature, bytes callData, bytes signature)
func (_Broker *BrokerFilterer) WatchCallFunction(opts *bind.WatchOpts, sink chan<- *BrokerCallFunction) (event.Subscription, error) {

	logs, sub, err := _Broker.contract.WatchLogs(opts, "CallFunction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerCallFunction)
				if err := _Broker.contract.UnpackLog(event, "CallFunction", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCallFunction is a log parse operation binding the contract event 0xf7a099b6e4317688ffa8d752134614f9ec1394f25cbfc1646032ddead07a1997.
//
// Solidity: event CallFunction(bytes32 userData1, bytes32 userData2, string functionSignature, bytes callData, bytes signature)
func (_Broker *BrokerFilterer) ParseCallFunction(log types.Log) (*BrokerCallFunction, error) {
	event := new(BrokerCallFunction)
	if err := _Broker.contract.UnpackLog(event, "CallFunction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerCancelOrderIterator is returned from FilterCancelOrder and is used to iterate over the raw logs and unpacked data for CancelOrder events raised by the Broker contract.
type BrokerCancelOrderIterator struct {
	Event *BrokerCancelOrder // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BrokerCancelOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerCancelOrder)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BrokerCancelOrder)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BrokerCancelOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerCancelOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerCancelOrder represents a CancelOrder event raised by the Broker contract.
type BrokerCancelOrder struct {
	OrderHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCancelOrder is a free log retrieval operation binding the contract event 0x42c76c81a7cba1b9c861353909a184e20747ab960332628dabcbb5852fc5cbb5.
//
// Solidity: event CancelOrder(bytes32 orderHash)
func (_Broker *BrokerFilterer) FilterCancelOrder(opts *bind.FilterOpts) (*BrokerCancelOrderIterator, error) {

	logs, sub, err := _Broker.contract.FilterLogs(opts, "CancelOrder")
	if err != nil {
		return nil, err
	}
	return &BrokerCancelOrderIterator{contract: _Broker.contract, event: "CancelOrder", logs: logs, sub: sub}, nil
}

// WatchCancelOrder is a free log subscription operation binding the contract event 0x42c76c81a7cba1b9c861353909a184e20747ab960332628dabcbb5852fc5cbb5.
//
// Solidity: event CancelOrder(bytes32 orderHash)
func (_Broker *BrokerFilterer) WatchCancelOrder(opts *bind.WatchOpts, sink chan<- *BrokerCancelOrder) (event.Subscription, error) {

	logs, sub, err := _Broker.contract.WatchLogs(opts, "CancelOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerCancelOrder)
				if err := _Broker.contract.UnpackLog(event, "CancelOrder", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCancelOrder is a log parse operation binding the contract event 0x42c76c81a7cba1b9c861353909a184e20747ab960332628dabcbb5852fc5cbb5.
//
// Solidity: event CancelOrder(bytes32 orderHash)
func (_Broker *BrokerFilterer) ParseCancelOrder(log types.Log) (*BrokerCancelOrder, error) {
	event := new(BrokerCancelOrder)
	if err := _Broker.contract.UnpackLog(event, "CancelOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Broker contract.
type BrokerDepositIterator struct {
	Event *BrokerDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BrokerDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BrokerDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BrokerDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerDeposit represents a Deposit event raised by the Broker contract.
type BrokerDeposit struct {
	Trader common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed trader, uint256 amount)
func (_Broker *BrokerFilterer) FilterDeposit(opts *bind.FilterOpts, trader []common.Address) (*BrokerDepositIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Broker.contract.FilterLogs(opts, "Deposit", traderRule)
	if err != nil {
		return nil, err
	}
	return &BrokerDepositIterator{contract: _Broker.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed trader, uint256 amount)
func (_Broker *BrokerFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BrokerDeposit, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Broker.contract.WatchLogs(opts, "Deposit", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerDeposit)
				if err := _Broker.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed trader, uint256 amount)
func (_Broker *BrokerFilterer) ParseDeposit(log types.Log) (*BrokerDeposit, error) {
	event := new(BrokerDeposit)
	if err := _Broker.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerFillOrderIterator is returned from FilterFillOrder and is used to iterate over the raw logs and unpacked data for FillOrder events raised by the Broker contract.
type BrokerFillOrderIterator struct {
	Event *BrokerFillOrder // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BrokerFillOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerFillOrder)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BrokerFillOrder)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BrokerFillOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerFillOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerFillOrder represents a FillOrder event raised by the Broker contract.
type BrokerFillOrder struct {
	OrderHash  [32]byte
	FillAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFillOrder is a free log retrieval operation binding the contract event 0x66bc40b7d864f86356146c435fb0178293d08d80e04a8fba27d0e372ffe2d82b.
//
// Solidity: event FillOrder(bytes32 orderHash, int256 fillAmount)
func (_Broker *BrokerFilterer) FilterFillOrder(opts *bind.FilterOpts) (*BrokerFillOrderIterator, error) {

	logs, sub, err := _Broker.contract.FilterLogs(opts, "FillOrder")
	if err != nil {
		return nil, err
	}
	return &BrokerFillOrderIterator{contract: _Broker.contract, event: "FillOrder", logs: logs, sub: sub}, nil
}

// WatchFillOrder is a free log subscription operation binding the contract event 0x66bc40b7d864f86356146c435fb0178293d08d80e04a8fba27d0e372ffe2d82b.
//
// Solidity: event FillOrder(bytes32 orderHash, int256 fillAmount)
func (_Broker *BrokerFilterer) WatchFillOrder(opts *bind.WatchOpts, sink chan<- *BrokerFillOrder) (event.Subscription, error) {

	logs, sub, err := _Broker.contract.WatchLogs(opts, "FillOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerFillOrder)
				if err := _Broker.contract.UnpackLog(event, "FillOrder", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFillOrder is a log parse operation binding the contract event 0x66bc40b7d864f86356146c435fb0178293d08d80e04a8fba27d0e372ffe2d82b.
//
// Solidity: event FillOrder(bytes32 orderHash, int256 fillAmount)
func (_Broker *BrokerFilterer) ParseFillOrder(log types.Log) (*BrokerFillOrder, error) {
	event := new(BrokerFillOrder)
	if err := _Broker.contract.UnpackLog(event, "FillOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerTradeFailedIterator is returned from FilterTradeFailed and is used to iterate over the raw logs and unpacked data for TradeFailed events raised by the Broker contract.
type BrokerTradeFailedIterator struct {
	Event *BrokerTradeFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BrokerTradeFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerTradeFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BrokerTradeFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BrokerTradeFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerTradeFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerTradeFailed represents a TradeFailed event raised by the Broker contract.
type BrokerTradeFailed struct {
	OrderHash [32]byte
	Order     Order
	Amount    *big.Int
	Reason    string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTradeFailed is a free log retrieval operation binding the contract event 0x1955905acc03ff235236ed3bc847a2b24e6a1a945754ffad043f7b8e01adaeb7.
//
// Solidity: event TradeFailed(bytes32 orderHash, (address,address,address,address,address,int256,int256,int256,int256,uint256,uint64,uint32,uint32,uint32,uint32) order, int256 amount, string reason)
func (_Broker *BrokerFilterer) FilterTradeFailed(opts *bind.FilterOpts) (*BrokerTradeFailedIterator, error) {

	logs, sub, err := _Broker.contract.FilterLogs(opts, "TradeFailed")
	if err != nil {
		return nil, err
	}
	return &BrokerTradeFailedIterator{contract: _Broker.contract, event: "TradeFailed", logs: logs, sub: sub}, nil
}

// WatchTradeFailed is a free log subscription operation binding the contract event 0x1955905acc03ff235236ed3bc847a2b24e6a1a945754ffad043f7b8e01adaeb7.
//
// Solidity: event TradeFailed(bytes32 orderHash, (address,address,address,address,address,int256,int256,int256,int256,uint256,uint64,uint32,uint32,uint32,uint32) order, int256 amount, string reason)
func (_Broker *BrokerFilterer) WatchTradeFailed(opts *bind.WatchOpts, sink chan<- *BrokerTradeFailed) (event.Subscription, error) {

	logs, sub, err := _Broker.contract.WatchLogs(opts, "TradeFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerTradeFailed)
				if err := _Broker.contract.UnpackLog(event, "TradeFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTradeFailed is a log parse operation binding the contract event 0x1955905acc03ff235236ed3bc847a2b24e6a1a945754ffad043f7b8e01adaeb7.
//
// Solidity: event TradeFailed(bytes32 orderHash, (address,address,address,address,address,int256,int256,int256,int256,uint256,uint64,uint32,uint32,uint32,uint32) order, int256 amount, string reason)
func (_Broker *BrokerFilterer) ParseTradeFailed(log types.Log) (*BrokerTradeFailed, error) {
	event := new(BrokerTradeFailed)
	if err := _Broker.contract.UnpackLog(event, "TradeFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerTradeSuccessIterator is returned from FilterTradeSuccess and is used to iterate over the raw logs and unpacked data for TradeSuccess events raised by the Broker contract.
type BrokerTradeSuccessIterator struct {
	Event *BrokerTradeSuccess // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BrokerTradeSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerTradeSuccess)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BrokerTradeSuccess)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BrokerTradeSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerTradeSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerTradeSuccess represents a TradeSuccess event raised by the Broker contract.
type BrokerTradeSuccess struct {
	OrderHash [32]byte
	Order     Order
	Amount    *big.Int
	GasReward *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTradeSuccess is a free log retrieval operation binding the contract event 0xa31c4dcf97bfab49d7f55b0dbc9a7a49d25cdfa4cead657a12315b48bf11f89a.
//
// Solidity: event TradeSuccess(bytes32 orderHash, (address,address,address,address,address,int256,int256,int256,int256,uint256,uint64,uint32,uint32,uint32,uint32) order, int256 amount, uint256 gasReward)
func (_Broker *BrokerFilterer) FilterTradeSuccess(opts *bind.FilterOpts) (*BrokerTradeSuccessIterator, error) {

	logs, sub, err := _Broker.contract.FilterLogs(opts, "TradeSuccess")
	if err != nil {
		return nil, err
	}
	return &BrokerTradeSuccessIterator{contract: _Broker.contract, event: "TradeSuccess", logs: logs, sub: sub}, nil
}

// WatchTradeSuccess is a free log subscription operation binding the contract event 0xa31c4dcf97bfab49d7f55b0dbc9a7a49d25cdfa4cead657a12315b48bf11f89a.
//
// Solidity: event TradeSuccess(bytes32 orderHash, (address,address,address,address,address,int256,int256,int256,int256,uint256,uint64,uint32,uint32,uint32,uint32) order, int256 amount, uint256 gasReward)
func (_Broker *BrokerFilterer) WatchTradeSuccess(opts *bind.WatchOpts, sink chan<- *BrokerTradeSuccess) (event.Subscription, error) {

	logs, sub, err := _Broker.contract.WatchLogs(opts, "TradeSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerTradeSuccess)
				if err := _Broker.contract.UnpackLog(event, "TradeSuccess", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTradeSuccess is a log parse operation binding the contract event 0xa31c4dcf97bfab49d7f55b0dbc9a7a49d25cdfa4cead657a12315b48bf11f89a.
//
// Solidity: event TradeSuccess(bytes32 orderHash, (address,address,address,address,address,int256,int256,int256,int256,uint256,uint64,uint32,uint32,uint32,uint32) order, int256 amount, uint256 gasReward)
func (_Broker *BrokerFilterer) ParseTradeSuccess(log types.Log) (*BrokerTradeSuccess, error) {
	event := new(BrokerTradeSuccess)
	if err := _Broker.contract.UnpackLog(event, "TradeSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Broker contract.
type BrokerTransferIterator struct {
	Event *BrokerTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BrokerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BrokerTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BrokerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerTransfer represents a Transfer event raised by the Broker contract.
type BrokerTransfer struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed recipient, uint256 amount)
func (_Broker *BrokerFilterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*BrokerTransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Broker.contract.FilterLogs(opts, "Transfer", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &BrokerTransferIterator{contract: _Broker.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed recipient, uint256 amount)
func (_Broker *BrokerFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BrokerTransfer, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Broker.contract.WatchLogs(opts, "Transfer", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerTransfer)
				if err := _Broker.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed recipient, uint256 amount)
func (_Broker *BrokerFilterer) ParseTransfer(log types.Log) (*BrokerTransfer, error) {
	event := new(BrokerTransfer)
	if err := _Broker.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Broker contract.
type BrokerWithdrawIterator struct {
	Event *BrokerWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BrokerWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BrokerWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BrokerWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerWithdraw represents a Withdraw event raised by the Broker contract.
type BrokerWithdraw struct {
	Trader common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed trader, uint256 amount)
func (_Broker *BrokerFilterer) FilterWithdraw(opts *bind.FilterOpts, trader []common.Address) (*BrokerWithdrawIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Broker.contract.FilterLogs(opts, "Withdraw", traderRule)
	if err != nil {
		return nil, err
	}
	return &BrokerWithdrawIterator{contract: _Broker.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed trader, uint256 amount)
func (_Broker *BrokerFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BrokerWithdraw, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Broker.contract.WatchLogs(opts, "Withdraw", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerWithdraw)
				if err := _Broker.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed trader, uint256 amount)
func (_Broker *BrokerFilterer) ParseWithdraw(log types.Log) (*BrokerWithdraw, error) {
	event := new(BrokerWithdraw)
	if err := _Broker.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
