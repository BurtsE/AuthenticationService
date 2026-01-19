package domain

import (
	"AuthenticationService/internal/config"
	"time"
)

var SessionDuration time.Duration = config.GetSessionDuration()

type Session struct {
	UserID         UserID
	RefreshTokenID string
	Fingerprint    string
	ExpiresAt      time.Time
	CreatedAt      time.Time
}

func (s *Session) IsExpired() bool {
	return time.Now().After(s.ExpiresAt)
}
