package config

import "os"

// Config is the global project config.
type Config struct {
	Server
	PostgreSQL
}

// New creates a new Config instance.
func New() *Config {
	return &Config{
		Server: Server{
			Addr: getEnv("SERVER_ADDR", ":8080"),
		},
		PostgreSQL: PostgreSQL{
			Host:     getEnv("POSTGRES_HOST", "locahost"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
			User:     getEnv("POSTGRES_USER", "postgres"),
			Password: getEnv("POSTGRES_PASSWORD", "123"),
			DbName:   getEnv("POSTGRES_DBNAME", "tasker"),
			SSLMode:  getEnv("POSTGRES_SSLMODE", "disable"),
		},
	}
}

// getEnv is the os.Getenv but with default value if environment variable wasn't set.
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}

	return value
}
