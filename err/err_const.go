package err

import "errors"

var (
	ErrMissingArgument = errors.New("Missing required argument")
	ErrInvalidFormat   = errors.New("Invalid format")
)
