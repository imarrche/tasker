package api

import "os"

// Config is the global project config.
type Config struct {
	serverConfig
}

// NewConfig creates a new Config instance.
func NewConfig() *Config {
	return &Config{
		serverConfig: serverConfig{
			Addr: getEnv("SERVER_ADDR", ":8080"),
		},
	}
}

// serverConfig is the config that keeps all server related settings.
type serverConfig struct {
	Addr string
}

// getEnv is the os.Getenv but with default value if environment variable wasn't set.
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}

	return value
}
