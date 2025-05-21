package main

import (
	"fmt"
	"log"
	"os"

	env "github.com/joho/godotenv"
	"github.com/mlucas4330/todo-go-cli/cmd"
	"github.com/mlucas4330/todo-go-cli/internal/db"
	"github.com/mlucas4330/todo-go-cli/internal/repository"
)

func main() {
	err := env.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", db_user, db_pass, db_port, db_name)

	database, err := db.New(connStr)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer database.Conn.Close()

	repo := repository.NewTaskRepository(database.Conn)

	cmd.Init(repo)
	cmd.Execute()
}
