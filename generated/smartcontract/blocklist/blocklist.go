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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newValue\",\"type\":\"string\"}],\"name\":\"addHostName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"authorizeUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorizedUsers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHostList\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hostlist\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"}],\"name\":\"unAuthorizeUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
