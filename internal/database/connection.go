package database

import (
	"database/sql"
	"fmt"

	"github.com/ricky/microservice_template/config"

	_ "github.com/lib/pq"
)

func NewConnection(dbConfig config.DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("database connected")

	return db, nil
}
