package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
	"github.com/kr/pretty"
	"github.com/mhd-sdk/go-chat/models"
)

var chanel = []models.Message{}
var clientsMap sync.Map

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/chat", websocket.New(func(c *websocket.Conn) {
		username := c.Query("username")
		isNameUniq := checkNameIsUnique(username)
		if !isNameUniq {
			c.WriteMessage(websocket.TextMessage, []byte("Username already taken"))
			fmt.Println("Username " + username + " already taken")
			c.Close()
			return
		}
		id := uuid.New().String()
		client := &models.Client{Conn: c, Id: id, Username: username}
		clientsMap.Store(username, client)
		defer func() {
			fmt.Println(username + "just left the chat")
			clientsMap.Delete(username)
		}()

		// broadcast new user
		clientsMap.Range(func(key, value interface{}) bool {
			client := value.(*models.Client)
			err := client.Conn.WriteMessage(websocket.TextMessage, []byte(username+" just joined the chat"))
			fmt.Println(username + " just joined the chat")
			if err != nil {
				log.Println("write:", err)
				return false
			}
			return true
		})

		for {

			// message handler
			{
				// Read the message from the client
				_, msg, err := c.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					break
				}
				pretty.Println("New message received from : " + username + " | " + string(msg))
				var newMessage models.Message = models.Message{Author: username, Content: string(msg), Date: "now"}
				chanel = append(chanel, newMessage)

				// Send the chat message to all clients
				clientsMap.Range(func(key, value interface{}) bool {
					client := value.(*models.Client)
					err = client.Conn.WriteJSON(chanel)
					if err != nil {
						log.Println("write:", err)
						return false
					}
					return true
				})
			}
		}
	}))

	log.Fatal(app.Listen(":3000"))
}

func checkNameIsUnique(name string) bool {
	_, ok := clientsMap.Load(name)
	return !ok
}
