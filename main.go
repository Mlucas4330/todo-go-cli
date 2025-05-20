package main

import (
	"github.com/mlucas4330/todo-go-cli/cmd"
	"github.com/mlucas4330/todo-go-cli/internal/db"
)

func main() {
	db.Connect()
	cmd.Execute()
}
