package config

import (
	"os"
)

type Config struct {
	ServerAddress string
	DatabaseDSN   string
}

func LoadConfig() *Config {
	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
<<<<<<< Updated upstream
		DatabaseDSN:   getEnv("DATABASE_DSN", "host=localhost user=postgres password=camila dbname=db_url port=5432 sslmode=disable"),
=======
		DatabaseDSN:   getEnv("DATABASE_DSN", "host=database user=lando password=lando dbname=proyectogo port=5432 sslmode=disable"),
>>>>>>> Stashed changes
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
