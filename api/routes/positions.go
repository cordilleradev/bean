package routes

import (
	"net/http"

	"github.com/cordilleradev/bean/common"
	"github.com/cordilleradev/bean/common/types"
	"github.com/gin-gonic/gin"
)

func Positions(clientMap *map[string]common.FuturesClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("userId")
		exchange := ctx.Param("exchange")

		if userId == "" || exchange == "" {
			ctx.JSON(http.StatusBadRequest,
				types.BlankParam([]string{"userId", "exchange"}),
			)
			return
		}

		client, exists := (*clientMap)[exchange]
		if !exists {
			ctx.JSON(http.StatusBadRequest, types.NoSuchExchange(exchange))
			return
		}

		positions, err := client.FetchPositions(userId)
		if err != nil {
			ctx.JSON(err.Code, err)
			return
		}

		ctx.JSON(200, positions)
	}
}
