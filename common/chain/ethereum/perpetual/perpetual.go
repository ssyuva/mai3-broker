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

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Trader      common.Address
	Broker      common.Address
	Relayer     common.Address
	Perpetual   common.Address
	Referrer    common.Address
	Amount      *big.Int
	PriceLimit  *big.Int
	Deadline    uint64
	Version     uint32
	OrderType   uint8
	IsCloseOnly bool
	Salt        uint64
	ChainID     *big.Int
}

// PerpetualABI is the input ABI used to generate the binding from.
const PerpetualABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"addedCash\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"mintedShare\",\"type\":\"int256\"}],\"name\":\"AddLiquidatity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"AdjustRiskSetting\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"ClaimFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"Clear\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fundingLoss\",\"type\":\"int256\"}],\"name\":\"ClosePositionByLiquidation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fundingLoss\",\"type\":\"int256\"}],\"name\":\"ClosePositionByTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"DonateInsuranceFund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"privilege\",\"type\":\"uint256\"}],\"name\":\"GrantPrivilege\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fee\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"LiquidateByAMM\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"LiquidateByTrader\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"}],\"name\":\"OpenPositionByLiquidation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"}],\"name\":\"OpenPositionByTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"returnedCash\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"burnedShare\",\"type\":\"int256\"}],\"name\":\"RemoveLiquidatity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"privilege\",\"type\":\"uint256\"}],\"name\":\"RevokePrivilege\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"positionAmount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"priceLimit\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fee\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"UpdateCoreSetting\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"minValue\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"maxValue\",\"type\":\"int256\"}],\"name\":\"UpdateRiskSetting\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"cashToAdd\",\"type\":\"int256\"}],\"name\":\"addLiquidatity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"newValue\",\"type\":\"int256\"}],\"name\":\"adjustRiskParameter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"availableMargin\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"broker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"perpetual\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"priceLimit\",\"type\":\"int256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"version\",\"type\":\"uint32\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isCloseOnly\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"salt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"brokerTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"claimFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"claimableFee\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"clear\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"donateInsuranceFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundingState\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"unitAccumulativeFunding\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"fundingRate\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"fundingTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"privilege\",\"type\":\"uint256\"}],\"name\":\"grantPrivilege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"information\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"underlyingAsset\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"int256[8]\",\"name\":\"coreParameter\",\"type\":\"int256[8]\"},{\"internalType\":\"int256[5]\",\"name\":\"riskParameter\",\"type\":\"int256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"shareToken\",\"type\":\"address\"},{\"internalType\":\"int256[7]\",\"name\":\"coreParams\",\"type\":\"int256[7]\"},{\"internalType\":\"int256[5]\",\"name\":\"riskParams\",\"type\":\"int256[5]\"},{\"internalType\":\"int256[5]\",\"name\":\"minRiskParamValues\",\"type\":\"int256[5]\"},{\"internalType\":\"int256[5]\",\"name\":\"maxRiskParamValues\",\"type\":\"int256[5]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"privilege\",\"type\":\"uint256\"}],\"name\":\"isGranted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"liquidateByAMM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"priceLimit\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"liquidateByTrader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"margin\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"marginAccount\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"positionAmount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"cashBalance\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"entryFundingLoss\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"shareToRemove\",\"type\":\"int256\"}],\"name\":\"removeLiquidatity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"privilege\",\"type\":\"uint256\"}],\"name\":\"revokePrivilege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"settle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shareToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shutdown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isEmergency\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isShuttingdown\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"insuranceFund\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"donatedInsuranceFund\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"markPrice\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"indexPrice\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"priceLimit\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"trade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"newValue\",\"type\":\"int256\"}],\"name\":\"updateCoreParameter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"newValue\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"minValue\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"maxValue\",\"type\":\"int256\"}],\"name\":\"updateRiskParameter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"withdrawableMargin\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"withdrawable\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// ClaimableFee is a free data retrieval call binding the contract method 0xa1c59afc.
//
// Solidity: function claimableFee(address claimer) constant returns(int256)
func (_Perpetual *PerpetualCaller) ClaimableFee(opts *bind.CallOpts, claimer common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Perpetual.contract.Call(opts, out, "claimableFee", claimer)
	return *ret0, err
}

// ClaimableFee is a free data retrieval call binding the contract method 0xa1c59afc.
//
// Solidity: function claimableFee(address claimer) constant returns(int256)
func (_Perpetual *PerpetualSession) ClaimableFee(claimer common.Address) (*big.Int, error) {
	return _Perpetual.Contract.ClaimableFee(&_Perpetual.CallOpts, claimer)
}

// ClaimableFee is a free data retrieval call binding the contract method 0xa1c59afc.
//
// Solidity: function claimableFee(address claimer) constant returns(int256)
func (_Perpetual *PerpetualCallerSession) ClaimableFee(claimer common.Address) (*big.Int, error) {
	return _Perpetual.Contract.ClaimableFee(&_Perpetual.CallOpts, claimer)
}

// FundingState is a free data retrieval call binding the contract method 0xb25cdccf.
//
// Solidity: function fundingState() constant returns(int256 unitAccumulativeFunding, int256 fundingRate, uint256 fundingTime)
func (_Perpetual *PerpetualCaller) FundingState(opts *bind.CallOpts) (struct {
	UnitAccumulativeFunding *big.Int
	FundingRate             *big.Int
	FundingTime             *big.Int
}, error) {
	ret := new(struct {
		UnitAccumulativeFunding *big.Int
		FundingRate             *big.Int
		FundingTime             *big.Int
	})
	out := ret
	err := _Perpetual.contract.Call(opts, out, "fundingState")
	return *ret, err
}

// FundingState is a free data retrieval call binding the contract method 0xb25cdccf.
//
// Solidity: function fundingState() constant returns(int256 unitAccumulativeFunding, int256 fundingRate, uint256 fundingTime)
func (_Perpetual *PerpetualSession) FundingState() (struct {
	UnitAccumulativeFunding *big.Int
	FundingRate             *big.Int
	FundingTime             *big.Int
}, error) {
	return _Perpetual.Contract.FundingState(&_Perpetual.CallOpts)
}

// FundingState is a free data retrieval call binding the contract method 0xb25cdccf.
//
// Solidity: function fundingState() constant returns(int256 unitAccumulativeFunding, int256 fundingRate, uint256 fundingTime)
func (_Perpetual *PerpetualCallerSession) FundingState() (struct {
	UnitAccumulativeFunding *big.Int
	FundingRate             *big.Int
	FundingTime             *big.Int
}, error) {
	return _Perpetual.Contract.FundingState(&_Perpetual.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() constant returns(address)
func (_Perpetual *PerpetualCaller) Governor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Perpetual.contract.Call(opts, out, "governor")
	return *ret0, err
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() constant returns(address)
func (_Perpetual *PerpetualSession) Governor() (common.Address, error) {
	return _Perpetual.Contract.Governor(&_Perpetual.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() constant returns(address)
func (_Perpetual *PerpetualCallerSession) Governor() (common.Address, error) {
	return _Perpetual.Contract.Governor(&_Perpetual.CallOpts)
}

// Information is a free data retrieval call binding the contract method 0xf47b7740.
//
// Solidity: function information() constant returns(string underlyingAsset, address collateral, address factory, address oracle, address operator, address vault, int256[8] coreParameter, int256[5] riskParameter)
func (_Perpetual *PerpetualCaller) Information(opts *bind.CallOpts) (struct {
	UnderlyingAsset string
	Collateral      common.Address
	Factory         common.Address
	Oracle          common.Address
	Operator        common.Address
	Vault           common.Address
	CoreParameter   [8]*big.Int
	RiskParameter   [5]*big.Int
}, error) {
	ret := new(struct {
		UnderlyingAsset string
		Collateral      common.Address
		Factory         common.Address
		Oracle          common.Address
		Operator        common.Address
		Vault           common.Address
		CoreParameter   [8]*big.Int
		RiskParameter   [5]*big.Int
	})
	out := ret
	err := _Perpetual.contract.Call(opts, out, "information")
	return *ret, err
}

// Information is a free data retrieval call binding the contract method 0xf47b7740.
//
// Solidity: function information() constant returns(string underlyingAsset, address collateral, address factory, address oracle, address operator, address vault, int256[8] coreParameter, int256[5] riskParameter)
func (_Perpetual *PerpetualSession) Information() (struct {
	UnderlyingAsset string
	Collateral      common.Address
	Factory         common.Address
	Oracle          common.Address
	Operator        common.Address
	Vault           common.Address
	CoreParameter   [8]*big.Int
	RiskParameter   [5]*big.Int
}, error) {
	return _Perpetual.Contract.Information(&_Perpetual.CallOpts)
}

// Information is a free data retrieval call binding the contract method 0xf47b7740.
//
// Solidity: function information() constant returns(string underlyingAsset, address collateral, address factory, address oracle, address operator, address vault, int256[8] coreParameter, int256[5] riskParameter)
func (_Perpetual *PerpetualCallerSession) Information() (struct {
	UnderlyingAsset string
	Collateral      common.Address
	Factory         common.Address
	Oracle          common.Address
	Operator        common.Address
	Vault           common.Address
	CoreParameter   [8]*big.Int
	RiskParameter   [5]*big.Int
}, error) {
	return _Perpetual.Contract.Information(&_Perpetual.CallOpts)
}

// IsGranted is a free data retrieval call binding the contract method 0xc2f3da1d.
//
// Solidity: function isGranted(address owner, address trader, uint256 privilege) constant returns(bool)
func (_Perpetual *PerpetualCaller) IsGranted(opts *bind.CallOpts, owner common.Address, trader common.Address, privilege *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Perpetual.contract.Call(opts, out, "isGranted", owner, trader, privilege)
	return *ret0, err
}

// IsGranted is a free data retrieval call binding the contract method 0xc2f3da1d.
//
// Solidity: function isGranted(address owner, address trader, uint256 privilege) constant returns(bool)
func (_Perpetual *PerpetualSession) IsGranted(owner common.Address, trader common.Address, privilege *big.Int) (bool, error) {
	return _Perpetual.Contract.IsGranted(&_Perpetual.CallOpts, owner, trader, privilege)
}

// IsGranted is a free data retrieval call binding the contract method 0xc2f3da1d.
//
// Solidity: function isGranted(address owner, address trader, uint256 privilege) constant returns(bool)
func (_Perpetual *PerpetualCallerSession) IsGranted(owner common.Address, trader common.Address, privilege *big.Int) (bool, error) {
	return _Perpetual.Contract.IsGranted(&_Perpetual.CallOpts, owner, trader, privilege)
}

// MarginAccount is a free data retrieval call binding the contract method 0x6280e472.
//
// Solidity: function marginAccount(address trader) constant returns(int256 positionAmount, int256 cashBalance, int256 entryFundingLoss)
func (_Perpetual *PerpetualCaller) MarginAccount(opts *bind.CallOpts, trader common.Address) (struct {
	PositionAmount   *big.Int
	CashBalance      *big.Int
	EntryFundingLoss *big.Int
}, error) {
	ret := new(struct {
		PositionAmount   *big.Int
		CashBalance      *big.Int
		EntryFundingLoss *big.Int
	})
	out := ret
	err := _Perpetual.contract.Call(opts, out, "marginAccount", trader)
	return *ret, err
}

// MarginAccount is a free data retrieval call binding the contract method 0x6280e472.
//
// Solidity: function marginAccount(address trader) constant returns(int256 positionAmount, int256 cashBalance, int256 entryFundingLoss)
func (_Perpetual *PerpetualSession) MarginAccount(trader common.Address) (struct {
	PositionAmount   *big.Int
	CashBalance      *big.Int
	EntryFundingLoss *big.Int
}, error) {
	return _Perpetual.Contract.MarginAccount(&_Perpetual.CallOpts, trader)
}

// MarginAccount is a free data retrieval call binding the contract method 0x6280e472.
//
// Solidity: function marginAccount(address trader) constant returns(int256 positionAmount, int256 cashBalance, int256 entryFundingLoss)
func (_Perpetual *PerpetualCallerSession) MarginAccount(trader common.Address) (struct {
	PositionAmount   *big.Int
	CashBalance      *big.Int
	EntryFundingLoss *big.Int
}, error) {
	return _Perpetual.Contract.MarginAccount(&_Perpetual.CallOpts, trader)
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() constant returns(address)
func (_Perpetual *PerpetualCaller) ShareToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Perpetual.contract.Call(opts, out, "shareToken")
	return *ret0, err
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() constant returns(address)
func (_Perpetual *PerpetualSession) ShareToken() (common.Address, error) {
	return _Perpetual.Contract.ShareToken(&_Perpetual.CallOpts)
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() constant returns(address)
func (_Perpetual *PerpetualCallerSession) ShareToken() (common.Address, error) {
	return _Perpetual.Contract.ShareToken(&_Perpetual.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(bool isEmergency, bool isShuttingdown, int256 insuranceFund, int256 donatedInsuranceFund, int256 markPrice, int256 indexPrice)
func (_Perpetual *PerpetualCaller) State(opts *bind.CallOpts) (struct {
	IsEmergency          bool
	IsShuttingdown       bool
	InsuranceFund        *big.Int
	DonatedInsuranceFund *big.Int
	MarkPrice            *big.Int
	IndexPrice           *big.Int
}, error) {
	ret := new(struct {
		IsEmergency          bool
		IsShuttingdown       bool
		InsuranceFund        *big.Int
		DonatedInsuranceFund *big.Int
		MarkPrice            *big.Int
		IndexPrice           *big.Int
	})
	out := ret
	err := _Perpetual.contract.Call(opts, out, "state")
	return *ret, err
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(bool isEmergency, bool isShuttingdown, int256 insuranceFund, int256 donatedInsuranceFund, int256 markPrice, int256 indexPrice)
func (_Perpetual *PerpetualSession) State() (struct {
	IsEmergency          bool
	IsShuttingdown       bool
	InsuranceFund        *big.Int
	DonatedInsuranceFund *big.Int
	MarkPrice            *big.Int
	IndexPrice           *big.Int
}, error) {
	return _Perpetual.Contract.State(&_Perpetual.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(bool isEmergency, bool isShuttingdown, int256 insuranceFund, int256 donatedInsuranceFund, int256 markPrice, int256 indexPrice)
func (_Perpetual *PerpetualCallerSession) State() (struct {
	IsEmergency          bool
	IsShuttingdown       bool
	InsuranceFund        *big.Int
	DonatedInsuranceFund *big.Int
	MarkPrice            *big.Int
	IndexPrice           *big.Int
}, error) {
	return _Perpetual.Contract.State(&_Perpetual.CallOpts)
}

// AddLiquidatity is a paid mutator transaction binding the contract method 0x9328f40b.
//
// Solidity: function addLiquidatity(int256 cashToAdd) returns()
func (_Perpetual *PerpetualTransactor) AddLiquidatity(opts *bind.TransactOpts, cashToAdd *big.Int) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "addLiquidatity", cashToAdd)
}

// AddLiquidatity is a paid mutator transaction binding the contract method 0x9328f40b.
//
// Solidity: function addLiquidatity(int256 cashToAdd) returns()
func (_Perpetual *PerpetualSession) AddLiquidatity(cashToAdd *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.AddLiquidatity(&_Perpetual.TransactOpts, cashToAdd)
}

// AddLiquidatity is a paid mutator transaction binding the contract method 0x9328f40b.
//
// Solidity: function addLiquidatity(int256 cashToAdd) returns()
func (_Perpetual *PerpetualTransactorSession) AddLiquidatity(cashToAdd *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.AddLiquidatity(&_Perpetual.TransactOpts, cashToAdd)
}

// AdjustRiskParameter is a paid mutator transaction binding the contract method 0x13ff3b2f.
//
// Solidity: function adjustRiskParameter(bytes32 key, int256 newValue) returns()
func (_Perpetual *PerpetualTransactor) AdjustRiskParameter(opts *bind.TransactOpts, key [32]byte, newValue *big.Int) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "adjustRiskParameter", key, newValue)
}

// AdjustRiskParameter is a paid mutator transaction binding the contract method 0x13ff3b2f.
//
// Solidity: function adjustRiskParameter(bytes32 key, int256 newValue) returns()
func (_Perpetual *PerpetualSession) AdjustRiskParameter(key [32]byte, newValue *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.AdjustRiskParameter(&_Perpetual.TransactOpts, key, newValue)
}

// AdjustRiskParameter is a paid mutator transaction binding the contract method 0x13ff3b2f.
//
// Solidity: function adjustRiskParameter(bytes32 key, int256 newValue) returns()
func (_Perpetual *PerpetualTransactorSession) AdjustRiskParameter(key [32]byte, newValue *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.AdjustRiskParameter(&_Perpetual.TransactOpts, key, newValue)
}

// AvailableMargin is a paid mutator transaction binding the contract method 0x711d4d55.
//
// Solidity: function availableMargin(address trader) returns(int256)
func (_Perpetual *PerpetualTransactor) AvailableMargin(opts *bind.TransactOpts, trader common.Address) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "availableMargin", trader)
}

// AvailableMargin is a paid mutator transaction binding the contract method 0x711d4d55.
//
// Solidity: function availableMargin(address trader) returns(int256)
func (_Perpetual *PerpetualSession) AvailableMargin(trader common.Address) (*types.Transaction, error) {
	return _Perpetual.Contract.AvailableMargin(&_Perpetual.TransactOpts, trader)
}

// AvailableMargin is a paid mutator transaction binding the contract method 0x711d4d55.
//
// Solidity: function availableMargin(address trader) returns(int256)
func (_Perpetual *PerpetualTransactorSession) AvailableMargin(trader common.Address) (*types.Transaction, error) {
	return _Perpetual.Contract.AvailableMargin(&_Perpetual.TransactOpts, trader)
}

// BrokerTrade is a paid mutator transaction binding the contract method 0xd82853bc.
//
// Solidity: function brokerTrade(Order order, int256 amount, bytes signature) returns()
func (_Perpetual *PerpetualTransactor) BrokerTrade(opts *bind.TransactOpts, order Order, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "brokerTrade", order, amount, signature)
}

// BrokerTrade is a paid mutator transaction binding the contract method 0xd82853bc.
//
// Solidity: function brokerTrade(Order order, int256 amount, bytes signature) returns()
func (_Perpetual *PerpetualSession) BrokerTrade(order Order, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Perpetual.Contract.BrokerTrade(&_Perpetual.TransactOpts, order, amount, signature)
}

// BrokerTrade is a paid mutator transaction binding the contract method 0xd82853bc.
//
// Solidity: function brokerTrade(Order order, int256 amount, bytes signature) returns()
func (_Perpetual *PerpetualTransactorSession) BrokerTrade(order Order, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Perpetual.Contract.BrokerTrade(&_Perpetual.TransactOpts, order, amount, signature)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x28846269.
//
// Solidity: function claimFee(address claimer, int256 amount) returns()
func (_Perpetual *PerpetualTransactor) ClaimFee(opts *bind.TransactOpts, claimer common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "claimFee", claimer, amount)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x28846269.
//
// Solidity: function claimFee(address claimer, int256 amount) returns()
func (_Perpetual *PerpetualSession) ClaimFee(claimer common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.ClaimFee(&_Perpetual.TransactOpts, claimer, amount)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x28846269.
//
// Solidity: function claimFee(address claimer, int256 amount) returns()
func (_Perpetual *PerpetualTransactorSession) ClaimFee(claimer common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.ClaimFee(&_Perpetual.TransactOpts, claimer, amount)
}

// Clear is a paid mutator transaction binding the contract method 0x3d0a4061.
//
// Solidity: function clear(address trader) returns()
func (_Perpetual *PerpetualTransactor) Clear(opts *bind.TransactOpts, trader common.Address) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "clear", trader)
}

// Clear is a paid mutator transaction binding the contract method 0x3d0a4061.
//
// Solidity: function clear(address trader) returns()
func (_Perpetual *PerpetualSession) Clear(trader common.Address) (*types.Transaction, error) {
	return _Perpetual.Contract.Clear(&_Perpetual.TransactOpts, trader)
}

// Clear is a paid mutator transaction binding the contract method 0x3d0a4061.
//
// Solidity: function clear(address trader) returns()
func (_Perpetual *PerpetualTransactorSession) Clear(trader common.Address) (*types.Transaction, error) {
	return _Perpetual.Contract.Clear(&_Perpetual.TransactOpts, trader)
}

// Deposit is a paid mutator transaction binding the contract method 0x1a8dea36.
//
// Solidity: function deposit(address trader, int256 amount) returns()
func (_Perpetual *PerpetualTransactor) Deposit(opts *bind.TransactOpts, trader common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "deposit", trader, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x1a8dea36.
//
// Solidity: function deposit(address trader, int256 amount) returns()
func (_Perpetual *PerpetualSession) Deposit(trader common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.Deposit(&_Perpetual.TransactOpts, trader, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x1a8dea36.
//
// Solidity: function deposit(address trader, int256 amount) returns()
func (_Perpetual *PerpetualTransactorSession) Deposit(trader common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.Deposit(&_Perpetual.TransactOpts, trader, amount)
}

// DonateInsuranceFund is a paid mutator transaction binding the contract method 0x6fca8b99.
//
// Solidity: function donateInsuranceFund(int256 amount) returns()
func (_Perpetual *PerpetualTransactor) DonateInsuranceFund(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "donateInsuranceFund", amount)
}

// DonateInsuranceFund is a paid mutator transaction binding the contract method 0x6fca8b99.
//
// Solidity: function donateInsuranceFund(int256 amount) returns()
func (_Perpetual *PerpetualSession) DonateInsuranceFund(amount *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.DonateInsuranceFund(&_Perpetual.TransactOpts, amount)
}

// DonateInsuranceFund is a paid mutator transaction binding the contract method 0x6fca8b99.
//
// Solidity: function donateInsuranceFund(int256 amount) returns()
func (_Perpetual *PerpetualTransactorSession) DonateInsuranceFund(amount *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.DonateInsuranceFund(&_Perpetual.TransactOpts, amount)
}

// GrantPrivilege is a paid mutator transaction binding the contract method 0x7063a24b.
//
// Solidity: function grantPrivilege(address owner, address trader, uint256 privilege) returns()
func (_Perpetual *PerpetualTransactor) GrantPrivilege(opts *bind.TransactOpts, owner common.Address, trader common.Address, privilege *big.Int) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "grantPrivilege", owner, trader, privilege)
}

// GrantPrivilege is a paid mutator transaction binding the contract method 0x7063a24b.
//
// Solidity: function grantPrivilege(address owner, address trader, uint256 privilege) returns()
func (_Perpetual *PerpetualSession) GrantPrivilege(owner common.Address, trader common.Address, privilege *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.GrantPrivilege(&_Perpetual.TransactOpts, owner, trader, privilege)
}

// GrantPrivilege is a paid mutator transaction binding the contract method 0x7063a24b.
//
// Solidity: function grantPrivilege(address owner, address trader, uint256 privilege) returns()
func (_Perpetual *PerpetualTransactorSession) GrantPrivilege(owner common.Address, trader common.Address, privilege *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.GrantPrivilege(&_Perpetual.TransactOpts, owner, trader, privilege)
}

// Initialize is a paid mutator transaction binding the contract method 0x49fcf496.
//
// Solidity: function initialize(address operator, address oracle, address governor, address shareToken, int256[7] coreParams, int256[5] riskParams, int256[5] minRiskParamValues, int256[5] maxRiskParamValues) returns()
func (_Perpetual *PerpetualTransactor) Initialize(opts *bind.TransactOpts, operator common.Address, oracle common.Address, governor common.Address, shareToken common.Address, coreParams [7]*big.Int, riskParams [5]*big.Int, minRiskParamValues [5]*big.Int, maxRiskParamValues [5]*big.Int) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "initialize", operator, oracle, governor, shareToken, coreParams, riskParams, minRiskParamValues, maxRiskParamValues)
}

// Initialize is a paid mutator transaction binding the contract method 0x49fcf496.
//
// Solidity: function initialize(address operator, address oracle, address governor, address shareToken, int256[7] coreParams, int256[5] riskParams, int256[5] minRiskParamValues, int256[5] maxRiskParamValues) returns()
func (_Perpetual *PerpetualSession) Initialize(operator common.Address, oracle common.Address, governor common.Address, shareToken common.Address, coreParams [7]*big.Int, riskParams [5]*big.Int, minRiskParamValues [5]*big.Int, maxRiskParamValues [5]*big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.Initialize(&_Perpetual.TransactOpts, operator, oracle, governor, shareToken, coreParams, riskParams, minRiskParamValues, maxRiskParamValues)
}

// Initialize is a paid mutator transaction binding the contract method 0x49fcf496.
//
// Solidity: function initialize(address operator, address oracle, address governor, address shareToken, int256[7] coreParams, int256[5] riskParams, int256[5] minRiskParamValues, int256[5] maxRiskParamValues) returns()
func (_Perpetual *PerpetualTransactorSession) Initialize(operator common.Address, oracle common.Address, governor common.Address, shareToken common.Address, coreParams [7]*big.Int, riskParams [5]*big.Int, minRiskParamValues [5]*big.Int, maxRiskParamValues [5]*big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.Initialize(&_Perpetual.TransactOpts, operator, oracle, governor, shareToken, coreParams, riskParams, minRiskParamValues, maxRiskParamValues)
}

// LiquidateByAMM is a paid mutator transaction binding the contract method 0x16da442c.
//
// Solidity: function liquidateByAMM(address trader, uint256 deadline) returns()
func (_Perpetual *PerpetualTransactor) LiquidateByAMM(opts *bind.TransactOpts, trader common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "liquidateByAMM", trader, deadline)
}

// LiquidateByAMM is a paid mutator transaction binding the contract method 0x16da442c.
//
// Solidity: function liquidateByAMM(address trader, uint256 deadline) returns()
func (_Perpetual *PerpetualSession) LiquidateByAMM(trader common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.LiquidateByAMM(&_Perpetual.TransactOpts, trader, deadline)
}

// LiquidateByAMM is a paid mutator transaction binding the contract method 0x16da442c.
//
// Solidity: function liquidateByAMM(address trader, uint256 deadline) returns()
func (_Perpetual *PerpetualTransactorSession) LiquidateByAMM(trader common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.LiquidateByAMM(&_Perpetual.TransactOpts, trader, deadline)
}

// LiquidateByTrader is a paid mutator transaction binding the contract method 0x6c5ce25f.
//
// Solidity: function liquidateByTrader(address trader, int256 amount, int256 priceLimit, uint256 deadline) returns()
func (_Perpetual *PerpetualTransactor) LiquidateByTrader(opts *bind.TransactOpts, trader common.Address, amount *big.Int, priceLimit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "liquidateByTrader", trader, amount, priceLimit, deadline)
}

// LiquidateByTrader is a paid mutator transaction binding the contract method 0x6c5ce25f.
//
// Solidity: function liquidateByTrader(address trader, int256 amount, int256 priceLimit, uint256 deadline) returns()
func (_Perpetual *PerpetualSession) LiquidateByTrader(trader common.Address, amount *big.Int, priceLimit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.LiquidateByTrader(&_Perpetual.TransactOpts, trader, amount, priceLimit, deadline)
}

// LiquidateByTrader is a paid mutator transaction binding the contract method 0x6c5ce25f.
//
// Solidity: function liquidateByTrader(address trader, int256 amount, int256 priceLimit, uint256 deadline) returns()
func (_Perpetual *PerpetualTransactorSession) LiquidateByTrader(trader common.Address, amount *big.Int, priceLimit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.LiquidateByTrader(&_Perpetual.TransactOpts, trader, amount, priceLimit, deadline)
}

// Margin is a paid mutator transaction binding the contract method 0xbeb3ed5d.
//
// Solidity: function margin(address trader) returns(int256)
func (_Perpetual *PerpetualTransactor) Margin(opts *bind.TransactOpts, trader common.Address) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "margin", trader)
}

// Margin is a paid mutator transaction binding the contract method 0xbeb3ed5d.
//
// Solidity: function margin(address trader) returns(int256)
func (_Perpetual *PerpetualSession) Margin(trader common.Address) (*types.Transaction, error) {
	return _Perpetual.Contract.Margin(&_Perpetual.TransactOpts, trader)
}

// Margin is a paid mutator transaction binding the contract method 0xbeb3ed5d.
//
// Solidity: function margin(address trader) returns(int256)
func (_Perpetual *PerpetualTransactorSession) Margin(trader common.Address) (*types.Transaction, error) {
	return _Perpetual.Contract.Margin(&_Perpetual.TransactOpts, trader)
}

// RemoveLiquidatity is a paid mutator transaction binding the contract method 0x02938dd4.
//
// Solidity: function removeLiquidatity(int256 shareToRemove) returns()
func (_Perpetual *PerpetualTransactor) RemoveLiquidatity(opts *bind.TransactOpts, shareToRemove *big.Int) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "removeLiquidatity", shareToRemove)
}

// RemoveLiquidatity is a paid mutator transaction binding the contract method 0x02938dd4.
//
// Solidity: function removeLiquidatity(int256 shareToRemove) returns()
func (_Perpetual *PerpetualSession) RemoveLiquidatity(shareToRemove *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.RemoveLiquidatity(&_Perpetual.TransactOpts, shareToRemove)
}

// RemoveLiquidatity is a paid mutator transaction binding the contract method 0x02938dd4.
//
// Solidity: function removeLiquidatity(int256 shareToRemove) returns()
func (_Perpetual *PerpetualTransactorSession) RemoveLiquidatity(shareToRemove *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.RemoveLiquidatity(&_Perpetual.TransactOpts, shareToRemove)
}

// RevokePrivilege is a paid mutator transaction binding the contract method 0x048b495a.
//
// Solidity: function revokePrivilege(address owner, address trader, uint256 privilege) returns()
func (_Perpetual *PerpetualTransactor) RevokePrivilege(opts *bind.TransactOpts, owner common.Address, trader common.Address, privilege *big.Int) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "revokePrivilege", owner, trader, privilege)
}

// RevokePrivilege is a paid mutator transaction binding the contract method 0x048b495a.
//
// Solidity: function revokePrivilege(address owner, address trader, uint256 privilege) returns()
func (_Perpetual *PerpetualSession) RevokePrivilege(owner common.Address, trader common.Address, privilege *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.RevokePrivilege(&_Perpetual.TransactOpts, owner, trader, privilege)
}

// RevokePrivilege is a paid mutator transaction binding the contract method 0x048b495a.
//
// Solidity: function revokePrivilege(address owner, address trader, uint256 privilege) returns()
func (_Perpetual *PerpetualTransactorSession) RevokePrivilege(owner common.Address, trader common.Address, privilege *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.RevokePrivilege(&_Perpetual.TransactOpts, owner, trader, privilege)
}

// Settle is a paid mutator transaction binding the contract method 0x6a256b29.
//
// Solidity: function settle(address trader) returns()
func (_Perpetual *PerpetualTransactor) Settle(opts *bind.TransactOpts, trader common.Address) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "settle", trader)
}

// Settle is a paid mutator transaction binding the contract method 0x6a256b29.
//
// Solidity: function settle(address trader) returns()
func (_Perpetual *PerpetualSession) Settle(trader common.Address) (*types.Transaction, error) {
	return _Perpetual.Contract.Settle(&_Perpetual.TransactOpts, trader)
}

// Settle is a paid mutator transaction binding the contract method 0x6a256b29.
//
// Solidity: function settle(address trader) returns()
func (_Perpetual *PerpetualTransactorSession) Settle(trader common.Address) (*types.Transaction, error) {
	return _Perpetual.Contract.Settle(&_Perpetual.TransactOpts, trader)
}

// Shutdown is a paid mutator transaction binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() returns()
func (_Perpetual *PerpetualTransactor) Shutdown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "shutdown")
}

// Shutdown is a paid mutator transaction binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() returns()
func (_Perpetual *PerpetualSession) Shutdown() (*types.Transaction, error) {
	return _Perpetual.Contract.Shutdown(&_Perpetual.TransactOpts)
}

// Shutdown is a paid mutator transaction binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() returns()
func (_Perpetual *PerpetualTransactorSession) Shutdown() (*types.Transaction, error) {
	return _Perpetual.Contract.Shutdown(&_Perpetual.TransactOpts)
}

// Trade is a paid mutator transaction binding the contract method 0xff7c29ec.
//
// Solidity: function trade(address trader, int256 amount, int256 priceLimit, uint256 deadline, address referrer) returns()
func (_Perpetual *PerpetualTransactor) Trade(opts *bind.TransactOpts, trader common.Address, amount *big.Int, priceLimit *big.Int, deadline *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "trade", trader, amount, priceLimit, deadline, referrer)
}

// Trade is a paid mutator transaction binding the contract method 0xff7c29ec.
//
// Solidity: function trade(address trader, int256 amount, int256 priceLimit, uint256 deadline, address referrer) returns()
func (_Perpetual *PerpetualSession) Trade(trader common.Address, amount *big.Int, priceLimit *big.Int, deadline *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Perpetual.Contract.Trade(&_Perpetual.TransactOpts, trader, amount, priceLimit, deadline, referrer)
}

// Trade is a paid mutator transaction binding the contract method 0xff7c29ec.
//
// Solidity: function trade(address trader, int256 amount, int256 priceLimit, uint256 deadline, address referrer) returns()
func (_Perpetual *PerpetualTransactorSession) Trade(trader common.Address, amount *big.Int, priceLimit *big.Int, deadline *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Perpetual.Contract.Trade(&_Perpetual.TransactOpts, trader, amount, priceLimit, deadline, referrer)
}

// UpdateCoreParameter is a paid mutator transaction binding the contract method 0x80bd474d.
//
// Solidity: function updateCoreParameter(bytes32 key, int256 newValue) returns()
func (_Perpetual *PerpetualTransactor) UpdateCoreParameter(opts *bind.TransactOpts, key [32]byte, newValue *big.Int) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "updateCoreParameter", key, newValue)
}

// UpdateCoreParameter is a paid mutator transaction binding the contract method 0x80bd474d.
//
// Solidity: function updateCoreParameter(bytes32 key, int256 newValue) returns()
func (_Perpetual *PerpetualSession) UpdateCoreParameter(key [32]byte, newValue *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.UpdateCoreParameter(&_Perpetual.TransactOpts, key, newValue)
}

// UpdateCoreParameter is a paid mutator transaction binding the contract method 0x80bd474d.
//
// Solidity: function updateCoreParameter(bytes32 key, int256 newValue) returns()
func (_Perpetual *PerpetualTransactorSession) UpdateCoreParameter(key [32]byte, newValue *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.UpdateCoreParameter(&_Perpetual.TransactOpts, key, newValue)
}

// UpdateRiskParameter is a paid mutator transaction binding the contract method 0x27111e67.
//
// Solidity: function updateRiskParameter(bytes32 key, int256 newValue, int256 minValue, int256 maxValue) returns()
func (_Perpetual *PerpetualTransactor) UpdateRiskParameter(opts *bind.TransactOpts, key [32]byte, newValue *big.Int, minValue *big.Int, maxValue *big.Int) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "updateRiskParameter", key, newValue, minValue, maxValue)
}

// UpdateRiskParameter is a paid mutator transaction binding the contract method 0x27111e67.
//
// Solidity: function updateRiskParameter(bytes32 key, int256 newValue, int256 minValue, int256 maxValue) returns()
func (_Perpetual *PerpetualSession) UpdateRiskParameter(key [32]byte, newValue *big.Int, minValue *big.Int, maxValue *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.UpdateRiskParameter(&_Perpetual.TransactOpts, key, newValue, minValue, maxValue)
}

// UpdateRiskParameter is a paid mutator transaction binding the contract method 0x27111e67.
//
// Solidity: function updateRiskParameter(bytes32 key, int256 newValue, int256 minValue, int256 maxValue) returns()
func (_Perpetual *PerpetualTransactorSession) UpdateRiskParameter(key [32]byte, newValue *big.Int, minValue *big.Int, maxValue *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.UpdateRiskParameter(&_Perpetual.TransactOpts, key, newValue, minValue, maxValue)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7da7d3f1.
//
// Solidity: function withdraw(address trader, int256 amount) returns()
func (_Perpetual *PerpetualTransactor) Withdraw(opts *bind.TransactOpts, trader common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "withdraw", trader, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7da7d3f1.
//
// Solidity: function withdraw(address trader, int256 amount) returns()
func (_Perpetual *PerpetualSession) Withdraw(trader common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.Withdraw(&_Perpetual.TransactOpts, trader, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7da7d3f1.
//
// Solidity: function withdraw(address trader, int256 amount) returns()
func (_Perpetual *PerpetualTransactorSession) Withdraw(trader common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Perpetual.Contract.Withdraw(&_Perpetual.TransactOpts, trader, amount)
}

// WithdrawableMargin is a paid mutator transaction binding the contract method 0x701d8b24.
//
// Solidity: function withdrawableMargin(address trader) returns(int256 withdrawable)
func (_Perpetual *PerpetualTransactor) WithdrawableMargin(opts *bind.TransactOpts, trader common.Address) (*types.Transaction, error) {
	return _Perpetual.contract.Transact(opts, "withdrawableMargin", trader)
}

// WithdrawableMargin is a paid mutator transaction binding the contract method 0x701d8b24.
//
// Solidity: function withdrawableMargin(address trader) returns(int256 withdrawable)
func (_Perpetual *PerpetualSession) WithdrawableMargin(trader common.Address) (*types.Transaction, error) {
	return _Perpetual.Contract.WithdrawableMargin(&_Perpetual.TransactOpts, trader)
}

// WithdrawableMargin is a paid mutator transaction binding the contract method 0x701d8b24.
//
// Solidity: function withdrawableMargin(address trader) returns(int256 withdrawable)
func (_Perpetual *PerpetualTransactorSession) WithdrawableMargin(trader common.Address) (*types.Transaction, error) {
	return _Perpetual.Contract.WithdrawableMargin(&_Perpetual.TransactOpts, trader)
}

// PerpetualAddLiquidatityIterator is returned from FilterAddLiquidatity and is used to iterate over the raw logs and unpacked data for AddLiquidatity events raised by the Perpetual contract.
type PerpetualAddLiquidatityIterator struct {
	Event *PerpetualAddLiquidatity // Event containing the contract specifics and raw log

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
func (it *PerpetualAddLiquidatityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpetualAddLiquidatity)
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
		it.Event = new(PerpetualAddLiquidatity)
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
func (it *PerpetualAddLiquidatityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpetualAddLiquidatityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpetualAddLiquidatity represents a AddLiquidatity event raised by the Perpetual contract.
type PerpetualAddLiquidatity struct {
	Trader      common.Address
	AddedCash   *big.Int
	MintedShare *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidatity is a free log retrieval operation binding the contract event 0xe148f39044da744c939a5c69461bb421044c9416b56f214ef09e7e5127629eb1.
//
// Solidity: event AddLiquidatity(address trader, int256 addedCash, int256 mintedShare)
func (_Perpetual *PerpetualFilterer) FilterAddLiquidatity(opts *bind.FilterOpts) (*PerpetualAddLiquidatityIterator, error) {

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "AddLiquidatity")
	if err != nil {
		return nil, err
	}
	return &PerpetualAddLiquidatityIterator{contract: _Perpetual.contract, event: "AddLiquidatity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidatity is a free log subscription operation binding the contract event 0xe148f39044da744c939a5c69461bb421044c9416b56f214ef09e7e5127629eb1.
//
// Solidity: event AddLiquidatity(address trader, int256 addedCash, int256 mintedShare)
func (_Perpetual *PerpetualFilterer) WatchAddLiquidatity(opts *bind.WatchOpts, sink chan<- *PerpetualAddLiquidatity) (event.Subscription, error) {

	logs, sub, err := _Perpetual.contract.WatchLogs(opts, "AddLiquidatity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpetualAddLiquidatity)
				if err := _Perpetual.contract.UnpackLog(event, "AddLiquidatity", log); err != nil {
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

// ParseAddLiquidatity is a log parse operation binding the contract event 0xe148f39044da744c939a5c69461bb421044c9416b56f214ef09e7e5127629eb1.
//
// Solidity: event AddLiquidatity(address trader, int256 addedCash, int256 mintedShare)
func (_Perpetual *PerpetualFilterer) ParseAddLiquidatity(log types.Log) (*PerpetualAddLiquidatity, error) {
	event := new(PerpetualAddLiquidatity)
	if err := _Perpetual.contract.UnpackLog(event, "AddLiquidatity", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PerpetualAdjustRiskSettingIterator is returned from FilterAdjustRiskSetting and is used to iterate over the raw logs and unpacked data for AdjustRiskSetting events raised by the Perpetual contract.
type PerpetualAdjustRiskSettingIterator struct {
	Event *PerpetualAdjustRiskSetting // Event containing the contract specifics and raw log

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
func (it *PerpetualAdjustRiskSettingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpetualAdjustRiskSetting)
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
		it.Event = new(PerpetualAdjustRiskSetting)
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
func (it *PerpetualAdjustRiskSettingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpetualAdjustRiskSettingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpetualAdjustRiskSetting represents a AdjustRiskSetting event raised by the Perpetual contract.
type PerpetualAdjustRiskSetting struct {
	Key   [32]byte
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdjustRiskSetting is a free log retrieval operation binding the contract event 0xa5de92b475d344c8314f3b21ad472ce28639ad724b5fc2a0391deb8588f159ea.
//
// Solidity: event AdjustRiskSetting(bytes32 key, int256 value)
func (_Perpetual *PerpetualFilterer) FilterAdjustRiskSetting(opts *bind.FilterOpts) (*PerpetualAdjustRiskSettingIterator, error) {

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "AdjustRiskSetting")
	if err != nil {
		return nil, err
	}
	return &PerpetualAdjustRiskSettingIterator{contract: _Perpetual.contract, event: "AdjustRiskSetting", logs: logs, sub: sub}, nil
}

// WatchAdjustRiskSetting is a free log subscription operation binding the contract event 0xa5de92b475d344c8314f3b21ad472ce28639ad724b5fc2a0391deb8588f159ea.
//
// Solidity: event AdjustRiskSetting(bytes32 key, int256 value)
func (_Perpetual *PerpetualFilterer) WatchAdjustRiskSetting(opts *bind.WatchOpts, sink chan<- *PerpetualAdjustRiskSetting) (event.Subscription, error) {

	logs, sub, err := _Perpetual.contract.WatchLogs(opts, "AdjustRiskSetting")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpetualAdjustRiskSetting)
				if err := _Perpetual.contract.UnpackLog(event, "AdjustRiskSetting", log); err != nil {
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

// ParseAdjustRiskSetting is a log parse operation binding the contract event 0xa5de92b475d344c8314f3b21ad472ce28639ad724b5fc2a0391deb8588f159ea.
//
// Solidity: event AdjustRiskSetting(bytes32 key, int256 value)
func (_Perpetual *PerpetualFilterer) ParseAdjustRiskSetting(log types.Log) (*PerpetualAdjustRiskSetting, error) {
	event := new(PerpetualAdjustRiskSetting)
	if err := _Perpetual.contract.UnpackLog(event, "AdjustRiskSetting", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PerpetualClaimFeeIterator is returned from FilterClaimFee and is used to iterate over the raw logs and unpacked data for ClaimFee events raised by the Perpetual contract.
type PerpetualClaimFeeIterator struct {
	Event *PerpetualClaimFee // Event containing the contract specifics and raw log

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
func (it *PerpetualClaimFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpetualClaimFee)
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
		it.Event = new(PerpetualClaimFee)
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
func (it *PerpetualClaimFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpetualClaimFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpetualClaimFee represents a ClaimFee event raised by the Perpetual contract.
type PerpetualClaimFee struct {
	Claimer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimFee is a free log retrieval operation binding the contract event 0x10688f72abe281c3b2c60342fd825674ce2e3d773fe911bc9d96efa899d0109d.
//
// Solidity: event ClaimFee(address claimer, int256 amount)
func (_Perpetual *PerpetualFilterer) FilterClaimFee(opts *bind.FilterOpts) (*PerpetualClaimFeeIterator, error) {

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "ClaimFee")
	if err != nil {
		return nil, err
	}
	return &PerpetualClaimFeeIterator{contract: _Perpetual.contract, event: "ClaimFee", logs: logs, sub: sub}, nil
}

// WatchClaimFee is a free log subscription operation binding the contract event 0x10688f72abe281c3b2c60342fd825674ce2e3d773fe911bc9d96efa899d0109d.
//
// Solidity: event ClaimFee(address claimer, int256 amount)
func (_Perpetual *PerpetualFilterer) WatchClaimFee(opts *bind.WatchOpts, sink chan<- *PerpetualClaimFee) (event.Subscription, error) {

	logs, sub, err := _Perpetual.contract.WatchLogs(opts, "ClaimFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpetualClaimFee)
				if err := _Perpetual.contract.UnpackLog(event, "ClaimFee", log); err != nil {
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

// ParseClaimFee is a log parse operation binding the contract event 0x10688f72abe281c3b2c60342fd825674ce2e3d773fe911bc9d96efa899d0109d.
//
// Solidity: event ClaimFee(address claimer, int256 amount)
func (_Perpetual *PerpetualFilterer) ParseClaimFee(log types.Log) (*PerpetualClaimFee, error) {
	event := new(PerpetualClaimFee)
	if err := _Perpetual.contract.UnpackLog(event, "ClaimFee", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PerpetualClearIterator is returned from FilterClear and is used to iterate over the raw logs and unpacked data for Clear events raised by the Perpetual contract.
type PerpetualClearIterator struct {
	Event *PerpetualClear // Event containing the contract specifics and raw log

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
func (it *PerpetualClearIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpetualClear)
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
		it.Event = new(PerpetualClear)
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
func (it *PerpetualClearIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpetualClearIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpetualClear represents a Clear event raised by the Perpetual contract.
type PerpetualClear struct {
	Trader common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClear is a free log retrieval operation binding the contract event 0x832e6903187c7c13e6dd325d0792d02ed4fb30e708d69386ece3ddc029fd9b4f.
//
// Solidity: event Clear(address trader)
func (_Perpetual *PerpetualFilterer) FilterClear(opts *bind.FilterOpts) (*PerpetualClearIterator, error) {

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "Clear")
	if err != nil {
		return nil, err
	}
	return &PerpetualClearIterator{contract: _Perpetual.contract, event: "Clear", logs: logs, sub: sub}, nil
}

// WatchClear is a free log subscription operation binding the contract event 0x832e6903187c7c13e6dd325d0792d02ed4fb30e708d69386ece3ddc029fd9b4f.
//
// Solidity: event Clear(address trader)
func (_Perpetual *PerpetualFilterer) WatchClear(opts *bind.WatchOpts, sink chan<- *PerpetualClear) (event.Subscription, error) {

	logs, sub, err := _Perpetual.contract.WatchLogs(opts, "Clear")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpetualClear)
				if err := _Perpetual.contract.UnpackLog(event, "Clear", log); err != nil {
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

// ParseClear is a log parse operation binding the contract event 0x832e6903187c7c13e6dd325d0792d02ed4fb30e708d69386ece3ddc029fd9b4f.
//
// Solidity: event Clear(address trader)
func (_Perpetual *PerpetualFilterer) ParseClear(log types.Log) (*PerpetualClear, error) {
	event := new(PerpetualClear)
	if err := _Perpetual.contract.UnpackLog(event, "Clear", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PerpetualClosePositionByLiquidationIterator is returned from FilterClosePositionByLiquidation and is used to iterate over the raw logs and unpacked data for ClosePositionByLiquidation events raised by the Perpetual contract.
type PerpetualClosePositionByLiquidationIterator struct {
	Event *PerpetualClosePositionByLiquidation // Event containing the contract specifics and raw log

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
func (it *PerpetualClosePositionByLiquidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpetualClosePositionByLiquidation)
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
		it.Event = new(PerpetualClosePositionByLiquidation)
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
func (it *PerpetualClosePositionByLiquidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpetualClosePositionByLiquidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpetualClosePositionByLiquidation represents a ClosePositionByLiquidation event raised by the Perpetual contract.
type PerpetualClosePositionByLiquidation struct {
	Trader      common.Address
	Amount      *big.Int
	Price       *big.Int
	FundingLoss *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClosePositionByLiquidation is a free log retrieval operation binding the contract event 0xabd2acbaab8c797dbdf9c9bd35a1513b44a62873014859fd3c3082dbc9b333ff.
//
// Solidity: event ClosePositionByLiquidation(address trader, int256 amount, int256 price, int256 fundingLoss)
func (_Perpetual *PerpetualFilterer) FilterClosePositionByLiquidation(opts *bind.FilterOpts) (*PerpetualClosePositionByLiquidationIterator, error) {

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "ClosePositionByLiquidation")
	if err != nil {
		return nil, err
	}
	return &PerpetualClosePositionByLiquidationIterator{contract: _Perpetual.contract, event: "ClosePositionByLiquidation", logs: logs, sub: sub}, nil
}

// WatchClosePositionByLiquidation is a free log subscription operation binding the contract event 0xabd2acbaab8c797dbdf9c9bd35a1513b44a62873014859fd3c3082dbc9b333ff.
//
// Solidity: event ClosePositionByLiquidation(address trader, int256 amount, int256 price, int256 fundingLoss)
func (_Perpetual *PerpetualFilterer) WatchClosePositionByLiquidation(opts *bind.WatchOpts, sink chan<- *PerpetualClosePositionByLiquidation) (event.Subscription, error) {

	logs, sub, err := _Perpetual.contract.WatchLogs(opts, "ClosePositionByLiquidation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpetualClosePositionByLiquidation)
				if err := _Perpetual.contract.UnpackLog(event, "ClosePositionByLiquidation", log); err != nil {
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

// ParseClosePositionByLiquidation is a log parse operation binding the contract event 0xabd2acbaab8c797dbdf9c9bd35a1513b44a62873014859fd3c3082dbc9b333ff.
//
// Solidity: event ClosePositionByLiquidation(address trader, int256 amount, int256 price, int256 fundingLoss)
func (_Perpetual *PerpetualFilterer) ParseClosePositionByLiquidation(log types.Log) (*PerpetualClosePositionByLiquidation, error) {
	event := new(PerpetualClosePositionByLiquidation)
	if err := _Perpetual.contract.UnpackLog(event, "ClosePositionByLiquidation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PerpetualClosePositionByTradeIterator is returned from FilterClosePositionByTrade and is used to iterate over the raw logs and unpacked data for ClosePositionByTrade events raised by the Perpetual contract.
type PerpetualClosePositionByTradeIterator struct {
	Event *PerpetualClosePositionByTrade // Event containing the contract specifics and raw log

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
func (it *PerpetualClosePositionByTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpetualClosePositionByTrade)
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
		it.Event = new(PerpetualClosePositionByTrade)
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
func (it *PerpetualClosePositionByTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpetualClosePositionByTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpetualClosePositionByTrade represents a ClosePositionByTrade event raised by the Perpetual contract.
type PerpetualClosePositionByTrade struct {
	Trader      common.Address
	Amount      *big.Int
	Price       *big.Int
	FundingLoss *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClosePositionByTrade is a free log retrieval operation binding the contract event 0x149cf94ac5b3602e20e97fe82a1d7e30bdbf553f6a59b7d10dbf77947fd7461e.
//
// Solidity: event ClosePositionByTrade(address trader, int256 amount, int256 price, int256 fundingLoss)
func (_Perpetual *PerpetualFilterer) FilterClosePositionByTrade(opts *bind.FilterOpts) (*PerpetualClosePositionByTradeIterator, error) {

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "ClosePositionByTrade")
	if err != nil {
		return nil, err
	}
	return &PerpetualClosePositionByTradeIterator{contract: _Perpetual.contract, event: "ClosePositionByTrade", logs: logs, sub: sub}, nil
}

// WatchClosePositionByTrade is a free log subscription operation binding the contract event 0x149cf94ac5b3602e20e97fe82a1d7e30bdbf553f6a59b7d10dbf77947fd7461e.
//
// Solidity: event ClosePositionByTrade(address trader, int256 amount, int256 price, int256 fundingLoss)
func (_Perpetual *PerpetualFilterer) WatchClosePositionByTrade(opts *bind.WatchOpts, sink chan<- *PerpetualClosePositionByTrade) (event.Subscription, error) {

	logs, sub, err := _Perpetual.contract.WatchLogs(opts, "ClosePositionByTrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpetualClosePositionByTrade)
				if err := _Perpetual.contract.UnpackLog(event, "ClosePositionByTrade", log); err != nil {
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

// ParseClosePositionByTrade is a log parse operation binding the contract event 0x149cf94ac5b3602e20e97fe82a1d7e30bdbf553f6a59b7d10dbf77947fd7461e.
//
// Solidity: event ClosePositionByTrade(address trader, int256 amount, int256 price, int256 fundingLoss)
func (_Perpetual *PerpetualFilterer) ParseClosePositionByTrade(log types.Log) (*PerpetualClosePositionByTrade, error) {
	event := new(PerpetualClosePositionByTrade)
	if err := _Perpetual.contract.UnpackLog(event, "ClosePositionByTrade", log); err != nil {
		return nil, err
	}
	return event, nil
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
	Trader common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xd8a6d38df847dcba70dfdeb4948fb1457d61a81d132801f40dc9c00d52dfd478.
//
// Solidity: event Deposit(address trader, int256 amount)
func (_Perpetual *PerpetualFilterer) FilterDeposit(opts *bind.FilterOpts) (*PerpetualDepositIterator, error) {

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &PerpetualDepositIterator{contract: _Perpetual.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xd8a6d38df847dcba70dfdeb4948fb1457d61a81d132801f40dc9c00d52dfd478.
//
// Solidity: event Deposit(address trader, int256 amount)
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
// Solidity: event Deposit(address trader, int256 amount)
func (_Perpetual *PerpetualFilterer) ParseDeposit(log types.Log) (*PerpetualDeposit, error) {
	event := new(PerpetualDeposit)
	if err := _Perpetual.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PerpetualDonateInsuranceFundIterator is returned from FilterDonateInsuranceFund and is used to iterate over the raw logs and unpacked data for DonateInsuranceFund events raised by the Perpetual contract.
type PerpetualDonateInsuranceFundIterator struct {
	Event *PerpetualDonateInsuranceFund // Event containing the contract specifics and raw log

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
func (it *PerpetualDonateInsuranceFundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpetualDonateInsuranceFund)
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
		it.Event = new(PerpetualDonateInsuranceFund)
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
func (it *PerpetualDonateInsuranceFundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpetualDonateInsuranceFundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpetualDonateInsuranceFund represents a DonateInsuranceFund event raised by the Perpetual contract.
type PerpetualDonateInsuranceFund struct {
	Trader common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDonateInsuranceFund is a free log retrieval operation binding the contract event 0x076a19d3bb1fbcd6fdee9888f4e0ab00a11cb2c4400a1d545eaeaf6bb5c25bc6.
//
// Solidity: event DonateInsuranceFund(address trader, int256 amount)
func (_Perpetual *PerpetualFilterer) FilterDonateInsuranceFund(opts *bind.FilterOpts) (*PerpetualDonateInsuranceFundIterator, error) {

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "DonateInsuranceFund")
	if err != nil {
		return nil, err
	}
	return &PerpetualDonateInsuranceFundIterator{contract: _Perpetual.contract, event: "DonateInsuranceFund", logs: logs, sub: sub}, nil
}

// WatchDonateInsuranceFund is a free log subscription operation binding the contract event 0x076a19d3bb1fbcd6fdee9888f4e0ab00a11cb2c4400a1d545eaeaf6bb5c25bc6.
//
// Solidity: event DonateInsuranceFund(address trader, int256 amount)
func (_Perpetual *PerpetualFilterer) WatchDonateInsuranceFund(opts *bind.WatchOpts, sink chan<- *PerpetualDonateInsuranceFund) (event.Subscription, error) {

	logs, sub, err := _Perpetual.contract.WatchLogs(opts, "DonateInsuranceFund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpetualDonateInsuranceFund)
				if err := _Perpetual.contract.UnpackLog(event, "DonateInsuranceFund", log); err != nil {
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

// ParseDonateInsuranceFund is a log parse operation binding the contract event 0x076a19d3bb1fbcd6fdee9888f4e0ab00a11cb2c4400a1d545eaeaf6bb5c25bc6.
//
// Solidity: event DonateInsuranceFund(address trader, int256 amount)
func (_Perpetual *PerpetualFilterer) ParseDonateInsuranceFund(log types.Log) (*PerpetualDonateInsuranceFund, error) {
	event := new(PerpetualDonateInsuranceFund)
	if err := _Perpetual.contract.UnpackLog(event, "DonateInsuranceFund", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PerpetualGrantPrivilegeIterator is returned from FilterGrantPrivilege and is used to iterate over the raw logs and unpacked data for GrantPrivilege events raised by the Perpetual contract.
type PerpetualGrantPrivilegeIterator struct {
	Event *PerpetualGrantPrivilege // Event containing the contract specifics and raw log

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
func (it *PerpetualGrantPrivilegeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpetualGrantPrivilege)
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
		it.Event = new(PerpetualGrantPrivilege)
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
func (it *PerpetualGrantPrivilegeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpetualGrantPrivilegeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpetualGrantPrivilege represents a GrantPrivilege event raised by the Perpetual contract.
type PerpetualGrantPrivilege struct {
	Owner     common.Address
	Trader    common.Address
	Privilege *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGrantPrivilege is a free log retrieval operation binding the contract event 0xb542e9113bf4db53a191bfb69cafe0c1d0647ca432aa8a6d36b6f9b04ec5473b.
//
// Solidity: event GrantPrivilege(address indexed owner, address indexed trader, uint256 privilege)
func (_Perpetual *PerpetualFilterer) FilterGrantPrivilege(opts *bind.FilterOpts, owner []common.Address, trader []common.Address) (*PerpetualGrantPrivilegeIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "GrantPrivilege", ownerRule, traderRule)
	if err != nil {
		return nil, err
	}
	return &PerpetualGrantPrivilegeIterator{contract: _Perpetual.contract, event: "GrantPrivilege", logs: logs, sub: sub}, nil
}

// WatchGrantPrivilege is a free log subscription operation binding the contract event 0xb542e9113bf4db53a191bfb69cafe0c1d0647ca432aa8a6d36b6f9b04ec5473b.
//
// Solidity: event GrantPrivilege(address indexed owner, address indexed trader, uint256 privilege)
func (_Perpetual *PerpetualFilterer) WatchGrantPrivilege(opts *bind.WatchOpts, sink chan<- *PerpetualGrantPrivilege, owner []common.Address, trader []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Perpetual.contract.WatchLogs(opts, "GrantPrivilege", ownerRule, traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpetualGrantPrivilege)
				if err := _Perpetual.contract.UnpackLog(event, "GrantPrivilege", log); err != nil {
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

// ParseGrantPrivilege is a log parse operation binding the contract event 0xb542e9113bf4db53a191bfb69cafe0c1d0647ca432aa8a6d36b6f9b04ec5473b.
//
// Solidity: event GrantPrivilege(address indexed owner, address indexed trader, uint256 privilege)
func (_Perpetual *PerpetualFilterer) ParseGrantPrivilege(log types.Log) (*PerpetualGrantPrivilege, error) {
	event := new(PerpetualGrantPrivilege)
	if err := _Perpetual.contract.UnpackLog(event, "GrantPrivilege", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PerpetualLiquidateByAMMIterator is returned from FilterLiquidateByAMM and is used to iterate over the raw logs and unpacked data for LiquidateByAMM events raised by the Perpetual contract.
type PerpetualLiquidateByAMMIterator struct {
	Event *PerpetualLiquidateByAMM // Event containing the contract specifics and raw log

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
func (it *PerpetualLiquidateByAMMIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpetualLiquidateByAMM)
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
		it.Event = new(PerpetualLiquidateByAMM)
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
func (it *PerpetualLiquidateByAMMIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpetualLiquidateByAMMIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpetualLiquidateByAMM represents a LiquidateByAMM event raised by the Perpetual contract.
type PerpetualLiquidateByAMM struct {
	Trader   common.Address
	Amount   *big.Int
	Price    *big.Int
	Fee      *big.Int
	Deadline *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLiquidateByAMM is a free log retrieval operation binding the contract event 0x6f7f426dbfbc2e084e8a286aa6ff573f30befbc4427e41611048846548d1a5f1.
//
// Solidity: event LiquidateByAMM(address indexed trader, int256 amount, int256 price, int256 fee, uint256 deadline)
func (_Perpetual *PerpetualFilterer) FilterLiquidateByAMM(opts *bind.FilterOpts, trader []common.Address) (*PerpetualLiquidateByAMMIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "LiquidateByAMM", traderRule)
	if err != nil {
		return nil, err
	}
	return &PerpetualLiquidateByAMMIterator{contract: _Perpetual.contract, event: "LiquidateByAMM", logs: logs, sub: sub}, nil
}

// WatchLiquidateByAMM is a free log subscription operation binding the contract event 0x6f7f426dbfbc2e084e8a286aa6ff573f30befbc4427e41611048846548d1a5f1.
//
// Solidity: event LiquidateByAMM(address indexed trader, int256 amount, int256 price, int256 fee, uint256 deadline)
func (_Perpetual *PerpetualFilterer) WatchLiquidateByAMM(opts *bind.WatchOpts, sink chan<- *PerpetualLiquidateByAMM, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Perpetual.contract.WatchLogs(opts, "LiquidateByAMM", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpetualLiquidateByAMM)
				if err := _Perpetual.contract.UnpackLog(event, "LiquidateByAMM", log); err != nil {
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

// ParseLiquidateByAMM is a log parse operation binding the contract event 0x6f7f426dbfbc2e084e8a286aa6ff573f30befbc4427e41611048846548d1a5f1.
//
// Solidity: event LiquidateByAMM(address indexed trader, int256 amount, int256 price, int256 fee, uint256 deadline)
func (_Perpetual *PerpetualFilterer) ParseLiquidateByAMM(log types.Log) (*PerpetualLiquidateByAMM, error) {
	event := new(PerpetualLiquidateByAMM)
	if err := _Perpetual.contract.UnpackLog(event, "LiquidateByAMM", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PerpetualLiquidateByTraderIterator is returned from FilterLiquidateByTrader and is used to iterate over the raw logs and unpacked data for LiquidateByTrader events raised by the Perpetual contract.
type PerpetualLiquidateByTraderIterator struct {
	Event *PerpetualLiquidateByTrader // Event containing the contract specifics and raw log

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
func (it *PerpetualLiquidateByTraderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpetualLiquidateByTrader)
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
		it.Event = new(PerpetualLiquidateByTrader)
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
func (it *PerpetualLiquidateByTraderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpetualLiquidateByTraderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpetualLiquidateByTrader represents a LiquidateByTrader event raised by the Perpetual contract.
type PerpetualLiquidateByTrader struct {
	Liquidator common.Address
	Trader     common.Address
	Amount     *big.Int
	Price      *big.Int
	Deadline   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLiquidateByTrader is a free log retrieval operation binding the contract event 0xb8105a35aa87a04868d98078a6291e279f44edfd3a0832e3e4a748803b030d7a.
//
// Solidity: event LiquidateByTrader(address indexed liquidator, address indexed trader, int256 amount, int256 price, uint256 deadline)
func (_Perpetual *PerpetualFilterer) FilterLiquidateByTrader(opts *bind.FilterOpts, liquidator []common.Address, trader []common.Address) (*PerpetualLiquidateByTraderIterator, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "LiquidateByTrader", liquidatorRule, traderRule)
	if err != nil {
		return nil, err
	}
	return &PerpetualLiquidateByTraderIterator{contract: _Perpetual.contract, event: "LiquidateByTrader", logs: logs, sub: sub}, nil
}

// WatchLiquidateByTrader is a free log subscription operation binding the contract event 0xb8105a35aa87a04868d98078a6291e279f44edfd3a0832e3e4a748803b030d7a.
//
// Solidity: event LiquidateByTrader(address indexed liquidator, address indexed trader, int256 amount, int256 price, uint256 deadline)
func (_Perpetual *PerpetualFilterer) WatchLiquidateByTrader(opts *bind.WatchOpts, sink chan<- *PerpetualLiquidateByTrader, liquidator []common.Address, trader []common.Address) (event.Subscription, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Perpetual.contract.WatchLogs(opts, "LiquidateByTrader", liquidatorRule, traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpetualLiquidateByTrader)
				if err := _Perpetual.contract.UnpackLog(event, "LiquidateByTrader", log); err != nil {
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

// ParseLiquidateByTrader is a log parse operation binding the contract event 0xb8105a35aa87a04868d98078a6291e279f44edfd3a0832e3e4a748803b030d7a.
//
// Solidity: event LiquidateByTrader(address indexed liquidator, address indexed trader, int256 amount, int256 price, uint256 deadline)
func (_Perpetual *PerpetualFilterer) ParseLiquidateByTrader(log types.Log) (*PerpetualLiquidateByTrader, error) {
	event := new(PerpetualLiquidateByTrader)
	if err := _Perpetual.contract.UnpackLog(event, "LiquidateByTrader", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PerpetualOpenPositionByLiquidationIterator is returned from FilterOpenPositionByLiquidation and is used to iterate over the raw logs and unpacked data for OpenPositionByLiquidation events raised by the Perpetual contract.
type PerpetualOpenPositionByLiquidationIterator struct {
	Event *PerpetualOpenPositionByLiquidation // Event containing the contract specifics and raw log

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
func (it *PerpetualOpenPositionByLiquidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpetualOpenPositionByLiquidation)
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
		it.Event = new(PerpetualOpenPositionByLiquidation)
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
func (it *PerpetualOpenPositionByLiquidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpetualOpenPositionByLiquidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpetualOpenPositionByLiquidation represents a OpenPositionByLiquidation event raised by the Perpetual contract.
type PerpetualOpenPositionByLiquidation struct {
	Trader common.Address
	Amount *big.Int
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOpenPositionByLiquidation is a free log retrieval operation binding the contract event 0x94cd570619ee4a4d76572118793021fe075463a74a86134e67c857708c4b0e61.
//
// Solidity: event OpenPositionByLiquidation(address trader, int256 amount, int256 price)
func (_Perpetual *PerpetualFilterer) FilterOpenPositionByLiquidation(opts *bind.FilterOpts) (*PerpetualOpenPositionByLiquidationIterator, error) {

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "OpenPositionByLiquidation")
	if err != nil {
		return nil, err
	}
	return &PerpetualOpenPositionByLiquidationIterator{contract: _Perpetual.contract, event: "OpenPositionByLiquidation", logs: logs, sub: sub}, nil
}

// WatchOpenPositionByLiquidation is a free log subscription operation binding the contract event 0x94cd570619ee4a4d76572118793021fe075463a74a86134e67c857708c4b0e61.
//
// Solidity: event OpenPositionByLiquidation(address trader, int256 amount, int256 price)
func (_Perpetual *PerpetualFilterer) WatchOpenPositionByLiquidation(opts *bind.WatchOpts, sink chan<- *PerpetualOpenPositionByLiquidation) (event.Subscription, error) {

	logs, sub, err := _Perpetual.contract.WatchLogs(opts, "OpenPositionByLiquidation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpetualOpenPositionByLiquidation)
				if err := _Perpetual.contract.UnpackLog(event, "OpenPositionByLiquidation", log); err != nil {
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

// ParseOpenPositionByLiquidation is a log parse operation binding the contract event 0x94cd570619ee4a4d76572118793021fe075463a74a86134e67c857708c4b0e61.
//
// Solidity: event OpenPositionByLiquidation(address trader, int256 amount, int256 price)
func (_Perpetual *PerpetualFilterer) ParseOpenPositionByLiquidation(log types.Log) (*PerpetualOpenPositionByLiquidation, error) {
	event := new(PerpetualOpenPositionByLiquidation)
	if err := _Perpetual.contract.UnpackLog(event, "OpenPositionByLiquidation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PerpetualOpenPositionByTradeIterator is returned from FilterOpenPositionByTrade and is used to iterate over the raw logs and unpacked data for OpenPositionByTrade events raised by the Perpetual contract.
type PerpetualOpenPositionByTradeIterator struct {
	Event *PerpetualOpenPositionByTrade // Event containing the contract specifics and raw log

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
func (it *PerpetualOpenPositionByTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpetualOpenPositionByTrade)
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
		it.Event = new(PerpetualOpenPositionByTrade)
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
func (it *PerpetualOpenPositionByTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpetualOpenPositionByTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpetualOpenPositionByTrade represents a OpenPositionByTrade event raised by the Perpetual contract.
type PerpetualOpenPositionByTrade struct {
	Trader common.Address
	Amount *big.Int
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOpenPositionByTrade is a free log retrieval operation binding the contract event 0xb59120e766f78f28f19e3656d568ce0dd0b17ebdfdc4d97b7ad08a3419d1caf9.
//
// Solidity: event OpenPositionByTrade(address trader, int256 amount, int256 price)
func (_Perpetual *PerpetualFilterer) FilterOpenPositionByTrade(opts *bind.FilterOpts) (*PerpetualOpenPositionByTradeIterator, error) {

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "OpenPositionByTrade")
	if err != nil {
		return nil, err
	}
	return &PerpetualOpenPositionByTradeIterator{contract: _Perpetual.contract, event: "OpenPositionByTrade", logs: logs, sub: sub}, nil
}

// WatchOpenPositionByTrade is a free log subscription operation binding the contract event 0xb59120e766f78f28f19e3656d568ce0dd0b17ebdfdc4d97b7ad08a3419d1caf9.
//
// Solidity: event OpenPositionByTrade(address trader, int256 amount, int256 price)
func (_Perpetual *PerpetualFilterer) WatchOpenPositionByTrade(opts *bind.WatchOpts, sink chan<- *PerpetualOpenPositionByTrade) (event.Subscription, error) {

	logs, sub, err := _Perpetual.contract.WatchLogs(opts, "OpenPositionByTrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpetualOpenPositionByTrade)
				if err := _Perpetual.contract.UnpackLog(event, "OpenPositionByTrade", log); err != nil {
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

// ParseOpenPositionByTrade is a log parse operation binding the contract event 0xb59120e766f78f28f19e3656d568ce0dd0b17ebdfdc4d97b7ad08a3419d1caf9.
//
// Solidity: event OpenPositionByTrade(address trader, int256 amount, int256 price)
func (_Perpetual *PerpetualFilterer) ParseOpenPositionByTrade(log types.Log) (*PerpetualOpenPositionByTrade, error) {
	event := new(PerpetualOpenPositionByTrade)
	if err := _Perpetual.contract.UnpackLog(event, "OpenPositionByTrade", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PerpetualRemoveLiquidatityIterator is returned from FilterRemoveLiquidatity and is used to iterate over the raw logs and unpacked data for RemoveLiquidatity events raised by the Perpetual contract.
type PerpetualRemoveLiquidatityIterator struct {
	Event *PerpetualRemoveLiquidatity // Event containing the contract specifics and raw log

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
func (it *PerpetualRemoveLiquidatityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpetualRemoveLiquidatity)
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
		it.Event = new(PerpetualRemoveLiquidatity)
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
func (it *PerpetualRemoveLiquidatityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpetualRemoveLiquidatityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpetualRemoveLiquidatity represents a RemoveLiquidatity event raised by the Perpetual contract.
type PerpetualRemoveLiquidatity struct {
	Trader       common.Address
	ReturnedCash *big.Int
	BurnedShare  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidatity is a free log retrieval operation binding the contract event 0xcb7e0e53409e241a54d296bfa7a83a012cd7f951ed5f8d5ca03d57aa2f6a3bf9.
//
// Solidity: event RemoveLiquidatity(address trader, int256 returnedCash, int256 burnedShare)
func (_Perpetual *PerpetualFilterer) FilterRemoveLiquidatity(opts *bind.FilterOpts) (*PerpetualRemoveLiquidatityIterator, error) {

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "RemoveLiquidatity")
	if err != nil {
		return nil, err
	}
	return &PerpetualRemoveLiquidatityIterator{contract: _Perpetual.contract, event: "RemoveLiquidatity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidatity is a free log subscription operation binding the contract event 0xcb7e0e53409e241a54d296bfa7a83a012cd7f951ed5f8d5ca03d57aa2f6a3bf9.
//
// Solidity: event RemoveLiquidatity(address trader, int256 returnedCash, int256 burnedShare)
func (_Perpetual *PerpetualFilterer) WatchRemoveLiquidatity(opts *bind.WatchOpts, sink chan<- *PerpetualRemoveLiquidatity) (event.Subscription, error) {

	logs, sub, err := _Perpetual.contract.WatchLogs(opts, "RemoveLiquidatity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpetualRemoveLiquidatity)
				if err := _Perpetual.contract.UnpackLog(event, "RemoveLiquidatity", log); err != nil {
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

// ParseRemoveLiquidatity is a log parse operation binding the contract event 0xcb7e0e53409e241a54d296bfa7a83a012cd7f951ed5f8d5ca03d57aa2f6a3bf9.
//
// Solidity: event RemoveLiquidatity(address trader, int256 returnedCash, int256 burnedShare)
func (_Perpetual *PerpetualFilterer) ParseRemoveLiquidatity(log types.Log) (*PerpetualRemoveLiquidatity, error) {
	event := new(PerpetualRemoveLiquidatity)
	if err := _Perpetual.contract.UnpackLog(event, "RemoveLiquidatity", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PerpetualRevokePrivilegeIterator is returned from FilterRevokePrivilege and is used to iterate over the raw logs and unpacked data for RevokePrivilege events raised by the Perpetual contract.
type PerpetualRevokePrivilegeIterator struct {
	Event *PerpetualRevokePrivilege // Event containing the contract specifics and raw log

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
func (it *PerpetualRevokePrivilegeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpetualRevokePrivilege)
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
		it.Event = new(PerpetualRevokePrivilege)
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
func (it *PerpetualRevokePrivilegeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpetualRevokePrivilegeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpetualRevokePrivilege represents a RevokePrivilege event raised by the Perpetual contract.
type PerpetualRevokePrivilege struct {
	Owner     common.Address
	Trader    common.Address
	Privilege *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRevokePrivilege is a free log retrieval operation binding the contract event 0x2e910d80b3ddddaff95f2e39c7c226987cd86d56b81c7352b8393efe6f091470.
//
// Solidity: event RevokePrivilege(address indexed owner, address indexed trader, uint256 privilege)
func (_Perpetual *PerpetualFilterer) FilterRevokePrivilege(opts *bind.FilterOpts, owner []common.Address, trader []common.Address) (*PerpetualRevokePrivilegeIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "RevokePrivilege", ownerRule, traderRule)
	if err != nil {
		return nil, err
	}
	return &PerpetualRevokePrivilegeIterator{contract: _Perpetual.contract, event: "RevokePrivilege", logs: logs, sub: sub}, nil
}

// WatchRevokePrivilege is a free log subscription operation binding the contract event 0x2e910d80b3ddddaff95f2e39c7c226987cd86d56b81c7352b8393efe6f091470.
//
// Solidity: event RevokePrivilege(address indexed owner, address indexed trader, uint256 privilege)
func (_Perpetual *PerpetualFilterer) WatchRevokePrivilege(opts *bind.WatchOpts, sink chan<- *PerpetualRevokePrivilege, owner []common.Address, trader []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Perpetual.contract.WatchLogs(opts, "RevokePrivilege", ownerRule, traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpetualRevokePrivilege)
				if err := _Perpetual.contract.UnpackLog(event, "RevokePrivilege", log); err != nil {
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

// ParseRevokePrivilege is a log parse operation binding the contract event 0x2e910d80b3ddddaff95f2e39c7c226987cd86d56b81c7352b8393efe6f091470.
//
// Solidity: event RevokePrivilege(address indexed owner, address indexed trader, uint256 privilege)
func (_Perpetual *PerpetualFilterer) ParseRevokePrivilege(log types.Log) (*PerpetualRevokePrivilege, error) {
	event := new(PerpetualRevokePrivilege)
	if err := _Perpetual.contract.UnpackLog(event, "RevokePrivilege", log); err != nil {
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
	Trader         common.Address
	PositionAmount *big.Int
	PriceLimit     *big.Int
	Fee            *big.Int
	Deadline       *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0x42fb2e7f753dd5d3da9571845ec946cce6f1e8ad4d0b4c0d2adb56b3cc878d36.
//
// Solidity: event Trade(address indexed trader, int256 positionAmount, int256 priceLimit, int256 fee, uint256 deadline)
func (_Perpetual *PerpetualFilterer) FilterTrade(opts *bind.FilterOpts, trader []common.Address) (*PerpetualTradeIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "Trade", traderRule)
	if err != nil {
		return nil, err
	}
	return &PerpetualTradeIterator{contract: _Perpetual.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0x42fb2e7f753dd5d3da9571845ec946cce6f1e8ad4d0b4c0d2adb56b3cc878d36.
//
// Solidity: event Trade(address indexed trader, int256 positionAmount, int256 priceLimit, int256 fee, uint256 deadline)
func (_Perpetual *PerpetualFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *PerpetualTrade, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Perpetual.contract.WatchLogs(opts, "Trade", traderRule)
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

// ParseTrade is a log parse operation binding the contract event 0x42fb2e7f753dd5d3da9571845ec946cce6f1e8ad4d0b4c0d2adb56b3cc878d36.
//
// Solidity: event Trade(address indexed trader, int256 positionAmount, int256 priceLimit, int256 fee, uint256 deadline)
func (_Perpetual *PerpetualFilterer) ParseTrade(log types.Log) (*PerpetualTrade, error) {
	event := new(PerpetualTrade)
	if err := _Perpetual.contract.UnpackLog(event, "Trade", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PerpetualUpdateCoreSettingIterator is returned from FilterUpdateCoreSetting and is used to iterate over the raw logs and unpacked data for UpdateCoreSetting events raised by the Perpetual contract.
type PerpetualUpdateCoreSettingIterator struct {
	Event *PerpetualUpdateCoreSetting // Event containing the contract specifics and raw log

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
func (it *PerpetualUpdateCoreSettingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpetualUpdateCoreSetting)
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
		it.Event = new(PerpetualUpdateCoreSetting)
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
func (it *PerpetualUpdateCoreSettingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpetualUpdateCoreSettingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpetualUpdateCoreSetting represents a UpdateCoreSetting event raised by the Perpetual contract.
type PerpetualUpdateCoreSetting struct {
	Key   [32]byte
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdateCoreSetting is a free log retrieval operation binding the contract event 0x8f1d4b4cbbbc7eea0d52f98cdf4aa59cf90f71b4d4f75d98c995cce58fcac029.
//
// Solidity: event UpdateCoreSetting(bytes32 key, int256 value)
func (_Perpetual *PerpetualFilterer) FilterUpdateCoreSetting(opts *bind.FilterOpts) (*PerpetualUpdateCoreSettingIterator, error) {

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "UpdateCoreSetting")
	if err != nil {
		return nil, err
	}
	return &PerpetualUpdateCoreSettingIterator{contract: _Perpetual.contract, event: "UpdateCoreSetting", logs: logs, sub: sub}, nil
}

// WatchUpdateCoreSetting is a free log subscription operation binding the contract event 0x8f1d4b4cbbbc7eea0d52f98cdf4aa59cf90f71b4d4f75d98c995cce58fcac029.
//
// Solidity: event UpdateCoreSetting(bytes32 key, int256 value)
func (_Perpetual *PerpetualFilterer) WatchUpdateCoreSetting(opts *bind.WatchOpts, sink chan<- *PerpetualUpdateCoreSetting) (event.Subscription, error) {

	logs, sub, err := _Perpetual.contract.WatchLogs(opts, "UpdateCoreSetting")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpetualUpdateCoreSetting)
				if err := _Perpetual.contract.UnpackLog(event, "UpdateCoreSetting", log); err != nil {
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

// ParseUpdateCoreSetting is a log parse operation binding the contract event 0x8f1d4b4cbbbc7eea0d52f98cdf4aa59cf90f71b4d4f75d98c995cce58fcac029.
//
// Solidity: event UpdateCoreSetting(bytes32 key, int256 value)
func (_Perpetual *PerpetualFilterer) ParseUpdateCoreSetting(log types.Log) (*PerpetualUpdateCoreSetting, error) {
	event := new(PerpetualUpdateCoreSetting)
	if err := _Perpetual.contract.UnpackLog(event, "UpdateCoreSetting", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PerpetualUpdateRiskSettingIterator is returned from FilterUpdateRiskSetting and is used to iterate over the raw logs and unpacked data for UpdateRiskSetting events raised by the Perpetual contract.
type PerpetualUpdateRiskSettingIterator struct {
	Event *PerpetualUpdateRiskSetting // Event containing the contract specifics and raw log

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
func (it *PerpetualUpdateRiskSettingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpetualUpdateRiskSetting)
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
		it.Event = new(PerpetualUpdateRiskSetting)
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
func (it *PerpetualUpdateRiskSettingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpetualUpdateRiskSettingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpetualUpdateRiskSetting represents a UpdateRiskSetting event raised by the Perpetual contract.
type PerpetualUpdateRiskSetting struct {
	Key      [32]byte
	Value    *big.Int
	MinValue *big.Int
	MaxValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateRiskSetting is a free log retrieval operation binding the contract event 0x1c9b4ff9a80e6d7986441e116860b3c64867108497a3acf1e71b22ada39a8e6d.
//
// Solidity: event UpdateRiskSetting(bytes32 key, int256 value, int256 minValue, int256 maxValue)
func (_Perpetual *PerpetualFilterer) FilterUpdateRiskSetting(opts *bind.FilterOpts) (*PerpetualUpdateRiskSettingIterator, error) {

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "UpdateRiskSetting")
	if err != nil {
		return nil, err
	}
	return &PerpetualUpdateRiskSettingIterator{contract: _Perpetual.contract, event: "UpdateRiskSetting", logs: logs, sub: sub}, nil
}

// WatchUpdateRiskSetting is a free log subscription operation binding the contract event 0x1c9b4ff9a80e6d7986441e116860b3c64867108497a3acf1e71b22ada39a8e6d.
//
// Solidity: event UpdateRiskSetting(bytes32 key, int256 value, int256 minValue, int256 maxValue)
func (_Perpetual *PerpetualFilterer) WatchUpdateRiskSetting(opts *bind.WatchOpts, sink chan<- *PerpetualUpdateRiskSetting) (event.Subscription, error) {

	logs, sub, err := _Perpetual.contract.WatchLogs(opts, "UpdateRiskSetting")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpetualUpdateRiskSetting)
				if err := _Perpetual.contract.UnpackLog(event, "UpdateRiskSetting", log); err != nil {
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

// ParseUpdateRiskSetting is a log parse operation binding the contract event 0x1c9b4ff9a80e6d7986441e116860b3c64867108497a3acf1e71b22ada39a8e6d.
//
// Solidity: event UpdateRiskSetting(bytes32 key, int256 value, int256 minValue, int256 maxValue)
func (_Perpetual *PerpetualFilterer) ParseUpdateRiskSetting(log types.Log) (*PerpetualUpdateRiskSetting, error) {
	event := new(PerpetualUpdateRiskSetting)
	if err := _Perpetual.contract.UnpackLog(event, "UpdateRiskSetting", log); err != nil {
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
	Trader common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x2d3cfd22ce461d7eafde7bb13c6af0c0d5ed08406a59166e093b4354cfd94ae2.
//
// Solidity: event Withdraw(address trader, int256 amount)
func (_Perpetual *PerpetualFilterer) FilterWithdraw(opts *bind.FilterOpts) (*PerpetualWithdrawIterator, error) {

	logs, sub, err := _Perpetual.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &PerpetualWithdrawIterator{contract: _Perpetual.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x2d3cfd22ce461d7eafde7bb13c6af0c0d5ed08406a59166e093b4354cfd94ae2.
//
// Solidity: event Withdraw(address trader, int256 amount)
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
// Solidity: event Withdraw(address trader, int256 amount)
func (_Perpetual *PerpetualFilterer) ParseWithdraw(log types.Log) (*PerpetualWithdraw, error) {
	event := new(PerpetualWithdraw)
	if err := _Perpetual.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
