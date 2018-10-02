package main

import (
	"github.com/joho/godotenv"
	"github.com/pamungkaski/camar"
	"github.com/pamungkaski/camar/client"
	"github.com/pamungkaski/camar/grabber"
	"github.com/pamungkaski/camar/handler"
	"github.com/pamungkaski/camar/listener"
	"github.com/pamungkaski/camar/recorder"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	apiKey := os.Getenv("API_KEY")
	apiKeySecret := os.Getenv("API_KEY_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	mg, err := recorder.NewMongoDB(username, password, host)
	if err != nil {
		log.Fatal(err)
	}

	twit := listener.NewListener(apiKey, apiKeySecret, accessToken, accessTokenSecret)

	usgs := grabber.NewGrabber("https://earthquake.usgs.gov/fdsnws/event/1/query", client.NewClient())

	cam := camar.NewDisasterReporter(twit, mg, usgs, 94119095)

	route := handler.NewRouter(cam)

	go http.ListenAndServe(":8080", route)

	cam.ListenTheEarth()

}
