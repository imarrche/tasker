package web

import "errors"

var (
	// ErrNameIsRequired is thrown when name field is not provided.
	ErrNameIsRequired = errors.New("name is required")
	// ErrNameIsTooLong is thrown when name field is too long.
	ErrNameIsTooLong = errors.New("name is too long")
	// ErrDescriptionIsTooLong is thrown when description field is too long.
	ErrDescriptionIsTooLong = errors.New("description is too long")
	// ErrTextIsRequired is thrown when comment field is not provided.
	ErrTextIsRequired = errors.New("text is required")
	// ErrTextIsTooLong is thrown when comment's text is too long.
	ErrTextIsTooLong = errors.New("text is too long")
	// ErrProjectIsRequired is thrown when project field is not provided.
	ErrProjectIsRequired = errors.New("project is required")
	// ErrColumnAlreadyExists is thrown when column with provided name already ixists.
	ErrColumnAlreadyExists = errors.New("column already ixists")
	// ErrLastColumn is thrown when user is trying to delete last column in the project.
	ErrLastColumn = errors.New("last column cannot be deleted")
	// ErrInvalidColumn is thrown when task has invalid column.
	ErrInvalidColumn = errors.New("invalid column")
	// ErrInvalidTask is thrown when comment has invalid task.
	ErrInvalidTask = errors.New("invalid task")
)
