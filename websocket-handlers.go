package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Erlaubt alle Verbindungen (ändere das für Sicherheit)
	},
}

// Handler für WebSocket-Verbindungen
func wsHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	for {

		text, timer := getMOTD()
		sendText := text+"!:timeSPLIT:!"+timer
		err := conn.WriteMessage(websocket.TextMessage, []byte(sendText))
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}

		time.Sleep(1 * time.Second)
	}
}
