package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/mhd-sdk/go-chat/pkg/models"
	"github.com/mhd-sdk/go-chat/pkg/services"
)

type wsMessage struct {
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
}

func WsHandler(c *websocket.Conn) {
	username := c.Params("username")
	isNameUnqique := services.IsUsernameAvailable(username)
	if !isNameUnqique {
		c.WriteMessage(websocket.TextMessage, []byte("Username is not unique"))
		c.Close()
		return
	}
	services.ConnectUser(username, c)

	defer services.DisconnectUser(username)

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		// parse to object
		wsMessage := wsMessage{}
		err = json.Unmarshal(msg, &wsMessage)

		if err != nil {
			log.Println("read:", err)
			break
		}
		fmt.Println(wsMessage)
		switch wsMessage.Action {
		case "addMessage":
			newMessage := models.Message{Author: username, Content: wsMessage.Payload.(string), Date: "now"}
			services.AddMessage(newMessage)
		case "changePixel":
			pixel := models.Pixel{}
			pixel.X = int(wsMessage.Payload.(map[string]interface{})["x"].(float64))
			pixel.Y = int(wsMessage.Payload.(map[string]interface{})["y"].(float64))
			pixel.Color = wsMessage.Payload.(map[string]interface{})["color"].(string)
			if err != nil {
				log.Println("read:", err)
				break
			}
			services.ChangePixelColor(pixel.X, pixel.Y, pixel.Color)
		}
	}
}
