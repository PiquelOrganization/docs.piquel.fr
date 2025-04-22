package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadConfig() *Config {
	godotenv.Load()

	log.Printf("[Config] Loading environment variables...")

	useGitString := getDefaultEnv("USE_GIT", "false")
	useGit := useGitString == "true"
	repository := ""

	if useGit {
		repository = getEnv("REPOSITORY")
	}

	return &Config{
		Domain:        getEnv("DOMAIN"),
		Host:          getEnv("HOST"),
		Port:          getDefaultEnv("PORT", "80"),
		DataPath:      getDefaultEnv("DATA_PATH", "/docs/data"),
		UseGit:        useGit,
		Repository:    repository,
		WebhookSecret: getEnv("WEBHOOK_SECRET"),
		HomePage:      strings.ToLower(getDefaultEnv("HOME_PAGE", "")),
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
