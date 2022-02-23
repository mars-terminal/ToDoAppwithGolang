package entities

import (
	"errors"
)

var (
	ErrInternal     = errors.New("internal error")
	ErrForbidden    = errors.New("access denied")
	ErrBadRequest   = errors.New("bad request")
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
)
