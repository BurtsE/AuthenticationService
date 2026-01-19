package user

import (
	"AuthenticationService/internal/domain"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func (s *Service) RefreshToken(
	ctx context.Context,
	refreshToken,
	fingerprint string,
) (string, string, error) {
	claims, err := s.tokenManager.Validate(refreshToken)
	if err != nil {
		return "", "", fmt.Errorf("%w: %s", domain.ErrInvalidCredentials, err)
	}

	user, err := s.userStorage.FindByEmail(ctx, claims.Email)
	if err != nil {
		return "", "", fmt.Errorf("%w: %s", domain.ErrDatabaseConflict, err)
	}
	if user == nil {
		return "", "", domain.ErrEntityNotFound
	}

	tokenID, err := uuid.Parse(claims.ID)
	if err != nil {
		return "", "", fmt.Errorf("%w: %s", domain.ErrInvalidToken, err)
	}

	session, err := s.sessionStorage.GetSessionByTokenID(ctx, tokenID)
	if err != nil {
		return "", "", fmt.Errorf("%w: %s", domain.ErrDatabaseConflict, err)
	}
	if session == nil {
		return "", "", domain.ErrEntityNotFound
	}

	err = s.sessionStorage.DeleteSession(ctx, session)
	if err != nil {
		return "", "", fmt.Errorf("%w: %s", domain.ErrDatabaseConflict, err)
	}

	if session.Fingerprint != fingerprint || session.IsExpired() {
		return "", "", fmt.Errorf("%w: %s: %v; %s;%s", domain.ErrInvalidRefreshSession,
			"session is expired", session.IsExpired(), session.Fingerprint, fingerprint)
	}

	accessToken, refreshToken, err := s.tokenManager.GenerateTokenPair(user.ID.String(), user.Email)
	if err != nil {
		return "", "", err
	}

	newClaims, err := s.tokenManager.Validate(refreshToken)
	if err != nil {
		return "", "", err
	}

	now := time.Now()
	newSession := &domain.Session{
		UserID:         user.ID,
		RefreshTokenID: newClaims.ID,
		Fingerprint:    fingerprint,
		CreatedAt:      now,
		ExpiresAt:      now.Add(domain.SessionDuration),
	}

	err = s.sessionStorage.CreateSession(ctx, newSession)
	if err != nil {
		return "", "", domain.ErrDatabaseConflict
	}

	return accessToken, refreshToken, nil

}
