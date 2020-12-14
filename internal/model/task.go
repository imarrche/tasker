package model

// Task is a project task that's being moved between columns according to its progress.
type Task struct {
	ID          int
	Name        string
	Description string
	Index       int
	Column      Column
}
