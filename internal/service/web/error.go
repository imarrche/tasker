package web

import "errors"

var (
	// ErrNameIsRequired is thrown when name field is not provided.
	ErrNameIsRequired = errors.New("name is required")
	// ErrNameIsTooLong is thrown when name field is too long.
	ErrNameIsTooLong = errors.New("name is too long")
	// ErrDescriptionIsTooLong is thrown when description field is too long.
	ErrDescriptionIsTooLong = errors.New("description is too long")
	// ErrTextIsRequired is thrown when text field is not provided.
	ErrTextIsRequired = errors.New("text is required")
	// ErrTextIsTooLong is thrown when text field is too long.
	ErrTextIsTooLong = errors.New("text is too long")
	// ErrColumnAlreadyExists is thrown when column with provided name already exists.
	ErrColumnAlreadyExists = errors.New("column already exists")
	// ErrLastColumn is thrown when deleting last project's column.
	ErrLastColumn = errors.New("last column cannot be deleted")
)
