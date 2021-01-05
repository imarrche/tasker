package inmem

import (
	"sync"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

// taskRepo is the task repository for in memory store.
type taskRepo struct {
	db *inMemoryDb
	m  sync.RWMutex
}

// newTaskRepo creates and returns a new taskRepo instance.
func newTaskRepo(db *inMemoryDb) *taskRepo { return &taskRepo{db: db} }

// GetByColumnID returns all tasks with specific column ID.
func (r *taskRepo) GetByColumnID(id int) ([]model.Task, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	if _, ok := r.db.columns[id]; !ok {
		return nil, store.ErrNotFound
	}

	ts := []model.Task{}
	for _, t := range r.db.tasks {
		if t.ColumnID == id {
			ts = append(ts, t)
		}
	}

	return ts, nil
}

// Create creates and returns a new task.
func (r *taskRepo) Create(t model.Task) (model.Task, error) {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.db.columns[t.ColumnID]; !ok {
		return model.Task{}, store.ErrDbQuery
	}

	t.ID = len(r.db.tasks) + 1
	r.db.tasks[t.ID] = t

	return t, nil
}

// GetByID returns the task with specific ID.
func (r *taskRepo) GetByID(id int) (model.Task, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	if t, ok := r.db.tasks[id]; ok {
		return t, nil
	}

	return model.Task{}, store.ErrNotFound
}

// GetByIndexAndColumnID returns the task with specific index and column ID.
func (r *taskRepo) GetByIndexAndColumnID(index, id int) (model.Task, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	for _, t := range r.db.tasks {
		if t.Index == index && t.ColumnID == id {
			return t, nil
		}
	}

	return model.Task{}, store.ErrNotFound
}

// Update updates the task.
func (r *taskRepo) Update(t model.Task) (model.Task, error) {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.db.tasks[t.ID]; !ok {
		return model.Task{}, store.ErrNotFound
	}
	if _, ok := r.db.columns[t.ColumnID]; !ok {
		return model.Task{}, store.ErrDbQuery
	}

	r.db.tasks[t.ID] = t

	return t, nil
}

// DeleteByID deletes the task with specific ID.
func (r *taskRepo) DeleteByID(id int) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.db.tasks[id]; !ok {
		return store.ErrNotFound
	}

	for commentID, comment := range r.db.comments {
		if comment.TaskID == id {
			delete(r.db.comments, commentID)
		}
	}
	delete(r.db.tasks, id)

	return nil
}
