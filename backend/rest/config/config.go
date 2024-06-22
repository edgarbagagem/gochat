package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBName                       string
	DBToken                      string
	Port                         string
	JWTSecret                    string
	GoogleApplicationCredentials string
	JWTExpirationInSeconds       int64
}

var Envs = InitConfig()

func InitConfig() Config {
	godotenv.Load()
	return Config{
		Port:                         getEnv("PORT", ":8080"),
		DBName:                       getEnv("DB_NAME", "dbname"),
		DBToken:                      getEnv("DB_TOKEN", "randomdatabasetoken"),
		JWTSecret:                    getEnv("JWT_SECRET", "randomjwtsecret"),
		JWTExpirationInSeconds:       getEnvAsInt("JWT_EXPIRATION_SECONDS", 3600),
		GoogleApplicationCredentials: getEnv("GOOGLE_APPLICATION_CREDENTIALS", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
