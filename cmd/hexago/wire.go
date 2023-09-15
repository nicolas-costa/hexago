//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"hexago/internal/application"
	"hexago/internal/infrastructure/controller"
	"hexago/internal/infrastructure/server"
)

func initialize() server.FiberServer {
	wire.Build(
		controller.NewHealthController,
		server.NewFiberServer,
		application.NewHealthService,
		wire.Bind(new(application.Checker), new(*application.HealthService)),
	)

	return server.FiberServer{}
}
