package server

import (
	"log"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Options struct {
	Public      string
	SecretKey   string
	FiberConfig fiber.Config
}

func New(options *Options) *fiber.App {
	app := fiber.New(options.FiberConfig, fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})
	game := NewGame()

	app.Static("/", options.Public)

	app.Get("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			if c.Query("k") != options.SecretKey {
				return fiber.ErrUnauthorized
			}

			return c.Next()
		}

		return fiber.ErrUpgradeRequired
	}, HandleWS(game))

	return app
}

func CreateAndListen(addr string, options *Options) {
	server := New(options)

	log.Fatal(server.Listen(addr))
}
