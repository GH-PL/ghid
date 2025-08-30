package errHandler

import "errors"

// Error
var (
	ErrEmptyArgument   = errors.New("Error: missing required argument <hash name>")
	ErrNotFoundHash    = errors.New("Error: not found type for this Hash")
	ErrNotExampleFound = errors.New("Error: no examples found")
	ErrNotFoundName    = errors.New("Error: not found this Hash for name")
)
