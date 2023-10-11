package main

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	loggerMiddleware "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mhd-sdk/go-chat/pkg/fiber/handlers"
	"github.com/mhd-sdk/go-chat/pkg/logging"
)

func main() {
	defer logging.Flush()
	app := fiber.New()
	app.Use(loggerMiddleware.New())
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/chat/:username", websocket.New(handlers.ChatHandler))

	log.Fatal(app.Listen(":3000"))
}
