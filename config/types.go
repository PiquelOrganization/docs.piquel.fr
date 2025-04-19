package config

type Config struct {
	Domain        string
	Host          string
	Port          string
	DataPath      string // where to store the pages for caching (should be a docker volume)
	UseGit        bool   // wether to use git (and thus the repository)
	Repository    string // the reposityory to get the pages from
	UseWebhook    bool   // wether to use a GitHub webhook to update the repository
	WebhookSecret string // the secret to use to validate the webhook
}
