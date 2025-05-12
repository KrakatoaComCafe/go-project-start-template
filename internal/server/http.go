package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/krakatoa/go-project-start-template/internal/config"
	"github.com/krakatoa/go-project-start-template/internal/handler"
	"go.uber.org/fx"
)

func NewMux() *http.ServeMux {
	return http.NewServeMux()
}

func RegisterRoutes(mux *http.ServeMux, healthHandler http.Handler, messageHandler *handler.MessageHandler) {
	mux.Handle("/health", healthHandler)
	mux.Handle("/message", messageHandler)
}

func StartServer(lc fx.Lifecycle, mux *http.ServeMux, appConfig *config.AppConfig) {
	addrs := fmt.Sprintf(":%s", appConfig.Port)
	srv := &http.Server{
		Addr:    addrs,
		Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("Starting server on port: ", addrs)
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("listen: %s", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down server...")
			return srv.Shutdown(ctx)
		},
	})
}
