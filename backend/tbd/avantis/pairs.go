package avantis

import (
	"context"
	"encoding/json"
	"io"
	"math/big"
	"os"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Feed struct {
	MaxDeviationP *big.Int
	FeedId        [32]byte
}

type BackupFeed struct {
	MaxDeviationP *big.Int
	FeedId        common.Address
}

type PairInfo struct {
	From                       string
	To                         string
	Feed                       Feed
	BackupFeed                 BackupFeed
	SpreadP                    *big.Int
	PriceImpactMultiplier      *big.Int
	SkewImpactMultiplier       *big.Int
	GroupIndex                 *big.Int
	FeeIndex                   *big.Int
	GroupOpenInterestPecentage *big.Int
	MaxWalletOI                *big.Int
}

type PairsCache struct {
	Client          *ethclient.Client
	PairStorageAbi  abi.ABI
	MulticallAbi    abi.ABI
	MulticallAddr   common.Address
	PairStorageAddr common.Address
	Cache           []PairInfo
	Mu              sync.RWMutex
}

func NewPairsCache(ethClientURL, multicallABIPath, pairStorageABIPath, multicallAddr, pairStorageAddr string) (*PairsCache, error) {
	client, err := ethclient.Dial(ethClientURL)
	if err != nil {
		return nil, err
	}

	multicallAbi, err := readABIFromFile(multicallABIPath)
	if err != nil {
		return nil, err
	}

	pairStorageAbi, err := readABIFromFile(pairStorageABIPath)
	if err != nil {
		return nil, err
	}

	cache := &PairsCache{
		Client:          client,
		PairStorageAbi:  pairStorageAbi,
		MulticallAbi:    multicallAbi,
		MulticallAddr:   common.HexToAddress(multicallAddr),
		PairStorageAddr: common.HexToAddress(pairStorageAddr),
		Cache:           make([]PairInfo, 0),
	}

	err = cache.LoadPairs()
	if err != nil {
		return nil, err
	}
	return cache, nil
}

func readABIFromFile(filePath string) (abi.ABI, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return abi.ABI{}, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return abi.ABI{}, err
	}

	var contractABI abi.ABI
	err = json.Unmarshal(byteValue, &contractABI)
	if err != nil {
		return abi.ABI{}, err
	}

	return contractABI, nil
}

func (pc *PairsCache) LoadPairs() error {
	ctx := context.Background()
	pairCountData, err := pc.PairStorageAbi.Pack("pairsCount")
	if err != nil {
		return err
	}

	callMsg := ethereum.CallMsg{To: &pc.PairStorageAddr, Data: pairCountData}
	result, err := pc.Client.CallContract(ctx, callMsg, nil)
	if err != nil {
		return err
	}

	var pairsCount *big.Int
	err = pc.PairStorageAbi.UnpackIntoInterface(&pairsCount, "pairsCount", result)
	if err != nil {
		return err
	}

	var calls []struct {
		Target   common.Address
		CallData []byte
	}

	for i := 0; i < int(pairsCount.Int64()); i++ {
		callData, err := pc.PairStorageAbi.Pack("pairs", big.NewInt(int64(i)))
		if err != nil {
			return err
		}
		calls = append(calls, struct {
			Target   common.Address
			CallData []byte
		}{
			Target:   pc.PairStorageAddr,
			CallData: callData,
		})
	}

	multicallData, err := pc.MulticallAbi.Pack("aggregate", calls)
	if err != nil {
		return err
	}

	callMsg = ethereum.CallMsg{To: &pc.MulticallAddr, Data: multicallData}
	result, err = pc.Client.CallContract(ctx, callMsg, nil)
	if err != nil {
		return err
	}

	var returnData struct {
		BlockNumber *big.Int
		ReturnData  [][]byte
	}

	err = pc.MulticallAbi.UnpackIntoInterface(&returnData, "aggregate", result)
	if err != nil {
		return err
	}

	pc.Mu.Lock()
	defer pc.Mu.Unlock()

	for _, data := range returnData.ReturnData {
		var pairInfo PairInfo
		err := pc.PairStorageAbi.UnpackIntoInterface(&pairInfo, "pairs", data)
		if err != nil {
			return err
		}

		pc.Cache = append(pc.Cache, pairInfo)
	}

	return nil
}

func (pc *PairsCache) ShowPairs() []PairInfo {
	pc.Mu.RLock()
	defer pc.Mu.RUnlock()
	copyCache := make([]PairInfo, len(pc.Cache))
	copy(copyCache, pc.Cache)
	return copyCache
}

func (pc *PairsCache) GetPair(index int) (PairInfo, bool) {
	pc.Mu.RLock()
	defer pc.Mu.RUnlock()
	if index < 0 || index >= len(pc.Cache) {
		return PairInfo{}, false
	}
	return pc.Cache[index], true
}
