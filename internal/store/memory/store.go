package memory

import "github.com/imarrche/tasker/internal/model"

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

// Store is an in memory store.
type Store struct {
	db       *inMemoryDb
	projects *ProjectRepository
	columns  *ColumnRepository
	tasks    *TaskRepository
	comments *CommentRepository
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

// Projects returns project repository.
func (s *Store) Projects() *ProjectRepository {
	if s.projects == nil {
		s.projects = NewProjectRepository(s.db)
	}

	return s.projects
}

// Columns returns column repository.
func (s *Store) Columns() *ColumnRepository {
	if s.columns == nil {
		s.columns = NewColumnRepository(s.db)
	}

	return s.columns
}

// Tasks returns task repository.
func (s *Store) Tasks() *TaskRepository {
	if s.tasks == nil {
		s.tasks = NewTaskRepository(s.db)
	}

	return s.tasks
}

// Comments returns comment repository.
func (s *Store) Comments() *CommentRepository {
	if s.comments == nil {
		s.comments = NewCommentRepository(s.db)
	}

	return s.comments
}

// Close closes the store.
func (s *Store) Close() error {
	return nil
}
