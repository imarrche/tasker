package web

import (
	"github.com/imarrche/tasker/internal/service"
	"github.com/imarrche/tasker/internal/store"
)

// Service is a web service.
type Service struct {
	store    store.Store
	projects service.ProjectService
	columns  service.ColumnService
	tasks    service.TaskService
}

// NewService creates and returns a new Service instance.
func NewService(s store.Store) *Service {
	return &Service{store: s}
}

// Projects returns project service.
func (s *Service) Projects() service.ProjectService {
	if s.projects == nil {
		s.projects = NewProjectService(s.store.Projects(), s.store.Columns())
	}

	return s.projects
}

// Columns returns column service.
func (s *Service) Columns() service.ColumnService {
	if s.columns == nil {
		s.columns = NewColumnService(s.store.Columns(), s.store.Tasks())
	}

	return s.columns
}

// Tasks returns task service.
func (s *Service) Tasks() service.TaskService {
	if s.tasks == nil {
		s.tasks = NewTaskService(s.store.Columns(), s.store.Tasks())
	}

	return s.tasks
}
