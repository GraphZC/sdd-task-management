package models

const (
	TaskStatusTodo      = "TODO"
	TaskStatusCompleted = "COMPLETED"
)

const (
	TaskPriorityLow    = 1
	TaskPriorityMedium = 2
	TaskPriorityHigh   = 3
)

type Task struct {
	ID          string `json:"id" db:"id"`
	UserID      string `json:"userId" db:"user_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Priority    int    `json:"priority" db:"priority"`
	Status      string `json:"status" db:"status"`
	CreatedAt   string `json:"createdAt" db:"created_at"`
	UpdatedAt   string `json:"updatedAt" db:"updated_at"`
}
