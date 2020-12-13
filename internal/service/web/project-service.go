package web

import (
	"sort"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/service"
	"github.com/imarrche/tasker/internal/store"
)

// ProjectService is a web project service.
type ProjectService struct {
	projectRepo store.ProjectRepository
	columnRepo  store.ColumnRepository
}

// NewProjectService creates and returns a new ProjectService instance.
func NewProjectService(pr store.ProjectRepository, cr store.ColumnRepository) service.ProjectService {
	return &ProjectService{projectRepo: pr, columnRepo: cr}
}

// GetAll returns all projects sorted alphabetically by name.
func (s *ProjectService) GetAll() ([]model.Project, error) {
	ps, err := s.projectRepo.GetAll()
	if err != nil {
		return nil, err
	}

	sort.SliceStable(ps, func(i, j int) bool {
		return ps[i].Name < ps[j].Name
	})

	return ps, nil
}

// Create creates a new project.
func (s *ProjectService) Create(p model.Project) (model.Project, error) {
	if err := s.Validate(p); err != nil {
		return model.Project{}, err
	}

	p, err := s.projectRepo.Create(p)
	if err != nil {
		return model.Project{}, err
	}

	_, err = s.columnRepo.Create(model.Column{Name: "default", Project: p})
	if err != nil {
		return model.Project{}, err
	}

	return p, nil
}

// GetByID returns project with specific ID.
func (s *ProjectService) GetByID(id int) (model.Project, error) {
	return s.projectRepo.GetByID(id)
}

// Update updates a project.
func (s *ProjectService) Update(p model.Project) error {
	return s.projectRepo.Update(p)
}

// Delete deletes a project.
func (s *ProjectService) Delete(p model.Project) error {
	return s.projectRepo.Delete(p)
}

// Validate validates a project.
func (s *ProjectService) Validate(p model.Project) error {
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
