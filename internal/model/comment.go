package model

import "time"

// Comment is a comment for a task.
type Comment struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	TaskID    int       `json:"task_id"`
}
