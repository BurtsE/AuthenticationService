package user

import (
	"AuthenticationService/internal/auth"
	"AuthenticationService/internal/service"
	"AuthenticationService/internal/storage"
)

const sessionsLimit = 4

var _ service.IUserService = (*Service)(nil)

type ITokenManager interface {
	GenerateTokenPair(userID string, email string) (string, string, error)
}
type Service struct {
	userStorage    storage.UserStorage
	sessionStorage storage.SessionStorage
	tokenManager   *auth.TokenManager
}

func NewUserService(
	userStorage storage.UserStorage,
	sessionStorage storage.SessionStorage,
	manager *auth.TokenManager) *Service {
	return &Service{
		userStorage:    userStorage,
		sessionStorage: sessionStorage,
		tokenManager:   manager,
	}
}
