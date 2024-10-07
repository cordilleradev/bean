package hyperliquid

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/cordilleradev/bean/common/types"
	"github.com/cordilleradev/bean/common/utils"
	ethCommon "github.com/ethereum/go-ethereum/common"
)

type HyperLiquidClient struct {
	clients    []*hyperLiquidManager
	priceCache *hyperLiquidPriceCache
	mu         sync.Mutex

	cacheLeaderboard bool
	leaderboardCache map[string]traderPerformance
	periodSession    []string
	roundRobinIndex  int
	httpClient       *http.Client
}

func NewHyperLiquidClient(cacheLeaderboard bool) (*HyperLiquidClient, error) {

	priceCache, priceCacheErr := newHyperLiquidPriceCache(
		hyperliquidApiUrl,
		hyperliquidWsUrl,
	)

	if priceCacheErr != nil {
		return nil, priceCacheErr
	}

	uiApiClient := newHyperLiquidManager(1200, hyperLiquidUiApi, false)
	defaultClient := newHyperLiquidManager(1200, hyperliquidApiUrl, true)

	clientList := []*hyperLiquidManager{
		defaultClient,
		uiApiClient,
	}

	return &HyperLiquidClient{
		clients:          clientList,
		priceCache:       priceCache,
		cacheLeaderboard: cacheLeaderboard,
		leaderboardCache: make(map[string]traderPerformance),
		httpClient:       &http.Client{},
	}, nil
}

func (hp *HyperLiquidClient) shouldRefreshCache(period string) bool {
	hp.mu.Lock()
	defer hp.mu.Unlock()

	needsRefresh := len(hp.leaderboardCache) == 0 ||
		len(hp.periodSession) == 4 ||
		utils.ContainsString(hp.periodSession, period)
	if needsRefresh {
		hp.periodSession = nil
	} else {
		hp.periodSession = append(hp.periodSession, period)
	}
	return needsRefresh
}

func (hp *HyperLiquidClient) cycle() *hyperLiquidManager {

	bestAllowance := 0
	bestManager := hp.clients[0]

	for _, m := range hp.clients[1:] {
		if (m.allowancePerMinute - m.resetRequests()) > bestAllowance {
			bestManager = m
		}
	}
	return bestManager
}

func (hp *HyperLiquidClient) ExchangeName() string {
	return "hyperliquid-v1-hyperliquid"
}

func (hp *HyperLiquidClient) GetSupportedMarginTypes() []types.MarginType {
	return []types.MarginType{types.Isolated, types.Cross}
}

func (hp *HyperLiquidClient) GetLeaderboardPeriods() types.SupportedPeriods {
	return types.NewSupportedPeriods([]string{"1d", "7d", "30d", "total"}, nil)
}

func (hp *HyperLiquidClient) GetSupportedLeaderboardFields() []types.LeaderboardField {
	return []types.LeaderboardField{
		types.PeriodPnlPercent,
		types.PeriodPnlAbsolute,
		types.Volume,
	}
}

func (hp *HyperLiquidClient) GetLeaderboard(period string) ([]types.Trader, *types.APIError) {
	allowedLeaderboardPeriods := hp.GetLeaderboardPeriods()
	if !utils.ContainsString(allowedLeaderboardPeriods.FixedPeriods, period) {
		return nil, types.InvalidPeriodError(
			period,
			fmt.Sprintf(
				"allowed periods are %s",
				strings.Join(
					allowedLeaderboardPeriods.FixedPeriods,
					", "),
			),
		)
	}

	var response map[string]traderPerformance
	var err error

	if hp.cacheLeaderboard && !hp.shouldRefreshCache(period) {
		hp.mu.Lock()
		response = hp.leaderboardCache
		hp.mu.Unlock()
	} else {
		response, err = leaderboardCall(hp.httpClient)
		if err != nil {
			return nil, types.FailedLeaderboardCall(err)
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

func (hp *HyperLiquidClient) FetchPositions(userId string) ([]types.FuturesPosition, *types.APIError) {
	if !ethCommon.IsHexAddress(userId) {
		return nil, types.InvalidUserId(userId)
	}

	var err error
	for i := 0; i < len(hp.clients); i++ {
		manager := hp.cycle()
		manager.waitForToken(clearingHouseStateWeight)

		resp, err := makeInfoRequest[clearingHouseState](
			manager.HttpClient,
			manager.ApiUrl,
			map[string]string{
				"type": "clearinghouseState",
				"user": userId,
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
				CurrentPrice:   hp.priceCache.getValue(position.Position.Coin), // Update with current price from relevant source
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
	return nil, types.FailedPositionsCall(err)
}
