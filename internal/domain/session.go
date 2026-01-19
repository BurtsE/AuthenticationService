package domain

import (
	"AuthenticationService/internal/config"
	"time"
)

var SessionDuration = config.GetSessionDuration()

type Session struct {
	UserID         UserID    `redis:"user_id"`
	RefreshTokenID string    `redis:"refresh_token_id"`
	Fingerprint    string    `redis:"fingerprint"`
	ExpiresAt      time.Time `redis:"expires_at"`
	CreatedAt      time.Time `json:"created_at"`
}

func (s *Session) IsExpired() bool {
	return time.Now().After(s.ExpiresAt)
}
