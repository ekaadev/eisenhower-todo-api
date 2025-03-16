package domain

import (
	"database/sql"
	"time"
)

type TodoType string

// type of todo
const (
	UrgentImportant       TodoType = "urgent_important"
	NotUrgentImportant    TodoType = "not_urgent_important"
	UrgentNotImportant    TodoType = "urgent_not_important"
	NotUrgentNotImportant TodoType = "not_urgent_not_important"
)

// Todo struct represents a table todos
type Todo struct {
	Id          int
	Title       string
	Description sql.NullString
	Type        TodoType
	IsDone      bool
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
