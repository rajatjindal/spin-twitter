package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/g8rswimmer/go-twitter/v2"
	"github.com/gorilla/mux"
	bothelper "github.com/rajatjindal/twitter-bot/v2/twitter"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stderr, "INCOMING - %s %s\n", r.Method, r.URL.Path)
		spinhttpclient := spinhttp.NewClient()
		botConf := &bothelper.BotConfig{
			Tokens: bothelper.Tokens{
				ConsumerKey:   "consumer-key",
				ConsumerToken: "consumer-token",
				Token:         "token",
				TokenSecret:   "token-secret",
			},
			WebhookConfig: bothelper.WebhookConfig{
				Environment:      "development",
				URL:              "https://webhook-hvhxj8cy.fermyon.app",
				Path:             "/webhook/twitter",
				OverWriteOnLimit: true,
				MaxAllowed:       1,
			},
		}

		bot, err := bothelper.NewBotWithClient(spinhttpclient, botConf)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if r.URL.Path == "/verify-registration" {
			fmt.Fprintln(os.Stderr, "verifying registration")
			err := bot.EnsureWebhookIsActive()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			return
		}

		router := mux.NewRouter().StrictSlash(true)
		router.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// this helps with CRC validation that Twitter do periodically
		router.Methods(http.MethodGet).Path(bot.WebhookPath()).HandlerFunc(bot.HandleCRCResponse)

		// this is your webhook handler
		wh := &Handler{
			bot: bot,
		}
		router.Methods(http.MethodPost).Path(bot.WebhookPath()).HandlerFunc(wh.ServeHTTP)

		router.ServeHTTP(w, r)
	})
}

func main() {}

type Handler struct {
	bot *bothelper.Bot
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(string(body))

	_, err = h.bot.AsOwnerOfApp().CreateTweet(context.TODO(), twitter.CreateTweetRequest{
		Text: "hello back to @rajatjindal",
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
