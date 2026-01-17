package main

import (
	"AuthenticationService/internal/auth"
	"AuthenticationService/internal/config"
	"AuthenticationService/internal/handlers"
	"AuthenticationService/internal/service/user"
	"AuthenticationService/internal/storage/postgres"
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

	postgresDb := postgres.NewDatabase(logger)
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
