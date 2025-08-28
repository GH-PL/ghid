package errHandler

import "errors"

// Error
var (
	ErrEmptyArgument = errors.New("Error: missing required argument <hash name>")
)
