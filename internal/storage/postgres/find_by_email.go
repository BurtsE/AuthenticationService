package postgres

import (
	"AuthenticationService/internal/domain"
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
)

func (d *DB) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	const query = `
		SELECT id, email, password_hash, email_verified, created_at
		FROM users
		WHERE email = $1
	`

	var user domain.User
	var id uuid.UUID
	err := d.pool.QueryRow(ctx, query, email).Scan(
		id,
		&user.Email,
		&user.PasswordHash,
		&user.EmailVerified,
		&user.CreatedAt,
	)
	user.ID = domain.UserID(id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
