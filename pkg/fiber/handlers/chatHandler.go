package handlers

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/kr/pretty"
	"github.com/mhd-sdk/go-chat/pkg/models"
	"github.com/mhd-sdk/go-chat/pkg/services/chatservice"
)

func ChatHandler(c *websocket.Conn) {
	username := c.Params("username")
	chatservice.ConnectUser(username, c)

	defer chatservice.DisconnectUser(username)

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		pretty.Println("New message received from : " + username + " | " + string(msg))
		newMessage := models.Message{Author: username, Content: string(msg), Date: "now"}
		chatservice.AddMessage(newMessage)

	}
}
