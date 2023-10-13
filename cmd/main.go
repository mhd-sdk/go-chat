package main

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	loggerMiddleware "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mhd-sdk/go-chat/pkg/fiber/handlers"
	"github.com/mhd-sdk/go-chat/pkg/services"
)

func main() {
	services.InitPixelMatrix()
	app := fiber.New()
	app.Use(loggerMiddleware.New())
	app.Use(cors.New())
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/pixelwar/:username", websocket.New(handlers.WsHandler))

	app.Get("/nameisunique/:username", func(c *fiber.Ctx) error {
		username := c.Params("username")
		if services.IsUsernameAvailable(username) {
			return c.JSON("true")
		}
		return c.JSON("false")
	})

	log.Fatal(app.Listen("127.0.0.1:3001"))
}
