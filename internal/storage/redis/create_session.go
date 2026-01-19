package redis

import (
	"AuthenticationService/internal/domain"
	"context"
	"github.com/google/uuid"
	"time"
)

func (s *SessionStorage) CreateSession(ctx context.Context, session *domain.Session) error {
	sessionKey := sessionKeyPrefix + session.RefreshTokenID
	userKey := userKeyPrefix + uuid.UUID(session.UserID).String()

	pipe := s.rdb.TxPipeline()

	pipe.HSet(ctx, sessionKey, *session)

	ttl := time.Until(session.ExpiresAt)
	if ttl > 0 {
		pipe.Expire(ctx, sessionKey, ttl)
	}

	pipe.SAdd(ctx, userKey, session.RefreshTokenID)

	if ttl > 0 {
		pipe.Expire(ctx, userKey, ttl+time.Hour)
	}

	_, err := pipe.Exec(ctx)

	return err
}
