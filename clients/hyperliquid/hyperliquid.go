package hyperliquid

import (
	"fmt"
	"math"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

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

	streamCancelMap map[string]chan struct{}
}

func NewHyperLiquidClient(cacheLeaderboard bool) (*HyperLiquidClient, error) {

	priceCache, priceCacheErr := newHyperLiquidPriceCache(
		hyperliquidApiUrl,
		hyperliquidWsUrl,
	)

	if priceCacheErr != nil {
		return nil, priceCacheErr
	}

	defaultClient := newHyperLiquidManager(1000, hyperliquidApiUrl, true)
	backupClient := newHyperLiquidManager(400, hyperLiquidUiApi, false)

	clientList := []*hyperLiquidManager{
		defaultClient,
		backupClient,
	}

	return &HyperLiquidClient{
		clients:          clientList,
		priceCache:       priceCache,
		cacheLeaderboard: cacheLeaderboard,
		leaderboardCache: make(map[string]traderPerformance),
		httpClient:       &http.Client{},
		streamCancelMap:  make(map[string]chan struct{}),
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

func (hp *HyperLiquidClient) cycle() []*hyperLiquidManager {

	bestAllowance := 0
	bestManager := hp.clients[0]

	for _, m := range hp.clients {
		allowance := m.allowancePerMinute - m.resetRequests()
		if (allowance) > bestAllowance {
			bestManager = m
			bestAllowance = allowance
		}
	}
	if bestManager.ApiUrl == hyperliquidApiUrl {
		return []*hyperLiquidManager{bestManager}
	}
	return []*hyperLiquidManager{bestManager, hp.clients[0]}
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
	sort.Slice(traders, func(i, j int) bool {
		return traders[i].PeriodPnlAbsolute > traders[j].PeriodPnlAbsolute
	})

	return traders, nil
}

func (hp *HyperLiquidClient) FetchPositions(userId string) ([]types.FuturesPosition, *types.APIError) {
	if !ethCommon.IsHexAddress(userId) {
		return nil, types.InvalidUserId(userId)
	}
	var err error
	manager := hp.cycle()
	for _, manager := range manager {
		manager.waitForToken(clearingHouseStateWeight)

		var resp clearingHouseState

		resp, err = makeInfoRequest[clearingHouseState](
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
			marginUsed := parseStringToFloat(position.Position.MarginUsed)
			leverageAmount := float64(position.Position.Leverage.Value)

			size := parseStringToFloat(position.Position.Szi)
			direction := types.Long
			if size < 0 {
				direction = types.Short
			}

			sizeUsd := entryPrice * size
			positionDetails := types.FuturesPosition{
				Market:         position.Position.Coin + "-USD",
				EntryPrice:     utils.RoundToNDecimalsOrSigFigs(entryPrice, 5),
				CurrentPrice:   hp.priceCache.getValue(position.Position.Coin),
				Status:         types.Open,
				Direction:      direction,
				SizeUsd:        utils.RoundToNDecimalsOrSigFigs(math.Abs(sizeUsd), 5),
				SizeToken:      utils.RoundToNDecimalsOrSigFigs(math.Abs(size), 5),
				UnrealizedPnl:  utils.RoundToNDecimalsOrSigFigs(unrealizedPnl, 5),
				LeverageAmount: leverageAmount,
			}

			switch position.Position.Leverage.Type {
			case "isolated":
				positionDetails.MarginType = types.Isolated
				positionDetails.CollateralTokenAmountUsd = marginUsed

			case "cross":
				positionDetails.MarginType = types.Cross
				positionDetails.HealthRatio = parseStringToFloat(position.Position.ReturnOnEquity)
				positionDetails.CrossMarginShare = marginUsed
				positionDetails.FreeCollateral = parseStringToFloat(resp.Withdrawable)
			}
			positions[index] = positionDetails
		}

		return positions, nil
	}

	return nil, types.FailedPositionsCall(err)
}

func (hp *HyperLiquidClient) StreamPositions(userId string, refreshWaitSeconds float64, positionStream chan types.FuturesResponse) *types.APIError {
	hp.mu.Lock()
	if _, exists := hp.streamCancelMap[userId]; exists {
		hp.mu.Unlock()
		return types.StreamAlreadyStarted(userId)
	}

	cancelChan := make(chan struct{})
	hp.streamCancelMap[userId] = cancelChan
	hp.mu.Unlock()

	initialPositions, err := hp.FetchPositions(userId)
	if err != nil {
		return err
	}

	positionStream <- types.FuturesResponse{
		Trader:             userId,
		Platform:           hp.ExchangeName(),
		UpdatedPositions:   initialPositions,
		IsInitialPositions: true,
		DetectedChanges:    make([]types.FuturesDelta, 0),
	}

	ticker := time.NewTicker(time.Duration(refreshWaitSeconds) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			positions, err := hp.FetchPositions(userId)
			if err != nil {
				continue
			}

			deltas := utils.CalculatePositionDeltas(initialPositions, positions)
			if len(deltas) != 0 {
				positionStream <- types.FuturesResponse{
					Trader:             userId,
					Platform:           hp.ExchangeName(),
					UpdatedPositions:   positions,
					IsInitialPositions: false,
					DetectedChanges:    deltas,
				}

				initialPositions = positions
			}
		case <-cancelChan:
			return nil
		}
	}
}

func (hp *HyperLiquidClient) CancelStream(userId string) *types.APIError {
	hp.mu.Lock()
	defer hp.mu.Unlock()

	cancelChan, exists := hp.streamCancelMap[userId]
	if !exists {
		return types.NoSuchStream(userId)
	}

	close(cancelChan)
	delete(hp.streamCancelMap, userId)
	return nil
}
