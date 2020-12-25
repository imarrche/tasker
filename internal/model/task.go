package model

// Task is a project task that's being moved across columns according to its progress
// and within column according to its priority.
type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Index       int    `json:"index"`
	ColumnID    int    `json:"column_id"`
}
