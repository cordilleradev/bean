package avantis

import (
	avantis_multicall_abis "github.com/cordilleradev/bean/common/abigen/avantis/multiCall"
	"github.com/cordilleradev/bean/common/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type avantisConnectionPool struct {
	pool *utils.EvmRpcPool[*avantis_multicall_abis.AvantisMulticallAbisCaller]
}

func newAvantisConnectionPool(rpcUrls []string, avantisMulticallContractAddress string) (*avantisConnectionPool, error) {
	contractAddress := common.HexToAddress(avantisMulticallContractAddress)

	clientCtor := func(
		address common.Address,
		client *ethclient.Client,
	) (*avantis_multicall_abis.AvantisMulticallAbisCaller, error) {
		return avantis_multicall_abis.NewAvantisMulticallAbisCaller(address, client)
	}
	pool, err := utils.NewCommonPoolHelper(rpcUrls, contractAddress, clientCtor)
	if err != nil {
		return nil, err
	}
	return &avantisConnectionPool{
		pool: pool,
	}, nil
}

func (a *avantisConnectionPool) getPositions(userAddress common.Address) []avantis_multicall_abis.IMulticallAggregatedTrade {
	// requestCount := 0
	for {
		// requestCount++
		caller := a.pool.NextCaller()
		if caller == nil {
			continue
		}
		positions, _, err := caller.GetPositions(nil, userAddress)
		if err == nil {
			// fmt.Printf(" Number of requests: %d\n", requestCount)
			return positions
		}
	}
}
