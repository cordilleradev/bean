package utils

import (
	"github.com/gorilla/websocket"
)

type ConnectionManager struct {
	Connections *ConcurrentMap[*websocket.Conn, *ConcurrentSet[string]]
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		Connections: NewConcurrentMap[*websocket.Conn, *ConcurrentSet[string]](),
	}
}
