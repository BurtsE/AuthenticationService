package redis

import (
	"AuthenticationService/internal/storage"
	"github.com/redis/go-redis/v9"
)

var _ storage.SessionStorage = (*SessionStorage)(nil)

const (
	sessionKeyPrefix = "session:"
	userKeyPrefix    = "user_sessions:"
)

// Storage logic
// session:{refresh_token_id} -> HASH
// user_sessions:{user_id} -> SET(refresh_token_id)

type SessionStorage struct {
	rdb *redis.Client
}

func NewSessionStorage(rdb *redis.Client) *SessionStorage {
	return &SessionStorage{rdb: rdb}
}

func (s *SessionStorage) Close() error {
	return s.rdb.Close()
}
