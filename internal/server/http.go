package server

import (
	"context"
	"log"
	"net/http"

	"go.uber.org/fx"
)

func NewMux() *http.ServeMux {
	return http.NewServeMux()
}

func RegisterRoutes(mux *http.ServeMux, healthHandler http.Handler) {
	mux.Handle("/health", healthHandler)
}

func StartServer(lc fx.Lifecycle, mux *http.ServeMux) {
	srv := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("Starting server on port :8000")
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
