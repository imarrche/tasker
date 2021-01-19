package service

import "github.com/imarrche/tasker/internal/model"

//go:generate mockgen -source=interface.go -destination=mocks/mock.go

// Service is the interface all services must implement.
type Service interface {
	Projects() ProjectService
	Columns() ColumnService
	Tasks() TaskService
	Comments() CommentService
}

// ProjectService is the interface all project services must implement.
type ProjectService interface {
	GetAll() ([]model.Project, error)
	Create(model.Project) (model.Project, error)
	GetByID(int) (model.Project, error)
	Update(model.Project) (model.Project, error)
	DeleteByID(int) error
	Validate(model.Project) error
}

// ColumnService is the interface all column services must implement.
type ColumnService interface {
	GetByProjectID(int) ([]model.Column, error)
	Create(model.Column) (model.Column, error)
	GetByID(int) (model.Column, error)
	Update(model.Column) (model.Column, error)
	MoveByID(int, bool) error
	DeleteByID(int) error
	Validate(model.Column) error
}

// TaskService is the interface all task services must implement.
type TaskService interface {
	GetByColumnID(int) ([]model.Task, error)
	Create(model.Task) (model.Task, error)
	GetByID(int) (model.Task, error)
	Update(model.Task) (model.Task, error)
	MoveToColumnByID(int, bool) error
	MoveByID(int, bool) error
	DeleteByID(int) error
	Validate(model.Task) error
}

// CommentService is the interface all comment services must implement.
type CommentService interface {
	GetByTaskID(int) ([]model.Comment, error)
	Create(model.Comment) (model.Comment, error)
	GetByID(int) (model.Comment, error)
	Update(model.Comment) (model.Comment, error)
	DeleteByID(int) error
	Validate(model.Comment) error
}
