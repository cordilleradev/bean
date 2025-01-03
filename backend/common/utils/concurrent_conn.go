package utils

import (
	"sync"

	"github.com/gorilla/websocket"
)

type ConcurrentConn struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

func NewConcurrentConn(conn *websocket.Conn) *ConcurrentConn {
	return &ConcurrentConn{
		conn: conn,
	}
}

func (c *ConcurrentConn) WriteJSON(v interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.conn.WriteJSON(v)
}

func (c *ConcurrentConn) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.conn.Close()
}

func (c *ConcurrentConn) ReadMessage() (messageType int, p []byte, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.conn.ReadMessage()
}
