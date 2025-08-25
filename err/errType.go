package err

import "errors"

// Error
var (
	ErrEmptyArgument   = errors.New("Error: missing required argument: hash name")
	ErrMissingArgument = errors.New("Missing required argument") 
	ErrInvalidFormat   = errors.New("Invalid format")           
)

// Not found
var (
	ErrNoSamplesFound = errors.New("No examples found")
)
