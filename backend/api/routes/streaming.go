package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cordilleradev/bean/common/types"
	"github.com/cordilleradev/bean/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebsocketMessage[T any] struct {
	Conn   *websocket.Conn
	Method string `json:"method"`
	Data   *T     `json:"query_data"`
}

func StartStreaming(
	upgrader websocket.Upgrader,
	manager *utils.ConnectionManager,
	messageChan chan WebsocketMessage[map[string][]string],
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
					var message WebsocketMessage[map[string][]string]
					_, msg, err := conn.ReadMessage()
					if err != nil {
						conn.WriteJSON(
							types.NewAPIError(
								http.StatusBadRequest,
								"failed_websocket_conn",
								err.Error(),
							),
						)
						break
					}

					err = json.Unmarshal(msg, &message)
					if err != nil {
						conn.WriteJSON(
							types.NewAPIError(
								http.StatusBadRequest,
								"invalid_message_format",
								"the message was formatted improperly",
							),
						)
						continue
					}

					if message.Method == "" || message.Data == nil {
						conn.WriteJSON(
							types.BlankParam([]string{"method", "query_data"}),
						)
						continue
					}

					if !utils.ContainsString([]string{"positions_query"}, message.Method) {
						conn.WriteJSON(
							types.NewAPIError(
								http.StatusBadRequest,
								"invalid_method",
								fmt.Sprintf("'%s' is an invalid method", message.Method),
							),
						)
						continue
					}
					messageChan <- message
				} else {
					return
				}
			}
			manager.Connections.Delete(conn)
			conn.Close()
		}()
	}
}
