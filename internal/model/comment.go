package model

import "time"

// Comment is a comment for a task.
type Comment struct {
	ID        int
	Text      string
	CreatedAt time.Time
	TaskID    int
}
