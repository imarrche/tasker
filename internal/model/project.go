package model

// Project is a project team is working on.
// Visually, it is a board with tasks in columns that represent tasks' progress.
type Project struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}
