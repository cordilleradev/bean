package kwenta

import (
	"math"
	"strconv"

	"github.com/cordilleradev/bean/common/types"
)

type kwentaClient struct {
	exchangeName         string
	leaderboardQueryName string
	leaderboardUrl       string
	marginType           types.MarginType
}

func newKwentaClient(
	exchangeName string,
	leaderboardUrl string,
	leaderboardQueryName string,
	marginType types.MarginType,
) (*kwentaClient, error) {
	return &kwentaClient{
		exchangeName:         exchangeName,
		marginType:           marginType,
		leaderboardUrl:       leaderboardUrl,
		leaderboardQueryName: leaderboardQueryName,
	}, nil
}

func (k *kwentaClient) ExchangeName() string {
	// Implementation here
	return k.exchangeName
}

func (k *kwentaClient) GetSupportedMarginTypes() []types.MarginType {
	// Implementation here
	return []types.MarginType{k.marginType}
}

func (k *kwentaClient) GetLeaderboardPeriods() types.SupportedPeriods {
	// Implementation here
	return types.NewSupportedPeriods([]string{"total"}, nil)
}

func (k *kwentaClient) GetSupportedLeaderboardFields() []types.LeaderboardField {
	return []types.LeaderboardField{
		types.PeriodPnlAbsolute,
		types.Volume,
		types.TotalTrades,
	}
}

func (k *kwentaClient) GetLeaderboard(period string) ([]types.Trader, *types.APIError) {
	if period != "total" {
		return nil, types.InvalidPeriodError(period, "allowed periods are total")
	}

	resp, err := fetchLeaderboardStats(k.leaderboardUrl, k.leaderboardQueryName)
	if err != nil {
		return nil, types.FailedLeaderboardCall(err)
	}

	traders := make([]types.Trader, len(resp.Top))
	for i, u := range resp.Top {

		pnlFloat, _ := strconv.ParseFloat(u.PnlWithFeesPaid, 64)
		pnl := pnlFloat / math.Pow(10, 18)

		volumeFloat, _ := strconv.ParseFloat(u.TotalVolume, 64)
		volume := volumeFloat / math.Pow(10, 18)

		totalTrades, _ := strconv.Atoi(u.TotalTrades)
		traders[i] = types.Trader{
			UserId:            u.AccountOwner,
			PeriodPnlAbsolute: pnl,
			Volume:            volume,
			TotalTrades:       totalTrades,
		}
	}

	return traders, nil
}

//implemented by each chain of client
// func (k *KwentaClient) FetchPositions(userId string) (*types.FuturesResponse, *types.APIError) {
// 	// Implementation here
// 	return nil, nil
// }
