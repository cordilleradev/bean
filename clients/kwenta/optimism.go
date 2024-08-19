package kwenta

import "github.com/cordilleradev/bean/common/types"

type KwentaOptimismClient struct {
	kwentaClient
}

func NewKwentaOptimismClient() (*KwentaOptimismClient, error) {
	opClient, err := newKwentaClient(
		"kwenta-v2-optimism",
		leaderboardApiUrlOp,
		"futuresStats",
		types.Isolated,
	)
	if err != nil {
		return nil, err
	}
	return &KwentaOptimismClient{
		kwentaClient: *opClient,
	}, nil
}

func (c *KwentaOptimismClient) FetchPositions(userId string) ([]types.FuturesResponse, *types.APIError) {
	// Implementation for fetching positions will go here
	return nil, nil
}
