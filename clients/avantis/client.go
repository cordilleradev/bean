package avantis

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type leaderBoardEntry struct {
	ID                    string    `json:"_id"`
	Trader                string    `json:"trader"`
	TotalWins             int       `json:"totalWins"`
	TotalProfits          float64   `json:"totalProfits"`
	TotalBase             float64   `json:"totalBase"`
	TotalTrades           int       `json:"totalTrades"`
	TotalPositionSizes    float64   `json:"totalPositionSizes"`
	TotalProfitPercentage float64   `json:"totalProfitPercentage"`
	WinRate               float64   `json:"winRate"`
	Rank                  int       `json:"rank"`
	LastUpdated           time.Time `json:"lastUpdated"`
	V                     int       `json:"__v"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
}

type leaderBoardResponse struct {
	LeaderBoard []leaderBoardEntry `json:"leaderBoard"`
}

func fetchLeaderBoard() (*leaderBoardResponse, error) {
	resp, err := http.Get(leaderboardUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	var leaderBoardResponse leaderBoardResponse
	if err := json.NewDecoder(resp.Body).Decode(&leaderBoardResponse); err != nil {
		return nil, err
	}
	return &leaderBoardResponse, nil
}
