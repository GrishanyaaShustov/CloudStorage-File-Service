package domain

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrConflict      = errors.New("conflict")
	ErrInvalidParent = errors.New("invalid parent")
	ErrInvalidState  = errors.New("invalid state")
)
