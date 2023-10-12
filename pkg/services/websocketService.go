package services

import (
	"fmt"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/mhd-sdk/go-chat/pkg/models"
)

var Mu sync.Mutex
var Channel = []models.Message{}
var CurrentLoggedUsersChanged = make(chan bool)
var ClientsMap sync.Map

func ConnectUser(username string, c *websocket.Conn) {
	fmt.Println("New user connected to chat: " + username)
	ClientsMap.Store(username, models.Client{Conn: c, Username: username})

	ClientsMap.Range(func(key, value interface{}) bool {
		client := value.(models.Client)
		var wsUpdateMessage = models.ChatHandlerWsUpdate{Messages: Channel, LoggedUsers: GetCurrentLoggedUsers(), PixelMatrix: PixelMatrix}
		err := client.Conn.WriteJSON(wsUpdateMessage)
		if err != nil {
			fmt.Println("Error while sending message to client")
			return false
		}
		return true
	})

}

func DisconnectUser(username string) {
	user, _ := ClientsMap.Load(username)
	ClientsMap.Delete(username)
	fmt.Println("User disconnected from chat: " + username)
	ClientsMap.Range(func(key, value interface{}) bool {
		client := value.(models.Client)
		var wsUpdateMessage = models.ChatHandlerWsUpdate{Messages: Channel, LoggedUsers: GetCurrentLoggedUsers(), PixelMatrix: PixelMatrix}
		err := client.Conn.WriteJSON(wsUpdateMessage)
		fmt.Println(username + " just left the chat")
		if err != nil {
			fmt.Println("Error while sending message to client")
			return false
		}
		return true
	})
	user.(models.Client).Conn.Close()
}

func GetCurrentLoggedUsers() []string {
	var users []string = []string{}
	ClientsMap.Range(func(key, value interface{}) bool {
		client := value.(models.Client)
		users = append(users, client.Username)
		return true
	})
	return users
}

func IsUsernameAvailable(username string) bool {
	_, ok := ClientsMap.Load(username)
	return !ok
}

func UpdateClients() {
	ClientsMap.Range(func(key, value interface{}) bool {
		client := value.(models.Client)
		var wsUpdateMessage = models.ChatHandlerWsUpdate{Messages: Channel, LoggedUsers: GetCurrentLoggedUsers(), PixelMatrix: PixelMatrix}
		err := client.Conn.WriteJSON(wsUpdateMessage)
		if err != nil {
			fmt.Println("Error while sending message to client")
			return false
		}
		return true
	})
}
