package redis

import (
	"AuthenticationService/internal/domain"
	"context"
	"github.com/google/uuid"
)

func (s *SessionStorage) DeleteAllUserSessions(ctx context.Context, userID domain.UserID) error {
	userKey := userKeyPrefix + uuid.UUID(userID).String()

	refreshTokenIDs, err := s.rdb.SMembers(ctx, userKey).Result()
	if err != nil {
		return err
	}

	if len(refreshTokenIDs) == 0 {
		return nil
	}

	pipe := s.rdb.TxPipeline()

	for _, tokenID := range refreshTokenIDs {
		pipe.Del(ctx, sessionKeyPrefix+tokenID)
	}

	pipe.Del(ctx, userKey)

	_, err = pipe.Exec(ctx)
	return err
}
