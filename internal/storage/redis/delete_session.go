package redis

import (
	"AuthenticationService/internal/domain"
	"context"
	"github.com/google/uuid"
)

func (s *SessionStorage) DeleteSession(ctx context.Context, session *domain.Session) error {
	sessionKey := sessionKeyPrefix + session.RefreshTokenID
	userKey := userKeyPrefix + uuid.UUID(session.UserID).String()

	pipe := s.rdb.TxPipeline()
	pipe.Del(ctx, sessionKey)
	pipe.SRem(ctx, userKey, session.RefreshTokenID)
	_, err := pipe.Exec(ctx)

	return err
}
