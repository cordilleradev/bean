package gains

import (
	"encoding/json"
	"io"
	"net/http"
)

type leaderboardEntry struct {
	Address     string  `json:"address"`
	Count       int     `json:"count"`
	CountWin    int     `json:"count_win"`
	CountLoss   int     `json:"count_loss"`
	AvgWin      float64 `json:"avg_win"`
	AvgLoss     float64 `json:"avg_loss"`
	TotalPnl    float64 `json:"total_pnl"`
	TotalPnlUsd float64 `json:"total_pnl_usd"`
}

type leaderboardResponse map[string][]leaderboardEntry

func fetchLeaderboard() (leaderboardResponse, error) {
	resp, err := http.Get(leaderboardUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var leaderboard leaderboardResponse
	err = json.Unmarshal(body, &leaderboard)
	if err != nil {
		return nil, err
	}

	return leaderboard, nil
}
