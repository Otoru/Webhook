package errors

import "errors"

var (
	ErrInvalidLogLevel = errors.New("invalid-log-level")
	ErrInvalidWorkdir  = errors.New("invalid-workdir")
	ErrUnexpectedError = errors.New("unexpected-error")
)
