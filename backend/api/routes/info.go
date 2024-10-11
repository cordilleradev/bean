package routes

import (
	"github.com/cordilleradev/bean/common"
	"github.com/cordilleradev/bean/common/types"
	"github.com/gin-gonic/gin"
)

type ExchangeInfo struct {
	SupportedMarginTypes []string               `json:"supported_margin_types"`
	LeaderboardPeriods   types.SupportedPeriods `json:"leaderboard_periods"`
	LeaderboardFields    []string               `json:"leaderboard_fields"`
}

type InfoResponse struct {
	ExchangeInfo map[string]ExchangeInfo `json:"exchange_info"`
}

func Info(clientMap *map[string]common.FuturesClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := InfoResponse{
			ExchangeInfo: make(map[string]ExchangeInfo),
		}
		for _, client := range *clientMap {
			marginTypes := client.GetSupportedMarginTypes()

			marginAsString := make([]string, len(marginTypes))
			for i, e := range marginTypes {
				marginAsString[i] = e.String()
			}

			leaderboardFields := client.GetSupportedLeaderboardFields()

			leaderboardAsString := make([]string, len(leaderboardFields))
			for i, f := range leaderboardFields {
				leaderboardAsString[i] = f.String()
			}

			response.ExchangeInfo[client.ExchangeName()] = ExchangeInfo{
				SupportedMarginTypes: marginAsString,
				LeaderboardPeriods:   client.GetLeaderboardPeriods(),
				LeaderboardFields:    leaderboardAsString,
			}
		}
		ctx.JSON(200, response)
	}
}
