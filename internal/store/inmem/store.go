package inmem

import (
	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

type inMemoryDb struct {
	projects map[int]model.Project
	columns  map[int]model.Column
	tasks    map[int]model.Task
	comments map[int]model.Comment
}

func newInMemoryDb() *inMemoryDb {
	return &inMemoryDb{
		projects: map[int]model.Project{},
		columns:  map[int]model.Column{},
		tasks:    map[int]model.Task{},
		comments: map[int]model.Comment{},
	}
}

// Store is the in memory store.
type Store struct {
	db          *inMemoryDb
	projectRepo *projectRepo
	columnRepo  *columnRepo
	taskRepo    *taskRepo
	commentRepo *commentRepo
}

// NewStore creates and returns a new Store instance.
func NewStore() *Store {
	return &Store{
		db: newInMemoryDb(),
	}
}

// Open opens the store.
func (s *Store) Open() error {
	return nil
}

// Projects returns the project repository.
func (s *Store) Projects() store.ProjectRepo {
	if s.projectRepo == nil {
		s.projectRepo = newProjectRepo(s.db)
	}

	return s.projectRepo
}

// Columns returns the column repository.
func (s *Store) Columns() store.ColumnRepo {
	if s.columnRepo == nil {
		s.columnRepo = newColumnRepo(s.db)
	}

	return s.columnRepo
}

// Tasks returns the task repository.
func (s *Store) Tasks() store.TaskRepo {
	if s.taskRepo == nil {
		s.taskRepo = newTaskRepo(s.db)
	}

	return s.taskRepo
}

// Comments returns the comment repository.
func (s *Store) Comments() store.CommentRepo {
	if s.commentRepo == nil {
		s.commentRepo = newCommentRepo(s.db)
	}

	return s.commentRepo
}

// Close closes the store.
func (s *Store) Close() error {
	return nil
}
