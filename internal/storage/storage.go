package storage

import (
	"AuthenticationService/internal/domain"
	"context"
)

type UserStorage interface {
	CreateUser(context.Context, *domain.User) error
	DeleteUser(context.Context, domain.UserID) error
	FindByID(context.Context, domain.UserID) (bool, error)
	FindByEmail(context.Context, string) (*domain.User, error)
}
