package model

import "time"

// Comment is a comment for a task.
type Comment struct {
	Text      string
	CreatedAt time.Time
	Task      Task
}
