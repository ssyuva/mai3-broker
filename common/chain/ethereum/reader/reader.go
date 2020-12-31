// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package reader

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

// MarginAccount is an auto generated low-level Go binding around an user-defined struct.
type MarginAccount struct {
	Cash     *big.Int
	Position *big.Int
}

// ReaderLiquidityPoolStorage is an auto generated low-level Go binding around an user-defined struct.
type ReaderLiquidityPoolStorage struct {
	Operator          common.Address
	CollateralToken   common.Address
	Vault             common.Address
	Governor          common.Address
	ShareToken        common.Address
	VaultFeeRate      *big.Int
	PoolCash          *big.Int
	FundingTime       *big.Int
	PerpetualStorages []ReaderPerpetualStorage
}

// ReaderPerpetualStorage is an auto generated low-level Go binding around an user-defined struct.
type ReaderPerpetualStorage struct {
	Symbol                  *big.Int
	UnderlyingAsset         string
	State                   uint8
	Oracle                  common.Address
	MarkPrice               *big.Int
	IndexPrice              *big.Int
	UnitAccumulativeFunding *big.Int
	InitialMarginRate       *big.Int
	MaintenanceMarginRate   *big.Int
	OperatorFeeRate         *big.Int
	LpFeeRate               *big.Int
	ReferrerRebateRate      *big.Int
	LiquidationPenaltyRate  *big.Int
	KeeperGasReward         *big.Int
	InsuranceFundRate       *big.Int
	InsuranceFundCap        *big.Int
	InsuranceFund           *big.Int
	DonatedInsuranceFund    *big.Int
	HalfSpread              *big.Int
	OpenSlippageFactor      *big.Int
	CloseSlippageFactor     *big.Int
	FundingRateLimit        *big.Int
	AmmMaxLeverage          *big.Int
	AmmCashBalance          *big.Int
	AmmPositionAmount       *big.Int
}

// ReaderABI is the input ABI used to generate the binding from.
const ReaderABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountStorage\",\"outputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"cash\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"position\",\"type\":\"int256\"}],\"internalType\":\"structMarginAccount\",\"name\":\"marginAccount\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidityPool\",\"type\":\"address\"}],\"name\":\"getLiquidityPoolStorage\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"shareToken\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"vaultFeeRate\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"poolCash\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"fundingTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"symbol\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"underlyingAsset\",\"type\":\"string\"},{\"internalType\":\"enumPerpetualState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"markPrice\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"indexPrice\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"unitAccumulativeFunding\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"initialMarginRate\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"maintenanceMarginRate\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"operatorFeeRate\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"lpFeeRate\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"referrerRebateRate\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"liquidationPenaltyRate\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"keeperGasReward\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"insuranceFundRate\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"insuranceFundCap\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"insuranceFund\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"donatedInsuranceFund\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"halfSpread\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"openSlippageFactor\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"closeSlippageFactor\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"fundingRateLimit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ammMaxLeverage\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ammCashBalance\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ammPositionAmount\",\"type\":\"int256\"}],\"internalType\":\"structReader.PerpetualStorage[]\",\"name\":\"perpetualStorages\",\"type\":\"tuple[]\"}],\"internalType\":\"structReader.LiquidityPoolStorage\",\"name\":\"pool\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Reader is an auto generated Go binding around an Ethereum contract.
type Reader struct {
	ReaderCaller     // Read-only binding to the contract
	ReaderTransactor // Write-only binding to the contract
	ReaderFilterer   // Log filterer for contract events
}

// ReaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReaderSession struct {
	Contract     *Reader           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReaderCallerSession struct {
	Contract *ReaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ReaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReaderTransactorSession struct {
	Contract     *ReaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReaderRaw struct {
	Contract *Reader // Generic contract binding to access the raw methods on
}

// ReaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReaderCallerRaw struct {
	Contract *ReaderCaller // Generic read-only contract binding to access the raw methods on
}

// ReaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReaderTransactorRaw struct {
	Contract *ReaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReader creates a new instance of Reader, bound to a specific deployed contract.
func NewReader(address common.Address, backend bind.ContractBackend) (*Reader, error) {
	contract, err := bindReader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reader{ReaderCaller: ReaderCaller{contract: contract}, ReaderTransactor: ReaderTransactor{contract: contract}, ReaderFilterer: ReaderFilterer{contract: contract}}, nil
}

// NewReaderCaller creates a new read-only instance of Reader, bound to a specific deployed contract.
func NewReaderCaller(address common.Address, caller bind.ContractCaller) (*ReaderCaller, error) {
	contract, err := bindReader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReaderCaller{contract: contract}, nil
}

// NewReaderTransactor creates a new write-only instance of Reader, bound to a specific deployed contract.
func NewReaderTransactor(address common.Address, transactor bind.ContractTransactor) (*ReaderTransactor, error) {
	contract, err := bindReader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReaderTransactor{contract: contract}, nil
}

// NewReaderFilterer creates a new log filterer instance of Reader, bound to a specific deployed contract.
func NewReaderFilterer(address common.Address, filterer bind.ContractFilterer) (*ReaderFilterer, error) {
	contract, err := bindReader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReaderFilterer{contract: contract}, nil
}

// bindReader binds a generic wrapper to an already deployed contract.
func bindReader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReaderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reader *ReaderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Reader.Contract.ReaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reader *ReaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reader.Contract.ReaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reader *ReaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reader.Contract.ReaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reader *ReaderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Reader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reader *ReaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reader *ReaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reader.Contract.contract.Transact(opts, method, params...)
}

// GetAccountStorage is a free data retrieval call binding the contract method 0xeb16510d.
//
// Solidity: function getAccountStorage(address liquidityPool, uint256 perpetualIndex, address account) constant returns(MarginAccount marginAccount)
func (_Reader *ReaderCaller) GetAccountStorage(opts *bind.CallOpts, liquidityPool common.Address, perpetualIndex *big.Int, account common.Address) (MarginAccount, error) {
	var (
		ret0 = new(MarginAccount)
	)
	out := ret0
	err := _Reader.contract.Call(opts, out, "getAccountStorage", liquidityPool, perpetualIndex, account)
	return *ret0, err
}

// GetAccountStorage is a free data retrieval call binding the contract method 0xeb16510d.
//
// Solidity: function getAccountStorage(address liquidityPool, uint256 perpetualIndex, address account) constant returns(MarginAccount marginAccount)
func (_Reader *ReaderSession) GetAccountStorage(liquidityPool common.Address, perpetualIndex *big.Int, account common.Address) (MarginAccount, error) {
	return _Reader.Contract.GetAccountStorage(&_Reader.CallOpts, liquidityPool, perpetualIndex, account)
}

// GetAccountStorage is a free data retrieval call binding the contract method 0xeb16510d.
//
// Solidity: function getAccountStorage(address liquidityPool, uint256 perpetualIndex, address account) constant returns(MarginAccount marginAccount)
func (_Reader *ReaderCallerSession) GetAccountStorage(liquidityPool common.Address, perpetualIndex *big.Int, account common.Address) (MarginAccount, error) {
	return _Reader.Contract.GetAccountStorage(&_Reader.CallOpts, liquidityPool, perpetualIndex, account)
}

// GetLiquidityPoolStorage is a free data retrieval call binding the contract method 0x574408c1.
//
// Solidity: function getLiquidityPoolStorage(address liquidityPool) constant returns(ReaderLiquidityPoolStorage pool)
func (_Reader *ReaderCaller) GetLiquidityPoolStorage(opts *bind.CallOpts, liquidityPool common.Address) (ReaderLiquidityPoolStorage, error) {
	var (
		ret0 = new(ReaderLiquidityPoolStorage)
	)
	out := ret0
	err := _Reader.contract.Call(opts, out, "getLiquidityPoolStorage", liquidityPool)
	return *ret0, err
}

// GetLiquidityPoolStorage is a free data retrieval call binding the contract method 0x574408c1.
//
// Solidity: function getLiquidityPoolStorage(address liquidityPool) constant returns(ReaderLiquidityPoolStorage pool)
func (_Reader *ReaderSession) GetLiquidityPoolStorage(liquidityPool common.Address) (ReaderLiquidityPoolStorage, error) {
	return _Reader.Contract.GetLiquidityPoolStorage(&_Reader.CallOpts, liquidityPool)
}

// GetLiquidityPoolStorage is a free data retrieval call binding the contract method 0x574408c1.
//
// Solidity: function getLiquidityPoolStorage(address liquidityPool) constant returns(ReaderLiquidityPoolStorage pool)
func (_Reader *ReaderCallerSession) GetLiquidityPoolStorage(liquidityPool common.Address) (ReaderLiquidityPoolStorage, error) {
	return _Reader.Contract.GetLiquidityPoolStorage(&_Reader.CallOpts, liquidityPool)
}