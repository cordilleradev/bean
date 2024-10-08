package common

import "github.com/cordilleradev/bean/common/types"

type FuturesClient interface {
	ExchangeName() string
	GetSupportedMarginTypes() []types.MarginType
	GetLeaderboardPeriods() types.SupportedPeriods
	GetSupportedLeaderboardFields() []types.LeaderboardField
	GetLeaderboard(period string) ([]types.Trader, *types.APIError)
	FetchPositions(userId string) ([]types.FuturesPosition, *types.APIError)
	StreamPositions(userId string, refreshWaitSeconds float64, positionStream chan types.FuturesResponse) *types.APIError
	CancelStream(userId string) *types.APIError
}
