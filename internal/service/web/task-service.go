package web

import (
	"sort"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

// taskService is the web task service.
type taskService struct {
	store store.Store
}

// newTaskService creates and returns a new taskService instance.
func newTaskService(s store.Store) *taskService {
	return &taskService{store: s}
}

// GetByColumnID returns all tasks with specific column ID sorted by index.
func (s *taskService) GetByColumnID(id int) ([]model.Task, error) {
	ts, err := s.store.Tasks().GetByColumnID(id)
	if err != nil {
		return nil, err
	}

	sort.SliceStable(ts, func(i, j int) bool {
		return ts[i].Index < ts[j].Index
	})

	return ts, nil
}

// Create creates a new task.
func (s *taskService) Create(t model.Task) (model.Task, error) {
	if err := s.Validate(t); err != nil {
		return model.Task{}, err
	}

	if _, err := s.store.Columns().GetByID(t.ColumnID); err != nil {
		return model.Task{}, err
	}

	return s.store.Tasks().Create(t)
}

// GetByID returns the task with specific ID.
func (s *taskService) GetByID(id int) (model.Task, error) {
	return s.store.Tasks().GetByID(id)
}

// Update updates a task.
func (s *taskService) Update(t model.Task) (model.Task, error) {
	if err := s.Validate(t); err != nil {
		return model.Task{}, err
	}

	return s.store.Tasks().Update(t)
}

func (s *taskService) MoveToColumnByID(id int, left bool) error {
	t, err := s.store.Tasks().GetByID(id)
	if err != nil {
		return err
	}
	c, err := s.store.Columns().GetByID(t.ColumnID)
	if err != nil {
		return err
	}

	nextIdx := c.Index + 1
	if left {
		nextIdx = c.Index - 1
	}
	nextColumn, err := s.store.Columns().GetByIndexAndProjectID(nextIdx, c.ProjectID)
	if err != nil {
		return err
	}
	nextColumnTasks, err := s.store.Tasks().GetByColumnID(nextColumn.ID)
	if err != nil {
		return err
	}

	tasks, err := s.store.Tasks().GetByColumnID(c.ID)
	if err != nil {
		return err
	}
	for _, task := range tasks {
		if task.Index > t.Index {
			task.Index--
			if _, err = s.store.Tasks().Update(task); err != nil {
				return err
			}
		}
	}

	t.ColumnID = nextColumn.ID
	t.Index = len(nextColumnTasks) + 1
	t, err = s.store.Tasks().Update(t)
	return err
}

func (s *taskService) MoveByID(id int, up bool) error {
	t, err := s.store.Tasks().GetByID(id)
	if err != nil {
		return err
	}

	nextIdx := t.Index + 1
	if up {
		nextIdx = t.Index - 1
	}
	nextTask, err := s.store.Tasks().GetByIndexAndColumnID(nextIdx, t.ColumnID)
	if err != nil {
		return err
	}

	if up {
		t.Index--
		nextTask.Index++
	} else {
		t.Index++
		nextTask.Index--
	}
	if _, err = s.store.Tasks().Update(nextTask); err != nil {
		return err
	}

	_, err = s.store.Tasks().Update(t)
	return err
}

// DeleteByID deletes the task with specific ID.
func (s *taskService) DeleteByID(id int) error {
	t, err := s.store.Tasks().GetByID(id)
	if err != nil {
		return err
	}

	ts, err := s.store.Tasks().GetByColumnID(t.ColumnID)
	for _, task := range ts {
		if task.Index > t.Index {
			task.Index--
			if _, err = s.store.Tasks().Update(task); err != nil {
				return err
			}
		}
	}

	return s.store.Tasks().DeleteByID(id)
}

// Validate validates a task.
func (s *taskService) Validate(t model.Task) error {
	if len(t.Name) == 0 {
		return ErrNameIsRequired
	} else if len(t.Name) > 500 {
		return ErrNameIsTooLong
	}

	if len(t.Description) > 5000 {
		return ErrDescriptionIsTooLong
	}

	return nil
}
