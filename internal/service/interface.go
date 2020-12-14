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

// ColumnService is an interface all column services must implement.
type ColumnService interface {
	GetAll() ([]model.Column, error)
	Create(model.Column) (model.Column, error)
	GetByID(int) (model.Column, error)
	Update(model.Column) error
	DeleteByID(int) error
	Validate(model.Column) error
}

// TaskService is an interface all task services must implement.
type TaskService interface {
	GetAll() ([]model.Task, error)
	Create(model.Task) (model.Task, error)
	GetByID(int) (model.Task, error)
	Update(model.Task) error
	DeleteByID(int) error
	Validate(model.Task) error
}
