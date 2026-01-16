package postgres

import (
	"AuthenticationService/internal/domain"
	"context"
)

func (d *DB) DeleteUser(ctx context.Context, id domain.UserID) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`
	_, err := d.pool.Exec(ctx, query, id)
	return err
}
