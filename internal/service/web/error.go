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
	ErrLastColumn = errors.New("last column can't be deleted")
	// ErrInvalidMove is thrown when model is moved to invalid position.
	ErrInvalidMove = errors.New("move can't be performed")
)

// IsValidationError checks whether error is validation related.
func IsValidationError(err error) bool {
	switch err {
	case ErrNameIsRequired, ErrNameIsTooLong:
		return true
	case ErrDescriptionIsTooLong:
		return true
	case ErrTextIsRequired, ErrTextIsTooLong:
		return true
	case ErrColumnAlreadyExists:
		return true
	default:
		return false
	}
}
