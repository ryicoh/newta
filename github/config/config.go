package config

import (
	"os"
)

type Config struct {
	Port                   string
	GithubWebhookSecretKey string
	SlackWebhookURL        string
}

func New() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	key := os.Getenv("GITHUB_WEBHOOK_SECRET_KEY")

	url := os.Getenv("SLACK_WEBHOOK_URL")

	return &Config{Port: port, GithubWebhookSecretKey: key, SlackWebhookURL: url}, nil
}
