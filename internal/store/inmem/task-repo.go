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
func newTaskRepo(db *inMemoryDb) *taskRepo {
	return &taskRepo{db: db}
}

// GetByColumnID returns all tasks with specific column ID.
func (r *taskRepo) GetByColumnID(id int) (tasks []model.Task, err error) {
	r.m.RLock()
	defer r.m.RUnlock()

	for _, t := range r.db.tasks {
		if t.ColumnID == id {
			tasks = append(tasks, t)
		}
	}

	return tasks, err
}

// Create creates and returns a new task.
func (r *taskRepo) Create(t model.Task) (model.Task, error) {
	r.m.Lock()
	defer r.m.Unlock()

	t.ID = len(r.db.tasks) + 1
	t.Index = len(r.db.tasks) + 1
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
func (r *taskRepo) Update(t model.Task) error {
	r.m.Lock()
	defer r.m.Unlock()

	if task, ok := r.db.tasks[t.ID]; ok {
		task.Name = t.Name
		task.Description = t.Description
		r.db.tasks[task.ID] = task
		return nil
	}

	return store.ErrNotFound
}

// DeleteByID deletes the task with specific ID.
func (r *taskRepo) DeleteByID(id int) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.db.tasks[id]; !ok {
		return store.ErrNotFound
	}

	for cid, c := range r.db.comments {
		if c.TaskID == id {
			delete(r.db.comments, cid)
		}
	}

	delete(r.db.tasks, id)
	return nil
}
