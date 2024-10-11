package utils

import (
	"fmt"
	"sort"

	"github.com/cordilleradev/bean/common/types"
)

func LeaderboardFieldSortAllowed(field string, allowedFields []types.LeaderboardField) (*types.LeaderboardField, *types.APIError) {
	leaderboardField, err := types.LeaderboardFieldFromString(field)

	if err != nil {
		return nil, err
	}

	for _, i := range allowedFields {
		if leaderboardField == i {
			return &i, nil
		}
	}

	return nil, types.NewAPIError(400, "field_not_supported",
		fmt.Sprintf("'%s' is not a supported leaderboard field for this exchange", field),
	)
}

func SortByFields(field types.LeaderboardField, traders []types.Trader, isAsc bool) []types.Trader {
	comparator := func(i, j int) bool { return false }
	switch field {
	case types.PeriodPnlPercent:
		if isAsc {
			comparator = func(i, j int) bool {
				return traders[i].PeriodPnlPercent < traders[j].PeriodPnlPercent
			}
		} else {
			comparator = func(i, j int) bool {
				return traders[i].PeriodPnlPercent > traders[j].PeriodPnlPercent
			}
		}
	case types.PeriodPnlAbsolute:
		if isAsc {
			comparator = func(i, j int) bool {
				return traders[i].PeriodPnlAbsolute < traders[j].PeriodPnlAbsolute
			}
		} else {
			comparator = func(i, j int) bool {
				return traders[i].PeriodPnlAbsolute > traders[j].PeriodPnlAbsolute
			}
		}
	case types.TotalTrades:
		if isAsc {
			comparator = func(i, j int) bool {
				return traders[i].TotalTrades < traders[j].TotalTrades
			}
		} else {
			comparator = func(i, j int) bool {
				return traders[i].TotalTrades > traders[j].TotalTrades
			}
		}
	case types.Wins:
		if isAsc {
			comparator = func(i, j int) bool {
				return traders[i].Wins < traders[j].Wins
			}
		} else {
			comparator = func(i, j int) bool {
				return traders[i].Wins > traders[j].Wins
			}
		}
	case types.Volume:
		if isAsc {
			comparator = func(i, j int) bool {
				return traders[i].Volume < traders[j].Volume
			}
		} else {
			comparator = func(i, j int) bool {
				return traders[i].Volume > traders[j].Volume
			}
		}
	case types.AvgWin:
		if isAsc {
			comparator = func(i, j int) bool {
				return traders[i].AvgWin < traders[j].AvgWin
			}
		} else {
			comparator = func(i, j int) bool {
				return traders[i].AvgWin > traders[j].AvgWin
			}
		}
	case types.AvgLoss:
		if isAsc {
			comparator = func(i, j int) bool {
				return traders[i].AvgLoss < traders[j].AvgLoss
			}
		} else {
			comparator = func(i, j int) bool {
				return traders[i].AvgLoss > traders[j].AvgLoss
			}
		}
	default:
		return traders
	}
	sort.SliceStable(traders, comparator)
	return traders
}
