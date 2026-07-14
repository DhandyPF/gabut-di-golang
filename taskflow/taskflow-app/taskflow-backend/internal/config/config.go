package config

import (
	"os"
)

// Config holds runtime configuration loaded from environment variables.
type Config struct {
	Port        string
	DatabaseDSN string
	JWTSecret   string
}

// Load reads configuration from environment variables, falling back to
// sensible defaults for local development.
func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseDSN: getEnv("DATABASE_DSN", "./taskflow.db"),
		JWTSecret:   getEnv("JWT_SECRET", "change-this-secret-in-production"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
