package domain

import "errors"

// Service errors
var (
	ErrUserAlreadyExists     = errors.New("user already exists")
	ErrEntityNotFound        = errors.New("entity not found")
	ErrInvalidEmail          = errors.New("invalid email")
	ErrWeakPassword          = errors.New("weak password")
	ErrInvalidCredentials    = errors.New("invalid credentials")
	ErrInvalidRefreshSession = errors.New("invalid refresh session")
	ErrDatabaseConflict      = errors.New("database error")
)

// Token errors
var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token expired")
)
