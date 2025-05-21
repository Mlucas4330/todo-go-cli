package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/mergestat/timediff"
	"github.com/mlucas4330/todo-go-cli/internal/util"
	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a task or all tasks",
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
				}
				return
			}

			w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, '\t', 0)

			fmt.Fprintln(w, "--- Task Details ---")
			fmt.Fprintf(w, "ID:\t%v\n", task.ID)
			fmt.Fprintf(w, "Title:\t%s\n", task.Title)
			fmt.Fprintf(w, "Category:\t%s\n", task.Category)
			fmt.Fprintf(w, "Description:\t%s\n", task.Description)
			if task.Amount.Valid {
				fmt.Fprintf(w, "Amount:\t%s\n", util.FormatCurrency(task.Amount.Int64))
			}
			if task.StartDate.Valid {
				fmt.Fprintf(w, "Start Date:\t%v\n", timediff.TimeDiff(task.StartDate.Time, timediff.WithLocale("pt-BR")))
			}
			if task.EndDate.Valid {
				fmt.Fprintf(w, "End Date:\t%v\n", timediff.TimeDiff(task.EndDate.Time, timediff.WithLocale("pt-BR")))
			}
			if task.NotificationDate.Valid {
				fmt.Fprintf(w, "Notification Date:\t%v\n", timediff.TimeDiff(task.NotificationDate.Time, timediff.WithLocale("pt-BR")))
			}
			fmt.Fprintln(w, "-----------------")

			w.Flush()
		} else {
			tasks, err := taskRepo.FindAll()

			if err != nil {
				log.Printf("Error fetching all tasks: %v\n", err)
				return
			}

			if len(tasks) == 0 {
				fmt.Println("No tasks found.")
				return
			}

			w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, '\t', 0)

			fmt.Fprintln(w, "--- All Tasks ---")

			for _, task := range tasks {
				fmt.Fprintf(w, "ID:\t%v\n", task.ID)
				fmt.Fprintf(w, "Title:\t%s\n", task.Title)
				fmt.Fprintf(w, "Category:\t%s\n", task.Category)
				fmt.Fprintf(w, "Description:\t%s\n", task.Description)
				if task.Amount.Valid {
					fmt.Fprintf(w, "Amount:\t%s\n", util.FormatCurrency(task.Amount.Int64))
				}
				if task.StartDate.Valid {
					fmt.Fprintf(w, "Start Date:\t%v\n", timediff.TimeDiff(task.StartDate.Time, timediff.WithLocale("pt-BR")))
				}
				if task.EndDate.Valid {
					fmt.Fprintf(w, "End Date:\t%v\n", timediff.TimeDiff(task.EndDate.Time, timediff.WithLocale("pt-BR")))
				}
				if task.NotificationDate.Valid {
					fmt.Fprintf(w, "Notification Date:\t%v\n", timediff.TimeDiff(task.NotificationDate.Time, timediff.WithLocale("pt-BR")))
				}
				fmt.Fprintln(w, "-----------------")
			}

			w.Flush()
		}
	},
}

func init() {
	readCmd.Flags().StringVar(&id, "id", "", "ID of the task to view")
}
