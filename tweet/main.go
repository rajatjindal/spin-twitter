package main

import (
	"context"
	"net/http"

	"github.com/dghubble/oauth1"
	"github.com/g8rswimmer/go-twitter/v2"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		config := oauth1.NewConfig("consumerKey", "consumerSecret")
		token := oauth1.NewToken("token", "tokenSecret")

		spinhttpclient := spinhttp.NewClient()
		ctx := context.WithValue(oauth1.NoContext, oauth1.HTTPClient, spinhttpclient)
		oauthClient := config.Client(ctx, token)

		// add client to use when making api call as owner of app
		client := &twitter.Client{
			Host:       "https://api.twitter.com",
			Client:     oauthClient,
			Authorizer: &noop{},
		}

		_, err := client.CreateTweet(context.TODO(), twitter.CreateTweetRequest{
			Text: "hello world using Fermyon Spin sdk",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

func main() {}

// github.com/g8rswimmer/go-twitter/v2 expects Authorizer to add auth header
// in this example here, we make use of oauth1 client to do this for us
// so adding this noop authorizer to satisfy the sdk
type noop struct{}

func (a noop) Add(req *http.Request) {}
