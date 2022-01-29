// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrapper

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

// WrapperMetaData contains all meta data concerning the Wrapper contract.
var WrapperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getBytesByHash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getBytesByName\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getHashByName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"schemaName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"schemaBody\",\"type\":\"bytes\"}],\"name\":\"save\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WrapperABI is the input ABI used to generate the binding from.
// Deprecated: Use WrapperMetaData.ABI instead.
var WrapperABI = WrapperMetaData.ABI

// Wrapper is an auto generated Go binding around an Ethereum contract.
type Wrapper struct {
	WrapperCaller     // Read-only binding to the contract
	WrapperTransactor // Write-only binding to the contract
	WrapperFilterer   // Log filterer for contract events
}

// WrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type WrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WrapperSession struct {
	Contract     *Wrapper          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WrapperCallerSession struct {
	Contract *WrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// WrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WrapperTransactorSession struct {
	Contract     *WrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// WrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type WrapperRaw struct {
	Contract *Wrapper // Generic contract binding to access the raw methods on
}

// WrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WrapperCallerRaw struct {
	Contract *WrapperCaller // Generic read-only contract binding to access the raw methods on
}

// WrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WrapperTransactorRaw struct {
	Contract *WrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWrapper creates a new instance of Wrapper, bound to a specific deployed contract.
func NewWrapper(address common.Address, backend bind.ContractBackend) (*Wrapper, error) {
	contract, err := bindWrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wrapper{WrapperCaller: WrapperCaller{contract: contract}, WrapperTransactor: WrapperTransactor{contract: contract}, WrapperFilterer: WrapperFilterer{contract: contract}}, nil
}

// NewWrapperCaller creates a new read-only instance of Wrapper, bound to a specific deployed contract.
func NewWrapperCaller(address common.Address, caller bind.ContractCaller) (*WrapperCaller, error) {
	contract, err := bindWrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WrapperCaller{contract: contract}, nil
}

// NewWrapperTransactor creates a new write-only instance of Wrapper, bound to a specific deployed contract.
func NewWrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*WrapperTransactor, error) {
	contract, err := bindWrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WrapperTransactor{contract: contract}, nil
}

// NewWrapperFilterer creates a new log filterer instance of Wrapper, bound to a specific deployed contract.
func NewWrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*WrapperFilterer, error) {
	contract, err := bindWrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WrapperFilterer{contract: contract}, nil
}

// bindWrapper binds a generic wrapper to an already deployed contract.
func bindWrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WrapperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wrapper *WrapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wrapper.Contract.WrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wrapper *WrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wrapper.Contract.WrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wrapper *WrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wrapper.Contract.WrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wrapper *WrapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wrapper *WrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wrapper *WrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wrapper.Contract.contract.Transact(opts, method, params...)
}

// GetBytesByHash is a free data retrieval call binding the contract method 0x78e4b929.
//
// Solidity: function getBytesByHash(bytes32 hash) view returns(bytes)
func (_Wrapper *WrapperCaller) GetBytesByHash(opts *bind.CallOpts, hash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Wrapper.contract.Call(opts, &out, "getBytesByHash", hash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetBytesByHash is a free data retrieval call binding the contract method 0x78e4b929.
//
// Solidity: function getBytesByHash(bytes32 hash) view returns(bytes)
func (_Wrapper *WrapperSession) GetBytesByHash(hash [32]byte) ([]byte, error) {
	return _Wrapper.Contract.GetBytesByHash(&_Wrapper.CallOpts, hash)
}

// GetBytesByHash is a free data retrieval call binding the contract method 0x78e4b929.
//
// Solidity: function getBytesByHash(bytes32 hash) view returns(bytes)
func (_Wrapper *WrapperCallerSession) GetBytesByHash(hash [32]byte) ([]byte, error) {
	return _Wrapper.Contract.GetBytesByHash(&_Wrapper.CallOpts, hash)
}

// GetBytesByName is a free data retrieval call binding the contract method 0x99f41794.
//
// Solidity: function getBytesByName(string name) view returns(bytes)
func (_Wrapper *WrapperCaller) GetBytesByName(opts *bind.CallOpts, name string) ([]byte, error) {
	var out []interface{}
	err := _Wrapper.contract.Call(opts, &out, "getBytesByName", name)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetBytesByName is a free data retrieval call binding the contract method 0x99f41794.
//
// Solidity: function getBytesByName(string name) view returns(bytes)
func (_Wrapper *WrapperSession) GetBytesByName(name string) ([]byte, error) {
	return _Wrapper.Contract.GetBytesByName(&_Wrapper.CallOpts, name)
}

// GetBytesByName is a free data retrieval call binding the contract method 0x99f41794.
//
// Solidity: function getBytesByName(string name) view returns(bytes)
func (_Wrapper *WrapperCallerSession) GetBytesByName(name string) ([]byte, error) {
	return _Wrapper.Contract.GetBytesByName(&_Wrapper.CallOpts, name)
}

// GetHashByName is a free data retrieval call binding the contract method 0x79c7db5a.
//
// Solidity: function getHashByName(string name) view returns(bytes32)
func (_Wrapper *WrapperCaller) GetHashByName(opts *bind.CallOpts, name string) ([32]byte, error) {
	var out []interface{}
	err := _Wrapper.contract.Call(opts, &out, "getHashByName", name)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHashByName is a free data retrieval call binding the contract method 0x79c7db5a.
//
// Solidity: function getHashByName(string name) view returns(bytes32)
func (_Wrapper *WrapperSession) GetHashByName(name string) ([32]byte, error) {
	return _Wrapper.Contract.GetHashByName(&_Wrapper.CallOpts, name)
}

// GetHashByName is a free data retrieval call binding the contract method 0x79c7db5a.
//
// Solidity: function getHashByName(string name) view returns(bytes32)
func (_Wrapper *WrapperCallerSession) GetHashByName(name string) ([32]byte, error) {
	return _Wrapper.Contract.GetHashByName(&_Wrapper.CallOpts, name)
}

// Save is a paid mutator transaction binding the contract method 0xcc9af980.
//
// Solidity: function save(string schemaName, bytes schemaBody) returns()
func (_Wrapper *WrapperTransactor) Save(opts *bind.TransactOpts, schemaName string, schemaBody []byte) (*types.Transaction, error) {
	return _Wrapper.contract.Transact(opts, "save", schemaName, schemaBody)
}

// Save is a paid mutator transaction binding the contract method 0xcc9af980.
//
// Solidity: function save(string schemaName, bytes schemaBody) returns()
func (_Wrapper *WrapperSession) Save(schemaName string, schemaBody []byte) (*types.Transaction, error) {
	return _Wrapper.Contract.Save(&_Wrapper.TransactOpts, schemaName, schemaBody)
}

// Save is a paid mutator transaction binding the contract method 0xcc9af980.
//
// Solidity: function save(string schemaName, bytes schemaBody) returns()
func (_Wrapper *WrapperTransactorSession) Save(schemaName string, schemaBody []byte) (*types.Transaction, error) {
	return _Wrapper.Contract.Save(&_Wrapper.TransactOpts, schemaName, schemaBody)
}
