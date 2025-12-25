package websocket

import (
	"log"
	"net/http"

	"github.com/aarthikrao/monorepo/common/routinepool"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Handler receives the connection wrapper and the message bytes. Handlers
// may call `Conn.WriteMessage` to send responses. Handlers are executed in
// worker goroutines provided by the given `routinepool.RoutinePool`.
// Note: the message slice is owned by the caller and must not be retained
// after the handler returns.
type Handler func(*Connection, []byte) error

// HandleWebSocket upgrades the HTTP connection to a WebSocket and processes incoming messages
// using the provided handler function within the given routine pool.
func HandleWebSocket(pool *routinepool.RoutinePool, handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("attempting websocket upgrade for", r.RemoteAddr, "path:", r.URL.Path)

		rawConn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("websocket upgrade error:", err, "path:", r.URL.Path, "remote:", r.RemoteAddr)
			return
		}
		// wrap the raw gorilla websocket connection with our safe Conn
		conn := newConn(rawConn)
		defer func() {
			err = conn.Close()
			log.Println("websocket connection closed for", r.RemoteAddr, err)
		}()

		for {
			_, message, err := rawConn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("websocket unexpected close error:", err, "remote:", r.RemoteAddr)
				} else {
					log.Println("websocket read error:", err, "remote:", r.RemoteAddr)
				}
				break
			}

			// execute handler with connection wrapper so it can respond
			pool.Submit(func() {
				if err := handler(conn, message); err != nil {
					log.Println("handler error:", err)
				}
			})
		}
	}
}
