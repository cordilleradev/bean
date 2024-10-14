package routes

import (
	"net/http"

	"github.com/cordilleradev/bean/common/types"
	"github.com/cordilleradev/bean/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebsocketMessage struct {
	Conn *websocket.Conn
	Data map[string][]string
}

func StartStreaming(
	upgrader websocket.Upgrader,
	manager *utils.ConnectionManager,
	messageChan chan WebsocketMessage,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				types.NewAPIError(
					http.StatusBadRequest,
					"websocket_upgrade_failed",
					err.Error(),
				),
			)
			return
		}
		manager.Connections.Store(conn, nil)

		go func() {
			for {
				_, exists := manager.Connections.Load(conn)
				if exists {
					var requestData map[string][]string
					err := conn.ReadJSON(&requestData)
					if err != nil {
						conn.WriteJSON(
							types.NewAPIError(
								http.StatusBadRequest,
								"invalid_stream_query",
								err.Error(),
							),
						)
						manager.Connections.Delete(conn)
						conn.Close()
						return
					}
					messageChan <- WebsocketMessage{
						Conn: conn,
						Data: requestData,
					}
				}
			}
		}()
	}
}
