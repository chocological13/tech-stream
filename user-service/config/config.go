package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	DatabaseUrl string
	Port        int
	JWTSecret   string
	Environment string
	RedisAddr   string
}

// Loads config from .env
func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file for user-service")
	}

	dsn := os.Getenv("DATABASE_URL")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	environment := os.Getenv("ENV")
	redisAddr := os.Getenv("REDIS_ADDR")

	return &Config{
		DatabaseUrl: dsn,
		Port:        port,
		JWTSecret:   jwtSecret,
		Environment: environment,
		RedisAddr:   redisAddr,
	}, nil

}
