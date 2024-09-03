// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gains_abis

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

// IAddressStoreAddresses is an auto generated low-level Go binding around an user-defined struct.
type IAddressStoreAddresses struct {
	Gns        common.Address
	GnsStaking common.Address
}

// IPairsStorageGroupLiquidationParams is an auto generated low-level Go binding around an user-defined struct.
type IPairsStorageGroupLiquidationParams struct {
	MaxLiqSpreadP      *big.Int
	StartLiqThresholdP *big.Int
	EndLiqThresholdP   *big.Int
	StartLeverage      *big.Int
	EndLeverage        *big.Int
}

// ITradingStorageCollateral is an auto generated low-level Go binding around an user-defined struct.
type ITradingStorageCollateral struct {
	Collateral     common.Address
	IsActive       bool
	Placeholder    *big.Int
	Precision      *big.Int
	PrecisionDelta *big.Int
}

// ITradingStorageCounter is an auto generated low-level Go binding around an user-defined struct.
type ITradingStorageCounter struct {
	CurrentIndex uint32
	OpenCount    uint32
	Placeholder  *big.Int
}

// ITradingStorageId is an auto generated low-level Go binding around an user-defined struct.
type ITradingStorageId struct {
	User  common.Address
	Index uint32
}

// ITradingStoragePendingOrder is an auto generated low-level Go binding around an user-defined struct.
type ITradingStoragePendingOrder struct {
	Trade        ITradingStorageTrade
	User         common.Address
	Index        uint32
	IsOpen       bool
	OrderType    uint8
	CreatedBlock uint32
	MaxSlippageP uint16
}

// ITradingStorageTrade is an auto generated low-level Go binding around an user-defined struct.
type ITradingStorageTrade struct {
	User             common.Address
	Index            uint32
	PairIndex        uint16
	Leverage         *big.Int
	Long             bool
	IsOpen           bool
	CollateralIndex  uint8
	TradeType        uint8
	CollateralAmount *big.Int
	OpenPrice        uint64
	Tp               uint64
	Sl               uint64
	Placeholder      *big.Int
}

// ITradingStorageTradeInfo is an auto generated low-level Go binding around an user-defined struct.
type ITradingStorageTradeInfo struct {
	CreatedBlock         uint32
	TpLastUpdatedBlock   uint32
	SlLastUpdatedBlock   uint32
	MaxSlippageP         uint16
	LastOiUpdateTs       *big.Int
	CollateralPriceUsd   *big.Int
	ContractsVersion     uint8
	LastPosIncreaseBlock uint32
	Placeholder          uint8
}

// GainsAbisMetaData contains all meta data concerning the GainsAbis contract.
var GainsAbisMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BelowMin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CollateralAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CollateralAlreadyDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DoesntExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCollateralIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInputLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxSlippageZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingCollaterals\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Overflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TradeInfoCollateralPriceUsdZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TradeOpenPriceZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TradePairNotListed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TradePositionSizeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TradeSlInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TradeTpInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongAccess\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongOrderType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongTradeType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroValue\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIAddressStore.Role\",\"name\":\"role\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"access\",\"type\":\"bool\"}],\"name\":\"AccessControlUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"gns\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gnsStaking\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structIAddressStore.Addresses\",\"name\":\"addresses\",\"type\":\"tuple\"}],\"name\":\"AddressesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gToken\",\"type\":\"address\"}],\"name\":\"CollateralAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"CollateralDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"CollateralUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gToken\",\"type\":\"address\"}],\"name\":\"GTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structITradingStorage.Id\",\"name\":\"tradeId\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"openPrice\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"tp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sl\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"maxSlippageP\",\"type\":\"uint16\"}],\"name\":\"OpenOrderDetailsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structITradingStorage.Id\",\"name\":\"orderId\",\"type\":\"tuple\"}],\"name\":\"PendingOrderClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"pairIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"leverage\",\"type\":\"uint24\"},{\"internalType\":\"bool\",\"name\":\"long\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"collateralIndex\",\"type\":\"uint8\"},{\"internalType\":\"enumITradingStorage.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"internalType\":\"uint120\",\"name\":\"collateralAmount\",\"type\":\"uint120\"},{\"internalType\":\"uint64\",\"name\":\"openPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"tp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sl\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"__placeholder\",\"type\":\"uint192\"}],\"internalType\":\"structITradingStorage.Trade\",\"name\":\"trade\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"enumITradingStorage.PendingOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"createdBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxSlippageP\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structITradingStorage.PendingOrder\",\"name\":\"pendingOrder\",\"type\":\"tuple\"}],\"name\":\"PendingOrderStored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structITradingStorage.Id\",\"name\":\"tradeId\",\"type\":\"tuple\"}],\"name\":\"TradeClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structITradingStorage.Id\",\"name\":\"tradeId\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint120\",\"name\":\"collateralAmount\",\"type\":\"uint120\"}],\"name\":\"TradeCollateralUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structITradingStorage.Id\",\"name\":\"tradeId\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"maxClosingSlippageP\",\"type\":\"uint16\"}],\"name\":\"TradeMaxClosingSlippagePUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structITradingStorage.Id\",\"name\":\"tradeId\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint120\",\"name\":\"collateralAmount\",\"type\":\"uint120\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"leverage\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"openPrice\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newTp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newSl\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPartialIncrease\",\"type\":\"bool\"}],\"name\":\"TradePositionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structITradingStorage.Id\",\"name\":\"tradeId\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newSl\",\"type\":\"uint64\"}],\"name\":\"TradeSlUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"pairIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"leverage\",\"type\":\"uint24\"},{\"internalType\":\"bool\",\"name\":\"long\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"collateralIndex\",\"type\":\"uint8\"},{\"internalType\":\"enumITradingStorage.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"internalType\":\"uint120\",\"name\":\"collateralAmount\",\"type\":\"uint120\"},{\"internalType\":\"uint64\",\"name\":\"openPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"tp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sl\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"__placeholder\",\"type\":\"uint192\"}],\"indexed\":false,\"internalType\":\"structITradingStorage.Trade\",\"name\":\"trade\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"createdBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"tpLastUpdatedBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"slLastUpdatedBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxSlippageP\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"lastOiUpdateTs\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"collateralPriceUsd\",\"type\":\"uint48\"},{\"internalType\":\"enumITradingStorage.ContractsVersion\",\"name\":\"contractsVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"lastPosIncreaseBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"__placeholder\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structITradingStorage.TradeInfo\",\"name\":\"tradeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint40\",\"name\":\"maxLiqSpreadP\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"startLiqThresholdP\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"endLiqThresholdP\",\"type\":\"uint40\"},{\"internalType\":\"uint24\",\"name\":\"startLeverage\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"endLeverage\",\"type\":\"uint24\"}],\"indexed\":false,\"internalType\":\"structIPairsStorage.GroupLiquidationParams\",\"name\":\"liquidationParams\",\"type\":\"tuple\"}],\"name\":\"TradeStored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structITradingStorage.Id\",\"name\":\"tradeId\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newTp\",\"type\":\"uint64\"}],\"name\":\"TradeTpUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumITradingStorage.TradingActivated\",\"name\":\"activated\",\"type\":\"uint8\"}],\"name\":\"TradingActivatedUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gToken\",\"type\":\"address\"}],\"name\":\"addCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structITradingStorage.Id\",\"name\":\"_orderId\",\"type\":\"tuple\"}],\"name\":\"closePendingOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structITradingStorage.Id\",\"name\":\"_tradeId\",\"type\":\"tuple\"}],\"name\":\"closeTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddresses\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"gns\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gnsStaking\",\"type\":\"address\"}],\"internalType\":\"structIAddressStore.Addresses\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getAllPendingOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"pairIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"leverage\",\"type\":\"uint24\"},{\"internalType\":\"bool\",\"name\":\"long\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"collateralIndex\",\"type\":\"uint8\"},{\"internalType\":\"enumITradingStorage.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"internalType\":\"uint120\",\"name\":\"collateralAmount\",\"type\":\"uint120\"},{\"internalType\":\"uint64\",\"name\":\"openPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"tp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sl\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"__placeholder\",\"type\":\"uint192\"}],\"internalType\":\"structITradingStorage.Trade\",\"name\":\"trade\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"enumITradingStorage.PendingOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"createdBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxSlippageP\",\"type\":\"uint16\"}],\"internalType\":\"structITradingStorage.PendingOrder[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getAllTradeInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"createdBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"tpLastUpdatedBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"slLastUpdatedBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxSlippageP\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"lastOiUpdateTs\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"collateralPriceUsd\",\"type\":\"uint48\"},{\"internalType\":\"enumITradingStorage.ContractsVersion\",\"name\":\"contractsVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"lastPosIncreaseBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"__placeholder\",\"type\":\"uint8\"}],\"internalType\":\"structITradingStorage.TradeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getAllTrades\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"pairIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"leverage\",\"type\":\"uint24\"},{\"internalType\":\"bool\",\"name\":\"long\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"collateralIndex\",\"type\":\"uint8\"},{\"internalType\":\"enumITradingStorage.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"internalType\":\"uint120\",\"name\":\"collateralAmount\",\"type\":\"uint120\"},{\"internalType\":\"uint64\",\"name\":\"openPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"tp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sl\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"__placeholder\",\"type\":\"uint192\"}],\"internalType\":\"structITradingStorage.Trade[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getAllTradesLiquidationParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint40\",\"name\":\"maxLiqSpreadP\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"startLiqThresholdP\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"endLiqThresholdP\",\"type\":\"uint40\"},{\"internalType\":\"uint24\",\"name\":\"startLeverage\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"endLeverage\",\"type\":\"uint24\"}],\"internalType\":\"structIPairsStorage.GroupLiquidationParams[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_index\",\"type\":\"uint8\"}],\"name\":\"getCollateral\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint88\",\"name\":\"__placeholder\",\"type\":\"uint88\"},{\"internalType\":\"uint128\",\"name\":\"precision\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"precisionDelta\",\"type\":\"uint128\"}],\"internalType\":\"structITradingStorage.Collateral\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"}],\"name\":\"getCollateralIndex\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollaterals\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint88\",\"name\":\"__placeholder\",\"type\":\"uint88\"},{\"internalType\":\"uint128\",\"name\":\"precision\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"precisionDelta\",\"type\":\"uint128\"}],\"internalType\":\"structITradingStorage.Collateral[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollateralsCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"},{\"internalType\":\"enumITradingStorage.CounterType\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"getCounters\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"currentIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"openCount\",\"type\":\"uint32\"},{\"internalType\":\"uint192\",\"name\":\"__placeholder\",\"type\":\"uint192\"}],\"internalType\":\"structITradingStorage.Counter\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentContractsVersion\",\"outputs\":[{\"internalType\":\"enumITradingStorage.ContractsVersion\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_collateralIndex\",\"type\":\"uint8\"}],\"name\":\"getGToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structITradingStorage.Id\",\"name\":\"_orderId\",\"type\":\"tuple\"}],\"name\":\"getPendingOrder\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"pairIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"leverage\",\"type\":\"uint24\"},{\"internalType\":\"bool\",\"name\":\"long\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"collateralIndex\",\"type\":\"uint8\"},{\"internalType\":\"enumITradingStorage.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"internalType\":\"uint120\",\"name\":\"collateralAmount\",\"type\":\"uint120\"},{\"internalType\":\"uint64\",\"name\":\"openPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"tp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sl\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"__placeholder\",\"type\":\"uint192\"}],\"internalType\":\"structITradingStorage.Trade\",\"name\":\"trade\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"enumITradingStorage.PendingOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"createdBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxSlippageP\",\"type\":\"uint16\"}],\"internalType\":\"structITradingStorage.PendingOrder\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getPendingOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"pairIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"leverage\",\"type\":\"uint24\"},{\"internalType\":\"bool\",\"name\":\"long\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"collateralIndex\",\"type\":\"uint8\"},{\"internalType\":\"enumITradingStorage.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"internalType\":\"uint120\",\"name\":\"collateralAmount\",\"type\":\"uint120\"},{\"internalType\":\"uint64\",\"name\":\"openPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"tp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sl\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"__placeholder\",\"type\":\"uint192\"}],\"internalType\":\"structITradingStorage.Trade\",\"name\":\"trade\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"enumITradingStorage.PendingOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"createdBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxSlippageP\",\"type\":\"uint16\"}],\"internalType\":\"structITradingStorage.PendingOrder[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_index\",\"type\":\"uint32\"}],\"name\":\"getTrade\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"pairIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"leverage\",\"type\":\"uint24\"},{\"internalType\":\"bool\",\"name\":\"long\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"collateralIndex\",\"type\":\"uint8\"},{\"internalType\":\"enumITradingStorage.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"internalType\":\"uint120\",\"name\":\"collateralAmount\",\"type\":\"uint120\"},{\"internalType\":\"uint64\",\"name\":\"openPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"tp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sl\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"__placeholder\",\"type\":\"uint192\"}],\"internalType\":\"structITradingStorage.Trade\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_index\",\"type\":\"uint32\"}],\"name\":\"getTradeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"createdBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"tpLastUpdatedBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"slLastUpdatedBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxSlippageP\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"lastOiUpdateTs\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"collateralPriceUsd\",\"type\":\"uint48\"},{\"internalType\":\"enumITradingStorage.ContractsVersion\",\"name\":\"contractsVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"lastPosIncreaseBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"__placeholder\",\"type\":\"uint8\"}],\"internalType\":\"structITradingStorage.TradeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"}],\"name\":\"getTradeInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"createdBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"tpLastUpdatedBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"slLastUpdatedBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxSlippageP\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"lastOiUpdateTs\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"collateralPriceUsd\",\"type\":\"uint48\"},{\"internalType\":\"enumITradingStorage.ContractsVersion\",\"name\":\"contractsVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"lastPosIncreaseBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"__placeholder\",\"type\":\"uint8\"}],\"internalType\":\"structITradingStorage.TradeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_index\",\"type\":\"uint32\"}],\"name\":\"getTradeLiquidationParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint40\",\"name\":\"maxLiqSpreadP\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"startLiqThresholdP\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"endLiqThresholdP\",\"type\":\"uint40\"},{\"internalType\":\"uint24\",\"name\":\"startLeverage\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"endLeverage\",\"type\":\"uint24\"}],\"internalType\":\"structIPairsStorage.GroupLiquidationParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structITradingStorage.Id\",\"name\":\"_tradeId\",\"type\":\"tuple\"},{\"internalType\":\"enumITradingStorage.PendingOrderType\",\"name\":\"_orderType\",\"type\":\"uint8\"}],\"name\":\"getTradePendingOrderBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"}],\"name\":\"getTraderStored\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_offset\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_limit\",\"type\":\"uint32\"}],\"name\":\"getTraders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"}],\"name\":\"getTrades\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"pairIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"leverage\",\"type\":\"uint24\"},{\"internalType\":\"bool\",\"name\":\"long\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"collateralIndex\",\"type\":\"uint8\"},{\"internalType\":\"enumITradingStorage.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"internalType\":\"uint120\",\"name\":\"collateralAmount\",\"type\":\"uint120\"},{\"internalType\":\"uint64\",\"name\":\"openPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"tp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sl\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"__placeholder\",\"type\":\"uint192\"}],\"internalType\":\"structITradingStorage.Trade[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"}],\"name\":\"getTradesLiquidationParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint40\",\"name\":\"maxLiqSpreadP\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"startLiqThresholdP\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"endLiqThresholdP\",\"type\":\"uint40\"},{\"internalType\":\"uint24\",\"name\":\"startLeverage\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"endLeverage\",\"type\":\"uint24\"}],\"internalType\":\"structIPairsStorage.GroupLiquidationParams[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTradingActivated\",\"outputs\":[{\"internalType\":\"enumITradingStorage.TradingActivated\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"enumIAddressStore.Role\",\"name\":\"_role\",\"type\":\"uint8\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rolesManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gns\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gnsStaking\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_collaterals\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_gTokens\",\"type\":\"address[]\"}],\"name\":\"initializeTradingStorage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_index\",\"type\":\"uint8\"}],\"name\":\"isCollateralActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_index\",\"type\":\"uint8\"}],\"name\":\"isCollateralListed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"},{\"internalType\":\"enumIAddressStore.Role[]\",\"name\":\"_roles\",\"type\":\"uint8[]\"},{\"internalType\":\"bool[]\",\"name\":\"_values\",\"type\":\"bool[]\"}],\"name\":\"setRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"pairIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"leverage\",\"type\":\"uint24\"},{\"internalType\":\"bool\",\"name\":\"long\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"collateralIndex\",\"type\":\"uint8\"},{\"internalType\":\"enumITradingStorage.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"internalType\":\"uint120\",\"name\":\"collateralAmount\",\"type\":\"uint120\"},{\"internalType\":\"uint64\",\"name\":\"openPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"tp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sl\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"__placeholder\",\"type\":\"uint192\"}],\"internalType\":\"structITradingStorage.Trade\",\"name\":\"trade\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"enumITradingStorage.PendingOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"createdBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxSlippageP\",\"type\":\"uint16\"}],\"internalType\":\"structITradingStorage.PendingOrder\",\"name\":\"_pendingOrder\",\"type\":\"tuple\"}],\"name\":\"storePendingOrder\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"pairIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"leverage\",\"type\":\"uint24\"},{\"internalType\":\"bool\",\"name\":\"long\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"collateralIndex\",\"type\":\"uint8\"},{\"internalType\":\"enumITradingStorage.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"internalType\":\"uint120\",\"name\":\"collateralAmount\",\"type\":\"uint120\"},{\"internalType\":\"uint64\",\"name\":\"openPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"tp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sl\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"__placeholder\",\"type\":\"uint192\"}],\"internalType\":\"structITradingStorage.Trade\",\"name\":\"trade\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"enumITradingStorage.PendingOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"createdBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxSlippageP\",\"type\":\"uint16\"}],\"internalType\":\"structITradingStorage.PendingOrder\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"pairIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"leverage\",\"type\":\"uint24\"},{\"internalType\":\"bool\",\"name\":\"long\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"collateralIndex\",\"type\":\"uint8\"},{\"internalType\":\"enumITradingStorage.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"internalType\":\"uint120\",\"name\":\"collateralAmount\",\"type\":\"uint120\"},{\"internalType\":\"uint64\",\"name\":\"openPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"tp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sl\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"__placeholder\",\"type\":\"uint192\"}],\"internalType\":\"structITradingStorage.Trade\",\"name\":\"_trade\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"createdBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"tpLastUpdatedBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"slLastUpdatedBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxSlippageP\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"lastOiUpdateTs\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"collateralPriceUsd\",\"type\":\"uint48\"},{\"internalType\":\"enumITradingStorage.ContractsVersion\",\"name\":\"contractsVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"lastPosIncreaseBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"__placeholder\",\"type\":\"uint8\"}],\"internalType\":\"structITradingStorage.TradeInfo\",\"name\":\"_tradeInfo\",\"type\":\"tuple\"}],\"name\":\"storeTrade\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"pairIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"leverage\",\"type\":\"uint24\"},{\"internalType\":\"bool\",\"name\":\"long\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"collateralIndex\",\"type\":\"uint8\"},{\"internalType\":\"enumITradingStorage.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"},{\"internalType\":\"uint120\",\"name\":\"collateralAmount\",\"type\":\"uint120\"},{\"internalType\":\"uint64\",\"name\":\"openPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"tp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sl\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"__placeholder\",\"type\":\"uint192\"}],\"internalType\":\"structITradingStorage.Trade\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_collateralIndex\",\"type\":\"uint8\"}],\"name\":\"toggleCollateralActiveState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gToken\",\"type\":\"address\"}],\"name\":\"updateGToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structITradingStorage.Id\",\"name\":\"_tradeId\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_openPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_tp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_sl\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"_maxSlippageP\",\"type\":\"uint16\"}],\"name\":\"updateOpenOrderDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structITradingStorage.Id\",\"name\":\"_tradeId\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"_collateralAmount\",\"type\":\"uint120\"}],\"name\":\"updateTradeCollateralAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structITradingStorage.Id\",\"name\":\"_tradeId\",\"type\":\"tuple\"},{\"internalType\":\"uint16\",\"name\":\"_maxSlippageP\",\"type\":\"uint16\"}],\"name\":\"updateTradeMaxClosingSlippageP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structITradingStorage.Id\",\"name\":\"_tradeId\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"_collateralAmount\",\"type\":\"uint120\"},{\"internalType\":\"uint24\",\"name\":\"_leverage\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"_openPrice\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_isPartialIncrease\",\"type\":\"bool\"}],\"name\":\"updateTradePosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structITradingStorage.Id\",\"name\":\"_tradeId\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_newSl\",\"type\":\"uint64\"}],\"name\":\"updateTradeSl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structITradingStorage.Id\",\"name\":\"_tradeId\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_newTp\",\"type\":\"uint64\"}],\"name\":\"updateTradeTp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumITradingStorage.TradingActivated\",\"name\":\"_activated\",\"type\":\"uint8\"}],\"name\":\"updateTradingActivated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GainsAbisABI is the input ABI used to generate the binding from.
// Deprecated: Use GainsAbisMetaData.ABI instead.
var GainsAbisABI = GainsAbisMetaData.ABI

// GainsAbis is an auto generated Go binding around an Ethereum contract.
type GainsAbis struct {
	GainsAbisCaller     // Read-only binding to the contract
	GainsAbisTransactor // Write-only binding to the contract
	GainsAbisFilterer   // Log filterer for contract events
}

// GainsAbisCaller is an auto generated read-only Go binding around an Ethereum contract.
type GainsAbisCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GainsAbisTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GainsAbisTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GainsAbisFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GainsAbisFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GainsAbisSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GainsAbisSession struct {
	Contract     *GainsAbis        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GainsAbisCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GainsAbisCallerSession struct {
	Contract *GainsAbisCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// GainsAbisTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GainsAbisTransactorSession struct {
	Contract     *GainsAbisTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// GainsAbisRaw is an auto generated low-level Go binding around an Ethereum contract.
type GainsAbisRaw struct {
	Contract *GainsAbis // Generic contract binding to access the raw methods on
}

// GainsAbisCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GainsAbisCallerRaw struct {
	Contract *GainsAbisCaller // Generic read-only contract binding to access the raw methods on
}

// GainsAbisTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GainsAbisTransactorRaw struct {
	Contract *GainsAbisTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGainsAbis creates a new instance of GainsAbis, bound to a specific deployed contract.
func NewGainsAbis(address common.Address, backend bind.ContractBackend) (*GainsAbis, error) {
	contract, err := bindGainsAbis(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GainsAbis{GainsAbisCaller: GainsAbisCaller{contract: contract}, GainsAbisTransactor: GainsAbisTransactor{contract: contract}, GainsAbisFilterer: GainsAbisFilterer{contract: contract}}, nil
}

// NewGainsAbisCaller creates a new read-only instance of GainsAbis, bound to a specific deployed contract.
func NewGainsAbisCaller(address common.Address, caller bind.ContractCaller) (*GainsAbisCaller, error) {
	contract, err := bindGainsAbis(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GainsAbisCaller{contract: contract}, nil
}

// NewGainsAbisTransactor creates a new write-only instance of GainsAbis, bound to a specific deployed contract.
func NewGainsAbisTransactor(address common.Address, transactor bind.ContractTransactor) (*GainsAbisTransactor, error) {
	contract, err := bindGainsAbis(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GainsAbisTransactor{contract: contract}, nil
}

// NewGainsAbisFilterer creates a new log filterer instance of GainsAbis, bound to a specific deployed contract.
func NewGainsAbisFilterer(address common.Address, filterer bind.ContractFilterer) (*GainsAbisFilterer, error) {
	contract, err := bindGainsAbis(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GainsAbisFilterer{contract: contract}, nil
}

// bindGainsAbis binds a generic wrapper to an already deployed contract.
func bindGainsAbis(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GainsAbisMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GainsAbis *GainsAbisRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GainsAbis.Contract.GainsAbisCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GainsAbis *GainsAbisRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GainsAbis.Contract.GainsAbisTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GainsAbis *GainsAbisRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GainsAbis.Contract.GainsAbisTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GainsAbis *GainsAbisCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GainsAbis.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GainsAbis *GainsAbisTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GainsAbis.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GainsAbis *GainsAbisTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GainsAbis.Contract.contract.Transact(opts, method, params...)
}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns((address,address))
func (_GainsAbis *GainsAbisCaller) GetAddresses(opts *bind.CallOpts) (IAddressStoreAddresses, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getAddresses")

	if err != nil {
		return *new(IAddressStoreAddresses), err
	}

	out0 := *abi.ConvertType(out[0], new(IAddressStoreAddresses)).(*IAddressStoreAddresses)

	return out0, err

}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns((address,address))
func (_GainsAbis *GainsAbisSession) GetAddresses() (IAddressStoreAddresses, error) {
	return _GainsAbis.Contract.GetAddresses(&_GainsAbis.CallOpts)
}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns((address,address))
func (_GainsAbis *GainsAbisCallerSession) GetAddresses() (IAddressStoreAddresses, error) {
	return _GainsAbis.Contract.GetAddresses(&_GainsAbis.CallOpts)
}

// GetAllPendingOrders is a free data retrieval call binding the contract method 0x2d11445f.
//
// Solidity: function getAllPendingOrders(uint256 _offset, uint256 _limit) view returns(((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192),address,uint32,bool,uint8,uint32,uint16)[])
func (_GainsAbis *GainsAbisCaller) GetAllPendingOrders(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int) ([]ITradingStoragePendingOrder, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getAllPendingOrders", _offset, _limit)

	if err != nil {
		return *new([]ITradingStoragePendingOrder), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITradingStoragePendingOrder)).(*[]ITradingStoragePendingOrder)

	return out0, err

}

// GetAllPendingOrders is a free data retrieval call binding the contract method 0x2d11445f.
//
// Solidity: function getAllPendingOrders(uint256 _offset, uint256 _limit) view returns(((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192),address,uint32,bool,uint8,uint32,uint16)[])
func (_GainsAbis *GainsAbisSession) GetAllPendingOrders(_offset *big.Int, _limit *big.Int) ([]ITradingStoragePendingOrder, error) {
	return _GainsAbis.Contract.GetAllPendingOrders(&_GainsAbis.CallOpts, _offset, _limit)
}

// GetAllPendingOrders is a free data retrieval call binding the contract method 0x2d11445f.
//
// Solidity: function getAllPendingOrders(uint256 _offset, uint256 _limit) view returns(((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192),address,uint32,bool,uint8,uint32,uint16)[])
func (_GainsAbis *GainsAbisCallerSession) GetAllPendingOrders(_offset *big.Int, _limit *big.Int) ([]ITradingStoragePendingOrder, error) {
	return _GainsAbis.Contract.GetAllPendingOrders(&_GainsAbis.CallOpts, _offset, _limit)
}

// GetAllTradeInfos is a free data retrieval call binding the contract method 0xeb50287f.
//
// Solidity: function getAllTradeInfos(uint256 _offset, uint256 _limit) view returns((uint32,uint32,uint32,uint16,uint48,uint48,uint8,uint32,uint8)[])
func (_GainsAbis *GainsAbisCaller) GetAllTradeInfos(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int) ([]ITradingStorageTradeInfo, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getAllTradeInfos", _offset, _limit)

	if err != nil {
		return *new([]ITradingStorageTradeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITradingStorageTradeInfo)).(*[]ITradingStorageTradeInfo)

	return out0, err

}

// GetAllTradeInfos is a free data retrieval call binding the contract method 0xeb50287f.
//
// Solidity: function getAllTradeInfos(uint256 _offset, uint256 _limit) view returns((uint32,uint32,uint32,uint16,uint48,uint48,uint8,uint32,uint8)[])
func (_GainsAbis *GainsAbisSession) GetAllTradeInfos(_offset *big.Int, _limit *big.Int) ([]ITradingStorageTradeInfo, error) {
	return _GainsAbis.Contract.GetAllTradeInfos(&_GainsAbis.CallOpts, _offset, _limit)
}

// GetAllTradeInfos is a free data retrieval call binding the contract method 0xeb50287f.
//
// Solidity: function getAllTradeInfos(uint256 _offset, uint256 _limit) view returns((uint32,uint32,uint32,uint16,uint48,uint48,uint8,uint32,uint8)[])
func (_GainsAbis *GainsAbisCallerSession) GetAllTradeInfos(_offset *big.Int, _limit *big.Int) ([]ITradingStorageTradeInfo, error) {
	return _GainsAbis.Contract.GetAllTradeInfos(&_GainsAbis.CallOpts, _offset, _limit)
}

// GetAllTrades is a free data retrieval call binding the contract method 0xdffd8a1f.
//
// Solidity: function getAllTrades(uint256 _offset, uint256 _limit) view returns((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192)[])
func (_GainsAbis *GainsAbisCaller) GetAllTrades(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int) ([]ITradingStorageTrade, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getAllTrades", _offset, _limit)

	if err != nil {
		return *new([]ITradingStorageTrade), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITradingStorageTrade)).(*[]ITradingStorageTrade)

	return out0, err

}

// GetAllTrades is a free data retrieval call binding the contract method 0xdffd8a1f.
//
// Solidity: function getAllTrades(uint256 _offset, uint256 _limit) view returns((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192)[])
func (_GainsAbis *GainsAbisSession) GetAllTrades(_offset *big.Int, _limit *big.Int) ([]ITradingStorageTrade, error) {
	return _GainsAbis.Contract.GetAllTrades(&_GainsAbis.CallOpts, _offset, _limit)
}

// GetAllTrades is a free data retrieval call binding the contract method 0xdffd8a1f.
//
// Solidity: function getAllTrades(uint256 _offset, uint256 _limit) view returns((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192)[])
func (_GainsAbis *GainsAbisCallerSession) GetAllTrades(_offset *big.Int, _limit *big.Int) ([]ITradingStorageTrade, error) {
	return _GainsAbis.Contract.GetAllTrades(&_GainsAbis.CallOpts, _offset, _limit)
}

// GetAllTradesLiquidationParams is a free data retrieval call binding the contract method 0xc2b96e65.
//
// Solidity: function getAllTradesLiquidationParams(uint256 _offset, uint256 _limit) view returns((uint40,uint40,uint40,uint24,uint24)[])
func (_GainsAbis *GainsAbisCaller) GetAllTradesLiquidationParams(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int) ([]IPairsStorageGroupLiquidationParams, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getAllTradesLiquidationParams", _offset, _limit)

	if err != nil {
		return *new([]IPairsStorageGroupLiquidationParams), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPairsStorageGroupLiquidationParams)).(*[]IPairsStorageGroupLiquidationParams)

	return out0, err

}

// GetAllTradesLiquidationParams is a free data retrieval call binding the contract method 0xc2b96e65.
//
// Solidity: function getAllTradesLiquidationParams(uint256 _offset, uint256 _limit) view returns((uint40,uint40,uint40,uint24,uint24)[])
func (_GainsAbis *GainsAbisSession) GetAllTradesLiquidationParams(_offset *big.Int, _limit *big.Int) ([]IPairsStorageGroupLiquidationParams, error) {
	return _GainsAbis.Contract.GetAllTradesLiquidationParams(&_GainsAbis.CallOpts, _offset, _limit)
}

// GetAllTradesLiquidationParams is a free data retrieval call binding the contract method 0xc2b96e65.
//
// Solidity: function getAllTradesLiquidationParams(uint256 _offset, uint256 _limit) view returns((uint40,uint40,uint40,uint24,uint24)[])
func (_GainsAbis *GainsAbisCallerSession) GetAllTradesLiquidationParams(_offset *big.Int, _limit *big.Int) ([]IPairsStorageGroupLiquidationParams, error) {
	return _GainsAbis.Contract.GetAllTradesLiquidationParams(&_GainsAbis.CallOpts, _offset, _limit)
}

// GetCollateral is a free data retrieval call binding the contract method 0xbb33a55b.
//
// Solidity: function getCollateral(uint8 _index) view returns((address,bool,uint88,uint128,uint128))
func (_GainsAbis *GainsAbisCaller) GetCollateral(opts *bind.CallOpts, _index uint8) (ITradingStorageCollateral, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getCollateral", _index)

	if err != nil {
		return *new(ITradingStorageCollateral), err
	}

	out0 := *abi.ConvertType(out[0], new(ITradingStorageCollateral)).(*ITradingStorageCollateral)

	return out0, err

}

// GetCollateral is a free data retrieval call binding the contract method 0xbb33a55b.
//
// Solidity: function getCollateral(uint8 _index) view returns((address,bool,uint88,uint128,uint128))
func (_GainsAbis *GainsAbisSession) GetCollateral(_index uint8) (ITradingStorageCollateral, error) {
	return _GainsAbis.Contract.GetCollateral(&_GainsAbis.CallOpts, _index)
}

// GetCollateral is a free data retrieval call binding the contract method 0xbb33a55b.
//
// Solidity: function getCollateral(uint8 _index) view returns((address,bool,uint88,uint128,uint128))
func (_GainsAbis *GainsAbisCallerSession) GetCollateral(_index uint8) (ITradingStorageCollateral, error) {
	return _GainsAbis.Contract.GetCollateral(&_GainsAbis.CallOpts, _index)
}

// GetCollateralIndex is a free data retrieval call binding the contract method 0x5c3ed7c3.
//
// Solidity: function getCollateralIndex(address _collateral) view returns(uint8)
func (_GainsAbis *GainsAbisCaller) GetCollateralIndex(opts *bind.CallOpts, _collateral common.Address) (uint8, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getCollateralIndex", _collateral)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetCollateralIndex is a free data retrieval call binding the contract method 0x5c3ed7c3.
//
// Solidity: function getCollateralIndex(address _collateral) view returns(uint8)
func (_GainsAbis *GainsAbisSession) GetCollateralIndex(_collateral common.Address) (uint8, error) {
	return _GainsAbis.Contract.GetCollateralIndex(&_GainsAbis.CallOpts, _collateral)
}

// GetCollateralIndex is a free data retrieval call binding the contract method 0x5c3ed7c3.
//
// Solidity: function getCollateralIndex(address _collateral) view returns(uint8)
func (_GainsAbis *GainsAbisCallerSession) GetCollateralIndex(_collateral common.Address) (uint8, error) {
	return _GainsAbis.Contract.GetCollateralIndex(&_GainsAbis.CallOpts, _collateral)
}

// GetCollaterals is a free data retrieval call binding the contract method 0x78b92636.
//
// Solidity: function getCollaterals() view returns((address,bool,uint88,uint128,uint128)[])
func (_GainsAbis *GainsAbisCaller) GetCollaterals(opts *bind.CallOpts) ([]ITradingStorageCollateral, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getCollaterals")

	if err != nil {
		return *new([]ITradingStorageCollateral), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITradingStorageCollateral)).(*[]ITradingStorageCollateral)

	return out0, err

}

// GetCollaterals is a free data retrieval call binding the contract method 0x78b92636.
//
// Solidity: function getCollaterals() view returns((address,bool,uint88,uint128,uint128)[])
func (_GainsAbis *GainsAbisSession) GetCollaterals() ([]ITradingStorageCollateral, error) {
	return _GainsAbis.Contract.GetCollaterals(&_GainsAbis.CallOpts)
}

// GetCollaterals is a free data retrieval call binding the contract method 0x78b92636.
//
// Solidity: function getCollaterals() view returns((address,bool,uint88,uint128,uint128)[])
func (_GainsAbis *GainsAbisCallerSession) GetCollaterals() ([]ITradingStorageCollateral, error) {
	return _GainsAbis.Contract.GetCollaterals(&_GainsAbis.CallOpts)
}

// GetCollateralsCount is a free data retrieval call binding the contract method 0xa3e15d09.
//
// Solidity: function getCollateralsCount() view returns(uint8)
func (_GainsAbis *GainsAbisCaller) GetCollateralsCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getCollateralsCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetCollateralsCount is a free data retrieval call binding the contract method 0xa3e15d09.
//
// Solidity: function getCollateralsCount() view returns(uint8)
func (_GainsAbis *GainsAbisSession) GetCollateralsCount() (uint8, error) {
	return _GainsAbis.Contract.GetCollateralsCount(&_GainsAbis.CallOpts)
}

// GetCollateralsCount is a free data retrieval call binding the contract method 0xa3e15d09.
//
// Solidity: function getCollateralsCount() view returns(uint8)
func (_GainsAbis *GainsAbisCallerSession) GetCollateralsCount() (uint8, error) {
	return _GainsAbis.Contract.GetCollateralsCount(&_GainsAbis.CallOpts)
}

// GetCounters is a free data retrieval call binding the contract method 0x0212f0d6.
//
// Solidity: function getCounters(address _trader, uint8 _type) view returns((uint32,uint32,uint192))
func (_GainsAbis *GainsAbisCaller) GetCounters(opts *bind.CallOpts, _trader common.Address, _type uint8) (ITradingStorageCounter, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getCounters", _trader, _type)

	if err != nil {
		return *new(ITradingStorageCounter), err
	}

	out0 := *abi.ConvertType(out[0], new(ITradingStorageCounter)).(*ITradingStorageCounter)

	return out0, err

}

// GetCounters is a free data retrieval call binding the contract method 0x0212f0d6.
//
// Solidity: function getCounters(address _trader, uint8 _type) view returns((uint32,uint32,uint192))
func (_GainsAbis *GainsAbisSession) GetCounters(_trader common.Address, _type uint8) (ITradingStorageCounter, error) {
	return _GainsAbis.Contract.GetCounters(&_GainsAbis.CallOpts, _trader, _type)
}

// GetCounters is a free data retrieval call binding the contract method 0x0212f0d6.
//
// Solidity: function getCounters(address _trader, uint8 _type) view returns((uint32,uint32,uint192))
func (_GainsAbis *GainsAbisCallerSession) GetCounters(_trader common.Address, _type uint8) (ITradingStorageCounter, error) {
	return _GainsAbis.Contract.GetCounters(&_GainsAbis.CallOpts, _trader, _type)
}

// GetCurrentContractsVersion is a free data retrieval call binding the contract method 0x9095b119.
//
// Solidity: function getCurrentContractsVersion() pure returns(uint8)
func (_GainsAbis *GainsAbisCaller) GetCurrentContractsVersion(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getCurrentContractsVersion")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetCurrentContractsVersion is a free data retrieval call binding the contract method 0x9095b119.
//
// Solidity: function getCurrentContractsVersion() pure returns(uint8)
func (_GainsAbis *GainsAbisSession) GetCurrentContractsVersion() (uint8, error) {
	return _GainsAbis.Contract.GetCurrentContractsVersion(&_GainsAbis.CallOpts)
}

// GetCurrentContractsVersion is a free data retrieval call binding the contract method 0x9095b119.
//
// Solidity: function getCurrentContractsVersion() pure returns(uint8)
func (_GainsAbis *GainsAbisCallerSession) GetCurrentContractsVersion() (uint8, error) {
	return _GainsAbis.Contract.GetCurrentContractsVersion(&_GainsAbis.CallOpts)
}

// GetGToken is a free data retrieval call binding the contract method 0x6a0aff41.
//
// Solidity: function getGToken(uint8 _collateralIndex) view returns(address)
func (_GainsAbis *GainsAbisCaller) GetGToken(opts *bind.CallOpts, _collateralIndex uint8) (common.Address, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getGToken", _collateralIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGToken is a free data retrieval call binding the contract method 0x6a0aff41.
//
// Solidity: function getGToken(uint8 _collateralIndex) view returns(address)
func (_GainsAbis *GainsAbisSession) GetGToken(_collateralIndex uint8) (common.Address, error) {
	return _GainsAbis.Contract.GetGToken(&_GainsAbis.CallOpts, _collateralIndex)
}

// GetGToken is a free data retrieval call binding the contract method 0x6a0aff41.
//
// Solidity: function getGToken(uint8 _collateralIndex) view returns(address)
func (_GainsAbis *GainsAbisCallerSession) GetGToken(_collateralIndex uint8) (common.Address, error) {
	return _GainsAbis.Contract.GetGToken(&_GainsAbis.CallOpts, _collateralIndex)
}

// GetPendingOrder is a free data retrieval call binding the contract method 0xc6e729bb.
//
// Solidity: function getPendingOrder((address,uint32) _orderId) view returns(((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192),address,uint32,bool,uint8,uint32,uint16))
func (_GainsAbis *GainsAbisCaller) GetPendingOrder(opts *bind.CallOpts, _orderId ITradingStorageId) (ITradingStoragePendingOrder, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getPendingOrder", _orderId)

	if err != nil {
		return *new(ITradingStoragePendingOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(ITradingStoragePendingOrder)).(*ITradingStoragePendingOrder)

	return out0, err

}

// GetPendingOrder is a free data retrieval call binding the contract method 0xc6e729bb.
//
// Solidity: function getPendingOrder((address,uint32) _orderId) view returns(((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192),address,uint32,bool,uint8,uint32,uint16))
func (_GainsAbis *GainsAbisSession) GetPendingOrder(_orderId ITradingStorageId) (ITradingStoragePendingOrder, error) {
	return _GainsAbis.Contract.GetPendingOrder(&_GainsAbis.CallOpts, _orderId)
}

// GetPendingOrder is a free data retrieval call binding the contract method 0xc6e729bb.
//
// Solidity: function getPendingOrder((address,uint32) _orderId) view returns(((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192),address,uint32,bool,uint8,uint32,uint16))
func (_GainsAbis *GainsAbisCallerSession) GetPendingOrder(_orderId ITradingStorageId) (ITradingStoragePendingOrder, error) {
	return _GainsAbis.Contract.GetPendingOrder(&_GainsAbis.CallOpts, _orderId)
}

// GetPendingOrders is a free data retrieval call binding the contract method 0x4c73cb25.
//
// Solidity: function getPendingOrders(address _user) view returns(((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192),address,uint32,bool,uint8,uint32,uint16)[])
func (_GainsAbis *GainsAbisCaller) GetPendingOrders(opts *bind.CallOpts, _user common.Address) ([]ITradingStoragePendingOrder, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getPendingOrders", _user)

	if err != nil {
		return *new([]ITradingStoragePendingOrder), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITradingStoragePendingOrder)).(*[]ITradingStoragePendingOrder)

	return out0, err

}

// GetPendingOrders is a free data retrieval call binding the contract method 0x4c73cb25.
//
// Solidity: function getPendingOrders(address _user) view returns(((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192),address,uint32,bool,uint8,uint32,uint16)[])
func (_GainsAbis *GainsAbisSession) GetPendingOrders(_user common.Address) ([]ITradingStoragePendingOrder, error) {
	return _GainsAbis.Contract.GetPendingOrders(&_GainsAbis.CallOpts, _user)
}

// GetPendingOrders is a free data retrieval call binding the contract method 0x4c73cb25.
//
// Solidity: function getPendingOrders(address _user) view returns(((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192),address,uint32,bool,uint8,uint32,uint16)[])
func (_GainsAbis *GainsAbisCallerSession) GetPendingOrders(_user common.Address) ([]ITradingStoragePendingOrder, error) {
	return _GainsAbis.Contract.GetPendingOrders(&_GainsAbis.CallOpts, _user)
}

// GetTrade is a free data retrieval call binding the contract method 0x15878e07.
//
// Solidity: function getTrade(address _trader, uint32 _index) view returns((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192))
func (_GainsAbis *GainsAbisCaller) GetTrade(opts *bind.CallOpts, _trader common.Address, _index uint32) (ITradingStorageTrade, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getTrade", _trader, _index)

	if err != nil {
		return *new(ITradingStorageTrade), err
	}

	out0 := *abi.ConvertType(out[0], new(ITradingStorageTrade)).(*ITradingStorageTrade)

	return out0, err

}

// GetTrade is a free data retrieval call binding the contract method 0x15878e07.
//
// Solidity: function getTrade(address _trader, uint32 _index) view returns((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192))
func (_GainsAbis *GainsAbisSession) GetTrade(_trader common.Address, _index uint32) (ITradingStorageTrade, error) {
	return _GainsAbis.Contract.GetTrade(&_GainsAbis.CallOpts, _trader, _index)
}

// GetTrade is a free data retrieval call binding the contract method 0x15878e07.
//
// Solidity: function getTrade(address _trader, uint32 _index) view returns((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192))
func (_GainsAbis *GainsAbisCallerSession) GetTrade(_trader common.Address, _index uint32) (ITradingStorageTrade, error) {
	return _GainsAbis.Contract.GetTrade(&_GainsAbis.CallOpts, _trader, _index)
}

// GetTradeInfo is a free data retrieval call binding the contract method 0x75cd812d.
//
// Solidity: function getTradeInfo(address _trader, uint32 _index) view returns((uint32,uint32,uint32,uint16,uint48,uint48,uint8,uint32,uint8))
func (_GainsAbis *GainsAbisCaller) GetTradeInfo(opts *bind.CallOpts, _trader common.Address, _index uint32) (ITradingStorageTradeInfo, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getTradeInfo", _trader, _index)

	if err != nil {
		return *new(ITradingStorageTradeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ITradingStorageTradeInfo)).(*ITradingStorageTradeInfo)

	return out0, err

}

// GetTradeInfo is a free data retrieval call binding the contract method 0x75cd812d.
//
// Solidity: function getTradeInfo(address _trader, uint32 _index) view returns((uint32,uint32,uint32,uint16,uint48,uint48,uint8,uint32,uint8))
func (_GainsAbis *GainsAbisSession) GetTradeInfo(_trader common.Address, _index uint32) (ITradingStorageTradeInfo, error) {
	return _GainsAbis.Contract.GetTradeInfo(&_GainsAbis.CallOpts, _trader, _index)
}

// GetTradeInfo is a free data retrieval call binding the contract method 0x75cd812d.
//
// Solidity: function getTradeInfo(address _trader, uint32 _index) view returns((uint32,uint32,uint32,uint16,uint48,uint48,uint8,uint32,uint8))
func (_GainsAbis *GainsAbisCallerSession) GetTradeInfo(_trader common.Address, _index uint32) (ITradingStorageTradeInfo, error) {
	return _GainsAbis.Contract.GetTradeInfo(&_GainsAbis.CallOpts, _trader, _index)
}

// GetTradeInfos is a free data retrieval call binding the contract method 0x0d1e3c94.
//
// Solidity: function getTradeInfos(address _trader) view returns((uint32,uint32,uint32,uint16,uint48,uint48,uint8,uint32,uint8)[])
func (_GainsAbis *GainsAbisCaller) GetTradeInfos(opts *bind.CallOpts, _trader common.Address) ([]ITradingStorageTradeInfo, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getTradeInfos", _trader)

	if err != nil {
		return *new([]ITradingStorageTradeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITradingStorageTradeInfo)).(*[]ITradingStorageTradeInfo)

	return out0, err

}

// GetTradeInfos is a free data retrieval call binding the contract method 0x0d1e3c94.
//
// Solidity: function getTradeInfos(address _trader) view returns((uint32,uint32,uint32,uint16,uint48,uint48,uint8,uint32,uint8)[])
func (_GainsAbis *GainsAbisSession) GetTradeInfos(_trader common.Address) ([]ITradingStorageTradeInfo, error) {
	return _GainsAbis.Contract.GetTradeInfos(&_GainsAbis.CallOpts, _trader)
}

// GetTradeInfos is a free data retrieval call binding the contract method 0x0d1e3c94.
//
// Solidity: function getTradeInfos(address _trader) view returns((uint32,uint32,uint32,uint16,uint48,uint48,uint8,uint32,uint8)[])
func (_GainsAbis *GainsAbisCallerSession) GetTradeInfos(_trader common.Address) ([]ITradingStorageTradeInfo, error) {
	return _GainsAbis.Contract.GetTradeInfos(&_GainsAbis.CallOpts, _trader)
}

// GetTradeLiquidationParams is a free data retrieval call binding the contract method 0x28dc892f.
//
// Solidity: function getTradeLiquidationParams(address _trader, uint32 _index) view returns((uint40,uint40,uint40,uint24,uint24))
func (_GainsAbis *GainsAbisCaller) GetTradeLiquidationParams(opts *bind.CallOpts, _trader common.Address, _index uint32) (IPairsStorageGroupLiquidationParams, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getTradeLiquidationParams", _trader, _index)

	if err != nil {
		return *new(IPairsStorageGroupLiquidationParams), err
	}

	out0 := *abi.ConvertType(out[0], new(IPairsStorageGroupLiquidationParams)).(*IPairsStorageGroupLiquidationParams)

	return out0, err

}

// GetTradeLiquidationParams is a free data retrieval call binding the contract method 0x28dc892f.
//
// Solidity: function getTradeLiquidationParams(address _trader, uint32 _index) view returns((uint40,uint40,uint40,uint24,uint24))
func (_GainsAbis *GainsAbisSession) GetTradeLiquidationParams(_trader common.Address, _index uint32) (IPairsStorageGroupLiquidationParams, error) {
	return _GainsAbis.Contract.GetTradeLiquidationParams(&_GainsAbis.CallOpts, _trader, _index)
}

// GetTradeLiquidationParams is a free data retrieval call binding the contract method 0x28dc892f.
//
// Solidity: function getTradeLiquidationParams(address _trader, uint32 _index) view returns((uint40,uint40,uint40,uint24,uint24))
func (_GainsAbis *GainsAbisCallerSession) GetTradeLiquidationParams(_trader common.Address, _index uint32) (IPairsStorageGroupLiquidationParams, error) {
	return _GainsAbis.Contract.GetTradeLiquidationParams(&_GainsAbis.CallOpts, _trader, _index)
}

// GetTradePendingOrderBlock is a free data retrieval call binding the contract method 0x067e84dd.
//
// Solidity: function getTradePendingOrderBlock((address,uint32) _tradeId, uint8 _orderType) view returns(uint256)
func (_GainsAbis *GainsAbisCaller) GetTradePendingOrderBlock(opts *bind.CallOpts, _tradeId ITradingStorageId, _orderType uint8) (*big.Int, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getTradePendingOrderBlock", _tradeId, _orderType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTradePendingOrderBlock is a free data retrieval call binding the contract method 0x067e84dd.
//
// Solidity: function getTradePendingOrderBlock((address,uint32) _tradeId, uint8 _orderType) view returns(uint256)
func (_GainsAbis *GainsAbisSession) GetTradePendingOrderBlock(_tradeId ITradingStorageId, _orderType uint8) (*big.Int, error) {
	return _GainsAbis.Contract.GetTradePendingOrderBlock(&_GainsAbis.CallOpts, _tradeId, _orderType)
}

// GetTradePendingOrderBlock is a free data retrieval call binding the contract method 0x067e84dd.
//
// Solidity: function getTradePendingOrderBlock((address,uint32) _tradeId, uint8 _orderType) view returns(uint256)
func (_GainsAbis *GainsAbisCallerSession) GetTradePendingOrderBlock(_tradeId ITradingStorageId, _orderType uint8) (*big.Int, error) {
	return _GainsAbis.Contract.GetTradePendingOrderBlock(&_GainsAbis.CallOpts, _tradeId, _orderType)
}

// GetTraderStored is a free data retrieval call binding the contract method 0xbed8d2da.
//
// Solidity: function getTraderStored(address _trader) view returns(bool)
func (_GainsAbis *GainsAbisCaller) GetTraderStored(opts *bind.CallOpts, _trader common.Address) (bool, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getTraderStored", _trader)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetTraderStored is a free data retrieval call binding the contract method 0xbed8d2da.
//
// Solidity: function getTraderStored(address _trader) view returns(bool)
func (_GainsAbis *GainsAbisSession) GetTraderStored(_trader common.Address) (bool, error) {
	return _GainsAbis.Contract.GetTraderStored(&_GainsAbis.CallOpts, _trader)
}

// GetTraderStored is a free data retrieval call binding the contract method 0xbed8d2da.
//
// Solidity: function getTraderStored(address _trader) view returns(bool)
func (_GainsAbis *GainsAbisCallerSession) GetTraderStored(_trader common.Address) (bool, error) {
	return _GainsAbis.Contract.GetTraderStored(&_GainsAbis.CallOpts, _trader)
}

// GetTraders is a free data retrieval call binding the contract method 0x0e503724.
//
// Solidity: function getTraders(uint32 _offset, uint32 _limit) view returns(address[])
func (_GainsAbis *GainsAbisCaller) GetTraders(opts *bind.CallOpts, _offset uint32, _limit uint32) ([]common.Address, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getTraders", _offset, _limit)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTraders is a free data retrieval call binding the contract method 0x0e503724.
//
// Solidity: function getTraders(uint32 _offset, uint32 _limit) view returns(address[])
func (_GainsAbis *GainsAbisSession) GetTraders(_offset uint32, _limit uint32) ([]common.Address, error) {
	return _GainsAbis.Contract.GetTraders(&_GainsAbis.CallOpts, _offset, _limit)
}

// GetTraders is a free data retrieval call binding the contract method 0x0e503724.
//
// Solidity: function getTraders(uint32 _offset, uint32 _limit) view returns(address[])
func (_GainsAbis *GainsAbisCallerSession) GetTraders(_offset uint32, _limit uint32) ([]common.Address, error) {
	return _GainsAbis.Contract.GetTraders(&_GainsAbis.CallOpts, _offset, _limit)
}

// GetTrades is a free data retrieval call binding the contract method 0x4bfad7c0.
//
// Solidity: function getTrades(address _trader) view returns((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192)[])
func (_GainsAbis *GainsAbisCaller) GetTrades(opts *bind.CallOpts, _trader common.Address) ([]ITradingStorageTrade, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getTrades", _trader)

	if err != nil {
		return *new([]ITradingStorageTrade), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITradingStorageTrade)).(*[]ITradingStorageTrade)

	return out0, err

}

// GetTrades is a free data retrieval call binding the contract method 0x4bfad7c0.
//
// Solidity: function getTrades(address _trader) view returns((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192)[])
func (_GainsAbis *GainsAbisSession) GetTrades(_trader common.Address) ([]ITradingStorageTrade, error) {
	return _GainsAbis.Contract.GetTrades(&_GainsAbis.CallOpts, _trader)
}

// GetTrades is a free data retrieval call binding the contract method 0x4bfad7c0.
//
// Solidity: function getTrades(address _trader) view returns((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192)[])
func (_GainsAbis *GainsAbisCallerSession) GetTrades(_trader common.Address) ([]ITradingStorageTrade, error) {
	return _GainsAbis.Contract.GetTrades(&_GainsAbis.CallOpts, _trader)
}

// GetTradesLiquidationParams is a free data retrieval call binding the contract method 0xf7746f3c.
//
// Solidity: function getTradesLiquidationParams(address _trader) view returns((uint40,uint40,uint40,uint24,uint24)[])
func (_GainsAbis *GainsAbisCaller) GetTradesLiquidationParams(opts *bind.CallOpts, _trader common.Address) ([]IPairsStorageGroupLiquidationParams, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getTradesLiquidationParams", _trader)

	if err != nil {
		return *new([]IPairsStorageGroupLiquidationParams), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPairsStorageGroupLiquidationParams)).(*[]IPairsStorageGroupLiquidationParams)

	return out0, err

}

// GetTradesLiquidationParams is a free data retrieval call binding the contract method 0xf7746f3c.
//
// Solidity: function getTradesLiquidationParams(address _trader) view returns((uint40,uint40,uint40,uint24,uint24)[])
func (_GainsAbis *GainsAbisSession) GetTradesLiquidationParams(_trader common.Address) ([]IPairsStorageGroupLiquidationParams, error) {
	return _GainsAbis.Contract.GetTradesLiquidationParams(&_GainsAbis.CallOpts, _trader)
}

// GetTradesLiquidationParams is a free data retrieval call binding the contract method 0xf7746f3c.
//
// Solidity: function getTradesLiquidationParams(address _trader) view returns((uint40,uint40,uint40,uint24,uint24)[])
func (_GainsAbis *GainsAbisCallerSession) GetTradesLiquidationParams(_trader common.Address) ([]IPairsStorageGroupLiquidationParams, error) {
	return _GainsAbis.Contract.GetTradesLiquidationParams(&_GainsAbis.CallOpts, _trader)
}

// GetTradingActivated is a free data retrieval call binding the contract method 0x4115c122.
//
// Solidity: function getTradingActivated() view returns(uint8)
func (_GainsAbis *GainsAbisCaller) GetTradingActivated(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "getTradingActivated")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTradingActivated is a free data retrieval call binding the contract method 0x4115c122.
//
// Solidity: function getTradingActivated() view returns(uint8)
func (_GainsAbis *GainsAbisSession) GetTradingActivated() (uint8, error) {
	return _GainsAbis.Contract.GetTradingActivated(&_GainsAbis.CallOpts)
}

// GetTradingActivated is a free data retrieval call binding the contract method 0x4115c122.
//
// Solidity: function getTradingActivated() view returns(uint8)
func (_GainsAbis *GainsAbisCallerSession) GetTradingActivated() (uint8, error) {
	return _GainsAbis.Contract.GetTradingActivated(&_GainsAbis.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x95a8c58d.
//
// Solidity: function hasRole(address _account, uint8 _role) view returns(bool)
func (_GainsAbis *GainsAbisCaller) HasRole(opts *bind.CallOpts, _account common.Address, _role uint8) (bool, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "hasRole", _account, _role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x95a8c58d.
//
// Solidity: function hasRole(address _account, uint8 _role) view returns(bool)
func (_GainsAbis *GainsAbisSession) HasRole(_account common.Address, _role uint8) (bool, error) {
	return _GainsAbis.Contract.HasRole(&_GainsAbis.CallOpts, _account, _role)
}

// HasRole is a free data retrieval call binding the contract method 0x95a8c58d.
//
// Solidity: function hasRole(address _account, uint8 _role) view returns(bool)
func (_GainsAbis *GainsAbisCallerSession) HasRole(_account common.Address, _role uint8) (bool, error) {
	return _GainsAbis.Contract.HasRole(&_GainsAbis.CallOpts, _account, _role)
}

// IsCollateralActive is a free data retrieval call binding the contract method 0x4d140218.
//
// Solidity: function isCollateralActive(uint8 _index) view returns(bool)
func (_GainsAbis *GainsAbisCaller) IsCollateralActive(opts *bind.CallOpts, _index uint8) (bool, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "isCollateralActive", _index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCollateralActive is a free data retrieval call binding the contract method 0x4d140218.
//
// Solidity: function isCollateralActive(uint8 _index) view returns(bool)
func (_GainsAbis *GainsAbisSession) IsCollateralActive(_index uint8) (bool, error) {
	return _GainsAbis.Contract.IsCollateralActive(&_GainsAbis.CallOpts, _index)
}

// IsCollateralActive is a free data retrieval call binding the contract method 0x4d140218.
//
// Solidity: function isCollateralActive(uint8 _index) view returns(bool)
func (_GainsAbis *GainsAbisCallerSession) IsCollateralActive(_index uint8) (bool, error) {
	return _GainsAbis.Contract.IsCollateralActive(&_GainsAbis.CallOpts, _index)
}

// IsCollateralListed is a free data retrieval call binding the contract method 0x1d2ffb42.
//
// Solidity: function isCollateralListed(uint8 _index) view returns(bool)
func (_GainsAbis *GainsAbisCaller) IsCollateralListed(opts *bind.CallOpts, _index uint8) (bool, error) {
	var out []interface{}
	err := _GainsAbis.contract.Call(opts, &out, "isCollateralListed", _index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCollateralListed is a free data retrieval call binding the contract method 0x1d2ffb42.
//
// Solidity: function isCollateralListed(uint8 _index) view returns(bool)
func (_GainsAbis *GainsAbisSession) IsCollateralListed(_index uint8) (bool, error) {
	return _GainsAbis.Contract.IsCollateralListed(&_GainsAbis.CallOpts, _index)
}

// IsCollateralListed is a free data retrieval call binding the contract method 0x1d2ffb42.
//
// Solidity: function isCollateralListed(uint8 _index) view returns(bool)
func (_GainsAbis *GainsAbisCallerSession) IsCollateralListed(_index uint8) (bool, error) {
	return _GainsAbis.Contract.IsCollateralListed(&_GainsAbis.CallOpts, _index)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xc6783af1.
//
// Solidity: function addCollateral(address _collateral, address _gToken) returns()
func (_GainsAbis *GainsAbisTransactor) AddCollateral(opts *bind.TransactOpts, _collateral common.Address, _gToken common.Address) (*types.Transaction, error) {
	return _GainsAbis.contract.Transact(opts, "addCollateral", _collateral, _gToken)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xc6783af1.
//
// Solidity: function addCollateral(address _collateral, address _gToken) returns()
func (_GainsAbis *GainsAbisSession) AddCollateral(_collateral common.Address, _gToken common.Address) (*types.Transaction, error) {
	return _GainsAbis.Contract.AddCollateral(&_GainsAbis.TransactOpts, _collateral, _gToken)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xc6783af1.
//
// Solidity: function addCollateral(address _collateral, address _gToken) returns()
func (_GainsAbis *GainsAbisTransactorSession) AddCollateral(_collateral common.Address, _gToken common.Address) (*types.Transaction, error) {
	return _GainsAbis.Contract.AddCollateral(&_GainsAbis.TransactOpts, _collateral, _gToken)
}

// ClosePendingOrder is a paid mutator transaction binding the contract method 0x4fb70bba.
//
// Solidity: function closePendingOrder((address,uint32) _orderId) returns()
func (_GainsAbis *GainsAbisTransactor) ClosePendingOrder(opts *bind.TransactOpts, _orderId ITradingStorageId) (*types.Transaction, error) {
	return _GainsAbis.contract.Transact(opts, "closePendingOrder", _orderId)
}

// ClosePendingOrder is a paid mutator transaction binding the contract method 0x4fb70bba.
//
// Solidity: function closePendingOrder((address,uint32) _orderId) returns()
func (_GainsAbis *GainsAbisSession) ClosePendingOrder(_orderId ITradingStorageId) (*types.Transaction, error) {
	return _GainsAbis.Contract.ClosePendingOrder(&_GainsAbis.TransactOpts, _orderId)
}

// ClosePendingOrder is a paid mutator transaction binding the contract method 0x4fb70bba.
//
// Solidity: function closePendingOrder((address,uint32) _orderId) returns()
func (_GainsAbis *GainsAbisTransactorSession) ClosePendingOrder(_orderId ITradingStorageId) (*types.Transaction, error) {
	return _GainsAbis.Contract.ClosePendingOrder(&_GainsAbis.TransactOpts, _orderId)
}

// CloseTrade is a paid mutator transaction binding the contract method 0x8583909b.
//
// Solidity: function closeTrade((address,uint32) _tradeId) returns()
func (_GainsAbis *GainsAbisTransactor) CloseTrade(opts *bind.TransactOpts, _tradeId ITradingStorageId) (*types.Transaction, error) {
	return _GainsAbis.contract.Transact(opts, "closeTrade", _tradeId)
}

// CloseTrade is a paid mutator transaction binding the contract method 0x8583909b.
//
// Solidity: function closeTrade((address,uint32) _tradeId) returns()
func (_GainsAbis *GainsAbisSession) CloseTrade(_tradeId ITradingStorageId) (*types.Transaction, error) {
	return _GainsAbis.Contract.CloseTrade(&_GainsAbis.TransactOpts, _tradeId)
}

// CloseTrade is a paid mutator transaction binding the contract method 0x8583909b.
//
// Solidity: function closeTrade((address,uint32) _tradeId) returns()
func (_GainsAbis *GainsAbisTransactorSession) CloseTrade(_tradeId ITradingStorageId) (*types.Transaction, error) {
	return _GainsAbis.Contract.CloseTrade(&_GainsAbis.TransactOpts, _tradeId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _rolesManager) returns()
func (_GainsAbis *GainsAbisTransactor) Initialize(opts *bind.TransactOpts, _rolesManager common.Address) (*types.Transaction, error) {
	return _GainsAbis.contract.Transact(opts, "initialize", _rolesManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _rolesManager) returns()
func (_GainsAbis *GainsAbisSession) Initialize(_rolesManager common.Address) (*types.Transaction, error) {
	return _GainsAbis.Contract.Initialize(&_GainsAbis.TransactOpts, _rolesManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _rolesManager) returns()
func (_GainsAbis *GainsAbisTransactorSession) Initialize(_rolesManager common.Address) (*types.Transaction, error) {
	return _GainsAbis.Contract.Initialize(&_GainsAbis.TransactOpts, _rolesManager)
}

// InitializeTradingStorage is a paid mutator transaction binding the contract method 0x1b7d88e5.
//
// Solidity: function initializeTradingStorage(address _gns, address _gnsStaking, address[] _collaterals, address[] _gTokens) returns()
func (_GainsAbis *GainsAbisTransactor) InitializeTradingStorage(opts *bind.TransactOpts, _gns common.Address, _gnsStaking common.Address, _collaterals []common.Address, _gTokens []common.Address) (*types.Transaction, error) {
	return _GainsAbis.contract.Transact(opts, "initializeTradingStorage", _gns, _gnsStaking, _collaterals, _gTokens)
}

// InitializeTradingStorage is a paid mutator transaction binding the contract method 0x1b7d88e5.
//
// Solidity: function initializeTradingStorage(address _gns, address _gnsStaking, address[] _collaterals, address[] _gTokens) returns()
func (_GainsAbis *GainsAbisSession) InitializeTradingStorage(_gns common.Address, _gnsStaking common.Address, _collaterals []common.Address, _gTokens []common.Address) (*types.Transaction, error) {
	return _GainsAbis.Contract.InitializeTradingStorage(&_GainsAbis.TransactOpts, _gns, _gnsStaking, _collaterals, _gTokens)
}

// InitializeTradingStorage is a paid mutator transaction binding the contract method 0x1b7d88e5.
//
// Solidity: function initializeTradingStorage(address _gns, address _gnsStaking, address[] _collaterals, address[] _gTokens) returns()
func (_GainsAbis *GainsAbisTransactorSession) InitializeTradingStorage(_gns common.Address, _gnsStaking common.Address, _collaterals []common.Address, _gTokens []common.Address) (*types.Transaction, error) {
	return _GainsAbis.Contract.InitializeTradingStorage(&_GainsAbis.TransactOpts, _gns, _gnsStaking, _collaterals, _gTokens)
}

// SetRoles is a paid mutator transaction binding the contract method 0x101e6503.
//
// Solidity: function setRoles(address[] _accounts, uint8[] _roles, bool[] _values) returns()
func (_GainsAbis *GainsAbisTransactor) SetRoles(opts *bind.TransactOpts, _accounts []common.Address, _roles []uint8, _values []bool) (*types.Transaction, error) {
	return _GainsAbis.contract.Transact(opts, "setRoles", _accounts, _roles, _values)
}

// SetRoles is a paid mutator transaction binding the contract method 0x101e6503.
//
// Solidity: function setRoles(address[] _accounts, uint8[] _roles, bool[] _values) returns()
func (_GainsAbis *GainsAbisSession) SetRoles(_accounts []common.Address, _roles []uint8, _values []bool) (*types.Transaction, error) {
	return _GainsAbis.Contract.SetRoles(&_GainsAbis.TransactOpts, _accounts, _roles, _values)
}

// SetRoles is a paid mutator transaction binding the contract method 0x101e6503.
//
// Solidity: function setRoles(address[] _accounts, uint8[] _roles, bool[] _values) returns()
func (_GainsAbis *GainsAbisTransactorSession) SetRoles(_accounts []common.Address, _roles []uint8, _values []bool) (*types.Transaction, error) {
	return _GainsAbis.Contract.SetRoles(&_GainsAbis.TransactOpts, _accounts, _roles, _values)
}

// StorePendingOrder is a paid mutator transaction binding the contract method 0x93f9384e.
//
// Solidity: function storePendingOrder(((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192),address,uint32,bool,uint8,uint32,uint16) _pendingOrder) returns(((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192),address,uint32,bool,uint8,uint32,uint16))
func (_GainsAbis *GainsAbisTransactor) StorePendingOrder(opts *bind.TransactOpts, _pendingOrder ITradingStoragePendingOrder) (*types.Transaction, error) {
	return _GainsAbis.contract.Transact(opts, "storePendingOrder", _pendingOrder)
}

// StorePendingOrder is a paid mutator transaction binding the contract method 0x93f9384e.
//
// Solidity: function storePendingOrder(((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192),address,uint32,bool,uint8,uint32,uint16) _pendingOrder) returns(((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192),address,uint32,bool,uint8,uint32,uint16))
func (_GainsAbis *GainsAbisSession) StorePendingOrder(_pendingOrder ITradingStoragePendingOrder) (*types.Transaction, error) {
	return _GainsAbis.Contract.StorePendingOrder(&_GainsAbis.TransactOpts, _pendingOrder)
}

// StorePendingOrder is a paid mutator transaction binding the contract method 0x93f9384e.
//
// Solidity: function storePendingOrder(((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192),address,uint32,bool,uint8,uint32,uint16) _pendingOrder) returns(((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192),address,uint32,bool,uint8,uint32,uint16))
func (_GainsAbis *GainsAbisTransactorSession) StorePendingOrder(_pendingOrder ITradingStoragePendingOrder) (*types.Transaction, error) {
	return _GainsAbis.Contract.StorePendingOrder(&_GainsAbis.TransactOpts, _pendingOrder)
}

// StoreTrade is a paid mutator transaction binding the contract method 0xd7ec0787.
//
// Solidity: function storeTrade((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192) _trade, (uint32,uint32,uint32,uint16,uint48,uint48,uint8,uint32,uint8) _tradeInfo) returns((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192))
func (_GainsAbis *GainsAbisTransactor) StoreTrade(opts *bind.TransactOpts, _trade ITradingStorageTrade, _tradeInfo ITradingStorageTradeInfo) (*types.Transaction, error) {
	return _GainsAbis.contract.Transact(opts, "storeTrade", _trade, _tradeInfo)
}

// StoreTrade is a paid mutator transaction binding the contract method 0xd7ec0787.
//
// Solidity: function storeTrade((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192) _trade, (uint32,uint32,uint32,uint16,uint48,uint48,uint8,uint32,uint8) _tradeInfo) returns((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192))
func (_GainsAbis *GainsAbisSession) StoreTrade(_trade ITradingStorageTrade, _tradeInfo ITradingStorageTradeInfo) (*types.Transaction, error) {
	return _GainsAbis.Contract.StoreTrade(&_GainsAbis.TransactOpts, _trade, _tradeInfo)
}

// StoreTrade is a paid mutator transaction binding the contract method 0xd7ec0787.
//
// Solidity: function storeTrade((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192) _trade, (uint32,uint32,uint32,uint16,uint48,uint48,uint8,uint32,uint8) _tradeInfo) returns((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192))
func (_GainsAbis *GainsAbisTransactorSession) StoreTrade(_trade ITradingStorageTrade, _tradeInfo ITradingStorageTradeInfo) (*types.Transaction, error) {
	return _GainsAbis.Contract.StoreTrade(&_GainsAbis.TransactOpts, _trade, _tradeInfo)
}

// ToggleCollateralActiveState is a paid mutator transaction binding the contract method 0x49f7895b.
//
// Solidity: function toggleCollateralActiveState(uint8 _collateralIndex) returns()
func (_GainsAbis *GainsAbisTransactor) ToggleCollateralActiveState(opts *bind.TransactOpts, _collateralIndex uint8) (*types.Transaction, error) {
	return _GainsAbis.contract.Transact(opts, "toggleCollateralActiveState", _collateralIndex)
}

// ToggleCollateralActiveState is a paid mutator transaction binding the contract method 0x49f7895b.
//
// Solidity: function toggleCollateralActiveState(uint8 _collateralIndex) returns()
func (_GainsAbis *GainsAbisSession) ToggleCollateralActiveState(_collateralIndex uint8) (*types.Transaction, error) {
	return _GainsAbis.Contract.ToggleCollateralActiveState(&_GainsAbis.TransactOpts, _collateralIndex)
}

// ToggleCollateralActiveState is a paid mutator transaction binding the contract method 0x49f7895b.
//
// Solidity: function toggleCollateralActiveState(uint8 _collateralIndex) returns()
func (_GainsAbis *GainsAbisTransactorSession) ToggleCollateralActiveState(_collateralIndex uint8) (*types.Transaction, error) {
	return _GainsAbis.Contract.ToggleCollateralActiveState(&_GainsAbis.TransactOpts, _collateralIndex)
}

// UpdateGToken is a paid mutator transaction binding the contract method 0x63450d74.
//
// Solidity: function updateGToken(address _collateral, address _gToken) returns()
func (_GainsAbis *GainsAbisTransactor) UpdateGToken(opts *bind.TransactOpts, _collateral common.Address, _gToken common.Address) (*types.Transaction, error) {
	return _GainsAbis.contract.Transact(opts, "updateGToken", _collateral, _gToken)
}

// UpdateGToken is a paid mutator transaction binding the contract method 0x63450d74.
//
// Solidity: function updateGToken(address _collateral, address _gToken) returns()
func (_GainsAbis *GainsAbisSession) UpdateGToken(_collateral common.Address, _gToken common.Address) (*types.Transaction, error) {
	return _GainsAbis.Contract.UpdateGToken(&_GainsAbis.TransactOpts, _collateral, _gToken)
}

// UpdateGToken is a paid mutator transaction binding the contract method 0x63450d74.
//
// Solidity: function updateGToken(address _collateral, address _gToken) returns()
func (_GainsAbis *GainsAbisTransactorSession) UpdateGToken(_collateral common.Address, _gToken common.Address) (*types.Transaction, error) {
	return _GainsAbis.Contract.UpdateGToken(&_GainsAbis.TransactOpts, _collateral, _gToken)
}

// UpdateOpenOrderDetails is a paid mutator transaction binding the contract method 0xeb2dfde8.
//
// Solidity: function updateOpenOrderDetails((address,uint32) _tradeId, uint64 _openPrice, uint64 _tp, uint64 _sl, uint16 _maxSlippageP) returns()
func (_GainsAbis *GainsAbisTransactor) UpdateOpenOrderDetails(opts *bind.TransactOpts, _tradeId ITradingStorageId, _openPrice uint64, _tp uint64, _sl uint64, _maxSlippageP uint16) (*types.Transaction, error) {
	return _GainsAbis.contract.Transact(opts, "updateOpenOrderDetails", _tradeId, _openPrice, _tp, _sl, _maxSlippageP)
}

// UpdateOpenOrderDetails is a paid mutator transaction binding the contract method 0xeb2dfde8.
//
// Solidity: function updateOpenOrderDetails((address,uint32) _tradeId, uint64 _openPrice, uint64 _tp, uint64 _sl, uint16 _maxSlippageP) returns()
func (_GainsAbis *GainsAbisSession) UpdateOpenOrderDetails(_tradeId ITradingStorageId, _openPrice uint64, _tp uint64, _sl uint64, _maxSlippageP uint16) (*types.Transaction, error) {
	return _GainsAbis.Contract.UpdateOpenOrderDetails(&_GainsAbis.TransactOpts, _tradeId, _openPrice, _tp, _sl, _maxSlippageP)
}

// UpdateOpenOrderDetails is a paid mutator transaction binding the contract method 0xeb2dfde8.
//
// Solidity: function updateOpenOrderDetails((address,uint32) _tradeId, uint64 _openPrice, uint64 _tp, uint64 _sl, uint16 _maxSlippageP) returns()
func (_GainsAbis *GainsAbisTransactorSession) UpdateOpenOrderDetails(_tradeId ITradingStorageId, _openPrice uint64, _tp uint64, _sl uint64, _maxSlippageP uint16) (*types.Transaction, error) {
	return _GainsAbis.Contract.UpdateOpenOrderDetails(&_GainsAbis.TransactOpts, _tradeId, _openPrice, _tp, _sl, _maxSlippageP)
}

// UpdateTradeCollateralAmount is a paid mutator transaction binding the contract method 0x5a68200d.
//
// Solidity: function updateTradeCollateralAmount((address,uint32) _tradeId, uint120 _collateralAmount) returns()
func (_GainsAbis *GainsAbisTransactor) UpdateTradeCollateralAmount(opts *bind.TransactOpts, _tradeId ITradingStorageId, _collateralAmount *big.Int) (*types.Transaction, error) {
	return _GainsAbis.contract.Transact(opts, "updateTradeCollateralAmount", _tradeId, _collateralAmount)
}

// UpdateTradeCollateralAmount is a paid mutator transaction binding the contract method 0x5a68200d.
//
// Solidity: function updateTradeCollateralAmount((address,uint32) _tradeId, uint120 _collateralAmount) returns()
func (_GainsAbis *GainsAbisSession) UpdateTradeCollateralAmount(_tradeId ITradingStorageId, _collateralAmount *big.Int) (*types.Transaction, error) {
	return _GainsAbis.Contract.UpdateTradeCollateralAmount(&_GainsAbis.TransactOpts, _tradeId, _collateralAmount)
}

// UpdateTradeCollateralAmount is a paid mutator transaction binding the contract method 0x5a68200d.
//
// Solidity: function updateTradeCollateralAmount((address,uint32) _tradeId, uint120 _collateralAmount) returns()
func (_GainsAbis *GainsAbisTransactorSession) UpdateTradeCollateralAmount(_tradeId ITradingStorageId, _collateralAmount *big.Int) (*types.Transaction, error) {
	return _GainsAbis.Contract.UpdateTradeCollateralAmount(&_GainsAbis.TransactOpts, _tradeId, _collateralAmount)
}

// UpdateTradeMaxClosingSlippageP is a paid mutator transaction binding the contract method 0x07d426fd.
//
// Solidity: function updateTradeMaxClosingSlippageP((address,uint32) _tradeId, uint16 _maxSlippageP) returns()
func (_GainsAbis *GainsAbisTransactor) UpdateTradeMaxClosingSlippageP(opts *bind.TransactOpts, _tradeId ITradingStorageId, _maxSlippageP uint16) (*types.Transaction, error) {
	return _GainsAbis.contract.Transact(opts, "updateTradeMaxClosingSlippageP", _tradeId, _maxSlippageP)
}

// UpdateTradeMaxClosingSlippageP is a paid mutator transaction binding the contract method 0x07d426fd.
//
// Solidity: function updateTradeMaxClosingSlippageP((address,uint32) _tradeId, uint16 _maxSlippageP) returns()
func (_GainsAbis *GainsAbisSession) UpdateTradeMaxClosingSlippageP(_tradeId ITradingStorageId, _maxSlippageP uint16) (*types.Transaction, error) {
	return _GainsAbis.Contract.UpdateTradeMaxClosingSlippageP(&_GainsAbis.TransactOpts, _tradeId, _maxSlippageP)
}

// UpdateTradeMaxClosingSlippageP is a paid mutator transaction binding the contract method 0x07d426fd.
//
// Solidity: function updateTradeMaxClosingSlippageP((address,uint32) _tradeId, uint16 _maxSlippageP) returns()
func (_GainsAbis *GainsAbisTransactorSession) UpdateTradeMaxClosingSlippageP(_tradeId ITradingStorageId, _maxSlippageP uint16) (*types.Transaction, error) {
	return _GainsAbis.Contract.UpdateTradeMaxClosingSlippageP(&_GainsAbis.TransactOpts, _tradeId, _maxSlippageP)
}

// UpdateTradePosition is a paid mutator transaction binding the contract method 0x72570e24.
//
// Solidity: function updateTradePosition((address,uint32) _tradeId, uint120 _collateralAmount, uint24 _leverage, uint64 _openPrice, bool _isPartialIncrease) returns()
func (_GainsAbis *GainsAbisTransactor) UpdateTradePosition(opts *bind.TransactOpts, _tradeId ITradingStorageId, _collateralAmount *big.Int, _leverage *big.Int, _openPrice uint64, _isPartialIncrease bool) (*types.Transaction, error) {
	return _GainsAbis.contract.Transact(opts, "updateTradePosition", _tradeId, _collateralAmount, _leverage, _openPrice, _isPartialIncrease)
}

// UpdateTradePosition is a paid mutator transaction binding the contract method 0x72570e24.
//
// Solidity: function updateTradePosition((address,uint32) _tradeId, uint120 _collateralAmount, uint24 _leverage, uint64 _openPrice, bool _isPartialIncrease) returns()
func (_GainsAbis *GainsAbisSession) UpdateTradePosition(_tradeId ITradingStorageId, _collateralAmount *big.Int, _leverage *big.Int, _openPrice uint64, _isPartialIncrease bool) (*types.Transaction, error) {
	return _GainsAbis.Contract.UpdateTradePosition(&_GainsAbis.TransactOpts, _tradeId, _collateralAmount, _leverage, _openPrice, _isPartialIncrease)
}

// UpdateTradePosition is a paid mutator transaction binding the contract method 0x72570e24.
//
// Solidity: function updateTradePosition((address,uint32) _tradeId, uint120 _collateralAmount, uint24 _leverage, uint64 _openPrice, bool _isPartialIncrease) returns()
func (_GainsAbis *GainsAbisTransactorSession) UpdateTradePosition(_tradeId ITradingStorageId, _collateralAmount *big.Int, _leverage *big.Int, _openPrice uint64, _isPartialIncrease bool) (*types.Transaction, error) {
	return _GainsAbis.Contract.UpdateTradePosition(&_GainsAbis.TransactOpts, _tradeId, _collateralAmount, _leverage, _openPrice, _isPartialIncrease)
}

// UpdateTradeSl is a paid mutator transaction binding the contract method 0x1053c279.
//
// Solidity: function updateTradeSl((address,uint32) _tradeId, uint64 _newSl) returns()
func (_GainsAbis *GainsAbisTransactor) UpdateTradeSl(opts *bind.TransactOpts, _tradeId ITradingStorageId, _newSl uint64) (*types.Transaction, error) {
	return _GainsAbis.contract.Transact(opts, "updateTradeSl", _tradeId, _newSl)
}

// UpdateTradeSl is a paid mutator transaction binding the contract method 0x1053c279.
//
// Solidity: function updateTradeSl((address,uint32) _tradeId, uint64 _newSl) returns()
func (_GainsAbis *GainsAbisSession) UpdateTradeSl(_tradeId ITradingStorageId, _newSl uint64) (*types.Transaction, error) {
	return _GainsAbis.Contract.UpdateTradeSl(&_GainsAbis.TransactOpts, _tradeId, _newSl)
}

// UpdateTradeSl is a paid mutator transaction binding the contract method 0x1053c279.
//
// Solidity: function updateTradeSl((address,uint32) _tradeId, uint64 _newSl) returns()
func (_GainsAbis *GainsAbisTransactorSession) UpdateTradeSl(_tradeId ITradingStorageId, _newSl uint64) (*types.Transaction, error) {
	return _GainsAbis.Contract.UpdateTradeSl(&_GainsAbis.TransactOpts, _tradeId, _newSl)
}

// UpdateTradeTp is a paid mutator transaction binding the contract method 0xb8f741d4.
//
// Solidity: function updateTradeTp((address,uint32) _tradeId, uint64 _newTp) returns()
func (_GainsAbis *GainsAbisTransactor) UpdateTradeTp(opts *bind.TransactOpts, _tradeId ITradingStorageId, _newTp uint64) (*types.Transaction, error) {
	return _GainsAbis.contract.Transact(opts, "updateTradeTp", _tradeId, _newTp)
}

// UpdateTradeTp is a paid mutator transaction binding the contract method 0xb8f741d4.
//
// Solidity: function updateTradeTp((address,uint32) _tradeId, uint64 _newTp) returns()
func (_GainsAbis *GainsAbisSession) UpdateTradeTp(_tradeId ITradingStorageId, _newTp uint64) (*types.Transaction, error) {
	return _GainsAbis.Contract.UpdateTradeTp(&_GainsAbis.TransactOpts, _tradeId, _newTp)
}

// UpdateTradeTp is a paid mutator transaction binding the contract method 0xb8f741d4.
//
// Solidity: function updateTradeTp((address,uint32) _tradeId, uint64 _newTp) returns()
func (_GainsAbis *GainsAbisTransactorSession) UpdateTradeTp(_tradeId ITradingStorageId, _newTp uint64) (*types.Transaction, error) {
	return _GainsAbis.Contract.UpdateTradeTp(&_GainsAbis.TransactOpts, _tradeId, _newTp)
}

// UpdateTradingActivated is a paid mutator transaction binding the contract method 0xb78f4b36.
//
// Solidity: function updateTradingActivated(uint8 _activated) returns()
func (_GainsAbis *GainsAbisTransactor) UpdateTradingActivated(opts *bind.TransactOpts, _activated uint8) (*types.Transaction, error) {
	return _GainsAbis.contract.Transact(opts, "updateTradingActivated", _activated)
}

// UpdateTradingActivated is a paid mutator transaction binding the contract method 0xb78f4b36.
//
// Solidity: function updateTradingActivated(uint8 _activated) returns()
func (_GainsAbis *GainsAbisSession) UpdateTradingActivated(_activated uint8) (*types.Transaction, error) {
	return _GainsAbis.Contract.UpdateTradingActivated(&_GainsAbis.TransactOpts, _activated)
}

// UpdateTradingActivated is a paid mutator transaction binding the contract method 0xb78f4b36.
//
// Solidity: function updateTradingActivated(uint8 _activated) returns()
func (_GainsAbis *GainsAbisTransactorSession) UpdateTradingActivated(_activated uint8) (*types.Transaction, error) {
	return _GainsAbis.Contract.UpdateTradingActivated(&_GainsAbis.TransactOpts, _activated)
}

// GainsAbisAccessControlUpdatedIterator is returned from FilterAccessControlUpdated and is used to iterate over the raw logs and unpacked data for AccessControlUpdated events raised by the GainsAbis contract.
type GainsAbisAccessControlUpdatedIterator struct {
	Event *GainsAbisAccessControlUpdated // Event containing the contract specifics and raw log

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
func (it *GainsAbisAccessControlUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GainsAbisAccessControlUpdated)
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
		it.Event = new(GainsAbisAccessControlUpdated)
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
func (it *GainsAbisAccessControlUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GainsAbisAccessControlUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GainsAbisAccessControlUpdated represents a AccessControlUpdated event raised by the GainsAbis contract.
type GainsAbisAccessControlUpdated struct {
	Target common.Address
	Role   uint8
	Access bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAccessControlUpdated is a free log retrieval operation binding the contract event 0x8d7fdec37f50c07219a6a0859420936836eb9254bf412035e3acede18b8b093d.
//
// Solidity: event AccessControlUpdated(address target, uint8 role, bool access)
func (_GainsAbis *GainsAbisFilterer) FilterAccessControlUpdated(opts *bind.FilterOpts) (*GainsAbisAccessControlUpdatedIterator, error) {

	logs, sub, err := _GainsAbis.contract.FilterLogs(opts, "AccessControlUpdated")
	if err != nil {
		return nil, err
	}
	return &GainsAbisAccessControlUpdatedIterator{contract: _GainsAbis.contract, event: "AccessControlUpdated", logs: logs, sub: sub}, nil
}

// WatchAccessControlUpdated is a free log subscription operation binding the contract event 0x8d7fdec37f50c07219a6a0859420936836eb9254bf412035e3acede18b8b093d.
//
// Solidity: event AccessControlUpdated(address target, uint8 role, bool access)
func (_GainsAbis *GainsAbisFilterer) WatchAccessControlUpdated(opts *bind.WatchOpts, sink chan<- *GainsAbisAccessControlUpdated) (event.Subscription, error) {

	logs, sub, err := _GainsAbis.contract.WatchLogs(opts, "AccessControlUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GainsAbisAccessControlUpdated)
				if err := _GainsAbis.contract.UnpackLog(event, "AccessControlUpdated", log); err != nil {
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

// ParseAccessControlUpdated is a log parse operation binding the contract event 0x8d7fdec37f50c07219a6a0859420936836eb9254bf412035e3acede18b8b093d.
//
// Solidity: event AccessControlUpdated(address target, uint8 role, bool access)
func (_GainsAbis *GainsAbisFilterer) ParseAccessControlUpdated(log types.Log) (*GainsAbisAccessControlUpdated, error) {
	event := new(GainsAbisAccessControlUpdated)
	if err := _GainsAbis.contract.UnpackLog(event, "AccessControlUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GainsAbisAddressesUpdatedIterator is returned from FilterAddressesUpdated and is used to iterate over the raw logs and unpacked data for AddressesUpdated events raised by the GainsAbis contract.
type GainsAbisAddressesUpdatedIterator struct {
	Event *GainsAbisAddressesUpdated // Event containing the contract specifics and raw log

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
func (it *GainsAbisAddressesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GainsAbisAddressesUpdated)
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
		it.Event = new(GainsAbisAddressesUpdated)
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
func (it *GainsAbisAddressesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GainsAbisAddressesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GainsAbisAddressesUpdated represents a AddressesUpdated event raised by the GainsAbis contract.
type GainsAbisAddressesUpdated struct {
	Addresses IAddressStoreAddresses
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddressesUpdated is a free log retrieval operation binding the contract event 0xe4f1f9461410dada4f4b49a4b363bdf35e6069fb5a0cea4b1147c32affbd954a.
//
// Solidity: event AddressesUpdated((address,address) addresses)
func (_GainsAbis *GainsAbisFilterer) FilterAddressesUpdated(opts *bind.FilterOpts) (*GainsAbisAddressesUpdatedIterator, error) {

	logs, sub, err := _GainsAbis.contract.FilterLogs(opts, "AddressesUpdated")
	if err != nil {
		return nil, err
	}
	return &GainsAbisAddressesUpdatedIterator{contract: _GainsAbis.contract, event: "AddressesUpdated", logs: logs, sub: sub}, nil
}

// WatchAddressesUpdated is a free log subscription operation binding the contract event 0xe4f1f9461410dada4f4b49a4b363bdf35e6069fb5a0cea4b1147c32affbd954a.
//
// Solidity: event AddressesUpdated((address,address) addresses)
func (_GainsAbis *GainsAbisFilterer) WatchAddressesUpdated(opts *bind.WatchOpts, sink chan<- *GainsAbisAddressesUpdated) (event.Subscription, error) {

	logs, sub, err := _GainsAbis.contract.WatchLogs(opts, "AddressesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GainsAbisAddressesUpdated)
				if err := _GainsAbis.contract.UnpackLog(event, "AddressesUpdated", log); err != nil {
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

// ParseAddressesUpdated is a log parse operation binding the contract event 0xe4f1f9461410dada4f4b49a4b363bdf35e6069fb5a0cea4b1147c32affbd954a.
//
// Solidity: event AddressesUpdated((address,address) addresses)
func (_GainsAbis *GainsAbisFilterer) ParseAddressesUpdated(log types.Log) (*GainsAbisAddressesUpdated, error) {
	event := new(GainsAbisAddressesUpdated)
	if err := _GainsAbis.contract.UnpackLog(event, "AddressesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GainsAbisCollateralAddedIterator is returned from FilterCollateralAdded and is used to iterate over the raw logs and unpacked data for CollateralAdded events raised by the GainsAbis contract.
type GainsAbisCollateralAddedIterator struct {
	Event *GainsAbisCollateralAdded // Event containing the contract specifics and raw log

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
func (it *GainsAbisCollateralAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GainsAbisCollateralAdded)
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
		it.Event = new(GainsAbisCollateralAdded)
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
func (it *GainsAbisCollateralAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GainsAbisCollateralAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GainsAbisCollateralAdded represents a CollateralAdded event raised by the GainsAbis contract.
type GainsAbisCollateralAdded struct {
	Collateral common.Address
	Index      uint8
	GToken     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollateralAdded is a free log retrieval operation binding the contract event 0xa02b5df63a0ca2660cbe23b5eb92c7f2ae514aee4a543a6032b38ef338865dbf.
//
// Solidity: event CollateralAdded(address collateral, uint8 index, address gToken)
func (_GainsAbis *GainsAbisFilterer) FilterCollateralAdded(opts *bind.FilterOpts) (*GainsAbisCollateralAddedIterator, error) {

	logs, sub, err := _GainsAbis.contract.FilterLogs(opts, "CollateralAdded")
	if err != nil {
		return nil, err
	}
	return &GainsAbisCollateralAddedIterator{contract: _GainsAbis.contract, event: "CollateralAdded", logs: logs, sub: sub}, nil
}

// WatchCollateralAdded is a free log subscription operation binding the contract event 0xa02b5df63a0ca2660cbe23b5eb92c7f2ae514aee4a543a6032b38ef338865dbf.
//
// Solidity: event CollateralAdded(address collateral, uint8 index, address gToken)
func (_GainsAbis *GainsAbisFilterer) WatchCollateralAdded(opts *bind.WatchOpts, sink chan<- *GainsAbisCollateralAdded) (event.Subscription, error) {

	logs, sub, err := _GainsAbis.contract.WatchLogs(opts, "CollateralAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GainsAbisCollateralAdded)
				if err := _GainsAbis.contract.UnpackLog(event, "CollateralAdded", log); err != nil {
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

// ParseCollateralAdded is a log parse operation binding the contract event 0xa02b5df63a0ca2660cbe23b5eb92c7f2ae514aee4a543a6032b38ef338865dbf.
//
// Solidity: event CollateralAdded(address collateral, uint8 index, address gToken)
func (_GainsAbis *GainsAbisFilterer) ParseCollateralAdded(log types.Log) (*GainsAbisCollateralAdded, error) {
	event := new(GainsAbisCollateralAdded)
	if err := _GainsAbis.contract.UnpackLog(event, "CollateralAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GainsAbisCollateralDisabledIterator is returned from FilterCollateralDisabled and is used to iterate over the raw logs and unpacked data for CollateralDisabled events raised by the GainsAbis contract.
type GainsAbisCollateralDisabledIterator struct {
	Event *GainsAbisCollateralDisabled // Event containing the contract specifics and raw log

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
func (it *GainsAbisCollateralDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GainsAbisCollateralDisabled)
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
		it.Event = new(GainsAbisCollateralDisabled)
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
func (it *GainsAbisCollateralDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GainsAbisCollateralDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GainsAbisCollateralDisabled represents a CollateralDisabled event raised by the GainsAbis contract.
type GainsAbisCollateralDisabled struct {
	Index uint8
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCollateralDisabled is a free log retrieval operation binding the contract event 0x09a6e6672fd5a685707eca1eeb3a3ef190ccf5ceaf9a78e410859f2d7983cc92.
//
// Solidity: event CollateralDisabled(uint8 index)
func (_GainsAbis *GainsAbisFilterer) FilterCollateralDisabled(opts *bind.FilterOpts) (*GainsAbisCollateralDisabledIterator, error) {

	logs, sub, err := _GainsAbis.contract.FilterLogs(opts, "CollateralDisabled")
	if err != nil {
		return nil, err
	}
	return &GainsAbisCollateralDisabledIterator{contract: _GainsAbis.contract, event: "CollateralDisabled", logs: logs, sub: sub}, nil
}

// WatchCollateralDisabled is a free log subscription operation binding the contract event 0x09a6e6672fd5a685707eca1eeb3a3ef190ccf5ceaf9a78e410859f2d7983cc92.
//
// Solidity: event CollateralDisabled(uint8 index)
func (_GainsAbis *GainsAbisFilterer) WatchCollateralDisabled(opts *bind.WatchOpts, sink chan<- *GainsAbisCollateralDisabled) (event.Subscription, error) {

	logs, sub, err := _GainsAbis.contract.WatchLogs(opts, "CollateralDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GainsAbisCollateralDisabled)
				if err := _GainsAbis.contract.UnpackLog(event, "CollateralDisabled", log); err != nil {
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

// ParseCollateralDisabled is a log parse operation binding the contract event 0x09a6e6672fd5a685707eca1eeb3a3ef190ccf5ceaf9a78e410859f2d7983cc92.
//
// Solidity: event CollateralDisabled(uint8 index)
func (_GainsAbis *GainsAbisFilterer) ParseCollateralDisabled(log types.Log) (*GainsAbisCollateralDisabled, error) {
	event := new(GainsAbisCollateralDisabled)
	if err := _GainsAbis.contract.UnpackLog(event, "CollateralDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GainsAbisCollateralUpdatedIterator is returned from FilterCollateralUpdated and is used to iterate over the raw logs and unpacked data for CollateralUpdated events raised by the GainsAbis contract.
type GainsAbisCollateralUpdatedIterator struct {
	Event *GainsAbisCollateralUpdated // Event containing the contract specifics and raw log

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
func (it *GainsAbisCollateralUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GainsAbisCollateralUpdated)
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
		it.Event = new(GainsAbisCollateralUpdated)
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
func (it *GainsAbisCollateralUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GainsAbisCollateralUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GainsAbisCollateralUpdated represents a CollateralUpdated event raised by the GainsAbis contract.
type GainsAbisCollateralUpdated struct {
	Index    uint8
	IsActive bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCollateralUpdated is a free log retrieval operation binding the contract event 0x98bbde8d067842c4760a76b32aebf2cd4feb8f07ddcf20d81c619c16f0242ecb.
//
// Solidity: event CollateralUpdated(uint8 indexed index, bool isActive)
func (_GainsAbis *GainsAbisFilterer) FilterCollateralUpdated(opts *bind.FilterOpts, index []uint8) (*GainsAbisCollateralUpdatedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _GainsAbis.contract.FilterLogs(opts, "CollateralUpdated", indexRule)
	if err != nil {
		return nil, err
	}
	return &GainsAbisCollateralUpdatedIterator{contract: _GainsAbis.contract, event: "CollateralUpdated", logs: logs, sub: sub}, nil
}

// WatchCollateralUpdated is a free log subscription operation binding the contract event 0x98bbde8d067842c4760a76b32aebf2cd4feb8f07ddcf20d81c619c16f0242ecb.
//
// Solidity: event CollateralUpdated(uint8 indexed index, bool isActive)
func (_GainsAbis *GainsAbisFilterer) WatchCollateralUpdated(opts *bind.WatchOpts, sink chan<- *GainsAbisCollateralUpdated, index []uint8) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _GainsAbis.contract.WatchLogs(opts, "CollateralUpdated", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GainsAbisCollateralUpdated)
				if err := _GainsAbis.contract.UnpackLog(event, "CollateralUpdated", log); err != nil {
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

// ParseCollateralUpdated is a log parse operation binding the contract event 0x98bbde8d067842c4760a76b32aebf2cd4feb8f07ddcf20d81c619c16f0242ecb.
//
// Solidity: event CollateralUpdated(uint8 indexed index, bool isActive)
func (_GainsAbis *GainsAbisFilterer) ParseCollateralUpdated(log types.Log) (*GainsAbisCollateralUpdated, error) {
	event := new(GainsAbisCollateralUpdated)
	if err := _GainsAbis.contract.UnpackLog(event, "CollateralUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GainsAbisGTokenUpdatedIterator is returned from FilterGTokenUpdated and is used to iterate over the raw logs and unpacked data for GTokenUpdated events raised by the GainsAbis contract.
type GainsAbisGTokenUpdatedIterator struct {
	Event *GainsAbisGTokenUpdated // Event containing the contract specifics and raw log

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
func (it *GainsAbisGTokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GainsAbisGTokenUpdated)
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
		it.Event = new(GainsAbisGTokenUpdated)
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
func (it *GainsAbisGTokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GainsAbisGTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GainsAbisGTokenUpdated represents a GTokenUpdated event raised by the GainsAbis contract.
type GainsAbisGTokenUpdated struct {
	Collateral common.Address
	Index      uint8
	GToken     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGTokenUpdated is a free log retrieval operation binding the contract event 0x347ad17cfe896bbbbdf75fa51fd03a1f1366df72ba0baf20ebed1ea1394a8ecd.
//
// Solidity: event GTokenUpdated(address collateral, uint8 index, address gToken)
func (_GainsAbis *GainsAbisFilterer) FilterGTokenUpdated(opts *bind.FilterOpts) (*GainsAbisGTokenUpdatedIterator, error) {

	logs, sub, err := _GainsAbis.contract.FilterLogs(opts, "GTokenUpdated")
	if err != nil {
		return nil, err
	}
	return &GainsAbisGTokenUpdatedIterator{contract: _GainsAbis.contract, event: "GTokenUpdated", logs: logs, sub: sub}, nil
}

// WatchGTokenUpdated is a free log subscription operation binding the contract event 0x347ad17cfe896bbbbdf75fa51fd03a1f1366df72ba0baf20ebed1ea1394a8ecd.
//
// Solidity: event GTokenUpdated(address collateral, uint8 index, address gToken)
func (_GainsAbis *GainsAbisFilterer) WatchGTokenUpdated(opts *bind.WatchOpts, sink chan<- *GainsAbisGTokenUpdated) (event.Subscription, error) {

	logs, sub, err := _GainsAbis.contract.WatchLogs(opts, "GTokenUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GainsAbisGTokenUpdated)
				if err := _GainsAbis.contract.UnpackLog(event, "GTokenUpdated", log); err != nil {
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

// ParseGTokenUpdated is a log parse operation binding the contract event 0x347ad17cfe896bbbbdf75fa51fd03a1f1366df72ba0baf20ebed1ea1394a8ecd.
//
// Solidity: event GTokenUpdated(address collateral, uint8 index, address gToken)
func (_GainsAbis *GainsAbisFilterer) ParseGTokenUpdated(log types.Log) (*GainsAbisGTokenUpdated, error) {
	event := new(GainsAbisGTokenUpdated)
	if err := _GainsAbis.contract.UnpackLog(event, "GTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GainsAbisInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GainsAbis contract.
type GainsAbisInitializedIterator struct {
	Event *GainsAbisInitialized // Event containing the contract specifics and raw log

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
func (it *GainsAbisInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GainsAbisInitialized)
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
		it.Event = new(GainsAbisInitialized)
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
func (it *GainsAbisInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GainsAbisInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GainsAbisInitialized represents a Initialized event raised by the GainsAbis contract.
type GainsAbisInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GainsAbis *GainsAbisFilterer) FilterInitialized(opts *bind.FilterOpts) (*GainsAbisInitializedIterator, error) {

	logs, sub, err := _GainsAbis.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GainsAbisInitializedIterator{contract: _GainsAbis.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GainsAbis *GainsAbisFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GainsAbisInitialized) (event.Subscription, error) {

	logs, sub, err := _GainsAbis.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GainsAbisInitialized)
				if err := _GainsAbis.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GainsAbis *GainsAbisFilterer) ParseInitialized(log types.Log) (*GainsAbisInitialized, error) {
	event := new(GainsAbisInitialized)
	if err := _GainsAbis.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GainsAbisOpenOrderDetailsUpdatedIterator is returned from FilterOpenOrderDetailsUpdated and is used to iterate over the raw logs and unpacked data for OpenOrderDetailsUpdated events raised by the GainsAbis contract.
type GainsAbisOpenOrderDetailsUpdatedIterator struct {
	Event *GainsAbisOpenOrderDetailsUpdated // Event containing the contract specifics and raw log

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
func (it *GainsAbisOpenOrderDetailsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GainsAbisOpenOrderDetailsUpdated)
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
		it.Event = new(GainsAbisOpenOrderDetailsUpdated)
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
func (it *GainsAbisOpenOrderDetailsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GainsAbisOpenOrderDetailsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GainsAbisOpenOrderDetailsUpdated represents a OpenOrderDetailsUpdated event raised by the GainsAbis contract.
type GainsAbisOpenOrderDetailsUpdated struct {
	TradeId      ITradingStorageId
	OpenPrice    uint64
	Tp           uint64
	Sl           uint64
	MaxSlippageP uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOpenOrderDetailsUpdated is a free log retrieval operation binding the contract event 0x57166866105b85933cf7d2f84637e524028a4ca84133309f14b2ce0dfc113498.
//
// Solidity: event OpenOrderDetailsUpdated((address,uint32) tradeId, uint64 openPrice, uint64 tp, uint64 sl, uint16 maxSlippageP)
func (_GainsAbis *GainsAbisFilterer) FilterOpenOrderDetailsUpdated(opts *bind.FilterOpts) (*GainsAbisOpenOrderDetailsUpdatedIterator, error) {

	logs, sub, err := _GainsAbis.contract.FilterLogs(opts, "OpenOrderDetailsUpdated")
	if err != nil {
		return nil, err
	}
	return &GainsAbisOpenOrderDetailsUpdatedIterator{contract: _GainsAbis.contract, event: "OpenOrderDetailsUpdated", logs: logs, sub: sub}, nil
}

// WatchOpenOrderDetailsUpdated is a free log subscription operation binding the contract event 0x57166866105b85933cf7d2f84637e524028a4ca84133309f14b2ce0dfc113498.
//
// Solidity: event OpenOrderDetailsUpdated((address,uint32) tradeId, uint64 openPrice, uint64 tp, uint64 sl, uint16 maxSlippageP)
func (_GainsAbis *GainsAbisFilterer) WatchOpenOrderDetailsUpdated(opts *bind.WatchOpts, sink chan<- *GainsAbisOpenOrderDetailsUpdated) (event.Subscription, error) {

	logs, sub, err := _GainsAbis.contract.WatchLogs(opts, "OpenOrderDetailsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GainsAbisOpenOrderDetailsUpdated)
				if err := _GainsAbis.contract.UnpackLog(event, "OpenOrderDetailsUpdated", log); err != nil {
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

// ParseOpenOrderDetailsUpdated is a log parse operation binding the contract event 0x57166866105b85933cf7d2f84637e524028a4ca84133309f14b2ce0dfc113498.
//
// Solidity: event OpenOrderDetailsUpdated((address,uint32) tradeId, uint64 openPrice, uint64 tp, uint64 sl, uint16 maxSlippageP)
func (_GainsAbis *GainsAbisFilterer) ParseOpenOrderDetailsUpdated(log types.Log) (*GainsAbisOpenOrderDetailsUpdated, error) {
	event := new(GainsAbisOpenOrderDetailsUpdated)
	if err := _GainsAbis.contract.UnpackLog(event, "OpenOrderDetailsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GainsAbisPendingOrderClosedIterator is returned from FilterPendingOrderClosed and is used to iterate over the raw logs and unpacked data for PendingOrderClosed events raised by the GainsAbis contract.
type GainsAbisPendingOrderClosedIterator struct {
	Event *GainsAbisPendingOrderClosed // Event containing the contract specifics and raw log

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
func (it *GainsAbisPendingOrderClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GainsAbisPendingOrderClosed)
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
		it.Event = new(GainsAbisPendingOrderClosed)
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
func (it *GainsAbisPendingOrderClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GainsAbisPendingOrderClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GainsAbisPendingOrderClosed represents a PendingOrderClosed event raised by the GainsAbis contract.
type GainsAbisPendingOrderClosed struct {
	OrderId ITradingStorageId
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPendingOrderClosed is a free log retrieval operation binding the contract event 0xf0e19a36a85c073783ad5d0a8026dffa190d250d673c8c80b687cbef125571f3.
//
// Solidity: event PendingOrderClosed((address,uint32) orderId)
func (_GainsAbis *GainsAbisFilterer) FilterPendingOrderClosed(opts *bind.FilterOpts) (*GainsAbisPendingOrderClosedIterator, error) {

	logs, sub, err := _GainsAbis.contract.FilterLogs(opts, "PendingOrderClosed")
	if err != nil {
		return nil, err
	}
	return &GainsAbisPendingOrderClosedIterator{contract: _GainsAbis.contract, event: "PendingOrderClosed", logs: logs, sub: sub}, nil
}

// WatchPendingOrderClosed is a free log subscription operation binding the contract event 0xf0e19a36a85c073783ad5d0a8026dffa190d250d673c8c80b687cbef125571f3.
//
// Solidity: event PendingOrderClosed((address,uint32) orderId)
func (_GainsAbis *GainsAbisFilterer) WatchPendingOrderClosed(opts *bind.WatchOpts, sink chan<- *GainsAbisPendingOrderClosed) (event.Subscription, error) {

	logs, sub, err := _GainsAbis.contract.WatchLogs(opts, "PendingOrderClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GainsAbisPendingOrderClosed)
				if err := _GainsAbis.contract.UnpackLog(event, "PendingOrderClosed", log); err != nil {
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

// ParsePendingOrderClosed is a log parse operation binding the contract event 0xf0e19a36a85c073783ad5d0a8026dffa190d250d673c8c80b687cbef125571f3.
//
// Solidity: event PendingOrderClosed((address,uint32) orderId)
func (_GainsAbis *GainsAbisFilterer) ParsePendingOrderClosed(log types.Log) (*GainsAbisPendingOrderClosed, error) {
	event := new(GainsAbisPendingOrderClosed)
	if err := _GainsAbis.contract.UnpackLog(event, "PendingOrderClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GainsAbisPendingOrderStoredIterator is returned from FilterPendingOrderStored and is used to iterate over the raw logs and unpacked data for PendingOrderStored events raised by the GainsAbis contract.
type GainsAbisPendingOrderStoredIterator struct {
	Event *GainsAbisPendingOrderStored // Event containing the contract specifics and raw log

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
func (it *GainsAbisPendingOrderStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GainsAbisPendingOrderStored)
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
		it.Event = new(GainsAbisPendingOrderStored)
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
func (it *GainsAbisPendingOrderStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GainsAbisPendingOrderStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GainsAbisPendingOrderStored represents a PendingOrderStored event raised by the GainsAbis contract.
type GainsAbisPendingOrderStored struct {
	PendingOrder ITradingStoragePendingOrder
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPendingOrderStored is a free log retrieval operation binding the contract event 0xc1f6d032e333e12d4ba1d8cdf8c4abc1bcaab7381a4eaa19a918a28f223f519d.
//
// Solidity: event PendingOrderStored(((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192),address,uint32,bool,uint8,uint32,uint16) pendingOrder)
func (_GainsAbis *GainsAbisFilterer) FilterPendingOrderStored(opts *bind.FilterOpts) (*GainsAbisPendingOrderStoredIterator, error) {

	logs, sub, err := _GainsAbis.contract.FilterLogs(opts, "PendingOrderStored")
	if err != nil {
		return nil, err
	}
	return &GainsAbisPendingOrderStoredIterator{contract: _GainsAbis.contract, event: "PendingOrderStored", logs: logs, sub: sub}, nil
}

// WatchPendingOrderStored is a free log subscription operation binding the contract event 0xc1f6d032e333e12d4ba1d8cdf8c4abc1bcaab7381a4eaa19a918a28f223f519d.
//
// Solidity: event PendingOrderStored(((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192),address,uint32,bool,uint8,uint32,uint16) pendingOrder)
func (_GainsAbis *GainsAbisFilterer) WatchPendingOrderStored(opts *bind.WatchOpts, sink chan<- *GainsAbisPendingOrderStored) (event.Subscription, error) {

	logs, sub, err := _GainsAbis.contract.WatchLogs(opts, "PendingOrderStored")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GainsAbisPendingOrderStored)
				if err := _GainsAbis.contract.UnpackLog(event, "PendingOrderStored", log); err != nil {
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

// ParsePendingOrderStored is a log parse operation binding the contract event 0xc1f6d032e333e12d4ba1d8cdf8c4abc1bcaab7381a4eaa19a918a28f223f519d.
//
// Solidity: event PendingOrderStored(((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192),address,uint32,bool,uint8,uint32,uint16) pendingOrder)
func (_GainsAbis *GainsAbisFilterer) ParsePendingOrderStored(log types.Log) (*GainsAbisPendingOrderStored, error) {
	event := new(GainsAbisPendingOrderStored)
	if err := _GainsAbis.contract.UnpackLog(event, "PendingOrderStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GainsAbisTradeClosedIterator is returned from FilterTradeClosed and is used to iterate over the raw logs and unpacked data for TradeClosed events raised by the GainsAbis contract.
type GainsAbisTradeClosedIterator struct {
	Event *GainsAbisTradeClosed // Event containing the contract specifics and raw log

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
func (it *GainsAbisTradeClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GainsAbisTradeClosed)
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
		it.Event = new(GainsAbisTradeClosed)
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
func (it *GainsAbisTradeClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GainsAbisTradeClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GainsAbisTradeClosed represents a TradeClosed event raised by the GainsAbis contract.
type GainsAbisTradeClosed struct {
	TradeId ITradingStorageId
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTradeClosed is a free log retrieval operation binding the contract event 0xedf2f9a86d6e2127c61aaaeb10a282ee4e0aa89ea19c7db37df80fece027a493.
//
// Solidity: event TradeClosed((address,uint32) tradeId)
func (_GainsAbis *GainsAbisFilterer) FilterTradeClosed(opts *bind.FilterOpts) (*GainsAbisTradeClosedIterator, error) {

	logs, sub, err := _GainsAbis.contract.FilterLogs(opts, "TradeClosed")
	if err != nil {
		return nil, err
	}
	return &GainsAbisTradeClosedIterator{contract: _GainsAbis.contract, event: "TradeClosed", logs: logs, sub: sub}, nil
}

// WatchTradeClosed is a free log subscription operation binding the contract event 0xedf2f9a86d6e2127c61aaaeb10a282ee4e0aa89ea19c7db37df80fece027a493.
//
// Solidity: event TradeClosed((address,uint32) tradeId)
func (_GainsAbis *GainsAbisFilterer) WatchTradeClosed(opts *bind.WatchOpts, sink chan<- *GainsAbisTradeClosed) (event.Subscription, error) {

	logs, sub, err := _GainsAbis.contract.WatchLogs(opts, "TradeClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GainsAbisTradeClosed)
				if err := _GainsAbis.contract.UnpackLog(event, "TradeClosed", log); err != nil {
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

// ParseTradeClosed is a log parse operation binding the contract event 0xedf2f9a86d6e2127c61aaaeb10a282ee4e0aa89ea19c7db37df80fece027a493.
//
// Solidity: event TradeClosed((address,uint32) tradeId)
func (_GainsAbis *GainsAbisFilterer) ParseTradeClosed(log types.Log) (*GainsAbisTradeClosed, error) {
	event := new(GainsAbisTradeClosed)
	if err := _GainsAbis.contract.UnpackLog(event, "TradeClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GainsAbisTradeCollateralUpdatedIterator is returned from FilterTradeCollateralUpdated and is used to iterate over the raw logs and unpacked data for TradeCollateralUpdated events raised by the GainsAbis contract.
type GainsAbisTradeCollateralUpdatedIterator struct {
	Event *GainsAbisTradeCollateralUpdated // Event containing the contract specifics and raw log

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
func (it *GainsAbisTradeCollateralUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GainsAbisTradeCollateralUpdated)
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
		it.Event = new(GainsAbisTradeCollateralUpdated)
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
func (it *GainsAbisTradeCollateralUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GainsAbisTradeCollateralUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GainsAbisTradeCollateralUpdated represents a TradeCollateralUpdated event raised by the GainsAbis contract.
type GainsAbisTradeCollateralUpdated struct {
	TradeId          ITradingStorageId
	CollateralAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTradeCollateralUpdated is a free log retrieval operation binding the contract event 0xce228a7b1b8e239798e94cb2ba581d57501692fc1d29719a891125f1f393826d.
//
// Solidity: event TradeCollateralUpdated((address,uint32) tradeId, uint120 collateralAmount)
func (_GainsAbis *GainsAbisFilterer) FilterTradeCollateralUpdated(opts *bind.FilterOpts) (*GainsAbisTradeCollateralUpdatedIterator, error) {

	logs, sub, err := _GainsAbis.contract.FilterLogs(opts, "TradeCollateralUpdated")
	if err != nil {
		return nil, err
	}
	return &GainsAbisTradeCollateralUpdatedIterator{contract: _GainsAbis.contract, event: "TradeCollateralUpdated", logs: logs, sub: sub}, nil
}

// WatchTradeCollateralUpdated is a free log subscription operation binding the contract event 0xce228a7b1b8e239798e94cb2ba581d57501692fc1d29719a891125f1f393826d.
//
// Solidity: event TradeCollateralUpdated((address,uint32) tradeId, uint120 collateralAmount)
func (_GainsAbis *GainsAbisFilterer) WatchTradeCollateralUpdated(opts *bind.WatchOpts, sink chan<- *GainsAbisTradeCollateralUpdated) (event.Subscription, error) {

	logs, sub, err := _GainsAbis.contract.WatchLogs(opts, "TradeCollateralUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GainsAbisTradeCollateralUpdated)
				if err := _GainsAbis.contract.UnpackLog(event, "TradeCollateralUpdated", log); err != nil {
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

// ParseTradeCollateralUpdated is a log parse operation binding the contract event 0xce228a7b1b8e239798e94cb2ba581d57501692fc1d29719a891125f1f393826d.
//
// Solidity: event TradeCollateralUpdated((address,uint32) tradeId, uint120 collateralAmount)
func (_GainsAbis *GainsAbisFilterer) ParseTradeCollateralUpdated(log types.Log) (*GainsAbisTradeCollateralUpdated, error) {
	event := new(GainsAbisTradeCollateralUpdated)
	if err := _GainsAbis.contract.UnpackLog(event, "TradeCollateralUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GainsAbisTradeMaxClosingSlippagePUpdatedIterator is returned from FilterTradeMaxClosingSlippagePUpdated and is used to iterate over the raw logs and unpacked data for TradeMaxClosingSlippagePUpdated events raised by the GainsAbis contract.
type GainsAbisTradeMaxClosingSlippagePUpdatedIterator struct {
	Event *GainsAbisTradeMaxClosingSlippagePUpdated // Event containing the contract specifics and raw log

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
func (it *GainsAbisTradeMaxClosingSlippagePUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GainsAbisTradeMaxClosingSlippagePUpdated)
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
		it.Event = new(GainsAbisTradeMaxClosingSlippagePUpdated)
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
func (it *GainsAbisTradeMaxClosingSlippagePUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GainsAbisTradeMaxClosingSlippagePUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GainsAbisTradeMaxClosingSlippagePUpdated represents a TradeMaxClosingSlippagePUpdated event raised by the GainsAbis contract.
type GainsAbisTradeMaxClosingSlippagePUpdated struct {
	TradeId             ITradingStorageId
	MaxClosingSlippageP uint16
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTradeMaxClosingSlippagePUpdated is a free log retrieval operation binding the contract event 0xb34e0065c48018b4a48b78c4729fd9ffd1968c59d6532e600c4afb42ce093da1.
//
// Solidity: event TradeMaxClosingSlippagePUpdated((address,uint32) tradeId, uint16 maxClosingSlippageP)
func (_GainsAbis *GainsAbisFilterer) FilterTradeMaxClosingSlippagePUpdated(opts *bind.FilterOpts) (*GainsAbisTradeMaxClosingSlippagePUpdatedIterator, error) {

	logs, sub, err := _GainsAbis.contract.FilterLogs(opts, "TradeMaxClosingSlippagePUpdated")
	if err != nil {
		return nil, err
	}
	return &GainsAbisTradeMaxClosingSlippagePUpdatedIterator{contract: _GainsAbis.contract, event: "TradeMaxClosingSlippagePUpdated", logs: logs, sub: sub}, nil
}

// WatchTradeMaxClosingSlippagePUpdated is a free log subscription operation binding the contract event 0xb34e0065c48018b4a48b78c4729fd9ffd1968c59d6532e600c4afb42ce093da1.
//
// Solidity: event TradeMaxClosingSlippagePUpdated((address,uint32) tradeId, uint16 maxClosingSlippageP)
func (_GainsAbis *GainsAbisFilterer) WatchTradeMaxClosingSlippagePUpdated(opts *bind.WatchOpts, sink chan<- *GainsAbisTradeMaxClosingSlippagePUpdated) (event.Subscription, error) {

	logs, sub, err := _GainsAbis.contract.WatchLogs(opts, "TradeMaxClosingSlippagePUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GainsAbisTradeMaxClosingSlippagePUpdated)
				if err := _GainsAbis.contract.UnpackLog(event, "TradeMaxClosingSlippagePUpdated", log); err != nil {
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

// ParseTradeMaxClosingSlippagePUpdated is a log parse operation binding the contract event 0xb34e0065c48018b4a48b78c4729fd9ffd1968c59d6532e600c4afb42ce093da1.
//
// Solidity: event TradeMaxClosingSlippagePUpdated((address,uint32) tradeId, uint16 maxClosingSlippageP)
func (_GainsAbis *GainsAbisFilterer) ParseTradeMaxClosingSlippagePUpdated(log types.Log) (*GainsAbisTradeMaxClosingSlippagePUpdated, error) {
	event := new(GainsAbisTradeMaxClosingSlippagePUpdated)
	if err := _GainsAbis.contract.UnpackLog(event, "TradeMaxClosingSlippagePUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GainsAbisTradePositionUpdatedIterator is returned from FilterTradePositionUpdated and is used to iterate over the raw logs and unpacked data for TradePositionUpdated events raised by the GainsAbis contract.
type GainsAbisTradePositionUpdatedIterator struct {
	Event *GainsAbisTradePositionUpdated // Event containing the contract specifics and raw log

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
func (it *GainsAbisTradePositionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GainsAbisTradePositionUpdated)
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
		it.Event = new(GainsAbisTradePositionUpdated)
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
func (it *GainsAbisTradePositionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GainsAbisTradePositionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GainsAbisTradePositionUpdated represents a TradePositionUpdated event raised by the GainsAbis contract.
type GainsAbisTradePositionUpdated struct {
	TradeId           ITradingStorageId
	CollateralAmount  *big.Int
	Leverage          *big.Int
	OpenPrice         uint64
	NewTp             uint64
	NewSl             uint64
	IsPartialIncrease bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTradePositionUpdated is a free log retrieval operation binding the contract event 0x3890801c6f1ea37e9eb6aaae6cb1e57eeb8fe67e3af47b2533ff1fc4d1031ede.
//
// Solidity: event TradePositionUpdated((address,uint32) tradeId, uint120 collateralAmount, uint24 leverage, uint64 openPrice, uint64 newTp, uint64 newSl, bool isPartialIncrease)
func (_GainsAbis *GainsAbisFilterer) FilterTradePositionUpdated(opts *bind.FilterOpts) (*GainsAbisTradePositionUpdatedIterator, error) {

	logs, sub, err := _GainsAbis.contract.FilterLogs(opts, "TradePositionUpdated")
	if err != nil {
		return nil, err
	}
	return &GainsAbisTradePositionUpdatedIterator{contract: _GainsAbis.contract, event: "TradePositionUpdated", logs: logs, sub: sub}, nil
}

// WatchTradePositionUpdated is a free log subscription operation binding the contract event 0x3890801c6f1ea37e9eb6aaae6cb1e57eeb8fe67e3af47b2533ff1fc4d1031ede.
//
// Solidity: event TradePositionUpdated((address,uint32) tradeId, uint120 collateralAmount, uint24 leverage, uint64 openPrice, uint64 newTp, uint64 newSl, bool isPartialIncrease)
func (_GainsAbis *GainsAbisFilterer) WatchTradePositionUpdated(opts *bind.WatchOpts, sink chan<- *GainsAbisTradePositionUpdated) (event.Subscription, error) {

	logs, sub, err := _GainsAbis.contract.WatchLogs(opts, "TradePositionUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GainsAbisTradePositionUpdated)
				if err := _GainsAbis.contract.UnpackLog(event, "TradePositionUpdated", log); err != nil {
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

// ParseTradePositionUpdated is a log parse operation binding the contract event 0x3890801c6f1ea37e9eb6aaae6cb1e57eeb8fe67e3af47b2533ff1fc4d1031ede.
//
// Solidity: event TradePositionUpdated((address,uint32) tradeId, uint120 collateralAmount, uint24 leverage, uint64 openPrice, uint64 newTp, uint64 newSl, bool isPartialIncrease)
func (_GainsAbis *GainsAbisFilterer) ParseTradePositionUpdated(log types.Log) (*GainsAbisTradePositionUpdated, error) {
	event := new(GainsAbisTradePositionUpdated)
	if err := _GainsAbis.contract.UnpackLog(event, "TradePositionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GainsAbisTradeSlUpdatedIterator is returned from FilterTradeSlUpdated and is used to iterate over the raw logs and unpacked data for TradeSlUpdated events raised by the GainsAbis contract.
type GainsAbisTradeSlUpdatedIterator struct {
	Event *GainsAbisTradeSlUpdated // Event containing the contract specifics and raw log

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
func (it *GainsAbisTradeSlUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GainsAbisTradeSlUpdated)
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
		it.Event = new(GainsAbisTradeSlUpdated)
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
func (it *GainsAbisTradeSlUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GainsAbisTradeSlUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GainsAbisTradeSlUpdated represents a TradeSlUpdated event raised by the GainsAbis contract.
type GainsAbisTradeSlUpdated struct {
	TradeId ITradingStorageId
	NewSl   uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTradeSlUpdated is a free log retrieval operation binding the contract event 0x38f5d5d40d9c4a41aa03d21461f1b07aa6b4ef035fb9d21f02d53a82c712a002.
//
// Solidity: event TradeSlUpdated((address,uint32) tradeId, uint64 newSl)
func (_GainsAbis *GainsAbisFilterer) FilterTradeSlUpdated(opts *bind.FilterOpts) (*GainsAbisTradeSlUpdatedIterator, error) {

	logs, sub, err := _GainsAbis.contract.FilterLogs(opts, "TradeSlUpdated")
	if err != nil {
		return nil, err
	}
	return &GainsAbisTradeSlUpdatedIterator{contract: _GainsAbis.contract, event: "TradeSlUpdated", logs: logs, sub: sub}, nil
}

// WatchTradeSlUpdated is a free log subscription operation binding the contract event 0x38f5d5d40d9c4a41aa03d21461f1b07aa6b4ef035fb9d21f02d53a82c712a002.
//
// Solidity: event TradeSlUpdated((address,uint32) tradeId, uint64 newSl)
func (_GainsAbis *GainsAbisFilterer) WatchTradeSlUpdated(opts *bind.WatchOpts, sink chan<- *GainsAbisTradeSlUpdated) (event.Subscription, error) {

	logs, sub, err := _GainsAbis.contract.WatchLogs(opts, "TradeSlUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GainsAbisTradeSlUpdated)
				if err := _GainsAbis.contract.UnpackLog(event, "TradeSlUpdated", log); err != nil {
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

// ParseTradeSlUpdated is a log parse operation binding the contract event 0x38f5d5d40d9c4a41aa03d21461f1b07aa6b4ef035fb9d21f02d53a82c712a002.
//
// Solidity: event TradeSlUpdated((address,uint32) tradeId, uint64 newSl)
func (_GainsAbis *GainsAbisFilterer) ParseTradeSlUpdated(log types.Log) (*GainsAbisTradeSlUpdated, error) {
	event := new(GainsAbisTradeSlUpdated)
	if err := _GainsAbis.contract.UnpackLog(event, "TradeSlUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GainsAbisTradeStoredIterator is returned from FilterTradeStored and is used to iterate over the raw logs and unpacked data for TradeStored events raised by the GainsAbis contract.
type GainsAbisTradeStoredIterator struct {
	Event *GainsAbisTradeStored // Event containing the contract specifics and raw log

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
func (it *GainsAbisTradeStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GainsAbisTradeStored)
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
		it.Event = new(GainsAbisTradeStored)
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
func (it *GainsAbisTradeStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GainsAbisTradeStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GainsAbisTradeStored represents a TradeStored event raised by the GainsAbis contract.
type GainsAbisTradeStored struct {
	Trade             ITradingStorageTrade
	TradeInfo         ITradingStorageTradeInfo
	LiquidationParams IPairsStorageGroupLiquidationParams
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTradeStored is a free log retrieval operation binding the contract event 0xb4c8599e992aeeb8f86e02eaee1061646ddf10a354b91f1daa776eb4595387a3.
//
// Solidity: event TradeStored((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192) trade, (uint32,uint32,uint32,uint16,uint48,uint48,uint8,uint32,uint8) tradeInfo, (uint40,uint40,uint40,uint24,uint24) liquidationParams)
func (_GainsAbis *GainsAbisFilterer) FilterTradeStored(opts *bind.FilterOpts) (*GainsAbisTradeStoredIterator, error) {

	logs, sub, err := _GainsAbis.contract.FilterLogs(opts, "TradeStored")
	if err != nil {
		return nil, err
	}
	return &GainsAbisTradeStoredIterator{contract: _GainsAbis.contract, event: "TradeStored", logs: logs, sub: sub}, nil
}

// WatchTradeStored is a free log subscription operation binding the contract event 0xb4c8599e992aeeb8f86e02eaee1061646ddf10a354b91f1daa776eb4595387a3.
//
// Solidity: event TradeStored((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192) trade, (uint32,uint32,uint32,uint16,uint48,uint48,uint8,uint32,uint8) tradeInfo, (uint40,uint40,uint40,uint24,uint24) liquidationParams)
func (_GainsAbis *GainsAbisFilterer) WatchTradeStored(opts *bind.WatchOpts, sink chan<- *GainsAbisTradeStored) (event.Subscription, error) {

	logs, sub, err := _GainsAbis.contract.WatchLogs(opts, "TradeStored")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GainsAbisTradeStored)
				if err := _GainsAbis.contract.UnpackLog(event, "TradeStored", log); err != nil {
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

// ParseTradeStored is a log parse operation binding the contract event 0xb4c8599e992aeeb8f86e02eaee1061646ddf10a354b91f1daa776eb4595387a3.
//
// Solidity: event TradeStored((address,uint32,uint16,uint24,bool,bool,uint8,uint8,uint120,uint64,uint64,uint64,uint192) trade, (uint32,uint32,uint32,uint16,uint48,uint48,uint8,uint32,uint8) tradeInfo, (uint40,uint40,uint40,uint24,uint24) liquidationParams)
func (_GainsAbis *GainsAbisFilterer) ParseTradeStored(log types.Log) (*GainsAbisTradeStored, error) {
	event := new(GainsAbisTradeStored)
	if err := _GainsAbis.contract.UnpackLog(event, "TradeStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GainsAbisTradeTpUpdatedIterator is returned from FilterTradeTpUpdated and is used to iterate over the raw logs and unpacked data for TradeTpUpdated events raised by the GainsAbis contract.
type GainsAbisTradeTpUpdatedIterator struct {
	Event *GainsAbisTradeTpUpdated // Event containing the contract specifics and raw log

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
func (it *GainsAbisTradeTpUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GainsAbisTradeTpUpdated)
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
		it.Event = new(GainsAbisTradeTpUpdated)
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
func (it *GainsAbisTradeTpUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GainsAbisTradeTpUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GainsAbisTradeTpUpdated represents a TradeTpUpdated event raised by the GainsAbis contract.
type GainsAbisTradeTpUpdated struct {
	TradeId ITradingStorageId
	NewTp   uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTradeTpUpdated is a free log retrieval operation binding the contract event 0x3d045f25e6a6757ae5ca79ce5d28d84d69713804353a02c521d6a5352c0f9e20.
//
// Solidity: event TradeTpUpdated((address,uint32) tradeId, uint64 newTp)
func (_GainsAbis *GainsAbisFilterer) FilterTradeTpUpdated(opts *bind.FilterOpts) (*GainsAbisTradeTpUpdatedIterator, error) {

	logs, sub, err := _GainsAbis.contract.FilterLogs(opts, "TradeTpUpdated")
	if err != nil {
		return nil, err
	}
	return &GainsAbisTradeTpUpdatedIterator{contract: _GainsAbis.contract, event: "TradeTpUpdated", logs: logs, sub: sub}, nil
}

// WatchTradeTpUpdated is a free log subscription operation binding the contract event 0x3d045f25e6a6757ae5ca79ce5d28d84d69713804353a02c521d6a5352c0f9e20.
//
// Solidity: event TradeTpUpdated((address,uint32) tradeId, uint64 newTp)
func (_GainsAbis *GainsAbisFilterer) WatchTradeTpUpdated(opts *bind.WatchOpts, sink chan<- *GainsAbisTradeTpUpdated) (event.Subscription, error) {

	logs, sub, err := _GainsAbis.contract.WatchLogs(opts, "TradeTpUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GainsAbisTradeTpUpdated)
				if err := _GainsAbis.contract.UnpackLog(event, "TradeTpUpdated", log); err != nil {
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

// ParseTradeTpUpdated is a log parse operation binding the contract event 0x3d045f25e6a6757ae5ca79ce5d28d84d69713804353a02c521d6a5352c0f9e20.
//
// Solidity: event TradeTpUpdated((address,uint32) tradeId, uint64 newTp)
func (_GainsAbis *GainsAbisFilterer) ParseTradeTpUpdated(log types.Log) (*GainsAbisTradeTpUpdated, error) {
	event := new(GainsAbisTradeTpUpdated)
	if err := _GainsAbis.contract.UnpackLog(event, "TradeTpUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GainsAbisTradingActivatedUpdatedIterator is returned from FilterTradingActivatedUpdated and is used to iterate over the raw logs and unpacked data for TradingActivatedUpdated events raised by the GainsAbis contract.
type GainsAbisTradingActivatedUpdatedIterator struct {
	Event *GainsAbisTradingActivatedUpdated // Event containing the contract specifics and raw log

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
func (it *GainsAbisTradingActivatedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GainsAbisTradingActivatedUpdated)
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
		it.Event = new(GainsAbisTradingActivatedUpdated)
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
func (it *GainsAbisTradingActivatedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GainsAbisTradingActivatedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GainsAbisTradingActivatedUpdated represents a TradingActivatedUpdated event raised by the GainsAbis contract.
type GainsAbisTradingActivatedUpdated struct {
	Activated uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTradingActivatedUpdated is a free log retrieval operation binding the contract event 0x4b502c3b75c299352edc7887297ae0f7c401ed654650a4c0e663458b6ed75fe4.
//
// Solidity: event TradingActivatedUpdated(uint8 activated)
func (_GainsAbis *GainsAbisFilterer) FilterTradingActivatedUpdated(opts *bind.FilterOpts) (*GainsAbisTradingActivatedUpdatedIterator, error) {

	logs, sub, err := _GainsAbis.contract.FilterLogs(opts, "TradingActivatedUpdated")
	if err != nil {
		return nil, err
	}
	return &GainsAbisTradingActivatedUpdatedIterator{contract: _GainsAbis.contract, event: "TradingActivatedUpdated", logs: logs, sub: sub}, nil
}

// WatchTradingActivatedUpdated is a free log subscription operation binding the contract event 0x4b502c3b75c299352edc7887297ae0f7c401ed654650a4c0e663458b6ed75fe4.
//
// Solidity: event TradingActivatedUpdated(uint8 activated)
func (_GainsAbis *GainsAbisFilterer) WatchTradingActivatedUpdated(opts *bind.WatchOpts, sink chan<- *GainsAbisTradingActivatedUpdated) (event.Subscription, error) {

	logs, sub, err := _GainsAbis.contract.WatchLogs(opts, "TradingActivatedUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GainsAbisTradingActivatedUpdated)
				if err := _GainsAbis.contract.UnpackLog(event, "TradingActivatedUpdated", log); err != nil {
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

// ParseTradingActivatedUpdated is a log parse operation binding the contract event 0x4b502c3b75c299352edc7887297ae0f7c401ed654650a4c0e663458b6ed75fe4.
//
// Solidity: event TradingActivatedUpdated(uint8 activated)
func (_GainsAbis *GainsAbisFilterer) ParseTradingActivatedUpdated(log types.Log) (*GainsAbisTradingActivatedUpdated, error) {
	event := new(GainsAbisTradingActivatedUpdated)
	if err := _GainsAbis.contract.UnpackLog(event, "TradingActivatedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
