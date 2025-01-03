package gmx

import (
	"fmt"
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
	streamCancelMap             *utils.ConcurrentMap[string, chan struct{}]
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
		streamCancelMap:             utils.NewConcurrentMap[string, chan struct{}](),
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

func (hp *gmxClient) ValidUserId(userId string) bool {
	return ethCommon.IsHexAddress(userId)
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

func (g *gmxClient) GetLeaderboard(period string, sortBy types.LeaderboardField, orderIsAsc bool) ([]types.Trader, *types.APIError) {

	fromTimestamp, err := utils.ConvertToUTCTimestamp(period)
	fmt.Println(fromTimestamp)
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
		absolutePnl := (u.RealizedPnl +
			u.RealizedPriceImpact -
			u.RealizedFees +
			u.StartUnrealizedPnl +
			u.StartUnrealizedImpact -
			u.StartUnrealizedFees)
		if u.MaxCapital != 0 {
			relativePnl = absolutePnl / u.MaxCapital
		}

		relativePnl = utils.RoundToNDecimalsOrSigFigs(relativePnl, 5)
		absolutePnl = utils.RoundToNDecimalsOrSigFigs(absolutePnl, 5)
		u.Volume = utils.RoundToNDecimalsOrSigFigs(u.Volume, 5)

		traders[i] = types.Trader{
			UserId:            u.ID,
			PeriodPnlPercent:  &relativePnl,
			PeriodPnlAbsolute: &absolutePnl,
			Volume:            &u.Volume,
			TotalTrades:       &u.ClosedCount,
			Wins:              &u.Wins,
		}
	}
	utils.SortByFields(sortBy, traders, orderIsAsc)
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
	g.streamCancelMap.Store(userId, make(chan struct{}))
	cancelChan, _ := g.streamCancelMap.Load(userId)

	initialPositions, err := g.FetchPositions(userId)
	if err != nil {
		g.streamCancelMap.Delete(userId)
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
	cancelChan, exists := g.streamCancelMap.Load(userId)
	if !exists {
		return types.NoSuchStream(userId)
	}

	close(cancelChan)
	g.streamCancelMap.Delete(userId)
	return nil
}

func (g *gmxClient) StreamExists(userId string) bool {
	_, exists := g.streamCancelMap.Load(userId)
	return exists
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

	var pnl float64
	if p.Flags.IsLong {
		pnl = (indexToken.MidPrice - entryPrice) * sizeInTokens
	} else {
		pnl = (entryPrice - indexToken.MidPrice) * sizeInTokens
	}
	pnl = utils.RoundToNDecimalsOrSigFigs(pnl, 5)
	direction := utils.IsLongAsType(p.Flags.IsLong)
	leverage := utils.RoundToNDecimalsOrSigFigs(sizeInUsd/collateralAmountUsd, 3)

	return types.FuturesPosition{
		Market:               indexToken.Symbol + "-" + "USD",
		EntryPrice:           entryPrice,
		CurrentPrice:         indexToken.MidPrice,
		Status:               types.Open,
		Direction:            direction,
		MarginType:           types.Isolated,
		SizeUsd:              sizeInUsd,
		SizeToken:            sizeInTokens,
		UnrealizedPnlUsd:     pnl,
		UnrealizedPnlPercent: utils.RoundToNDecimalsOrSigFigs(pnl/collateralAmountUsd, 5),
		LeverageAmount:       leverage,

		CollateralToken:          collateralToken.Symbol,
		CollateralTokenAmount:    collateralAmount,
		CollateralTokenAmountUsd: collateralAmountUsd,
	}
}
