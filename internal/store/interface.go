package store

import "github.com/imarrche/tasker/internal/model"

//go:generate mockgen -source=interface.go -destination=mocks/mock.go

// Store is the interface all stores must implement.
type Store interface {
	Open() error
	Projects() ProjectRepo
	Columns() ColumnRepo
	Tasks() TaskRepo
	Comments() CommentRepo
	Close() error
}

// ProjectRepo is the interface all project repositories must implement.
type ProjectRepo interface {
	GetAll() ([]model.Project, error)
	Create(model.Project) (model.Project, error)
	GetByID(int) (model.Project, error)
	Update(model.Project) error
	DeleteByID(int) error
}

// ColumnRepo is the interface all column repositories must implement.
type ColumnRepo interface {
	GetByProjectID(int) ([]model.Column, error)
	Create(model.Column) (model.Column, error)
	GetByID(int) (model.Column, error)
	GetByIndexAndProjectID(int, int) (model.Column, error)
	Update(model.Column) error
	DeleteByID(int) error
}

// TaskRepo is the interface all task repositories must implement.
type TaskRepo interface {
	GetByColumnID(int) ([]model.Task, error)
	Create(model.Task) (model.Task, error)
	GetByID(int) (model.Task, error)
	GetByIndexAndColumnID(int, int) (model.Task, error)
	Update(model.Task) error
	DeleteByID(int) error
}

// CommentRepo is the interface all comment repositories must implement.
type CommentRepo interface {
	GetByTaskID(int) ([]model.Comment, error)
	Create(model.Comment) (model.Comment, error)
	GetByID(int) (model.Comment, error)
	Update(model.Comment) error
	DeleteByID(int) error
}
