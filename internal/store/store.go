package store

// Store is an interface all stores must implement.
type Store interface {
	Open() error
	Projects() ProjectRepository
	Columns() ColumnRepository
	Tasks() TaskRepository
	Comments() CommentRepository
	Close() error
}
