package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() *Config {
	godotenv.Load()

	log.Printf("[Config] Loading environment variables...")

	return &Config{
		Domain: getEnv("DOMAIN"),
		Host:   getEnv("HOST"),
		Port:   getDefaultEnv("PORT", "80"),
	}
}

func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	log.Fatalf("Environment variable %s is not set", key)
	return ""
}

func getDefaultEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
