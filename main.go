package main

import (
	"log"

	"github.com/mlucas4330/todo-go-cli/cmd"
	"github.com/mlucas4330/todo-go-cli/internal/db"
	"github.com/mlucas4330/todo-go-cli/internal/repository"
)

func main() {
	database, err := db.New("connStr")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer database.Conn.Close()

	repo := repository.NewTaskRepository(database.Conn)

	cmd.Init(repo)
	cmd.Execute()
}
