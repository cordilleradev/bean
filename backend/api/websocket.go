package api

import (
	"fmt"
	"time"

	"github.com/cordilleradev/bean/common/utils"
	"github.com/gorilla/websocket"
)

func (api *ApiInstance) HearbeatHandler() {
	hitConnections := utils.NewConcurrentSet[*websocket.Conn]()
	api.connMap.Connections.Range(func(conn *websocket.Conn, _ *utils.ConcurrentSet[string]) bool {
		if !hitConnections.Contains(conn) {
			if conn != nil {
				err := conn.WriteJSON(map[string]string{
					"method": "heartbeat",
				})
				if err != nil {
					conn.Close()
					api.connMap.Connections.Delete(conn)
				}
				hitConnections.Add(conn)
			}

		}
		return true
	})
}

func (api *ApiInstance) StartStreamHandler() {
	ticker := time.NewTicker(40 * time.Second)
	defer ticker.Stop()

	func() {
		for {
			select {
			case <-ticker.C:
				go api.HearbeatHandler()

			case positionUpdate := <-api.positionChan:
				go fmt.Printf(
					"found position response for %s\n",
					positionUpdate.Platform+"-"+positionUpdate.Trader,
				)

			case incomingMessage := <-api.incomingMessageChan:
				go fmt.Println(incomingMessage.Data)
			}
		}
	}()
}
