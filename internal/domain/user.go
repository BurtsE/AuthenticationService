package domain

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type UserID uuid.UUID

func (u UserID) String() string {
	return uuid.UUID(u).String()
}

func (u UserID) MarshalBinary() ([]byte, error) {
	return uuid.UUID(u).MarshalBinary()
}

func (u *UserID) UnmarshalBinary(data []byte) error {
	var id uuid.UUID
	if err := id.UnmarshalBinary(data); err != nil {
		return err
	}
	*u = UserID(id)
	return nil
}

func (u *UserID) UnmarshalText(text []byte) error {
	id, err := uuid.FromBytes(text)
	if err != nil {
		return fmt.Errorf("invalid UUID: %w", err)
	}
	*u = UserID(id)
	return nil
}

type User struct {
	ID            UserID
	Email         string
	PasswordHash  string
	EmailVerified bool
	CreatedAt     time.Time
}
