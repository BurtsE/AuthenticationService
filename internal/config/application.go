package config

func GetApplicationPort() string {
	return getEnv("APPLICATION_PORT", ":8080")
}

func GetApplicationEnvironment() string {
	return getEnv("APPLICATION_ENVIRONMENT", "development")
}
