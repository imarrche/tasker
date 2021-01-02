package inmem

import (
	"sync"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

// commentRepo is the comment repository for in memory store.
type commentRepo struct {
	db *inMemoryDb
	m  sync.RWMutex
}

// newCommentRepo creates and returns a new commentRepo instance.
func newCommentRepo(db *inMemoryDb) *commentRepo { return &commentRepo{db: db} }

// GetByTaskID returns all comments with specific task ID.
func (r *commentRepo) GetByTaskID(id int) ([]model.Comment, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	if _, ok := r.db.tasks[id]; !ok {
		return nil, store.ErrNotFound
	}

	cs := []model.Comment{}
	for _, c := range r.db.comments {
		if c.TaskID == id {
			cs = append(cs, c)
		}
	}

	return cs, nil
}

// Create creates and returns a new comment.
func (r *commentRepo) Create(c model.Comment) (model.Comment, error) {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.db.tasks[c.TaskID]; !ok {
		return model.Comment{}, store.ErrDbQuery
	}

	c.ID = len(r.db.comments) + 1
	r.db.comments[c.ID] = c

	return c, nil
}

// GetByID returns the comment with specific ID.
func (r *commentRepo) GetByID(id int) (model.Comment, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	if c, ok := r.db.comments[id]; ok {
		return c, nil
	}

	return model.Comment{}, store.ErrNotFound
}

// Update updates the comment.
func (r *commentRepo) Update(c model.Comment) (model.Comment, error) {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.db.comments[c.ID]; !ok {
		return model.Comment{}, store.ErrNotFound
	}
	if _, ok := r.db.tasks[c.TaskID]; !ok {
		return model.Comment{}, store.ErrDbQuery
	}

	r.db.comments[c.ID] = c

	return c, nil
}

// DeleteByID deletes the comment with specific ID.
func (r *commentRepo) DeleteByID(id int) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.db.comments[id]; !ok {
		return store.ErrNotFound
	}

	delete(r.db.comments, id)

	return nil
}
