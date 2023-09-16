//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"hexago/internal/application"
	"hexago/internal/infrastructure/controller"
	"hexago/internal/infrastructure/repositories"
	"hexago/internal/infrastructure/server"
)

func initialize() server.FiberServer {
	wire.Build(
		controller.NewHealthController,
		controller.NewCoinController,
		server.NewFiberServer,
		application.NewHealthService,
		wire.Bind(new(application.Checker), new(*application.HealthService)),
		application.NewCoinService,
		wire.Bind(new(application.Pricer), new(*application.CoinService)),
		repositories.NewCoinRepository,
		wire.Bind(new(repositories.CoinRepository), new(*repositories.CoinGeckoClient)),
	)

	return server.FiberServer{}
}
