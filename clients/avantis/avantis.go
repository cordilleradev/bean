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
	return types.SupportedPeriods{
		FixedPeriods:  []string{"total"},
		CustomPeriods: nil,
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
		}
	}

	return traders, nil
}

func (a *AvantisClient) FetchPositions(userId string) ([]types.FuturesPosition, *types.APIError) {
	// Implementation here
	return nil, nil
}
