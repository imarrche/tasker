package model

// Task is a project task that's being moved across columns according to its progress
// and within column according to its priority.
type Task struct {
	ID          int
	Name        string
	Description string
	Index       int
	ColumnID    int
}
