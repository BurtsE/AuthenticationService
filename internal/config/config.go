package config

import (
	"fmt"
	"os"
)

func GetApplicationPort() string {
	return getEnv("APPLICATION_PORT", ":8080")
}

func GetApplicationEnvironment() string {
	return getEnv("APPLICATION_ENVIRONMENT", "development")
}

func GetPostgresUrl() string {
	host := getEnv("POSTGRES_HOST", "localhost")
	port := getEnv("POSTGRES_PORT", "5432")
	user := getEnv("POSTGRES_USER", "postgres")
	password := getEnv("POSTGRES_PASSWORD", "postgres")
	database := getEnv("POSTGRES_DB", "postgres")
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
