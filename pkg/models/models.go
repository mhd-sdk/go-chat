package models

import "github.com/gofiber/contrib/websocket"

type Client struct {
	Username string `json:"username"`
	Conn     *websocket.Conn
}

type Message struct {
	Author  string `json:"author"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

type UserList []string

type ChatHandlerWsUpdate struct {
	Messages    []Message `json:"messages"`
	LoggedUsers UserList  `json:"loggedUsers"`
}
