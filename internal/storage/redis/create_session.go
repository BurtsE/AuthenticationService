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

	//values := map[string]interface{}{
	//	"user_id":          uuid.UUID(session.UserID).String(),
	//	"refresh_token_id": session.RefreshTokenID,
	//	"fingerprint":      session.Fingerprint,
	//	"expires_at":       session.ExpiresAt.Unix(),
	//	"created_at":       session.CreatedAt.Unix(),
	//}

	pipe := s.rdb.TxPipeline()

	//pipe.HSet(ctx, sessionKey, values)
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
