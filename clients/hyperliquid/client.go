package hyperliquid

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type request struct {
	Type string `json:"type"`
}

type performanceStats struct {
}

type response struct {
	LeaderboardRows []struct {
		EthAddress         string          `json:"ethAddress"`
		AccountValue       string          `json:"accountValue"`
		WindowPerformances [][]interface{} `json:"windowPerformances"`
		Prize              int             `json:"prize"`
		DisplayName        interface{}     `json:"displayName"`
	} `json:"leaderboardRows"`
}

func fetchPeriodStats() (*response, error) {
	reqBody := request{
		Type: "leaderboard",
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api-ui.hyperliquid.xyz/info", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("authority", "api-ui.hyperliquid.xyz")
	req.Header.Set("method", "POST")
	req.Header.Set("path", "/info")
	req.Header.Set("scheme", "https")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-encoding", "gzip, deflate, br, zstd")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("content-length", "22")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("origin", "https://app.hyperliquid.xyz")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("referer", "https://app.hyperliquid.xyz/")
	req.Header.Set("sec-ch-ua", `"Not)A;Brand";v="99", "Google Chrome";v="127", "Chromium";v="127"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch leaderboard")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res response
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
