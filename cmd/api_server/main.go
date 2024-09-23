package main

import (
	//"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/ricky/microservice_template/config"
	"github.com/ricky/microservice_template/internal/database"
	"github.com/ricky/microservice_template/internal/handler"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
    if err != nil {
        fmt.Printf("Error loading .env file: %s", err)
    }
	db, err := database.NewConnection(config.GetDbConfig())
	if err != nil {
        fmt.Printf("Failed to connect to database: %s", err)
    }
	defer db.Close()

	//ctx := context.Background()
	if err := run(db); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(db *sql.DB) error {
	fmt.Println("Starting application")

	r := handler.NewServer(db)

	if err:= http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("error listening and serving: %s\n", err)
	}
	return nil
}