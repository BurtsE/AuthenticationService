package main

import (
	"AuthenticationService/internal/auth"
	"AuthenticationService/internal/config"
	"AuthenticationService/internal/handlers"
	"AuthenticationService/internal/service/user"
	"AuthenticationService/internal/storage/postgres"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := logrus.New()
	if config.GetApplicationEnvironment() == "development" {
		logger.SetLevel(logrus.DebugLevel)
	}

	logFile, err := os.OpenFile("/root/logs/logFile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Warn(err)
	} else {
		logger.SetOutput(logFile)
	}
	defer logFile.Close()

	privateKey, err := config.ReadPrivateKey()
	if err != nil {
		logger.Fatal(err)
	}
	publicKey, err := config.ReadPublicKey()
	if err != nil {
		logger.Fatal(err)
	}
	tokenManager := auth.NewTokenManager(logger, privateKey, publicKey)

	// Configure PostgreSQL Database
	cfg, err := pgxpool.ParseConfig(config.GetPostgresUrl())
	if err != nil {
		logger.Fatal(err)
	}
	cfg.ConnConfig.Logger = logrusadapter.NewLogger(logger)
	cfg.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		logger.Info("connected to database")
		return nil
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		logger.Fatal(err)
	}
	postgresDb := postgres.NewDatabase(logger, pool)

	// Configure  service, handlers
	service := user.NewUserService(postgresDb, tokenManager)
	userHandler := handlers.NewUserHandler(logger, service, tokenManager)

	// Channel for kill signal
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := userHandler.Start(); err != nil {
			logger.Errorf("Failed to start server: %v", err)
			sigchan <- syscall.SIGTERM
		}
	}()
	<-sigchan

	// Clearing resources
	logger.Info("Closing database connections")
	postgresDb.Close()

	logger.Info("Shutting down...")

}
