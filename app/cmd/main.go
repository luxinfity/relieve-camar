package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"os"
	"os/signal"
	"syscall"
	"log"
	"github.com/joho/godotenv"
	"github.com/pamungkaski/camar/listener"
)

func main() {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	apiKeySecret := os.Getenv("API_KEY_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	usgsUserIdString := os.Getenv("USGS_USER_ID")

	twit := listener.NewListener(apiKey, apiKeySecret, accessToken, accessTokenSecret)

	fmt.Println("Starting Stream...")

	fmt.Println(usgsUserIdString)
	// FILTER
	params := &twitter.StreamFilterParams{
		Follow:        []string{usgsUserIdString},
		StallWarnings: twitter.Bool(true),
	}
	stream, err := twit.Streams.Filter(params)
	if err != nil {
		log.Fatal(err)
	}

	for message := range stream.Messages {
		fmt.Println(message)
	}

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	stream.Stop()
}
