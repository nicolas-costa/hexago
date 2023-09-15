//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"hexago/internal/infrastructure/controller"
	"hexago/internal/infrastructure/server"
)

func initialize() server.FiberServer {
	wire.Build(controller.NewHealthController, server.NewFiberServer)

	return server.FiberServer{}
}
