package service

import (
	"AuthenticationService/internal/model"
	"context"
)

type IUserService interface {
	CreateUser(context.Context, *model.User) error
	DeleteUser(context.Context, model.UserID) error
}
