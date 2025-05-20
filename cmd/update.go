package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updated a task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Updating a task:", args[0])
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
