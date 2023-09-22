package models

import "github.com/gofiber/contrib/websocket"

type Client struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Conn     *websocket.Conn
}

type Message struct {
	Author  string `json:"author"`
	Content string `json:"content"`
	Date    string `json:"date"`
}
