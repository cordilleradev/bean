package utils

import (
	"context"
	"errors"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EvmRpcPool[T any] struct {
	connections []T
	rpcUrls     []string
	mu          sync.Mutex
	index       int
}

// Initialize CommonPoolHelper with a constructor
func NewCommonPoolHelper[T any](rpcUrls []string, contractAddress common.Address, clientCtor func(common.Address, *ethclient.Client) (T, error)) (*EvmRpcPool[T], error) {
	callers := returnCallers(rpcUrls, contractAddress, clientCtor)
	if len(callers) == 0 {
		return nil, errors.New("no valid RPCs provided")
	}
	return &EvmRpcPool[T]{
		connections: callers,
		rpcUrls:     rpcUrls,
		index:       0,
	}, nil
}

// Create RPC clients from URLs
func returnCallers[T any](rpcUrls []string, contractAddress common.Address, clientCtor func(common.Address, *ethclient.Client) (T, error)) []T {
	var callers []T
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
			client, err := clientCtor(contractAddress, web3)
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

// Return the number of RPC clients in the pool
func (p *EvmRpcPool[T]) NumCallers() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return len(p.connections)
}

// Return the next RPC client in a round-robin fashion
func (p *EvmRpcPool[T]) NextCaller() T {
	p.mu.Lock()
	defer p.mu.Unlock()
	if len(p.connections) == 0 {
		var zero T
		return zero
	}
	caller := p.connections[p.index]
	p.index = (p.index + 1) % len(p.connections)
	return caller
}

// Show the RPC URLs that are being connected to successfully
func (p *EvmRpcPool[T]) ShowRpcs() []string {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.rpcUrls
}
