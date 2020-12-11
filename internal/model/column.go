package model

// Column is a column in a project board for grouping tasks by their progress.
type Column struct {
	Name    string
	Index   int
	Project Project
}
