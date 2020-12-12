package store

import "github.com/imarrche/tasker/internal/model"

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
	Delete(model.Project) error
}

// ColumnRepository is an interface all column repositories must implement.
type ColumnRepository interface {
	GetAll() ([]model.Column, error)
	Create(model.Project) (model.Column, error)
	GetByID(int) (model.Column, error)
	Update(model.Column) error
	Delete(model.Column) error
}

// TaskRepository is an interface all task repositories must implement.
type TaskRepository interface {
	GetAll() ([]model.Task, error)
	Create(model.Task) (model.Task, error)
	GetByID(int) (model.Task, error)
	Update(model.Task) error
	Delete(model.Task) error
}

// CommentRepository is an interface all comment repositories must implement.
type CommentRepository interface {
	GetAll() ([]model.Comment, error)
	Create(model.Comment) (model.Comment, error)
	GetByID(int) (model.Comment, error)
	Update(model.Comment) error
	Delete(model.Comment) error
}
