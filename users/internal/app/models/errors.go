package models

import "errors"

var (
	ErrNotFound         = errors.New("not found")
	ErrAlreadyExists    = errors.New("already exists")
	ErrUnauthenticated  = errors.New("unauthenticated")
	ErrPermissionDenied = errors.New("permission denied")
)
