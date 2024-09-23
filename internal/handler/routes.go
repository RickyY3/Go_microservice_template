package handler

import (
	"database/sql"

	"github.com/ricky/microservice_template/internal/database"
	"github.com/ricky/microservice_template/internal/service"

	"github.com/go-chi/chi/v5"
)



func addRoutes(mux *chi.Mux, db *sql.DB) {
	userHandler := &UserHandler{
		userService: &service.UserService{
			Repository: &database.UserRepository{
				Db: db,
			},
		},
	}

	mux.Post("/create-user", userHandler.createUser)
	mux.Put("/{userId}/update", userHandler.updateUser)
	mux.Get("/{userId}/get", userHandler.getUser)
	mux.Delete("/{userId}/delete", userHandler.deleteUser)
}


