package database

import (
	"database/sql"
	"fmt"

	"github.com/ricky/microservice_template/internal/model"
)

type UserRepository struct {
	Db *sql.DB
}

func (ur *UserRepository) CreateUser(user model.User) error {
	query := "INSERT INTO users (user_name, pass, email) VALUES ($1, $2, $3)"
	_, err := ur.Db.Exec(query, user.Username, user.Pass, user.Email)
	if err != nil {
        return err
    }
	fmt.Println("Successfully created user.")
	return nil
}
