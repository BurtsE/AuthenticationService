package user

import (
	"AuthenticationService/internal/domain"
	"AuthenticationService/internal/dto"
	"context"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) AuthorizeUser(ctx context.Context, dto dto.AuthorizeUserRequest) (string, string, error) {
	user, err := s.db.FindByEmail(ctx, dto.Email)
	if err != nil {
		return "", "", err
	}

	if user == nil {
		return "", "", domain.ErrInvalidCredentials
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(dto.Password)); err != nil {
		return "", "", domain.ErrInvalidCredentials
	}

	return s.tokenManager.GenerateTokenPair(user.ID.String(), user.Email)
}
