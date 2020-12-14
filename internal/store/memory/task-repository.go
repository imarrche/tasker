package memory

import (
	"sync"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

// TaskRepository is an in memory task repository.
type TaskRepository struct {
	db *inMemoryDb
	m  sync.RWMutex
}

// NewTaskRepository creates and returns a new TaskRepository instance.
func NewTaskRepository(db *inMemoryDb) *TaskRepository {
	return &TaskRepository{db: db}
}

// GetAll returns all tasks.
func (r *TaskRepository) GetAll() ([]model.Task, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	tasks := []model.Task{}
	for _, t := range r.db.tasks {
		tasks = append(tasks, t)
	}

	return tasks, nil
}

// GetAllByColumnID returns all tasks that belong to specific column.
func (r *TaskRepository) GetAllByColumnID(id int) ([]model.Task, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	tasks := []model.Task{}
	for _, t := range r.db.tasks {
		if t.Column.ID == id {
			tasks = append(tasks, t)
		}
	}

	return tasks, nil
}

// Create creates and returns a new task.
func (r *TaskRepository) Create(t model.Task) (model.Task, error) {
	r.m.Lock()
	defer r.m.Unlock()

	t.ID = len(r.db.tasks) + 1
	r.db.tasks[t.ID] = t

	return t, nil
}

// GetByID returns a task with specific ID.
func (r *TaskRepository) GetByID(id int) (model.Task, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	if t, ok := r.db.tasks[id]; ok {
		return t, nil
	}

	return model.Task{}, store.ErrNotFound
}

// Update updates a task.
func (r *TaskRepository) Update(t model.Task) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.db.tasks[t.ID]; ok {
		r.db.tasks[t.ID] = t
		return nil
	}

	return store.ErrNotFound
}

// DeleteByID deletes a task with specific ID.
func (r *TaskRepository) DeleteByID(id int) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.db.tasks[id]; ok {
		delete(r.db.tasks, id)
		return nil
	}

	return store.ErrNotFound
}
