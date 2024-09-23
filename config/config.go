package config

import (
	"fmt"
	"os"
	"strconv"
)

type AppConfig struct {
	DbConn DBConfig
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func GetDbConfig() DBConfig{
	return DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     getPortOrDefault(os.Getenv("DB_PORT")),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

func getPortOrDefault(port string) int {
	value, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("Cannot get port from .env, not an integer, using default 5432", err.Error())
		return 5432
	}
	return value
}