package server

import (
	"github.com/gofiber/fiber/v2"
	"hexago/internal/infrastructure/controller"
)

type FiberServer struct {
	server *fiber.App
}

func NewFiberServer(healthController *controller.HealthController) FiberServer {
	server := fiber.New()

	server.Get("/health", healthController.Check)

	return FiberServer{server: server}
}

func (f *FiberServer) Start() error {
	// @todo: add to env vars
	return f.server.Listen(":3000")
}
