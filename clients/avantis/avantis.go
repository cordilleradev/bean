package avantis

import (
	avantis_multicall_abis "github.com/cordilleradev/bean/common/abigen/avantis/multiCall"
	"github.com/cordilleradev/bean/common/types"
	"github.com/cordilleradev/bean/common/utils"
	ethCommon "github.com/ethereum/go-ethereum/common"
)

type AvantisClient struct {
	connectionPool *avantisConnectionPool
}

func NewAvantisClient(rpcs []string) (*AvantisClient, error) {
	connectionPool, err := newAvantisConnectionPool(rpcs, avantisMultiCallContractAddress)
	if err != nil {
		return nil, err
	}

	return &AvantisClient{connectionPool: connectionPool}, nil
}

func (a *AvantisClient) ExchangeName() string {
	// Implementation here
	return "avantis-v1-base"
}

func (a *AvantisClient) GetSupportedMarginTypes() []types.MarginType {
	return []types.MarginType{types.Isolated}
}

func (a *AvantisClient) GetLeaderboardPeriods() types.SupportedPeriods {
	return types.NewSupportedPeriods([]string{"total"}, nil)
}

func (a *AvantisClient) GetSupportedLeaderboardFields() []types.LeaderboardField {
	return []types.LeaderboardField{
		types.PeriodPnlPercent,
		types.PeriodPnlAbsolute,
		types.Volume,
		types.TotalTrades,
		types.Wins,
	}
}

func (a *AvantisClient) GetLeaderboard(period string) ([]types.Trader, *types.APIError) {
	resp, err := fetchLeaderBoard()
	if err != nil {
		return nil, types.FailedLeaderboardCall(err)
	}

	traders := make([]types.Trader, len(resp.LeaderBoard))

	for i, u := range resp.LeaderBoard {
		traders[i] = types.Trader{
			UserId:            u.Trader,
			PeriodPnlPercent:  u.TotalProfitPercentage / 100,
			PeriodPnlAbsolute: u.TotalProfits,
			Volume:            u.TotalPositionSizes,
			TotalTrades:       u.TotalTrades,
			Wins:              u.TotalWins,
		}
	}

	return traders, nil
}

func (a *AvantisClient) FetchPositions(userId string) (*types.FuturesResponse, *types.APIError) {
	if !ethCommon.IsHexAddress(userId) {
		return nil, types.InvalidUserId(userId)
	}

	rawPositions := a.connectionPool.getPositions(ethCommon.HexToAddress(userId))

	positionsList := make([]types.FuturesPosition, 0)
	for _, p := range rawPositions {
		if p.Trade.Trader.Hex() != "0x0000000000000000000000000000000000000000" {
			positionsList = append(positionsList, formatFuturesPositions(p))
		}
	}

	return &types.FuturesResponse{
		Trader:    userId,
		Platform:  a.ExchangeName(),
		Positions: positionsList,
	}, nil
}

func formatFuturesPositions(p avantis_multicall_abis.IMulticallAggregatedTrade) types.FuturesPosition {

	entryPrice := utils.BigIntToRelevantFloat(p.Trade.OpenPrice, 10, 5)
	leverage := utils.BigIntToRelevantFloat(p.Trade.Leverage, 10, 5)

	positionSize := utils.BigIntToRelevantFloat(p.TradeInfo.OpenInterestUSDC, 6, 5)

	return types.FuturesPosition{
		// Basic fields
		Market:       "",
		Status:       types.Open,
		Direction:    utils.IsLongAsType(p.Trade.Buy),
		MarginType:   types.Isolated,
		SizeUsd:      positionSize, //not sure this is how it works
		EntryPrice:   entryPrice,
		CurrentPrice: 0,

		// Isolated Margin fields
		CollateralToken:          "USDC",
		CollateralTokenAmount:    0,
		CollateralTokenAmountUsd: 0,
		LeverageAmount:           leverage,
	}
}
