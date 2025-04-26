package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Domain        string
	Host          string
	Port          string
	DataPath      string // where to clone the repository
	Repository    string // the reposityory to get the pages from
	WebhookSecret string // the secret to use to validate the webhook
	HomePage      string // the page to render at /
}

func LoadConfig() *Config {
	godotenv.Load()

	log.Printf("[Config] Loading environment variables...")

	return &Config{
		Domain:        getEnv("DOMAIN"),
		Host:          getEnv("HOST"),
		Port:          getDefaultEnv("PORT", "80"),
		DataPath:      getDefaultEnv("DATA_PATH", "/docs/data"),
		Repository:    getEnv("REPOSITORY"),
		WebhookSecret: getEnv("WEBHOOK_SECRET"),
		HomePage:      getDefaultEnv("HOME_PAGE", "index"),
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
