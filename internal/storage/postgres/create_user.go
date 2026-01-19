package postgres

import (
	"AuthenticationService/internal/domain"
	"context"
)

func (d *UserStorage) CreateUser(ctx context.Context, user *domain.User) error {
	const query = `
		INSERT INTO users(id, email, password_hash, email_verified, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	_, err := d.pool.Exec(ctx, query,
		user.ID,
		user.Email,
		user.PasswordHash,
		user.EmailVerified,
		user.CreatedAt,
	)

	return err
}
