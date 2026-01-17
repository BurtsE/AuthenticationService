package domain

import (
	"github.com/google/uuid"
	"time"
)

type UserID uuid.UUID

func (u UserID) String() string {
	return uuid.UUID(u).String()
}

type User struct {
	ID            UserID
	Email         string
	PasswordHash  string
	EmailVerified bool
	CreatedAt     time.Time
}
