package main

import (
	"AuthenticationService/internal/auth"
	"AuthenticationService/internal/config"
	"AuthenticationService/internal/handlers"
	"AuthenticationService/internal/service/user"
	"AuthenticationService/internal/storage/postgres"
	redisInternal "AuthenticationService/internal/storage/redis"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/redis/go-redis/v9"
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

	//configure redis storage
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.GetRedisUrl(),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if err = rdb.Ping(context.Background()).Err(); err != nil {
		logger.Fatal(err)
	}
	sessionStorage := redisInternal.NewSessionStorage(rdb)

	// Configure  service, handlers
	service := user.NewUserService(postgresDb, sessionStorage, tokenManager)
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
	_ = sessionStorage.Close()

	logger.Info("Shutting down...")

}
