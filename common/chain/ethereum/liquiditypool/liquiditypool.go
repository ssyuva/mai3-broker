// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package liquiditypool

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
	Trader         common.Address
	Broker         common.Address
	Relayer        common.Address
	Referrer       common.Address
	LiquidityPool  common.Address
	PerpetualIndex *big.Int
	Amount         *big.Int
	PriceLimit     *big.Int
	MinTradeAmount *big.Int
	TradeGasLimit  *big.Int
	ChainID        *big.Int
	Data           [32]byte
}

// LiquidityPoolABI is the input ABI used to generate the binding from.
const LiquidityPoolABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"addedCash\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"mintedShare\",\"type\":\"int256\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"AdjustPerpetualRiskSetting\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"ClaimFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"Clear\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"shareToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256[8]\",\"name\":\"coreParams\",\"type\":\"int256[8]\"},{\"indexed\":false,\"internalType\":\"int256[5]\",\"name\":\"riskParams\",\"type\":\"int256[5]\"}],\"name\":\"CreatePerpetual\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"DonateInsuranceFund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Finalize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"}],\"name\":\"Liquidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"returnedCash\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"burnedShare\",\"type\":\"int256\"}],\"name\":\"RemoveLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"Settle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"positionAmount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fee\",\"type\":\"int256\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"UpdateLiquidityPoolParameter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"UpdatePerpetualParameter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"minValue\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"maxValue\",\"type\":\"int256\"}],\"name\":\"UpdatePerpetualRiskParameter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"}],\"name\":\"activeAccountCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"cashToAdd\",\"type\":\"int256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"newValue\",\"type\":\"int256\"}],\"name\":\"adjustPerpetualRiskParameter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"broker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"priceLimit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"minTradeAmount\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"tradeGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"brokerTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"claimFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"claimableFee\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"clear\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"int256[8]\",\"name\":\"coreParams\",\"type\":\"int256[8]\"},{\"internalType\":\"int256[5]\",\"name\":\"riskParams\",\"type\":\"int256[5]\"},{\"internalType\":\"int256[5]\",\"name\":\"minRiskParamValues\",\"type\":\"int256[5]\"},{\"internalType\":\"int256[5]\",\"name\":\"maxRiskParamValues\",\"type\":\"int256[5]\"}],\"name\":\"createPerpetual\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"donateInsuranceFund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"shareToken\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"liquidateByAMM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"priceLimit\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"liquidateByTrader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityPoolInfo\",\"outputs\":[{\"internalType\":\"address[6]\",\"name\":\"addresses\",\"type\":\"address[6]\"},{\"internalType\":\"int256[7]\",\"name\":\"nums\",\"type\":\"int256[7]\"},{\"internalType\":\"uint256\",\"name\":\"perpetualCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fundingTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"listActiveAccounts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"result\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"marginAccount\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"cashBalance\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"positionAmount\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"}],\"name\":\"perpetualInfo\",\"outputs\":[{\"internalType\":\"enumPerpetualState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"int256[17]\",\"name\":\"nums\",\"type\":\"int256[17]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"shareToRemove\",\"type\":\"int256\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"settle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"priceLimit\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isCloseOnly\",\"type\":\"bool\"}],\"name\":\"trade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"newValue\",\"type\":\"int256\"}],\"name\":\"updateLiquidityPoolParameter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"newValue\",\"type\":\"int256\"}],\"name\":\"updatePerpetualParameter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"newValue\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"minValue\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"maxValue\",\"type\":\"int256\"}],\"name\":\"updatePerpetualRiskParameter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// LiquidityPool is an auto generated Go binding around an Ethereum contract.
type LiquidityPool struct {
	LiquidityPoolCaller     // Read-only binding to the contract
	LiquidityPoolTransactor // Write-only binding to the contract
	LiquidityPoolFilterer   // Log filterer for contract events
}

// LiquidityPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidityPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidityPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidityPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidityPoolSession struct {
	Contract     *LiquidityPool    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LiquidityPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidityPoolCallerSession struct {
	Contract *LiquidityPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// LiquidityPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidityPoolTransactorSession struct {
	Contract     *LiquidityPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// LiquidityPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidityPoolRaw struct {
	Contract *LiquidityPool // Generic contract binding to access the raw methods on
}

// LiquidityPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidityPoolCallerRaw struct {
	Contract *LiquidityPoolCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidityPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidityPoolTransactorRaw struct {
	Contract *LiquidityPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquidityPool creates a new instance of LiquidityPool, bound to a specific deployed contract.
func NewLiquidityPool(address common.Address, backend bind.ContractBackend) (*LiquidityPool, error) {
	contract, err := bindLiquidityPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LiquidityPool{LiquidityPoolCaller: LiquidityPoolCaller{contract: contract}, LiquidityPoolTransactor: LiquidityPoolTransactor{contract: contract}, LiquidityPoolFilterer: LiquidityPoolFilterer{contract: contract}}, nil
}

// NewLiquidityPoolCaller creates a new read-only instance of LiquidityPool, bound to a specific deployed contract.
func NewLiquidityPoolCaller(address common.Address, caller bind.ContractCaller) (*LiquidityPoolCaller, error) {
	contract, err := bindLiquidityPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolCaller{contract: contract}, nil
}

// NewLiquidityPoolTransactor creates a new write-only instance of LiquidityPool, bound to a specific deployed contract.
func NewLiquidityPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidityPoolTransactor, error) {
	contract, err := bindLiquidityPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolTransactor{contract: contract}, nil
}

// NewLiquidityPoolFilterer creates a new log filterer instance of LiquidityPool, bound to a specific deployed contract.
func NewLiquidityPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidityPoolFilterer, error) {
	contract, err := bindLiquidityPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolFilterer{contract: contract}, nil
}

// bindLiquidityPool binds a generic wrapper to an already deployed contract.
func bindLiquidityPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LiquidityPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidityPool *LiquidityPoolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LiquidityPool.Contract.LiquidityPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidityPool *LiquidityPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPool.Contract.LiquidityPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidityPool *LiquidityPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidityPool.Contract.LiquidityPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidityPool *LiquidityPoolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LiquidityPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidityPool *LiquidityPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidityPool *LiquidityPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidityPool.Contract.contract.Transact(opts, method, params...)
}

// ActiveAccountCount is a free data retrieval call binding the contract method 0x6fbef638.
//
// Solidity: function activeAccountCount(uint256 perpetualIndex) constant returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) ActiveAccountCount(opts *bind.CallOpts, perpetualIndex *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityPool.contract.Call(opts, out, "activeAccountCount", perpetualIndex)
	return *ret0, err
}

// ActiveAccountCount is a free data retrieval call binding the contract method 0x6fbef638.
//
// Solidity: function activeAccountCount(uint256 perpetualIndex) constant returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) ActiveAccountCount(perpetualIndex *big.Int) (*big.Int, error) {
	return _LiquidityPool.Contract.ActiveAccountCount(&_LiquidityPool.CallOpts, perpetualIndex)
}

// ActiveAccountCount is a free data retrieval call binding the contract method 0x6fbef638.
//
// Solidity: function activeAccountCount(uint256 perpetualIndex) constant returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) ActiveAccountCount(perpetualIndex *big.Int) (*big.Int, error) {
	return _LiquidityPool.Contract.ActiveAccountCount(&_LiquidityPool.CallOpts, perpetualIndex)
}

// ClaimableFee is a free data retrieval call binding the contract method 0xa1c59afc.
//
// Solidity: function claimableFee(address claimer) constant returns(int256)
func (_LiquidityPool *LiquidityPoolCaller) ClaimableFee(opts *bind.CallOpts, claimer common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LiquidityPool.contract.Call(opts, out, "claimableFee", claimer)
	return *ret0, err
}

// ClaimableFee is a free data retrieval call binding the contract method 0xa1c59afc.
//
// Solidity: function claimableFee(address claimer) constant returns(int256)
func (_LiquidityPool *LiquidityPoolSession) ClaimableFee(claimer common.Address) (*big.Int, error) {
	return _LiquidityPool.Contract.ClaimableFee(&_LiquidityPool.CallOpts, claimer)
}

// ClaimableFee is a free data retrieval call binding the contract method 0xa1c59afc.
//
// Solidity: function claimableFee(address claimer) constant returns(int256)
func (_LiquidityPool *LiquidityPoolCallerSession) ClaimableFee(claimer common.Address) (*big.Int, error) {
	return _LiquidityPool.Contract.ClaimableFee(&_LiquidityPool.CallOpts, claimer)
}

// LiquidityPoolInfo is a free data retrieval call binding the contract method 0xc83937d1.
//
// Solidity: function liquidityPoolInfo() constant returns(address[6] addresses, int256[7] nums, uint256 perpetualCount, uint256 fundingTime)
func (_LiquidityPool *LiquidityPoolCaller) LiquidityPoolInfo(opts *bind.CallOpts) (struct {
	Addresses      [6]common.Address
	Nums           [7]*big.Int
	PerpetualCount *big.Int
	FundingTime    *big.Int
}, error) {
	ret := new(struct {
		Addresses      [6]common.Address
		Nums           [7]*big.Int
		PerpetualCount *big.Int
		FundingTime    *big.Int
	})
	out := ret
	err := _LiquidityPool.contract.Call(opts, out, "liquidityPoolInfo")
	return *ret, err
}

// LiquidityPoolInfo is a free data retrieval call binding the contract method 0xc83937d1.
//
// Solidity: function liquidityPoolInfo() constant returns(address[6] addresses, int256[7] nums, uint256 perpetualCount, uint256 fundingTime)
func (_LiquidityPool *LiquidityPoolSession) LiquidityPoolInfo() (struct {
	Addresses      [6]common.Address
	Nums           [7]*big.Int
	PerpetualCount *big.Int
	FundingTime    *big.Int
}, error) {
	return _LiquidityPool.Contract.LiquidityPoolInfo(&_LiquidityPool.CallOpts)
}

// LiquidityPoolInfo is a free data retrieval call binding the contract method 0xc83937d1.
//
// Solidity: function liquidityPoolInfo() constant returns(address[6] addresses, int256[7] nums, uint256 perpetualCount, uint256 fundingTime)
func (_LiquidityPool *LiquidityPoolCallerSession) LiquidityPoolInfo() (struct {
	Addresses      [6]common.Address
	Nums           [7]*big.Int
	PerpetualCount *big.Int
	FundingTime    *big.Int
}, error) {
	return _LiquidityPool.Contract.LiquidityPoolInfo(&_LiquidityPool.CallOpts)
}

// ListActiveAccounts is a free data retrieval call binding the contract method 0x13f07f45.
//
// Solidity: function listActiveAccounts(uint256 perpetualIndex, uint256 start, uint256 end) constant returns(address[] result)
func (_LiquidityPool *LiquidityPoolCaller) ListActiveAccounts(opts *bind.CallOpts, perpetualIndex *big.Int, start *big.Int, end *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _LiquidityPool.contract.Call(opts, out, "listActiveAccounts", perpetualIndex, start, end)
	return *ret0, err
}

// ListActiveAccounts is a free data retrieval call binding the contract method 0x13f07f45.
//
// Solidity: function listActiveAccounts(uint256 perpetualIndex, uint256 start, uint256 end) constant returns(address[] result)
func (_LiquidityPool *LiquidityPoolSession) ListActiveAccounts(perpetualIndex *big.Int, start *big.Int, end *big.Int) ([]common.Address, error) {
	return _LiquidityPool.Contract.ListActiveAccounts(&_LiquidityPool.CallOpts, perpetualIndex, start, end)
}

// ListActiveAccounts is a free data retrieval call binding the contract method 0x13f07f45.
//
// Solidity: function listActiveAccounts(uint256 perpetualIndex, uint256 start, uint256 end) constant returns(address[] result)
func (_LiquidityPool *LiquidityPoolCallerSession) ListActiveAccounts(perpetualIndex *big.Int, start *big.Int, end *big.Int) ([]common.Address, error) {
	return _LiquidityPool.Contract.ListActiveAccounts(&_LiquidityPool.CallOpts, perpetualIndex, start, end)
}

// MarginAccount is a free data retrieval call binding the contract method 0xc18f0d54.
//
// Solidity: function marginAccount(uint256 perpetualIndex, address trader) constant returns(int256 cashBalance, int256 positionAmount)
func (_LiquidityPool *LiquidityPoolCaller) MarginAccount(opts *bind.CallOpts, perpetualIndex *big.Int, trader common.Address) (struct {
	CashBalance    *big.Int
	PositionAmount *big.Int
}, error) {
	ret := new(struct {
		CashBalance    *big.Int
		PositionAmount *big.Int
	})
	out := ret
	err := _LiquidityPool.contract.Call(opts, out, "marginAccount", perpetualIndex, trader)
	return *ret, err
}

// MarginAccount is a free data retrieval call binding the contract method 0xc18f0d54.
//
// Solidity: function marginAccount(uint256 perpetualIndex, address trader) constant returns(int256 cashBalance, int256 positionAmount)
func (_LiquidityPool *LiquidityPoolSession) MarginAccount(perpetualIndex *big.Int, trader common.Address) (struct {
	CashBalance    *big.Int
	PositionAmount *big.Int
}, error) {
	return _LiquidityPool.Contract.MarginAccount(&_LiquidityPool.CallOpts, perpetualIndex, trader)
}

// MarginAccount is a free data retrieval call binding the contract method 0xc18f0d54.
//
// Solidity: function marginAccount(uint256 perpetualIndex, address trader) constant returns(int256 cashBalance, int256 positionAmount)
func (_LiquidityPool *LiquidityPoolCallerSession) MarginAccount(perpetualIndex *big.Int, trader common.Address) (struct {
	CashBalance    *big.Int
	PositionAmount *big.Int
}, error) {
	return _LiquidityPool.Contract.MarginAccount(&_LiquidityPool.CallOpts, perpetualIndex, trader)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x5cfb7f8f.
//
// Solidity: function addLiquidity(int256 cashToAdd) returns()
func (_LiquidityPool *LiquidityPoolTransactor) AddLiquidity(opts *bind.TransactOpts, cashToAdd *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "addLiquidity", cashToAdd)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x5cfb7f8f.
//
// Solidity: function addLiquidity(int256 cashToAdd) returns()
func (_LiquidityPool *LiquidityPoolSession) AddLiquidity(cashToAdd *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.AddLiquidity(&_LiquidityPool.TransactOpts, cashToAdd)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x5cfb7f8f.
//
// Solidity: function addLiquidity(int256 cashToAdd) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) AddLiquidity(cashToAdd *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.AddLiquidity(&_LiquidityPool.TransactOpts, cashToAdd)
}

// AdjustPerpetualRiskParameter is a paid mutator transaction binding the contract method 0x882eaa77.
//
// Solidity: function adjustPerpetualRiskParameter(uint256 perpetualIndex, bytes32 key, int256 newValue) returns()
func (_LiquidityPool *LiquidityPoolTransactor) AdjustPerpetualRiskParameter(opts *bind.TransactOpts, perpetualIndex *big.Int, key [32]byte, newValue *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "adjustPerpetualRiskParameter", perpetualIndex, key, newValue)
}

// AdjustPerpetualRiskParameter is a paid mutator transaction binding the contract method 0x882eaa77.
//
// Solidity: function adjustPerpetualRiskParameter(uint256 perpetualIndex, bytes32 key, int256 newValue) returns()
func (_LiquidityPool *LiquidityPoolSession) AdjustPerpetualRiskParameter(perpetualIndex *big.Int, key [32]byte, newValue *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.AdjustPerpetualRiskParameter(&_LiquidityPool.TransactOpts, perpetualIndex, key, newValue)
}

// AdjustPerpetualRiskParameter is a paid mutator transaction binding the contract method 0x882eaa77.
//
// Solidity: function adjustPerpetualRiskParameter(uint256 perpetualIndex, bytes32 key, int256 newValue) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) AdjustPerpetualRiskParameter(perpetualIndex *big.Int, key [32]byte, newValue *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.AdjustPerpetualRiskParameter(&_LiquidityPool.TransactOpts, perpetualIndex, key, newValue)
}

// BrokerTrade is a paid mutator transaction binding the contract method 0xd55bbd3b.
//
// Solidity: function brokerTrade(Order order, int256 amount, bytes signature) returns()
func (_LiquidityPool *LiquidityPoolTransactor) BrokerTrade(opts *bind.TransactOpts, order Order, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "brokerTrade", order, amount, signature)
}

// BrokerTrade is a paid mutator transaction binding the contract method 0xd55bbd3b.
//
// Solidity: function brokerTrade(Order order, int256 amount, bytes signature) returns()
func (_LiquidityPool *LiquidityPoolSession) BrokerTrade(order Order, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _LiquidityPool.Contract.BrokerTrade(&_LiquidityPool.TransactOpts, order, amount, signature)
}

// BrokerTrade is a paid mutator transaction binding the contract method 0xd55bbd3b.
//
// Solidity: function brokerTrade(Order order, int256 amount, bytes signature) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) BrokerTrade(order Order, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _LiquidityPool.Contract.BrokerTrade(&_LiquidityPool.TransactOpts, order, amount, signature)
}

// ClaimFee is a paid mutator transaction binding the contract method 0xeeaa184f.
//
// Solidity: function claimFee(int256 amount) returns()
func (_LiquidityPool *LiquidityPoolTransactor) ClaimFee(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "claimFee", amount)
}

// ClaimFee is a paid mutator transaction binding the contract method 0xeeaa184f.
//
// Solidity: function claimFee(int256 amount) returns()
func (_LiquidityPool *LiquidityPoolSession) ClaimFee(amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.ClaimFee(&_LiquidityPool.TransactOpts, amount)
}

// ClaimFee is a paid mutator transaction binding the contract method 0xeeaa184f.
//
// Solidity: function claimFee(int256 amount) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) ClaimFee(amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.ClaimFee(&_LiquidityPool.TransactOpts, amount)
}

// Clear is a paid mutator transaction binding the contract method 0xbbb364f2.
//
// Solidity: function clear(uint256 perpetualIndex, address trader) returns()
func (_LiquidityPool *LiquidityPoolTransactor) Clear(opts *bind.TransactOpts, perpetualIndex *big.Int, trader common.Address) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "clear", perpetualIndex, trader)
}

// Clear is a paid mutator transaction binding the contract method 0xbbb364f2.
//
// Solidity: function clear(uint256 perpetualIndex, address trader) returns()
func (_LiquidityPool *LiquidityPoolSession) Clear(perpetualIndex *big.Int, trader common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.Clear(&_LiquidityPool.TransactOpts, perpetualIndex, trader)
}

// Clear is a paid mutator transaction binding the contract method 0xbbb364f2.
//
// Solidity: function clear(uint256 perpetualIndex, address trader) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) Clear(perpetualIndex *big.Int, trader common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.Clear(&_LiquidityPool.TransactOpts, perpetualIndex, trader)
}

// CreatePerpetual is a paid mutator transaction binding the contract method 0x71c1d898.
//
// Solidity: function createPerpetual(address oracle, int256[8] coreParams, int256[5] riskParams, int256[5] minRiskParamValues, int256[5] maxRiskParamValues) returns()
func (_LiquidityPool *LiquidityPoolTransactor) CreatePerpetual(opts *bind.TransactOpts, oracle common.Address, coreParams [8]*big.Int, riskParams [5]*big.Int, minRiskParamValues [5]*big.Int, maxRiskParamValues [5]*big.Int) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "createPerpetual", oracle, coreParams, riskParams, minRiskParamValues, maxRiskParamValues)
}

// CreatePerpetual is a paid mutator transaction binding the contract method 0x71c1d898.
//
// Solidity: function createPerpetual(address oracle, int256[8] coreParams, int256[5] riskParams, int256[5] minRiskParamValues, int256[5] maxRiskParamValues) returns()
func (_LiquidityPool *LiquidityPoolSession) CreatePerpetual(oracle common.Address, coreParams [8]*big.Int, riskParams [5]*big.Int, minRiskParamValues [5]*big.Int, maxRiskParamValues [5]*big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.CreatePerpetual(&_LiquidityPool.TransactOpts, oracle, coreParams, riskParams, minRiskParamValues, maxRiskParamValues)
}

// CreatePerpetual is a paid mutator transaction binding the contract method 0x71c1d898.
//
// Solidity: function createPerpetual(address oracle, int256[8] coreParams, int256[5] riskParams, int256[5] minRiskParamValues, int256[5] maxRiskParamValues) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) CreatePerpetual(oracle common.Address, coreParams [8]*big.Int, riskParams [5]*big.Int, minRiskParamValues [5]*big.Int, maxRiskParamValues [5]*big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.CreatePerpetual(&_LiquidityPool.TransactOpts, oracle, coreParams, riskParams, minRiskParamValues, maxRiskParamValues)
}

// Deposit is a paid mutator transaction binding the contract method 0x78f140ea.
//
// Solidity: function deposit(uint256 perpetualIndex, address trader, int256 amount) returns()
func (_LiquidityPool *LiquidityPoolTransactor) Deposit(opts *bind.TransactOpts, perpetualIndex *big.Int, trader common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "deposit", perpetualIndex, trader, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x78f140ea.
//
// Solidity: function deposit(uint256 perpetualIndex, address trader, int256 amount) returns()
func (_LiquidityPool *LiquidityPoolSession) Deposit(perpetualIndex *big.Int, trader common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.Deposit(&_LiquidityPool.TransactOpts, perpetualIndex, trader, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x78f140ea.
//
// Solidity: function deposit(uint256 perpetualIndex, address trader, int256 amount) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) Deposit(perpetualIndex *big.Int, trader common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.Deposit(&_LiquidityPool.TransactOpts, perpetualIndex, trader, amount)
}

// DonateInsuranceFund is a paid mutator transaction binding the contract method 0x6fca8b99.
//
// Solidity: function donateInsuranceFund(int256 amount) returns()
func (_LiquidityPool *LiquidityPoolTransactor) DonateInsuranceFund(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "donateInsuranceFund", amount)
}

// DonateInsuranceFund is a paid mutator transaction binding the contract method 0x6fca8b99.
//
// Solidity: function donateInsuranceFund(int256 amount) returns()
func (_LiquidityPool *LiquidityPoolSession) DonateInsuranceFund(amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.DonateInsuranceFund(&_LiquidityPool.TransactOpts, amount)
}

// DonateInsuranceFund is a paid mutator transaction binding the contract method 0x6fca8b99.
//
// Solidity: function donateInsuranceFund(int256 amount) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) DonateInsuranceFund(amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.DonateInsuranceFund(&_LiquidityPool.TransactOpts, amount)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_LiquidityPool *LiquidityPoolTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_LiquidityPool *LiquidityPoolSession) Finalize() (*types.Transaction, error) {
	return _LiquidityPool.Contract.Finalize(&_LiquidityPool.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) Finalize() (*types.Transaction, error) {
	return _LiquidityPool.Contract.Finalize(&_LiquidityPool.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address operator, address collateral, address governor, address shareToken) returns()
func (_LiquidityPool *LiquidityPoolTransactor) Initialize(opts *bind.TransactOpts, operator common.Address, collateral common.Address, governor common.Address, shareToken common.Address) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "initialize", operator, collateral, governor, shareToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address operator, address collateral, address governor, address shareToken) returns()
func (_LiquidityPool *LiquidityPoolSession) Initialize(operator common.Address, collateral common.Address, governor common.Address, shareToken common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.Initialize(&_LiquidityPool.TransactOpts, operator, collateral, governor, shareToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address operator, address collateral, address governor, address shareToken) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) Initialize(operator common.Address, collateral common.Address, governor common.Address, shareToken common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.Initialize(&_LiquidityPool.TransactOpts, operator, collateral, governor, shareToken)
}

// LiquidateByAMM is a paid mutator transaction binding the contract method 0x650bfc56.
//
// Solidity: function liquidateByAMM(uint256 perpetualIndex, address trader, uint256 deadline) returns()
func (_LiquidityPool *LiquidityPoolTransactor) LiquidateByAMM(opts *bind.TransactOpts, perpetualIndex *big.Int, trader common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "liquidateByAMM", perpetualIndex, trader, deadline)
}

// LiquidateByAMM is a paid mutator transaction binding the contract method 0x650bfc56.
//
// Solidity: function liquidateByAMM(uint256 perpetualIndex, address trader, uint256 deadline) returns()
func (_LiquidityPool *LiquidityPoolSession) LiquidateByAMM(perpetualIndex *big.Int, trader common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.LiquidateByAMM(&_LiquidityPool.TransactOpts, perpetualIndex, trader, deadline)
}

// LiquidateByAMM is a paid mutator transaction binding the contract method 0x650bfc56.
//
// Solidity: function liquidateByAMM(uint256 perpetualIndex, address trader, uint256 deadline) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) LiquidateByAMM(perpetualIndex *big.Int, trader common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.LiquidateByAMM(&_LiquidityPool.TransactOpts, perpetualIndex, trader, deadline)
}

// LiquidateByTrader is a paid mutator transaction binding the contract method 0xa74d270c.
//
// Solidity: function liquidateByTrader(uint256 perpetualIndex, address trader, int256 amount, int256 priceLimit, uint256 deadline) returns()
func (_LiquidityPool *LiquidityPoolTransactor) LiquidateByTrader(opts *bind.TransactOpts, perpetualIndex *big.Int, trader common.Address, amount *big.Int, priceLimit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "liquidateByTrader", perpetualIndex, trader, amount, priceLimit, deadline)
}

// LiquidateByTrader is a paid mutator transaction binding the contract method 0xa74d270c.
//
// Solidity: function liquidateByTrader(uint256 perpetualIndex, address trader, int256 amount, int256 priceLimit, uint256 deadline) returns()
func (_LiquidityPool *LiquidityPoolSession) LiquidateByTrader(perpetualIndex *big.Int, trader common.Address, amount *big.Int, priceLimit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.LiquidateByTrader(&_LiquidityPool.TransactOpts, perpetualIndex, trader, amount, priceLimit, deadline)
}

// LiquidateByTrader is a paid mutator transaction binding the contract method 0xa74d270c.
//
// Solidity: function liquidateByTrader(uint256 perpetualIndex, address trader, int256 amount, int256 priceLimit, uint256 deadline) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) LiquidateByTrader(perpetualIndex *big.Int, trader common.Address, amount *big.Int, priceLimit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.LiquidateByTrader(&_LiquidityPool.TransactOpts, perpetualIndex, trader, amount, priceLimit, deadline)
}

// PerpetualInfo is a paid mutator transaction binding the contract method 0x4562d3d0.
//
// Solidity: function perpetualInfo(uint256 perpetualIndex) returns(uint8 state, address oracle, int256[17] nums)
func (_LiquidityPool *LiquidityPoolTransactor) PerpetualInfo(opts *bind.TransactOpts, perpetualIndex *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "perpetualInfo", perpetualIndex)
}

// PerpetualInfo is a paid mutator transaction binding the contract method 0x4562d3d0.
//
// Solidity: function perpetualInfo(uint256 perpetualIndex) returns(uint8 state, address oracle, int256[17] nums)
func (_LiquidityPool *LiquidityPoolSession) PerpetualInfo(perpetualIndex *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.PerpetualInfo(&_LiquidityPool.TransactOpts, perpetualIndex)
}

// PerpetualInfo is a paid mutator transaction binding the contract method 0x4562d3d0.
//
// Solidity: function perpetualInfo(uint256 perpetualIndex) returns(uint8 state, address oracle, int256[17] nums)
func (_LiquidityPool *LiquidityPoolTransactorSession) PerpetualInfo(perpetualIndex *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.PerpetualInfo(&_LiquidityPool.TransactOpts, perpetualIndex)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xcd26551a.
//
// Solidity: function removeLiquidity(int256 shareToRemove) returns()
func (_LiquidityPool *LiquidityPoolTransactor) RemoveLiquidity(opts *bind.TransactOpts, shareToRemove *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "removeLiquidity", shareToRemove)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xcd26551a.
//
// Solidity: function removeLiquidity(int256 shareToRemove) returns()
func (_LiquidityPool *LiquidityPoolSession) RemoveLiquidity(shareToRemove *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.RemoveLiquidity(&_LiquidityPool.TransactOpts, shareToRemove)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xcd26551a.
//
// Solidity: function removeLiquidity(int256 shareToRemove) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) RemoveLiquidity(shareToRemove *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.RemoveLiquidity(&_LiquidityPool.TransactOpts, shareToRemove)
}

// Settle is a paid mutator transaction binding the contract method 0x962d1938.
//
// Solidity: function settle(uint256 perpetualIndex, address trader) returns()
func (_LiquidityPool *LiquidityPoolTransactor) Settle(opts *bind.TransactOpts, perpetualIndex *big.Int, trader common.Address) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "settle", perpetualIndex, trader)
}

// Settle is a paid mutator transaction binding the contract method 0x962d1938.
//
// Solidity: function settle(uint256 perpetualIndex, address trader) returns()
func (_LiquidityPool *LiquidityPoolSession) Settle(perpetualIndex *big.Int, trader common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.Settle(&_LiquidityPool.TransactOpts, perpetualIndex, trader)
}

// Settle is a paid mutator transaction binding the contract method 0x962d1938.
//
// Solidity: function settle(uint256 perpetualIndex, address trader) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) Settle(perpetualIndex *big.Int, trader common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.Settle(&_LiquidityPool.TransactOpts, perpetualIndex, trader)
}

// Trade is a paid mutator transaction binding the contract method 0xfe92fd69.
//
// Solidity: function trade(uint256 perpetualIndex, address trader, int256 amount, int256 priceLimit, uint256 deadline, address referrer, bool isCloseOnly) returns()
func (_LiquidityPool *LiquidityPoolTransactor) Trade(opts *bind.TransactOpts, perpetualIndex *big.Int, trader common.Address, amount *big.Int, priceLimit *big.Int, deadline *big.Int, referrer common.Address, isCloseOnly bool) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "trade", perpetualIndex, trader, amount, priceLimit, deadline, referrer, isCloseOnly)
}

// Trade is a paid mutator transaction binding the contract method 0xfe92fd69.
//
// Solidity: function trade(uint256 perpetualIndex, address trader, int256 amount, int256 priceLimit, uint256 deadline, address referrer, bool isCloseOnly) returns()
func (_LiquidityPool *LiquidityPoolSession) Trade(perpetualIndex *big.Int, trader common.Address, amount *big.Int, priceLimit *big.Int, deadline *big.Int, referrer common.Address, isCloseOnly bool) (*types.Transaction, error) {
	return _LiquidityPool.Contract.Trade(&_LiquidityPool.TransactOpts, perpetualIndex, trader, amount, priceLimit, deadline, referrer, isCloseOnly)
}

// Trade is a paid mutator transaction binding the contract method 0xfe92fd69.
//
// Solidity: function trade(uint256 perpetualIndex, address trader, int256 amount, int256 priceLimit, uint256 deadline, address referrer, bool isCloseOnly) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) Trade(perpetualIndex *big.Int, trader common.Address, amount *big.Int, priceLimit *big.Int, deadline *big.Int, referrer common.Address, isCloseOnly bool) (*types.Transaction, error) {
	return _LiquidityPool.Contract.Trade(&_LiquidityPool.TransactOpts, perpetualIndex, trader, amount, priceLimit, deadline, referrer, isCloseOnly)
}

// UpdateLiquidityPoolParameter is a paid mutator transaction binding the contract method 0x43e068c1.
//
// Solidity: function updateLiquidityPoolParameter(bytes32 key, int256 newValue) returns()
func (_LiquidityPool *LiquidityPoolTransactor) UpdateLiquidityPoolParameter(opts *bind.TransactOpts, key [32]byte, newValue *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "updateLiquidityPoolParameter", key, newValue)
}

// UpdateLiquidityPoolParameter is a paid mutator transaction binding the contract method 0x43e068c1.
//
// Solidity: function updateLiquidityPoolParameter(bytes32 key, int256 newValue) returns()
func (_LiquidityPool *LiquidityPoolSession) UpdateLiquidityPoolParameter(key [32]byte, newValue *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.UpdateLiquidityPoolParameter(&_LiquidityPool.TransactOpts, key, newValue)
}

// UpdateLiquidityPoolParameter is a paid mutator transaction binding the contract method 0x43e068c1.
//
// Solidity: function updateLiquidityPoolParameter(bytes32 key, int256 newValue) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) UpdateLiquidityPoolParameter(key [32]byte, newValue *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.UpdateLiquidityPoolParameter(&_LiquidityPool.TransactOpts, key, newValue)
}

// UpdatePerpetualParameter is a paid mutator transaction binding the contract method 0x3e7a121a.
//
// Solidity: function updatePerpetualParameter(uint256 perpetualIndex, bytes32 key, int256 newValue) returns()
func (_LiquidityPool *LiquidityPoolTransactor) UpdatePerpetualParameter(opts *bind.TransactOpts, perpetualIndex *big.Int, key [32]byte, newValue *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "updatePerpetualParameter", perpetualIndex, key, newValue)
}

// UpdatePerpetualParameter is a paid mutator transaction binding the contract method 0x3e7a121a.
//
// Solidity: function updatePerpetualParameter(uint256 perpetualIndex, bytes32 key, int256 newValue) returns()
func (_LiquidityPool *LiquidityPoolSession) UpdatePerpetualParameter(perpetualIndex *big.Int, key [32]byte, newValue *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.UpdatePerpetualParameter(&_LiquidityPool.TransactOpts, perpetualIndex, key, newValue)
}

// UpdatePerpetualParameter is a paid mutator transaction binding the contract method 0x3e7a121a.
//
// Solidity: function updatePerpetualParameter(uint256 perpetualIndex, bytes32 key, int256 newValue) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) UpdatePerpetualParameter(perpetualIndex *big.Int, key [32]byte, newValue *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.UpdatePerpetualParameter(&_LiquidityPool.TransactOpts, perpetualIndex, key, newValue)
}

// UpdatePerpetualRiskParameter is a paid mutator transaction binding the contract method 0x663e017a.
//
// Solidity: function updatePerpetualRiskParameter(uint256 perpetualIndex, bytes32 key, int256 newValue, int256 minValue, int256 maxValue) returns()
func (_LiquidityPool *LiquidityPoolTransactor) UpdatePerpetualRiskParameter(opts *bind.TransactOpts, perpetualIndex *big.Int, key [32]byte, newValue *big.Int, minValue *big.Int, maxValue *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "updatePerpetualRiskParameter", perpetualIndex, key, newValue, minValue, maxValue)
}

// UpdatePerpetualRiskParameter is a paid mutator transaction binding the contract method 0x663e017a.
//
// Solidity: function updatePerpetualRiskParameter(uint256 perpetualIndex, bytes32 key, int256 newValue, int256 minValue, int256 maxValue) returns()
func (_LiquidityPool *LiquidityPoolSession) UpdatePerpetualRiskParameter(perpetualIndex *big.Int, key [32]byte, newValue *big.Int, minValue *big.Int, maxValue *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.UpdatePerpetualRiskParameter(&_LiquidityPool.TransactOpts, perpetualIndex, key, newValue, minValue, maxValue)
}

// UpdatePerpetualRiskParameter is a paid mutator transaction binding the contract method 0x663e017a.
//
// Solidity: function updatePerpetualRiskParameter(uint256 perpetualIndex, bytes32 key, int256 newValue, int256 minValue, int256 maxValue) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) UpdatePerpetualRiskParameter(perpetualIndex *big.Int, key [32]byte, newValue *big.Int, minValue *big.Int, maxValue *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.UpdatePerpetualRiskParameter(&_LiquidityPool.TransactOpts, perpetualIndex, key, newValue, minValue, maxValue)
}

// Withdraw is a paid mutator transaction binding the contract method 0x6ef05a40.
//
// Solidity: function withdraw(uint256 perpetualIndex, address trader, int256 amount) returns()
func (_LiquidityPool *LiquidityPoolTransactor) Withdraw(opts *bind.TransactOpts, perpetualIndex *big.Int, trader common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "withdraw", perpetualIndex, trader, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x6ef05a40.
//
// Solidity: function withdraw(uint256 perpetualIndex, address trader, int256 amount) returns()
func (_LiquidityPool *LiquidityPoolSession) Withdraw(perpetualIndex *big.Int, trader common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.Withdraw(&_LiquidityPool.TransactOpts, perpetualIndex, trader, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x6ef05a40.
//
// Solidity: function withdraw(uint256 perpetualIndex, address trader, int256 amount) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) Withdraw(perpetualIndex *big.Int, trader common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.Withdraw(&_LiquidityPool.TransactOpts, perpetualIndex, trader, amount)
}

// LiquidityPoolAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the LiquidityPool contract.
type LiquidityPoolAddLiquidityIterator struct {
	Event *LiquidityPoolAddLiquidity // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolAddLiquidity)
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
		it.Event = new(LiquidityPoolAddLiquidity)
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
func (it *LiquidityPoolAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolAddLiquidity represents a AddLiquidity event raised by the LiquidityPool contract.
type LiquidityPoolAddLiquidity struct {
	Trader      common.Address
	AddedCash   *big.Int
	MintedShare *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x04f3570d3bae00aef882e41946412ebe2d1c55b473e18af97148c3dde1b08f85.
//
// Solidity: event AddLiquidity(address trader, int256 addedCash, int256 mintedShare)
func (_LiquidityPool *LiquidityPoolFilterer) FilterAddLiquidity(opts *bind.FilterOpts) (*LiquidityPoolAddLiquidityIterator, error) {

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "AddLiquidity")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolAddLiquidityIterator{contract: _LiquidityPool.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x04f3570d3bae00aef882e41946412ebe2d1c55b473e18af97148c3dde1b08f85.
//
// Solidity: event AddLiquidity(address trader, int256 addedCash, int256 mintedShare)
func (_LiquidityPool *LiquidityPoolFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *LiquidityPoolAddLiquidity) (event.Subscription, error) {

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "AddLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolAddLiquidity)
				if err := _LiquidityPool.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x04f3570d3bae00aef882e41946412ebe2d1c55b473e18af97148c3dde1b08f85.
//
// Solidity: event AddLiquidity(address trader, int256 addedCash, int256 mintedShare)
func (_LiquidityPool *LiquidityPoolFilterer) ParseAddLiquidity(log types.Log) (*LiquidityPoolAddLiquidity, error) {
	event := new(LiquidityPoolAddLiquidity)
	if err := _LiquidityPool.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityPoolAdjustPerpetualRiskSettingIterator is returned from FilterAdjustPerpetualRiskSetting and is used to iterate over the raw logs and unpacked data for AdjustPerpetualRiskSetting events raised by the LiquidityPool contract.
type LiquidityPoolAdjustPerpetualRiskSettingIterator struct {
	Event *LiquidityPoolAdjustPerpetualRiskSetting // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolAdjustPerpetualRiskSettingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolAdjustPerpetualRiskSetting)
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
		it.Event = new(LiquidityPoolAdjustPerpetualRiskSetting)
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
func (it *LiquidityPoolAdjustPerpetualRiskSettingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolAdjustPerpetualRiskSettingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolAdjustPerpetualRiskSetting represents a AdjustPerpetualRiskSetting event raised by the LiquidityPool contract.
type LiquidityPoolAdjustPerpetualRiskSetting struct {
	PerpetualIndex *big.Int
	Key            [32]byte
	Value          *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAdjustPerpetualRiskSetting is a free log retrieval operation binding the contract event 0x308c838c4da0fbc62d294d47b98b3c612a1a3c4efe5c8971714849cfe80da909.
//
// Solidity: event AdjustPerpetualRiskSetting(uint256 perpetualIndex, bytes32 key, int256 value)
func (_LiquidityPool *LiquidityPoolFilterer) FilterAdjustPerpetualRiskSetting(opts *bind.FilterOpts) (*LiquidityPoolAdjustPerpetualRiskSettingIterator, error) {

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "AdjustPerpetualRiskSetting")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolAdjustPerpetualRiskSettingIterator{contract: _LiquidityPool.contract, event: "AdjustPerpetualRiskSetting", logs: logs, sub: sub}, nil
}

// WatchAdjustPerpetualRiskSetting is a free log subscription operation binding the contract event 0x308c838c4da0fbc62d294d47b98b3c612a1a3c4efe5c8971714849cfe80da909.
//
// Solidity: event AdjustPerpetualRiskSetting(uint256 perpetualIndex, bytes32 key, int256 value)
func (_LiquidityPool *LiquidityPoolFilterer) WatchAdjustPerpetualRiskSetting(opts *bind.WatchOpts, sink chan<- *LiquidityPoolAdjustPerpetualRiskSetting) (event.Subscription, error) {

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "AdjustPerpetualRiskSetting")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolAdjustPerpetualRiskSetting)
				if err := _LiquidityPool.contract.UnpackLog(event, "AdjustPerpetualRiskSetting", log); err != nil {
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

// ParseAdjustPerpetualRiskSetting is a log parse operation binding the contract event 0x308c838c4da0fbc62d294d47b98b3c612a1a3c4efe5c8971714849cfe80da909.
//
// Solidity: event AdjustPerpetualRiskSetting(uint256 perpetualIndex, bytes32 key, int256 value)
func (_LiquidityPool *LiquidityPoolFilterer) ParseAdjustPerpetualRiskSetting(log types.Log) (*LiquidityPoolAdjustPerpetualRiskSetting, error) {
	event := new(LiquidityPoolAdjustPerpetualRiskSetting)
	if err := _LiquidityPool.contract.UnpackLog(event, "AdjustPerpetualRiskSetting", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityPoolClaimFeeIterator is returned from FilterClaimFee and is used to iterate over the raw logs and unpacked data for ClaimFee events raised by the LiquidityPool contract.
type LiquidityPoolClaimFeeIterator struct {
	Event *LiquidityPoolClaimFee // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolClaimFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolClaimFee)
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
		it.Event = new(LiquidityPoolClaimFee)
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
func (it *LiquidityPoolClaimFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolClaimFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolClaimFee represents a ClaimFee event raised by the LiquidityPool contract.
type LiquidityPoolClaimFee struct {
	Claimer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimFee is a free log retrieval operation binding the contract event 0x10688f72abe281c3b2c60342fd825674ce2e3d773fe911bc9d96efa899d0109d.
//
// Solidity: event ClaimFee(address claimer, int256 amount)
func (_LiquidityPool *LiquidityPoolFilterer) FilterClaimFee(opts *bind.FilterOpts) (*LiquidityPoolClaimFeeIterator, error) {

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "ClaimFee")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolClaimFeeIterator{contract: _LiquidityPool.contract, event: "ClaimFee", logs: logs, sub: sub}, nil
}

// WatchClaimFee is a free log subscription operation binding the contract event 0x10688f72abe281c3b2c60342fd825674ce2e3d773fe911bc9d96efa899d0109d.
//
// Solidity: event ClaimFee(address claimer, int256 amount)
func (_LiquidityPool *LiquidityPoolFilterer) WatchClaimFee(opts *bind.WatchOpts, sink chan<- *LiquidityPoolClaimFee) (event.Subscription, error) {

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "ClaimFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolClaimFee)
				if err := _LiquidityPool.contract.UnpackLog(event, "ClaimFee", log); err != nil {
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
func (_LiquidityPool *LiquidityPoolFilterer) ParseClaimFee(log types.Log) (*LiquidityPoolClaimFee, error) {
	event := new(LiquidityPoolClaimFee)
	if err := _LiquidityPool.contract.UnpackLog(event, "ClaimFee", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityPoolClearIterator is returned from FilterClear and is used to iterate over the raw logs and unpacked data for Clear events raised by the LiquidityPool contract.
type LiquidityPoolClearIterator struct {
	Event *LiquidityPoolClear // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolClearIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolClear)
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
		it.Event = new(LiquidityPoolClear)
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
func (it *LiquidityPoolClearIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolClearIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolClear represents a Clear event raised by the LiquidityPool contract.
type LiquidityPoolClear struct {
	PerpetualIndex *big.Int
	Trader         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterClear is a free log retrieval operation binding the contract event 0x57bee722e110bffffed7bde2bceab2452cea26c7a9bb412059063dd37bdbb7e3.
//
// Solidity: event Clear(uint256 perpetualIndex, address trader)
func (_LiquidityPool *LiquidityPoolFilterer) FilterClear(opts *bind.FilterOpts) (*LiquidityPoolClearIterator, error) {

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "Clear")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolClearIterator{contract: _LiquidityPool.contract, event: "Clear", logs: logs, sub: sub}, nil
}

// WatchClear is a free log subscription operation binding the contract event 0x57bee722e110bffffed7bde2bceab2452cea26c7a9bb412059063dd37bdbb7e3.
//
// Solidity: event Clear(uint256 perpetualIndex, address trader)
func (_LiquidityPool *LiquidityPoolFilterer) WatchClear(opts *bind.WatchOpts, sink chan<- *LiquidityPoolClear) (event.Subscription, error) {

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "Clear")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolClear)
				if err := _LiquidityPool.contract.UnpackLog(event, "Clear", log); err != nil {
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

// ParseClear is a log parse operation binding the contract event 0x57bee722e110bffffed7bde2bceab2452cea26c7a9bb412059063dd37bdbb7e3.
//
// Solidity: event Clear(uint256 perpetualIndex, address trader)
func (_LiquidityPool *LiquidityPoolFilterer) ParseClear(log types.Log) (*LiquidityPoolClear, error) {
	event := new(LiquidityPoolClear)
	if err := _LiquidityPool.contract.UnpackLog(event, "Clear", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityPoolCreatePerpetualIterator is returned from FilterCreatePerpetual and is used to iterate over the raw logs and unpacked data for CreatePerpetual events raised by the LiquidityPool contract.
type LiquidityPoolCreatePerpetualIterator struct {
	Event *LiquidityPoolCreatePerpetual // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolCreatePerpetualIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolCreatePerpetual)
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
		it.Event = new(LiquidityPoolCreatePerpetual)
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
func (it *LiquidityPoolCreatePerpetualIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolCreatePerpetualIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolCreatePerpetual represents a CreatePerpetual event raised by the LiquidityPool contract.
type LiquidityPoolCreatePerpetual struct {
	PerpetualIndex *big.Int
	Governor       common.Address
	ShareToken     common.Address
	Operator       common.Address
	Oracle         common.Address
	Collateral     common.Address
	CoreParams     [8]*big.Int
	RiskParams     [5]*big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCreatePerpetual is a free log retrieval operation binding the contract event 0xdc99186d975bfd4c4da0121cef3e940d589fade4fd503e37b7f9a0d08f73fe16.
//
// Solidity: event CreatePerpetual(uint256 perpetualIndex, address governor, address shareToken, address operator, address oracle, address collateral, int256[8] coreParams, int256[5] riskParams)
func (_LiquidityPool *LiquidityPoolFilterer) FilterCreatePerpetual(opts *bind.FilterOpts) (*LiquidityPoolCreatePerpetualIterator, error) {

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "CreatePerpetual")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolCreatePerpetualIterator{contract: _LiquidityPool.contract, event: "CreatePerpetual", logs: logs, sub: sub}, nil
}

// WatchCreatePerpetual is a free log subscription operation binding the contract event 0xdc99186d975bfd4c4da0121cef3e940d589fade4fd503e37b7f9a0d08f73fe16.
//
// Solidity: event CreatePerpetual(uint256 perpetualIndex, address governor, address shareToken, address operator, address oracle, address collateral, int256[8] coreParams, int256[5] riskParams)
func (_LiquidityPool *LiquidityPoolFilterer) WatchCreatePerpetual(opts *bind.WatchOpts, sink chan<- *LiquidityPoolCreatePerpetual) (event.Subscription, error) {

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "CreatePerpetual")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolCreatePerpetual)
				if err := _LiquidityPool.contract.UnpackLog(event, "CreatePerpetual", log); err != nil {
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

// ParseCreatePerpetual is a log parse operation binding the contract event 0xdc99186d975bfd4c4da0121cef3e940d589fade4fd503e37b7f9a0d08f73fe16.
//
// Solidity: event CreatePerpetual(uint256 perpetualIndex, address governor, address shareToken, address operator, address oracle, address collateral, int256[8] coreParams, int256[5] riskParams)
func (_LiquidityPool *LiquidityPoolFilterer) ParseCreatePerpetual(log types.Log) (*LiquidityPoolCreatePerpetual, error) {
	event := new(LiquidityPoolCreatePerpetual)
	if err := _LiquidityPool.contract.UnpackLog(event, "CreatePerpetual", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityPoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the LiquidityPool contract.
type LiquidityPoolDepositIterator struct {
	Event *LiquidityPoolDeposit // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolDeposit)
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
		it.Event = new(LiquidityPoolDeposit)
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
func (it *LiquidityPoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolDeposit represents a Deposit event raised by the LiquidityPool contract.
type LiquidityPoolDeposit struct {
	PerpetualIndex *big.Int
	Trader         common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xd778d84255d1b73e8206f8b933689456acbbac9b48c1d13448aa204215969c54.
//
// Solidity: event Deposit(uint256 perpetualIndex, address trader, int256 amount)
func (_LiquidityPool *LiquidityPoolFilterer) FilterDeposit(opts *bind.FilterOpts) (*LiquidityPoolDepositIterator, error) {

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolDepositIterator{contract: _LiquidityPool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xd778d84255d1b73e8206f8b933689456acbbac9b48c1d13448aa204215969c54.
//
// Solidity: event Deposit(uint256 perpetualIndex, address trader, int256 amount)
func (_LiquidityPool *LiquidityPoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *LiquidityPoolDeposit) (event.Subscription, error) {

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolDeposit)
				if err := _LiquidityPool.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xd778d84255d1b73e8206f8b933689456acbbac9b48c1d13448aa204215969c54.
//
// Solidity: event Deposit(uint256 perpetualIndex, address trader, int256 amount)
func (_LiquidityPool *LiquidityPoolFilterer) ParseDeposit(log types.Log) (*LiquidityPoolDeposit, error) {
	event := new(LiquidityPoolDeposit)
	if err := _LiquidityPool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityPoolDonateInsuranceFundIterator is returned from FilterDonateInsuranceFund and is used to iterate over the raw logs and unpacked data for DonateInsuranceFund events raised by the LiquidityPool contract.
type LiquidityPoolDonateInsuranceFundIterator struct {
	Event *LiquidityPoolDonateInsuranceFund // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolDonateInsuranceFundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolDonateInsuranceFund)
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
		it.Event = new(LiquidityPoolDonateInsuranceFund)
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
func (it *LiquidityPoolDonateInsuranceFundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolDonateInsuranceFundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolDonateInsuranceFund represents a DonateInsuranceFund event raised by the LiquidityPool contract.
type LiquidityPoolDonateInsuranceFund struct {
	Trader common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDonateInsuranceFund is a free log retrieval operation binding the contract event 0x076a19d3bb1fbcd6fdee9888f4e0ab00a11cb2c4400a1d545eaeaf6bb5c25bc6.
//
// Solidity: event DonateInsuranceFund(address trader, int256 amount)
func (_LiquidityPool *LiquidityPoolFilterer) FilterDonateInsuranceFund(opts *bind.FilterOpts) (*LiquidityPoolDonateInsuranceFundIterator, error) {

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "DonateInsuranceFund")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolDonateInsuranceFundIterator{contract: _LiquidityPool.contract, event: "DonateInsuranceFund", logs: logs, sub: sub}, nil
}

// WatchDonateInsuranceFund is a free log subscription operation binding the contract event 0x076a19d3bb1fbcd6fdee9888f4e0ab00a11cb2c4400a1d545eaeaf6bb5c25bc6.
//
// Solidity: event DonateInsuranceFund(address trader, int256 amount)
func (_LiquidityPool *LiquidityPoolFilterer) WatchDonateInsuranceFund(opts *bind.WatchOpts, sink chan<- *LiquidityPoolDonateInsuranceFund) (event.Subscription, error) {

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "DonateInsuranceFund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolDonateInsuranceFund)
				if err := _LiquidityPool.contract.UnpackLog(event, "DonateInsuranceFund", log); err != nil {
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
func (_LiquidityPool *LiquidityPoolFilterer) ParseDonateInsuranceFund(log types.Log) (*LiquidityPoolDonateInsuranceFund, error) {
	event := new(LiquidityPoolDonateInsuranceFund)
	if err := _LiquidityPool.contract.UnpackLog(event, "DonateInsuranceFund", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityPoolFinalizeIterator is returned from FilterFinalize and is used to iterate over the raw logs and unpacked data for Finalize events raised by the LiquidityPool contract.
type LiquidityPoolFinalizeIterator struct {
	Event *LiquidityPoolFinalize // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolFinalizeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolFinalize)
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
		it.Event = new(LiquidityPoolFinalize)
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
func (it *LiquidityPoolFinalizeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolFinalizeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolFinalize represents a Finalize event raised by the LiquidityPool contract.
type LiquidityPoolFinalize struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFinalize is a free log retrieval operation binding the contract event 0xc5454d111913d0c92fa9088b73be5c3fc91d1eb84db52a8a8485154f05d73f2e.
//
// Solidity: event Finalize()
func (_LiquidityPool *LiquidityPoolFilterer) FilterFinalize(opts *bind.FilterOpts) (*LiquidityPoolFinalizeIterator, error) {

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "Finalize")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolFinalizeIterator{contract: _LiquidityPool.contract, event: "Finalize", logs: logs, sub: sub}, nil
}

// WatchFinalize is a free log subscription operation binding the contract event 0xc5454d111913d0c92fa9088b73be5c3fc91d1eb84db52a8a8485154f05d73f2e.
//
// Solidity: event Finalize()
func (_LiquidityPool *LiquidityPoolFilterer) WatchFinalize(opts *bind.WatchOpts, sink chan<- *LiquidityPoolFinalize) (event.Subscription, error) {

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "Finalize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolFinalize)
				if err := _LiquidityPool.contract.UnpackLog(event, "Finalize", log); err != nil {
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

// ParseFinalize is a log parse operation binding the contract event 0xc5454d111913d0c92fa9088b73be5c3fc91d1eb84db52a8a8485154f05d73f2e.
//
// Solidity: event Finalize()
func (_LiquidityPool *LiquidityPoolFilterer) ParseFinalize(log types.Log) (*LiquidityPoolFinalize, error) {
	event := new(LiquidityPoolFinalize)
	if err := _LiquidityPool.contract.UnpackLog(event, "Finalize", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityPoolLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the LiquidityPool contract.
type LiquidityPoolLiquidateIterator struct {
	Event *LiquidityPoolLiquidate // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolLiquidate)
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
		it.Event = new(LiquidityPoolLiquidate)
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
func (it *LiquidityPoolLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolLiquidate represents a Liquidate event raised by the LiquidityPool contract.
type LiquidityPoolLiquidate struct {
	PerpetualIndex *big.Int
	Liquidator     common.Address
	Trader         common.Address
	Amount         *big.Int
	Price          *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0xc8f9ffc2832736f60661d787f9e35925e293ade314949b157fc13899574011d1.
//
// Solidity: event Liquidate(uint256 perpetualIndex, address indexed liquidator, address indexed trader, int256 amount, int256 price)
func (_LiquidityPool *LiquidityPoolFilterer) FilterLiquidate(opts *bind.FilterOpts, liquidator []common.Address, trader []common.Address) (*LiquidityPoolLiquidateIterator, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "Liquidate", liquidatorRule, traderRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolLiquidateIterator{contract: _LiquidityPool.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0xc8f9ffc2832736f60661d787f9e35925e293ade314949b157fc13899574011d1.
//
// Solidity: event Liquidate(uint256 perpetualIndex, address indexed liquidator, address indexed trader, int256 amount, int256 price)
func (_LiquidityPool *LiquidityPoolFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *LiquidityPoolLiquidate, liquidator []common.Address, trader []common.Address) (event.Subscription, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "Liquidate", liquidatorRule, traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolLiquidate)
				if err := _LiquidityPool.contract.UnpackLog(event, "Liquidate", log); err != nil {
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

// ParseLiquidate is a log parse operation binding the contract event 0xc8f9ffc2832736f60661d787f9e35925e293ade314949b157fc13899574011d1.
//
// Solidity: event Liquidate(uint256 perpetualIndex, address indexed liquidator, address indexed trader, int256 amount, int256 price)
func (_LiquidityPool *LiquidityPoolFilterer) ParseLiquidate(log types.Log) (*LiquidityPoolLiquidate, error) {
	event := new(LiquidityPoolLiquidate)
	if err := _LiquidityPool.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityPoolRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the LiquidityPool contract.
type LiquidityPoolRemoveLiquidityIterator struct {
	Event *LiquidityPoolRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolRemoveLiquidity)
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
		it.Event = new(LiquidityPoolRemoveLiquidity)
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
func (it *LiquidityPoolRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolRemoveLiquidity represents a RemoveLiquidity event raised by the LiquidityPool contract.
type LiquidityPoolRemoveLiquidity struct {
	Trader       common.Address
	ReturnedCash *big.Int
	BurnedShare  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x01f9cda55a171066e2652d83c252608e17b9f8a22400477e57065c51682133ab.
//
// Solidity: event RemoveLiquidity(address trader, int256 returnedCash, int256 burnedShare)
func (_LiquidityPool *LiquidityPoolFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts) (*LiquidityPoolRemoveLiquidityIterator, error) {

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "RemoveLiquidity")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolRemoveLiquidityIterator{contract: _LiquidityPool.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x01f9cda55a171066e2652d83c252608e17b9f8a22400477e57065c51682133ab.
//
// Solidity: event RemoveLiquidity(address trader, int256 returnedCash, int256 burnedShare)
func (_LiquidityPool *LiquidityPoolFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *LiquidityPoolRemoveLiquidity) (event.Subscription, error) {

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "RemoveLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolRemoveLiquidity)
				if err := _LiquidityPool.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0x01f9cda55a171066e2652d83c252608e17b9f8a22400477e57065c51682133ab.
//
// Solidity: event RemoveLiquidity(address trader, int256 returnedCash, int256 burnedShare)
func (_LiquidityPool *LiquidityPoolFilterer) ParseRemoveLiquidity(log types.Log) (*LiquidityPoolRemoveLiquidity, error) {
	event := new(LiquidityPoolRemoveLiquidity)
	if err := _LiquidityPool.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityPoolSettleIterator is returned from FilterSettle and is used to iterate over the raw logs and unpacked data for Settle events raised by the LiquidityPool contract.
type LiquidityPoolSettleIterator struct {
	Event *LiquidityPoolSettle // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolSettle)
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
		it.Event = new(LiquidityPoolSettle)
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
func (it *LiquidityPoolSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolSettle represents a Settle event raised by the LiquidityPool contract.
type LiquidityPoolSettle struct {
	PerpetualIndex *big.Int
	Trader         common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSettle is a free log retrieval operation binding the contract event 0x632450284489b209b98b764aad06c5d46017343cca821d115314e9f085f82355.
//
// Solidity: event Settle(uint256 perpetualIndex, address trader, int256 amount)
func (_LiquidityPool *LiquidityPoolFilterer) FilterSettle(opts *bind.FilterOpts) (*LiquidityPoolSettleIterator, error) {

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "Settle")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolSettleIterator{contract: _LiquidityPool.contract, event: "Settle", logs: logs, sub: sub}, nil
}

// WatchSettle is a free log subscription operation binding the contract event 0x632450284489b209b98b764aad06c5d46017343cca821d115314e9f085f82355.
//
// Solidity: event Settle(uint256 perpetualIndex, address trader, int256 amount)
func (_LiquidityPool *LiquidityPoolFilterer) WatchSettle(opts *bind.WatchOpts, sink chan<- *LiquidityPoolSettle) (event.Subscription, error) {

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "Settle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolSettle)
				if err := _LiquidityPool.contract.UnpackLog(event, "Settle", log); err != nil {
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

// ParseSettle is a log parse operation binding the contract event 0x632450284489b209b98b764aad06c5d46017343cca821d115314e9f085f82355.
//
// Solidity: event Settle(uint256 perpetualIndex, address trader, int256 amount)
func (_LiquidityPool *LiquidityPoolFilterer) ParseSettle(log types.Log) (*LiquidityPoolSettle, error) {
	event := new(LiquidityPoolSettle)
	if err := _LiquidityPool.contract.UnpackLog(event, "Settle", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityPoolTradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the LiquidityPool contract.
type LiquidityPoolTradeIterator struct {
	Event *LiquidityPoolTrade // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolTrade)
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
		it.Event = new(LiquidityPoolTrade)
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
func (it *LiquidityPoolTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolTrade represents a Trade event raised by the LiquidityPool contract.
type LiquidityPoolTrade struct {
	PerpetualIndex *big.Int
	Trader         common.Address
	PositionAmount *big.Int
	Price          *big.Int
	Fee            *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0x4c45b3a2c402695ee5576ee9d35638db8221c2094be1dbc75362d265f00ff081.
//
// Solidity: event Trade(uint256 perpetualIndex, address indexed trader, int256 positionAmount, int256 price, int256 fee)
func (_LiquidityPool *LiquidityPoolFilterer) FilterTrade(opts *bind.FilterOpts, trader []common.Address) (*LiquidityPoolTradeIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "Trade", traderRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolTradeIterator{contract: _LiquidityPool.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0x4c45b3a2c402695ee5576ee9d35638db8221c2094be1dbc75362d265f00ff081.
//
// Solidity: event Trade(uint256 perpetualIndex, address indexed trader, int256 positionAmount, int256 price, int256 fee)
func (_LiquidityPool *LiquidityPoolFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *LiquidityPoolTrade, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "Trade", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolTrade)
				if err := _LiquidityPool.contract.UnpackLog(event, "Trade", log); err != nil {
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

// ParseTrade is a log parse operation binding the contract event 0x4c45b3a2c402695ee5576ee9d35638db8221c2094be1dbc75362d265f00ff081.
//
// Solidity: event Trade(uint256 perpetualIndex, address indexed trader, int256 positionAmount, int256 price, int256 fee)
func (_LiquidityPool *LiquidityPoolFilterer) ParseTrade(log types.Log) (*LiquidityPoolTrade, error) {
	event := new(LiquidityPoolTrade)
	if err := _LiquidityPool.contract.UnpackLog(event, "Trade", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityPoolUpdateLiquidityPoolParameterIterator is returned from FilterUpdateLiquidityPoolParameter and is used to iterate over the raw logs and unpacked data for UpdateLiquidityPoolParameter events raised by the LiquidityPool contract.
type LiquidityPoolUpdateLiquidityPoolParameterIterator struct {
	Event *LiquidityPoolUpdateLiquidityPoolParameter // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolUpdateLiquidityPoolParameterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolUpdateLiquidityPoolParameter)
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
		it.Event = new(LiquidityPoolUpdateLiquidityPoolParameter)
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
func (it *LiquidityPoolUpdateLiquidityPoolParameterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolUpdateLiquidityPoolParameterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolUpdateLiquidityPoolParameter represents a UpdateLiquidityPoolParameter event raised by the LiquidityPool contract.
type LiquidityPoolUpdateLiquidityPoolParameter struct {
	Key   [32]byte
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdateLiquidityPoolParameter is a free log retrieval operation binding the contract event 0xb6bac3474a1e91ded0abbb7b4c2e5601610ebe0aa3fcf36c9cdae981ea8d2eb8.
//
// Solidity: event UpdateLiquidityPoolParameter(bytes32 key, int256 value)
func (_LiquidityPool *LiquidityPoolFilterer) FilterUpdateLiquidityPoolParameter(opts *bind.FilterOpts) (*LiquidityPoolUpdateLiquidityPoolParameterIterator, error) {

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "UpdateLiquidityPoolParameter")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolUpdateLiquidityPoolParameterIterator{contract: _LiquidityPool.contract, event: "UpdateLiquidityPoolParameter", logs: logs, sub: sub}, nil
}

// WatchUpdateLiquidityPoolParameter is a free log subscription operation binding the contract event 0xb6bac3474a1e91ded0abbb7b4c2e5601610ebe0aa3fcf36c9cdae981ea8d2eb8.
//
// Solidity: event UpdateLiquidityPoolParameter(bytes32 key, int256 value)
func (_LiquidityPool *LiquidityPoolFilterer) WatchUpdateLiquidityPoolParameter(opts *bind.WatchOpts, sink chan<- *LiquidityPoolUpdateLiquidityPoolParameter) (event.Subscription, error) {

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "UpdateLiquidityPoolParameter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolUpdateLiquidityPoolParameter)
				if err := _LiquidityPool.contract.UnpackLog(event, "UpdateLiquidityPoolParameter", log); err != nil {
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

// ParseUpdateLiquidityPoolParameter is a log parse operation binding the contract event 0xb6bac3474a1e91ded0abbb7b4c2e5601610ebe0aa3fcf36c9cdae981ea8d2eb8.
//
// Solidity: event UpdateLiquidityPoolParameter(bytes32 key, int256 value)
func (_LiquidityPool *LiquidityPoolFilterer) ParseUpdateLiquidityPoolParameter(log types.Log) (*LiquidityPoolUpdateLiquidityPoolParameter, error) {
	event := new(LiquidityPoolUpdateLiquidityPoolParameter)
	if err := _LiquidityPool.contract.UnpackLog(event, "UpdateLiquidityPoolParameter", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityPoolUpdatePerpetualParameterIterator is returned from FilterUpdatePerpetualParameter and is used to iterate over the raw logs and unpacked data for UpdatePerpetualParameter events raised by the LiquidityPool contract.
type LiquidityPoolUpdatePerpetualParameterIterator struct {
	Event *LiquidityPoolUpdatePerpetualParameter // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolUpdatePerpetualParameterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolUpdatePerpetualParameter)
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
		it.Event = new(LiquidityPoolUpdatePerpetualParameter)
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
func (it *LiquidityPoolUpdatePerpetualParameterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolUpdatePerpetualParameterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolUpdatePerpetualParameter represents a UpdatePerpetualParameter event raised by the LiquidityPool contract.
type LiquidityPoolUpdatePerpetualParameter struct {
	PerpetualIndex *big.Int
	Key            [32]byte
	Value          *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdatePerpetualParameter is a free log retrieval operation binding the contract event 0xf801cb33bd95a4f489fe50413d2f2aaa6b47b47742febbeef10dea3860bfe591.
//
// Solidity: event UpdatePerpetualParameter(uint256 perpetualIndex, bytes32 key, int256 value)
func (_LiquidityPool *LiquidityPoolFilterer) FilterUpdatePerpetualParameter(opts *bind.FilterOpts) (*LiquidityPoolUpdatePerpetualParameterIterator, error) {

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "UpdatePerpetualParameter")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolUpdatePerpetualParameterIterator{contract: _LiquidityPool.contract, event: "UpdatePerpetualParameter", logs: logs, sub: sub}, nil
}

// WatchUpdatePerpetualParameter is a free log subscription operation binding the contract event 0xf801cb33bd95a4f489fe50413d2f2aaa6b47b47742febbeef10dea3860bfe591.
//
// Solidity: event UpdatePerpetualParameter(uint256 perpetualIndex, bytes32 key, int256 value)
func (_LiquidityPool *LiquidityPoolFilterer) WatchUpdatePerpetualParameter(opts *bind.WatchOpts, sink chan<- *LiquidityPoolUpdatePerpetualParameter) (event.Subscription, error) {

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "UpdatePerpetualParameter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolUpdatePerpetualParameter)
				if err := _LiquidityPool.contract.UnpackLog(event, "UpdatePerpetualParameter", log); err != nil {
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

// ParseUpdatePerpetualParameter is a log parse operation binding the contract event 0xf801cb33bd95a4f489fe50413d2f2aaa6b47b47742febbeef10dea3860bfe591.
//
// Solidity: event UpdatePerpetualParameter(uint256 perpetualIndex, bytes32 key, int256 value)
func (_LiquidityPool *LiquidityPoolFilterer) ParseUpdatePerpetualParameter(log types.Log) (*LiquidityPoolUpdatePerpetualParameter, error) {
	event := new(LiquidityPoolUpdatePerpetualParameter)
	if err := _LiquidityPool.contract.UnpackLog(event, "UpdatePerpetualParameter", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityPoolUpdatePerpetualRiskParameterIterator is returned from FilterUpdatePerpetualRiskParameter and is used to iterate over the raw logs and unpacked data for UpdatePerpetualRiskParameter events raised by the LiquidityPool contract.
type LiquidityPoolUpdatePerpetualRiskParameterIterator struct {
	Event *LiquidityPoolUpdatePerpetualRiskParameter // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolUpdatePerpetualRiskParameterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolUpdatePerpetualRiskParameter)
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
		it.Event = new(LiquidityPoolUpdatePerpetualRiskParameter)
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
func (it *LiquidityPoolUpdatePerpetualRiskParameterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolUpdatePerpetualRiskParameterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolUpdatePerpetualRiskParameter represents a UpdatePerpetualRiskParameter event raised by the LiquidityPool contract.
type LiquidityPoolUpdatePerpetualRiskParameter struct {
	PerpetualIndex *big.Int
	Key            [32]byte
	Value          *big.Int
	MinValue       *big.Int
	MaxValue       *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdatePerpetualRiskParameter is a free log retrieval operation binding the contract event 0x418f81a495d2ca3f2b1f7e0dbf81c2c6b512e26567c1ae1af234ebdff11c317a.
//
// Solidity: event UpdatePerpetualRiskParameter(uint256 perpetualIndex, bytes32 key, int256 value, int256 minValue, int256 maxValue)
func (_LiquidityPool *LiquidityPoolFilterer) FilterUpdatePerpetualRiskParameter(opts *bind.FilterOpts) (*LiquidityPoolUpdatePerpetualRiskParameterIterator, error) {

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "UpdatePerpetualRiskParameter")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolUpdatePerpetualRiskParameterIterator{contract: _LiquidityPool.contract, event: "UpdatePerpetualRiskParameter", logs: logs, sub: sub}, nil
}

// WatchUpdatePerpetualRiskParameter is a free log subscription operation binding the contract event 0x418f81a495d2ca3f2b1f7e0dbf81c2c6b512e26567c1ae1af234ebdff11c317a.
//
// Solidity: event UpdatePerpetualRiskParameter(uint256 perpetualIndex, bytes32 key, int256 value, int256 minValue, int256 maxValue)
func (_LiquidityPool *LiquidityPoolFilterer) WatchUpdatePerpetualRiskParameter(opts *bind.WatchOpts, sink chan<- *LiquidityPoolUpdatePerpetualRiskParameter) (event.Subscription, error) {

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "UpdatePerpetualRiskParameter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolUpdatePerpetualRiskParameter)
				if err := _LiquidityPool.contract.UnpackLog(event, "UpdatePerpetualRiskParameter", log); err != nil {
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

// ParseUpdatePerpetualRiskParameter is a log parse operation binding the contract event 0x418f81a495d2ca3f2b1f7e0dbf81c2c6b512e26567c1ae1af234ebdff11c317a.
//
// Solidity: event UpdatePerpetualRiskParameter(uint256 perpetualIndex, bytes32 key, int256 value, int256 minValue, int256 maxValue)
func (_LiquidityPool *LiquidityPoolFilterer) ParseUpdatePerpetualRiskParameter(log types.Log) (*LiquidityPoolUpdatePerpetualRiskParameter, error) {
	event := new(LiquidityPoolUpdatePerpetualRiskParameter)
	if err := _LiquidityPool.contract.UnpackLog(event, "UpdatePerpetualRiskParameter", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LiquidityPoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the LiquidityPool contract.
type LiquidityPoolWithdrawIterator struct {
	Event *LiquidityPoolWithdraw // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolWithdraw)
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
		it.Event = new(LiquidityPoolWithdraw)
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
func (it *LiquidityPoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolWithdraw represents a Withdraw event raised by the LiquidityPool contract.
type LiquidityPoolWithdraw struct {
	PerpetualIndex *big.Int
	Trader         common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xe197d5ea53e580dbe195873fc9d78d1b6582c3c76a3a019571613fafa130b492.
//
// Solidity: event Withdraw(uint256 perpetualIndex, address trader, int256 amount)
func (_LiquidityPool *LiquidityPoolFilterer) FilterWithdraw(opts *bind.FilterOpts) (*LiquidityPoolWithdrawIterator, error) {

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolWithdrawIterator{contract: _LiquidityPool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xe197d5ea53e580dbe195873fc9d78d1b6582c3c76a3a019571613fafa130b492.
//
// Solidity: event Withdraw(uint256 perpetualIndex, address trader, int256 amount)
func (_LiquidityPool *LiquidityPoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *LiquidityPoolWithdraw) (event.Subscription, error) {

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolWithdraw)
				if err := _LiquidityPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xe197d5ea53e580dbe195873fc9d78d1b6582c3c76a3a019571613fafa130b492.
//
// Solidity: event Withdraw(uint256 perpetualIndex, address trader, int256 amount)
func (_LiquidityPool *LiquidityPoolFilterer) ParseWithdraw(log types.Log) (*LiquidityPoolWithdraw, error) {
	event := new(LiquidityPoolWithdraw)
	if err := _LiquidityPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
