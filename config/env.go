package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
}

var ENVS = initConfig()

func initConfig() Config {
	return Config{
		DatabaseURL: getenv("DATABASE_URL", "DATABASE URL IS MISSING"),
	}
}

func getenv(key, fallback string) string {

	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
