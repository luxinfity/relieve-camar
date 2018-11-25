package listener

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/pamungkaski/camar"
)

type TwitterListener interface {
	ListenToQuake()
}

type Listener struct {
	camar     camar.Camar
	client    *twitter.Client
	twitterID int64
}

func NewListener(consumerKey, consumerSecret, accessToken, accessSecret string) *twitter.Client {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)
	return client
}

func (l *Listener) ListenToQuake() {
	fmt.Println("Starting Stream...")

	// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		if l.validateTwit(tweet) {
			if err := l.camar.ListenTheEarth(); err != nil {
				log.Println(err)
			}
		}
	}

	usgsStringId := strconv.FormatInt(l.twitterID, 10)
	// FILTER
	params := &twitter.StreamFilterParams{
		Follow:        []string{usgsStringId},
		StallWarnings: twitter.Bool(true),
	}

	stream, err := l.client.Streams.Filter(params)
	if err != nil {
		log.Fatal(err)
	}

	// Receive messages until stopped or stream quits
	go demux.HandleChan(stream.Messages)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	stream.Stop()
}

func (l *Listener) validateTwit(tweet *twitter.Tweet) bool {
	twit := tweet.Text

	if tweet.User.ID != l.twitterID {
		return false
	}

	word := strings.Split(twit, " ")
	if word[0] != "#Gempa" {
		return false
	}

	date := strings.Split(word[2], "/")

	if len(date) != 3 {
		return false
	}

	return true
}
