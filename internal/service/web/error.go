package web

import "errors"

var (
	// ErrNameIsRequired is thrown when name field is not provided.
	ErrNameIsRequired = errors.New("name is required")
	// ErrNameIsTooLong is thrown when name field is too long.
	ErrNameIsTooLong = errors.New("name is too long")
	// ErrDescriptionIsTooLong is thrown when description field is too long.
	ErrDescriptionIsTooLong = errors.New("description is too long")
	// ErrProjectIsRequired is thrown when proejct field is not provided.
	ErrProjectIsRequired = errors.New("project is required")
	// ErrColumnAlreadyExists is thrown when column with provided name already ixists.
	ErrColumnAlreadyExists = errors.New("column already ixists")
	// ErrLastColumn is thrown when user is trying to delete last column in the project.
	ErrLastColumn = errors.New("last column cannot be deleted")
)
