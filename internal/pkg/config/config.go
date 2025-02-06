package config

import (
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
}

func NewConfig() *Config {
	return loadFromEnv()
}

func loadFromEnv() *Config {
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
	}
}
