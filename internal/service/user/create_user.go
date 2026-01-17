package user

import (
	"AuthenticationService/internal/domain"
	"AuthenticationService/internal/dto"
	"AuthenticationService/internal/validation"
	"context"
	"fmt"
)

func (s *Service) CreateUser(ctx context.Context, request dto.CreateUserRequest) error {
	if !validation.ValidEmail(request.Email) {
		return domain.ErrInvalidEmail
	}
	if !validation.ValidPassword(request.Password) {
		return domain.ErrWeakPassword
	}

	existingUser, err := s.db.FindByEmail(ctx, request.Email)
	if err != nil {
		return fmt.Errorf("%w: %s", domain.ErrDatabaseConflict, err)
	}
	if existingUser != nil {
		return domain.ErrUserAlreadyExists
	}

	user := request.ToEntity()

	err = s.db.CreateUser(ctx, &user)
	if err != nil {
		return fmt.Errorf("%w: %s", domain.ErrDatabaseConflict, err)
	}
	return nil
}
