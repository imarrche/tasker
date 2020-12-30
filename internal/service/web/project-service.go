package web

import (
	"database/sql"
	"sort"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

// projectService is the web project service.
type projectService struct {
	store store.Store
}

// newProjectService creates and returns a new projectService instance.
func newProjectService(s store.Store) *projectService {
	return &projectService{store: s}
}

// GetAll returns all projects sorted alphabetically by name.
func (s *projectService) GetAll() ([]model.Project, error) {
	ps, err := s.store.Projects().GetAll()
	if err != nil {
		return nil, err
	}

	sort.SliceStable(ps, func(i, j int) bool {
		return ps[i].Name < ps[j].Name
	})

	return ps, nil
}

// Create creates a new project.
func (s *projectService) Create(p model.Project) (model.Project, error) {
	if err := s.Validate(p); err != nil {
		return model.Project{}, err
	}

	p, err := s.store.Projects().Create(p)
	if err != nil {
		return model.Project{}, err
	}

	_, err = s.store.Columns().Create(
		model.Column{Name: "default", Index: 1, ProjectID: p.ID},
	)
	if err != nil {
		return model.Project{}, err
	}

	return p, nil
}

// GetByID returns the project with specific ID.
func (s *projectService) GetByID(id int) (model.Project, error) {
	return s.store.Projects().GetByID(id)
}

// Update updates a project.
func (s *projectService) Update(p model.Project) (model.Project, error) {
	if err := s.Validate(p); err != nil {
		return model.Project{}, err
	}

	_, err := s.store.Projects().GetByID(p.ID)
	if err == sql.ErrNoRows {
		return model.Project{}, store.ErrNotFound
	}
	if err != nil {
		return model.Project{}, err
	}

	return s.store.Projects().Update(p)
}

// DeleteByID deletes the project with specific ID.
func (s *projectService) DeleteByID(id int) error {
	_, err := s.store.Projects().GetByID(id)
	if err == sql.ErrNoRows {
		return store.ErrNotFound
	}
	if err != nil {
		return err
	}

	return s.store.Projects().DeleteByID(id)
}

// Validate validates a project.
func (s *projectService) Validate(p model.Project) error {
	if len(p.Name) == 0 {
		return ErrNameIsRequired
	} else if len(p.Name) > 500 {
		return ErrNameIsTooLong
	}

	if len(p.Description) > 1000 {
		return ErrDescriptionIsTooLong
	}

	return nil
}
