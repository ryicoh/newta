package handler

import (
	"log"
	"net/http"

	"github.com/ryicoh/newta/github/usecase"
	"gopkg.in/go-playground/webhooks.v5/github"
)

type GithubHandler struct {
	hook    *github.Webhook
	usecase *usecase.Usecase
}

func NewGithubHandler(hook *github.Webhook, usecase *usecase.Usecase) *GithubHandler {
	return &GithubHandler{hook, usecase}
}

func (h *GithubHandler) WebHookHandler(rw http.ResponseWriter, r *http.Request) {
	payload, err := h.hook.Parse(r, github.PullRequestEvent)
	if err != nil {
		if err == github.ErrEventNotFound {
			return
		}
		log.Println(err)
		return
	}

	pullRequest, ok := payload.(github.PullRequestPayload)
	if !ok {
		log.Println(err)
		return
	}

	if len(pullRequest.PullRequest.Assignees) > 0 && len(pullRequest.PullRequest.Labels) > 0 {
		var assignees, labels []string
		for _, assignee := range pullRequest.PullRequest.Assignees {
			assignees = append(assignees, assignee.Login)
		}
		for _, label := range pullRequest.PullRequest.Labels {
			labels = append(labels, label.Name)
		}

		err := h.usecase.SendNotificationAssigned(pullRequest.PullRequest.HTMLURL, assignees, labels)
		if err != nil {
			log.Println(err)
		}

	}
}
