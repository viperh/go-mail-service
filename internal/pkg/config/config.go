package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	SMTPTarget   string
	SMTPPort     int
	SMTPUser     string
	SMTPPassword string

	APIPort string

	JwtSecret string
}

func NewConfig() *Config {
	return loadFromEnv()
}

func loadFromEnv() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Invalid port in .env file")
	}

	return &Config{
		SMTPTarget:   os.Getenv("TARGET"),
		SMTPPort:     port,
		SMTPUser:     os.Getenv("USER"),
		SMTPPassword: os.Getenv("PASSWORD"),
		APIPort:      os.Getenv("API_PORT"),
		JwtSecret:    os.Getenv("JWT_SECRET"),
	}
}
