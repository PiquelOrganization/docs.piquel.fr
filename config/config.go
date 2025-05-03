package config

import (
	"log"
	"os"
	"sync"

	"github.com/PiquelOrganization/docs.piquel.fr/utils"
	"github.com/joho/godotenv"
)

type Config struct {
	Envs   Envs       // the technical server configuration
	Config DocsConfig // the documentation specific configuration
}

type Envs struct {
	Domain        string
	Host          string
	Port          string
	DataPath      string // where to clone the repository
	Repository    string // the reposityory to get the pages from
	WebhookSecret string // the secret to use to validate the webhook
}

type DocsConfig struct {
	sync.Mutex
	HomePage       string `json,yaml:"home_page"`       // the page to render at /
	HighlightStyle string `json,yaml:"highlight_style"` // the style used to highlight code
	Root           string `json,yaml:"root"`            // the root used to return
	UseTailwind    bool   `json,yaml:"use_tailwind"`    // wether to add classes to reverse tailwind config
}

func LoadConfig() *Config {
	godotenv.Load()

	log.Printf("[Config] Loading environment variables...")

	return &Config{
		Envs: Envs{
			Domain:        getEnv("DOMAIN"),
			Host:          getEnv("HOST"),
			Port:          getDefaultEnv("PORT", "80"),
			DataPath:      utils.FormatLocalPathString(getDefaultEnv("DATA_PATH", "/docs/data"), ""),
			Repository:    getEnv("REPOSITORY"),
			WebhookSecret: getEnv("WEBHOOK_SECRET")},
		Config: DocsConfig{
			HomePage:       "index.md",
			HighlightStyle: "tokyonight",
			Root:           "",
		},
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
