// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package perpetual

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PerpetualOrder is an auto generated low-level Go binding around an user-defined struct.
type PerpetualOrder struct {
	Trader    common.Address
	Broker    common.Address
	Perpetual common.Address
	Price     *big.Int
	Amount    *big.Int
	ExpiredAt uint64
	Version   uint32
	Category  uint8
	CloseOnly bool
	Inversed  bool
	Salt      uint64
	ChainId   uint64
	Signature PerpetualOrderSignature
}

// PerpetualOrderSignature is an auto generated low-level Go binding around an user-defined struct.
type PerpetualOrderSignature struct {
	Config [32]byte
	R      [32]byte
	S      [32]byte
}

// PerpetualABI is the input ABI used to generate the binding from.
const PerpetualABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"broker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"perpetual\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"uint64\",\"name\":\"expiredAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"version\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"category\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"closeOnly\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"inversed\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"config\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structPerpetual.OrderSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structPerpetual.Order\",\"name\":\"\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"Match\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Trade2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"availableMargin\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"broker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"perpetual\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"uint64\",\"name\":\"expiredAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"version\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"category\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"closeOnly\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"inversed\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"config\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structPerpetual.OrderSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structPerpetual.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"int256[]\",\"name\":\"amounts\",\"type\":\"int256[]\"},{\"internalType\":\"int256[]\",\"name\":\"gases\",\"type\":\"int256[]\"},{\"internalType\":\"enumPerpetual.FailureOption\",\"name\":\"option\",\"type\":\"uint8\"}],\"name\":\"batchTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initialMargin\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastFundingState\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maintenanceMargin\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"margin\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"marginAccount\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"setAvailableMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"setInitialMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"setMaintenanceMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"setMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"setWithdrawableMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"priceLimit\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"priceLimit\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"trade2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawableMargin\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Perpetual is an auto generated Go binding around an Ethereum contract.
type Perpetual struct {
	PerpetualCaller     // Read-only binding to the contract
	PerpetualTransactor // Write-only binding to the contract
	PerpetualFilterer   // Log filterer for contract events
}

// PerpetualCaller is an auto generated read-only Go binding around an Ethereum contract.
type PerpetualCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerpetualTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PerpetualTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerpetualFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PerpetualFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerpetualSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PerpetualSession struct {
	Contract     *Perpetual        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PerpetualCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PerpetualCallerSession struct {
	Contract *PerpetualCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PerpetualTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PerpetualTransactorSession struct {
	Contract     *PerpetualTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PerpetualRaw is an auto generated low-level Go binding around an Ethereum contract.
type PerpetualRaw struct {
	Contract *Perpetual // Generic contract binding to access the raw methods on
}

// PerpetualCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PerpetualCallerRaw struct {
	Contract *PerpetualCaller // Generic read-only contract binding to access the raw methods on
}

// PerpetualTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PerpetualTransactorRaw struct {
	Contract *PerpetualTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPerpetual creates a new instance of Perpetual, bound to a specific deployed contract.
func NewPerpetual(address common.Address, backend bind.ContractBackend) (*Perpetual, error) {
	contract, err := bindPerpetual(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Perpetual{PerpetualCaller: PerpetualCaller{contract: contract}, PerpetualTransactor: PerpetualTransactor{contract: contract}, PerpetualFilterer: PerpetualFilterer{contract: contract}}, nil
}

// NewPerpetualCaller creates a new read-only instance of Perpetual, bound to a specific deployed contract.
func NewPerpetualCaller(address common.Address, caller bind.ContractCaller) (*PerpetualCaller, error) {
	contract, err := bindPerpetual(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PerpetualCaller{contract: contract}, nil
}

// NewPerpetualTransactor creates a new write-only instance of Perpetual, bound to a specific deployed contract.
func NewPerpetualTransactor(address common.Address, transactor bind.ContractTransactor) (*PerpetualTransactor, error) {
	contract, err := bindPerpetual(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PerpetualTransactor{contract: contract}, nil
}

// NewPerpetualFilterer creates a new log filterer instance of Perpetual, bound to a specific deployed contract.
func NewPerpetualFilterer(address common.Address, filterer bind.ContractFilterer) (*PerpetualFilterer, error) {
	contract, err := bindPerpetual(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PerpetualFilterer{contract: contract}, nil
}

// bindPerpetual binds a generic wrapper to an already deployed contract.
func bindPerpetual(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PerpetualABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Perpetual *PerpetualRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Perpetual.Contract.PerpetualCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Perpetual *PerpetualRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Perpetual.Contract.PerpetualTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Perpetual *PerpetualRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Perpetual.Contract.PerpetualTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Perpetual *PerpetualCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Perpetual.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Perpetual *PerpetualTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Perpetual.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Perpetual *PerpetualTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Perpetual.Contract.contract.Transact(opts, method, params...)
}

// AvailableMargin is a free data retrieval call binding the contract method 0x711d4d55.
//
// Solidity: function availableMargin(address ) constant returns(int256)
func (_Perpetual *PerpetualCaller) AvailableMargin(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Perpetual.contract.Call(opts, out, "availableMargin", arg0)
	return *ret0, err
}

// AvailableMargin is a free data retrieval call binding the contract method 0x711d4d55.
//
// Solidity: function availableMargin(address ) constant returns(int256)
func (_Perpetual *PerpetualSession) AvailableMargin(arg0 common.Address) (*big.Int, error) {
	return _Perpetual.Contract.AvailableMargin(&_Perpetual.CallOpts, arg0)
}

// AvailableMargin is a free data retrieval call binding the contract method 0x711d4d55.
//
// Solidity: function availableMargin(address ) constant returns(int256)
func (_Perpetual *PerpetualCallerSession) AvailableMargin(arg0 common.Address) (*big.Int, error) {
	return _Perpetual.Contract.AvailableMargin(&_Perpetual.CallOpts, arg0)
}

// InitialMargin is a free data retrieval call binding the contract method 0xad1144eb.
//
// Solidity: function initialMargin(address ) constant returns(int256)
func (_Perpetual *PerpetualCaller) InitialMargin(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Perpetual.contract.Call(opts, out, "initialMargin", arg0)
	return *ret0, err
}

// InitialMargin is a free data retrieval call binding the contract method 0xad1144eb.
//
// Solidity: function initialMargin(address ) constant returns(int256)
func (_Perpetual *PerpetualSession) InitialMargin(arg0 common.Address) (*big.Int, error) {
	return _Perpetual.Contract.InitialMargin(&_Perpetual.CallOpts, arg0)
}

// InitialMargin is a free data retrieval call binding the contract method 0xad1144eb.
//
// Solidity: function initialMargin(address ) constant returns(int256)
func (_Perpetual *PerpetualCallerSession) InitialMargin(arg0 common.Address) (*big.Int, error) {
	return _Perpetual.Contract.InitialMargin(&_Perpetual.CallOpts, arg0)
}

// LastFundingState is a free data retrieval call binding the contract method 0x06a7570c.
//
// Solidity: function lastFundingState() constant returns(int256, uint256)
func (_Perpetual *PerpetualCaller) LastFundingState(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Perpetual.contract.Call(opts, out, "lastFundingState")
	return *ret0, *ret1, err
}

// LastFundingState is a free data retrieval call binding the contract method 0x06a7570c.
//
// Solidity: function lastFundingState() constant returns(int256, uint256)
func (_Perpetual *PerpetualSession) LastFundingState() (*big.Int, *big.Int, error) {
	return _Perpetual.Contract.LastFundingState(&_Perpetual.CallOpts)
}

// LastFundingState is a free data retrieval call binding the contract method 0x06a7570c.
//
// Solidity: function lastFundingState() constant returns(int256, uint256)
func (_Perpetual *PerpetualCallerSession) LastFundingState() (*big.Int, *big.Int, error) {
	return _Perpetual.Contract.LastFundingState(&_Perpetual.CallOpts)
}

// MaintenanceMargin is a free data retrieval call binding the contract method 0x6027f7a8.
//
// Solidity: function maintenanceMargin(address ) constant returns(int256)
func (_Perpetual *PerpetualCaller) MaintenanceMargin(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Perpetual.contract.Call(opts, out, "maintenanceMargin", arg0)
	return *ret0, err
}

// MaintenanceMargin is a free data retrieval call binding the contract method 0x6027f7a8.
//
// Solidity: function maintenanceMargin(address ) constant returns(int256)
func (_Perpetual *PerpetualSession) MaintenanceMargin(arg0 common.Address) (*big.Int, error) {
	return _Perpetual.Contract.MaintenanceMargin(&_Perpetual.CallOpts, arg0)
}

// MaintenanceMargin is a free data retrieval call binding the contract method 0x6027f7a8.
//
// Solidity: function maintenanceMargin(address ) constant returns(int256)
func (_Perpetual *PerpetualCallerSession) MaintenanceMargin(arg0 common.Address) (*big.Int, error) {
	return _Perpetual.Contract.MaintenanceMargin(&_Perpetual.CallOpts, arg0)
}

// Margin is a free data retrieval call binding the contract method 0xbeb3ed5d.
//
// Solidity: function margin(address ) constant returns(int256)
func (_Perpetual *PerpetualCaller) Margin(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Perpetual.contract.Call(opts, out, "margin", arg0)
	return *ret0, err
}

// Margin is a free data retrieval call binding the contract method 0xbeb3ed5d.
//
// Solidity: function margin(address ) constant returns(int256)
func (_Perpetual *PerpetualSession) Margin(arg0 common.Address) (*big.Int, error) {
	return _Perpetual.Contract.Margin(&_Perpetual.CallOpts, arg0)
}

// Margin is a free data retrieval call binding the contract method 0xbeb3ed5d.
//
// Solidity: function margin(address ) constant returns(int256)
func (_Perpetual *PerpetualCallerSession) Margin(arg0 common.Address) (*big.Int, error) {
	return _Perpetual.Contract.Margin(&_Perpetual.CallOpts, arg0)
}

// MarginAccount is a free data retrieval call binding the contract method 0x6280e472.
//
// Solidity: function marginAccount(address ) constant returns(int256, int256, int256)
func (_Perpetual *PerpetualCaller) MarginAccount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Perpetual.contract.Call(opts, out, "marginAccount", arg0)
	return *ret0, *ret1, *ret2, err
}

// MarginAccount is a free data retrieval call binding the contract method 0x6280e472.
//
// Solidity: function marginAccount(address ) constant returns(int256, int256, int256)
func (_Perpetual *PerpetualSession) MarginAccount(arg0 common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Perpetual.Contract.MarginAccount(&_Perpetual.CallOpts, arg0)
}

// MarginAccount is a free data retrieval call binding the contract method 0x6280e472.
//
// Solidity: function marginAccount(address ) constant returns(int256, int256, int256)
func (_Perpetual *PerpetualCallerSession) MarginAccount(arg0 common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Perpetual.Contract.MarginAccount(&_Perpetual.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_Perpetual *PerpetualCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Perpetual.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_Perpetual *PerpetualSession) Oracle() (common.Address, error) {
	return _Perpetual.Contract.Oracle(&_Perpetual.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_Perpetual *PerpetualCallerSession) Oracle() (common.Address, error) {
	return _Perpetual.Contract.Oracle(&_Perpetual.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(int256, uint256, int256, uint256)
func (_Perpetual *PerpetualCaller) Price(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Perpetual.contract.Call(opts, out, "price")
	return *ret0, *ret1, *ret2, *ret3, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(int256, uint256, int256, uint256)
func (_Perpetual *PerpetualSession) Price() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Perpetual.Contract.Price(&_Perpetual.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(int256, uint256, int256, uint256)
func (_Perpetual *PerpetualCallerSession) Price() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Perpetual.Contract.Price(&_Perpetual.CallOpts)
}

// WithdrawableMargin is a free data retrieval call binding the contract method 0x701d8b24.
//
// Solidity: function withdrawableMargin(address ) constant returns(int256)
func (_Perpetual *PerpetualCaller) WithdrawableMargin(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Perpetual.contract.Call(opts, out, "withdrawableMargin", arg0)
	return *ret0, err
}

// WithdrawableMargin is a free data retrieval call binding the contract method 0x701d8b24.
//
// Solidity: function withdrawableMargin(address ) constant returns(int256)
func (_Perpetual *PerpetualSession) WithdrawableMargin(arg0 common.Address) (*big.Int, error) {
	return _Perpetual.Contract.WithdrawableMargin(&_Perpetual.CallOpts, arg0)
}

// WithdrawableMargin is a free data retrieval call binding the contract method 0x701d8b24.
//
// Solidity: function withdrawableMargin(address ) constant returns(int256)
func (_Perpetual *PerpetualCallerSession) WithdrawableMargin(arg0 common.Address) (*big.Int, error) {
	return _Perpetual.Contract.WithdrawableMargin(&_Perpetual.CallOpts, arg0)
}

// BatchTrade is a paid mutator transaction binding the contract method 0x48e9ad09.
//
// Solidity: function batchTrade([]PerpetualOrder orders, int256[] amounts, int256[] gases, uint8 option) returns()
func (_Perpetual *PerpetualTransactor) BatchTrade(opts *bind.TransactOpts, orders []PerpetualOrder, amounts []*big.Int, gases []*big.Int, option uint8) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "batchTrade", orders, amounts, gases, option)
}

// BatchTrade is a paid mutator transaction binding the contract method 0x48e9ad09.
//
// Solidity: function batchTrade([]PerpetualOrder orders, int256[] amounts, int256[] gases, uint8 option) returns()
func (_Perpetual *PerpetualSession) BatchTrade(orders []PerpetualOrder, amounts []*big.Int, gases []*big.Int, option uint8) (*types.Transaction, error) {
	return _Perpetual.Contract.BatchTrade(&_Perpetual.TransactOpts, orders, amounts, gases, option)
}

// BatchTrade is a paid mutator transaction binding the contract method 0x48e9ad09.
//
// Solidity: function batchTrade([]PerpetualOrder orders, int256[] amounts, int256[] gases, uint8 option) returns()
func (_Perpetual *PerpetualTransactorSession) BatchTrade(orders []PerpetualOrder, amounts []*big.Int, gases []*big.Int, option uint8) (*types.Transaction, error) {
	return _Perpetual.Contract.BatchTrade(&_Perpetual.TransactOpts, orders, amounts, gases, option)
}

// Deposit is a paid mutator transaction binding the contract method 0xf04991f0.
//
// Solidity: function deposit(int256 amount) returns()
func (_Perpetual *PerpetualTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xf04991f0.
//
// Solidity: function deposit(int256 amount) returns()
func (_Perpetual *PerpetualSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.Deposit(&_Perpetual.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xf04991f0.
//
// Solidity: function deposit(int256 amount) returns()
func (_Perpetual *PerpetualTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.Deposit(&_Perpetual.TransactOpts, amount)
}

// SetAvailableMargin is a paid mutator transaction binding the contract method 0xe201c4d9.
//
// Solidity: function setAvailableMargin(int256 value) returns()
func (_Perpetual *PerpetualTransactor) SetAvailableMargin(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "setAvailableMargin", value)
}

// SetAvailableMargin is a paid mutator transaction binding the contract method 0xe201c4d9.
//
// Solidity: function setAvailableMargin(int256 value) returns()
func (_Perpetual *PerpetualSession) SetAvailableMargin(value *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.SetAvailableMargin(&_Perpetual.TransactOpts, value)
}

// SetAvailableMargin is a paid mutator transaction binding the contract method 0xe201c4d9.
//
// Solidity: function setAvailableMargin(int256 value) returns()
func (_Perpetual *PerpetualTransactorSession) SetAvailableMargin(value *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.SetAvailableMargin(&_Perpetual.TransactOpts, value)
}

// SetInitialMargin is a paid mutator transaction binding the contract method 0x5451f406.
//
// Solidity: function setInitialMargin(int256 value) returns()
func (_Perpetual *PerpetualTransactor) SetInitialMargin(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "setInitialMargin", value)
}

// SetInitialMargin is a paid mutator transaction binding the contract method 0x5451f406.
//
// Solidity: function setInitialMargin(int256 value) returns()
func (_Perpetual *PerpetualSession) SetInitialMargin(value *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.SetInitialMargin(&_Perpetual.TransactOpts, value)
}

// SetInitialMargin is a paid mutator transaction binding the contract method 0x5451f406.
//
// Solidity: function setInitialMargin(int256 value) returns()
func (_Perpetual *PerpetualTransactorSession) SetInitialMargin(value *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.SetInitialMargin(&_Perpetual.TransactOpts, value)
}

// SetMaintenanceMargin is a paid mutator transaction binding the contract method 0x5eeb605c.
//
// Solidity: function setMaintenanceMargin(int256 value) returns()
func (_Perpetual *PerpetualTransactor) SetMaintenanceMargin(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "setMaintenanceMargin", value)
}

// SetMaintenanceMargin is a paid mutator transaction binding the contract method 0x5eeb605c.
//
// Solidity: function setMaintenanceMargin(int256 value) returns()
func (_Perpetual *PerpetualSession) SetMaintenanceMargin(value *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.SetMaintenanceMargin(&_Perpetual.TransactOpts, value)
}

// SetMaintenanceMargin is a paid mutator transaction binding the contract method 0x5eeb605c.
//
// Solidity: function setMaintenanceMargin(int256 value) returns()
func (_Perpetual *PerpetualTransactorSession) SetMaintenanceMargin(value *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.SetMaintenanceMargin(&_Perpetual.TransactOpts, value)
}

// SetMargin is a paid mutator transaction binding the contract method 0x09d34c3b.
//
// Solidity: function setMargin(int256 value) returns()
func (_Perpetual *PerpetualTransactor) SetMargin(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "setMargin", value)
}

// SetMargin is a paid mutator transaction binding the contract method 0x09d34c3b.
//
// Solidity: function setMargin(int256 value) returns()
func (_Perpetual *PerpetualSession) SetMargin(value *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.SetMargin(&_Perpetual.TransactOpts, value)
}

// SetMargin is a paid mutator transaction binding the contract method 0x09d34c3b.
//
// Solidity: function setMargin(int256 value) returns()
func (_Perpetual *PerpetualTransactorSession) SetMargin(value *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.SetMargin(&_Perpetual.TransactOpts, value)
}

// SetWithdrawableMargin is a paid mutator transaction binding the contract method 0x5c4b7d57.
//
// Solidity: function setWithdrawableMargin(int256 value) returns()
func (_Perpetual *PerpetualTransactor) SetWithdrawableMargin(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "setWithdrawableMargin", value)
}

// SetWithdrawableMargin is a paid mutator transaction binding the contract method 0x5c4b7d57.
//
// Solidity: function setWithdrawableMargin(int256 value) returns()
func (_Perpetual *PerpetualSession) SetWithdrawableMargin(value *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.SetWithdrawableMargin(&_Perpetual.TransactOpts, value)
}

// SetWithdrawableMargin is a paid mutator transaction binding the contract method 0x5c4b7d57.
//
// Solidity: function setWithdrawableMargin(int256 value) returns()
func (_Perpetual *PerpetualTransactorSession) SetWithdrawableMargin(value *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.SetWithdrawableMargin(&_Perpetual.TransactOpts, value)
}

// Trade is a paid mutator transaction binding the contract method 0x00fe3dc4.
//
// Solidity: function trade(int256 amount, int256 priceLimit, uint256 deadline) returns()
func (_Perpetual *PerpetualTransactor) Trade(opts *bind.TransactOpts, amount *big.Int, priceLimit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "trade", amount, priceLimit, deadline)
}

// Trade is a paid mutator transaction binding the contract method 0x00fe3dc4.
//
// Solidity: function trade(int256 amount, int256 priceLimit, uint256 deadline) returns()
func (_Perpetual *PerpetualSession) Trade(amount *big.Int, priceLimit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.Trade(&_Perpetual.TransactOpts, amount, priceLimit, deadline)
}

// Trade is a paid mutator transaction binding the contract method 0x00fe3dc4.
//
// Solidity: function trade(int256 amount, int256 priceLimit, uint256 deadline) returns()
func (_Perpetual *PerpetualTransactorSession) Trade(amount *big.Int, priceLimit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.Trade(&_Perpetual.TransactOpts, amount, priceLimit, deadline)
}

// Trade2 is a paid mutator transaction binding the contract method 0x934589ba.
//
// Solidity: function trade2(address trader, int256 amount, int256 priceLimit, uint256 deadline, bytes data) returns()
func (_Perpetual *PerpetualTransactor) Trade2(opts *bind.TransactOpts, trader common.Address, amount *big.Int, priceLimit *big.Int, deadline *big.Int, data []byte) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "trade2", trader, amount, priceLimit, deadline, data)
}

// Trade2 is a paid mutator transaction binding the contract method 0x934589ba.
//
// Solidity: function trade2(address trader, int256 amount, int256 priceLimit, uint256 deadline, bytes data) returns()
func (_Perpetual *PerpetualSession) Trade2(trader common.Address, amount *big.Int, priceLimit *big.Int, deadline *big.Int, data []byte) (*types.Transaction, error) {
	return _Perpetual.Contract.Trade2(&_Perpetual.TransactOpts, trader, amount, priceLimit, deadline, data)
}

// Trade2 is a paid mutator transaction binding the contract method 0x934589ba.
//
// Solidity: function trade2(address trader, int256 amount, int256 priceLimit, uint256 deadline, bytes data) returns()
func (_Perpetual *PerpetualTransactorSession) Trade2(trader common.Address, amount *big.Int, priceLimit *big.Int, deadline *big.Int, data []byte) (*types.Transaction, error) {
	return _Perpetual.Contract.Trade2(&_Perpetual.TransactOpts, trader, amount, priceLimit, deadline, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7e62eab8.
//
// Solidity: function withdraw(int256 amount) returns()
func (_Perpetual *PerpetualTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7e62eab8.
//
// Solidity: function withdraw(int256 amount) returns()
func (_Perpetual *PerpetualSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.Withdraw(&_Perpetual.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7e62eab8.
//
// Solidity: function withdraw(int256 amount) returns()
func (_Perpetual *PerpetualTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.Withdraw(&_Perpetual.TransactOpts, amount)
}

// PerpetualDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Perpetual contract.
type PerpetualDepositIterator struct {
	Event *PerpetualDeposit // Event containing the contract specifics and raw log

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
func (it *PerpetualDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpetualDeposit)
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
		it.Event = new(PerpetualDeposit)
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
func (it *PerpetualDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpetualDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpetualDeposit represents a Deposit event raised by the Perpetual contract.
type PerpetualDeposit struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xd8a6d38df847dcba70dfdeb4948fb1457d61a81d132801f40dc9c00d52dfd478.
//
// Solidity: event Deposit(address , int256 )
func (_Perpetual *PerpetualFilterer) FilterDeposit(opts *bind.FilterOpts) (*PerpetualDepositIterator, error) {

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &PerpetualDepositIterator{contract: _Perpetual.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xd8a6d38df847dcba70dfdeb4948fb1457d61a81d132801f40dc9c00d52dfd478.
//
// Solidity: event Deposit(address , int256 )
func (_Perpetual *PerpetualFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *PerpetualDeposit) (event.Subscription, error) {

	logs, sub, err := _Perpetual.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpetualDeposit)
				if err := _Perpetual.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xd8a6d38df847dcba70dfdeb4948fb1457d61a81d132801f40dc9c00d52dfd478.
//
// Solidity: event Deposit(address , int256 )
func (_Perpetual *PerpetualFilterer) ParseDeposit(log types.Log) (*PerpetualDeposit, error) {
	event := new(PerpetualDeposit)
	if err := _Perpetual.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PerpetualMatchIterator is returned from FilterMatch and is used to iterate over the raw logs and unpacked data for Match events raised by the Perpetual contract.
type PerpetualMatchIterator struct {
	Event *PerpetualMatch // Event containing the contract specifics and raw log

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
func (it *PerpetualMatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpetualMatch)
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
		it.Event = new(PerpetualMatch)
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
func (it *PerpetualMatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpetualMatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpetualMatch represents a Match event raised by the Perpetual contract.
type PerpetualMatch struct {
	Arg0 PerpetualOrder
	Arg1 *big.Int
	Arg2 *big.Int
	Arg3 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMatch is a free log retrieval operation binding the contract event 0x9762105b16239d9062b56255920e8d1dca13b49c6a4b3062ea0ab085b6d238b9.
//
// Solidity: event Match(PerpetualOrder , int256 , int256 , bool )
func (_Perpetual *PerpetualFilterer) FilterMatch(opts *bind.FilterOpts) (*PerpetualMatchIterator, error) {

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "Match")
	if err != nil {
		return nil, err
	}
	return &PerpetualMatchIterator{contract: _Perpetual.contract, event: "Match", logs: logs, sub: sub}, nil
}

// WatchMatch is a free log subscription operation binding the contract event 0x9762105b16239d9062b56255920e8d1dca13b49c6a4b3062ea0ab085b6d238b9.
//
// Solidity: event Match(PerpetualOrder , int256 , int256 , bool )
func (_Perpetual *PerpetualFilterer) WatchMatch(opts *bind.WatchOpts, sink chan<- *PerpetualMatch) (event.Subscription, error) {

	logs, sub, err := _Perpetual.contract.WatchLogs(opts, "Match")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpetualMatch)
				if err := _Perpetual.contract.UnpackLog(event, "Match", log); err != nil {
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

// ParseMatch is a log parse operation binding the contract event 0x9762105b16239d9062b56255920e8d1dca13b49c6a4b3062ea0ab085b6d238b9.
//
// Solidity: event Match(PerpetualOrder , int256 , int256 , bool )
func (_Perpetual *PerpetualFilterer) ParseMatch(log types.Log) (*PerpetualMatch, error) {
	event := new(PerpetualMatch)
	if err := _Perpetual.contract.UnpackLog(event, "Match", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PerpetualTradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the Perpetual contract.
type PerpetualTradeIterator struct {
	Event *PerpetualTrade // Event containing the contract specifics and raw log

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
func (it *PerpetualTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpetualTrade)
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
		it.Event = new(PerpetualTrade)
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
func (it *PerpetualTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpetualTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpetualTrade represents a Trade event raised by the Perpetual contract.
type PerpetualTrade struct {
	Trader common.Address
	Amount *big.Int
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0xd5be22e057354a2acd7baadad3a29aa0d7b3b4504dcc1b1f01efb55b9ec9bfe5.
//
// Solidity: event Trade(address trader, int256 amount, int256 price)
func (_Perpetual *PerpetualFilterer) FilterTrade(opts *bind.FilterOpts) (*PerpetualTradeIterator, error) {

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "Trade")
	if err != nil {
		return nil, err
	}
	return &PerpetualTradeIterator{contract: _Perpetual.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0xd5be22e057354a2acd7baadad3a29aa0d7b3b4504dcc1b1f01efb55b9ec9bfe5.
//
// Solidity: event Trade(address trader, int256 amount, int256 price)
func (_Perpetual *PerpetualFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *PerpetualTrade) (event.Subscription, error) {

	logs, sub, err := _Perpetual.contract.WatchLogs(opts, "Trade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpetualTrade)
				if err := _Perpetual.contract.UnpackLog(event, "Trade", log); err != nil {
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

// ParseTrade is a log parse operation binding the contract event 0xd5be22e057354a2acd7baadad3a29aa0d7b3b4504dcc1b1f01efb55b9ec9bfe5.
//
// Solidity: event Trade(address trader, int256 amount, int256 price)
func (_Perpetual *PerpetualFilterer) ParseTrade(log types.Log) (*PerpetualTrade, error) {
	event := new(PerpetualTrade)
	if err := _Perpetual.contract.UnpackLog(event, "Trade", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PerpetualTrade2Iterator is returned from FilterTrade2 and is used to iterate over the raw logs and unpacked data for Trade2 events raised by the Perpetual contract.
type PerpetualTrade2Iterator struct {
	Event *PerpetualTrade2 // Event containing the contract specifics and raw log

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
func (it *PerpetualTrade2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpetualTrade2)
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
		it.Event = new(PerpetualTrade2)
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
func (it *PerpetualTrade2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpetualTrade2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpetualTrade2 represents a Trade2 event raised by the Perpetual contract.
type PerpetualTrade2 struct {
	Trader common.Address
	Amount *big.Int
	Price  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTrade2 is a free log retrieval operation binding the contract event 0xf23239ab792a36d5f2240e5d2278631e240e3926d7dffb6c0b903a5165535439.
//
// Solidity: event Trade2(address trader, int256 amount, int256 price, bytes data)
func (_Perpetual *PerpetualFilterer) FilterTrade2(opts *bind.FilterOpts) (*PerpetualTrade2Iterator, error) {

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "Trade2")
	if err != nil {
		return nil, err
	}
	return &PerpetualTrade2Iterator{contract: _Perpetual.contract, event: "Trade2", logs: logs, sub: sub}, nil
}

// WatchTrade2 is a free log subscription operation binding the contract event 0xf23239ab792a36d5f2240e5d2278631e240e3926d7dffb6c0b903a5165535439.
//
// Solidity: event Trade2(address trader, int256 amount, int256 price, bytes data)
func (_Perpetual *PerpetualFilterer) WatchTrade2(opts *bind.WatchOpts, sink chan<- *PerpetualTrade2) (event.Subscription, error) {

	logs, sub, err := _Perpetual.contract.WatchLogs(opts, "Trade2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpetualTrade2)
				if err := _Perpetual.contract.UnpackLog(event, "Trade2", log); err != nil {
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

// ParseTrade2 is a log parse operation binding the contract event 0xf23239ab792a36d5f2240e5d2278631e240e3926d7dffb6c0b903a5165535439.
//
// Solidity: event Trade2(address trader, int256 amount, int256 price, bytes data)
func (_Perpetual *PerpetualFilterer) ParseTrade2(log types.Log) (*PerpetualTrade2, error) {
	event := new(PerpetualTrade2)
	if err := _Perpetual.contract.UnpackLog(event, "Trade2", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PerpetualWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Perpetual contract.
type PerpetualWithdrawIterator struct {
	Event *PerpetualWithdraw // Event containing the contract specifics and raw log

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
func (it *PerpetualWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpetualWithdraw)
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
		it.Event = new(PerpetualWithdraw)
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
func (it *PerpetualWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpetualWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpetualWithdraw represents a Withdraw event raised by the Perpetual contract.
type PerpetualWithdraw struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x2d3cfd22ce461d7eafde7bb13c6af0c0d5ed08406a59166e093b4354cfd94ae2.
//
// Solidity: event Withdraw(address , int256 )
func (_Perpetual *PerpetualFilterer) FilterWithdraw(opts *bind.FilterOpts) (*PerpetualWithdrawIterator, error) {

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &PerpetualWithdrawIterator{contract: _Perpetual.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x2d3cfd22ce461d7eafde7bb13c6af0c0d5ed08406a59166e093b4354cfd94ae2.
//
// Solidity: event Withdraw(address , int256 )
func (_Perpetual *PerpetualFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PerpetualWithdraw) (event.Subscription, error) {

	logs, sub, err := _Perpetual.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpetualWithdraw)
				if err := _Perpetual.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x2d3cfd22ce461d7eafde7bb13c6af0c0d5ed08406a59166e093b4354cfd94ae2.
//
// Solidity: event Withdraw(address , int256 )
func (_Perpetual *PerpetualFilterer) ParseWithdraw(log types.Log) (*PerpetualWithdraw, error) {
	event := new(PerpetualWithdraw)
	if err := _Perpetual.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
