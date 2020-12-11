package model

// Task is a project task that's being moved between columns according to its progress.
type Task struct {
	Name        string
	Description string
	Index       string
	Column      Column
}
