package user

import (
	"AuthenticationService/internal/service"
	"AuthenticationService/internal/storage"
)

var _ service.IUserService = (*Service)(nil)

type ITokenManager interface {
	GenerateTokenPair(userID string, email string) (string, string, error)
}
type Service struct {
	db           storage.UserStorage
	tokenManager ITokenManager
}

func NewUserService(db storage.UserStorage, manager ITokenManager) *Service {
	return &Service{
		db:           db,
		tokenManager: manager,
	}
}
