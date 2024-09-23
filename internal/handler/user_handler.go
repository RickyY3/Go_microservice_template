package handler

import (
	"fmt"
	"net/http"

	"github.com/ricky/microservice_template/internal/model"
	"github.com/ricky/microservice_template/internal/service"
	"github.com/ricky/microservice_template/pkg/utils"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	userService *service.UserService
}

func (uh *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var upsertUserRequest model.UpsertUserRequest
	err := utils.Decode(&upsertUserRequest, r)
	if err != nil {
		errorMessage := fmt.Sprintf("Invalid JSON: %s", err.Error())
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}

	err = uh.userService.CreateUser(ctx, upsertUserRequest)
	fmt.Println(err)
}

func (uh *UserHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := chi.URLParam(r, "userId")

	var upsertUserRequest model.UpsertUserRequest
	err := utils.Decode(&upsertUserRequest, r)
	if err != nil {
		errorMessage := fmt.Sprintf("Invalid JSON: %s", err.Error())
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}

	err = uh.userService.UpdateUser(ctx, userId, upsertUserRequest)
}

func (uh *UserHandler) getUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := chi.URLParam(r, "userId")

	user, err := uh.userService.GetUser(ctx, userId)

	print(user.Username)
	print(err)
}

func (uh *UserHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := chi.URLParam(r, "userId")

	err := uh.userService.DeleteUser(ctx, userId)

	print(err)
}