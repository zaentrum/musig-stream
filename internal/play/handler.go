// Package play exposes the placeholder playback surface for musig-stream.
//
// audio Opus HLS, gapless playback
package play

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/zaentrum/musig-stream/internal/config"
)

func Routes(_ config.Config) http.Handler {
	r := chi.NewRouter()
	r.Get("/*", placeholder)
	return r
}

func placeholder(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusNotImplemented)
	_, _ = w.Write([]byte("musig-stream playback not yet implemented; scaffold only\n"))
}
