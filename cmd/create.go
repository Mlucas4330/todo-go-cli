package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating a task:", args[0])
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
