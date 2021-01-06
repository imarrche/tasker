package inmem

import (
	"time"

	"github.com/imarrche/tasker/internal/model"
)

// TestStoreWithFixtures creates and returns in memory store instance with fixtures
// for testing.
func TestStoreWithFixtures() *Store {
	s := NewStore()
	s.db = &inMemoryDb{
		projects: map[int]model.Project{
			1: {ID: 1, Name: "Project 1"},
			2: {ID: 2, Name: "Project 2"},
		},
		columns: map[int]model.Column{
			1: {ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
			2: {ID: 2, Name: "Column 2", Index: 2, ProjectID: 1},
			3: {ID: 3, Name: "Column 3", Index: 1, ProjectID: 2},
		},
		tasks: map[int]model.Task{
			1: {ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
			2: {ID: 2, Name: "Task 2", Index: 2, ColumnID: 1},
			3: {ID: 3, Name: "Task 3", Index: 1, ColumnID: 2},
		},
		comments: map[int]model.Comment{
			1: {ID: 1, Text: "Comment 1", CreatedAt: time.Now(), TaskID: 1},
			2: {ID: 2, Text: "Comment 2", CreatedAt: time.Now(), TaskID: 1},
			3: {ID: 3, Text: "Comment 3", CreatedAt: time.Now(), TaskID: 2},
		},
	}

	return s
}
