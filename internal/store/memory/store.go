package memory

import "github.com/imarrche/tasker/internal/model"

// Store is an in memory store.
type Store struct {
	projects map[int]model.Project
	columns  map[int]model.Column
	tasks    map[int]model.Task
}

// NewStore creates and returns a new Store instance.
func NewStore() *Store {
	return &Store{
		projects: map[int]model.Project{},
		columns:  map[int]model.Column{},
		tasks:    map[int]model.Task{},
	}
}
