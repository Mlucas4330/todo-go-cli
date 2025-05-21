package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View a task or all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("id") {
			if id == "" {
				fmt.Println("Error: --id flag requires a value.")
				return
			}

			task, err := taskRepo.FindById(id)

			if err != nil {
				if err.Error() == "sql: no rows in result set" {
					fmt.Printf("Task with ID '%s' not found.\n", id)
				} else {
					log.Printf("Error fetching task by ID '%s': %v\n", id, err)
					fmt.Println("An error occurred while fetching the task.")
				}
				return
			}

			fmt.Println("--- Task Details ---")
			fmt.Printf("ID: %v\n", task.ID)
			fmt.Printf("	Title: %s\n", task.Title)
			fmt.Printf("	Description: %s\n", task.Description)
			fmt.Printf("	Amount: %v\n", task.Amount)
			fmt.Printf("	Category: %s\n", task.Category)
			fmt.Printf("	Start Date: %v\n", task.StartDate)
			fmt.Printf("	End Date: %v\n", task.EndDate)
			fmt.Printf("	Notification Date: %v\n", task.NotificationDate)
			fmt.Println("--------------------")
		} else {
			tasks, err := taskRepo.FindAll()

			if err != nil {
				log.Printf("Error fetching all tasks: %v\n", err)
				fmt.Println("An error occurred while fetching tasks.")
				return
			}

			if len(tasks) == 0 {
				fmt.Println("No tasks found.")
				return
			}

			fmt.Println("--- All Tasks ---")
			for _, task := range tasks {
				fmt.Printf("ID: %v\n", task.ID)
				fmt.Printf("  Title: %s\n", task.Title)
				fmt.Printf("  Category: %s\n", task.Category)
				fmt.Printf("	Description: %s\n", task.Description)
				fmt.Printf("	Amount: %v\n", task.Amount)
				fmt.Printf("	Category: %s\n", task.Category)
				fmt.Printf("	Start Date: %v\n", task.StartDate)
				fmt.Printf("	End Date: %v\n", task.EndDate)
				fmt.Printf("	Notification Date: %v\n", task.NotificationDate)
				fmt.Println("-----------------")
			}
		}
	},
}

func init() {
	viewCmd.Flags().StringVar(&id, "id", "", "ID of the task to view")

	rootCmd.AddCommand(viewCmd)
}
