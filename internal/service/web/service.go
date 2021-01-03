package web

import (
	"github.com/imarrche/tasker/internal/service"
	"github.com/imarrche/tasker/internal/store"
)

// Service is the web service.
type Service struct {
	store    store.Store
	projects *projectService
	columns  *columnService
	tasks    *taskService
	comments *commentService
}

// NewService creates and returns a new Service instance.
func NewService(s store.Store) *Service { return &Service{store: s} }

// Projects returns the project service.
func (s *Service) Projects() service.ProjectService {
	if s.projects == nil {
		s.projects = newProjectService(s.store)
	}

	return s.projects
}

// Columns returns the column service.
func (s *Service) Columns() service.ColumnService {
	if s.columns == nil {
		s.columns = newColumnService(s.store)
	}

	return s.columns
}

// Tasks returns the task service.
func (s *Service) Tasks() service.TaskService {
	if s.tasks == nil {
		s.tasks = newTaskService(s.store)
	}

	return s.tasks
}

// Comments returns the comment service.
func (s *Service) Comments() service.CommentService {
	if s.comments == nil {
		s.comments = newCommentService(s.store)
	}

	return s.comments
}
