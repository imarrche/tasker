package web

import (
	"sort"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/service"
	"github.com/imarrche/tasker/internal/store"
)

// ColumnService is a web column service.
type ColumnService struct {
	projectRepo store.ProjectRepository
	columnRepo  store.ColumnRepository
	taskRepo    store.TaskRepository
}

// NewColumnService creates and returns a new ColumnService instance.
func NewColumnService(cr store.ColumnRepository, tr store.TaskRepository) service.ColumnService {
	return &ColumnService{columnRepo: cr, taskRepo: tr}
}

// GetAll returns all columns sorted alphabetically by name.
func (s *ColumnService) GetAll() ([]model.Column, error) {
	cs, err := s.columnRepo.GetAll()
	if err != nil {
		return nil, err
	}

	sort.SliceStable(cs, func(i, j int) bool {
		return cs[i].Index < cs[j].Index
	})

	return cs, nil
}

// Create creates a new column.
func (s *ColumnService) Create(c model.Column) (model.Column, error) {
	if err := s.Validate(c); err != nil {
		return model.Column{}, err
	}

	return s.columnRepo.Create(c)
}

// GetByID returns column with specific ID.
func (s *ColumnService) GetByID(id int) (model.Column, error) {
	return s.columnRepo.GetByID(id)
}

// Update updates a column.
func (s *ColumnService) Update(c model.Column) error {
	if err := s.Validate(c); err != nil {
		return err
	}

	return s.columnRepo.Update(c)
}

// DeleteByID deletes a column with specific ID.
func (s *ColumnService) DeleteByID(id int) error {
	c, err := s.columnRepo.GetByID(id)
	if err != nil {
		return err
	}

	cs, err := s.columnRepo.GetAllByProjectID(c.Project.ID)
	if err != nil {
		return err
	}
	if len(cs) == 1 {
		return ErrLastColumn
	}

	_, err = s.taskRepo.GetAllByColumnID(c.ID)
	if err != nil {
		return err
	}
	// TODO: move tasks to the column to the left.

	return s.columnRepo.DeleteByID(id)
}

// Validate validates a column.
func (s *ColumnService) Validate(c model.Column) error {
	if len(c.Name) == 0 {
		return ErrNameIsRequired
	} else if len(c.Name) > 255 {
		return ErrNameIsTooLong
	}

	if c.Project.ID == 0 {
		return ErrProjectIsRequired
	}

	cs, err := s.columnRepo.GetAllByProjectID(c.Project.ID)
	if err != nil {
		return err
	}
	for _, column := range cs {
		if column.Name == c.Name {
			return ErrColumnAlreadyExists
		}
	}

	return nil
}
