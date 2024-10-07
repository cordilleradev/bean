package gains

import (
	"fmt"
	"strings"

	"github.com/cordilleradev/bean/common/types"
	"github.com/cordilleradev/bean/common/utils"
)

type GainsClient struct{}

func NewGainsClient() (*GainsClient, error) {
	// Perform any necessary initialization here
	return &GainsClient{}, nil
}
func (g *GainsClient) ExchangeName() string {
	return "gains-v1-arbitrum"
}

func (g *GainsClient) GetSupportedMarginTypes() []types.MarginType {
	return []types.MarginType{types.Cross, types.Isolated}
}

func (g *GainsClient) GetLeaderboardPeriods() types.SupportedPeriods {
	return types.NewSupportedPeriods([]string{"1d", "7d", "30d", "90d"}, nil)
}

func (g *GainsClient) GetSupportedLeaderboardFields() []types.LeaderboardField {
	return []types.LeaderboardField{
		types.PeriodPnlAbsolute,
		types.TotalTrades,
		types.Wins,
		types.AvgWin,
		types.AvgLoss,
	}
}

func (g *GainsClient) GetLeaderboard(period string) ([]types.Trader, *types.APIError) {
	allowedLeaderboardPeriods := g.GetLeaderboardPeriods().FixedPeriods
	if !utils.ContainsString(allowedLeaderboardPeriods, period) {
		return nil, types.InvalidPeriodError(
			period,
			fmt.Sprintf(
				"allowed periods are %s",
				strings.Join(
					allowedLeaderboardPeriods,
					", "),
			),
		)
	}

	leaderboardKey := period[:len(period)-1]

	leaderboardResponse, err := fetchLeaderboard()
	if err != nil {
		return nil, types.FailedLeaderboardCall(err)
	}

	traders := make([]types.Trader, len(leaderboardResponse[leaderboardKey]))

	for i, u := range leaderboardResponse[leaderboardKey] {
		traders[i] = types.Trader{
			UserId:            u.Address,
			PeriodPnlAbsolute: u.TotalPnlUsd,
			AvgWin:            u.AvgWin,
			AvgLoss:           u.AvgLoss,
			TotalTrades:       u.Count,
			Wins:              u.CountWin,
		}
	}

	return traders, nil
}

func (g *GainsClient) FetchPositions(userId string) ([]types.FuturesPosition, *types.APIError) {
	// Implementation here

	return nil, nil
}
