package service

import "github.com/imarrche/tasker/internal/model"

// Service is an interface all services must implement.
type Service interface {
	Projects() ProjectService
}

// ProjectService is an interface all project services must implement.
type ProjectService interface {
	GetAll() ([]model.Project, error)
	Create(model.Project) (model.Project, error)
	GetByID(int) (model.Project, error)
	Update(model.Project) error
	DeleteByID(int) error
	Validate(model.Project) error
}
