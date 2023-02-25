package costformation

import (
	"errors"
)

var (
	ErrInvalidReader     = errors.New("invalid or nil reader source")
	ErrInvalidWriter     = errors.New("invalid or nil writer destination")
	ErrUnableToRead      = errors.New("unable to read from file")
	ErrUnableToWrite     = errors.New("unable to write to destination")
	ErrInvalidDefinition = errors.New("invalid or nil Definition")
)
