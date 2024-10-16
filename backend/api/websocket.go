package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cordilleradev/bean/api/routes"
	"github.com/cordilleradev/bean/common/types"
	"github.com/cordilleradev/bean/common/utils"
)

func (api *ApiInstance) StartStreamHandler() {
	ticker := time.NewTicker(40 * time.Second)
	defer ticker.Stop()

	func() {
		for {
			select {
			case <-ticker.C:
				go api.HearbeatHandler()

			case positionUpdate := <-api.positionChan:
				go api.PositionUpdater(positionUpdate)

			case incomingMessage := <-api.incomingMessageChan:
				go api.RegisterNewStreams(incomingMessage)
			}
		}
	}()
}

func (api *ApiInstance) HearbeatHandler() {
	hitConnections := utils.NewConcurrentSet[*utils.ConcurrentConn]()
	api.connMap.Connections.Range(func(conn *utils.ConcurrentConn, _ *utils.ConcurrentSet[string]) bool {
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

func (api ApiInstance) isNewPositionMessageValid(incomingMessage routes.WebsocketMessage[map[string][]string]) bool {
	for exchange, userIds := range *incomingMessage.Data {
		client, exists := (*api.clientMap)[exchange]
		if !exists {
			incomingMessage.Conn.WriteJSON(
				types.NoSuchExchange(exchange),
			)
			return false
		}
		if len(userIds) == 0 {
			incomingMessage.Conn.WriteJSON(
				types.NewAPIError(
					http.StatusBadRequest,
					"no_user_ids_provided",
					fmt.Sprintf("list of user_ids for exchange '%s' was empty", exchange),
				),
			)
			return false
		}
		for _, userId := range userIds {
			if !client.ValidUserId(userId) {
				incomingMessage.Conn.WriteJSON(
					types.InvalidUserId(userId),
				)
				return false
			}
		}
	}
	return true
}

func (api *ApiInstance) RegisterNewStreams(incomingMessage routes.WebsocketMessage[map[string][]string]) {
	streamSet, e := api.connMap.Connections.Load(incomingMessage.Conn)
	fmt.Println(streamSet.String())
	if e {
		if api.isNewPositionMessageValid(incomingMessage) {
			for exchange, userIds := range *incomingMessage.Data {
				client, _ := (*api.clientMap)[exchange]
				for _, userId := range userIds {
					key := exchange + "-" + userId
					if !streamSet.Contains(key) {
						streamSet.Add(key)
						if client.StreamExists(userId) {
							positions, err := client.FetchPositions(userId)
							if err != nil {
								incomingMessage.Conn.WriteJSON(err)
							} else {
								incomingMessage.Conn.WriteJSON(
									types.FuturesResponse{
										Trader:             userId,
										Platform:           exchange,
										UpdatedPositions:   positions,
										DetectedChanges:    make([]types.FuturesDelta, 0),
										IsInitialPositions: true,
									},
								)
							}
						} else {
							go client.StreamPositions(userId, 5, api.positionChan)
						}
					} else {
						incomingMessage.Conn.WriteJSON(
							types.NewAPIError(
								http.StatusBadRequest,
								"already_streaming",
								fmt.Sprintf("stream already initialized for '%s'", key),
							),
						)
					}
				}
			}
		}

	}
}

func (api *ApiInstance) PositionUpdater(update types.FuturesResponse) {
	key := update.Platform + "-" + update.Trader
	api.connMap.Connections.Range(func(conn *utils.ConcurrentConn, subscribedUserIds *utils.ConcurrentSet[string]) bool {
		subscribedUserIds.Range(func(item string) bool {
			if item == key {
				go conn.WriteJSON(update)
			}
			return true
		})
		return true
	})
}
