package websocket

import (
	"sync"

	gw "github.com/gorilla/websocket"
)

// Connection is a thin wrapper around *gorilla/websocket.Connection that provides a
// write mutex so handlers can safely write concurrently from different
// goroutines. It intentionally exposes only the minimal methods needed by
// handlers.
type Connection struct {
	raw *gw.Conn
	mu  sync.Mutex
}

func newConn(c *gw.Conn) *Connection {
	return &Connection{raw: c}
}

// WriteMessage sends a complete message. It is safe for concurrent use by
// different goroutines; a mutex serializes writes.
func (c *Connection) WriteMessage(messageType int, data []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.raw.WriteMessage(messageType, data)
}

// Close closes the underlying connection.
func (c *Connection) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.raw.Close()
}
