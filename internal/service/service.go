package service

import (
	"AuthenticationService/internal/domain"
	"AuthenticationService/internal/dto"
	"context"
)

type IUserService interface {
	CreateUser(context.Context, dto.CreateUserRequest) error
	DeleteUser(context.Context, domain.UserID) error
}
