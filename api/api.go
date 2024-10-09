package api

import (
	"github.com/cordilleradev/bean/api/routes"
	"github.com/cordilleradev/bean/common"
	"github.com/cordilleradev/bean/common/types"
	"github.com/gin-gonic/gin"
)

type ApiInstance struct {
	clientMap    *map[string]common.FuturesClient
	positionChan chan types.FuturesResponse
}

func NewApiInstance(clientMap *map[string]common.FuturesClient) *ApiInstance {
	return &ApiInstance{
		clientMap:    clientMap,
		positionChan: make(chan types.FuturesResponse),
	}
}

func (api *ApiInstance) Run(isProd bool) {
	if isProd {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET(
		"/exchange_info",
		routes.Info(api.clientMap),
	)

	// Live leaderboard is stateless, but does all the sorting etc by itself,
	// in production, this is cached and refreshed by the worker, and then
	// fetched by the database.
	router.GET(
		"/live_leaderboard/:exchange",
		routes.LiveLeaderboard(api.clientMap),
	)

	router.GET(
		"/positions/:exchange/:userId",
		routes.Positions(api.clientMap),
	)

	router.Run(":8080")

}
