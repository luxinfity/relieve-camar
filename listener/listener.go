package listener

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func NewListener(consumerKey, consumerSecret, accessToken, accessSecret string) *twitter.Client {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)
	return client
}
