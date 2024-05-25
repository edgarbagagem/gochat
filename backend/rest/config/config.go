package config

import "os"

type Config struct {
	DBName    string
	DBToken   string
	Port      string
	JWTSecret string
}

var Envs = InitConfig()

func InitConfig() Config {
	return Config{
		Port:      getEnv("PORT", ":8080"),
		DBName:    getEnv("DB_NAME", "gochat"),
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
