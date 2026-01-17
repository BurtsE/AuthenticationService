package postgres

import (
	"AuthenticationService/internal/storage"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

var _ storage.UserStorage = (*DB)(nil)

type DB struct {
	log  *logrus.Logger
	pool *pgxpool.Pool
}

func NewDatabase(
	logger *logrus.Logger,
	pool *pgxpool.Pool,
) *DB {
	return &DB{
		pool: pool,
		log:  logger,
	}
}

func (d *DB) Close() {
	d.pool.Close()
}
