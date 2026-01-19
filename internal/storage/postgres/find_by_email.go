package postgres

import (
	"AuthenticationService/internal/domain"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

func (d *UserStorage) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	const query = `
		SELECT id, email, password_hash, email_verified, created_at
		FROM users
		WHERE email = $1
	`

	var user domain.User
	var id uuid.UUID
	err := d.pool.QueryRow(ctx, query, email).Scan(
		&id,
		&user.Email,
		&user.PasswordHash,
		&user.EmailVerified,
		&user.CreatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	user.ID = domain.UserID(id)

	return &user, nil
}
