package memory

import (
	"sync"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

// ColumnRepository is an in memory column repository.
type ColumnRepository struct {
	store *Store
	m     sync.RWMutex
}

// NewColumnRepository creates and returns a new ColumnRepository instance.
func NewColumnRepository() *ColumnRepository {
	return &ColumnRepository{store: NewStore()}
}

// GetAll returns all columns.
func (r *ColumnRepository) GetAll() ([]model.Column, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	columns := []model.Column{}
	for _, c := range r.store.columns {
		columns = append(columns, c)
	}

	return columns, nil
}

// Create creates and returns a new column.
func (r *ColumnRepository) Create(c model.Column) (model.Column, error) {
	r.m.Lock()
	defer r.m.Unlock()

	c.ID = len(r.store.columns) + 1
	r.store.columns[c.ID] = c

	return c, nil
}

// GetByID returns a column with specific ID.
func (r *ColumnRepository) GetByID(id int) (model.Column, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	if c, ok := r.store.columns[id]; ok {
		return c, nil
	}

	return model.Column{}, store.ErrNotFound
}

// Update updates a column.
func (r *ColumnRepository) Update(c model.Column) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.store.columns[c.ID]; ok {
		r.store.columns[c.ID] = c
		return nil
	}

	return store.ErrNotFound
}

// Delete deletes a column.
func (r *ColumnRepository) Delete(c model.Column) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.store.columns[c.ID]; ok {
		delete(r.store.columns, c.ID)
		return nil
	}

	return store.ErrNotFound
}
