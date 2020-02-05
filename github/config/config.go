package config

import (
	"errors"
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
	if key == "" {
		return nil, errors.New("GITHUB_WEBHOOK_SECRET_KEY is empty")
	}

	url := os.Getenv("SLACK_WEBHOOK_URL")
	if url == "" {
		return nil, errors.New("SLACK_WEBHOOK_URL is empty")
	}

	return &Config{Port: port, GithubWebhookSecretKey: key, SlackWebhookURL: url}, nil
}
