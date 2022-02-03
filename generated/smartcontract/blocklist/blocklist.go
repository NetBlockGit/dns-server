// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blocklist

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BlocklistMetaData contains all meta data concerning the Blocklist contract.
var BlocklistMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"}],\"name\":\"HostNameAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"HostNameDeleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newValue\",\"type\":\"string\"}],\"name\":\"addHostName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"authorizeUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorizedUsers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHostList\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hostlist\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeHostList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"unAuthorizeUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BlocklistABI is the input ABI used to generate the binding from.
// Deprecated: Use BlocklistMetaData.ABI instead.
var BlocklistABI = BlocklistMetaData.ABI

// Blocklist is an auto generated Go binding around an Ethereum contract.
type Blocklist struct {
	BlocklistCaller     // Read-only binding to the contract
	BlocklistTransactor // Write-only binding to the contract
	BlocklistFilterer   // Log filterer for contract events
}

// BlocklistCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlocklistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlocklistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlocklistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlocklistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlocklistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlocklistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlocklistSession struct {
	Contract     *Blocklist        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlocklistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlocklistCallerSession struct {
	Contract *BlocklistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BlocklistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlocklistTransactorSession struct {
	Contract     *BlocklistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BlocklistRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlocklistRaw struct {
	Contract *Blocklist // Generic contract binding to access the raw methods on
}

// BlocklistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlocklistCallerRaw struct {
	Contract *BlocklistCaller // Generic read-only contract binding to access the raw methods on
}

// BlocklistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlocklistTransactorRaw struct {
	Contract *BlocklistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlocklist creates a new instance of Blocklist, bound to a specific deployed contract.
func NewBlocklist(address common.Address, backend bind.ContractBackend) (*Blocklist, error) {
	contract, err := bindBlocklist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blocklist{BlocklistCaller: BlocklistCaller{contract: contract}, BlocklistTransactor: BlocklistTransactor{contract: contract}, BlocklistFilterer: BlocklistFilterer{contract: contract}}, nil
}

// NewBlocklistCaller creates a new read-only instance of Blocklist, bound to a specific deployed contract.
func NewBlocklistCaller(address common.Address, caller bind.ContractCaller) (*BlocklistCaller, error) {
	contract, err := bindBlocklist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlocklistCaller{contract: contract}, nil
}

// NewBlocklistTransactor creates a new write-only instance of Blocklist, bound to a specific deployed contract.
func NewBlocklistTransactor(address common.Address, transactor bind.ContractTransactor) (*BlocklistTransactor, error) {
	contract, err := bindBlocklist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlocklistTransactor{contract: contract}, nil
}

// NewBlocklistFilterer creates a new log filterer instance of Blocklist, bound to a specific deployed contract.
func NewBlocklistFilterer(address common.Address, filterer bind.ContractFilterer) (*BlocklistFilterer, error) {
	contract, err := bindBlocklist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlocklistFilterer{contract: contract}, nil
}

// bindBlocklist binds a generic wrapper to an already deployed contract.
func bindBlocklist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlocklistABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blocklist *BlocklistRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blocklist.Contract.BlocklistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blocklist *BlocklistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blocklist.Contract.BlocklistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blocklist *BlocklistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blocklist.Contract.BlocklistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blocklist *BlocklistCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blocklist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blocklist *BlocklistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blocklist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blocklist *BlocklistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blocklist.Contract.contract.Transact(opts, method, params...)
}

// AuthorizedUsers is a free data retrieval call binding the contract method 0x1828983a.
//
// Solidity: function authorizedUsers(address ) view returns(bool)
func (_Blocklist *BlocklistCaller) AuthorizedUsers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Blocklist.contract.Call(opts, &out, "authorizedUsers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthorizedUsers is a free data retrieval call binding the contract method 0x1828983a.
//
// Solidity: function authorizedUsers(address ) view returns(bool)
func (_Blocklist *BlocklistSession) AuthorizedUsers(arg0 common.Address) (bool, error) {
	return _Blocklist.Contract.AuthorizedUsers(&_Blocklist.CallOpts, arg0)
}

// AuthorizedUsers is a free data retrieval call binding the contract method 0x1828983a.
//
// Solidity: function authorizedUsers(address ) view returns(bool)
func (_Blocklist *BlocklistCallerSession) AuthorizedUsers(arg0 common.Address) (bool, error) {
	return _Blocklist.Contract.AuthorizedUsers(&_Blocklist.CallOpts, arg0)
}

// GetHostList is a free data retrieval call binding the contract method 0x8e076cf0.
//
// Solidity: function getHostList() view returns(string[])
func (_Blocklist *BlocklistCaller) GetHostList(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Blocklist.contract.Call(opts, &out, "getHostList")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetHostList is a free data retrieval call binding the contract method 0x8e076cf0.
//
// Solidity: function getHostList() view returns(string[])
func (_Blocklist *BlocklistSession) GetHostList() ([]string, error) {
	return _Blocklist.Contract.GetHostList(&_Blocklist.CallOpts)
}

// GetHostList is a free data retrieval call binding the contract method 0x8e076cf0.
//
// Solidity: function getHostList() view returns(string[])
func (_Blocklist *BlocklistCallerSession) GetHostList() ([]string, error) {
	return _Blocklist.Contract.GetHostList(&_Blocklist.CallOpts)
}

// Hostlist is a free data retrieval call binding the contract method 0xa67d1448.
//
// Solidity: function hostlist(uint256 ) view returns(string)
func (_Blocklist *BlocklistCaller) Hostlist(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Blocklist.contract.Call(opts, &out, "hostlist", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Hostlist is a free data retrieval call binding the contract method 0xa67d1448.
//
// Solidity: function hostlist(uint256 ) view returns(string)
func (_Blocklist *BlocklistSession) Hostlist(arg0 *big.Int) (string, error) {
	return _Blocklist.Contract.Hostlist(&_Blocklist.CallOpts, arg0)
}

// Hostlist is a free data retrieval call binding the contract method 0xa67d1448.
//
// Solidity: function hostlist(uint256 ) view returns(string)
func (_Blocklist *BlocklistCallerSession) Hostlist(arg0 *big.Int) (string, error) {
	return _Blocklist.Contract.Hostlist(&_Blocklist.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blocklist *BlocklistCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blocklist.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blocklist *BlocklistSession) Owner() (common.Address, error) {
	return _Blocklist.Contract.Owner(&_Blocklist.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blocklist *BlocklistCallerSession) Owner() (common.Address, error) {
	return _Blocklist.Contract.Owner(&_Blocklist.CallOpts)
}

// AddHostName is a paid mutator transaction binding the contract method 0x09d11e30.
//
// Solidity: function addHostName(string newValue) returns()
func (_Blocklist *BlocklistTransactor) AddHostName(opts *bind.TransactOpts, newValue string) (*types.Transaction, error) {
	return _Blocklist.contract.Transact(opts, "addHostName", newValue)
}

// AddHostName is a paid mutator transaction binding the contract method 0x09d11e30.
//
// Solidity: function addHostName(string newValue) returns()
func (_Blocklist *BlocklistSession) AddHostName(newValue string) (*types.Transaction, error) {
	return _Blocklist.Contract.AddHostName(&_Blocklist.TransactOpts, newValue)
}

// AddHostName is a paid mutator transaction binding the contract method 0x09d11e30.
//
// Solidity: function addHostName(string newValue) returns()
func (_Blocklist *BlocklistTransactorSession) AddHostName(newValue string) (*types.Transaction, error) {
	return _Blocklist.Contract.AddHostName(&_Blocklist.TransactOpts, newValue)
}

// AuthorizeUser is a paid mutator transaction binding the contract method 0x67c2a360.
//
// Solidity: function authorizeUser(address userAddr) returns()
func (_Blocklist *BlocklistTransactor) AuthorizeUser(opts *bind.TransactOpts, userAddr common.Address) (*types.Transaction, error) {
	return _Blocklist.contract.Transact(opts, "authorizeUser", userAddr)
}

// AuthorizeUser is a paid mutator transaction binding the contract method 0x67c2a360.
//
// Solidity: function authorizeUser(address userAddr) returns()
func (_Blocklist *BlocklistSession) AuthorizeUser(userAddr common.Address) (*types.Transaction, error) {
	return _Blocklist.Contract.AuthorizeUser(&_Blocklist.TransactOpts, userAddr)
}

// AuthorizeUser is a paid mutator transaction binding the contract method 0x67c2a360.
//
// Solidity: function authorizeUser(address userAddr) returns()
func (_Blocklist *BlocklistTransactorSession) AuthorizeUser(userAddr common.Address) (*types.Transaction, error) {
	return _Blocklist.Contract.AuthorizeUser(&_Blocklist.TransactOpts, userAddr)
}

// RemoveHostList is a paid mutator transaction binding the contract method 0x9b567cb6.
//
// Solidity: function removeHostList(uint256 index) returns()
func (_Blocklist *BlocklistTransactor) RemoveHostList(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _Blocklist.contract.Transact(opts, "removeHostList", index)
}

// RemoveHostList is a paid mutator transaction binding the contract method 0x9b567cb6.
//
// Solidity: function removeHostList(uint256 index) returns()
func (_Blocklist *BlocklistSession) RemoveHostList(index *big.Int) (*types.Transaction, error) {
	return _Blocklist.Contract.RemoveHostList(&_Blocklist.TransactOpts, index)
}

// RemoveHostList is a paid mutator transaction binding the contract method 0x9b567cb6.
//
// Solidity: function removeHostList(uint256 index) returns()
func (_Blocklist *BlocklistTransactorSession) RemoveHostList(index *big.Int) (*types.Transaction, error) {
	return _Blocklist.Contract.RemoveHostList(&_Blocklist.TransactOpts, index)
}

// UnAuthorizeUser is a paid mutator transaction binding the contract method 0x09151c7f.
//
// Solidity: function unAuthorizeUser(address userAddr) returns()
func (_Blocklist *BlocklistTransactor) UnAuthorizeUser(opts *bind.TransactOpts, userAddr common.Address) (*types.Transaction, error) {
	return _Blocklist.contract.Transact(opts, "unAuthorizeUser", userAddr)
}

// UnAuthorizeUser is a paid mutator transaction binding the contract method 0x09151c7f.
//
// Solidity: function unAuthorizeUser(address userAddr) returns()
func (_Blocklist *BlocklistSession) UnAuthorizeUser(userAddr common.Address) (*types.Transaction, error) {
	return _Blocklist.Contract.UnAuthorizeUser(&_Blocklist.TransactOpts, userAddr)
}

// UnAuthorizeUser is a paid mutator transaction binding the contract method 0x09151c7f.
//
// Solidity: function unAuthorizeUser(address userAddr) returns()
func (_Blocklist *BlocklistTransactorSession) UnAuthorizeUser(userAddr common.Address) (*types.Transaction, error) {
	return _Blocklist.Contract.UnAuthorizeUser(&_Blocklist.TransactOpts, userAddr)
}

// BlocklistHostNameAddedIterator is returned from FilterHostNameAdded and is used to iterate over the raw logs and unpacked data for HostNameAdded events raised by the Blocklist contract.
type BlocklistHostNameAddedIterator struct {
	Event *BlocklistHostNameAdded // Event containing the contract specifics and raw log

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
func (it *BlocklistHostNameAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlocklistHostNameAdded)
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
		it.Event = new(BlocklistHostNameAdded)
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
func (it *BlocklistHostNameAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlocklistHostNameAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlocklistHostNameAdded represents a HostNameAdded event raised by the Blocklist contract.
type BlocklistHostNameAdded struct {
	Hostname string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterHostNameAdded is a free log retrieval operation binding the contract event 0x29205eb0c7ad84d9cfe64c7e47012225e8130a13a12e7fbf079f6ffc1b813a98.
//
// Solidity: event HostNameAdded(string hostname)
func (_Blocklist *BlocklistFilterer) FilterHostNameAdded(opts *bind.FilterOpts) (*BlocklistHostNameAddedIterator, error) {

	logs, sub, err := _Blocklist.contract.FilterLogs(opts, "HostNameAdded")
	if err != nil {
		return nil, err
	}
	return &BlocklistHostNameAddedIterator{contract: _Blocklist.contract, event: "HostNameAdded", logs: logs, sub: sub}, nil
}

// WatchHostNameAdded is a free log subscription operation binding the contract event 0x29205eb0c7ad84d9cfe64c7e47012225e8130a13a12e7fbf079f6ffc1b813a98.
//
// Solidity: event HostNameAdded(string hostname)
func (_Blocklist *BlocklistFilterer) WatchHostNameAdded(opts *bind.WatchOpts, sink chan<- *BlocklistHostNameAdded) (event.Subscription, error) {

	logs, sub, err := _Blocklist.contract.WatchLogs(opts, "HostNameAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlocklistHostNameAdded)
				if err := _Blocklist.contract.UnpackLog(event, "HostNameAdded", log); err != nil {
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

// ParseHostNameAdded is a log parse operation binding the contract event 0x29205eb0c7ad84d9cfe64c7e47012225e8130a13a12e7fbf079f6ffc1b813a98.
//
// Solidity: event HostNameAdded(string hostname)
func (_Blocklist *BlocklistFilterer) ParseHostNameAdded(log types.Log) (*BlocklistHostNameAdded, error) {
	event := new(BlocklistHostNameAdded)
	if err := _Blocklist.contract.UnpackLog(event, "HostNameAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlocklistHostNameDeletedIterator is returned from FilterHostNameDeleted and is used to iterate over the raw logs and unpacked data for HostNameDeleted events raised by the Blocklist contract.
type BlocklistHostNameDeletedIterator struct {
	Event *BlocklistHostNameDeleted // Event containing the contract specifics and raw log

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
func (it *BlocklistHostNameDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlocklistHostNameDeleted)
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
		it.Event = new(BlocklistHostNameDeleted)
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
func (it *BlocklistHostNameDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlocklistHostNameDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlocklistHostNameDeleted represents a HostNameDeleted event raised by the Blocklist contract.
type BlocklistHostNameDeleted struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterHostNameDeleted is a free log retrieval operation binding the contract event 0x9838c313878774c7095f2ec13276b88235ab0bc2b3ac66435711b6407134fedb.
//
// Solidity: event HostNameDeleted(uint256 index)
func (_Blocklist *BlocklistFilterer) FilterHostNameDeleted(opts *bind.FilterOpts) (*BlocklistHostNameDeletedIterator, error) {

	logs, sub, err := _Blocklist.contract.FilterLogs(opts, "HostNameDeleted")
	if err != nil {
		return nil, err
	}
	return &BlocklistHostNameDeletedIterator{contract: _Blocklist.contract, event: "HostNameDeleted", logs: logs, sub: sub}, nil
}

// WatchHostNameDeleted is a free log subscription operation binding the contract event 0x9838c313878774c7095f2ec13276b88235ab0bc2b3ac66435711b6407134fedb.
//
// Solidity: event HostNameDeleted(uint256 index)
func (_Blocklist *BlocklistFilterer) WatchHostNameDeleted(opts *bind.WatchOpts, sink chan<- *BlocklistHostNameDeleted) (event.Subscription, error) {

	logs, sub, err := _Blocklist.contract.WatchLogs(opts, "HostNameDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlocklistHostNameDeleted)
				if err := _Blocklist.contract.UnpackLog(event, "HostNameDeleted", log); err != nil {
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

// ParseHostNameDeleted is a log parse operation binding the contract event 0x9838c313878774c7095f2ec13276b88235ab0bc2b3ac66435711b6407134fedb.
//
// Solidity: event HostNameDeleted(uint256 index)
func (_Blocklist *BlocklistFilterer) ParseHostNameDeleted(log types.Log) (*BlocklistHostNameDeleted, error) {
	event := new(BlocklistHostNameDeleted)
	if err := _Blocklist.contract.UnpackLog(event, "HostNameDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
