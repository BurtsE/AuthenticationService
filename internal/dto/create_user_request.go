package dto

import (
	"AuthenticationService/internal/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *CreateUserRequest) ToEntity() domain.User {
	user := domain.User{
		ID:            domain.UserID(uuid.New()),
		Email:         r.Email,
		EmailVerified: false,
		CreatedAt:     time.Now(),
	}
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.MinCost)
	user.PasswordHash = string(passwordHash)
	return user
}
