package postgres

import (
	"AuthenticationService/internal/config"
	"AuthenticationService/internal/storage"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

var _ storage.UserStorage = (*DB)(nil)

type DB struct {
	log  *logrus.Logger
	pool *pgxpool.Pool
}

func NewDatabase(logger *logrus.Logger) *DB {
	cfg, err := pgxpool.ParseConfig(config.GetPostgresUrl())
	if err != nil {
		logger.Fatal(err)
	}

	cfg.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		logger.Info("connected to database")
		return nil
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		logger.Fatal(err)
	}
	return &DB{
		pool: pool,
		log:  logger,
	}
}

func (d *DB) Close() {
	d.pool.Close()
}
