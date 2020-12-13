package store

import "github.com/imarrche/tasker/internal/model"

//go:generate mockgen -source=interface.go -destination=mocks/mock.go

// Store is an interface all stores must implement.
type Store interface {
	Open() error
	Projects() ProjectRepository
	Columns() ColumnRepository
	Tasks() TaskRepository
	Comments() CommentRepository
	Close() error
}

// ProjectRepository is an interface all project repositories must implement.
type ProjectRepository interface {
	GetAll() ([]model.Project, error)
	Create(model.Project) (model.Project, error)
	GetByID(int) (model.Project, error)
	Update(model.Project) error
	DeleteByID(int) error
}

// ColumnRepository is an interface all column repositories must implement.
type ColumnRepository interface {
	GetAll() ([]model.Column, error)
	Create(model.Column) (model.Column, error)
	GetByID(int) (model.Column, error)
	Update(model.Column) error
	DeleteByID(int) error
}

// TaskRepository is an interface all task repositories must implement.
type TaskRepository interface {
	GetAll() ([]model.Task, error)
	Create(model.Task) (model.Task, error)
	GetByID(int) (model.Task, error)
	Update(model.Task) error
	DeleteByID(int) error
}

// CommentRepository is an interface all comment repositories must implement.
type CommentRepository interface {
	GetAll() ([]model.Comment, error)
	Create(model.Comment) (model.Comment, error)
	GetByID(int) (model.Comment, error)
	Update(model.Comment) error
	DeleteByID(int) error
}
