package kwenta

import (
	"github.com/cordilleradev/bean/common/types"
)

type KwentaBaseClient struct {
	kwentaClient
}

func NewKwentaBaseClient() (*KwentaBaseClient, error) {
	baseClient, err := newKwentaClient(
		"kwenta-v3-base",
		leaderboardApiUrlBase,
		"perpsV3Stats",
		types.Cross,
	)
	if err != nil {
		return nil, err
	}

	return &KwentaBaseClient{
		kwentaClient: *baseClient,
	}, nil

}

func (c *KwentaBaseClient) FetchPositions(userId string) (*types.FuturesResponse, *types.APIError) {
	return nil, nil
}
