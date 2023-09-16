package server

import (
	"github.com/gofiber/fiber/v2"
	"hexago/internal/infrastructure/controller"
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

	api := server.Group("/api")

	api.Get("/coin", coinController.CheckPrice)

	return FiberServer{server: server}
}

func (f *FiberServer) Start() error {
	// @todo: add to env vars
	return f.server.Listen(":3000")
}
