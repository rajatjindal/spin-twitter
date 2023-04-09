module github.com/rajatjindal/spin-twitter/webhook

go 1.20

require (
	github.com/fermyon/spin/sdk/go v1.0.0
	github.com/g8rswimmer/go-twitter/v2 v2.1.5
	github.com/gorilla/mux v1.8.0
	github.com/rajatjindal/twitter-bot/v2 v2.0.0-20230409054342-ddab727b085d
)

require (
	github.com/dghubble/oauth1 v0.7.1 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
)

replace github.com/fermyon/spin/sdk/go => ../../../fermyon/spin/sdk/go
