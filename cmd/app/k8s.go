package app

import (
	"context"
	"net/http"
	"sync/atomic"

	"github.com/AxulReich/kitchen/internal/pkg/logger"
	"github.com/gorilla/mux"
)

// Router register necessary routes and returns an instance of a router.
func Router(ctx context.Context) *mux.Router {
	isReady := &atomic.Value{}
	isReady.Store(false)

	go func() {
		logger.Error(ctx, "Readyz probe is negative by default...")
		isReady.Store(true)
		logger.Error(ctx, "Readyz probe is positive.")
	}()

	r := mux.NewRouter()

	r.HandleFunc("/healthz", healthz)
	r.HandleFunc("/readyz", readyz(isReady))
	return r
}

// healthz is a liveness probe.
func healthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// readyz is a readiness probe.
func readyz(isReady *atomic.Value) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		if isReady == nil || !isReady.Load().(bool) {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
