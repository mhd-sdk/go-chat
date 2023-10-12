package services

import (
	"fmt"

	"github.com/mhd-sdk/go-chat/pkg/models"
)

func AddMessage(message models.Message) {
	Mu.Lock()
	Channel = append(Channel, message)
	fmt.Println("New message added to chat: " + message.Content)
	Mu.Unlock()
	ClientsMap.Range(func(key, value interface{}) bool {
		client := value.(models.Client)
		wsUpdateMessage := models.ChatHandlerWsUpdate{Messages: Channel, LoggedUsers: GetCurrentLoggedUsers(), PixelMatrix: PixelMatrix}
		err := client.Conn.WriteJSON(wsUpdateMessage)
		if err != nil {
			fmt.Println("Error while sending message to client")
			return false
		}
		return true
	})
}

func GetMessages() []models.Message {
	return Channel
}
