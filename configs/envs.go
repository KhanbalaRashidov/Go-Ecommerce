package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var Envs = initConfig()

type Config struct {
	Host       string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

func initConfig() Config {
	godotenv.Load()

	return Config{
		Host:       getEnv("HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "mypassword"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "ecom"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
