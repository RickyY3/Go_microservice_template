package internal_utils

import "github.com/ricky/microservice_template/internal/model"

func MapUser(request model.UpsertUserRequest) model.User{
	return model.User{
		Username: *request.Username,
		Pass: *request.Pass,
		Email: *request.Email,
	}
}
