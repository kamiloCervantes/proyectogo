package config

import (
	"os"
	e "proyectogo/internal/utils/env"
)

type Config struct {
	ServerAddress string
	DatabaseDSN   string
}

func LoadConfig() *Config {

	Env := e.NewEnv()
	Env.Env()

	serverPort := Env.ServerPortQuery
	if Env.Mode == "WRITE" {
		serverPort = Env.ServerPortWrite
	}
	return &Config{

		ServerAddress: getEnv("SERVER_ADDRESS", ":"+serverPort),
		DatabaseDSN:   getEnv("DATABASE_DSN", "host="+Env.DbServer+" user="+Env.DbUser+" password="+Env.DbPassword+" dbname="+Env.DbName+" port="+Env.DbPort+" sslmode=disable"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
