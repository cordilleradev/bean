package gmx

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cordilleradev/bean/common"
	"github.com/cordilleradev/bean/common/types"
	"github.com/cordilleradev/bean/common/utils"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/machinebox/graphql"
)

type gmxClient struct {
	common.FuturesClient
	exchangeName  string
	gqlClient     *graphql.Client
	marketNameMap map[string]string
	tokenMap      map[string]tokenInfo
	// mutex         *sync.Mutex
	tokensUrl string
}

func newGmxClient(
	exchangeName string,
	indexerUrl string,
	tokensUrl string,
) (*gmxClient, error) {
	client := graphql.NewClient(indexerUrl)

	tokenMap, marketNameMap, err := getMarkets(client, tokensUrl)

	if err != nil {
		return nil, err
	}

	tokenMapJson, err := json.MarshalIndent(tokenMap, "", "  ")
	if err != nil {
		return nil, err
	}
	marketNameMapJson, err := json.MarshalIndent(marketNameMap, "", "  ")
	if err != nil {
		return nil, err
	}
	fmt.Println("Token Map:", string(tokenMapJson))
	fmt.Println("Market Name Map:", string(marketNameMapJson))

	return &gmxClient{
		exchangeName:  exchangeName,
		gqlClient:     client,
		tokenMap:      tokenMap,
		marketNameMap: marketNameMap,
		// mutex:         &sync.Mutex{},
		tokensUrl: tokensUrl,
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
		}
	}

	return traders, nil
}

func (g *gmxClient) FetchPositions(userId string) ([]types.FuturesResponse, *types.APIError) {

	if !ethCommon.IsHexAddress(userId) {
		return nil, types.InvalidUserId(userId)
	}

	return nil, nil
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
