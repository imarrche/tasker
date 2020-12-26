package pg

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver.

	"github.com/imarrche/tasker/internal/config"
)

// Store is PostgreSQL store.
type Store struct {
	config config.PostgreSQL
	db     *sql.DB
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

// Close closes a connection with PostgreSQL.
func (s *Store) Close() error {
	return s.db.Close()
}
