package postgres

import (
	"AuthenticationService/internal/model"
	"context"
)

func (d *DB) DeleteUser(ctx context.Context, id model.UserID) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`
	_, err := d.pool.Exec(ctx, query, id)
	return err
}
