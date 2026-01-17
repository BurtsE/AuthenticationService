package config

import "fmt"

func GetPostgresUrl() string {
	host := getEnv("POSTGRES_HOST", "localhost")
	port := getEnv("POSTGRES_PORT", "5432")
	user := getEnv("POSTGRES_USER", "postgres")
	password := getEnv("POSTGRES_PASSWORD", "postgres")
	database := getEnv("POSTGRES_DB", "postgres")
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)
}
