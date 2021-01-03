package web

import (
	"sort"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

// columnService is the web column service.
type columnService struct {
	store store.Store
}

// newColumnService creates and returns a new columnService instance.
func newColumnService(s store.Store) *columnService {
	return &columnService{store: s}
}

// GetByProjectID returns all columns with specific project ID sorted by index.
func (s *columnService) GetByProjectID(id int) ([]model.Column, error) {
	cs, err := s.store.Columns().GetByProjectID(id)
	if err != nil {
		return nil, err
	}

	sort.SliceStable(cs, func(i, j int) bool {
		return cs[i].Index < cs[j].Index
	})

	return cs, nil
}

// Create creates a new column.
func (s *columnService) Create(c model.Column) (model.Column, error) {
	if err := s.Validate(c); err != nil {
		return model.Column{}, err
	}

	cs, err := s.store.Columns().GetByProjectID(c.ProjectID)
	if err != nil {
		return model.Column{}, err
	}
	c.Index = len(cs) + 1

	return s.store.Columns().Create(c)
}

// GetByID returns the column with specific ID.
func (s *columnService) GetByID(id int) (model.Column, error) {
	return s.store.Columns().GetByID(id)
}

// Update updates a column.
func (s *columnService) Update(c model.Column) (model.Column, error) {
	column, err := s.store.Columns().GetByID(c.ID)
	if err != nil {
		return model.Column{}, err
	}

	column.Name = c.Name
	if err := s.Validate(column); err != nil {
		return model.Column{}, err
	}

	return s.store.Columns().Update(column)
}

// MoveByID moves the column with specific ID left/right.
func (s *columnService) MoveByID(id int, left bool) error {
	c, err := s.store.Columns().GetByID(id)
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

	if left {
		c.Index--
		nextColumn.Index++
	} else {
		c.Index++
		nextColumn.Index--
	}
	if _, err = s.store.Columns().Update(nextColumn); err != nil {
		return err
	}
	_, err = s.store.Columns().Update(c)

	return err
}

// DeleteByID deletes the column with specific ID.
func (s *columnService) DeleteByID(id int) error {
	c, err := s.store.Columns().GetByID(id)
	if err != nil {
		return err
	}
	cs, err := s.store.Columns().GetByProjectID(c.ProjectID)
	if err != nil {
		return err
	}
	if len(cs) == 1 {
		return ErrLastColumn
	}

	nextIdx := c.Index - 1
	if nextIdx == 0 {
		nextIdx = 2
	}
	tasks, err := s.store.Tasks().GetByColumnID(c.ID)
	if err != nil {
		return err
	}
	var nextColumn model.Column
	for _, column := range cs {
		if column.Index == nextIdx {
			nextColumn = column
			break
		}
	}
	nextColumnTasks, err := s.store.Tasks().GetByColumnID(nextColumn.ID)
	if err != nil {
		return err
	}
	nextIdx = len(nextColumnTasks) + 1
	for _, t := range tasks {
		t.ColumnID = nextColumn.ID
		t.Index = nextIdx
		if _, err = s.store.Tasks().Update(t); err != nil {
			return err
		}
		nextIdx++
	}

	for _, column := range cs {
		if column.Index > c.Index {
			column.Index--
			if _, err = s.store.Columns().Update(column); err != nil {
				return err
			}
		}
	}

	return s.store.Columns().DeleteByID(id)
}

// Validate validates a column.
func (s *columnService) Validate(c model.Column) error {
	if len(c.Name) == 0 {
		return ErrNameIsRequired
	} else if len(c.Name) > 255 {
		return ErrNameIsTooLong
	}

	cs, err := s.store.Columns().GetByProjectID(c.ProjectID)
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
