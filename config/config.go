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
	HomePage       string `yaml:"home_page" json:"home_page"`             // the page to render at /
	HighlightStyle string `yaml:"highlight_style" json:"highlight_style"` // The name of the style used to format code blocks
	Root           string `yaml:"root" json:"root"`                       // this will be prepended to any local URLs in the markdown
	UseTailwind    bool   `yaml:"tailwind" json:"tailwind"`               // wether to use tailwind classes and settings (notably restore the proper size of titles)
	FullPage       bool   `yaml:"full_page" json:"full_page"`             // wether to render a full page (add <!DOCTYPE html> to the top of the page
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
			HomePage:       "index",
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
