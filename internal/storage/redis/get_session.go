package redis

import (
	"AuthenticationService/internal/domain"
	"context"
	"github.com/google/uuid"
)

func (s *SessionStorage) GetSessionByTokenID(ctx context.Context, tokenID uuid.UUID) (*domain.Session, error) {
	key := sessionKeyPrefix + tokenID.String()

	session := &domain.Session{}
	err := s.rdb.HGetAll(ctx, key).Scan(session)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *SessionStorage) UserSessionsCount(ctx context.Context, id domain.UserID) (int, error) {
	key := userKeyPrefix + id.String()

	refreshTokenIDs, err := s.rdb.SMembers(ctx, key).Result()
	if err != nil {
		return -1, err
	}

	return len(refreshTokenIDs), nil
}
