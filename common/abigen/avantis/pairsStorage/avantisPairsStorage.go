// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package avantis_pairs_abis

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
	_ = abi.ConvertType
)

// IPairStorageBackupFeed is an auto generated low-level Go binding around an user-defined struct.
type IPairStorageBackupFeed struct {
	MaxDeviationP *big.Int
	FeedId        common.Address
}

// IPairStorageFeed is an auto generated low-level Go binding around an user-defined struct.
type IPairStorageFeed struct {
	MaxDeviationP *big.Int
	FeedId        [32]byte
}

// AvantisPairsAbisMetaData contains all meta data concerning the AvantisPairsAbis contract.
var AvantisPairsAbisMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"blockOILimit\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"groupMaxOI\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"groupOI\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"guaranteedSlEnabled\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isUSDCAligned\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lossProtectionMultiplier\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tier\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxWalletOI\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairBackupFeed\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.BackupFeed\",\"components\":[{\"name\":\"maxDeviationP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feedId\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairCloseFeeP\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairFeed\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.Feed\",\"components\":[{\"name\":\"maxDeviationP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feedId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairGroupIndex\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairJob\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pairLimitOrderFeeP\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairMaxLeverage\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairMaxOI\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairMinLevPosUSDC\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairMinLeverage\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairOpenFeeP\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairPriceImpactMultiplier\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairSkewImpactMultiplier\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairSpreadP\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairsCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateGroupOI\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BlockOILimitsSet\",\"inputs\":[{\"name\":\"pairIndex\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"limits\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeAdded\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeUpdated\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GroupAdded\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GroupUpdated\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LossProtectionAdded\",\"inputs\":[{\"name\":\"pairIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"tier\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"multiplier\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderLimitsSet\",\"inputs\":[{\"name\":\"pairIndex\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"limits\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PairAdded\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"to\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PairUpdated\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SkewFeeAdded\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SkewFeeUpdated\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// AvantisPairsAbisABI is the input ABI used to generate the binding from.
// Deprecated: Use AvantisPairsAbisMetaData.ABI instead.
var AvantisPairsAbisABI = AvantisPairsAbisMetaData.ABI

// AvantisPairsAbis is an auto generated Go binding around an Ethereum contract.
type AvantisPairsAbis struct {
	AvantisPairsAbisCaller     // Read-only binding to the contract
	AvantisPairsAbisTransactor // Write-only binding to the contract
	AvantisPairsAbisFilterer   // Log filterer for contract events
}

// AvantisPairsAbisCaller is an auto generated read-only Go binding around an Ethereum contract.
type AvantisPairsAbisCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvantisPairsAbisTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AvantisPairsAbisTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvantisPairsAbisFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AvantisPairsAbisFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvantisPairsAbisSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AvantisPairsAbisSession struct {
	Contract     *AvantisPairsAbis // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AvantisPairsAbisCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AvantisPairsAbisCallerSession struct {
	Contract *AvantisPairsAbisCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// AvantisPairsAbisTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AvantisPairsAbisTransactorSession struct {
	Contract     *AvantisPairsAbisTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AvantisPairsAbisRaw is an auto generated low-level Go binding around an Ethereum contract.
type AvantisPairsAbisRaw struct {
	Contract *AvantisPairsAbis // Generic contract binding to access the raw methods on
}

// AvantisPairsAbisCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AvantisPairsAbisCallerRaw struct {
	Contract *AvantisPairsAbisCaller // Generic read-only contract binding to access the raw methods on
}

// AvantisPairsAbisTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AvantisPairsAbisTransactorRaw struct {
	Contract *AvantisPairsAbisTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAvantisPairsAbis creates a new instance of AvantisPairsAbis, bound to a specific deployed contract.
func NewAvantisPairsAbis(address common.Address, backend bind.ContractBackend) (*AvantisPairsAbis, error) {
	contract, err := bindAvantisPairsAbis(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AvantisPairsAbis{AvantisPairsAbisCaller: AvantisPairsAbisCaller{contract: contract}, AvantisPairsAbisTransactor: AvantisPairsAbisTransactor{contract: contract}, AvantisPairsAbisFilterer: AvantisPairsAbisFilterer{contract: contract}}, nil
}

// NewAvantisPairsAbisCaller creates a new read-only instance of AvantisPairsAbis, bound to a specific deployed contract.
func NewAvantisPairsAbisCaller(address common.Address, caller bind.ContractCaller) (*AvantisPairsAbisCaller, error) {
	contract, err := bindAvantisPairsAbis(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AvantisPairsAbisCaller{contract: contract}, nil
}

// NewAvantisPairsAbisTransactor creates a new write-only instance of AvantisPairsAbis, bound to a specific deployed contract.
func NewAvantisPairsAbisTransactor(address common.Address, transactor bind.ContractTransactor) (*AvantisPairsAbisTransactor, error) {
	contract, err := bindAvantisPairsAbis(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AvantisPairsAbisTransactor{contract: contract}, nil
}

// NewAvantisPairsAbisFilterer creates a new log filterer instance of AvantisPairsAbis, bound to a specific deployed contract.
func NewAvantisPairsAbisFilterer(address common.Address, filterer bind.ContractFilterer) (*AvantisPairsAbisFilterer, error) {
	contract, err := bindAvantisPairsAbis(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AvantisPairsAbisFilterer{contract: contract}, nil
}

// bindAvantisPairsAbis binds a generic wrapper to an already deployed contract.
func bindAvantisPairsAbis(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AvantisPairsAbisMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AvantisPairsAbis *AvantisPairsAbisRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AvantisPairsAbis.Contract.AvantisPairsAbisCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AvantisPairsAbis *AvantisPairsAbisRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AvantisPairsAbis.Contract.AvantisPairsAbisTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AvantisPairsAbis *AvantisPairsAbisRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AvantisPairsAbis.Contract.AvantisPairsAbisTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AvantisPairsAbis *AvantisPairsAbisCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AvantisPairsAbis.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AvantisPairsAbis *AvantisPairsAbisTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AvantisPairsAbis.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AvantisPairsAbis *AvantisPairsAbisTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AvantisPairsAbis.Contract.contract.Transact(opts, method, params...)
}

// BlockOILimit is a free data retrieval call binding the contract method 0x9ae6b9e0.
//
// Solidity: function blockOILimit(uint256 _pairIndex) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCaller) BlockOILimit(opts *bind.CallOpts, _pairIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisPairsAbis.contract.Call(opts, &out, "blockOILimit", _pairIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockOILimit is a free data retrieval call binding the contract method 0x9ae6b9e0.
//
// Solidity: function blockOILimit(uint256 _pairIndex) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisSession) BlockOILimit(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.BlockOILimit(&_AvantisPairsAbis.CallOpts, _pairIndex)
}

// BlockOILimit is a free data retrieval call binding the contract method 0x9ae6b9e0.
//
// Solidity: function blockOILimit(uint256 _pairIndex) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCallerSession) BlockOILimit(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.BlockOILimit(&_AvantisPairsAbis.CallOpts, _pairIndex)
}

// GroupMaxOI is a free data retrieval call binding the contract method 0x8c9a0ea4.
//
// Solidity: function groupMaxOI(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCaller) GroupMaxOI(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisPairsAbis.contract.Call(opts, &out, "groupMaxOI", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GroupMaxOI is a free data retrieval call binding the contract method 0x8c9a0ea4.
//
// Solidity: function groupMaxOI(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisSession) GroupMaxOI(arg0 *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.GroupMaxOI(&_AvantisPairsAbis.CallOpts, arg0)
}

// GroupMaxOI is a free data retrieval call binding the contract method 0x8c9a0ea4.
//
// Solidity: function groupMaxOI(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCallerSession) GroupMaxOI(arg0 *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.GroupMaxOI(&_AvantisPairsAbis.CallOpts, arg0)
}

// GroupOI is a free data retrieval call binding the contract method 0x5e4f8f59.
//
// Solidity: function groupOI(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCaller) GroupOI(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisPairsAbis.contract.Call(opts, &out, "groupOI", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GroupOI is a free data retrieval call binding the contract method 0x5e4f8f59.
//
// Solidity: function groupOI(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisSession) GroupOI(arg0 *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.GroupOI(&_AvantisPairsAbis.CallOpts, arg0)
}

// GroupOI is a free data retrieval call binding the contract method 0x5e4f8f59.
//
// Solidity: function groupOI(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCallerSession) GroupOI(arg0 *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.GroupOI(&_AvantisPairsAbis.CallOpts, arg0)
}

// GuaranteedSlEnabled is a free data retrieval call binding the contract method 0x24abd3fb.
//
// Solidity: function guaranteedSlEnabled(uint256 ) view returns(bool)
func (_AvantisPairsAbis *AvantisPairsAbisCaller) GuaranteedSlEnabled(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _AvantisPairsAbis.contract.Call(opts, &out, "guaranteedSlEnabled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GuaranteedSlEnabled is a free data retrieval call binding the contract method 0x24abd3fb.
//
// Solidity: function guaranteedSlEnabled(uint256 ) view returns(bool)
func (_AvantisPairsAbis *AvantisPairsAbisSession) GuaranteedSlEnabled(arg0 *big.Int) (bool, error) {
	return _AvantisPairsAbis.Contract.GuaranteedSlEnabled(&_AvantisPairsAbis.CallOpts, arg0)
}

// GuaranteedSlEnabled is a free data retrieval call binding the contract method 0x24abd3fb.
//
// Solidity: function guaranteedSlEnabled(uint256 ) view returns(bool)
func (_AvantisPairsAbis *AvantisPairsAbisCallerSession) GuaranteedSlEnabled(arg0 *big.Int) (bool, error) {
	return _AvantisPairsAbis.Contract.GuaranteedSlEnabled(&_AvantisPairsAbis.CallOpts, arg0)
}

// IsUSDCAligned is a free data retrieval call binding the contract method 0x6dc37de3.
//
// Solidity: function isUSDCAligned(uint256 _pairIndex) view returns(bool)
func (_AvantisPairsAbis *AvantisPairsAbisCaller) IsUSDCAligned(opts *bind.CallOpts, _pairIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _AvantisPairsAbis.contract.Call(opts, &out, "isUSDCAligned", _pairIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUSDCAligned is a free data retrieval call binding the contract method 0x6dc37de3.
//
// Solidity: function isUSDCAligned(uint256 _pairIndex) view returns(bool)
func (_AvantisPairsAbis *AvantisPairsAbisSession) IsUSDCAligned(_pairIndex *big.Int) (bool, error) {
	return _AvantisPairsAbis.Contract.IsUSDCAligned(&_AvantisPairsAbis.CallOpts, _pairIndex)
}

// IsUSDCAligned is a free data retrieval call binding the contract method 0x6dc37de3.
//
// Solidity: function isUSDCAligned(uint256 _pairIndex) view returns(bool)
func (_AvantisPairsAbis *AvantisPairsAbisCallerSession) IsUSDCAligned(_pairIndex *big.Int) (bool, error) {
	return _AvantisPairsAbis.Contract.IsUSDCAligned(&_AvantisPairsAbis.CallOpts, _pairIndex)
}

// LossProtectionMultiplier is a free data retrieval call binding the contract method 0x2b1cbe9d.
//
// Solidity: function lossProtectionMultiplier(uint256 _pairIndex, uint256 _tier) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCaller) LossProtectionMultiplier(opts *bind.CallOpts, _pairIndex *big.Int, _tier *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisPairsAbis.contract.Call(opts, &out, "lossProtectionMultiplier", _pairIndex, _tier)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LossProtectionMultiplier is a free data retrieval call binding the contract method 0x2b1cbe9d.
//
// Solidity: function lossProtectionMultiplier(uint256 _pairIndex, uint256 _tier) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisSession) LossProtectionMultiplier(_pairIndex *big.Int, _tier *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.LossProtectionMultiplier(&_AvantisPairsAbis.CallOpts, _pairIndex, _tier)
}

// LossProtectionMultiplier is a free data retrieval call binding the contract method 0x2b1cbe9d.
//
// Solidity: function lossProtectionMultiplier(uint256 _pairIndex, uint256 _tier) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCallerSession) LossProtectionMultiplier(_pairIndex *big.Int, _tier *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.LossProtectionMultiplier(&_AvantisPairsAbis.CallOpts, _pairIndex, _tier)
}

// MaxWalletOI is a free data retrieval call binding the contract method 0x56e5536c.
//
// Solidity: function maxWalletOI(uint256 _pairIndex) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCaller) MaxWalletOI(opts *bind.CallOpts, _pairIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisPairsAbis.contract.Call(opts, &out, "maxWalletOI", _pairIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWalletOI is a free data retrieval call binding the contract method 0x56e5536c.
//
// Solidity: function maxWalletOI(uint256 _pairIndex) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisSession) MaxWalletOI(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.MaxWalletOI(&_AvantisPairsAbis.CallOpts, _pairIndex)
}

// MaxWalletOI is a free data retrieval call binding the contract method 0x56e5536c.
//
// Solidity: function maxWalletOI(uint256 _pairIndex) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCallerSession) MaxWalletOI(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.MaxWalletOI(&_AvantisPairsAbis.CallOpts, _pairIndex)
}

// PairBackupFeed is a free data retrieval call binding the contract method 0x685e905b.
//
// Solidity: function pairBackupFeed(uint256 ) view returns((uint256,address))
func (_AvantisPairsAbis *AvantisPairsAbisCaller) PairBackupFeed(opts *bind.CallOpts, arg0 *big.Int) (IPairStorageBackupFeed, error) {
	var out []interface{}
	err := _AvantisPairsAbis.contract.Call(opts, &out, "pairBackupFeed", arg0)

	if err != nil {
		return *new(IPairStorageBackupFeed), err
	}

	out0 := *abi.ConvertType(out[0], new(IPairStorageBackupFeed)).(*IPairStorageBackupFeed)

	return out0, err

}

// PairBackupFeed is a free data retrieval call binding the contract method 0x685e905b.
//
// Solidity: function pairBackupFeed(uint256 ) view returns((uint256,address))
func (_AvantisPairsAbis *AvantisPairsAbisSession) PairBackupFeed(arg0 *big.Int) (IPairStorageBackupFeed, error) {
	return _AvantisPairsAbis.Contract.PairBackupFeed(&_AvantisPairsAbis.CallOpts, arg0)
}

// PairBackupFeed is a free data retrieval call binding the contract method 0x685e905b.
//
// Solidity: function pairBackupFeed(uint256 ) view returns((uint256,address))
func (_AvantisPairsAbis *AvantisPairsAbisCallerSession) PairBackupFeed(arg0 *big.Int) (IPairStorageBackupFeed, error) {
	return _AvantisPairsAbis.Contract.PairBackupFeed(&_AvantisPairsAbis.CallOpts, arg0)
}

// PairCloseFeeP is a free data retrieval call binding the contract method 0x836a341a.
//
// Solidity: function pairCloseFeeP(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCaller) PairCloseFeeP(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisPairsAbis.contract.Call(opts, &out, "pairCloseFeeP", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairCloseFeeP is a free data retrieval call binding the contract method 0x836a341a.
//
// Solidity: function pairCloseFeeP(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisSession) PairCloseFeeP(arg0 *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairCloseFeeP(&_AvantisPairsAbis.CallOpts, arg0)
}

// PairCloseFeeP is a free data retrieval call binding the contract method 0x836a341a.
//
// Solidity: function pairCloseFeeP(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCallerSession) PairCloseFeeP(arg0 *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairCloseFeeP(&_AvantisPairsAbis.CallOpts, arg0)
}

// PairFeed is a free data retrieval call binding the contract method 0x1eaa005c.
//
// Solidity: function pairFeed(uint256 ) view returns((uint256,bytes32))
func (_AvantisPairsAbis *AvantisPairsAbisCaller) PairFeed(opts *bind.CallOpts, arg0 *big.Int) (IPairStorageFeed, error) {
	var out []interface{}
	err := _AvantisPairsAbis.contract.Call(opts, &out, "pairFeed", arg0)

	if err != nil {
		return *new(IPairStorageFeed), err
	}

	out0 := *abi.ConvertType(out[0], new(IPairStorageFeed)).(*IPairStorageFeed)

	return out0, err

}

// PairFeed is a free data retrieval call binding the contract method 0x1eaa005c.
//
// Solidity: function pairFeed(uint256 ) view returns((uint256,bytes32))
func (_AvantisPairsAbis *AvantisPairsAbisSession) PairFeed(arg0 *big.Int) (IPairStorageFeed, error) {
	return _AvantisPairsAbis.Contract.PairFeed(&_AvantisPairsAbis.CallOpts, arg0)
}

// PairFeed is a free data retrieval call binding the contract method 0x1eaa005c.
//
// Solidity: function pairFeed(uint256 ) view returns((uint256,bytes32))
func (_AvantisPairsAbis *AvantisPairsAbisCallerSession) PairFeed(arg0 *big.Int) (IPairStorageFeed, error) {
	return _AvantisPairsAbis.Contract.PairFeed(&_AvantisPairsAbis.CallOpts, arg0)
}

// PairGroupIndex is a free data retrieval call binding the contract method 0x3c44d0a3.
//
// Solidity: function pairGroupIndex(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCaller) PairGroupIndex(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisPairsAbis.contract.Call(opts, &out, "pairGroupIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairGroupIndex is a free data retrieval call binding the contract method 0x3c44d0a3.
//
// Solidity: function pairGroupIndex(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisSession) PairGroupIndex(arg0 *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairGroupIndex(&_AvantisPairsAbis.CallOpts, arg0)
}

// PairGroupIndex is a free data retrieval call binding the contract method 0x3c44d0a3.
//
// Solidity: function pairGroupIndex(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCallerSession) PairGroupIndex(arg0 *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairGroupIndex(&_AvantisPairsAbis.CallOpts, arg0)
}

// PairLimitOrderFeeP is a free data retrieval call binding the contract method 0x649a6488.
//
// Solidity: function pairLimitOrderFeeP(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCaller) PairLimitOrderFeeP(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisPairsAbis.contract.Call(opts, &out, "pairLimitOrderFeeP", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairLimitOrderFeeP is a free data retrieval call binding the contract method 0x649a6488.
//
// Solidity: function pairLimitOrderFeeP(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisSession) PairLimitOrderFeeP(arg0 *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairLimitOrderFeeP(&_AvantisPairsAbis.CallOpts, arg0)
}

// PairLimitOrderFeeP is a free data retrieval call binding the contract method 0x649a6488.
//
// Solidity: function pairLimitOrderFeeP(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCallerSession) PairLimitOrderFeeP(arg0 *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairLimitOrderFeeP(&_AvantisPairsAbis.CallOpts, arg0)
}

// PairMaxLeverage is a free data retrieval call binding the contract method 0x281b693c.
//
// Solidity: function pairMaxLeverage(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCaller) PairMaxLeverage(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisPairsAbis.contract.Call(opts, &out, "pairMaxLeverage", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairMaxLeverage is a free data retrieval call binding the contract method 0x281b693c.
//
// Solidity: function pairMaxLeverage(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisSession) PairMaxLeverage(arg0 *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairMaxLeverage(&_AvantisPairsAbis.CallOpts, arg0)
}

// PairMaxLeverage is a free data retrieval call binding the contract method 0x281b693c.
//
// Solidity: function pairMaxLeverage(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCallerSession) PairMaxLeverage(arg0 *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairMaxLeverage(&_AvantisPairsAbis.CallOpts, arg0)
}

// PairMaxOI is a free data retrieval call binding the contract method 0xb2e1b2d6.
//
// Solidity: function pairMaxOI(uint256 _pairIndex) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCaller) PairMaxOI(opts *bind.CallOpts, _pairIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisPairsAbis.contract.Call(opts, &out, "pairMaxOI", _pairIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairMaxOI is a free data retrieval call binding the contract method 0xb2e1b2d6.
//
// Solidity: function pairMaxOI(uint256 _pairIndex) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisSession) PairMaxOI(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairMaxOI(&_AvantisPairsAbis.CallOpts, _pairIndex)
}

// PairMaxOI is a free data retrieval call binding the contract method 0xb2e1b2d6.
//
// Solidity: function pairMaxOI(uint256 _pairIndex) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCallerSession) PairMaxOI(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairMaxOI(&_AvantisPairsAbis.CallOpts, _pairIndex)
}

// PairMinLevPosUSDC is a free data retrieval call binding the contract method 0x238cd23f.
//
// Solidity: function pairMinLevPosUSDC(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCaller) PairMinLevPosUSDC(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisPairsAbis.contract.Call(opts, &out, "pairMinLevPosUSDC", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairMinLevPosUSDC is a free data retrieval call binding the contract method 0x238cd23f.
//
// Solidity: function pairMinLevPosUSDC(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisSession) PairMinLevPosUSDC(arg0 *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairMinLevPosUSDC(&_AvantisPairsAbis.CallOpts, arg0)
}

// PairMinLevPosUSDC is a free data retrieval call binding the contract method 0x238cd23f.
//
// Solidity: function pairMinLevPosUSDC(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCallerSession) PairMinLevPosUSDC(arg0 *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairMinLevPosUSDC(&_AvantisPairsAbis.CallOpts, arg0)
}

// PairMinLeverage is a free data retrieval call binding the contract method 0x59a992d0.
//
// Solidity: function pairMinLeverage(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCaller) PairMinLeverage(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisPairsAbis.contract.Call(opts, &out, "pairMinLeverage", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairMinLeverage is a free data retrieval call binding the contract method 0x59a992d0.
//
// Solidity: function pairMinLeverage(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisSession) PairMinLeverage(arg0 *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairMinLeverage(&_AvantisPairsAbis.CallOpts, arg0)
}

// PairMinLeverage is a free data retrieval call binding the contract method 0x59a992d0.
//
// Solidity: function pairMinLeverage(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCallerSession) PairMinLeverage(arg0 *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairMinLeverage(&_AvantisPairsAbis.CallOpts, arg0)
}

// PairOpenFeeP is a free data retrieval call binding the contract method 0xfb3e519d.
//
// Solidity: function pairOpenFeeP(uint256 , uint256 , bool ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCaller) PairOpenFeeP(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 bool) (*big.Int, error) {
	var out []interface{}
	err := _AvantisPairsAbis.contract.Call(opts, &out, "pairOpenFeeP", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairOpenFeeP is a free data retrieval call binding the contract method 0xfb3e519d.
//
// Solidity: function pairOpenFeeP(uint256 , uint256 , bool ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisSession) PairOpenFeeP(arg0 *big.Int, arg1 *big.Int, arg2 bool) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairOpenFeeP(&_AvantisPairsAbis.CallOpts, arg0, arg1, arg2)
}

// PairOpenFeeP is a free data retrieval call binding the contract method 0xfb3e519d.
//
// Solidity: function pairOpenFeeP(uint256 , uint256 , bool ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCallerSession) PairOpenFeeP(arg0 *big.Int, arg1 *big.Int, arg2 bool) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairOpenFeeP(&_AvantisPairsAbis.CallOpts, arg0, arg1, arg2)
}

// PairPriceImpactMultiplier is a free data retrieval call binding the contract method 0x5fc8cd28.
//
// Solidity: function pairPriceImpactMultiplier(uint256 _pairIndex) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCaller) PairPriceImpactMultiplier(opts *bind.CallOpts, _pairIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisPairsAbis.contract.Call(opts, &out, "pairPriceImpactMultiplier", _pairIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairPriceImpactMultiplier is a free data retrieval call binding the contract method 0x5fc8cd28.
//
// Solidity: function pairPriceImpactMultiplier(uint256 _pairIndex) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisSession) PairPriceImpactMultiplier(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairPriceImpactMultiplier(&_AvantisPairsAbis.CallOpts, _pairIndex)
}

// PairPriceImpactMultiplier is a free data retrieval call binding the contract method 0x5fc8cd28.
//
// Solidity: function pairPriceImpactMultiplier(uint256 _pairIndex) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCallerSession) PairPriceImpactMultiplier(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairPriceImpactMultiplier(&_AvantisPairsAbis.CallOpts, _pairIndex)
}

// PairSkewImpactMultiplier is a free data retrieval call binding the contract method 0xe8b50f84.
//
// Solidity: function pairSkewImpactMultiplier(uint256 _pairIndex) view returns(int256)
func (_AvantisPairsAbis *AvantisPairsAbisCaller) PairSkewImpactMultiplier(opts *bind.CallOpts, _pairIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisPairsAbis.contract.Call(opts, &out, "pairSkewImpactMultiplier", _pairIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairSkewImpactMultiplier is a free data retrieval call binding the contract method 0xe8b50f84.
//
// Solidity: function pairSkewImpactMultiplier(uint256 _pairIndex) view returns(int256)
func (_AvantisPairsAbis *AvantisPairsAbisSession) PairSkewImpactMultiplier(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairSkewImpactMultiplier(&_AvantisPairsAbis.CallOpts, _pairIndex)
}

// PairSkewImpactMultiplier is a free data retrieval call binding the contract method 0xe8b50f84.
//
// Solidity: function pairSkewImpactMultiplier(uint256 _pairIndex) view returns(int256)
func (_AvantisPairsAbis *AvantisPairsAbisCallerSession) PairSkewImpactMultiplier(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairSkewImpactMultiplier(&_AvantisPairsAbis.CallOpts, _pairIndex)
}

// PairSpreadP is a free data retrieval call binding the contract method 0xa1d54e9b.
//
// Solidity: function pairSpreadP(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCaller) PairSpreadP(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisPairsAbis.contract.Call(opts, &out, "pairSpreadP", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairSpreadP is a free data retrieval call binding the contract method 0xa1d54e9b.
//
// Solidity: function pairSpreadP(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisSession) PairSpreadP(arg0 *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairSpreadP(&_AvantisPairsAbis.CallOpts, arg0)
}

// PairSpreadP is a free data retrieval call binding the contract method 0xa1d54e9b.
//
// Solidity: function pairSpreadP(uint256 ) view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCallerSession) PairSpreadP(arg0 *big.Int) (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairSpreadP(&_AvantisPairsAbis.CallOpts, arg0)
}

// PairsCount is a free data retrieval call binding the contract method 0xb81b2b71.
//
// Solidity: function pairsCount() view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCaller) PairsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AvantisPairsAbis.contract.Call(opts, &out, "pairsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairsCount is a free data retrieval call binding the contract method 0xb81b2b71.
//
// Solidity: function pairsCount() view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisSession) PairsCount() (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairsCount(&_AvantisPairsAbis.CallOpts)
}

// PairsCount is a free data retrieval call binding the contract method 0xb81b2b71.
//
// Solidity: function pairsCount() view returns(uint256)
func (_AvantisPairsAbis *AvantisPairsAbisCallerSession) PairsCount() (*big.Int, error) {
	return _AvantisPairsAbis.Contract.PairsCount(&_AvantisPairsAbis.CallOpts)
}

// PairJob is a paid mutator transaction binding the contract method 0x302f81fc.
//
// Solidity: function pairJob(uint256 ) returns(string, string, bytes32, address, uint256)
func (_AvantisPairsAbis *AvantisPairsAbisTransactor) PairJob(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _AvantisPairsAbis.contract.Transact(opts, "pairJob", arg0)
}

// PairJob is a paid mutator transaction binding the contract method 0x302f81fc.
//
// Solidity: function pairJob(uint256 ) returns(string, string, bytes32, address, uint256)
func (_AvantisPairsAbis *AvantisPairsAbisSession) PairJob(arg0 *big.Int) (*types.Transaction, error) {
	return _AvantisPairsAbis.Contract.PairJob(&_AvantisPairsAbis.TransactOpts, arg0)
}

// PairJob is a paid mutator transaction binding the contract method 0x302f81fc.
//
// Solidity: function pairJob(uint256 ) returns(string, string, bytes32, address, uint256)
func (_AvantisPairsAbis *AvantisPairsAbisTransactorSession) PairJob(arg0 *big.Int) (*types.Transaction, error) {
	return _AvantisPairsAbis.Contract.PairJob(&_AvantisPairsAbis.TransactOpts, arg0)
}

// UpdateGroupOI is a paid mutator transaction binding the contract method 0x4ffe8aec.
//
// Solidity: function updateGroupOI(uint256 , uint256 , bool , bool ) returns()
func (_AvantisPairsAbis *AvantisPairsAbisTransactor) UpdateGroupOI(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int, arg2 bool, arg3 bool) (*types.Transaction, error) {
	return _AvantisPairsAbis.contract.Transact(opts, "updateGroupOI", arg0, arg1, arg2, arg3)
}

// UpdateGroupOI is a paid mutator transaction binding the contract method 0x4ffe8aec.
//
// Solidity: function updateGroupOI(uint256 , uint256 , bool , bool ) returns()
func (_AvantisPairsAbis *AvantisPairsAbisSession) UpdateGroupOI(arg0 *big.Int, arg1 *big.Int, arg2 bool, arg3 bool) (*types.Transaction, error) {
	return _AvantisPairsAbis.Contract.UpdateGroupOI(&_AvantisPairsAbis.TransactOpts, arg0, arg1, arg2, arg3)
}

// UpdateGroupOI is a paid mutator transaction binding the contract method 0x4ffe8aec.
//
// Solidity: function updateGroupOI(uint256 , uint256 , bool , bool ) returns()
func (_AvantisPairsAbis *AvantisPairsAbisTransactorSession) UpdateGroupOI(arg0 *big.Int, arg1 *big.Int, arg2 bool, arg3 bool) (*types.Transaction, error) {
	return _AvantisPairsAbis.Contract.UpdateGroupOI(&_AvantisPairsAbis.TransactOpts, arg0, arg1, arg2, arg3)
}

// AvantisPairsAbisBlockOILimitsSetIterator is returned from FilterBlockOILimitsSet and is used to iterate over the raw logs and unpacked data for BlockOILimitsSet events raised by the AvantisPairsAbis contract.
type AvantisPairsAbisBlockOILimitsSetIterator struct {
	Event *AvantisPairsAbisBlockOILimitsSet // Event containing the contract specifics and raw log

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
func (it *AvantisPairsAbisBlockOILimitsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisPairsAbisBlockOILimitsSet)
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
		it.Event = new(AvantisPairsAbisBlockOILimitsSet)
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
func (it *AvantisPairsAbisBlockOILimitsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisPairsAbisBlockOILimitsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisPairsAbisBlockOILimitsSet represents a BlockOILimitsSet event raised by the AvantisPairsAbis contract.
type AvantisPairsAbisBlockOILimitsSet struct {
	PairIndex []*big.Int
	Limits    []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlockOILimitsSet is a free log retrieval operation binding the contract event 0xc5c5af5d05689513c60f635389c8449ac23f2b277af3a15f88c21c6de36fcf14.
//
// Solidity: event BlockOILimitsSet(uint256[] pairIndex, uint256[] limits)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) FilterBlockOILimitsSet(opts *bind.FilterOpts) (*AvantisPairsAbisBlockOILimitsSetIterator, error) {

	logs, sub, err := _AvantisPairsAbis.contract.FilterLogs(opts, "BlockOILimitsSet")
	if err != nil {
		return nil, err
	}
	return &AvantisPairsAbisBlockOILimitsSetIterator{contract: _AvantisPairsAbis.contract, event: "BlockOILimitsSet", logs: logs, sub: sub}, nil
}

// WatchBlockOILimitsSet is a free log subscription operation binding the contract event 0xc5c5af5d05689513c60f635389c8449ac23f2b277af3a15f88c21c6de36fcf14.
//
// Solidity: event BlockOILimitsSet(uint256[] pairIndex, uint256[] limits)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) WatchBlockOILimitsSet(opts *bind.WatchOpts, sink chan<- *AvantisPairsAbisBlockOILimitsSet) (event.Subscription, error) {

	logs, sub, err := _AvantisPairsAbis.contract.WatchLogs(opts, "BlockOILimitsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisPairsAbisBlockOILimitsSet)
				if err := _AvantisPairsAbis.contract.UnpackLog(event, "BlockOILimitsSet", log); err != nil {
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

// ParseBlockOILimitsSet is a log parse operation binding the contract event 0xc5c5af5d05689513c60f635389c8449ac23f2b277af3a15f88c21c6de36fcf14.
//
// Solidity: event BlockOILimitsSet(uint256[] pairIndex, uint256[] limits)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) ParseBlockOILimitsSet(log types.Log) (*AvantisPairsAbisBlockOILimitsSet, error) {
	event := new(AvantisPairsAbisBlockOILimitsSet)
	if err := _AvantisPairsAbis.contract.UnpackLog(event, "BlockOILimitsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvantisPairsAbisFeeAddedIterator is returned from FilterFeeAdded and is used to iterate over the raw logs and unpacked data for FeeAdded events raised by the AvantisPairsAbis contract.
type AvantisPairsAbisFeeAddedIterator struct {
	Event *AvantisPairsAbisFeeAdded // Event containing the contract specifics and raw log

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
func (it *AvantisPairsAbisFeeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisPairsAbisFeeAdded)
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
		it.Event = new(AvantisPairsAbisFeeAdded)
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
func (it *AvantisPairsAbisFeeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisPairsAbisFeeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisPairsAbisFeeAdded represents a FeeAdded event raised by the AvantisPairsAbis contract.
type AvantisPairsAbisFeeAdded struct {
	Index *big.Int
	Name  string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFeeAdded is a free log retrieval operation binding the contract event 0x482049823c85e038e099fe4f2b901487c4800def71c9a3f5bae2de8381ec54f6.
//
// Solidity: event FeeAdded(uint256 index, string name)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) FilterFeeAdded(opts *bind.FilterOpts) (*AvantisPairsAbisFeeAddedIterator, error) {

	logs, sub, err := _AvantisPairsAbis.contract.FilterLogs(opts, "FeeAdded")
	if err != nil {
		return nil, err
	}
	return &AvantisPairsAbisFeeAddedIterator{contract: _AvantisPairsAbis.contract, event: "FeeAdded", logs: logs, sub: sub}, nil
}

// WatchFeeAdded is a free log subscription operation binding the contract event 0x482049823c85e038e099fe4f2b901487c4800def71c9a3f5bae2de8381ec54f6.
//
// Solidity: event FeeAdded(uint256 index, string name)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) WatchFeeAdded(opts *bind.WatchOpts, sink chan<- *AvantisPairsAbisFeeAdded) (event.Subscription, error) {

	logs, sub, err := _AvantisPairsAbis.contract.WatchLogs(opts, "FeeAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisPairsAbisFeeAdded)
				if err := _AvantisPairsAbis.contract.UnpackLog(event, "FeeAdded", log); err != nil {
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

// ParseFeeAdded is a log parse operation binding the contract event 0x482049823c85e038e099fe4f2b901487c4800def71c9a3f5bae2de8381ec54f6.
//
// Solidity: event FeeAdded(uint256 index, string name)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) ParseFeeAdded(log types.Log) (*AvantisPairsAbisFeeAdded, error) {
	event := new(AvantisPairsAbisFeeAdded)
	if err := _AvantisPairsAbis.contract.UnpackLog(event, "FeeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvantisPairsAbisFeeUpdatedIterator is returned from FilterFeeUpdated and is used to iterate over the raw logs and unpacked data for FeeUpdated events raised by the AvantisPairsAbis contract.
type AvantisPairsAbisFeeUpdatedIterator struct {
	Event *AvantisPairsAbisFeeUpdated // Event containing the contract specifics and raw log

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
func (it *AvantisPairsAbisFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisPairsAbisFeeUpdated)
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
		it.Event = new(AvantisPairsAbisFeeUpdated)
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
func (it *AvantisPairsAbisFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisPairsAbisFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisPairsAbisFeeUpdated represents a FeeUpdated event raised by the AvantisPairsAbis contract.
type AvantisPairsAbisFeeUpdated struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFeeUpdated is a free log retrieval operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 index)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) FilterFeeUpdated(opts *bind.FilterOpts) (*AvantisPairsAbisFeeUpdatedIterator, error) {

	logs, sub, err := _AvantisPairsAbis.contract.FilterLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return &AvantisPairsAbisFeeUpdatedIterator{contract: _AvantisPairsAbis.contract, event: "FeeUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeUpdated is a free log subscription operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 index)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) WatchFeeUpdated(opts *bind.WatchOpts, sink chan<- *AvantisPairsAbisFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _AvantisPairsAbis.contract.WatchLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisPairsAbisFeeUpdated)
				if err := _AvantisPairsAbis.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
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

// ParseFeeUpdated is a log parse operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 index)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) ParseFeeUpdated(log types.Log) (*AvantisPairsAbisFeeUpdated, error) {
	event := new(AvantisPairsAbisFeeUpdated)
	if err := _AvantisPairsAbis.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvantisPairsAbisGroupAddedIterator is returned from FilterGroupAdded and is used to iterate over the raw logs and unpacked data for GroupAdded events raised by the AvantisPairsAbis contract.
type AvantisPairsAbisGroupAddedIterator struct {
	Event *AvantisPairsAbisGroupAdded // Event containing the contract specifics and raw log

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
func (it *AvantisPairsAbisGroupAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisPairsAbisGroupAdded)
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
		it.Event = new(AvantisPairsAbisGroupAdded)
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
func (it *AvantisPairsAbisGroupAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisPairsAbisGroupAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisPairsAbisGroupAdded represents a GroupAdded event raised by the AvantisPairsAbis contract.
type AvantisPairsAbisGroupAdded struct {
	Index *big.Int
	Name  string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGroupAdded is a free log retrieval operation binding the contract event 0xaf17de8e82beccc440012117a600dc37e26925225d0f1ee192fc107eb3dcbca4.
//
// Solidity: event GroupAdded(uint256 index, string name)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) FilterGroupAdded(opts *bind.FilterOpts) (*AvantisPairsAbisGroupAddedIterator, error) {

	logs, sub, err := _AvantisPairsAbis.contract.FilterLogs(opts, "GroupAdded")
	if err != nil {
		return nil, err
	}
	return &AvantisPairsAbisGroupAddedIterator{contract: _AvantisPairsAbis.contract, event: "GroupAdded", logs: logs, sub: sub}, nil
}

// WatchGroupAdded is a free log subscription operation binding the contract event 0xaf17de8e82beccc440012117a600dc37e26925225d0f1ee192fc107eb3dcbca4.
//
// Solidity: event GroupAdded(uint256 index, string name)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) WatchGroupAdded(opts *bind.WatchOpts, sink chan<- *AvantisPairsAbisGroupAdded) (event.Subscription, error) {

	logs, sub, err := _AvantisPairsAbis.contract.WatchLogs(opts, "GroupAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisPairsAbisGroupAdded)
				if err := _AvantisPairsAbis.contract.UnpackLog(event, "GroupAdded", log); err != nil {
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

// ParseGroupAdded is a log parse operation binding the contract event 0xaf17de8e82beccc440012117a600dc37e26925225d0f1ee192fc107eb3dcbca4.
//
// Solidity: event GroupAdded(uint256 index, string name)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) ParseGroupAdded(log types.Log) (*AvantisPairsAbisGroupAdded, error) {
	event := new(AvantisPairsAbisGroupAdded)
	if err := _AvantisPairsAbis.contract.UnpackLog(event, "GroupAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvantisPairsAbisGroupUpdatedIterator is returned from FilterGroupUpdated and is used to iterate over the raw logs and unpacked data for GroupUpdated events raised by the AvantisPairsAbis contract.
type AvantisPairsAbisGroupUpdatedIterator struct {
	Event *AvantisPairsAbisGroupUpdated // Event containing the contract specifics and raw log

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
func (it *AvantisPairsAbisGroupUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisPairsAbisGroupUpdated)
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
		it.Event = new(AvantisPairsAbisGroupUpdated)
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
func (it *AvantisPairsAbisGroupUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisPairsAbisGroupUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisPairsAbisGroupUpdated represents a GroupUpdated event raised by the AvantisPairsAbis contract.
type AvantisPairsAbisGroupUpdated struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGroupUpdated is a free log retrieval operation binding the contract event 0xcfde8f228364c70f12cbbac5a88fc91ceca76dd750ac93364991a333b34afb8e.
//
// Solidity: event GroupUpdated(uint256 index)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) FilterGroupUpdated(opts *bind.FilterOpts) (*AvantisPairsAbisGroupUpdatedIterator, error) {

	logs, sub, err := _AvantisPairsAbis.contract.FilterLogs(opts, "GroupUpdated")
	if err != nil {
		return nil, err
	}
	return &AvantisPairsAbisGroupUpdatedIterator{contract: _AvantisPairsAbis.contract, event: "GroupUpdated", logs: logs, sub: sub}, nil
}

// WatchGroupUpdated is a free log subscription operation binding the contract event 0xcfde8f228364c70f12cbbac5a88fc91ceca76dd750ac93364991a333b34afb8e.
//
// Solidity: event GroupUpdated(uint256 index)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) WatchGroupUpdated(opts *bind.WatchOpts, sink chan<- *AvantisPairsAbisGroupUpdated) (event.Subscription, error) {

	logs, sub, err := _AvantisPairsAbis.contract.WatchLogs(opts, "GroupUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisPairsAbisGroupUpdated)
				if err := _AvantisPairsAbis.contract.UnpackLog(event, "GroupUpdated", log); err != nil {
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

// ParseGroupUpdated is a log parse operation binding the contract event 0xcfde8f228364c70f12cbbac5a88fc91ceca76dd750ac93364991a333b34afb8e.
//
// Solidity: event GroupUpdated(uint256 index)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) ParseGroupUpdated(log types.Log) (*AvantisPairsAbisGroupUpdated, error) {
	event := new(AvantisPairsAbisGroupUpdated)
	if err := _AvantisPairsAbis.contract.UnpackLog(event, "GroupUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvantisPairsAbisLossProtectionAddedIterator is returned from FilterLossProtectionAdded and is used to iterate over the raw logs and unpacked data for LossProtectionAdded events raised by the AvantisPairsAbis contract.
type AvantisPairsAbisLossProtectionAddedIterator struct {
	Event *AvantisPairsAbisLossProtectionAdded // Event containing the contract specifics and raw log

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
func (it *AvantisPairsAbisLossProtectionAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisPairsAbisLossProtectionAdded)
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
		it.Event = new(AvantisPairsAbisLossProtectionAdded)
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
func (it *AvantisPairsAbisLossProtectionAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisPairsAbisLossProtectionAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisPairsAbisLossProtectionAdded represents a LossProtectionAdded event raised by the AvantisPairsAbis contract.
type AvantisPairsAbisLossProtectionAdded struct {
	PairIndex  *big.Int
	Tier       []*big.Int
	Multiplier []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLossProtectionAdded is a free log retrieval operation binding the contract event 0xf66559a3ee7945e0e49289392f86db1a3a3049abc2a1c89664a702fa04273227.
//
// Solidity: event LossProtectionAdded(uint256 pairIndex, uint256[] tier, uint256[] multiplier)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) FilterLossProtectionAdded(opts *bind.FilterOpts) (*AvantisPairsAbisLossProtectionAddedIterator, error) {

	logs, sub, err := _AvantisPairsAbis.contract.FilterLogs(opts, "LossProtectionAdded")
	if err != nil {
		return nil, err
	}
	return &AvantisPairsAbisLossProtectionAddedIterator{contract: _AvantisPairsAbis.contract, event: "LossProtectionAdded", logs: logs, sub: sub}, nil
}

// WatchLossProtectionAdded is a free log subscription operation binding the contract event 0xf66559a3ee7945e0e49289392f86db1a3a3049abc2a1c89664a702fa04273227.
//
// Solidity: event LossProtectionAdded(uint256 pairIndex, uint256[] tier, uint256[] multiplier)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) WatchLossProtectionAdded(opts *bind.WatchOpts, sink chan<- *AvantisPairsAbisLossProtectionAdded) (event.Subscription, error) {

	logs, sub, err := _AvantisPairsAbis.contract.WatchLogs(opts, "LossProtectionAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisPairsAbisLossProtectionAdded)
				if err := _AvantisPairsAbis.contract.UnpackLog(event, "LossProtectionAdded", log); err != nil {
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

// ParseLossProtectionAdded is a log parse operation binding the contract event 0xf66559a3ee7945e0e49289392f86db1a3a3049abc2a1c89664a702fa04273227.
//
// Solidity: event LossProtectionAdded(uint256 pairIndex, uint256[] tier, uint256[] multiplier)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) ParseLossProtectionAdded(log types.Log) (*AvantisPairsAbisLossProtectionAdded, error) {
	event := new(AvantisPairsAbisLossProtectionAdded)
	if err := _AvantisPairsAbis.contract.UnpackLog(event, "LossProtectionAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvantisPairsAbisOrderLimitsSetIterator is returned from FilterOrderLimitsSet and is used to iterate over the raw logs and unpacked data for OrderLimitsSet events raised by the AvantisPairsAbis contract.
type AvantisPairsAbisOrderLimitsSetIterator struct {
	Event *AvantisPairsAbisOrderLimitsSet // Event containing the contract specifics and raw log

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
func (it *AvantisPairsAbisOrderLimitsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisPairsAbisOrderLimitsSet)
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
		it.Event = new(AvantisPairsAbisOrderLimitsSet)
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
func (it *AvantisPairsAbisOrderLimitsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisPairsAbisOrderLimitsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisPairsAbisOrderLimitsSet represents a OrderLimitsSet event raised by the AvantisPairsAbis contract.
type AvantisPairsAbisOrderLimitsSet struct {
	PairIndex []*big.Int
	Limits    []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderLimitsSet is a free log retrieval operation binding the contract event 0x5ffcf71171c83ae021779d84f12b918f2bb07077c6d3047c5d3203022317f6a3.
//
// Solidity: event OrderLimitsSet(uint256[] pairIndex, uint256[] limits)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) FilterOrderLimitsSet(opts *bind.FilterOpts) (*AvantisPairsAbisOrderLimitsSetIterator, error) {

	logs, sub, err := _AvantisPairsAbis.contract.FilterLogs(opts, "OrderLimitsSet")
	if err != nil {
		return nil, err
	}
	return &AvantisPairsAbisOrderLimitsSetIterator{contract: _AvantisPairsAbis.contract, event: "OrderLimitsSet", logs: logs, sub: sub}, nil
}

// WatchOrderLimitsSet is a free log subscription operation binding the contract event 0x5ffcf71171c83ae021779d84f12b918f2bb07077c6d3047c5d3203022317f6a3.
//
// Solidity: event OrderLimitsSet(uint256[] pairIndex, uint256[] limits)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) WatchOrderLimitsSet(opts *bind.WatchOpts, sink chan<- *AvantisPairsAbisOrderLimitsSet) (event.Subscription, error) {

	logs, sub, err := _AvantisPairsAbis.contract.WatchLogs(opts, "OrderLimitsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisPairsAbisOrderLimitsSet)
				if err := _AvantisPairsAbis.contract.UnpackLog(event, "OrderLimitsSet", log); err != nil {
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

// ParseOrderLimitsSet is a log parse operation binding the contract event 0x5ffcf71171c83ae021779d84f12b918f2bb07077c6d3047c5d3203022317f6a3.
//
// Solidity: event OrderLimitsSet(uint256[] pairIndex, uint256[] limits)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) ParseOrderLimitsSet(log types.Log) (*AvantisPairsAbisOrderLimitsSet, error) {
	event := new(AvantisPairsAbisOrderLimitsSet)
	if err := _AvantisPairsAbis.contract.UnpackLog(event, "OrderLimitsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvantisPairsAbisPairAddedIterator is returned from FilterPairAdded and is used to iterate over the raw logs and unpacked data for PairAdded events raised by the AvantisPairsAbis contract.
type AvantisPairsAbisPairAddedIterator struct {
	Event *AvantisPairsAbisPairAdded // Event containing the contract specifics and raw log

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
func (it *AvantisPairsAbisPairAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisPairsAbisPairAdded)
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
		it.Event = new(AvantisPairsAbisPairAdded)
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
func (it *AvantisPairsAbisPairAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisPairsAbisPairAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisPairsAbisPairAdded represents a PairAdded event raised by the AvantisPairsAbis contract.
type AvantisPairsAbisPairAdded struct {
	Index *big.Int
	From  string
	To    string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPairAdded is a free log retrieval operation binding the contract event 0x3adfd40f2b74073df2a84238acdb7f460565a557b3cc13bddc8833289bf38e09.
//
// Solidity: event PairAdded(uint256 index, string from, string to)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) FilterPairAdded(opts *bind.FilterOpts) (*AvantisPairsAbisPairAddedIterator, error) {

	logs, sub, err := _AvantisPairsAbis.contract.FilterLogs(opts, "PairAdded")
	if err != nil {
		return nil, err
	}
	return &AvantisPairsAbisPairAddedIterator{contract: _AvantisPairsAbis.contract, event: "PairAdded", logs: logs, sub: sub}, nil
}

// WatchPairAdded is a free log subscription operation binding the contract event 0x3adfd40f2b74073df2a84238acdb7f460565a557b3cc13bddc8833289bf38e09.
//
// Solidity: event PairAdded(uint256 index, string from, string to)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) WatchPairAdded(opts *bind.WatchOpts, sink chan<- *AvantisPairsAbisPairAdded) (event.Subscription, error) {

	logs, sub, err := _AvantisPairsAbis.contract.WatchLogs(opts, "PairAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisPairsAbisPairAdded)
				if err := _AvantisPairsAbis.contract.UnpackLog(event, "PairAdded", log); err != nil {
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

// ParsePairAdded is a log parse operation binding the contract event 0x3adfd40f2b74073df2a84238acdb7f460565a557b3cc13bddc8833289bf38e09.
//
// Solidity: event PairAdded(uint256 index, string from, string to)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) ParsePairAdded(log types.Log) (*AvantisPairsAbisPairAdded, error) {
	event := new(AvantisPairsAbisPairAdded)
	if err := _AvantisPairsAbis.contract.UnpackLog(event, "PairAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvantisPairsAbisPairUpdatedIterator is returned from FilterPairUpdated and is used to iterate over the raw logs and unpacked data for PairUpdated events raised by the AvantisPairsAbis contract.
type AvantisPairsAbisPairUpdatedIterator struct {
	Event *AvantisPairsAbisPairUpdated // Event containing the contract specifics and raw log

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
func (it *AvantisPairsAbisPairUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisPairsAbisPairUpdated)
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
		it.Event = new(AvantisPairsAbisPairUpdated)
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
func (it *AvantisPairsAbisPairUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisPairsAbisPairUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisPairsAbisPairUpdated represents a PairUpdated event raised by the AvantisPairsAbis contract.
type AvantisPairsAbisPairUpdated struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPairUpdated is a free log retrieval operation binding the contract event 0x123a1b961ae93e7acda9790b318237b175b45ac09277cd3614305d8baa3f1953.
//
// Solidity: event PairUpdated(uint256 index)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) FilterPairUpdated(opts *bind.FilterOpts) (*AvantisPairsAbisPairUpdatedIterator, error) {

	logs, sub, err := _AvantisPairsAbis.contract.FilterLogs(opts, "PairUpdated")
	if err != nil {
		return nil, err
	}
	return &AvantisPairsAbisPairUpdatedIterator{contract: _AvantisPairsAbis.contract, event: "PairUpdated", logs: logs, sub: sub}, nil
}

// WatchPairUpdated is a free log subscription operation binding the contract event 0x123a1b961ae93e7acda9790b318237b175b45ac09277cd3614305d8baa3f1953.
//
// Solidity: event PairUpdated(uint256 index)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) WatchPairUpdated(opts *bind.WatchOpts, sink chan<- *AvantisPairsAbisPairUpdated) (event.Subscription, error) {

	logs, sub, err := _AvantisPairsAbis.contract.WatchLogs(opts, "PairUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisPairsAbisPairUpdated)
				if err := _AvantisPairsAbis.contract.UnpackLog(event, "PairUpdated", log); err != nil {
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

// ParsePairUpdated is a log parse operation binding the contract event 0x123a1b961ae93e7acda9790b318237b175b45ac09277cd3614305d8baa3f1953.
//
// Solidity: event PairUpdated(uint256 index)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) ParsePairUpdated(log types.Log) (*AvantisPairsAbisPairUpdated, error) {
	event := new(AvantisPairsAbisPairUpdated)
	if err := _AvantisPairsAbis.contract.UnpackLog(event, "PairUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvantisPairsAbisSkewFeeAddedIterator is returned from FilterSkewFeeAdded and is used to iterate over the raw logs and unpacked data for SkewFeeAdded events raised by the AvantisPairsAbis contract.
type AvantisPairsAbisSkewFeeAddedIterator struct {
	Event *AvantisPairsAbisSkewFeeAdded // Event containing the contract specifics and raw log

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
func (it *AvantisPairsAbisSkewFeeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisPairsAbisSkewFeeAdded)
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
		it.Event = new(AvantisPairsAbisSkewFeeAdded)
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
func (it *AvantisPairsAbisSkewFeeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisPairsAbisSkewFeeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisPairsAbisSkewFeeAdded represents a SkewFeeAdded event raised by the AvantisPairsAbis contract.
type AvantisPairsAbisSkewFeeAdded struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSkewFeeAdded is a free log retrieval operation binding the contract event 0x5d7a16e490fc66ca6522d4ba0437ac1a3b97cf25666340a19e408176295826d7.
//
// Solidity: event SkewFeeAdded(uint256 index)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) FilterSkewFeeAdded(opts *bind.FilterOpts) (*AvantisPairsAbisSkewFeeAddedIterator, error) {

	logs, sub, err := _AvantisPairsAbis.contract.FilterLogs(opts, "SkewFeeAdded")
	if err != nil {
		return nil, err
	}
	return &AvantisPairsAbisSkewFeeAddedIterator{contract: _AvantisPairsAbis.contract, event: "SkewFeeAdded", logs: logs, sub: sub}, nil
}

// WatchSkewFeeAdded is a free log subscription operation binding the contract event 0x5d7a16e490fc66ca6522d4ba0437ac1a3b97cf25666340a19e408176295826d7.
//
// Solidity: event SkewFeeAdded(uint256 index)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) WatchSkewFeeAdded(opts *bind.WatchOpts, sink chan<- *AvantisPairsAbisSkewFeeAdded) (event.Subscription, error) {

	logs, sub, err := _AvantisPairsAbis.contract.WatchLogs(opts, "SkewFeeAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisPairsAbisSkewFeeAdded)
				if err := _AvantisPairsAbis.contract.UnpackLog(event, "SkewFeeAdded", log); err != nil {
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

// ParseSkewFeeAdded is a log parse operation binding the contract event 0x5d7a16e490fc66ca6522d4ba0437ac1a3b97cf25666340a19e408176295826d7.
//
// Solidity: event SkewFeeAdded(uint256 index)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) ParseSkewFeeAdded(log types.Log) (*AvantisPairsAbisSkewFeeAdded, error) {
	event := new(AvantisPairsAbisSkewFeeAdded)
	if err := _AvantisPairsAbis.contract.UnpackLog(event, "SkewFeeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvantisPairsAbisSkewFeeUpdatedIterator is returned from FilterSkewFeeUpdated and is used to iterate over the raw logs and unpacked data for SkewFeeUpdated events raised by the AvantisPairsAbis contract.
type AvantisPairsAbisSkewFeeUpdatedIterator struct {
	Event *AvantisPairsAbisSkewFeeUpdated // Event containing the contract specifics and raw log

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
func (it *AvantisPairsAbisSkewFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisPairsAbisSkewFeeUpdated)
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
		it.Event = new(AvantisPairsAbisSkewFeeUpdated)
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
func (it *AvantisPairsAbisSkewFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisPairsAbisSkewFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisPairsAbisSkewFeeUpdated represents a SkewFeeUpdated event raised by the AvantisPairsAbis contract.
type AvantisPairsAbisSkewFeeUpdated struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSkewFeeUpdated is a free log retrieval operation binding the contract event 0x049d8e73ae29fd805de1f8eae372a5020742554b37118c1e33687e89ef3027ff.
//
// Solidity: event SkewFeeUpdated(uint256 index)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) FilterSkewFeeUpdated(opts *bind.FilterOpts) (*AvantisPairsAbisSkewFeeUpdatedIterator, error) {

	logs, sub, err := _AvantisPairsAbis.contract.FilterLogs(opts, "SkewFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &AvantisPairsAbisSkewFeeUpdatedIterator{contract: _AvantisPairsAbis.contract, event: "SkewFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchSkewFeeUpdated is a free log subscription operation binding the contract event 0x049d8e73ae29fd805de1f8eae372a5020742554b37118c1e33687e89ef3027ff.
//
// Solidity: event SkewFeeUpdated(uint256 index)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) WatchSkewFeeUpdated(opts *bind.WatchOpts, sink chan<- *AvantisPairsAbisSkewFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _AvantisPairsAbis.contract.WatchLogs(opts, "SkewFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisPairsAbisSkewFeeUpdated)
				if err := _AvantisPairsAbis.contract.UnpackLog(event, "SkewFeeUpdated", log); err != nil {
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

// ParseSkewFeeUpdated is a log parse operation binding the contract event 0x049d8e73ae29fd805de1f8eae372a5020742554b37118c1e33687e89ef3027ff.
//
// Solidity: event SkewFeeUpdated(uint256 index)
func (_AvantisPairsAbis *AvantisPairsAbisFilterer) ParseSkewFeeUpdated(log types.Log) (*AvantisPairsAbisSkewFeeUpdated, error) {
	event := new(AvantisPairsAbisSkewFeeUpdated)
	if err := _AvantisPairsAbis.contract.UnpackLog(event, "SkewFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
