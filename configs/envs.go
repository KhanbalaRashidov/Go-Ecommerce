package configs

import (
	"github.com/joho/godotenv"
	"os"
)

var Envs = initConfig()

type Config struct {
	Host string
	Port string
}

func initConfig() Config {
	godotenv.Load()

	return Config{
		Host: getEnv("HOST", "http://localhost"),
		Port: getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
