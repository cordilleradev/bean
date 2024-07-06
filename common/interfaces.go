package common

import "github.com/cordilleradev/bean/common/types"

type FuturesClient interface {
	GetLeaderboardPeriods() []string
	GetLeaderboard(period string) ([]types.Trader, *types.APIError)
	FetchPositions(userId string) ([]types.FuturesPosition, *types.APIError)
}
