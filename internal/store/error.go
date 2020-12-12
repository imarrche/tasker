package store

import "errors"

var (
	// ErrNotFound is thrown when specified record was not found in a store.
	ErrNotFound = errors.New("not found")
)
