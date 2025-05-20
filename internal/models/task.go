package models

import "time"

type Category int

const (
	Work Category = iota
	Personal
	Shopping
	Others
)

func (c Category) String() string {
	return [...]string{"Work", "Personal", "Shopping", "Others"}[c]
}

type Task struct {
	ID               int
	Title            string
	Description      string
	Amount           int64
	Category         Category
	StartDate        time.Time
	EndDate          time.Time
	NotificationDate time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
