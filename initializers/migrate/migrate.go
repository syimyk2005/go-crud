package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		panic("DB_URL is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to DB: %v", err))
	}
	
	defer db.Close()

	if err := goose.Up(db, "migrations"); err != nil {
		panic(fmt.Sprintf("Failed to run migrations: %v", err))
	}

	fmt.Println("Migrations applied successfully!")
}
