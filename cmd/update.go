package cmd

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/mlucas4330/todo-go-cli/internal/util"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task",
	Run: func(cmd *cobra.Command, args []string) {
		var setClauses []string
		var values []any

		if cmd.Flags().Changed("title") {
			setClauses = append(setClauses, "title")
			values = append(values, title)
		}
		if cmd.Flags().Changed("category") {
			setClauses = append(setClauses, "category")
			values = append(values, categoryStr)
		}
		if cmd.Flags().Changed("description") {
			setClauses = append(setClauses, "description")
			values = append(values, description)
		}
		if cmd.Flags().Changed("amount") {
			amountInt, err := strconv.Atoi(amountStr)
			if err != nil {
				log.Fatalf("Error converting amount '%s' to integer: %v", amountStr, err)
				return
			}
			setClauses = append(setClauses, "amount")
			values = append(values, amountInt)
		}
		if cmd.Flags().Changed("start-date") {
			startDateParsed, err := util.ParseDate(startDateStr)
			if err != nil {
				log.Fatalf("Error parsing end date: %v", err)
			}
			setClauses = append(setClauses, "start_date")
			values = append(values, startDateParsed)
		}
		if cmd.Flags().Changed("end-date") {
			endDateParsed, err := util.ParseDate(endDateStr)
			if err != nil {
				log.Fatalf("Error parsing notification date: %v", err)
			}
			setClauses = append(setClauses, "end_date")
			values = append(values, endDateParsed)
		}
		if cmd.Flags().Changed("notification-date") {
			notificationDateParsed, err := util.ParseDate(notificationDateStr)
			if err != nil {
				log.Fatalf("Error parsing notification date: %v", err)
			}
			setClauses = append(setClauses, "notification_date")
			values = append(values, notificationDateParsed)
		}

		if len(setClauses) == 0 {
			fmt.Println("No fields to update. Please use flags like --title, --category, --description, etc.")
			return
		}

		setClauses = append(setClauses, "updated_at")
		values = append(values, time.Now())

		err := taskRepo.Update(id, setClauses, values)
		if err != nil {
			fmt.Printf("Error updating task: %v\n", err)
			return
		}

		fmt.Println("Task updated successfully!")
	},
}

func init() {
	updateCmd.Flags().StringVar(&id, "id", "", "ID of the task to update")
	updateCmd.Flags().StringVar(&title, "title", "", "Title of the task")
	updateCmd.Flags().StringVar(&categoryStr, "category", "", "Category (Work, Personal, Shopping, Others)")
	updateCmd.Flags().StringVar(&description, "description", "", "Description of the task")
	updateCmd.Flags().StringVar(&amountStr, "amount", "", "Amount associated with the task")
	updateCmd.Flags().StringVar(&startDateStr, "start-date", "", "Start date (YYYY-MM-DD)")
	updateCmd.Flags().StringVar(&endDateStr, "end-date", "", "End date (YYYY-MM-DD)")
	updateCmd.Flags().StringVar(&notificationDateStr, "notification-date", "", "End date (YYYY-MM-DD)")
	updateCmd.MarkFlagRequired("id")
}
