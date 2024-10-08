package gmx

import (
	"sort"
	"sync"
	"time"

	"github.com/cordilleradev/bean/common"
	gmx_abis "github.com/cordilleradev/bean/common/abigen/gmx"
	"github.com/cordilleradev/bean/common/types"
	"github.com/cordilleradev/bean/common/utils"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/machinebox/graphql"
)

type gmxClient struct {
	common.FuturesClient
	exchangeName                string
	gqlClient                   *graphql.Client
	connectionPool              *gmxConnectionPool
	gmxReaderContractAddress    string
	gmxDataStoreContractAddress string
	marketManager               *marketManager
	tokensUrl                   string
	streamCancelMap             map[string]chan struct{}
	mu                          sync.Mutex
}

func newGmxClient(
	exchangeName string,
	indexerUrl string,
	tokensUrl string,
	pricesUrl string,
	rpcs []string,
	gmxReaderContractAddress string,
	gmxDataStoreContractAddress string,
) (*gmxClient, error) {
	client := graphql.NewClient(indexerUrl)

	marketManager, err := newMarketManager(client, pricesUrl, tokensUrl)
	if err != nil {
		return nil, err
	}

	connectionPool, err := newGmxConnectionPool(rpcs, gmxReaderContractAddress)
	if err != nil {
		return nil, err
	}

	return &gmxClient{
		exchangeName:                exchangeName,
		gqlClient:                   client,
		connectionPool:              connectionPool,
		gmxReaderContractAddress:    gmxReaderContractAddress,
		gmxDataStoreContractAddress: gmxDataStoreContractAddress,
		tokensUrl:                   tokensUrl,
		marketManager:               marketManager,
		streamCancelMap:             make(map[string]chan struct{}),
	}, nil
}

func (g *gmxClient) ExchangeName() string {
	return g.exchangeName
}

func (g *gmxClient) GetSupportedMarginTypes() []types.MarginType {
	return []types.MarginType{types.Isolated}
}

func (g *gmxClient) GetLeaderboardPeriods() types.SupportedPeriods {
	return types.NewSupportedPeriods(
		[]string{"total"},
		&types.CustomPeriods{
			MinDays: 1,
			MaxDays: 90,
		},
	)
}

func (g *gmxClient) GetSupportedLeaderboardFields() []types.LeaderboardField {
	return []types.LeaderboardField{
		types.PeriodPnlPercent,
		types.PeriodPnlAbsolute,
		types.Volume,
		types.TotalTrades,
		types.Wins,
	}
}

func (g *gmxClient) GetLeaderboard(period string) ([]types.Trader, *types.APIError) {
	fromTimestamp, err := utils.ConvertToUTCTimestamp(period)
	if err != nil {
		return nil, types.InvalidPeriodError(period, err.Error())
	}
	if fromTimestamp > 0 && (time.Now().Unix()-fromTimestamp)/86400 > 90 {
		return nil, types.InvalidPeriodError(period, "must be between 1 - 90 days")
	}
	stats, err := fetchPeriodAccountStats(g.gqlClient, fromTimestamp)
	if err != nil {
		return nil, types.FailedLeaderboardCall(err)
	}
	traders := make([]types.Trader, len(stats))
	for i, u := range stats {
		relativePnl := 0.0
		absolutePnl := u.RealizedPnl - u.RealizedFees + u.RealizedPriceImpact
		if u.MaxCapital != 0 {
			relativePnl = absolutePnl / u.MaxCapital
		}

		traders[i] = types.Trader{
			UserId:            u.ID,
			PeriodPnlPercent:  relativePnl,
			PeriodPnlAbsolute: absolutePnl,
			Volume:            u.Volume,
			TotalTrades:       u.ClosedCount,
			Wins:              u.Wins,
		}
	}

	sort.SliceStable(traders, func(i, j int) bool {
		return traders[i].PeriodPnlAbsolute > traders[j].PeriodPnlAbsolute
	})

	return traders, nil
}

func (g *gmxClient) FetchPositions(userId string) ([]types.FuturesPosition, *types.APIError) {
	if !ethCommon.IsHexAddress(userId) {
		return nil, types.InvalidUserId(userId)
	}

	positionsRaw := g.connectionPool.GetPositions(
		ethCommon.HexToAddress(userId),
		g.gmxDataStoreContractAddress,
	)

	positionsList := make([]types.FuturesPosition, len(positionsRaw))
	for i, p := range positionsRaw {
		positionsList[i] = g.formatToFuturesPosition(p)
	}

	return positionsList, nil
}

func (g *gmxClient) StreamPositions(userId string, refreshWaitSeconds float64, positionStream chan types.FuturesResponse) *types.APIError {
	g.mu.Lock()
	if _, exists := g.streamCancelMap[userId]; exists {
		g.mu.Unlock()
		return types.StreamAlreadyStarted(userId)
	}

	cancelChan := make(chan struct{})
	g.streamCancelMap[userId] = cancelChan
	g.mu.Unlock()

	initialPositions, err := g.FetchPositions(userId)
	if err != nil {
		return err
	}

	positionStream <- types.FuturesResponse{
		Trader:             userId,
		Platform:           g.ExchangeName(),
		UpdatedPositions:   initialPositions,
		IsInitialPositions: true,
		DetectedChanges:    make([]types.FuturesDelta, 0),
	}

	ticker := time.NewTicker(
		time.Duration(refreshWaitSeconds) * time.Second,
	)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			positions, _ := g.FetchPositions(userId)

			deltas := utils.CalculatePositionDeltas(initialPositions, positions)
			if len(deltas) != 0 {
				positionStream <- types.FuturesResponse{
					Trader:             userId,
					Platform:           g.ExchangeName(),
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

func (g *gmxClient) CancelStream(userId string) *types.APIError {
	g.mu.Lock()
	defer g.mu.Unlock()

	cancelChan, exists := g.streamCancelMap[userId]
	if !exists {
		return types.NoSuchStream(userId)
	}

	close(cancelChan)
	delete(g.streamCancelMap, userId)
	return nil
}

func (g *gmxClient) formatToFuturesPosition(p gmx_abis.PositionProps) types.FuturesPosition {
	market, _ := g.marketManager.getMarket(p.Addresses.Market.Hex())
	collateralToken, _ := g.marketManager.getToken(p.Addresses.CollateralToken.String())
	indexToken, _ := g.marketManager.getToken(market.IndexToken)

	sizeInUsd := utils.BigIntToRelevantFloat(p.Numbers.SizeInUsd, 30, 5)
	sizeInTokens := utils.BigIntToRelevantFloat(p.Numbers.SizeInTokens, float64(indexToken.Decimals), 5)
	entryPrice := utils.RoundToNDecimalsOrSigFigs(sizeInUsd/sizeInTokens, 5)
	collateralAmount := utils.BigIntToRelevantFloat(p.Numbers.CollateralAmount, float64(collateralToken.Decimals), 5)
	collateralAmountUsd := utils.RoundToNDecimalsOrSigFigs(collateralAmount*collateralToken.MidPrice, 5)
	leverage := utils.RoundToNDecimalsOrSigFigs(sizeInUsd/collateralAmountUsd, 3)

	var pnl float64
	if p.Flags.IsLong {
		pnl = (indexToken.MidPrice - entryPrice) * sizeInTokens
	} else {
		pnl = (entryPrice - indexToken.MidPrice) * sizeInTokens
	}
	pnl = utils.RoundToNDecimalsOrSigFigs(pnl, 5)
	direction := utils.IsLongAsType(p.Flags.IsLong)

	return types.FuturesPosition{
		Market:         indexToken.Symbol + "-" + "USD",
		EntryPrice:     entryPrice,
		CurrentPrice:   indexToken.MidPrice,
		Status:         types.Open,
		Direction:      direction,
		MarginType:     types.Isolated,
		SizeUsd:        sizeInUsd,
		SizeToken:      sizeInTokens,
		UnrealizedPnl:  pnl,
		LeverageAmount: leverage,

		CollateralToken:          collateralToken.Symbol,
		CollateralTokenAmount:    collateralAmount,
		CollateralTokenAmountUsd: collateralAmountUsd,
	}
}
