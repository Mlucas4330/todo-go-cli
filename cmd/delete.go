package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deleting a task:", args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
