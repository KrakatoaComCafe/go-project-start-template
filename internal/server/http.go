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

type ServerParams struct {
	fx.In

	Lc        fx.Lifecycle
	Handler   http.Handler `name:"appHandler"`
	AppConfig *config.AppConfig
}

func StartServer(sp ServerParams) {
	addrs := fmt.Sprintf(":%s", sp.AppConfig.Port)
	srv := &http.Server{
		Addr:    addrs,
		Handler: sp.Handler,
	}

	sp.Lc.Append(fx.Hook{
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
