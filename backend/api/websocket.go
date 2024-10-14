package api

import (
	"github.com/cordilleradev/bean/common/utils"
	"github.com/gorilla/websocket"
)

func (api *ApiInstance) HearbeatHandler() {
	hitConnections := utils.NewConcurrentSet[*websocket.Conn]()
	api.connMap.Connections.Range(func(conn *websocket.Conn, _ *utils.ConcurrentSet[string]) bool {
		if !hitConnections.Contains(conn) {
			go func() {
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
			}()
		}
		return true
	})
}
