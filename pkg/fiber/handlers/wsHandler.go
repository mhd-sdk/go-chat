package handlers

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/mhd-sdk/go-chat/pkg/models"
	"github.com/mhd-sdk/go-chat/pkg/services"
)

type wsMessage struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

func WsHandler(c *websocket.Conn) {
	username := c.Params("username")
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
		switch wsMessage.Action {
		case "addMessage":
			newMessage := models.Message{Author: username, Content: wsMessage.Data.(string), Date: "now"}
			services.AddMessage(newMessage)
		case "changePixel":
			pixel := models.Pixel{}
			pixel.X = int(wsMessage.Data.(map[string]interface{})["x"].(float64))
			pixel.Y = int(wsMessage.Data.(map[string]interface{})["y"].(float64))
			pixel.Color = wsMessage.Data.(map[string]interface{})["color"].(string)
			if err != nil {
				log.Println("read:", err)
				break
			}
			services.ChangePixelColor(pixel.X, pixel.Y, pixel.Color)
		}
	}
}
