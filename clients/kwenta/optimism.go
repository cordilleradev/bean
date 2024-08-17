package kwenta

import "github.com/cordilleradev/bean/common/types"

type KwentaOptimisimClient struct {
	kwentaClient
}

func NewKwentaOptimismClient() (*KwentaOptimisimClient, error) {
	opClient, err := newKwentaClient(
		"kwenta-v2-optimism",
		leaderboardApiUrlOp,
		"futuresStats",
		types.Isolated,
	)
	if err != nil {
		return nil, err
	}
	return &KwentaOptimisimClient{
		kwentaClient: *opClient,
	}, nil
}

func (c *KwentaOptimisimClient) FetchPositions(userId string) ([]types.FuturesResponse, *types.APIError) {
	// Implementation for fetching positions will go here
	return nil, nil
}
