package domain

import (
	"github.com/google/uuid"
	"time"
)

type UserID uuid.UUID
type User struct {
	ID            UserID
	Email         string
	PasswordHash  string
	EmailVerified bool
	CreatedAt     time.Time
}
