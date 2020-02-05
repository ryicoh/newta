package usecase

import (
	"fmt"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/ryicoh/newta/pkg/github/config"
)

type Usecase struct {
	slackWebhookURL string
}

func NewUsecase(c *config.Config) *Usecase {
	return &Usecase{slackWebhookURL: c.SlackWebhookURL}
}

func (u *Usecase) SendNotificationAssigned(pullRequestURL string, users []string, labels []string) []error {
	attachment := slack.Attachment{}
	for _, user := range users {
		for _, label := range labels {
			attachment.
				AddField(slack.Field{Title: "Assignee", Value: user}).
				AddField(slack.Field{Title: "Label", Value: label})
		}
	}

	payload := slack.Payload{
		Text:        fmt.Sprintf("<%s|PR>にアサインされました.", pullRequestURL),
		IconEmoji:   ":monkey_face:",
		Attachments: []slack.Attachment{attachment},
	}

	errs := slack.Send(u.slackWebhookURL, "", payload)
	if len(errs) > 0 {
		return errs
	}

	return nil
}
