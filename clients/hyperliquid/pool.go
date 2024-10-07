package hyperliquid

import (
	"net/http"
	"sync"

	"github.com/cordilleradev/bean/common/types"
)

type HyperLiquidPool struct {
	clients    []*HyperLiquidManager
	priceCache *HyperLiquidPriceCache
	mu         sync.Mutex

	cacheLeaderboard bool
	leaderboardCache map[string]TraderPerformance
	periodSession    []string
	roundRobinIndex  int
	httpClient       *http.Client
}

func NewHyperLiquidPool(cacheLeaderboard bool) (*HyperLiquidPool, error) {

	priceCache, priceCacheErr := NewHyperLiquidPriceCache(
		hyperliquidApiUrl,
		hyperliquidWsUrl,
	)

	if priceCacheErr != nil {
		return nil, priceCacheErr
	}

	uiApiClient := NewHyperLiquidManager(1200, hyperLiquidUiApi, false)
	defaultClient := NewHyperLiquidManager(1200, hyperliquidApiUrl, true)

	clientList := []*HyperLiquidManager{
		defaultClient,
		uiApiClient,
	}

	return &HyperLiquidPool{
		clients:          clientList,
		priceCache:       priceCache,
		cacheLeaderboard: cacheLeaderboard,
		leaderboardCache: make(map[string]TraderPerformance),
		httpClient:       &http.Client{},
	}, nil
}

func (hp *HyperLiquidPool) shouldRefreshCache(period string) bool {
	hp.mu.Lock()
	defer hp.mu.Unlock()

	needsRefresh := len(hp.leaderboardCache) == 0 ||
		len(hp.periodSession) == 4 ||
		contains(hp.periodSession, period)
	if needsRefresh {
		hp.periodSession = nil
	} else {
		hp.periodSession = append(hp.periodSession, period)
	}
	return needsRefresh
}

func contains(slice []string, item string) bool {
	for _, elem := range slice {
		if elem == item {
			return true
		}
	}
	return false
}

func (hp *HyperLiquidPool) Leaderboard(period string) ([]types.Trader, error) {
	var response map[string]TraderPerformance
	var err error

	if hp.cacheLeaderboard && !hp.shouldRefreshCache(period) {
		hp.mu.Lock()
		response = hp.leaderboardCache
		hp.mu.Unlock()
	} else {
		response, err = leaderboardCall(hp.httpClient)
		if err != nil {
			return nil, err
		}
		hp.mu.Lock()
		hp.leaderboardCache = response
		hp.mu.Unlock()
	}

	traders := make([]types.Trader, len(response))

	index := 0
	for userId, item := range response {

		periodPnlPercent := 0.0
		periodPnlAbsolute := 0.0
		periodVolume := 0.0

		switch period {
		case "1d":
			periodPnlAbsolute = item.Day.Pnl
			periodPnlPercent = item.Day.Roi
			periodVolume = item.Day.Vlm
		case "7d":
			periodPnlAbsolute = item.Week.Pnl
			periodPnlPercent = item.Week.Roi
			periodVolume = item.Week.Vlm
		case "30d":
			periodPnlAbsolute = item.Month.Pnl
			periodPnlPercent = item.Month.Roi
			periodVolume = item.Month.Vlm
		case "total":
			periodPnlAbsolute = item.AllTime.Pnl
			periodPnlPercent = item.AllTime.Roi
			periodVolume = item.AllTime.Vlm
		}

		traders[index] = types.Trader{
			UserId:            userId,
			PeriodPnlPercent:  periodPnlPercent,
			PeriodPnlAbsolute: periodPnlAbsolute,
			Volume:            periodVolume,
		}

		index++
	}

	return traders, nil
}

func (hp *HyperLiquidPool) cycle() *HyperLiquidManager {

	bestAllowance := 0
	bestManager := hp.clients[0]

	for _, m := range hp.clients[1:] {
		if (m.allowancePerMinute - m.resetRequests()) > bestAllowance {
			bestManager = m
		}
		// fmt.Println("Allowance left", bestManager.ApiUrl, m.allowancePerMinute-m.resetRequests())
	}
	// fmt.Println("Going with", bestManager.ApiUrl)
	return bestManager
}

func (hp *HyperLiquidPool) Positions(userID string) ([]types.FuturesPosition, error) {
	var err error

	for i := 0; i < len(hp.clients); i++ {
		manager := hp.cycle()
		manager.waitForToken(ClearingHouseStateWeight)

		resp, err := makeInfoRequest[ClearingHouseState](
			manager.HttpClient,
			manager.ApiUrl,
			map[string]string{
				"type": "clearinghouseState",
				"user": userID,
			})

		if err != nil {
			continue
		}

		positions := make([]types.FuturesPosition, len(resp.AssetPositions))

		for index, position := range resp.AssetPositions {
			entryPrice := parseStringToFloat(position.Position.EntryPx)
			unrealizedPnl := parseStringToFloat(position.Position.UnrealizedPnl)
			positionValue := parseStringToFloat(position.Position.PositionValue)
			marginUsed := parseStringToFloat(position.Position.MarginUsed)
			leverageAmount := float64(position.Position.Leverage.Value)

			size := parseStringToFloat(position.Position.Szi)
			direction := types.Long
			if size < 0 {
				direction = types.Short
			}

			positionDetails := types.FuturesPosition{
				Market:         position.Position.Coin + "-USD",
				EntryPrice:     entryPrice,
				CurrentPrice:   hp.priceCache.GetValue(position.Position.Coin), // Update with current price from relevant source
				Status:         types.Open,                                     // Assuming status is open for placeholder
				Direction:      direction,                                      // Update with actual direction from relevant data
				SizeUsd:        positionValue,
				UnrealizedPnl:  unrealizedPnl,
				LeverageAmount: leverageAmount,
			}

			switch position.Position.Leverage.Type {
			case "isolated":
				positionDetails.MarginType = types.Isolated
				positionDetails.CollateralTokenAmountUsd = marginUsed

			case "cross":
				positionDetails.MarginType = types.Cross
				positionDetails.HealthRatio = parseStringToFloat(position.Position.ReturnOnEquity) // Assuming Health Ratio calculated similarly
				positionDetails.CrossMarginShare = marginUsed                                      // Example share of cross margin
				positionDetails.FreeCollateral = parseStringToFloat(resp.Withdrawable)             // Assuming freeCollateral is related to withdrawable balance
			}
			positions[index] = positionDetails
		}

		return positions, nil
	}
	return nil, err
}
