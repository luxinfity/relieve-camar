package main

import (
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/dghubble/go-twitter/twitter"
	"os"
	"os/signal"
	"syscall"
	"github.com/joho/godotenv"
	"github.com/dghubble/oauth1"
)

func main() {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	apiKeySecret := os.Getenv("API_KEY_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	fmt.Println("Starting Stream...")
	config := oauth1.NewConfig(apiKey, apiKeySecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	twitClient := twitter.NewClient(httpClient)

	// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Println(tweet)
	}

	// FILTER
	params := &twitter.StreamFilterParams{
		Follow: []string{"382449035"},
		StallWarnings: twitter.Bool(true),
	}

	stream, err := twitClient.Streams.Filter(params)
	if err != nil {
		log.Fatal(err)
	}

	// Receive messages until stopped or stream quits
	go demux.HandleChan(stream.Messages)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Fatal(<-ch)

	fmt.Println("Stopping Stream...")
	stream.Stop()
}
