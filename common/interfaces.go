package common

import "github.com/cordilleradev/bean/common/types"

type FuturesClient interface {
	GetSupportedMarginTypes() []string
	GetLeaderboardPeriods() []string
	GetLeaderboard(period string, limit int) ([]types.Trader, *types.APIError)
	FetchPositions(userId string) ([]types.FuturesPosition, *types.APIError)
}
