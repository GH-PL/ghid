package errHandler

import "errors"

// Error
var (
	ErrEmptyArgument   = errors.New("Error: missing required argument <hash name>")
	ErrNotFoundHash    = errors.New("Not found type for this Hash")
	ErrNotExampleFound = errors.New("No examples found")
)
