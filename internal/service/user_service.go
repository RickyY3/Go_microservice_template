package service

import (
	"context"

	"github.com/ricky/microservice_template/internal/database"
	"github.com/ricky/microservice_template/internal/model"
	internal_utils "github.com/ricky/microservice_template/internal/utils"
)

type UserService struct {
	Repository *database.UserRepository
}

func (s *UserService) CreateUser(ctx context.Context, upsertUserRequest model.UpsertUserRequest) error {
	user := internal_utils.MapUser(upsertUserRequest)
	err := s.Repository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdateUser(ctx context.Context, userId string, upsertUserRequest model.UpsertUserRequest) error {
	// convert upsertUserRequest to user
	return nil
}

func (s *UserService) GetUser(ctx context.Context, userId string) (model.User, error) {
	var user model.User
	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, userId string) error {
	return nil
}
