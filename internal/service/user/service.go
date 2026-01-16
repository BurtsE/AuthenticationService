package user

import (
	"AuthenticationService/internal/service"
	"AuthenticationService/internal/storage"
)

var _ service.IUserService = (*Service)(nil)

type Service struct {
	db storage.UserStorage
}

func NewUserService(db storage.UserStorage) *Service {
	return &Service{db: db}
}
