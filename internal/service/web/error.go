package web

import "errors"

var (
	// ErrNameIsRequired is thrown when name field was not provided.
	ErrNameIsRequired = errors.New("name is required")
	// ErrNameIsTooLong is thrown when name field is too long.
	ErrNameIsTooLong = errors.New("name is too long")
	// ErrDescriptionIsTooLong is thrown when description field is too long.
	ErrDescriptionIsTooLong = errors.New("description is too long")
)
