package domain

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidEmail      = errors.New("invalid email")
	ErrWeakPassword      = errors.New("weak password")
)
