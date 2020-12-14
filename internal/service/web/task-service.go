package web

import (
	"sort"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/service"
	"github.com/imarrche/tasker/internal/store"
)

// TaskService is a web task service.
type TaskService struct {
	columnRepo store.ColumnRepository
	taskRepo   store.TaskRepository
}

// NewTaskService creates and returns a new TaskService instance.
func NewTaskService(cr store.ColumnRepository, tr store.TaskRepository) service.TaskService {
	return &TaskService{columnRepo: cr, taskRepo: tr}
}

// GetAll returns all tasks sorted by index.
func (s *TaskService) GetAll() ([]model.Task, error) {
	ts, err := s.taskRepo.GetAll()
	if err != nil {
		return nil, err
	}

	sort.SliceStable(ts, func(i, j int) bool {
		return ts[i].Index < ts[j].Index
	})

	return ts, nil
}

// Create creates a new task.
func (s *TaskService) Create(t model.Task) (model.Task, error) {
	if err := s.Validate(t); err != nil {
		return model.Task{}, err
	}

	return s.taskRepo.Create(t)
}

// GetByID returns task with specific ID.
func (s *TaskService) GetByID(id int) (model.Task, error) {
	return s.taskRepo.GetByID(id)
}

// Update updates a task.
func (s *TaskService) Update(t model.Task) error {
	if err := s.Validate(t); err != nil {
		return err
	}

	return s.taskRepo.Update(t)
}

// DeleteByID deletes a task with specific ID.
func (s *TaskService) DeleteByID(id int) error {
	// TODO: delete all task's comments also.
	return s.taskRepo.DeleteByID(id)
}

// Validate validates a task.
func (s *TaskService) Validate(t model.Task) error {
	if len(t.Name) == 0 {
		return ErrNameIsRequired
	} else if len(t.Name) > 500 {
		return ErrNameIsTooLong
	}

	if len(t.Description) > 5000 {
		return ErrDescriptionIsTooLong
	}

	if _, err := s.columnRepo.GetByID(t.Column.ID); err != nil {
		return ErrInvalidColumn
	}

	return nil
}
