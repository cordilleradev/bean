package avantis

import (
	"github.com/cordilleradev/bean/common/types"
)

type AvantisClient struct {
}

func NewAvantisClient() (*AvantisClient, error) {
	return &AvantisClient{}, nil
}

func (a *AvantisClient) ExchangeName() string {
	// Implementation here
	return "avantis-v1-base"
}

func (a *AvantisClient) GetSupportedMarginTypes() []types.MarginType {
	return []types.MarginType{types.Isolated}
}

func (a *AvantisClient) GetLeaderboardPeriods() types.SupportedPeriods {
	return types.NewSupportedPeriods([]string{"total"}, nil)
}

func (a *AvantisClient) GetSupportedLeaderboardFields() []types.LeaderboardField {
	return []types.LeaderboardField{
		types.PeriodPnlPercent,
		types.PeriodPnlAbsolute,
		types.Volume,
		types.TotalTrades,
		types.Wins,
	}
}

func (a *AvantisClient) GetLeaderboard(period string) ([]types.Trader, *types.APIError) {
	resp, err := fetchLeaderBoard()
	if err != nil {
		return nil, types.FailedLeaderboardCall(err)
	}

	traders := make([]types.Trader, len(resp.LeaderBoard))

	for i, u := range resp.LeaderBoard {
		traders[i] = types.Trader{
			UserId:            u.Trader,
			PeriodPnlPercent:  u.TotalProfitPercentage / 100,
			PeriodPnlAbsolute: u.TotalProfits,
			Volume:            u.TotalPositionSizes,
			TotalTrades:       u.TotalTrades,
			Wins:              u.TotalWins,
		}
	}

	return traders, nil
}

func (a *AvantisClient) FetchPositions(userId string) ([]types.FuturesPosition, *types.APIError) {
	// Implementation here
	return nil, nil
}
