package gmx

import (
	"math/big"

	gmx_abis "github.com/cordilleradev/bean/common/abigen/gmx"
	"github.com/cordilleradev/bean/common/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type gmxConnectionPool struct {
	pool *utils.EvmRpcPool[*gmx_abis.GmxAbisCaller]
}

func newGmxConnectionPool(rpcUrls []string, gmxReaderContractAddress string) (*gmxConnectionPool, error) {
	readerAddress := common.HexToAddress(gmxReaderContractAddress)
	clientCtor := func(
		address common.Address,
		client *ethclient.Client,
	) (*gmx_abis.GmxAbisCaller, error) {
		return gmx_abis.NewGmxAbisCaller(address, client)
	}
	pool, err := utils.NewCommonPoolHelper(rpcUrls, readerAddress, clientCtor)
	if err != nil {
		return nil, err
	}
	return &gmxConnectionPool{
		pool: pool,
	}, nil
}

func (p *gmxConnectionPool) GetPositions(userAddress common.Address, gmxDataStoreContractAddress string) []gmx_abis.PositionProps {
	dataStoreAddress := common.HexToAddress(gmxDataStoreContractAddress)

	for {
		caller := p.pool.NextCaller()
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
