package fxapp

import (
	"github.com/krakatoa/go-project-start-template/internal/config"
	"github.com/krakatoa/go-project-start-template/internal/handler"
	"github.com/krakatoa/go-project-start-template/internal/repository"
	"github.com/krakatoa/go-project-start-template/internal/server"
	"github.com/krakatoa/go-project-start-template/internal/server/middleware"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		server.NewMux,
	),
	fx.Provide(
		config.NewAppConfig,
		repository.NewMessageRepository,
		handler.NewHealthHandler,
		newMessageHandler,
		middleware.ApplyMidleware,
	),
	fx.Invoke(
		server.RegisterRoutes,
		server.StartServer,
	),
)

func newMessageHandler(repo *repository.MessageRepository) *handler.MessageHandler {
	return handler.NewMessageHandler(repo)
}
