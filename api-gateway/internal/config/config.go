package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Address            string
	Port               int
	UserServiceAddress string
	JWTSecret          string
	Env                string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Error parsing PORT")
	}

	return &Config{
		Address:            fmt.Sprintf(":%d", 8080),
		Port:               port,
		UserServiceAddress: os.Getenv("ADDRESS_USER_SERVICE"),
		JWTSecret:          os.Getenv("JWT_SECRET"),
		Env:                os.Getenv("ENV"),
	}, nil
}
