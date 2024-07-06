package gmx

import (
	"github.com/cordilleradev/bean/common"
	"github.com/cordilleradev/bean/common/types"
)

type GmxClient struct {
	common.FuturesClient
}

func (g *GmxClient) GetSupportedMarginTypes() []string {
	return supportedMarginTypes
}

func (g *GmxClient) GetLeaderboardPeriods() []string {
	return periods
}

func (g *GmxClient) GetLeaderboard(period string, limit int) ([]types.Trader, *types.APIError) {

}

func (g *GmxClient) FetchPositions(userId string) ([]types.FuturesPosition, *types.APIError) {

}
