package cmd

import (
	"log"

	"github.com/mlucas4330/todo-go-cli/internal/repository"
	"github.com/spf13/cobra"
)

var (
	id                  string
	title               string
	categoryStr         string
	description         string
	amountStr           string
	startDateStr        string
	endDateStr          string
	notificationDateStr string
)

var (
	taskRepo *repository.TaskRepository
)

func Init(repo *repository.TaskRepository) {
	taskRepo = repo

	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(viewCmd)
}

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "A simple CLI todo app",
	Long:  "A fast and minimal CLI app to manage your daily tasks.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("CLI failed:", err)
	}
}
