package api

import (
	"net/http"
	"time"

	"github.com/cordilleradev/bean/api/routes"
	"github.com/cordilleradev/bean/common"
	"github.com/cordilleradev/bean/common/types"
	"github.com/cordilleradev/bean/common/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ApiInstance struct {
	clientMap           *map[string]common.FuturesClient
	positionChan        chan types.FuturesResponse
	incomingMessageChan chan routes.WebsocketMessage
	connMap             *utils.ConnectionManager
}

func NewApiInstance(clientMap *map[string]common.FuturesClient) *ApiInstance {
	return &ApiInstance{
		clientMap:           clientMap,
		positionChan:        make(chan types.FuturesResponse),
		incomingMessageChan: make(chan routes.WebsocketMessage),
		connMap:             utils.NewConnectionManager(),
	}
}

func (api *ApiInstance) Run(isProd bool) {
	if isProd {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	// Add CORS middleware to allow all origins
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET(
		"/exchange_info",
		routes.Info(api.clientMap),
	)

	router.GET(
		"/live_leaderboard/:exchange",
		routes.LiveLeaderboard(api.clientMap),
	)

	router.GET(
		"/positions/:exchange/:userId",
		routes.Positions(api.clientMap),
	)

	router.GET(
		"/stream",
		routes.StartStreaming(upgrader, api.connMap, api.incomingMessageChan),
	)

	go api.StartStreamHandler()
	if isProd {
		router.Run(":10000")
	} else {
		router.Run(":8080")
	}

}
