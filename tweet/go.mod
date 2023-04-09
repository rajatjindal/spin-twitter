module github.com/rajatjindal/spin-twitter/tweet

go 1.20

require (
	github.com/dghubble/oauth1 v0.7.2
	github.com/fermyon/spin/sdk/go v1.0.0
	github.com/g8rswimmer/go-twitter/v2 v2.1.5
)

require github.com/julienschmidt/httprouter v1.3.0 // indirect

replace github.com/fermyon/spin/sdk/go => github.com/fermyon/spin/sdk/go v1.0.0-rc.1.0.20230407105950-a713eee17b52
