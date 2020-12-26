package model

// Project is a project team is working on.
// Visually, it is a board with tasks in columns that represent tasks' progress.
type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
