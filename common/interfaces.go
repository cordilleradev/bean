package common

import "github.com/cordilleradev/bean/common/types"

type FuturesClient interface {
	ExchangeName() string
	GetSupportedMarginTypes() []types.MarginType
	GetLeaderboardPeriods() types.SupportedPeriods
	GetSupportedLeaderboardFields() []types.LeaderboardField
	FetchPositions(userId string) ([]types.FuturesPosition, *types.APIError)
	GetLeaderboard(period string, sortBy types.LeaderboardField, orderIsAsc bool) ([]types.Trader, *types.APIError)
	StreamPositions(userId string, refreshWaitSeconds float64, positionStream chan types.FuturesResponse) *types.APIError
	StreamExists(userId string) bool
	CancelStream(userId string) *types.APIError
}
