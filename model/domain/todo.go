package domain

type TodoType string

// type of todo
const (
	ImportantUrgent       TodoType = "important_urgent"
	ImportantNotUrgent    TodoType = "important_not_urgent"
	NotImportantUrgent    TodoType = "not_important_urgent"
	NotImportantNotUrgent TodoType = "not_important_not_urgent"
)

// Todo struct represents a table todos
type Todo struct {
	Id          int
	Title       string
	Description string
	Type        TodoType
	IsDone      bool
	CreatedAt   string
	UpdatedAt   string
}
