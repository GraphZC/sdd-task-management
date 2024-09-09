package exceptions

import "errors"

var (
	ErrInvalidPriority = errors.New("invalid priority")
	ErrInvalidStatus   = errors.New("invalid status")
	ErrTaskNotFound    = errors.New("task not found")
)
