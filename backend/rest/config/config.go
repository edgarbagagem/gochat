package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBName    string
	DBToken   string
	Port      string
	JWTSecret string
}

var Envs = InitConfig()

func InitConfig() Config {
	godotenv.Load()
	return Config{
		Port:      getEnv("PORT", ":8080"),
		DBName:    getEnv("DB_NAME", "dbname"),
		DBToken:   getEnv("DB_TOKEN", "randomdatabasetoken"),
		JWTSecret: getEnv("JWT_SECRET", "randomjwtsecret"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
