package storage

import (
	"AuthenticationService/internal/domain"
	"context"
	"github.com/google/uuid"
)

type UserStorage interface {
	CreateUser(context.Context, *domain.User) error
	DeleteUser(context.Context, domain.UserID) error
	FindByEmail(context.Context, string) (*domain.User, error)
}

type SessionStorage interface {
	GetSessionByTokenID(context.Context, uuid.UUID) (*domain.Session, error)
	CreateSession(context.Context, *domain.Session) error
	DeleteSession(context.Context, *domain.Session) error
	UserSessionsCount(context.Context, domain.UserID) (int, error)
	DeleteAllUserSessions(context.Context, domain.UserID) error
}
