// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package factory

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

// FactoryABI is the input ABI used to generate the binding from.
const FactoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"shareToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"perpetual\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"vaultFeeRate_\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"AddVersion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"perpetual\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"shareToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256[7]\",\"name\":\"coreParams\",\"type\":\"int256[7]\"},{\"indexed\":false,\"internalType\":\"int256[5]\",\"name\":\"riskParams\",\"type\":\"int256[5]\"}],\"name\":\"CreatePerpetual\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"RevokeVersion\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"activeProxy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"int256[7]\",\"name\":\"coreParams\",\"type\":\"int256[7]\"},{\"internalType\":\"int256[5]\",\"name\":\"riskParams\",\"type\":\"int256[5]\"},{\"internalType\":\"int256[5]\",\"name\":\"minRiskParamValues\",\"type\":\"int256[5]\"},{\"internalType\":\"int256[5]\",\"name\":\"maxRiskParamValues\",\"type\":\"int256[5]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"createPerpetual\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"deactiveProxy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"listPerpetuals\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPerpetualCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultFeeRate\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Factory is an auto generated Go binding around an Ethereum contract.
type Factory struct {
	FactoryCaller     // Read-only binding to the contract
	FactoryTransactor // Write-only binding to the contract
	FactoryFilterer   // Log filterer for contract events
}

// FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactorySession struct {
	Contract     *Factory          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactoryCallerSession struct {
	Contract *FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactoryTransactorSession struct {
	Contract     *FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactoryRaw struct {
	Contract *Factory // Generic contract binding to access the raw methods on
}

// FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactoryCallerRaw struct {
	Contract *FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactoryTransactorRaw struct {
	Contract *FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactory creates a new instance of Factory, bound to a specific deployed contract.
func NewFactory(address common.Address, backend bind.ContractBackend) (*Factory, error) {
	contract, err := bindFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// NewFactoryCaller creates a new read-only instance of Factory, bound to a specific deployed contract.
func NewFactoryCaller(address common.Address, caller bind.ContractCaller) (*FactoryCaller, error) {
	contract, err := bindFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryCaller{contract: contract}, nil
}

// NewFactoryTransactor creates a new write-only instance of Factory, bound to a specific deployed contract.
func NewFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FactoryTransactor, error) {
	contract, err := bindFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryTransactor{contract: contract}, nil
}

// NewFactoryFilterer creates a new log filterer instance of Factory, bound to a specific deployed contract.
func NewFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FactoryFilterer, error) {
	contract, err := bindFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryFilterer{contract: contract}, nil
}

// bindFactory binds a generic wrapper to an already deployed contract.
func bindFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transact(opts, method, params...)
}

// ListPerpetuals is a free data retrieval call binding the contract method 0x670e682c.
//
// Solidity: function listPerpetuals(uint256 begin, uint256 end) constant returns(address[])
func (_Factory *FactoryCaller) ListPerpetuals(opts *bind.CallOpts, begin *big.Int, end *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Factory.contract.Call(opts, out, "listPerpetuals", begin, end)
	return *ret0, err
}

// ListPerpetuals is a free data retrieval call binding the contract method 0x670e682c.
//
// Solidity: function listPerpetuals(uint256 begin, uint256 end) constant returns(address[])
func (_Factory *FactorySession) ListPerpetuals(begin *big.Int, end *big.Int) ([]common.Address, error) {
	return _Factory.Contract.ListPerpetuals(&_Factory.CallOpts, begin, end)
}

// ListPerpetuals is a free data retrieval call binding the contract method 0x670e682c.
//
// Solidity: function listPerpetuals(uint256 begin, uint256 end) constant returns(address[])
func (_Factory *FactoryCallerSession) ListPerpetuals(begin *big.Int, end *big.Int) ([]common.Address, error) {
	return _Factory.Contract.ListPerpetuals(&_Factory.CallOpts, begin, end)
}

// TotalPerpetualCount is a free data retrieval call binding the contract method 0x49789c17.
//
// Solidity: function totalPerpetualCount() constant returns(uint256)
func (_Factory *FactoryCaller) TotalPerpetualCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Factory.contract.Call(opts, out, "totalPerpetualCount")
	return *ret0, err
}

// TotalPerpetualCount is a free data retrieval call binding the contract method 0x49789c17.
//
// Solidity: function totalPerpetualCount() constant returns(uint256)
func (_Factory *FactorySession) TotalPerpetualCount() (*big.Int, error) {
	return _Factory.Contract.TotalPerpetualCount(&_Factory.CallOpts)
}

// TotalPerpetualCount is a free data retrieval call binding the contract method 0x49789c17.
//
// Solidity: function totalPerpetualCount() constant returns(uint256)
func (_Factory *FactoryCallerSession) TotalPerpetualCount() (*big.Int, error) {
	return _Factory.Contract.TotalPerpetualCount(&_Factory.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() constant returns(address)
func (_Factory *FactoryCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Factory.contract.Call(opts, out, "vault")
	return *ret0, err
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() constant returns(address)
func (_Factory *FactorySession) Vault() (common.Address, error) {
	return _Factory.Contract.Vault(&_Factory.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() constant returns(address)
func (_Factory *FactoryCallerSession) Vault() (common.Address, error) {
	return _Factory.Contract.Vault(&_Factory.CallOpts)
}

// VaultFeeRate is a free data retrieval call binding the contract method 0x96550e8b.
//
// Solidity: function vaultFeeRate() constant returns(int256)
func (_Factory *FactoryCaller) VaultFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Factory.contract.Call(opts, out, "vaultFeeRate")
	return *ret0, err
}

// VaultFeeRate is a free data retrieval call binding the contract method 0x96550e8b.
//
// Solidity: function vaultFeeRate() constant returns(int256)
func (_Factory *FactorySession) VaultFeeRate() (*big.Int, error) {
	return _Factory.Contract.VaultFeeRate(&_Factory.CallOpts)
}

// VaultFeeRate is a free data retrieval call binding the contract method 0x96550e8b.
//
// Solidity: function vaultFeeRate() constant returns(int256)
func (_Factory *FactoryCallerSession) VaultFeeRate() (*big.Int, error) {
	return _Factory.Contract.VaultFeeRate(&_Factory.CallOpts)
}

// ActiveProxy is a paid mutator transaction binding the contract method 0xd3b3eb2b.
//
// Solidity: function activeProxy(address trader) returns(bool)
func (_Factory *FactoryTransactor) ActiveProxy(opts *bind.TransactOpts, trader common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "activeProxy", trader)
}

// ActiveProxy is a paid mutator transaction binding the contract method 0xd3b3eb2b.
//
// Solidity: function activeProxy(address trader) returns(bool)
func (_Factory *FactorySession) ActiveProxy(trader common.Address) (*types.Transaction, error) {
	return _Factory.Contract.ActiveProxy(&_Factory.TransactOpts, trader)
}

// ActiveProxy is a paid mutator transaction binding the contract method 0xd3b3eb2b.
//
// Solidity: function activeProxy(address trader) returns(bool)
func (_Factory *FactoryTransactorSession) ActiveProxy(trader common.Address) (*types.Transaction, error) {
	return _Factory.Contract.ActiveProxy(&_Factory.TransactOpts, trader)
}

// CreatePerpetual is a paid mutator transaction binding the contract method 0x53a1c053.
//
// Solidity: function createPerpetual(address oracle, int256[7] coreParams, int256[5] riskParams, int256[5] minRiskParamValues, int256[5] maxRiskParamValues, uint256 nonce) returns(address)
func (_Factory *FactoryTransactor) CreatePerpetual(opts *bind.TransactOpts, oracle common.Address, coreParams [7]*big.Int, riskParams [5]*big.Int, minRiskParamValues [5]*big.Int, maxRiskParamValues [5]*big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "createPerpetual", oracle, coreParams, riskParams, minRiskParamValues, maxRiskParamValues, nonce)
}

// CreatePerpetual is a paid mutator transaction binding the contract method 0x53a1c053.
//
// Solidity: function createPerpetual(address oracle, int256[7] coreParams, int256[5] riskParams, int256[5] minRiskParamValues, int256[5] maxRiskParamValues, uint256 nonce) returns(address)
func (_Factory *FactorySession) CreatePerpetual(oracle common.Address, coreParams [7]*big.Int, riskParams [5]*big.Int, minRiskParamValues [5]*big.Int, maxRiskParamValues [5]*big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _Factory.Contract.CreatePerpetual(&_Factory.TransactOpts, oracle, coreParams, riskParams, minRiskParamValues, maxRiskParamValues, nonce)
}

// CreatePerpetual is a paid mutator transaction binding the contract method 0x53a1c053.
//
// Solidity: function createPerpetual(address oracle, int256[7] coreParams, int256[5] riskParams, int256[5] minRiskParamValues, int256[5] maxRiskParamValues, uint256 nonce) returns(address)
func (_Factory *FactoryTransactorSession) CreatePerpetual(oracle common.Address, coreParams [7]*big.Int, riskParams [5]*big.Int, minRiskParamValues [5]*big.Int, maxRiskParamValues [5]*big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _Factory.Contract.CreatePerpetual(&_Factory.TransactOpts, oracle, coreParams, riskParams, minRiskParamValues, maxRiskParamValues, nonce)
}

// DeactiveProxy is a paid mutator transaction binding the contract method 0x5b110896.
//
// Solidity: function deactiveProxy(address trader) returns(bool)
func (_Factory *FactoryTransactor) DeactiveProxy(opts *bind.TransactOpts, trader common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "deactiveProxy", trader)
}

// DeactiveProxy is a paid mutator transaction binding the contract method 0x5b110896.
//
// Solidity: function deactiveProxy(address trader) returns(bool)
func (_Factory *FactorySession) DeactiveProxy(trader common.Address) (*types.Transaction, error) {
	return _Factory.Contract.DeactiveProxy(&_Factory.TransactOpts, trader)
}

// DeactiveProxy is a paid mutator transaction binding the contract method 0x5b110896.
//
// Solidity: function deactiveProxy(address trader) returns(bool)
func (_Factory *FactoryTransactorSession) DeactiveProxy(trader common.Address) (*types.Transaction, error) {
	return _Factory.Contract.DeactiveProxy(&_Factory.TransactOpts, trader)
}

// FactoryAddVersionIterator is returned from FilterAddVersion and is used to iterate over the raw logs and unpacked data for AddVersion events raised by the Factory contract.
type FactoryAddVersionIterator struct {
	Event *FactoryAddVersion // Event containing the contract specifics and raw log

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
func (it *FactoryAddVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryAddVersion)
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
		it.Event = new(FactoryAddVersion)
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
func (it *FactoryAddVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryAddVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryAddVersion represents a AddVersion event raised by the Factory contract.
type FactoryAddVersion struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddVersion is a free log retrieval operation binding the contract event 0x51ea2037319ede313dbdbd0b178b14cf5fb14befaeaab5e24ed1d72f70ffc567.
//
// Solidity: event AddVersion(address implementation)
func (_Factory *FactoryFilterer) FilterAddVersion(opts *bind.FilterOpts) (*FactoryAddVersionIterator, error) {

	logs, sub, err := _Factory.contract.FilterLogs(opts, "AddVersion")
	if err != nil {
		return nil, err
	}
	return &FactoryAddVersionIterator{contract: _Factory.contract, event: "AddVersion", logs: logs, sub: sub}, nil
}

// WatchAddVersion is a free log subscription operation binding the contract event 0x51ea2037319ede313dbdbd0b178b14cf5fb14befaeaab5e24ed1d72f70ffc567.
//
// Solidity: event AddVersion(address implementation)
func (_Factory *FactoryFilterer) WatchAddVersion(opts *bind.WatchOpts, sink chan<- *FactoryAddVersion) (event.Subscription, error) {

	logs, sub, err := _Factory.contract.WatchLogs(opts, "AddVersion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryAddVersion)
				if err := _Factory.contract.UnpackLog(event, "AddVersion", log); err != nil {
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

// ParseAddVersion is a log parse operation binding the contract event 0x51ea2037319ede313dbdbd0b178b14cf5fb14befaeaab5e24ed1d72f70ffc567.
//
// Solidity: event AddVersion(address implementation)
func (_Factory *FactoryFilterer) ParseAddVersion(log types.Log) (*FactoryAddVersion, error) {
	event := new(FactoryAddVersion)
	if err := _Factory.contract.UnpackLog(event, "AddVersion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FactoryCreatePerpetualIterator is returned from FilterCreatePerpetual and is used to iterate over the raw logs and unpacked data for CreatePerpetual events raised by the Factory contract.
type FactoryCreatePerpetualIterator struct {
	Event *FactoryCreatePerpetual // Event containing the contract specifics and raw log

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
func (it *FactoryCreatePerpetualIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryCreatePerpetual)
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
		it.Event = new(FactoryCreatePerpetual)
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
func (it *FactoryCreatePerpetualIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryCreatePerpetualIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryCreatePerpetual represents a CreatePerpetual event raised by the Factory contract.
type FactoryCreatePerpetual struct {
	Perpetual  common.Address
	Governor   common.Address
	ShareToken common.Address
	Operator   common.Address
	Collateral common.Address
	Oracle     common.Address
	CoreParams [7]*big.Int
	RiskParams [5]*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreatePerpetual is a free log retrieval operation binding the contract event 0x00fed7a4317354cc24fd703396ae9c240add67e3eb2ca2ebe1400cf9218e0fdd.
//
// Solidity: event CreatePerpetual(address perpetual, address governor, address shareToken, address operator, address collateral, address oracle, int256[7] coreParams, int256[5] riskParams)
func (_Factory *FactoryFilterer) FilterCreatePerpetual(opts *bind.FilterOpts) (*FactoryCreatePerpetualIterator, error) {

	logs, sub, err := _Factory.contract.FilterLogs(opts, "CreatePerpetual")
	if err != nil {
		return nil, err
	}
	return &FactoryCreatePerpetualIterator{contract: _Factory.contract, event: "CreatePerpetual", logs: logs, sub: sub}, nil
}

// WatchCreatePerpetual is a free log subscription operation binding the contract event 0x00fed7a4317354cc24fd703396ae9c240add67e3eb2ca2ebe1400cf9218e0fdd.
//
// Solidity: event CreatePerpetual(address perpetual, address governor, address shareToken, address operator, address collateral, address oracle, int256[7] coreParams, int256[5] riskParams)
func (_Factory *FactoryFilterer) WatchCreatePerpetual(opts *bind.WatchOpts, sink chan<- *FactoryCreatePerpetual) (event.Subscription, error) {

	logs, sub, err := _Factory.contract.WatchLogs(opts, "CreatePerpetual")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryCreatePerpetual)
				if err := _Factory.contract.UnpackLog(event, "CreatePerpetual", log); err != nil {
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

// ParseCreatePerpetual is a log parse operation binding the contract event 0x00fed7a4317354cc24fd703396ae9c240add67e3eb2ca2ebe1400cf9218e0fdd.
//
// Solidity: event CreatePerpetual(address perpetual, address governor, address shareToken, address operator, address collateral, address oracle, int256[7] coreParams, int256[5] riskParams)
func (_Factory *FactoryFilterer) ParseCreatePerpetual(log types.Log) (*FactoryCreatePerpetual, error) {
	event := new(FactoryCreatePerpetual)
	if err := _Factory.contract.UnpackLog(event, "CreatePerpetual", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FactoryRevokeVersionIterator is returned from FilterRevokeVersion and is used to iterate over the raw logs and unpacked data for RevokeVersion events raised by the Factory contract.
type FactoryRevokeVersionIterator struct {
	Event *FactoryRevokeVersion // Event containing the contract specifics and raw log

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
func (it *FactoryRevokeVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryRevokeVersion)
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
		it.Event = new(FactoryRevokeVersion)
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
func (it *FactoryRevokeVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryRevokeVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryRevokeVersion represents a RevokeVersion event raised by the Factory contract.
type FactoryRevokeVersion struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRevokeVersion is a free log retrieval operation binding the contract event 0x3833a162b1aa25175fd0a1a978bb44a5af7cf069da101d07cae98234ec66bd31.
//
// Solidity: event RevokeVersion(address implementation)
func (_Factory *FactoryFilterer) FilterRevokeVersion(opts *bind.FilterOpts) (*FactoryRevokeVersionIterator, error) {

	logs, sub, err := _Factory.contract.FilterLogs(opts, "RevokeVersion")
	if err != nil {
		return nil, err
	}
	return &FactoryRevokeVersionIterator{contract: _Factory.contract, event: "RevokeVersion", logs: logs, sub: sub}, nil
}

// WatchRevokeVersion is a free log subscription operation binding the contract event 0x3833a162b1aa25175fd0a1a978bb44a5af7cf069da101d07cae98234ec66bd31.
//
// Solidity: event RevokeVersion(address implementation)
func (_Factory *FactoryFilterer) WatchRevokeVersion(opts *bind.WatchOpts, sink chan<- *FactoryRevokeVersion) (event.Subscription, error) {

	logs, sub, err := _Factory.contract.WatchLogs(opts, "RevokeVersion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryRevokeVersion)
				if err := _Factory.contract.UnpackLog(event, "RevokeVersion", log); err != nil {
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

// ParseRevokeVersion is a log parse operation binding the contract event 0x3833a162b1aa25175fd0a1a978bb44a5af7cf069da101d07cae98234ec66bd31.
//
// Solidity: event RevokeVersion(address implementation)
func (_Factory *FactoryFilterer) ParseRevokeVersion(log types.Log) (*FactoryRevokeVersion, error) {
	event := new(FactoryRevokeVersion)
	if err := _Factory.contract.UnpackLog(event, "RevokeVersion", log); err != nil {
		return nil, err
	}
	return event, nil
}
