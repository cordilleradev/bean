package gmx

import (
	"context"
	"errors"
	"math/big"
	"sync"

	gmx_abis "github.com/cordilleradev/bean/common/abigen/gmx"
	"github.com/cordilleradev/bean/common/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type gmxConnectionPool struct {
	connections []*gmx_abis.GmxAbisCaller
	mu          sync.Mutex
	index       int
}

func newGmxConnectionPool(rpcUrls []string, gmxReaderContractAddress string) (*gmxConnectionPool, error) {
	callers := returnCallers(rpcUrls, gmxReaderContractAddress)
	if len(callers) == 0 {
		return nil, errors.New("no valid rpcs provided")
	}
	return &gmxConnectionPool{
		connections: callers,
		index:       0,
	}, nil
}

func returnCallers(rpcUrls []string, gmxReaderContractAddress string) []*gmx_abis.GmxAbisCaller {
	var callers []*gmx_abis.GmxAbisCaller
	readerAddress := common.HexToAddress(gmxReaderContractAddress)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, url := range rpcUrls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			web3, err := ethclient.Dial(url)
			if err != nil {
				return
			}
			_, err = web3.BlockNumber(context.Background())
			if err != nil {
				return
			}
			client, err := gmx_abis.NewGmxAbisCaller(readerAddress, web3)
			if err != nil {
				return
			}
			mu.Lock()
			callers = append(callers, client)
			mu.Unlock()
		}(url)
	}

	wg.Wait()
	return callers
}

func (p *gmxConnectionPool) numCallers() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return len(p.connections)
}

func (p *gmxConnectionPool) nextCaller() *gmx_abis.GmxAbisCaller {
	p.mu.Lock()
	defer p.mu.Unlock()
	if len(p.connections) == 0 {
		return nil
	}
	caller := p.connections[p.index]
	p.index = (p.index + 1) % len(p.connections)
	return caller
}

func (p *gmxConnectionPool) getPositions(userAddress common.Address, gmxDataStoreContractAddress string) []gmx_abis.PositionProps {
	dataStoreAddress := common.HexToAddress(gmxDataStoreContractAddress)

	for {
		caller := p.nextCaller()
		if caller == nil {
			continue
		}
		positionProps, err := caller.GetAccountPositions(
			&bind.CallOpts{Pending: false},
			dataStoreAddress,
			userAddress,
			&big.Int{},
			utils.MaxBigint(),
		)
		if err == nil {
			return positionProps
		}
	}
}
