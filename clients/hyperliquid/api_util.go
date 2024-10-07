package hyperliquid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type WindowPerformance struct {
	Timeframe string `json:"timeframe"`
	Pnl       string `json:"pnl"`
	Roi       string `json:"roi"`
	Vlm       string `json:"vlm"`
}

type LeaderboardRow struct {
	EthAddress         string          `json:"ethAddress"`
	AccountValue       string          `json:"accountValue"`
	WindowPerformances [][]interface{} `json:"windowPerformances"`
	Prize              int             `json:"prize"`
	DisplayName        *string         `json:"displayName"`
}

type RawLeaderboardResponse struct {
	LeaderboardRows []LeaderboardRow `json:"leaderboardRows"`
}

type UserPosition struct {
	Type     string          `json:"type"`
	Position PositionDetails `json:"position"`
}

type PositionDetails struct {
	Coin           string  `json:"coin"`
	Szi            string  `json:"szi"`
	Leverage       Lever   `json:"leverage"`
	EntryPx        string  `json:"entryPx"`
	PositionValue  string  `json:"positionValue"`
	UnrealizedPnl  string  `json:"unrealizedPnl"`
	ReturnOnEquity string  `json:"returnOnEquity"`
	LiquidationPx  *string `json:"liquidationPx"`
	MarginUsed     string  `json:"marginUsed"`
	MaxLeverage    int     `json:"maxLeverage"`
	CumFunding     Funding `json:"cumFunding"`
}

type Lever struct {
	Type   string  `json:"type"`
	Value  int     `json:"value"`
	RawUsd *string `json:"rawUsd,omitempty"`
}

type Funding struct {
	AllTime     string `json:"allTime"`
	SinceOpen   string `json:"sinceOpen"`
	SinceChange string `json:"sinceChange"`
}

type ClearingHouseState struct {
	AssetPositions             []UserPosition `json:"assetPositions"`
	CrossMaintenanceMarginUsed string         `json:"crossMaintenanceMarginUsed"`
	CrossMarginSummary         MarginSummary  `json:"crossMarginSummary"`
	MarginSummary              MarginSummary  `json:"marginSummary"`
	Time                       int64          `json:"time"`
	Withdrawable               string         `json:"withdrawable"`
}

type MarginSummary struct {
	AccountValue    string `json:"accountValue"`
	TotalNtlPos     string `json:"totalNtlPos"`
	TotalRawUsd     string `json:"totalRawUsd"`
	TotalMarginUsed string `json:"totalMarginUsed"`
}

type TraderPerformance struct {
	Day     PeriodPerformance
	Week    PeriodPerformance
	Month   PeriodPerformance
	AllTime PeriodPerformance
}

type PeriodPerformance struct {
	Pnl float64 `json:"pnl"`
	Roi float64 `json:"roi"`
	Vlm float64 `json:"vlm"`
}

type RequestWeight int

const (
	OpenOrdersWeight         RequestWeight = 20
	AllMidsWeight            RequestWeight = 2
	PriceFetchWeight         RequestWeight = 2
	ClearingHouseStateWeight               = 2
)

func parseStringToFloat(str string) float64 {
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Printf("Error converting string to float64: %v", err)
		return 0.0
	}
	return value
}

func leaderboardCall(client *http.Client) (map[string]TraderPerformance, error) {
	req, err := http.NewRequest("GET", leaderboardUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Origin", "https://app.hyperliquid.xyz")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch leaderboard: %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rawResponse RawLeaderboardResponse
	if err := json.Unmarshal(body, &rawResponse); err != nil {
		return nil, err
	}

	performanceList := make(
		map[string]TraderPerformance,
		len(rawResponse.LeaderboardRows),
	)
	for _, item := range rawResponse.LeaderboardRows {
		windowPerformances := item.WindowPerformances
		traderPerformance := TraderPerformance{
			Day: PeriodPerformance{
				Pnl: parseStringToFloat(windowPerformances[0][1].(map[string]interface{})["pnl"].(string)),
				Roi: parseStringToFloat(windowPerformances[0][1].(map[string]interface{})["roi"].(string)),
				Vlm: parseStringToFloat(windowPerformances[0][1].(map[string]interface{})["vlm"].(string)),
			},
			Week: PeriodPerformance{
				Pnl: parseStringToFloat(windowPerformances[1][1].(map[string]interface{})["pnl"].(string)),
				Roi: parseStringToFloat(windowPerformances[1][1].(map[string]interface{})["roi"].(string)),
				Vlm: parseStringToFloat(windowPerformances[1][1].(map[string]interface{})["vlm"].(string)),
			},
			Month: PeriodPerformance{
				Pnl: parseStringToFloat(windowPerformances[2][1].(map[string]interface{})["pnl"].(string)),
				Roi: parseStringToFloat(windowPerformances[2][1].(map[string]interface{})["roi"].(string)),
				Vlm: parseStringToFloat(windowPerformances[2][1].(map[string]interface{})["vlm"].(string)),
			},
			AllTime: PeriodPerformance{
				Pnl: parseStringToFloat(windowPerformances[3][1].(map[string]interface{})["pnl"].(string)),
				Roi: parseStringToFloat(windowPerformances[3][1].(map[string]interface{})["roi"].(string)),
				Vlm: parseStringToFloat(windowPerformances[3][1].(map[string]interface{})["vlm"].(string)),
			},
		}
		performanceList[item.EthAddress] = traderPerformance
	}
	return performanceList, nil
}

func makeInfoRequest[T any](client *http.Client, url string, requestBody map[string]string) (T, error) {
	var result T

	reqBody, err := json.Marshal(requestBody)
	if err != nil {
		return result, fmt.Errorf("failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", url, io.NopCloser(bytes.NewReader(reqBody)))
	if err != nil {
		return result, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return result, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("bad response status: %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("failed to read response body: %v", err)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return result, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return result, nil
}
