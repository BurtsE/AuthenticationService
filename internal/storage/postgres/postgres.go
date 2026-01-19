package postgres

import (
	"AuthenticationService/internal/storage"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

var _ storage.UserStorage = (*UserStorage)(nil)

type UserStorage struct {
	log  *logrus.Logger
	pool *pgxpool.Pool
}

func NewDatabase(
	logger *logrus.Logger,
	pool *pgxpool.Pool,
) *UserStorage {
	return &UserStorage{
		pool: pool,
		log:  logger,
	}
}

func (d *UserStorage) Close() {
	d.pool.Close()
}
