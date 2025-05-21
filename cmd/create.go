package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/mlucas4330/todo-go-cli/internal/model"
	"github.com/mlucas4330/todo-go-cli/internal/util"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a task",
	Run: func(cmd *cobra.Command, args []string) {
		category := model.ParseCategory(categoryStr)
		var startDate, endDate, notificationDate sql.NullTime
		var amount sql.NullInt64
		var err error

		if cmd.Flags().Changed("start-date") {
			parsedDate, err := util.ParseDate(startDateStr)
			if err != nil {
				log.Fatalf("Error parsing start date: %v", err)
			}
			startDate = sql.NullTime{Time: parsedDate, Valid: true}
		}

		if cmd.Flags().Changed("end-date") {
			parsedDate, err := util.ParseDate(endDateStr)
			if err != nil {
				log.Fatalf("Error parsing end date: %v", err)
			}
			endDate = sql.NullTime{Time: parsedDate, Valid: true}
		}

		if cmd.Flags().Changed("notification-date") {
			parsedDate, err := util.ParseDate(notificationDateStr)
			if err != nil {
				log.Fatalf("Error parsing notification date: %v", err)
			}
			notificationDate = sql.NullTime{Time: parsedDate, Valid: true}
		}

		if cmd.Flags().Changed("amount") {
			parsedAmount, err := strconv.Atoi(amountStr)
			if err != nil {
				log.Fatalf("Error converting amount: %v", err)
			}
			amount = sql.NullInt64{Int64: int64(parsedAmount), Valid: true}
		}

		task := model.Task{
			Title:            title,
			Description:      description,
			Amount:           amount,
			Category:         category,
			StartDate:        startDate,
			EndDate:          endDate,
			NotificationDate: notificationDate,
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		}

		err = taskRepo.Create(&task)
		if err != nil {
			fmt.Printf("Error creating task: %v\n", err)
			return
		}

		fmt.Println("Task created successfully!")
	},
}

func init() {
	createCmd.Flags().StringVar(&title, "title", "", "Title of the task")
	createCmd.Flags().StringVar(&categoryStr, "category", "", "Category (Work, Personal, Shopping, Others)")
	createCmd.Flags().StringVar(&description, "description", "", "Description of the task")
	createCmd.Flags().StringVar(&amountStr, "amount", "", "Amount associated with the task")
	createCmd.Flags().StringVar(&startDateStr, "start-date", "", "Start date (YYYY-MM-DD HH:MM:SS)")
	createCmd.Flags().StringVar(&endDateStr, "end-date", "", "End date (YYYY-MM-DD HH:MM:SS)")
	createCmd.Flags().StringVar(&notificationDateStr, "notification-date", "", "End date (YYYY-MM-DD HH:MM:SS)")
	createCmd.MarkFlagRequired("title")
	createCmd.MarkFlagRequired("description")
	createCmd.MarkFlagRequired("category")
}
