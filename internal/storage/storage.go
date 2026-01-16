package storage

import (
	"AuthenticationService/internal/model"
	"context"
)

type UserStorage interface {
	CreateUser(context.Context, *model.User) error
	DeleteUser(context.Context, model.UserID) error
}
