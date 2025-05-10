package fxapp

import (
	"github.com/krakatoa/go-project-start-template/internal/handler"
	"github.com/krakatoa/go-project-start-template/internal/server"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		server.NewMux,
	),
	fx.Provide(
		handler.NewHealthHandler,
	),
	fx.Invoke(
		server.RegisterRoutes,
		server.StartServer,
	),
)
