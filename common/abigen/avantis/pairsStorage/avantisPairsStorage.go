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

// IPairStorageBackupFeed is an auto generated low-level Go binding around an user-defined struct.
type IPairStorageBackupFeed struct {
	MaxDeviationP *big.Int
	FeedId        common.Address
}

// IPairStorageFee is an auto generated low-level Go binding around an user-defined struct.
type IPairStorageFee struct {
	Name           string
	OpenFeeP       *big.Int
	CloseFeeP      *big.Int
	LimitOrderFeeP *big.Int
	MinLevPosUSDC  *big.Int
}

// IPairStorageFeed is an auto generated low-level Go binding around an user-defined struct.
type IPairStorageFeed struct {
	MaxDeviationP *big.Int
	FeedId        [32]byte
}

// IPairStorageGroup is an auto generated low-level Go binding around an user-defined struct.
type IPairStorageGroup struct {
	Name             string
	MinLeverage      *big.Int
	MaxLeverage      *big.Int
	MaxOpenInterestP *big.Int
}

// IPairStoragePair is an auto generated low-level Go binding around an user-defined struct.
type IPairStoragePair struct {
	From                       string
	To                         string
	Feed                       IPairStorageFeed
	BackupFeed                 IPairStorageBackupFeed
	SpreadP                    *big.Int
	PriceImpactMultiplier      *big.Int
	SkewImpactMultiplier       *big.Int
	GroupIndex                 *big.Int
	FeeIndex                   *big.Int
	GroupOpenInterestPecentage *big.Int
	MaxWalletOI                *big.Int
}

// IPairStorageSkewFee is an auto generated low-level Go binding around an user-defined struct.
type IPairStorageSkewFee struct {
	EqParams [10][2]*big.Int
}

// AvantisMulticallAbisMetaData contains all meta data concerning the AvantisMulticallAbis contract.
var AvantisMulticallAbisMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addFee\",\"inputs\":[{\"name\":\"_fee\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.Fee\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"openFeeP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"closeFeeP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limitOrderFeeP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minLevPosUSDC\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addGroup\",\"inputs\":[{\"name\":\"_group\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.Group\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"minLeverage\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxLeverage\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxOpenInterestP\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addPair\",\"inputs\":[{\"name\":\"_pair\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.Pair\",\"components\":[{\"name\":\"from\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"to\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"feed\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.Feed\",\"components\":[{\"name\":\"maxDeviationP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feedId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"backupFeed\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.BackupFeed\",\"components\":[{\"name\":\"maxDeviationP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feedId\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"spreadP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priceImpactMultiplier\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"skewImpactMultiplier\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"groupIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"groupOpenInterestPecentage\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxWalletOI\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addSkewOpenFees\",\"inputs\":[{\"name\":\"_skewFee\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.SkewFee\",\"components\":[{\"name\":\"eqParams\",\"type\":\"int256[2][10]\",\"internalType\":\"int256[2][10]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blockOILimit\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentOrderId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delistPair\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fees\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"openFeeP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"closeFeeP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limitOrderFeeP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minLevPosUSDC\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feesCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"groupMaxOI\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"groupOI\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"groupOIs\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"groups\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"minLeverage\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxLeverage\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxOpenInterestP\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"groupsCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"guaranteedSlEnabled\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_storageT\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_currentOrderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPairListed\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isUSDCAligned\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lossProtection\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lossProtectionMultiplier\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tier\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxWalletOI\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairBackupFeed\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.BackupFeed\",\"components\":[{\"name\":\"maxDeviationP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feedId\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairCloseFeeP\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairFeed\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.Feed\",\"components\":[{\"name\":\"maxDeviationP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feedId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairGroupIndex\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairJob\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pairLimitOrderFeeP\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairMaxLeverage\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairMaxOI\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairMinLevPosUSDC\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairMinLeverage\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairOpenFeeP\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_leveragedPosition\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_buy\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairPriceImpactMultiplier\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairSkewImpactMultiplier\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairSpreadP\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairs\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"from\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"to\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"feed\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.Feed\",\"components\":[{\"name\":\"maxDeviationP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feedId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"backupFeed\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.BackupFeed\",\"components\":[{\"name\":\"maxDeviationP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feedId\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"spreadP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priceImpactMultiplier\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"skewImpactMultiplier\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"groupIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"groupOpenInterestPecentage\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxWalletOI\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairsBackend\",\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.Pair\",\"components\":[{\"name\":\"from\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"to\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"feed\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.Feed\",\"components\":[{\"name\":\"maxDeviationP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feedId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"backupFeed\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.BackupFeed\",\"components\":[{\"name\":\"maxDeviationP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feedId\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"spreadP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priceImpactMultiplier\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"skewImpactMultiplier\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"groupIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"groupOpenInterestPecentage\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxWalletOI\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.Group\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"minLeverage\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxLeverage\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxOpenInterestP\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.Fee\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"openFeeP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"closeFeeP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limitOrderFeeP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minLevPosUSDC\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairsCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setBlockOILImits\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_limits\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"skewedFeesCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"storageT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITradingStorage\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"udpateSkewOpenFees\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_skewFee\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.SkewFee\",\"components\":[{\"name\":\"eqParams\",\"type\":\"int256[2][10]\",\"internalType\":\"int256[2][10]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateFee\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.Fee\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"openFeeP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"closeFeeP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limitOrderFeeP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minLevPosUSDC\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateGroup\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_group\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.Group\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"minLeverage\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxLeverage\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxOpenInterestP\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateGroupOI\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_long\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_increase\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateLossProtectionMultiplier\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tier\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_multiplierPercent\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePair\",\"inputs\":[{\"name\":\"_pairIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_pair\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.Pair\",\"components\":[{\"name\":\"from\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"to\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"feed\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.Feed\",\"components\":[{\"name\":\"maxDeviationP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feedId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"backupFeed\",\"type\":\"tuple\",\"internalType\":\"structIPairStorage.BackupFeed\",\"components\":[{\"name\":\"maxDeviationP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feedId\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"spreadP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priceImpactMultiplier\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"skewImpactMultiplier\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"groupIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"groupOpenInterestPecentage\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxWalletOI\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BlockOILimitsSet\",\"inputs\":[{\"name\":\"pairIndex\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"limits\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeAdded\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeUpdated\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GroupAdded\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GroupUpdated\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LossProtectionAdded\",\"inputs\":[{\"name\":\"pairIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"tier\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"multiplier\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderLimitsSet\",\"inputs\":[{\"name\":\"pairIndex\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"limits\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PairAdded\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"to\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PairUpdated\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SkewFeeAdded\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SkewFeeUpdated\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
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

// BlockOILimit is a free data retrieval call binding the contract method 0x9ae6b9e0.
//
// Solidity: function blockOILimit(uint256 ) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) BlockOILimit(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "blockOILimit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockOILimit is a free data retrieval call binding the contract method 0x9ae6b9e0.
//
// Solidity: function blockOILimit(uint256 ) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) BlockOILimit(arg0 *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.BlockOILimit(&_AvantisMulticallAbis.CallOpts, arg0)
}

// BlockOILimit is a free data retrieval call binding the contract method 0x9ae6b9e0.
//
// Solidity: function blockOILimit(uint256 ) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) BlockOILimit(arg0 *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.BlockOILimit(&_AvantisMulticallAbis.CallOpts, arg0)
}

// CurrentOrderId is a free data retrieval call binding the contract method 0x925931fc.
//
// Solidity: function currentOrderId() view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) CurrentOrderId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "currentOrderId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentOrderId is a free data retrieval call binding the contract method 0x925931fc.
//
// Solidity: function currentOrderId() view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) CurrentOrderId() (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.CurrentOrderId(&_AvantisMulticallAbis.CallOpts)
}

// CurrentOrderId is a free data retrieval call binding the contract method 0x925931fc.
//
// Solidity: function currentOrderId() view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) CurrentOrderId() (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.CurrentOrderId(&_AvantisMulticallAbis.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x4acc79ed.
//
// Solidity: function fees(uint256 ) view returns(string name, uint256 openFeeP, uint256 closeFeeP, uint256 limitOrderFeeP, uint256 minLevPosUSDC)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) Fees(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name           string
	OpenFeeP       *big.Int
	CloseFeeP      *big.Int
	LimitOrderFeeP *big.Int
	MinLevPosUSDC  *big.Int
}, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "fees", arg0)

	outstruct := new(struct {
		Name           string
		OpenFeeP       *big.Int
		CloseFeeP      *big.Int
		LimitOrderFeeP *big.Int
		MinLevPosUSDC  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.OpenFeeP = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CloseFeeP = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LimitOrderFeeP = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MinLevPosUSDC = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Fees is a free data retrieval call binding the contract method 0x4acc79ed.
//
// Solidity: function fees(uint256 ) view returns(string name, uint256 openFeeP, uint256 closeFeeP, uint256 limitOrderFeeP, uint256 minLevPosUSDC)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) Fees(arg0 *big.Int) (struct {
	Name           string
	OpenFeeP       *big.Int
	CloseFeeP      *big.Int
	LimitOrderFeeP *big.Int
	MinLevPosUSDC  *big.Int
}, error) {
	return _AvantisMulticallAbis.Contract.Fees(&_AvantisMulticallAbis.CallOpts, arg0)
}

// Fees is a free data retrieval call binding the contract method 0x4acc79ed.
//
// Solidity: function fees(uint256 ) view returns(string name, uint256 openFeeP, uint256 closeFeeP, uint256 limitOrderFeeP, uint256 minLevPosUSDC)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) Fees(arg0 *big.Int) (struct {
	Name           string
	OpenFeeP       *big.Int
	CloseFeeP      *big.Int
	LimitOrderFeeP *big.Int
	MinLevPosUSDC  *big.Int
}, error) {
	return _AvantisMulticallAbis.Contract.Fees(&_AvantisMulticallAbis.CallOpts, arg0)
}

// FeesCount is a free data retrieval call binding the contract method 0x658de48a.
//
// Solidity: function feesCount() view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) FeesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "feesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeesCount is a free data retrieval call binding the contract method 0x658de48a.
//
// Solidity: function feesCount() view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) FeesCount() (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.FeesCount(&_AvantisMulticallAbis.CallOpts)
}

// FeesCount is a free data retrieval call binding the contract method 0x658de48a.
//
// Solidity: function feesCount() view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) FeesCount() (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.FeesCount(&_AvantisMulticallAbis.CallOpts)
}

// GroupMaxOI is a free data retrieval call binding the contract method 0x8c9a0ea4.
//
// Solidity: function groupMaxOI(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) GroupMaxOI(opts *bind.CallOpts, _pairIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "groupMaxOI", _pairIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GroupMaxOI is a free data retrieval call binding the contract method 0x8c9a0ea4.
//
// Solidity: function groupMaxOI(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) GroupMaxOI(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GroupMaxOI(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// GroupMaxOI is a free data retrieval call binding the contract method 0x8c9a0ea4.
//
// Solidity: function groupMaxOI(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) GroupMaxOI(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GroupMaxOI(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// GroupOI is a free data retrieval call binding the contract method 0x5e4f8f59.
//
// Solidity: function groupOI(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) GroupOI(opts *bind.CallOpts, _pairIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "groupOI", _pairIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GroupOI is a free data retrieval call binding the contract method 0x5e4f8f59.
//
// Solidity: function groupOI(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) GroupOI(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GroupOI(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// GroupOI is a free data retrieval call binding the contract method 0x5e4f8f59.
//
// Solidity: function groupOI(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) GroupOI(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GroupOI(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// GroupOIs is a free data retrieval call binding the contract method 0xa1ead064.
//
// Solidity: function groupOIs(uint256 , uint256 ) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) GroupOIs(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "groupOIs", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GroupOIs is a free data retrieval call binding the contract method 0xa1ead064.
//
// Solidity: function groupOIs(uint256 , uint256 ) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) GroupOIs(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GroupOIs(&_AvantisMulticallAbis.CallOpts, arg0, arg1)
}

// GroupOIs is a free data retrieval call binding the contract method 0xa1ead064.
//
// Solidity: function groupOIs(uint256 , uint256 ) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) GroupOIs(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GroupOIs(&_AvantisMulticallAbis.CallOpts, arg0, arg1)
}

// Groups is a free data retrieval call binding the contract method 0x96324bd4.
//
// Solidity: function groups(uint256 ) view returns(string name, uint256 minLeverage, uint256 maxLeverage, uint256 maxOpenInterestP)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) Groups(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name             string
	MinLeverage      *big.Int
	MaxLeverage      *big.Int
	MaxOpenInterestP *big.Int
}, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "groups", arg0)

	outstruct := new(struct {
		Name             string
		MinLeverage      *big.Int
		MaxLeverage      *big.Int
		MaxOpenInterestP *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.MinLeverage = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MaxLeverage = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxOpenInterestP = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Groups is a free data retrieval call binding the contract method 0x96324bd4.
//
// Solidity: function groups(uint256 ) view returns(string name, uint256 minLeverage, uint256 maxLeverage, uint256 maxOpenInterestP)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) Groups(arg0 *big.Int) (struct {
	Name             string
	MinLeverage      *big.Int
	MaxLeverage      *big.Int
	MaxOpenInterestP *big.Int
}, error) {
	return _AvantisMulticallAbis.Contract.Groups(&_AvantisMulticallAbis.CallOpts, arg0)
}

// Groups is a free data retrieval call binding the contract method 0x96324bd4.
//
// Solidity: function groups(uint256 ) view returns(string name, uint256 minLeverage, uint256 maxLeverage, uint256 maxOpenInterestP)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) Groups(arg0 *big.Int) (struct {
	Name             string
	MinLeverage      *big.Int
	MaxLeverage      *big.Int
	MaxOpenInterestP *big.Int
}, error) {
	return _AvantisMulticallAbis.Contract.Groups(&_AvantisMulticallAbis.CallOpts, arg0)
}

// GroupsCount is a free data retrieval call binding the contract method 0x885e2750.
//
// Solidity: function groupsCount() view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) GroupsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "groupsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GroupsCount is a free data retrieval call binding the contract method 0x885e2750.
//
// Solidity: function groupsCount() view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) GroupsCount() (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GroupsCount(&_AvantisMulticallAbis.CallOpts)
}

// GroupsCount is a free data retrieval call binding the contract method 0x885e2750.
//
// Solidity: function groupsCount() view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) GroupsCount() (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.GroupsCount(&_AvantisMulticallAbis.CallOpts)
}

// GuaranteedSlEnabled is a free data retrieval call binding the contract method 0x24abd3fb.
//
// Solidity: function guaranteedSlEnabled(uint256 _pairIndex) view returns(bool)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) GuaranteedSlEnabled(opts *bind.CallOpts, _pairIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "guaranteedSlEnabled", _pairIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GuaranteedSlEnabled is a free data retrieval call binding the contract method 0x24abd3fb.
//
// Solidity: function guaranteedSlEnabled(uint256 _pairIndex) view returns(bool)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) GuaranteedSlEnabled(_pairIndex *big.Int) (bool, error) {
	return _AvantisMulticallAbis.Contract.GuaranteedSlEnabled(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// GuaranteedSlEnabled is a free data retrieval call binding the contract method 0x24abd3fb.
//
// Solidity: function guaranteedSlEnabled(uint256 _pairIndex) view returns(bool)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) GuaranteedSlEnabled(_pairIndex *big.Int) (bool, error) {
	return _AvantisMulticallAbis.Contract.GuaranteedSlEnabled(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// IsPairListed is a free data retrieval call binding the contract method 0x1628bfeb.
//
// Solidity: function isPairListed(string , string ) view returns(bool)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) IsPairListed(opts *bind.CallOpts, arg0 string, arg1 string) (bool, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "isPairListed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPairListed is a free data retrieval call binding the contract method 0x1628bfeb.
//
// Solidity: function isPairListed(string , string ) view returns(bool)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) IsPairListed(arg0 string, arg1 string) (bool, error) {
	return _AvantisMulticallAbis.Contract.IsPairListed(&_AvantisMulticallAbis.CallOpts, arg0, arg1)
}

// IsPairListed is a free data retrieval call binding the contract method 0x1628bfeb.
//
// Solidity: function isPairListed(string , string ) view returns(bool)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) IsPairListed(arg0 string, arg1 string) (bool, error) {
	return _AvantisMulticallAbis.Contract.IsPairListed(&_AvantisMulticallAbis.CallOpts, arg0, arg1)
}

// IsUSDCAligned is a free data retrieval call binding the contract method 0x6dc37de3.
//
// Solidity: function isUSDCAligned(uint256 _pairIndex) view returns(bool)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) IsUSDCAligned(opts *bind.CallOpts, _pairIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "isUSDCAligned", _pairIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUSDCAligned is a free data retrieval call binding the contract method 0x6dc37de3.
//
// Solidity: function isUSDCAligned(uint256 _pairIndex) view returns(bool)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) IsUSDCAligned(_pairIndex *big.Int) (bool, error) {
	return _AvantisMulticallAbis.Contract.IsUSDCAligned(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// IsUSDCAligned is a free data retrieval call binding the contract method 0x6dc37de3.
//
// Solidity: function isUSDCAligned(uint256 _pairIndex) view returns(bool)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) IsUSDCAligned(_pairIndex *big.Int) (bool, error) {
	return _AvantisMulticallAbis.Contract.IsUSDCAligned(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// LossProtection is a free data retrieval call binding the contract method 0x15b40d91.
//
// Solidity: function lossProtection(uint256 , uint256 ) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) LossProtection(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "lossProtection", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LossProtection is a free data retrieval call binding the contract method 0x15b40d91.
//
// Solidity: function lossProtection(uint256 , uint256 ) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) LossProtection(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.LossProtection(&_AvantisMulticallAbis.CallOpts, arg0, arg1)
}

// LossProtection is a free data retrieval call binding the contract method 0x15b40d91.
//
// Solidity: function lossProtection(uint256 , uint256 ) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) LossProtection(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.LossProtection(&_AvantisMulticallAbis.CallOpts, arg0, arg1)
}

// LossProtectionMultiplier is a free data retrieval call binding the contract method 0x2b1cbe9d.
//
// Solidity: function lossProtectionMultiplier(uint256 _pairIndex, uint256 _tier) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) LossProtectionMultiplier(opts *bind.CallOpts, _pairIndex *big.Int, _tier *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "lossProtectionMultiplier", _pairIndex, _tier)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LossProtectionMultiplier is a free data retrieval call binding the contract method 0x2b1cbe9d.
//
// Solidity: function lossProtectionMultiplier(uint256 _pairIndex, uint256 _tier) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) LossProtectionMultiplier(_pairIndex *big.Int, _tier *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.LossProtectionMultiplier(&_AvantisMulticallAbis.CallOpts, _pairIndex, _tier)
}

// LossProtectionMultiplier is a free data retrieval call binding the contract method 0x2b1cbe9d.
//
// Solidity: function lossProtectionMultiplier(uint256 _pairIndex, uint256 _tier) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) LossProtectionMultiplier(_pairIndex *big.Int, _tier *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.LossProtectionMultiplier(&_AvantisMulticallAbis.CallOpts, _pairIndex, _tier)
}

// MaxWalletOI is a free data retrieval call binding the contract method 0x56e5536c.
//
// Solidity: function maxWalletOI(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) MaxWalletOI(opts *bind.CallOpts, _pairIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "maxWalletOI", _pairIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWalletOI is a free data retrieval call binding the contract method 0x56e5536c.
//
// Solidity: function maxWalletOI(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) MaxWalletOI(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.MaxWalletOI(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// MaxWalletOI is a free data retrieval call binding the contract method 0x56e5536c.
//
// Solidity: function maxWalletOI(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) MaxWalletOI(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.MaxWalletOI(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairBackupFeed is a free data retrieval call binding the contract method 0x685e905b.
//
// Solidity: function pairBackupFeed(uint256 _pairIndex) view returns((uint256,address))
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) PairBackupFeed(opts *bind.CallOpts, _pairIndex *big.Int) (IPairStorageBackupFeed, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "pairBackupFeed", _pairIndex)

	if err != nil {
		return *new(IPairStorageBackupFeed), err
	}

	out0 := *abi.ConvertType(out[0], new(IPairStorageBackupFeed)).(*IPairStorageBackupFeed)

	return out0, err

}

// PairBackupFeed is a free data retrieval call binding the contract method 0x685e905b.
//
// Solidity: function pairBackupFeed(uint256 _pairIndex) view returns((uint256,address))
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) PairBackupFeed(_pairIndex *big.Int) (IPairStorageBackupFeed, error) {
	return _AvantisMulticallAbis.Contract.PairBackupFeed(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairBackupFeed is a free data retrieval call binding the contract method 0x685e905b.
//
// Solidity: function pairBackupFeed(uint256 _pairIndex) view returns((uint256,address))
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) PairBackupFeed(_pairIndex *big.Int) (IPairStorageBackupFeed, error) {
	return _AvantisMulticallAbis.Contract.PairBackupFeed(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairCloseFeeP is a free data retrieval call binding the contract method 0x836a341a.
//
// Solidity: function pairCloseFeeP(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) PairCloseFeeP(opts *bind.CallOpts, _pairIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "pairCloseFeeP", _pairIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairCloseFeeP is a free data retrieval call binding the contract method 0x836a341a.
//
// Solidity: function pairCloseFeeP(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) PairCloseFeeP(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairCloseFeeP(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairCloseFeeP is a free data retrieval call binding the contract method 0x836a341a.
//
// Solidity: function pairCloseFeeP(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) PairCloseFeeP(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairCloseFeeP(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairFeed is a free data retrieval call binding the contract method 0x1eaa005c.
//
// Solidity: function pairFeed(uint256 _pairIndex) view returns((uint256,bytes32))
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) PairFeed(opts *bind.CallOpts, _pairIndex *big.Int) (IPairStorageFeed, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "pairFeed", _pairIndex)

	if err != nil {
		return *new(IPairStorageFeed), err
	}

	out0 := *abi.ConvertType(out[0], new(IPairStorageFeed)).(*IPairStorageFeed)

	return out0, err

}

// PairFeed is a free data retrieval call binding the contract method 0x1eaa005c.
//
// Solidity: function pairFeed(uint256 _pairIndex) view returns((uint256,bytes32))
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) PairFeed(_pairIndex *big.Int) (IPairStorageFeed, error) {
	return _AvantisMulticallAbis.Contract.PairFeed(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairFeed is a free data retrieval call binding the contract method 0x1eaa005c.
//
// Solidity: function pairFeed(uint256 _pairIndex) view returns((uint256,bytes32))
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) PairFeed(_pairIndex *big.Int) (IPairStorageFeed, error) {
	return _AvantisMulticallAbis.Contract.PairFeed(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairGroupIndex is a free data retrieval call binding the contract method 0x3c44d0a3.
//
// Solidity: function pairGroupIndex(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) PairGroupIndex(opts *bind.CallOpts, _pairIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "pairGroupIndex", _pairIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairGroupIndex is a free data retrieval call binding the contract method 0x3c44d0a3.
//
// Solidity: function pairGroupIndex(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) PairGroupIndex(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairGroupIndex(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairGroupIndex is a free data retrieval call binding the contract method 0x3c44d0a3.
//
// Solidity: function pairGroupIndex(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) PairGroupIndex(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairGroupIndex(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairLimitOrderFeeP is a free data retrieval call binding the contract method 0x649a6488.
//
// Solidity: function pairLimitOrderFeeP(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) PairLimitOrderFeeP(opts *bind.CallOpts, _pairIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "pairLimitOrderFeeP", _pairIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairLimitOrderFeeP is a free data retrieval call binding the contract method 0x649a6488.
//
// Solidity: function pairLimitOrderFeeP(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) PairLimitOrderFeeP(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairLimitOrderFeeP(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairLimitOrderFeeP is a free data retrieval call binding the contract method 0x649a6488.
//
// Solidity: function pairLimitOrderFeeP(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) PairLimitOrderFeeP(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairLimitOrderFeeP(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairMaxLeverage is a free data retrieval call binding the contract method 0x281b693c.
//
// Solidity: function pairMaxLeverage(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) PairMaxLeverage(opts *bind.CallOpts, _pairIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "pairMaxLeverage", _pairIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairMaxLeverage is a free data retrieval call binding the contract method 0x281b693c.
//
// Solidity: function pairMaxLeverage(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) PairMaxLeverage(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairMaxLeverage(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairMaxLeverage is a free data retrieval call binding the contract method 0x281b693c.
//
// Solidity: function pairMaxLeverage(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) PairMaxLeverage(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairMaxLeverage(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairMaxOI is a free data retrieval call binding the contract method 0xb2e1b2d6.
//
// Solidity: function pairMaxOI(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) PairMaxOI(opts *bind.CallOpts, _pairIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "pairMaxOI", _pairIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairMaxOI is a free data retrieval call binding the contract method 0xb2e1b2d6.
//
// Solidity: function pairMaxOI(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) PairMaxOI(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairMaxOI(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairMaxOI is a free data retrieval call binding the contract method 0xb2e1b2d6.
//
// Solidity: function pairMaxOI(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) PairMaxOI(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairMaxOI(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairMinLevPosUSDC is a free data retrieval call binding the contract method 0x238cd23f.
//
// Solidity: function pairMinLevPosUSDC(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) PairMinLevPosUSDC(opts *bind.CallOpts, _pairIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "pairMinLevPosUSDC", _pairIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairMinLevPosUSDC is a free data retrieval call binding the contract method 0x238cd23f.
//
// Solidity: function pairMinLevPosUSDC(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) PairMinLevPosUSDC(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairMinLevPosUSDC(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairMinLevPosUSDC is a free data retrieval call binding the contract method 0x238cd23f.
//
// Solidity: function pairMinLevPosUSDC(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) PairMinLevPosUSDC(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairMinLevPosUSDC(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairMinLeverage is a free data retrieval call binding the contract method 0x59a992d0.
//
// Solidity: function pairMinLeverage(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) PairMinLeverage(opts *bind.CallOpts, _pairIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "pairMinLeverage", _pairIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairMinLeverage is a free data retrieval call binding the contract method 0x59a992d0.
//
// Solidity: function pairMinLeverage(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) PairMinLeverage(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairMinLeverage(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairMinLeverage is a free data retrieval call binding the contract method 0x59a992d0.
//
// Solidity: function pairMinLeverage(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) PairMinLeverage(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairMinLeverage(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairOpenFeeP is a free data retrieval call binding the contract method 0xfb3e519d.
//
// Solidity: function pairOpenFeeP(uint256 _pairIndex, uint256 _leveragedPosition, bool _buy) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) PairOpenFeeP(opts *bind.CallOpts, _pairIndex *big.Int, _leveragedPosition *big.Int, _buy bool) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "pairOpenFeeP", _pairIndex, _leveragedPosition, _buy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairOpenFeeP is a free data retrieval call binding the contract method 0xfb3e519d.
//
// Solidity: function pairOpenFeeP(uint256 _pairIndex, uint256 _leveragedPosition, bool _buy) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) PairOpenFeeP(_pairIndex *big.Int, _leveragedPosition *big.Int, _buy bool) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairOpenFeeP(&_AvantisMulticallAbis.CallOpts, _pairIndex, _leveragedPosition, _buy)
}

// PairOpenFeeP is a free data retrieval call binding the contract method 0xfb3e519d.
//
// Solidity: function pairOpenFeeP(uint256 _pairIndex, uint256 _leveragedPosition, bool _buy) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) PairOpenFeeP(_pairIndex *big.Int, _leveragedPosition *big.Int, _buy bool) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairOpenFeeP(&_AvantisMulticallAbis.CallOpts, _pairIndex, _leveragedPosition, _buy)
}

// PairPriceImpactMultiplier is a free data retrieval call binding the contract method 0x5fc8cd28.
//
// Solidity: function pairPriceImpactMultiplier(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) PairPriceImpactMultiplier(opts *bind.CallOpts, _pairIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "pairPriceImpactMultiplier", _pairIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairPriceImpactMultiplier is a free data retrieval call binding the contract method 0x5fc8cd28.
//
// Solidity: function pairPriceImpactMultiplier(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) PairPriceImpactMultiplier(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairPriceImpactMultiplier(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairPriceImpactMultiplier is a free data retrieval call binding the contract method 0x5fc8cd28.
//
// Solidity: function pairPriceImpactMultiplier(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) PairPriceImpactMultiplier(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairPriceImpactMultiplier(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairSkewImpactMultiplier is a free data retrieval call binding the contract method 0xe8b50f84.
//
// Solidity: function pairSkewImpactMultiplier(uint256 _pairIndex) view returns(int256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) PairSkewImpactMultiplier(opts *bind.CallOpts, _pairIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "pairSkewImpactMultiplier", _pairIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairSkewImpactMultiplier is a free data retrieval call binding the contract method 0xe8b50f84.
//
// Solidity: function pairSkewImpactMultiplier(uint256 _pairIndex) view returns(int256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) PairSkewImpactMultiplier(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairSkewImpactMultiplier(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairSkewImpactMultiplier is a free data retrieval call binding the contract method 0xe8b50f84.
//
// Solidity: function pairSkewImpactMultiplier(uint256 _pairIndex) view returns(int256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) PairSkewImpactMultiplier(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairSkewImpactMultiplier(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairSpreadP is a free data retrieval call binding the contract method 0xa1d54e9b.
//
// Solidity: function pairSpreadP(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) PairSpreadP(opts *bind.CallOpts, _pairIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "pairSpreadP", _pairIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairSpreadP is a free data retrieval call binding the contract method 0xa1d54e9b.
//
// Solidity: function pairSpreadP(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) PairSpreadP(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairSpreadP(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// PairSpreadP is a free data retrieval call binding the contract method 0xa1d54e9b.
//
// Solidity: function pairSpreadP(uint256 _pairIndex) view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) PairSpreadP(_pairIndex *big.Int) (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairSpreadP(&_AvantisMulticallAbis.CallOpts, _pairIndex)
}

// Pairs is a free data retrieval call binding the contract method 0xb91ac788.
//
// Solidity: function pairs(uint256 ) view returns(string from, string to, (uint256,bytes32) feed, (uint256,address) backupFeed, uint256 spreadP, uint256 priceImpactMultiplier, int256 skewImpactMultiplier, uint256 groupIndex, uint256 feeIndex, uint256 groupOpenInterestPecentage, uint256 maxWalletOI)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) Pairs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	From                       string
	To                         string
	Feed                       IPairStorageFeed
	BackupFeed                 IPairStorageBackupFeed
	SpreadP                    *big.Int
	PriceImpactMultiplier      *big.Int
	SkewImpactMultiplier       *big.Int
	GroupIndex                 *big.Int
	FeeIndex                   *big.Int
	GroupOpenInterestPecentage *big.Int
	MaxWalletOI                *big.Int
}, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "pairs", arg0)

	outstruct := new(struct {
		From                       string
		To                         string
		Feed                       IPairStorageFeed
		BackupFeed                 IPairStorageBackupFeed
		SpreadP                    *big.Int
		PriceImpactMultiplier      *big.Int
		SkewImpactMultiplier       *big.Int
		GroupIndex                 *big.Int
		FeeIndex                   *big.Int
		GroupOpenInterestPecentage *big.Int
		MaxWalletOI                *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.From = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.To = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Feed = *abi.ConvertType(out[2], new(IPairStorageFeed)).(*IPairStorageFeed)
	outstruct.BackupFeed = *abi.ConvertType(out[3], new(IPairStorageBackupFeed)).(*IPairStorageBackupFeed)
	outstruct.SpreadP = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.PriceImpactMultiplier = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.SkewImpactMultiplier = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.GroupIndex = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.FeeIndex = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.GroupOpenInterestPecentage = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.MaxWalletOI = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Pairs is a free data retrieval call binding the contract method 0xb91ac788.
//
// Solidity: function pairs(uint256 ) view returns(string from, string to, (uint256,bytes32) feed, (uint256,address) backupFeed, uint256 spreadP, uint256 priceImpactMultiplier, int256 skewImpactMultiplier, uint256 groupIndex, uint256 feeIndex, uint256 groupOpenInterestPecentage, uint256 maxWalletOI)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) Pairs(arg0 *big.Int) (struct {
	From                       string
	To                         string
	Feed                       IPairStorageFeed
	BackupFeed                 IPairStorageBackupFeed
	SpreadP                    *big.Int
	PriceImpactMultiplier      *big.Int
	SkewImpactMultiplier       *big.Int
	GroupIndex                 *big.Int
	FeeIndex                   *big.Int
	GroupOpenInterestPecentage *big.Int
	MaxWalletOI                *big.Int
}, error) {
	return _AvantisMulticallAbis.Contract.Pairs(&_AvantisMulticallAbis.CallOpts, arg0)
}

// Pairs is a free data retrieval call binding the contract method 0xb91ac788.
//
// Solidity: function pairs(uint256 ) view returns(string from, string to, (uint256,bytes32) feed, (uint256,address) backupFeed, uint256 spreadP, uint256 priceImpactMultiplier, int256 skewImpactMultiplier, uint256 groupIndex, uint256 feeIndex, uint256 groupOpenInterestPecentage, uint256 maxWalletOI)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) Pairs(arg0 *big.Int) (struct {
	From                       string
	To                         string
	Feed                       IPairStorageFeed
	BackupFeed                 IPairStorageBackupFeed
	SpreadP                    *big.Int
	PriceImpactMultiplier      *big.Int
	SkewImpactMultiplier       *big.Int
	GroupIndex                 *big.Int
	FeeIndex                   *big.Int
	GroupOpenInterestPecentage *big.Int
	MaxWalletOI                *big.Int
}, error) {
	return _AvantisMulticallAbis.Contract.Pairs(&_AvantisMulticallAbis.CallOpts, arg0)
}

// PairsBackend is a free data retrieval call binding the contract method 0x9567dccf.
//
// Solidity: function pairsBackend(uint256 _index) view returns((string,string,(uint256,bytes32),(uint256,address),uint256,uint256,int256,uint256,uint256,uint256,uint256), (string,uint256,uint256,uint256), (string,uint256,uint256,uint256,uint256))
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) PairsBackend(opts *bind.CallOpts, _index *big.Int) (IPairStoragePair, IPairStorageGroup, IPairStorageFee, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "pairsBackend", _index)

	if err != nil {
		return *new(IPairStoragePair), *new(IPairStorageGroup), *new(IPairStorageFee), err
	}

	out0 := *abi.ConvertType(out[0], new(IPairStoragePair)).(*IPairStoragePair)
	out1 := *abi.ConvertType(out[1], new(IPairStorageGroup)).(*IPairStorageGroup)
	out2 := *abi.ConvertType(out[2], new(IPairStorageFee)).(*IPairStorageFee)

	return out0, out1, out2, err

}

// PairsBackend is a free data retrieval call binding the contract method 0x9567dccf.
//
// Solidity: function pairsBackend(uint256 _index) view returns((string,string,(uint256,bytes32),(uint256,address),uint256,uint256,int256,uint256,uint256,uint256,uint256), (string,uint256,uint256,uint256), (string,uint256,uint256,uint256,uint256))
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) PairsBackend(_index *big.Int) (IPairStoragePair, IPairStorageGroup, IPairStorageFee, error) {
	return _AvantisMulticallAbis.Contract.PairsBackend(&_AvantisMulticallAbis.CallOpts, _index)
}

// PairsBackend is a free data retrieval call binding the contract method 0x9567dccf.
//
// Solidity: function pairsBackend(uint256 _index) view returns((string,string,(uint256,bytes32),(uint256,address),uint256,uint256,int256,uint256,uint256,uint256,uint256), (string,uint256,uint256,uint256), (string,uint256,uint256,uint256,uint256))
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) PairsBackend(_index *big.Int) (IPairStoragePair, IPairStorageGroup, IPairStorageFee, error) {
	return _AvantisMulticallAbis.Contract.PairsBackend(&_AvantisMulticallAbis.CallOpts, _index)
}

// PairsCount is a free data retrieval call binding the contract method 0xb81b2b71.
//
// Solidity: function pairsCount() view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) PairsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "pairsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairsCount is a free data retrieval call binding the contract method 0xb81b2b71.
//
// Solidity: function pairsCount() view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) PairsCount() (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairsCount(&_AvantisMulticallAbis.CallOpts)
}

// PairsCount is a free data retrieval call binding the contract method 0xb81b2b71.
//
// Solidity: function pairsCount() view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) PairsCount() (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.PairsCount(&_AvantisMulticallAbis.CallOpts)
}

// SkewedFeesCount is a free data retrieval call binding the contract method 0x6d8b5224.
//
// Solidity: function skewedFeesCount() view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) SkewedFeesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "skewedFeesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SkewedFeesCount is a free data retrieval call binding the contract method 0x6d8b5224.
//
// Solidity: function skewedFeesCount() view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) SkewedFeesCount() (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.SkewedFeesCount(&_AvantisMulticallAbis.CallOpts)
}

// SkewedFeesCount is a free data retrieval call binding the contract method 0x6d8b5224.
//
// Solidity: function skewedFeesCount() view returns(uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) SkewedFeesCount() (*big.Int, error) {
	return _AvantisMulticallAbis.Contract.SkewedFeesCount(&_AvantisMulticallAbis.CallOpts)
}

// StorageT is a free data retrieval call binding the contract method 0x16fff074.
//
// Solidity: function storageT() view returns(address)
func (_AvantisMulticallAbis *AvantisMulticallAbisCaller) StorageT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AvantisMulticallAbis.contract.Call(opts, &out, "storageT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StorageT is a free data retrieval call binding the contract method 0x16fff074.
//
// Solidity: function storageT() view returns(address)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) StorageT() (common.Address, error) {
	return _AvantisMulticallAbis.Contract.StorageT(&_AvantisMulticallAbis.CallOpts)
}

// StorageT is a free data retrieval call binding the contract method 0x16fff074.
//
// Solidity: function storageT() view returns(address)
func (_AvantisMulticallAbis *AvantisMulticallAbisCallerSession) StorageT() (common.Address, error) {
	return _AvantisMulticallAbis.Contract.StorageT(&_AvantisMulticallAbis.CallOpts)
}

// AddFee is a paid mutator transaction binding the contract method 0x7f684baf.
//
// Solidity: function addFee((string,uint256,uint256,uint256,uint256) _fee) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactor) AddFee(opts *bind.TransactOpts, _fee IPairStorageFee) (*types.Transaction, error) {
	return _AvantisMulticallAbis.contract.Transact(opts, "addFee", _fee)
}

// AddFee is a paid mutator transaction binding the contract method 0x7f684baf.
//
// Solidity: function addFee((string,uint256,uint256,uint256,uint256) _fee) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) AddFee(_fee IPairStorageFee) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.AddFee(&_AvantisMulticallAbis.TransactOpts, _fee)
}

// AddFee is a paid mutator transaction binding the contract method 0x7f684baf.
//
// Solidity: function addFee((string,uint256,uint256,uint256,uint256) _fee) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorSession) AddFee(_fee IPairStorageFee) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.AddFee(&_AvantisMulticallAbis.TransactOpts, _fee)
}

// AddGroup is a paid mutator transaction binding the contract method 0x72fce6f0.
//
// Solidity: function addGroup((string,uint256,uint256,uint256) _group) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactor) AddGroup(opts *bind.TransactOpts, _group IPairStorageGroup) (*types.Transaction, error) {
	return _AvantisMulticallAbis.contract.Transact(opts, "addGroup", _group)
}

// AddGroup is a paid mutator transaction binding the contract method 0x72fce6f0.
//
// Solidity: function addGroup((string,uint256,uint256,uint256) _group) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) AddGroup(_group IPairStorageGroup) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.AddGroup(&_AvantisMulticallAbis.TransactOpts, _group)
}

// AddGroup is a paid mutator transaction binding the contract method 0x72fce6f0.
//
// Solidity: function addGroup((string,uint256,uint256,uint256) _group) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorSession) AddGroup(_group IPairStorageGroup) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.AddGroup(&_AvantisMulticallAbis.TransactOpts, _group)
}

// AddPair is a paid mutator transaction binding the contract method 0xe70ab5cd.
//
// Solidity: function addPair((string,string,(uint256,bytes32),(uint256,address),uint256,uint256,int256,uint256,uint256,uint256,uint256) _pair) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactor) AddPair(opts *bind.TransactOpts, _pair IPairStoragePair) (*types.Transaction, error) {
	return _AvantisMulticallAbis.contract.Transact(opts, "addPair", _pair)
}

// AddPair is a paid mutator transaction binding the contract method 0xe70ab5cd.
//
// Solidity: function addPair((string,string,(uint256,bytes32),(uint256,address),uint256,uint256,int256,uint256,uint256,uint256,uint256) _pair) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) AddPair(_pair IPairStoragePair) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.AddPair(&_AvantisMulticallAbis.TransactOpts, _pair)
}

// AddPair is a paid mutator transaction binding the contract method 0xe70ab5cd.
//
// Solidity: function addPair((string,string,(uint256,bytes32),(uint256,address),uint256,uint256,int256,uint256,uint256,uint256,uint256) _pair) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorSession) AddPair(_pair IPairStoragePair) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.AddPair(&_AvantisMulticallAbis.TransactOpts, _pair)
}

// AddSkewOpenFees is a paid mutator transaction binding the contract method 0x1a67bacc.
//
// Solidity: function addSkewOpenFees((int256[2][10]) _skewFee) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactor) AddSkewOpenFees(opts *bind.TransactOpts, _skewFee IPairStorageSkewFee) (*types.Transaction, error) {
	return _AvantisMulticallAbis.contract.Transact(opts, "addSkewOpenFees", _skewFee)
}

// AddSkewOpenFees is a paid mutator transaction binding the contract method 0x1a67bacc.
//
// Solidity: function addSkewOpenFees((int256[2][10]) _skewFee) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) AddSkewOpenFees(_skewFee IPairStorageSkewFee) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.AddSkewOpenFees(&_AvantisMulticallAbis.TransactOpts, _skewFee)
}

// AddSkewOpenFees is a paid mutator transaction binding the contract method 0x1a67bacc.
//
// Solidity: function addSkewOpenFees((int256[2][10]) _skewFee) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorSession) AddSkewOpenFees(_skewFee IPairStorageSkewFee) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.AddSkewOpenFees(&_AvantisMulticallAbis.TransactOpts, _skewFee)
}

// DelistPair is a paid mutator transaction binding the contract method 0x74db8c89.
//
// Solidity: function delistPair(uint256 _pairIndex) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactor) DelistPair(opts *bind.TransactOpts, _pairIndex *big.Int) (*types.Transaction, error) {
	return _AvantisMulticallAbis.contract.Transact(opts, "delistPair", _pairIndex)
}

// DelistPair is a paid mutator transaction binding the contract method 0x74db8c89.
//
// Solidity: function delistPair(uint256 _pairIndex) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) DelistPair(_pairIndex *big.Int) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.DelistPair(&_AvantisMulticallAbis.TransactOpts, _pairIndex)
}

// DelistPair is a paid mutator transaction binding the contract method 0x74db8c89.
//
// Solidity: function delistPair(uint256 _pairIndex) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorSession) DelistPair(_pairIndex *big.Int) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.DelistPair(&_AvantisMulticallAbis.TransactOpts, _pairIndex)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _storageT, uint256 _currentOrderId) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactor) Initialize(opts *bind.TransactOpts, _storageT common.Address, _currentOrderId *big.Int) (*types.Transaction, error) {
	return _AvantisMulticallAbis.contract.Transact(opts, "initialize", _storageT, _currentOrderId)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _storageT, uint256 _currentOrderId) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) Initialize(_storageT common.Address, _currentOrderId *big.Int) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.Initialize(&_AvantisMulticallAbis.TransactOpts, _storageT, _currentOrderId)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _storageT, uint256 _currentOrderId) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorSession) Initialize(_storageT common.Address, _currentOrderId *big.Int) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.Initialize(&_AvantisMulticallAbis.TransactOpts, _storageT, _currentOrderId)
}

// PairJob is a paid mutator transaction binding the contract method 0x302f81fc.
//
// Solidity: function pairJob(uint256 _pairIndex) returns(string, string, bytes32, address, uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactor) PairJob(opts *bind.TransactOpts, _pairIndex *big.Int) (*types.Transaction, error) {
	return _AvantisMulticallAbis.contract.Transact(opts, "pairJob", _pairIndex)
}

// PairJob is a paid mutator transaction binding the contract method 0x302f81fc.
//
// Solidity: function pairJob(uint256 _pairIndex) returns(string, string, bytes32, address, uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) PairJob(_pairIndex *big.Int) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.PairJob(&_AvantisMulticallAbis.TransactOpts, _pairIndex)
}

// PairJob is a paid mutator transaction binding the contract method 0x302f81fc.
//
// Solidity: function pairJob(uint256 _pairIndex) returns(string, string, bytes32, address, uint256)
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorSession) PairJob(_pairIndex *big.Int) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.PairJob(&_AvantisMulticallAbis.TransactOpts, _pairIndex)
}

// SetBlockOILImits is a paid mutator transaction binding the contract method 0x10ed4088.
//
// Solidity: function setBlockOILImits(uint256[] _pairIndex, uint256[] _limits) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactor) SetBlockOILImits(opts *bind.TransactOpts, _pairIndex []*big.Int, _limits []*big.Int) (*types.Transaction, error) {
	return _AvantisMulticallAbis.contract.Transact(opts, "setBlockOILImits", _pairIndex, _limits)
}

// SetBlockOILImits is a paid mutator transaction binding the contract method 0x10ed4088.
//
// Solidity: function setBlockOILImits(uint256[] _pairIndex, uint256[] _limits) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) SetBlockOILImits(_pairIndex []*big.Int, _limits []*big.Int) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.SetBlockOILImits(&_AvantisMulticallAbis.TransactOpts, _pairIndex, _limits)
}

// SetBlockOILImits is a paid mutator transaction binding the contract method 0x10ed4088.
//
// Solidity: function setBlockOILImits(uint256[] _pairIndex, uint256[] _limits) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorSession) SetBlockOILImits(_pairIndex []*big.Int, _limits []*big.Int) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.SetBlockOILImits(&_AvantisMulticallAbis.TransactOpts, _pairIndex, _limits)
}

// UdpateSkewOpenFees is a paid mutator transaction binding the contract method 0x3d7a5f1a.
//
// Solidity: function udpateSkewOpenFees(uint256 _pairIndex, (int256[2][10]) _skewFee) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactor) UdpateSkewOpenFees(opts *bind.TransactOpts, _pairIndex *big.Int, _skewFee IPairStorageSkewFee) (*types.Transaction, error) {
	return _AvantisMulticallAbis.contract.Transact(opts, "udpateSkewOpenFees", _pairIndex, _skewFee)
}

// UdpateSkewOpenFees is a paid mutator transaction binding the contract method 0x3d7a5f1a.
//
// Solidity: function udpateSkewOpenFees(uint256 _pairIndex, (int256[2][10]) _skewFee) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) UdpateSkewOpenFees(_pairIndex *big.Int, _skewFee IPairStorageSkewFee) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.UdpateSkewOpenFees(&_AvantisMulticallAbis.TransactOpts, _pairIndex, _skewFee)
}

// UdpateSkewOpenFees is a paid mutator transaction binding the contract method 0x3d7a5f1a.
//
// Solidity: function udpateSkewOpenFees(uint256 _pairIndex, (int256[2][10]) _skewFee) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorSession) UdpateSkewOpenFees(_pairIndex *big.Int, _skewFee IPairStorageSkewFee) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.UdpateSkewOpenFees(&_AvantisMulticallAbis.TransactOpts, _pairIndex, _skewFee)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9a7c200e.
//
// Solidity: function updateFee(uint256 _id, (string,uint256,uint256,uint256,uint256) _fee) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactor) UpdateFee(opts *bind.TransactOpts, _id *big.Int, _fee IPairStorageFee) (*types.Transaction, error) {
	return _AvantisMulticallAbis.contract.Transact(opts, "updateFee", _id, _fee)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9a7c200e.
//
// Solidity: function updateFee(uint256 _id, (string,uint256,uint256,uint256,uint256) _fee) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) UpdateFee(_id *big.Int, _fee IPairStorageFee) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.UpdateFee(&_AvantisMulticallAbis.TransactOpts, _id, _fee)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9a7c200e.
//
// Solidity: function updateFee(uint256 _id, (string,uint256,uint256,uint256,uint256) _fee) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorSession) UpdateFee(_id *big.Int, _fee IPairStorageFee) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.UpdateFee(&_AvantisMulticallAbis.TransactOpts, _id, _fee)
}

// UpdateGroup is a paid mutator transaction binding the contract method 0x5a3e32e0.
//
// Solidity: function updateGroup(uint256 _id, (string,uint256,uint256,uint256) _group) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactor) UpdateGroup(opts *bind.TransactOpts, _id *big.Int, _group IPairStorageGroup) (*types.Transaction, error) {
	return _AvantisMulticallAbis.contract.Transact(opts, "updateGroup", _id, _group)
}

// UpdateGroup is a paid mutator transaction binding the contract method 0x5a3e32e0.
//
// Solidity: function updateGroup(uint256 _id, (string,uint256,uint256,uint256) _group) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) UpdateGroup(_id *big.Int, _group IPairStorageGroup) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.UpdateGroup(&_AvantisMulticallAbis.TransactOpts, _id, _group)
}

// UpdateGroup is a paid mutator transaction binding the contract method 0x5a3e32e0.
//
// Solidity: function updateGroup(uint256 _id, (string,uint256,uint256,uint256) _group) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorSession) UpdateGroup(_id *big.Int, _group IPairStorageGroup) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.UpdateGroup(&_AvantisMulticallAbis.TransactOpts, _id, _group)
}

// UpdateGroupOI is a paid mutator transaction binding the contract method 0x4ffe8aec.
//
// Solidity: function updateGroupOI(uint256 _pairIndex, uint256 _amount, bool _long, bool _increase) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactor) UpdateGroupOI(opts *bind.TransactOpts, _pairIndex *big.Int, _amount *big.Int, _long bool, _increase bool) (*types.Transaction, error) {
	return _AvantisMulticallAbis.contract.Transact(opts, "updateGroupOI", _pairIndex, _amount, _long, _increase)
}

// UpdateGroupOI is a paid mutator transaction binding the contract method 0x4ffe8aec.
//
// Solidity: function updateGroupOI(uint256 _pairIndex, uint256 _amount, bool _long, bool _increase) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) UpdateGroupOI(_pairIndex *big.Int, _amount *big.Int, _long bool, _increase bool) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.UpdateGroupOI(&_AvantisMulticallAbis.TransactOpts, _pairIndex, _amount, _long, _increase)
}

// UpdateGroupOI is a paid mutator transaction binding the contract method 0x4ffe8aec.
//
// Solidity: function updateGroupOI(uint256 _pairIndex, uint256 _amount, bool _long, bool _increase) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorSession) UpdateGroupOI(_pairIndex *big.Int, _amount *big.Int, _long bool, _increase bool) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.UpdateGroupOI(&_AvantisMulticallAbis.TransactOpts, _pairIndex, _amount, _long, _increase)
}

// UpdateLossProtectionMultiplier is a paid mutator transaction binding the contract method 0x9586621f.
//
// Solidity: function updateLossProtectionMultiplier(uint256 _pairIndex, uint256[] _tier, uint256[] _multiplierPercent) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactor) UpdateLossProtectionMultiplier(opts *bind.TransactOpts, _pairIndex *big.Int, _tier []*big.Int, _multiplierPercent []*big.Int) (*types.Transaction, error) {
	return _AvantisMulticallAbis.contract.Transact(opts, "updateLossProtectionMultiplier", _pairIndex, _tier, _multiplierPercent)
}

// UpdateLossProtectionMultiplier is a paid mutator transaction binding the contract method 0x9586621f.
//
// Solidity: function updateLossProtectionMultiplier(uint256 _pairIndex, uint256[] _tier, uint256[] _multiplierPercent) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) UpdateLossProtectionMultiplier(_pairIndex *big.Int, _tier []*big.Int, _multiplierPercent []*big.Int) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.UpdateLossProtectionMultiplier(&_AvantisMulticallAbis.TransactOpts, _pairIndex, _tier, _multiplierPercent)
}

// UpdateLossProtectionMultiplier is a paid mutator transaction binding the contract method 0x9586621f.
//
// Solidity: function updateLossProtectionMultiplier(uint256 _pairIndex, uint256[] _tier, uint256[] _multiplierPercent) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorSession) UpdateLossProtectionMultiplier(_pairIndex *big.Int, _tier []*big.Int, _multiplierPercent []*big.Int) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.UpdateLossProtectionMultiplier(&_AvantisMulticallAbis.TransactOpts, _pairIndex, _tier, _multiplierPercent)
}

// UpdatePair is a paid mutator transaction binding the contract method 0xf645aff5.
//
// Solidity: function updatePair(uint256 _pairIndex, (string,string,(uint256,bytes32),(uint256,address),uint256,uint256,int256,uint256,uint256,uint256,uint256) _pair) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactor) UpdatePair(opts *bind.TransactOpts, _pairIndex *big.Int, _pair IPairStoragePair) (*types.Transaction, error) {
	return _AvantisMulticallAbis.contract.Transact(opts, "updatePair", _pairIndex, _pair)
}

// UpdatePair is a paid mutator transaction binding the contract method 0xf645aff5.
//
// Solidity: function updatePair(uint256 _pairIndex, (string,string,(uint256,bytes32),(uint256,address),uint256,uint256,int256,uint256,uint256,uint256,uint256) _pair) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisSession) UpdatePair(_pairIndex *big.Int, _pair IPairStoragePair) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.UpdatePair(&_AvantisMulticallAbis.TransactOpts, _pairIndex, _pair)
}

// UpdatePair is a paid mutator transaction binding the contract method 0xf645aff5.
//
// Solidity: function updatePair(uint256 _pairIndex, (string,string,(uint256,bytes32),(uint256,address),uint256,uint256,int256,uint256,uint256,uint256,uint256) _pair) returns()
func (_AvantisMulticallAbis *AvantisMulticallAbisTransactorSession) UpdatePair(_pairIndex *big.Int, _pair IPairStoragePair) (*types.Transaction, error) {
	return _AvantisMulticallAbis.Contract.UpdatePair(&_AvantisMulticallAbis.TransactOpts, _pairIndex, _pair)
}

// AvantisMulticallAbisBlockOILimitsSetIterator is returned from FilterBlockOILimitsSet and is used to iterate over the raw logs and unpacked data for BlockOILimitsSet events raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisBlockOILimitsSetIterator struct {
	Event *AvantisMulticallAbisBlockOILimitsSet // Event containing the contract specifics and raw log

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
func (it *AvantisMulticallAbisBlockOILimitsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisMulticallAbisBlockOILimitsSet)
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
		it.Event = new(AvantisMulticallAbisBlockOILimitsSet)
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
func (it *AvantisMulticallAbisBlockOILimitsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisMulticallAbisBlockOILimitsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisMulticallAbisBlockOILimitsSet represents a BlockOILimitsSet event raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisBlockOILimitsSet struct {
	PairIndex []*big.Int
	Limits    []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlockOILimitsSet is a free log retrieval operation binding the contract event 0xc5c5af5d05689513c60f635389c8449ac23f2b277af3a15f88c21c6de36fcf14.
//
// Solidity: event BlockOILimitsSet(uint256[] pairIndex, uint256[] limits)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) FilterBlockOILimitsSet(opts *bind.FilterOpts) (*AvantisMulticallAbisBlockOILimitsSetIterator, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.FilterLogs(opts, "BlockOILimitsSet")
	if err != nil {
		return nil, err
	}
	return &AvantisMulticallAbisBlockOILimitsSetIterator{contract: _AvantisMulticallAbis.contract, event: "BlockOILimitsSet", logs: logs, sub: sub}, nil
}

// WatchBlockOILimitsSet is a free log subscription operation binding the contract event 0xc5c5af5d05689513c60f635389c8449ac23f2b277af3a15f88c21c6de36fcf14.
//
// Solidity: event BlockOILimitsSet(uint256[] pairIndex, uint256[] limits)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) WatchBlockOILimitsSet(opts *bind.WatchOpts, sink chan<- *AvantisMulticallAbisBlockOILimitsSet) (event.Subscription, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.WatchLogs(opts, "BlockOILimitsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisMulticallAbisBlockOILimitsSet)
				if err := _AvantisMulticallAbis.contract.UnpackLog(event, "BlockOILimitsSet", log); err != nil {
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
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) ParseBlockOILimitsSet(log types.Log) (*AvantisMulticallAbisBlockOILimitsSet, error) {
	event := new(AvantisMulticallAbisBlockOILimitsSet)
	if err := _AvantisMulticallAbis.contract.UnpackLog(event, "BlockOILimitsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvantisMulticallAbisFeeAddedIterator is returned from FilterFeeAdded and is used to iterate over the raw logs and unpacked data for FeeAdded events raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisFeeAddedIterator struct {
	Event *AvantisMulticallAbisFeeAdded // Event containing the contract specifics and raw log

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
func (it *AvantisMulticallAbisFeeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisMulticallAbisFeeAdded)
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
		it.Event = new(AvantisMulticallAbisFeeAdded)
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
func (it *AvantisMulticallAbisFeeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisMulticallAbisFeeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisMulticallAbisFeeAdded represents a FeeAdded event raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisFeeAdded struct {
	Index *big.Int
	Name  string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFeeAdded is a free log retrieval operation binding the contract event 0x482049823c85e038e099fe4f2b901487c4800def71c9a3f5bae2de8381ec54f6.
//
// Solidity: event FeeAdded(uint256 index, string name)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) FilterFeeAdded(opts *bind.FilterOpts) (*AvantisMulticallAbisFeeAddedIterator, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.FilterLogs(opts, "FeeAdded")
	if err != nil {
		return nil, err
	}
	return &AvantisMulticallAbisFeeAddedIterator{contract: _AvantisMulticallAbis.contract, event: "FeeAdded", logs: logs, sub: sub}, nil
}

// WatchFeeAdded is a free log subscription operation binding the contract event 0x482049823c85e038e099fe4f2b901487c4800def71c9a3f5bae2de8381ec54f6.
//
// Solidity: event FeeAdded(uint256 index, string name)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) WatchFeeAdded(opts *bind.WatchOpts, sink chan<- *AvantisMulticallAbisFeeAdded) (event.Subscription, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.WatchLogs(opts, "FeeAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisMulticallAbisFeeAdded)
				if err := _AvantisMulticallAbis.contract.UnpackLog(event, "FeeAdded", log); err != nil {
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
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) ParseFeeAdded(log types.Log) (*AvantisMulticallAbisFeeAdded, error) {
	event := new(AvantisMulticallAbisFeeAdded)
	if err := _AvantisMulticallAbis.contract.UnpackLog(event, "FeeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvantisMulticallAbisFeeUpdatedIterator is returned from FilterFeeUpdated and is used to iterate over the raw logs and unpacked data for FeeUpdated events raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisFeeUpdatedIterator struct {
	Event *AvantisMulticallAbisFeeUpdated // Event containing the contract specifics and raw log

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
func (it *AvantisMulticallAbisFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisMulticallAbisFeeUpdated)
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
		it.Event = new(AvantisMulticallAbisFeeUpdated)
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
func (it *AvantisMulticallAbisFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisMulticallAbisFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisMulticallAbisFeeUpdated represents a FeeUpdated event raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisFeeUpdated struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFeeUpdated is a free log retrieval operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 index)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) FilterFeeUpdated(opts *bind.FilterOpts) (*AvantisMulticallAbisFeeUpdatedIterator, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.FilterLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return &AvantisMulticallAbisFeeUpdatedIterator{contract: _AvantisMulticallAbis.contract, event: "FeeUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeUpdated is a free log subscription operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 index)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) WatchFeeUpdated(opts *bind.WatchOpts, sink chan<- *AvantisMulticallAbisFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.WatchLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisMulticallAbisFeeUpdated)
				if err := _AvantisMulticallAbis.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
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
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) ParseFeeUpdated(log types.Log) (*AvantisMulticallAbisFeeUpdated, error) {
	event := new(AvantisMulticallAbisFeeUpdated)
	if err := _AvantisMulticallAbis.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvantisMulticallAbisGroupAddedIterator is returned from FilterGroupAdded and is used to iterate over the raw logs and unpacked data for GroupAdded events raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisGroupAddedIterator struct {
	Event *AvantisMulticallAbisGroupAdded // Event containing the contract specifics and raw log

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
func (it *AvantisMulticallAbisGroupAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisMulticallAbisGroupAdded)
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
		it.Event = new(AvantisMulticallAbisGroupAdded)
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
func (it *AvantisMulticallAbisGroupAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisMulticallAbisGroupAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisMulticallAbisGroupAdded represents a GroupAdded event raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisGroupAdded struct {
	Index *big.Int
	Name  string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGroupAdded is a free log retrieval operation binding the contract event 0xaf17de8e82beccc440012117a600dc37e26925225d0f1ee192fc107eb3dcbca4.
//
// Solidity: event GroupAdded(uint256 index, string name)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) FilterGroupAdded(opts *bind.FilterOpts) (*AvantisMulticallAbisGroupAddedIterator, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.FilterLogs(opts, "GroupAdded")
	if err != nil {
		return nil, err
	}
	return &AvantisMulticallAbisGroupAddedIterator{contract: _AvantisMulticallAbis.contract, event: "GroupAdded", logs: logs, sub: sub}, nil
}

// WatchGroupAdded is a free log subscription operation binding the contract event 0xaf17de8e82beccc440012117a600dc37e26925225d0f1ee192fc107eb3dcbca4.
//
// Solidity: event GroupAdded(uint256 index, string name)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) WatchGroupAdded(opts *bind.WatchOpts, sink chan<- *AvantisMulticallAbisGroupAdded) (event.Subscription, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.WatchLogs(opts, "GroupAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisMulticallAbisGroupAdded)
				if err := _AvantisMulticallAbis.contract.UnpackLog(event, "GroupAdded", log); err != nil {
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
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) ParseGroupAdded(log types.Log) (*AvantisMulticallAbisGroupAdded, error) {
	event := new(AvantisMulticallAbisGroupAdded)
	if err := _AvantisMulticallAbis.contract.UnpackLog(event, "GroupAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvantisMulticallAbisGroupUpdatedIterator is returned from FilterGroupUpdated and is used to iterate over the raw logs and unpacked data for GroupUpdated events raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisGroupUpdatedIterator struct {
	Event *AvantisMulticallAbisGroupUpdated // Event containing the contract specifics and raw log

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
func (it *AvantisMulticallAbisGroupUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisMulticallAbisGroupUpdated)
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
		it.Event = new(AvantisMulticallAbisGroupUpdated)
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
func (it *AvantisMulticallAbisGroupUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisMulticallAbisGroupUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisMulticallAbisGroupUpdated represents a GroupUpdated event raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisGroupUpdated struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGroupUpdated is a free log retrieval operation binding the contract event 0xcfde8f228364c70f12cbbac5a88fc91ceca76dd750ac93364991a333b34afb8e.
//
// Solidity: event GroupUpdated(uint256 index)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) FilterGroupUpdated(opts *bind.FilterOpts) (*AvantisMulticallAbisGroupUpdatedIterator, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.FilterLogs(opts, "GroupUpdated")
	if err != nil {
		return nil, err
	}
	return &AvantisMulticallAbisGroupUpdatedIterator{contract: _AvantisMulticallAbis.contract, event: "GroupUpdated", logs: logs, sub: sub}, nil
}

// WatchGroupUpdated is a free log subscription operation binding the contract event 0xcfde8f228364c70f12cbbac5a88fc91ceca76dd750ac93364991a333b34afb8e.
//
// Solidity: event GroupUpdated(uint256 index)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) WatchGroupUpdated(opts *bind.WatchOpts, sink chan<- *AvantisMulticallAbisGroupUpdated) (event.Subscription, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.WatchLogs(opts, "GroupUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisMulticallAbisGroupUpdated)
				if err := _AvantisMulticallAbis.contract.UnpackLog(event, "GroupUpdated", log); err != nil {
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
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) ParseGroupUpdated(log types.Log) (*AvantisMulticallAbisGroupUpdated, error) {
	event := new(AvantisMulticallAbisGroupUpdated)
	if err := _AvantisMulticallAbis.contract.UnpackLog(event, "GroupUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvantisMulticallAbisInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisInitializedIterator struct {
	Event *AvantisMulticallAbisInitialized // Event containing the contract specifics and raw log

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
func (it *AvantisMulticallAbisInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisMulticallAbisInitialized)
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
		it.Event = new(AvantisMulticallAbisInitialized)
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
func (it *AvantisMulticallAbisInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisMulticallAbisInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisMulticallAbisInitialized represents a Initialized event raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) FilterInitialized(opts *bind.FilterOpts) (*AvantisMulticallAbisInitializedIterator, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AvantisMulticallAbisInitializedIterator{contract: _AvantisMulticallAbis.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AvantisMulticallAbisInitialized) (event.Subscription, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisMulticallAbisInitialized)
				if err := _AvantisMulticallAbis.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) ParseInitialized(log types.Log) (*AvantisMulticallAbisInitialized, error) {
	event := new(AvantisMulticallAbisInitialized)
	if err := _AvantisMulticallAbis.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvantisMulticallAbisLossProtectionAddedIterator is returned from FilterLossProtectionAdded and is used to iterate over the raw logs and unpacked data for LossProtectionAdded events raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisLossProtectionAddedIterator struct {
	Event *AvantisMulticallAbisLossProtectionAdded // Event containing the contract specifics and raw log

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
func (it *AvantisMulticallAbisLossProtectionAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisMulticallAbisLossProtectionAdded)
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
		it.Event = new(AvantisMulticallAbisLossProtectionAdded)
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
func (it *AvantisMulticallAbisLossProtectionAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisMulticallAbisLossProtectionAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisMulticallAbisLossProtectionAdded represents a LossProtectionAdded event raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisLossProtectionAdded struct {
	PairIndex  *big.Int
	Tier       []*big.Int
	Multiplier []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLossProtectionAdded is a free log retrieval operation binding the contract event 0xf66559a3ee7945e0e49289392f86db1a3a3049abc2a1c89664a702fa04273227.
//
// Solidity: event LossProtectionAdded(uint256 pairIndex, uint256[] tier, uint256[] multiplier)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) FilterLossProtectionAdded(opts *bind.FilterOpts) (*AvantisMulticallAbisLossProtectionAddedIterator, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.FilterLogs(opts, "LossProtectionAdded")
	if err != nil {
		return nil, err
	}
	return &AvantisMulticallAbisLossProtectionAddedIterator{contract: _AvantisMulticallAbis.contract, event: "LossProtectionAdded", logs: logs, sub: sub}, nil
}

// WatchLossProtectionAdded is a free log subscription operation binding the contract event 0xf66559a3ee7945e0e49289392f86db1a3a3049abc2a1c89664a702fa04273227.
//
// Solidity: event LossProtectionAdded(uint256 pairIndex, uint256[] tier, uint256[] multiplier)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) WatchLossProtectionAdded(opts *bind.WatchOpts, sink chan<- *AvantisMulticallAbisLossProtectionAdded) (event.Subscription, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.WatchLogs(opts, "LossProtectionAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisMulticallAbisLossProtectionAdded)
				if err := _AvantisMulticallAbis.contract.UnpackLog(event, "LossProtectionAdded", log); err != nil {
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
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) ParseLossProtectionAdded(log types.Log) (*AvantisMulticallAbisLossProtectionAdded, error) {
	event := new(AvantisMulticallAbisLossProtectionAdded)
	if err := _AvantisMulticallAbis.contract.UnpackLog(event, "LossProtectionAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvantisMulticallAbisOrderLimitsSetIterator is returned from FilterOrderLimitsSet and is used to iterate over the raw logs and unpacked data for OrderLimitsSet events raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisOrderLimitsSetIterator struct {
	Event *AvantisMulticallAbisOrderLimitsSet // Event containing the contract specifics and raw log

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
func (it *AvantisMulticallAbisOrderLimitsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisMulticallAbisOrderLimitsSet)
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
		it.Event = new(AvantisMulticallAbisOrderLimitsSet)
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
func (it *AvantisMulticallAbisOrderLimitsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisMulticallAbisOrderLimitsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisMulticallAbisOrderLimitsSet represents a OrderLimitsSet event raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisOrderLimitsSet struct {
	PairIndex []*big.Int
	Limits    []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderLimitsSet is a free log retrieval operation binding the contract event 0x5ffcf71171c83ae021779d84f12b918f2bb07077c6d3047c5d3203022317f6a3.
//
// Solidity: event OrderLimitsSet(uint256[] pairIndex, uint256[] limits)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) FilterOrderLimitsSet(opts *bind.FilterOpts) (*AvantisMulticallAbisOrderLimitsSetIterator, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.FilterLogs(opts, "OrderLimitsSet")
	if err != nil {
		return nil, err
	}
	return &AvantisMulticallAbisOrderLimitsSetIterator{contract: _AvantisMulticallAbis.contract, event: "OrderLimitsSet", logs: logs, sub: sub}, nil
}

// WatchOrderLimitsSet is a free log subscription operation binding the contract event 0x5ffcf71171c83ae021779d84f12b918f2bb07077c6d3047c5d3203022317f6a3.
//
// Solidity: event OrderLimitsSet(uint256[] pairIndex, uint256[] limits)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) WatchOrderLimitsSet(opts *bind.WatchOpts, sink chan<- *AvantisMulticallAbisOrderLimitsSet) (event.Subscription, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.WatchLogs(opts, "OrderLimitsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisMulticallAbisOrderLimitsSet)
				if err := _AvantisMulticallAbis.contract.UnpackLog(event, "OrderLimitsSet", log); err != nil {
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
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) ParseOrderLimitsSet(log types.Log) (*AvantisMulticallAbisOrderLimitsSet, error) {
	event := new(AvantisMulticallAbisOrderLimitsSet)
	if err := _AvantisMulticallAbis.contract.UnpackLog(event, "OrderLimitsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvantisMulticallAbisPairAddedIterator is returned from FilterPairAdded and is used to iterate over the raw logs and unpacked data for PairAdded events raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisPairAddedIterator struct {
	Event *AvantisMulticallAbisPairAdded // Event containing the contract specifics and raw log

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
func (it *AvantisMulticallAbisPairAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisMulticallAbisPairAdded)
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
		it.Event = new(AvantisMulticallAbisPairAdded)
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
func (it *AvantisMulticallAbisPairAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisMulticallAbisPairAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisMulticallAbisPairAdded represents a PairAdded event raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisPairAdded struct {
	Index *big.Int
	From  string
	To    string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPairAdded is a free log retrieval operation binding the contract event 0x3adfd40f2b74073df2a84238acdb7f460565a557b3cc13bddc8833289bf38e09.
//
// Solidity: event PairAdded(uint256 index, string from, string to)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) FilterPairAdded(opts *bind.FilterOpts) (*AvantisMulticallAbisPairAddedIterator, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.FilterLogs(opts, "PairAdded")
	if err != nil {
		return nil, err
	}
	return &AvantisMulticallAbisPairAddedIterator{contract: _AvantisMulticallAbis.contract, event: "PairAdded", logs: logs, sub: sub}, nil
}

// WatchPairAdded is a free log subscription operation binding the contract event 0x3adfd40f2b74073df2a84238acdb7f460565a557b3cc13bddc8833289bf38e09.
//
// Solidity: event PairAdded(uint256 index, string from, string to)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) WatchPairAdded(opts *bind.WatchOpts, sink chan<- *AvantisMulticallAbisPairAdded) (event.Subscription, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.WatchLogs(opts, "PairAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisMulticallAbisPairAdded)
				if err := _AvantisMulticallAbis.contract.UnpackLog(event, "PairAdded", log); err != nil {
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
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) ParsePairAdded(log types.Log) (*AvantisMulticallAbisPairAdded, error) {
	event := new(AvantisMulticallAbisPairAdded)
	if err := _AvantisMulticallAbis.contract.UnpackLog(event, "PairAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvantisMulticallAbisPairUpdatedIterator is returned from FilterPairUpdated and is used to iterate over the raw logs and unpacked data for PairUpdated events raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisPairUpdatedIterator struct {
	Event *AvantisMulticallAbisPairUpdated // Event containing the contract specifics and raw log

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
func (it *AvantisMulticallAbisPairUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisMulticallAbisPairUpdated)
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
		it.Event = new(AvantisMulticallAbisPairUpdated)
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
func (it *AvantisMulticallAbisPairUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisMulticallAbisPairUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisMulticallAbisPairUpdated represents a PairUpdated event raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisPairUpdated struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPairUpdated is a free log retrieval operation binding the contract event 0x123a1b961ae93e7acda9790b318237b175b45ac09277cd3614305d8baa3f1953.
//
// Solidity: event PairUpdated(uint256 index)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) FilterPairUpdated(opts *bind.FilterOpts) (*AvantisMulticallAbisPairUpdatedIterator, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.FilterLogs(opts, "PairUpdated")
	if err != nil {
		return nil, err
	}
	return &AvantisMulticallAbisPairUpdatedIterator{contract: _AvantisMulticallAbis.contract, event: "PairUpdated", logs: logs, sub: sub}, nil
}

// WatchPairUpdated is a free log subscription operation binding the contract event 0x123a1b961ae93e7acda9790b318237b175b45ac09277cd3614305d8baa3f1953.
//
// Solidity: event PairUpdated(uint256 index)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) WatchPairUpdated(opts *bind.WatchOpts, sink chan<- *AvantisMulticallAbisPairUpdated) (event.Subscription, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.WatchLogs(opts, "PairUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisMulticallAbisPairUpdated)
				if err := _AvantisMulticallAbis.contract.UnpackLog(event, "PairUpdated", log); err != nil {
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
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) ParsePairUpdated(log types.Log) (*AvantisMulticallAbisPairUpdated, error) {
	event := new(AvantisMulticallAbisPairUpdated)
	if err := _AvantisMulticallAbis.contract.UnpackLog(event, "PairUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvantisMulticallAbisSkewFeeAddedIterator is returned from FilterSkewFeeAdded and is used to iterate over the raw logs and unpacked data for SkewFeeAdded events raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisSkewFeeAddedIterator struct {
	Event *AvantisMulticallAbisSkewFeeAdded // Event containing the contract specifics and raw log

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
func (it *AvantisMulticallAbisSkewFeeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisMulticallAbisSkewFeeAdded)
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
		it.Event = new(AvantisMulticallAbisSkewFeeAdded)
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
func (it *AvantisMulticallAbisSkewFeeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisMulticallAbisSkewFeeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisMulticallAbisSkewFeeAdded represents a SkewFeeAdded event raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisSkewFeeAdded struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSkewFeeAdded is a free log retrieval operation binding the contract event 0x5d7a16e490fc66ca6522d4ba0437ac1a3b97cf25666340a19e408176295826d7.
//
// Solidity: event SkewFeeAdded(uint256 index)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) FilterSkewFeeAdded(opts *bind.FilterOpts) (*AvantisMulticallAbisSkewFeeAddedIterator, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.FilterLogs(opts, "SkewFeeAdded")
	if err != nil {
		return nil, err
	}
	return &AvantisMulticallAbisSkewFeeAddedIterator{contract: _AvantisMulticallAbis.contract, event: "SkewFeeAdded", logs: logs, sub: sub}, nil
}

// WatchSkewFeeAdded is a free log subscription operation binding the contract event 0x5d7a16e490fc66ca6522d4ba0437ac1a3b97cf25666340a19e408176295826d7.
//
// Solidity: event SkewFeeAdded(uint256 index)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) WatchSkewFeeAdded(opts *bind.WatchOpts, sink chan<- *AvantisMulticallAbisSkewFeeAdded) (event.Subscription, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.WatchLogs(opts, "SkewFeeAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisMulticallAbisSkewFeeAdded)
				if err := _AvantisMulticallAbis.contract.UnpackLog(event, "SkewFeeAdded", log); err != nil {
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
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) ParseSkewFeeAdded(log types.Log) (*AvantisMulticallAbisSkewFeeAdded, error) {
	event := new(AvantisMulticallAbisSkewFeeAdded)
	if err := _AvantisMulticallAbis.contract.UnpackLog(event, "SkewFeeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvantisMulticallAbisSkewFeeUpdatedIterator is returned from FilterSkewFeeUpdated and is used to iterate over the raw logs and unpacked data for SkewFeeUpdated events raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisSkewFeeUpdatedIterator struct {
	Event *AvantisMulticallAbisSkewFeeUpdated // Event containing the contract specifics and raw log

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
func (it *AvantisMulticallAbisSkewFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvantisMulticallAbisSkewFeeUpdated)
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
		it.Event = new(AvantisMulticallAbisSkewFeeUpdated)
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
func (it *AvantisMulticallAbisSkewFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvantisMulticallAbisSkewFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvantisMulticallAbisSkewFeeUpdated represents a SkewFeeUpdated event raised by the AvantisMulticallAbis contract.
type AvantisMulticallAbisSkewFeeUpdated struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSkewFeeUpdated is a free log retrieval operation binding the contract event 0x049d8e73ae29fd805de1f8eae372a5020742554b37118c1e33687e89ef3027ff.
//
// Solidity: event SkewFeeUpdated(uint256 index)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) FilterSkewFeeUpdated(opts *bind.FilterOpts) (*AvantisMulticallAbisSkewFeeUpdatedIterator, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.FilterLogs(opts, "SkewFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &AvantisMulticallAbisSkewFeeUpdatedIterator{contract: _AvantisMulticallAbis.contract, event: "SkewFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchSkewFeeUpdated is a free log subscription operation binding the contract event 0x049d8e73ae29fd805de1f8eae372a5020742554b37118c1e33687e89ef3027ff.
//
// Solidity: event SkewFeeUpdated(uint256 index)
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) WatchSkewFeeUpdated(opts *bind.WatchOpts, sink chan<- *AvantisMulticallAbisSkewFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _AvantisMulticallAbis.contract.WatchLogs(opts, "SkewFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvantisMulticallAbisSkewFeeUpdated)
				if err := _AvantisMulticallAbis.contract.UnpackLog(event, "SkewFeeUpdated", log); err != nil {
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
func (_AvantisMulticallAbis *AvantisMulticallAbisFilterer) ParseSkewFeeUpdated(log types.Log) (*AvantisMulticallAbisSkewFeeUpdated, error) {
	event := new(AvantisMulticallAbisSkewFeeUpdated)
	if err := _AvantisMulticallAbis.contract.UnpackLog(event, "SkewFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
