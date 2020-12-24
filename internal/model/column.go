package model

// Column is a column in a project board for grouping tasks by their progress.
type Column struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Index     int    `json:"index"`
	ProjectID int    `json:"project_id"`
}
