package model

import (
	"database/sql"
	"strings"
	"time"
)

type Category int

const (
	Work Category = iota
	Personal
	Shopping
	Others
	Unknown
)

func (c Category) String() string {
	switch c {
	case Work:
		return "Work"
	case Personal:
		return "Personal"
	case Shopping:
		return "Shopping"
	case Others:
		return "Others"
	default:
		return "Unknown"
	}
}

func ParseCategory(input string) Category {
	clean := strings.TrimSpace(strings.ToLower(input))

	switch clean {
	case "work":
		return Work
	case "personal":
		return Personal
	case "shopping":
		return Shopping
	case "others":
		return Others
	default:
		return Unknown
	}
}

type Task struct {
	ID               int           `json:"id"`
	Title            string        `json:"title"`
	Description      string        `json:"description"`
	Amount           sql.NullInt64 `json:"amount"`
	Category         Category      `json:"category"`
	StartDate        sql.NullTime  `json:"start_date"`
	EndDate          sql.NullTime  `json:"end_date"`
	NotificationDate sql.NullTime  `json:"notification_date"`
	CreatedAt        time.Time     `json:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at"`
}
