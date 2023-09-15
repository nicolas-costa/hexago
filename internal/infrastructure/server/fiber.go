package server

import (
	"github.com/gofiber/fiber/v2"
)

func NewFiberServer() *fiber.App {
	server := fiber.New()

	server.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"health": true,
		})
	})

	return server
}
