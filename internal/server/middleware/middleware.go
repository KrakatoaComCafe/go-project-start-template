package middleware

import (
	"net/http"

	"go.uber.org/fx"
)

type MiddlewareOut struct {
	fx.Out

	Handler http.Handler `name:"appHandler"`
}

func ApplyMidleware(mux *http.ServeMux) MiddlewareOut {
	return MiddlewareOut{
		Handler: chainMiddleWare(mux,
			Cors,
			Logging,
		),
	}
}

func chainMiddleWare(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}
