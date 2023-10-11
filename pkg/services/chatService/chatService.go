package chatservice

import (
	"fmt"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/mhd-sdk/go-chat/pkg/logging"
	"github.com/mhd-sdk/go-chat/pkg/models"
)

var mu sync.Mutex
var channel = []models.Message{}
var CurrentLoggedUsersChanged = make(chan bool)
var clientsMap sync.Map

func ConnectUser(username string, c *websocket.Conn) {
	logging.Logger.Info("New user connected to chat: " + username)
	clientsMap.Store(username, models.Client{Conn: c, Username: username})

	clientsMap.Range(func(key, value interface{}) bool {
		client := value.(models.Client)
		var wsUpdateMessage = models.ChatHandlerWsUpdate{Messages: channel, LoggedUsers: GetCurrentLoggedUsers()}
		err := client.Conn.WriteJSON(wsUpdateMessage)
		if err != nil {
			logging.Logger.Error("Error while sending message to client")
			return false
		}
		return true
	})

}

func DisconnectUser(username string) {
	user, _ := clientsMap.Load(username)
	clientsMap.Delete(username)
	logging.Logger.Info("User disconnected from chat: " + username)
	// broadcast user left
	clientsMap.Range(func(key, value interface{}) bool {
		client := value.(models.Client)
		var wsUpdateMessage = models.ChatHandlerWsUpdate{Messages: channel, LoggedUsers: GetCurrentLoggedUsers()}
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
	mu.Lock()
	channel = append(channel, message)
	logging.Logger.Info("New message added to chat: " + message.Content)
	mu.Unlock()
	clientsMap.Range(func(key, value interface{}) bool {
		client := value.(models.Client)
		wsUpdateMessage := models.ChatHandlerWsUpdate{Messages: channel, LoggedUsers: GetCurrentLoggedUsers()}
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
	clientsMap.Range(func(key, value interface{}) bool {
		client := value.(models.Client)
		users = append(users, client.Username)
		return true
	})
	return users
}
