package costformation

import (
	"errors"
)

var (
	ErrUnableToRead      = errors.New("unable to read from file")
	ErrUnableToWrite     = errors.New("unable to write to destination")
	ErrInvalidDefinition = errors.New("invalid or nil Definition")
)
