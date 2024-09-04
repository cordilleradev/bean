// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package avantis_multicall_abis

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

// IMulticallAggregatedOrder is an auto generated low-level Go binding around an user-defined struct.
type IMulticallAggregatedOrder struct {
	Order            ITradingStorageOpenLimitOrder
	LiquidationPrice *big.Int
}

// IMulticallAggregatedTrade is an auto generated low-level Go binding around an user-defined struct.
type IMulticallAggregatedTrade struct {
	Trade            ITradingStorageTrade
	TradeInfo        ITradingStorageTradeInfo
	RolloverFee      *big.Int
	LiquidationPrice *big.Int
}

// ITradingStorageOpenLimitOrder is an auto generated low-level Go binding around an user-defined struct.
type ITradingStorageOpenLimitOrder struct {
	Trader       common.Address
	PairIndex    *big.Int
	Index        *big.Int
	PositionSize *big.Int
	Buy          bool
	Leverage     *big.Int
	Tp           *big.Int
	Sl           *big.Int
	Price        *big.Int
	SlippageP    *big.Int
	Block        *big.Int
	ExecutionFee *big.Int
}

// ITradingStorageTrade is an auto generated low-level Go binding around an user-defined struct.
type ITradingStorageTrade struct {
	Trader           common.Address
	PairIndex        *big.Int
	Index            *big.Int
	InitialPosToken  *big.Int
	PositionSizeUSDC *big.Int
	OpenPrice        *big.Int
	Buy              bool
	Leverage         *big.Int
	Tp               *big.Int
	Sl               *big.Int
	Timestamp        *big.Int
}

// ITradingStorageTradeInfo is an auto generated low-level Go binding around an user-defined struct.
type ITradingStorageTradeInfo struct {
	OpenInterestUSDC  *big.Int
	TpLastUpdated     *big.Int
	SlLastUpdated     *big.Int
	BeingMarketClosed bool
	LossProtection    *big.Int
}

// Multicall3Call is an auto generated low-level Go binding around an user-defined struct.
type Multicall3Call struct {
	Target   common.Address
	CallData []byte
}

// Multicall3Call3 is an auto generated low-level Go binding around an user-defined struct.
type Multicall3Call3 struct {
	Target       common.Address
	AllowFailure bool
	CallData     []byte
}

// Multicall3Call3Value is an auto generated low-level Go binding around an user-defined struct.
type Multicall3Call3Value struct {
	Target       common.Address
	AllowFailure bool
	Value        *big.Int
	CallData     []byte
}

// Multicall3Result is an auto generated low-level Go binding around an user-defined struct.
type Multicall3Result struct {
	Success    bool
	ReturnData []byte
}

// AvantisMulticallAbisMetaData contains all meta data concerning the AvantisMulticallAbis contract.
var AvantisMulticallAbisMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_storage\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pairInfos\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pairsStorage\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_trading\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"aggregate\",\"inputs\":[{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structMulticall3.Call[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"returnData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"aggregate3\",\"inputs\":[{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structMulticall3.Call3[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowFailure\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"returnData\",\"type\":\"tuple[]\",\"internalType\":\"structMulticall3.Result[]\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"aggregate3Value\",\"inputs\":[{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structMulticall3.Call3Value[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowFailure\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"returnData\",\"type\":\"tuple[]\",\"internalType\":\"structMulticall3.Result[]\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"blockAndAggregate\",\"inputs\":[{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structMulticall3.Call[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"returnData\",\"type\":\"tuple[]\",\"internalType\":\"structMulticall3.Result[]\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getBasefee\",\"inputs\":[],\"outputs\":[{\"name\":\"basefee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlockHash\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlockNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"chainid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentBlockCoinbase\",\"inputs\":[],\"outputs\":[{\"name\":\"coinbase\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentBlockDifficulty\",\"inputs\":[],\"outputs\":[{\"name\":\"difficulty\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentBlockGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"gaslimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentBlockTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEthBalance\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFirstEmptyTradeIndexes\",\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"firstEmptyTradeIndexes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLastBlockHash\",\"inputs\":[],\"outputs\":[{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLongShortRatios\",\"inputs\":[],\"outputs\":[{\"name\":\"longRatio\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"shortRatio\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMargins\",\"inputs\":[],\"outputs\":[{\"name\":\"rolloverFeePerBlockP\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"rolloverFeePerBlockLong\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"rolloverFeePerBlockShort\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOpenLimitOrdersCounts\",\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"openLimitOrdersCounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPositions\",\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIMulticall.AggregatedTrade[]\",\"components\":[{\"name\":\"trade\",\"type\":\"tuple\",\"internalType\":\"structITradingStorage.Trade\",\"components\":[{\"name\":\"trader\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initialPosToken\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"positionSizeUSDC\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"openPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"buy\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"leverage\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sl\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"tradeInfo\",\"type\":\"tuple\",\"internalType\":\"structITradingStorage.TradeInfo\",\"components\":[{\"name\":\"openInterestUSDC\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tpLastUpdated\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slLastUpdated\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beingMarketClosed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lossProtection\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"rolloverFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"liquidationPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIMulticall.AggregatedOrder[]\",\"components\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structITradingStorage.OpenLimitOrder\",\"components\":[{\"name\":\"trader\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"positionSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"buy\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"leverage\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sl\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slippageP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"block\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"liquidationPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tryAggregate\",\"inputs\":[{\"name\":\"requireSuccess\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structMulticall3.Call[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"returnData\",\"type\":\"tuple[]\",\"internalType\":\"structMulticall3.Result[]\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"tryBlockAndAggregate\",\"inputs\":[{\"name\":\"requireSuccess\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structMulticall3.Call[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"returnData\",\"type\":\"tuple[]\",\"internalType\":\"structMulticall3.Result[]\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"updateContracts\",\"inputs\":[{\"name\":\"_storage\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pairInfos\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pairsStorage\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_trading\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// AvantisMulticallAbisABI is the input ABI used to generate the binding from.
// Deprecated: Use AvantisMulticallAbisMetaData.ABI instead.
var AvantisMulticallAbisABI = AvantisMulticallAbisMetaData.ABI

// AvantisMulticallAbis is an auto generated Go binding around an Ethereum contract.
type AvantisMulticallAbis struct {
	AvantisMulticallAbisCaller     // Read-only binding to the contract
	AvantisMulticallAbisTransactor // Write-only binding to the contract
	AvantisMulticallAbisFilterer   // Log filterer for contract events
}

// AvantisMulticallAbisCaller is an auto generated read-only Go binding around an Ethereum contract.
type AvantisMulticallAbisCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvantisMulticallAbisTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AvantisMulticallAbisTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvantisMulticallAbisFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AvantisMulticallAbisFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvantisMulticallAbisSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AvantisMulticallAbisSession struct {
	Contract     *AvantisMulticallAbis // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AvantisMulticallAbisCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AvantisMulticallAbisCallerSession struct {
	Contract *AvantisMulticallAbisCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// AvantisMulticallAbisTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AvantisMulticallAbisTransactorSession struct {
	Contract     *AvantisMulticallAbisTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// AvantisMulticallAbisRaw is an auto generated low-level Go binding around an Ethereum contract.
type AvantisMulticallAbisRaw struct {
	Contract *AvantisMulticallAbis // Generic contract binding to access the raw methods on
}

// AvantisMulticallAbisCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AvantisMulticallAbisCallerRaw struct {
	Contract *AvantisMulticallAbisCaller // Generic read-only contract binding to access the raw methods on
}

// AvantisMulticallAbisTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AvantisMulticallAbisTransactorRaw struct {
	Contract *AvantisMulticallAbisTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAvantisMulticallAbis creates a new instance of AvantisMulticallAbis, bound to a specific deployed contract.
func NewAvantisMulticallAbis(address common.Address, backend bind.ContractBackend) (*AvantisMulticallAbis, error) {
	contract, err := bindAvantisMulticallAbis(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AvantisMulticallAbis{AvantisMulticallAbisCaller: AvantisMulticallAbisCaller{contract: contract}, AvantisMulticallAbisTransactor: AvantisMulticallAbisTransactor{contract: contract}, AvantisMulticallAbisFilterer: AvantisMulticallAbisFilterer{contract: contract}}, nil
}

// NewAvantisMulticallAbisCaller creates a new read-only instance of AvantisMulticallAbis, bound to a specific deployed contract.
func NewAvantisMulticallAbisCaller(address common.Address, caller bind.ContractCaller) (*AvantisMulticallAbisCaller, error) {
	contract, err := bindAvantisMulticallAbis(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AvantisMulticallAbisCaller{contract: contract}, nil
}

// NewAvantisMulticallAbisTransactor creates a new write-only instance of AvantisMulticallAbis, bound to a specific deployed contract.
func NewAvantisMulticallAbisTransactor(address common.Address, transactor bind.ContractTransactor) (*AvantisMulticallAbisTransactor, error) {
	contract, err := bindAvantisMulticallAbis(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AvantisMulticallAbisTransactor{contract: contract}, nil
}

// NewAvantisMulticallAbisFilterer creates a new log filterer instance of AvantisMulticallAbis, bound to a specific deployed contract.
func NewAvantisMulticallAbisFilterer(address common.Address, filterer bind.ContractFilterer) (*AvantisMulticallAbisFilterer, error) {
	contract, err := bindAvantisMulticallAbis(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AvantisMulticallAbisFilterer{contract: contract}, nil
}

// bindAvantisMulticallAbis binds a generic wrapper to an already deployed contract.
func bindAvantisMulticallAbis(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AvantisMulticallAbisMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AvantisMulticallAbis *AvantisMulticallAbisRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AvantisMulticallAbis.Contract.AvantisMulticallAbisCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AvantisMulticallAbis *AvantisMulticallAbisRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.AvantisMulticallAbisTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AvantisMulticallAbis *AvantisMulticallAbisRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.AvantisMulticallAbisTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AvantisMulticallAbis.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.contract.Transact(opts, method, params...)
}

// GetBasefee is a free data retrieval call binding the contract method 0x3e64a696.
//
// Solidity: function getBasefee() view returns(uint256 basefee)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) GetBasefee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "getBasefee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBasefee is a free data retrieval call binding the contract method 0x3e64a696.
//
// Solidity: function getBasefee() view returns(uint256 basefee)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) GetBasefee() (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GetBasefee(&_AvantisMulticallAbis.CallOpts)
}

// GetBasefee is a free data retrieval call binding the contract method 0x3e64a696.
//
// Solidity: function getBasefee() view returns(uint256 basefee)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) GetBasefee() (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GetBasefee(&_AvantisMulticallAbis.CallOpts)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) GetBlockHash(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "getBlockHash", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _AvantisMulticallAbis.Contract.GetBlockHash(&_AvantisMulticallAbis.CallOpts, blockNumber)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _AvantisMulticallAbis.Contract.GetBlockHash(&_AvantisMulticallAbis.CallOpts, blockNumber)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) GetBlockNumber() (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GetBlockNumber(&_AvantisMulticallAbis.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) GetBlockNumber() (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GetBlockNumber(&_AvantisMulticallAbis.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainid)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainid)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) GetChainId() (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GetChainId(&_AvantisMulticallAbis.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainid)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) GetChainId() (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GetChainId(&_AvantisMulticallAbis.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) GetCurrentBlockCoinbase(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "getCurrentBlockCoinbase")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _AvantisMulticallAbis.Contract.GetCurrentBlockCoinbase(&_AvantisMulticallAbis.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _AvantisMulticallAbis.Contract.GetCurrentBlockCoinbase(&_AvantisMulticallAbis.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) GetCurrentBlockDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "getCurrentBlockDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GetCurrentBlockDifficulty(&_AvantisMulticallAbis.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GetCurrentBlockDifficulty(&_AvantisMulticallAbis.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) GetCurrentBlockGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "getCurrentBlockGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GetCurrentBlockGasLimit(&_AvantisMulticallAbis.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GetCurrentBlockGasLimit(&_AvantisMulticallAbis.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) GetCurrentBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "getCurrentBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GetCurrentBlockTimestamp(&_AvantisMulticallAbis.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GetCurrentBlockTimestamp(&_AvantisMulticallAbis.CallOpts)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) GetEthBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "getEthBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GetEthBalance(&_AvantisMulticallAbis.CallOpts, addr)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GetEthBalance(&_AvantisMulticallAbis.CallOpts, addr)
}

// GetFirstEmptyTradeIndexes is a free data retrieval call binding the contract method 0x977ca604.
//
// Solidity: function getFirstEmptyTradeIndexes(address userAddress) view returns(uint256[] firstEmptyTradeIndexes)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) GetFirstEmptyTradeIndexes(opts *bind.CallOpts, userAddress common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "getFirstEmptyTradeIndexes", userAddress)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetFirstEmptyTradeIndexes is a free data retrieval call binding the contract method 0x977ca604.
//
// Solidity: function getFirstEmptyTradeIndexes(address userAddress) view returns(uint256[] firstEmptyTradeIndexes)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) GetFirstEmptyTradeIndexes(userAddress common.Address) ([]*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GetFirstEmptyTradeIndexes(&_AvantisMulticallAbis.CallOpts, userAddress)
}

// GetFirstEmptyTradeIndexes is a free data retrieval call binding the contract method 0x977ca604.
//
// Solidity: function getFirstEmptyTradeIndexes(address userAddress) view returns(uint256[] firstEmptyTradeIndexes)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) GetFirstEmptyTradeIndexes(userAddress common.Address) ([]*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GetFirstEmptyTradeIndexes(&_AvantisMulticallAbis.CallOpts, userAddress)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) GetLastBlockHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "getLastBlockHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) GetLastBlockHash() ([32]byte, error) {
	return _AvantisMulticallAbis.Contract.GetLastBlockHash(&_AvantisMulticallAbis.CallOpts)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) GetLastBlockHash() ([32]byte, error) {
	return _AvantisMulticallAbis.Contract.GetLastBlockHash(&_AvantisMulticallAbis.CallOpts)
}

// GetLongShortRatios is a free data retrieval call binding the contract method 0x582b175f.
//
// Solidity: function getLongShortRatios() view returns(uint256[] longRatio, uint256[] shortRatio)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) GetLongShortRatios(opts *bind.CallOpts) (struct {
	LongRatio  []*big.Int
	ShortRatio []*big.Int
}, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "getLongShortRatios")

	outstruct := new(struct {
		LongRatio  []*big.Int
		ShortRatio []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LongRatio = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.ShortRatio = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetLongShortRatios is a free data retrieval call binding the contract method 0x582b175f.
//
// Solidity: function getLongShortRatios() view returns(uint256[] longRatio, uint256[] shortRatio)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) GetLongShortRatios() (struct {
	LongRatio  []*big.Int
	ShortRatio []*big.Int
}, error) {
	return _AvantisMulticallAbis.Contract.GetLongShortRatios(&_AvantisMulticallAbis.CallOpts)
}

// GetLongShortRatios is a free data retrieval call binding the contract method 0x582b175f.
//
// Solidity: function getLongShortRatios() view returns(uint256[] longRatio, uint256[] shortRatio)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) GetLongShortRatios() (struct {
	LongRatio  []*big.Int
	ShortRatio []*big.Int
}, error) {
	return _AvantisMulticallAbis.Contract.GetLongShortRatios(&_AvantisMulticallAbis.CallOpts)
}

// GetMargins is a free data retrieval call binding the contract method 0x8693da43.
//
// Solidity: function getMargins() view returns(uint256[] rolloverFeePerBlockP, uint256[] rolloverFeePerBlockLong, uint256[] rolloverFeePerBlockShort)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) GetMargins(opts *bind.CallOpts) (struct {
	RolloverFeePerBlockP     []*big.Int
	RolloverFeePerBlockLong  []*big.Int
	RolloverFeePerBlockShort []*big.Int
}, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "getMargins")

	outstruct := new(struct {
		RolloverFeePerBlockP     []*big.Int
		RolloverFeePerBlockLong  []*big.Int
		RolloverFeePerBlockShort []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RolloverFeePerBlockP = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.RolloverFeePerBlockLong = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.RolloverFeePerBlockShort = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetMargins is a free data retrieval call binding the contract method 0x8693da43.
//
// Solidity: function getMargins() view returns(uint256[] rolloverFeePerBlockP, uint256[] rolloverFeePerBlockLong, uint256[] rolloverFeePerBlockShort)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) GetMargins() (struct {
	RolloverFeePerBlockP     []*big.Int
	RolloverFeePerBlockLong  []*big.Int
	RolloverFeePerBlockShort []*big.Int
}, error) {
	return _AvantisMulticallAbis.Contract.GetMargins(&_AvantisMulticallAbis.CallOpts)
}

// GetMargins is a free data retrieval call binding the contract method 0x8693da43.
//
// Solidity: function getMargins() view returns(uint256[] rolloverFeePerBlockP, uint256[] rolloverFeePerBlockLong, uint256[] rolloverFeePerBlockShort)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) GetMargins() (struct {
	RolloverFeePerBlockP     []*big.Int
	RolloverFeePerBlockLong  []*big.Int
	RolloverFeePerBlockShort []*big.Int
}, error) {
	return _AvantisMulticallAbis.Contract.GetMargins(&_AvantisMulticallAbis.CallOpts)
}

// GetOpenLimitOrdersCounts is a free data retrieval call binding the contract method 0xed4a96a9.
//
// Solidity: function getOpenLimitOrdersCounts(address userAddress) view returns(uint256[] openLimitOrdersCounts)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) GetOpenLimitOrdersCounts(opts *bind.CallOpts, userAddress common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "getOpenLimitOrdersCounts", userAddress)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOpenLimitOrdersCounts is a free data retrieval call binding the contract method 0xed4a96a9.
//
// Solidity: function getOpenLimitOrdersCounts(address userAddress) view returns(uint256[] openLimitOrdersCounts)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) GetOpenLimitOrdersCounts(userAddress common.Address) ([]*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GetOpenLimitOrdersCounts(&_AvantisMulticallAbis.CallOpts, userAddress)
}

// GetOpenLimitOrdersCounts is a free data retrieval call binding the contract method 0xed4a96a9.
//
// Solidity: function getOpenLimitOrdersCounts(address userAddress) view returns(uint256[] openLimitOrdersCounts)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) GetOpenLimitOrdersCounts(userAddress common.Address) ([]*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GetOpenLimitOrdersCounts(&_AvantisMulticallAbis.CallOpts, userAddress)
}

// GetPositions is a free data retrieval call binding the contract method 0x3eeb530e.
//
// Solidity: function getPositions(address userAddress) view returns(((address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,bool,uint256),uint256,uint256)[], ((address,uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256)[])
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) GetPositions(opts *bind.CallOpts, userAddress common.Address) ([]IMulticallAggregatedTrade, []IMulticallAggregatedOrder, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "getPositions", userAddress)

	if err != nil {
		return *new([]IMulticallAggregatedTrade), *new([]IMulticallAggregatedOrder), err
	}

	out0 := *abi.ConvertType(out[0], new([]IMulticallAggregatedTrade)).(*[]IMulticallAggregatedTrade)
	out1 := *abi.ConvertType(out[1], new([]IMulticallAggregatedOrder)).(*[]IMulticallAggregatedOrder)

	return out0, out1, err

}

// GetPositions is a free data retrieval call binding the contract method 0x3eeb530e.
//
// Solidity: function getPositions(address userAddress) view returns(((address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,bool,uint256),uint256,uint256)[], ((address,uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256)[])
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) GetPositions(userAddress common.Address) ([]IMulticallAggregatedTrade, []IMulticallAggregatedOrder, error) {
	return _AvantisMulticallAbis.Contract.GetPositions(&_AvantisMulticallAbis.CallOpts, userAddress)
}

// GetPositions is a free data retrieval call binding the contract method 0x3eeb530e.
//
// Solidity: function getPositions(address userAddress) view returns(((address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256),(uint256,uint256,uint256,bool,uint256),uint256,uint256)[], ((address,uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256)[])
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) GetPositions(userAddress common.Address) ([]IMulticallAggregatedTrade, []IMulticallAggregatedOrder, error) {
	return _AvantisMulticallAbis.Contract.GetPositions(&_AvantisMulticallAbis.CallOpts, userAddress)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes[] returnData)
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactor) Aggregate(opts *bind.TransactOpts, calls []Multicall3Call) (*types.Transaction, error) {
	return _AvantisMulticallAbis.contract.Transact(opts, "aggregate", calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes[] returnData)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) Aggregate(calls []Multicall3Call) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.Aggregate(&_AvantisMulticallAbis.TransactOpts, calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes[] returnData)
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorSession) Aggregate(calls []Multicall3Call) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.Aggregate(&_AvantisMulticallAbis.TransactOpts, calls)
}

// Aggregate3 is a paid mutator transaction binding the contract method 0x82ad56cb.
//
// Solidity: function aggregate3((address,bool,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactor) Aggregate3(opts *bind.TransactOpts, calls []Multicall3Call3) (*types.Transaction, error) {
	return _AvantisMulticallAbis.contract.Transact(opts, "aggregate3", calls)
}

// Aggregate3 is a paid mutator transaction binding the contract method 0x82ad56cb.
//
// Solidity: function aggregate3((address,bool,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) Aggregate3(calls []Multicall3Call3) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.Aggregate3(&_AvantisMulticallAbis.TransactOpts, calls)
}

// Aggregate3 is a paid mutator transaction binding the contract method 0x82ad56cb.
//
// Solidity: function aggregate3((address,bool,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorSession) Aggregate3(calls []Multicall3Call3) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.Aggregate3(&_AvantisMulticallAbis.TransactOpts, calls)
}

// Aggregate3Value is a paid mutator transaction binding the contract method 0x174dea71.
//
// Solidity: function aggregate3Value((address,bool,uint256,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactor) Aggregate3Value(opts *bind.TransactOpts, calls []Multicall3Call3Value) (*types.Transaction, error) {
	return _AvantisMulticallAbis.contract.Transact(opts, "aggregate3Value", calls)
}

// Aggregate3Value is a paid mutator transaction binding the contract method 0x174dea71.
//
// Solidity: function aggregate3Value((address,bool,uint256,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) Aggregate3Value(calls []Multicall3Call3Value) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.Aggregate3Value(&_AvantisMulticallAbis.TransactOpts, calls)
}

// Aggregate3Value is a paid mutator transaction binding the contract method 0x174dea71.
//
// Solidity: function aggregate3Value((address,bool,uint256,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorSession) Aggregate3Value(calls []Multicall3Call3Value) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.Aggregate3Value(&_AvantisMulticallAbis.TransactOpts, calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactor) BlockAndAggregate(opts *bind.TransactOpts, calls []Multicall3Call) (*types.Transaction, error) {
	return _AvantisMulticallAbis.contract.Transact(opts, "blockAndAggregate", calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) BlockAndAggregate(calls []Multicall3Call) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.BlockAndAggregate(&_AvantisMulticallAbis.TransactOpts, calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorSession) BlockAndAggregate(calls []Multicall3Call) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.BlockAndAggregate(&_AvantisMulticallAbis.TransactOpts, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactor) TryAggregate(opts *bind.TransactOpts, requireSuccess bool, calls []Multicall3Call) (*types.Transaction, error) {
	return _AvantisMulticallAbis.contract.Transact(opts, "tryAggregate", requireSuccess, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) TryAggregate(requireSuccess bool, calls []Multicall3Call) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.TryAggregate(&_AvantisMulticallAbis.TransactOpts, requireSuccess, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorSession) TryAggregate(requireSuccess bool, calls []Multicall3Call) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.TryAggregate(&_AvantisMulticallAbis.TransactOpts, requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactor) TryBlockAndAggregate(opts *bind.TransactOpts, requireSuccess bool, calls []Multicall3Call) (*types.Transaction, error) {
	return _AvantisMulticallAbis.contract.Transact(opts, "tryBlockAndAggregate", requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) TryBlockAndAggregate(requireSuccess bool, calls []Multicall3Call) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.TryBlockAndAggregate(&_AvantisMulticallAbis.TransactOpts, requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorSession) TryBlockAndAggregate(requireSuccess bool, calls []Multicall3Call) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.TryBlockAndAggregate(&_AvantisMulticallAbis.TransactOpts, requireSuccess, calls)
}

// UpdateContracts is a paid mutator transaction binding the contract method 0x0b751279.
//
// Solidity: function updateContracts(address _storage, address _pairInfos, address _pairsStorage, address _trading) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactor) UpdateContracts(opts *bind.TransactOpts, _storage common.Address, _pairInfos common.Address, _pairsStorage common.Address, _trading common.Address) (*types.Transaction, error) {
	return _AvantisMulticallAbis.contract.Transact(opts, "updateContracts", _storage, _pairInfos, _pairsStorage, _trading)
}

// UpdateContracts is a paid mutator transaction binding the contract method 0x0b751279.
//
// Solidity: function updateContracts(address _storage, address _pairInfos, address _pairsStorage, address _trading) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) UpdateContracts(_storage common.Address, _pairInfos common.Address, _pairsStorage common.Address, _trading common.Address) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.UpdateContracts(&_AvantisMulticallAbis.TransactOpts, _storage, _pairInfos, _pairsStorage, _trading)
}

// UpdateContracts is a paid mutator transaction binding the contract method 0x0b751279.
//
// Solidity: function updateContracts(address _storage, address _pairInfos, address _pairsStorage, address _trading) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorSession) UpdateContracts(_storage common.Address, _pairInfos common.Address, _pairsStorage common.Address, _trading common.Address) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.UpdateContracts(&_AvantisMulticallAbis.TransactOpts, _storage, _pairInfos, _pairsStorage, _trading)
}
