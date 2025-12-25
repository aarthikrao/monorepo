package main

import (
	"github.com/aarthikrao/monorepo/common/routinepool"
	"github.com/aarthikrao/monorepo/common/websocket"
	"github.com/gin-gonic/gin"
	gw "github.com/gorilla/websocket"
)

func main() {
	pool := routinepool.New(5)
	websocketHandler := websocket.HandleWebSocket(pool, func(c *websocket.Connection, message []byte) error {
		// Process the incoming message and echo it back
		println("Received message:", string(message))
		// echo back as a text message
		return c.WriteMessage(gw.TextMessage, message)
	})

	router := gin.Default()
	router.GET("/ws", gin.WrapF(websocketHandler))
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	router.Run(":8080")

	pool.CloseAndWait()
}
