package memory

import (
	"sync"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

// ColumnRepository is an in memory column repository.
type ColumnRepository struct {
	db *inMemoryDb
	m  sync.RWMutex
}

// NewColumnRepository creates and returns a new ColumnRepository instance.
func NewColumnRepository(db *inMemoryDb) *ColumnRepository {
	return &ColumnRepository{db: db}
}

// GetAll returns all columns.
func (r *ColumnRepository) GetAll() ([]model.Column, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	columns := []model.Column{}
	for _, c := range r.db.columns {
		columns = append(columns, c)
	}

	return columns, nil
}

// GetAllByProjectID returns all columns that belong to specific project.
func (r *ColumnRepository) GetAllByProjectID(id int) ([]model.Column, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	columns := []model.Column{}
	for _, c := range r.db.columns {
		if c.Project.ID == id {
			columns = append(columns, c)
		}
	}

	return columns, nil
}

// Create creates and returns a new column.
func (r *ColumnRepository) Create(c model.Column) (model.Column, error) {
	r.m.Lock()
	defer r.m.Unlock()

	c.ID = len(r.db.columns) + 1
	r.db.columns[c.ID] = c

	return c, nil
}

// GetByID returns a column with specific ID.
func (r *ColumnRepository) GetByID(id int) (model.Column, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	if c, ok := r.db.columns[id]; ok {
		return c, nil
	}

	return model.Column{}, store.ErrNotFound
}

// Update updates a column.
func (r *ColumnRepository) Update(c model.Column) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.db.columns[c.ID]; ok {
		r.db.columns[c.ID] = c
		return nil
	}

	return store.ErrNotFound
}

// DeleteByID deletes a column with specific ID.
func (r *ColumnRepository) DeleteByID(id int) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.db.columns[id]; ok {
		delete(r.db.columns, id)
		return nil
	}

	return store.ErrNotFound
}
