package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	AppName      string
	Env          string
	RapidAPIKey  string
	RapidAPIHost string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &Config{
		Port:         getEnv("PORT", "8080"),
		AppName:      getEnv("APP_NAME", "price-comparison-engine"),
		Env:          getEnv("ENV", "development"),
		RapidAPIKey:  getEnv("RAPIDAPI_KEY", ""),
		RapidAPIHost: getEnv("RAPIDAPI_HOST", ""),
	}
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}