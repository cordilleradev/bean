package types

type Trader struct {
	UserId            string  `json:"user_id"`
	PeriodPnlPercent  float64 `json:"period_pnl_percent,omitempty"`
	PeriodPnlAbsolute float64 `json:"period_pnl_absolute,omitempty"`
	TotalTrades       int     `json:"total_trades,omitempty"`
	Wins              int     `json:"wins,omitempty"`
	Volume            float64 `json:"volume,omitempty"`
	AvgWin            float64 `json:"avg_win,omitempty"`
	AvgLoss           float64 `json:"avg_loss,omitempty"`
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
)
