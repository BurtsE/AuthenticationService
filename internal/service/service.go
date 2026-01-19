package service

import (
	"AuthenticationService/internal/domain"
	"AuthenticationService/internal/dto"
	"context"
)

type IUserService interface {
	CreateUser(context.Context, dto.CreateUserRequest) error
	DeleteUser(context.Context, domain.UserID) error

	AuthenticateUser(ctx context.Context, dto dto.AuthorizeUserRequest, fingerprint string) (string, string, error)
	RefreshToken(ctx context.Context, refreshToken, fingerprint string) (string, string, error)
}
