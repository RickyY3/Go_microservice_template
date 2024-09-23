package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func addMiddlewares(mux *chi.Mux) {
	mux.Use(middleware.Logger)
	mux.Use(requestIdIntercepter)
}

func requestIdIntercepter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := r.Header.Get("requestId")
		if requestId == "" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "requestId", requestId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}