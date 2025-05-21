package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		err := taskRepo.Delete(id)

		if err != nil {
			fmt.Printf("Error deleting task: %v\n", err)
			return
		}

		fmt.Println("Task deleted successfully!")
	},
}

func init() {
	deleteCmd.Flags().StringVar(&id, "id", "", "ID of the task to delete")
	deleteCmd.MarkFlagRequired("id")
}
