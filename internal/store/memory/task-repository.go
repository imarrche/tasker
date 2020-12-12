package memory

import (
	"sync"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

// TaskRepository is an in memory task repository.
type TaskRepository struct {
	store *Store
	m     sync.RWMutex
}

// NewTaskRepository creates and returns a new TaskRepository instance.
func NewTaskRepository() *TaskRepository {
	return &TaskRepository{store: NewStore()}
}

// GetAll returns all Tasks.
func (r *TaskRepository) GetAll() ([]model.Task, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	tasks := []model.Task{}
	for _, t := range r.store.tasks {
		tasks = append(tasks, t)
	}

	return tasks, nil
}

// Create creates and returns a new task.
func (r *TaskRepository) Create(t model.Task) (model.Task, error) {
	r.m.Lock()
	defer r.m.Unlock()

	t.ID = len(r.store.tasks) + 1
	r.store.tasks[t.ID] = t

	return t, nil
}

// GetByID returns a task with specific ID.
func (r *TaskRepository) GetByID(id int) (model.Task, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	if t, ok := r.store.tasks[id]; ok {
		return t, nil
	}

	return model.Task{}, store.ErrNotFound
}

// Update updates a task.
func (r *TaskRepository) Update(t model.Task) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.store.tasks[t.ID]; ok {
		r.store.tasks[t.ID] = t
		return nil
	}

	return store.ErrNotFound
}

// Delete deletes a task.
func (r *TaskRepository) Delete(t model.Task) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.store.tasks[t.ID]; ok {
		delete(r.store.tasks, t.ID)
		return nil
	}

	return store.ErrNotFound
}
