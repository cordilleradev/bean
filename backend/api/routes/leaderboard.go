package routes

import (
	"net/http"
	"strconv"

	"github.com/cordilleradev/bean/common"
	"github.com/cordilleradev/bean/common/types"
	"github.com/cordilleradev/bean/common/utils"
	"github.com/gin-gonic/gin"
)

func LiveLeaderboard(clientMap *map[string]common.FuturesClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		order := ctx.Query("order")
		limit := ctx.Query("limit")
		period := ctx.Query("period")
		sortBy := ctx.Query("sort_by")
		exchange := ctx.Param("exchange")

		if exchange == "" || period == "" {
			ctx.JSON(http.StatusBadRequest,
				types.BlankParam([]string{"exchange", "period"}),
			)
			return
		}

		client, exists := (*clientMap)[exchange]

		if !exists {
			ctx.JSON(
				http.StatusBadRequest,
				types.NoSuchExchange(exchange),
			)
			return
		}

		if order != "" && !utils.ContainsString([]string{"asc", "desc"}, order) {
			ctx.JSON(
				http.StatusBadRequest,
				types.NewAPIError(http.StatusBadRequest, "invalid_order_value", "order value must be 'asc' or 'desc'"),
			)
			return
		}

		limitInt := 100

		if limit == "all" {
			limitInt = -1
		} else if limit != "" {
			var err error
			limitInt, err = strconv.Atoi(limit)
			if err != nil {
				ctx.JSON(
					http.StatusBadRequest,
					types.NewAPIError(http.StatusBadRequest, "invalid_limit_value", "limit must be an integer"),
				)
				return
			}
		}
		isAsc := order == "asc"

		if sortBy == "" {
			sortBy = types.PeriodPnlAbsolute.String()
		}

		sortByField, fieldErr := utils.LeaderboardFieldSortAllowed(sortBy, client.GetSupportedLeaderboardFields())
		if fieldErr != nil {
			ctx.JSON(http.StatusBadRequest, fieldErr)
			return
		}

		leaderboard, err := client.GetLeaderboard(period, *sortByField, isAsc)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		leaderboardLen := len(leaderboard)

		if limitInt == -1 || leaderboardLen < limitInt {
			limitInt = leaderboardLen
		}

		ctx.JSON(http.StatusOK, leaderboard[0:limitInt])
	}
}
