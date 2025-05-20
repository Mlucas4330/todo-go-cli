package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view [task ID]",
	Short: "Views a task by its ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID := args[0]
		fmt.Printf("🔍 Viewing task with ID: %s\n", taskID)
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
