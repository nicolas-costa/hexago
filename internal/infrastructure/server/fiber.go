package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "hexago/docs"
	"hexago/internal/infrastructure/controllers"
)

type FiberServer struct {
	server *fiber.App
}

func NewFiberServer(
	healthController *controller.HealthController,
	coinController *controller.CoinController,
) FiberServer {
	server := fiber.New()

	server.Get("/health", healthController.Check)
	server.Get("/swagger/*", swagger.HandlerDefault)

	api := server.Group("/api")

	api.Get("/coin", coinController.CheckPrice)

	return FiberServer{server: server}
}

func (f *FiberServer) Start() error {
	// @todo: add to env vars
	return f.server.Listen(":3000")
}
