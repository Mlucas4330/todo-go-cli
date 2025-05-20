package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "A simple CLI todo app",
	Long:  "A fast and minimal CLI app to manage your daily tasks.",
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		log.Fatal("Cobra cli failed:", err)
	}
}
