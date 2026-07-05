// Package http wires the HTTP router for musig-stream.
package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/zaentrum/musig-stream/internal/config"
	"github.com/zaentrum/musig-stream/internal/play"
)

func New(cfg config.Config) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Get("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})
	r.Handle("/metrics", promhttp.Handler())

	r.Mount("/play", play.Routes(cfg))
	return r
}
