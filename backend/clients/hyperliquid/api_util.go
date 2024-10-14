package hyperliquid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type windowPerformance struct {
	Timeframe string `json:"timeframe"`
	Pnl       string `json:"pnl"`
	Roi       string `json:"roi"`
	Vlm       string `json:"vlm"`
}

type leaderboardRow struct {
	EthAddress         string          `json:"ethAddress"`
	AccountValue       string          `json:"accountValue"`
	WindowPerformances [][]interface{} `json:"windowPerformances"`
	Prize              int             `json:"prize"`
	DisplayName        *string         `json:"displayName"`
}

type rawLeaderboardResponse struct {
	LeaderboardRows []leaderboardRow `json:"leaderboardRows"`
}

type userPosition struct {
	Type     string          `json:"type"`
	Position positionDetails `json:"position"`
}

type positionDetails struct {
	Coin           string  `json:"coin"`
	Szi            string  `json:"szi"`
	Leverage       lever   `json:"leverage"`
	EntryPx        string  `json:"entryPx"`
	PositionValue  string  `json:"positionValue"`
	UnrealizedPnl  string  `json:"unrealizedPnl"`
	ReturnOnEquity string  `json:"returnOnEquity"`
	LiquidationPx  *string `json:"liquidationPx"`
	MarginUsed     string  `json:"marginUsed"`
	MaxLeverage    int     `json:"maxLeverage"`
	CumFunding     funding `json:"cumFunding"`
}

type lever struct {
	Type   string  `json:"type"`
	Value  int     `json:"value"`
	RawUsd *string `json:"rawUsd,omitempty"`
}

type funding struct {
	AllTime     string `json:"allTime"`
	SinceOpen   string `json:"sinceOpen"`
	SinceChange string `json:"sinceChange"`
}

type clearingHouseState struct {
	AssetPositions             []userPosition `json:"assetPositions"`
	CrossMaintenanceMarginUsed string         `json:"crossMaintenanceMarginUsed"`
	CrossMarginSummary         marginSummary  `json:"crossMarginSummary"`
	MarginSummary              marginSummary  `json:"marginSummary"`
	Time                       int64          `json:"time"`
	Withdrawable               string         `json:"withdrawable"`
}

type marginSummary struct {
	AccountValue    string `json:"accountValue"`
	TotalNtlPos     string `json:"totalNtlPos"`
	TotalRawUsd     string `json:"totalRawUsd"`
	TotalMarginUsed string `json:"totalMarginUsed"`
}

type traderPerformance struct {
	AccountValue float64
	Day          periodPerformance
	Week         periodPerformance
	Month        periodPerformance
	AllTime      periodPerformance
}

type periodPerformance struct {
	Pnl float64 `json:"pnl"`
	Roi float64 `json:"roi"`
	Vlm float64 `json:"vlm"`
}

type requestWeight int

const (
	openOrdersWeight         requestWeight = 20
	allMidsWeight            requestWeight = 2
	priceFetchWeight         requestWeight = 2
	clearingHouseStateWeight               = 2
)

func parseStringToFloat(str string) float64 {
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Printf("Error converting string to float64: %v", err)
		return 0.0
	}
	return value
}

func leaderboardCall(client *http.Client) (map[string]traderPerformance, error) {
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

	var rawResponse rawLeaderboardResponse
	if err := json.Unmarshal(body, &rawResponse); err != nil {
		return nil, err
	}

	performanceList := make(
		map[string]traderPerformance,
		len(rawResponse.LeaderboardRows),
	)
	for _, item := range rawResponse.LeaderboardRows {
		accountValue := parseStringToFloat(item.AccountValue)
		windowPerformances := item.WindowPerformances
		traderPerformance := traderPerformance{
			Day: periodPerformance{
				Pnl: parseStringToFloat(windowPerformances[0][1].(map[string]interface{})["pnl"].(string)),
				Roi: parseStringToFloat(windowPerformances[0][1].(map[string]interface{})["roi"].(string)),
				Vlm: parseStringToFloat(windowPerformances[0][1].(map[string]interface{})["vlm"].(string)),
			},
			Week: periodPerformance{
				Pnl: parseStringToFloat(windowPerformances[1][1].(map[string]interface{})["pnl"].(string)),
				Roi: parseStringToFloat(windowPerformances[1][1].(map[string]interface{})["roi"].(string)),
				Vlm: parseStringToFloat(windowPerformances[1][1].(map[string]interface{})["vlm"].(string)),
			},
			Month: periodPerformance{
				Pnl: parseStringToFloat(windowPerformances[2][1].(map[string]interface{})["pnl"].(string)),
				Roi: parseStringToFloat(windowPerformances[2][1].(map[string]interface{})["roi"].(string)),
				Vlm: parseStringToFloat(windowPerformances[2][1].(map[string]interface{})["vlm"].(string)),
			},
			AllTime: periodPerformance{
				Pnl: parseStringToFloat(windowPerformances[3][1].(map[string]interface{})["pnl"].(string)),
				Roi: parseStringToFloat(windowPerformances[3][1].(map[string]interface{})["roi"].(string)),
				Vlm: parseStringToFloat(windowPerformances[3][1].(map[string]interface{})["vlm"].(string)),
			},
			AccountValue: accountValue,
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
