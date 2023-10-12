package websocketService

import (
	"fmt"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/mhd-sdk/go-chat/pkg/logging"
	"github.com/mhd-sdk/go-chat/pkg/models"
)

var Mu sync.Mutex
var Channel = []models.Message{}
var CurrentLoggedUsersChanged = make(chan bool)
var ClientsMap sync.Map

func ConnectUser(username string, c *websocket.Conn) {
	logging.Logger.Info("New user connected to chat: " + username)
	ClientsMap.Store(username, models.Client{Conn: c, Username: username})

	ClientsMap.Range(func(key, value interface{}) bool {
		client := value.(models.Client)
		var wsUpdateMessage = models.ChatHandlerWsUpdate{Messages: Channel, LoggedUsers: GetCurrentLoggedUsers()}
		err := client.Conn.WriteJSON(wsUpdateMessage)
		if err != nil {
			logging.Logger.Error("Error while sending message to client")
			return false
		}
		return true
	})

}

func DisconnectUser(username string) {
	user, _ := ClientsMap.Load(username)
	ClientsMap.Delete(username)
	logging.Logger.Info("User disconnected from chat: " + username)
	// broadcast user left
	ClientsMap.Range(func(key, value interface{}) bool {
		client := value.(models.Client)
		var wsUpdateMessage = models.ChatHandlerWsUpdate{Messages: Channel, LoggedUsers: GetCurrentLoggedUsers()}
		err := client.Conn.WriteJSON(wsUpdateMessage)
		fmt.Println(username + " just left the chat")
		if err != nil {
			logging.Logger.Error("Error while sending message to client")
			return false
		}
		return true
	})
	user.(*models.Client).Conn.Close()
}

func AddMessage(message models.Message) {
	Mu.Lock()
	Channel = append(Channel, message)
	logging.Logger.Info("New message added to chat: " + message.Content)
	Mu.Unlock()
	ClientsMap.Range(func(key, value interface{}) bool {
		client := value.(models.Client)
		wsUpdateMessage := models.ChatHandlerWsUpdate{Messages: Channel, LoggedUsers: GetCurrentLoggedUsers()}
		err := client.Conn.WriteJSON(wsUpdateMessage)
		if err != nil {
			logging.Logger.Error("Error while sending message to client")
			return false
		}
		return true
	})
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
