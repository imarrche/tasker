package store

import "github.com/imarrche/tasker/internal/model"

// ProjectRepository is an interface all project repositories must implement.
type ProjectRepository interface {
	GetAll() ([]model.Project, error)
	Create(model.Project) (model.Project, error)
	GetByID(int) (model.Project, error)
	Update(model.Project) error
	Delete(model.Project) error
}
