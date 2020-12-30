package pg

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver.

	"github.com/imarrche/tasker/internal/config"
	"github.com/imarrche/tasker/internal/store"
)

// Store is PostgreSQL store.
type Store struct {
	config      config.PostgreSQL
	db          *sql.DB
	projectRepo *projectRepo
	columnRepo  *columnRepo
	taskRepo    *taskRepo
	commentRepo *commentRepo
}

// New creates new Store instance.
func New(config config.PostgreSQL) *Store {
	return &Store{config: config}
}

// Open open a connection with PostgreSQL.
func (s *Store) Open() error {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		s.config.Host, s.config.Port, s.config.User,
		s.config.Password, s.config.DbName, s.config.SSLMode,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}

	s.db = db

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

// Close closes a connection with PostgreSQL.
func (s *Store) Close() error {
	return s.db.Close()
}
