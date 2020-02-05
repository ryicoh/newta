package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/ryicoh/newta/github/config"
	"github.com/ryicoh/newta/github/handler"
	"github.com/ryicoh/newta/github/usecase"
	"go.uber.org/dig"
	"gopkg.in/go-playground/webhooks.v5/github"
)

func main() {
	log.Println("This is Newta")

	c := dig.New()

	c.Provide(config.New)

	c.Provide(func(c *config.Config) (*github.Webhook, error) {
		hook, err := github.New(github.Options.Secret(c.GithubWebhookSecretKey))
		if err != nil {
			log.Printf("err: %s", err)
			return nil, nil
		}
		return hook, nil
	})

	c.Provide(usecase.NewUsecase)

	c.Provide(handler.NewGithubHandler)

	err := c.Invoke(func(h *handler.GithubHandler, c *config.Config) {

		http.HandleFunc("/", h.WebHookHandler)

		wg := &sync.WaitGroup{}

		wg.Add(1)
		go func() {
			log.Fatal(http.ListenAndServe(":"+c.Port, nil))
			wg.Done()
		}()

		fmt.Println("HTTP server Started")

		wg.Wait()

		fmt.Println("HTTP server Ended")
	})

	if err != nil {
		log.Fatal(err)
	}
}
