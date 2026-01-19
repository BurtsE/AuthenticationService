package user

import (
	"AuthenticationService/internal/domain"
	"AuthenticationService/internal/dto"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (s *Service) AuthenticateUser(ctx context.Context, dto dto.AuthorizeUserRequest, fingerprint string,
) (string, string, error) {
	user, err := s.userStorage.FindByEmail(ctx, dto.Email)
	if err != nil {
		return "", "", err
	}
	if user == nil {
		return "", "", domain.ErrInvalidCredentials
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(dto.Password)); err != nil {
		return "", "", fmt.Errorf("%w: %s", domain.ErrInvalidCredentials, err)
	}

	sessionCount, err := s.sessionStorage.UserSessionsCount(ctx, user.ID)
	if err != nil {
		return "", "", fmt.Errorf("%w: %s", domain.ErrDatabaseConflict, err)
	}

	if sessionCount > sessionsLimit {
		err = s.sessionStorage.DeleteAllUserSessions(ctx, user.ID)
		if err != nil {
			return "", "", fmt.Errorf("%w: %s", domain.ErrDatabaseConflict, err)
		}
	}

	accessToken, refreshToken, err := s.tokenManager.GenerateTokenPair(user.ID.String(), user.Email)
	if err != nil {
		return "", "", err
	}

	claims, err := s.tokenManager.Validate(refreshToken)
	if err != nil {
		return "", "", err
	}

	now := time.Now()
	newSession := &domain.Session{
		UserID:         user.ID,
		RefreshTokenID: claims.ID,
		Fingerprint:    fingerprint,
		CreatedAt:      now,
		ExpiresAt:      now.Add(domain.SessionDuration),
	}

	err = s.sessionStorage.CreateSession(ctx, newSession)
	if err != nil {
		return "", "", fmt.Errorf("%w: %s", domain.ErrDatabaseConflict, err)
	}

	return accessToken, refreshToken, nil
}
