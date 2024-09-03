package gmx

import (
	"time"

	"github.com/cordilleradev/bean/common"
	"github.com/cordilleradev/bean/common/types"
	"github.com/cordilleradev/bean/common/utils"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/machinebox/graphql"
)

type gmxClient struct {
	common.FuturesClient
	exchangeName                string
	gqlClient                   *graphql.Client
	marketNameMap               map[string]market
	tokenMap                    map[string]tokenInfo
	connectionPool              *gmxConnectionPool
	gmxReaderContractAddress    string
	gmxDataStoreContractAddress string
	priceCache                  *priceCache
	tokensUrl                   string
}

func newGmxClient(
	exchangeName string,
	indexerUrl string,
	tokensUrl string,
	pricesUrl string,
	waitPeriod float64,
	rpcs []string,
	gmxReaderContractAddress string,
	gmxDataStoreContractAddress string,

) (*gmxClient, error) {
	client := graphql.NewClient(indexerUrl)
	tokenMap, marketNameMap, err := getMarkets(client, tokensUrl)
	if err != nil {
		return nil, err
	}

	priceCache, err := newPriceCache(waitPeriod, pricesUrl)
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
		tokenMap:                    tokenMap,
		marketNameMap:               marketNameMap,
		connectionPool:              connectionPool,
		gmxReaderContractAddress:    gmxReaderContractAddress,
		gmxDataStoreContractAddress: gmxDataStoreContractAddress,
		tokensUrl:                   tokensUrl,
		priceCache:                  priceCache,
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
	return traders, nil
}

func (g *gmxClient) FetchPositions(userId string) (*types.FuturesResponse, *types.APIError) {

	if !ethCommon.IsHexAddress(userId) {
		return nil, types.InvalidUserId(userId)
	}

	positionsRaw := g.connectionPool.getPositions(
		ethCommon.HexToAddress(userId),
		g.gmxDataStoreContractAddress,
	)

	positionsList := make([]types.FuturesPosition, len(positionsRaw))

	for i, p := range positionsRaw {
		market := g.marketNameMap[p.Addresses.Market.Hex()]
		collateralTokenInfo := g.tokenMap[p.Addresses.CollateralToken.Hex()]

		marketTokenAddress := market.LongToken
		marketTokenInfo := g.tokenMap[marketTokenAddress]

		sizeInUsd := utils.BigIntToRelevantFloat(p.Numbers.SizeInUsd, 30, 5)

		sizeInTokens := utils.BigIntToRelevantFloat(p.Numbers.SizeInTokens, float64(marketTokenInfo.Decimals), 5)

		entryPrice := utils.RoundToNDecimals(sizeInUsd/sizeInTokens, 4)

		collateralAmount := utils.BigIntToRelevantFloat(p.Numbers.CollateralAmount, float64(collateralTokenInfo.Decimals), 5)

		collateralTokenPrice := g.priceCache.getPrice(p.Addresses.CollateralToken.Hex())
		marketTokenPrice := g.priceCache.getPrice(marketTokenAddress)

		leverage := utils.RoundToNDecimals(
			(marketTokenPrice*sizeInTokens)/(collateralTokenPrice*collateralAmount),
			5,
		)

		positionsList[i] = types.FuturesPosition{
			// Basic fields
			Market:       market.Name,
			Status:       types.Open,
			Direction:    utils.IsLongAsType(p.Flags.IsLong),
			MarginType:   types.Isolated,
			SizeUsd:      sizeInUsd,
			EntryPrice:   entryPrice,
			CurrentPrice: marketTokenPrice,

			// Isolated Margin fields
			CollateralToken:          collateralTokenInfo.Symbol,
			CollateralTokenAmount:    collateralAmount,
			CollateralTokenAmountUsd: utils.RoundToNDecimals(collateralAmount*collateralTokenPrice, 5),
			LeverageAmount:           leverage,
		}
	}

	return &types.FuturesResponse{
		Trader:    userId,
		Platform:  g.ExchangeName(),
		Positions: positionsList,
	}, nil
}

// func (g *gmxClient) UpdateMarketsAndTokens() error {
// 	tokens, markets, err := getMarkets(g.gqlClient, g.tokensUrl)

// 	if err != nil {
// 		return err
// 	}

// 	g.mutex.Lock()
// 	g.tokenMap = tokens
// 	g.marketNameMap = markets
// 	g.mutex.Unlock()

// 	return nil
// }
