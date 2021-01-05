package inmem

import (
	"sync"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

// columnRepo is the column repository for in memory store.
type columnRepo struct {
	db *inMemoryDb
	m  sync.RWMutex
}

// newColumnRepo creates and returns a new columnRepo instance.
func newColumnRepo(db *inMemoryDb) *columnRepo { return &columnRepo{db: db} }

// GetByProjectID returns all columns with specific project ID.
func (r *columnRepo) GetByProjectID(id int) ([]model.Column, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	if _, ok := r.db.projects[id]; !ok {
		return nil, store.ErrNotFound
	}

	cs := []model.Column{}
	for _, c := range r.db.columns {
		if c.ProjectID == id {
			cs = append(cs, c)
		}
	}

	return cs, nil
}

// Create creates and returns a new column.
func (r *columnRepo) Create(c model.Column) (model.Column, error) {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.db.projects[c.ProjectID]; !ok {
		return model.Column{}, store.ErrDbQuery
	}

	c.ID = len(r.db.columns) + 1
	r.db.columns[c.ID] = c

	return c, nil
}

// GetByID returns the column with specific ID.
func (r *columnRepo) GetByID(id int) (model.Column, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	if c, ok := r.db.columns[id]; ok {
		return c, nil
	}

	return model.Column{}, store.ErrNotFound
}

// GetByIndexAndProjectID returns the column with specific index and project ID.
func (r *columnRepo) GetByIndexAndProjectID(index, id int) (model.Column, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	for _, c := range r.db.columns {
		if c.Index == index && c.ProjectID == id {
			return c, nil
		}
	}

	return model.Column{}, store.ErrNotFound
}

// Update updates the column.
func (r *columnRepo) Update(c model.Column) (model.Column, error) {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.db.columns[c.ID]; !ok {
		return model.Column{}, store.ErrNotFound
	}
	if _, ok := r.db.projects[c.ProjectID]; !ok {
		return model.Column{}, store.ErrDbQuery
	}

	r.db.columns[c.ID] = c

	return c, nil
}

// DeleteByID deletes the column with specific ID.
func (r *columnRepo) DeleteByID(id int) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.db.columns[id]; !ok {
		return store.ErrNotFound
	}

	for taskID, task := range r.db.tasks {
		if task.ColumnID == id {
			for commentID, comment := range r.db.comments {
				if comment.TaskID == taskID {
					delete(r.db.comments, commentID)
				}
			}
			delete(r.db.tasks, taskID)
		}
	}
	delete(r.db.columns, id)

	return nil
}
