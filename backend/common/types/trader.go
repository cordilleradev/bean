package types

import (
	"fmt"
)

type Trader struct {
	UserId            string   `json:"user_id"`
	PeriodPnlPercent  *float64 `json:"period_pnl_percent,omitempty"`
	PeriodPnlAbsolute *float64 `json:"period_pnl_absolute,omitempty"`
	TotalTrades       *int     `json:"total_trades,omitempty"`
	Wins              *int     `json:"wins,omitempty"`
	Volume            *float64 `json:"volume,omitempty"`
	AvgWin            *float64 `json:"avg_win,omitempty"`
	AvgLoss           *float64 `json:"avg_loss,omitempty"`
	AccountValue      *float64 `json:"account_value,omitempty"`
}

type LeaderboardField string

const (
	PeriodPnlPercent  LeaderboardField = "period_pnl_percent"
	PeriodPnlAbsolute LeaderboardField = "period_pnl_absolute"
	TotalTrades       LeaderboardField = "total_trades"
	Wins              LeaderboardField = "wins"
	Volume            LeaderboardField = "volume"
	AvgWin            LeaderboardField = "avg_win"
	AvgLoss           LeaderboardField = "avg_loss"
	AccountValue      LeaderboardField = "account_value"
)

func (lf LeaderboardField) String() string {
	return string(lf)
}

func LeaderboardFieldFromString(s string) (LeaderboardField, *APIError) {
	switch s {
	case "period_pnl_percent":
		return PeriodPnlPercent, nil
	case "period_pnl_absolute":
		return PeriodPnlAbsolute, nil
	case "total_trades":
		return TotalTrades, nil
	case "wins":
		return Wins, nil
	case "volume":
		return Volume, nil
	case "avg_win":
		return AvgWin, nil
	case "avg_loss":
		return AvgLoss, nil
	case "account_value":
		return AccountValue, nil
	default:
		return "", NewAPIError(
			400,
			"invalid_leaderboard_field",
			fmt.Sprintf("'%s' is not a valid leaderboard field", s),
		)
	}
}
