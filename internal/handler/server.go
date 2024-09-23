package handler

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewServer(db *sql.DB) http.Handler {
	mux := chi.NewRouter()

	addMiddlewares(mux)
	addRoutes(mux, db)

	return mux
}
