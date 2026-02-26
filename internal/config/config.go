package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName     string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		// Log a fatal error if the .env file is not found or cannot be loaded
		// In production, variables are typically set in the environment directly.
		log.Fatalf("Error loading .env file: %s", err)
	}

	dbPortInt, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Invalid DB_PORT")
	}

	return &Config{
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     dbPortInt,
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
	}
}
