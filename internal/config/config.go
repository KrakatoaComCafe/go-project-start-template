package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port string
}

func NewAppConfig() *AppConfig {
	if err := godotenv.Load("conf/.env"); err != nil {
		log.Println("File .env not found, using default values")
	}
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}
	return &AppConfig{
		Port: port,
	}
}
