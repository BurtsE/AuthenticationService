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

	//if len(data) == 0 {
	//	return nil, nil
	//}
	//
	//userUUID, err := uuid.Parse(data["user_id"])
	//if err != nil {
	//	return nil, err
	//}
	//
	//expiresAtUnix, err := strconv.ParseInt(data["expires_at"], 10, 64)
	//if err != nil {
	//	return nil, err
	//}
	//
	//createdAtUnix, err := strconv.ParseInt(data["created_at"], 10, 64)
	//if err != nil {
	//	return nil, err
	//}
	//
	//session := &domain.Session{
	//	UserID:         domain.UserID(userUUID),
	//	RefreshTokenID: data["refresh_token_id"],
	//	Fingerprint:    data["fingerprint"],
	//	ExpiresAt:      time.Unix(expiresAtUnix, 0),
	//	CreatedAt:      time.Unix(createdAtUnix, 0),
	//}

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
