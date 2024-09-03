package common

import "github.com/cordilleradev/bean/common/types"

type FuturesClient interface {
	ExchangeName() string
	GetSupportedMarginTypes() []types.MarginType
	GetLeaderboardPeriods() types.SupportedPeriods
	GetSupportedLeaderboardFields() []types.LeaderboardField
	GetLeaderboard(period string) ([]types.Trader, *types.APIError)
	FetchPositions(userId string) (*types.FuturesResponse, *types.APIError)
}
