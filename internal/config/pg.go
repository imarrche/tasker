package config

// PostgreSQL is the config for PostgreSQL database.
type PostgreSQL struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSLMode  string
}
