package hyperliquid

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cordilleradev/bean/common/types"
	"github.com/cordilleradev/bean/common/utils"
)

type HyperliquidClient struct {
}

func NewHyperliquidClient() (*HyperliquidClient, error) {
	client := &HyperliquidClient{}
	return client, nil
}

func (c *HyperliquidClient) ExchangeName() string {
	return "hyperliquid-v1-hyperliquid"
}

func (c *HyperliquidClient) GetSupportedMarginTypes() []types.MarginType {
	return []types.MarginType{types.Isolated, types.Cross}
}

func (c *HyperliquidClient) GetLeaderboardPeriods() types.SupportedPeriods {
	return types.NewSupportedPeriods([]string{"1d", "7d", "30d", "total"}, nil)
}

func (c *HyperliquidClient) GetSupportedLeaderboardFields() []types.LeaderboardField {
	return []types.LeaderboardField{
		types.PeriodPnlPercent,
		types.PeriodPnlAbsolute,
		types.Volume,
	}
}

func (c *HyperliquidClient) GetLeaderboard(period string) ([]types.Trader, *types.APIError) {
	allowedLeaderboardPeriods := c.GetLeaderboardPeriods()
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

	rawData, err := fetchPeriodStats()
	if err != nil {
		return nil, types.FailedLeaderboardCall(err)
	}
	traderList := make([]types.Trader, len(rawData.LeaderboardRows))

	referredKey := timePeriodMap[period]

	for i, u := range rawData.LeaderboardRows {
		for _, performance := range u.WindowPerformances {
			if performance[0] == referredKey {
				stats := performance[1].(map[string]interface{})
				roi, _ := strconv.ParseFloat(stats["roi"].(string), 64)
				pnl, _ := strconv.ParseFloat(stats["pnl"].(string), 64)
				vlm, _ := strconv.ParseFloat(stats["vlm"].(string), 64)
				traderList[i] = types.Trader{
					UserId:            u.EthAddress,
					PeriodPnlPercent:  roi,
					PeriodPnlAbsolute: pnl,
					Volume:            vlm,
				}
				break
			}
		}
	}
	return traderList, nil
}

func (c *HyperliquidClient) FetchPositions(userId string) (*types.FuturesResponse, *types.APIError) {

	return nil, nil
}
