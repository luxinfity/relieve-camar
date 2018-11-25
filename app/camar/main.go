package main

import (
	"log"
	"os"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
	"github.com/pamungkaski/camar"
	"github.com/pamungkaski/camar/client"
	"github.com/pamungkaski/camar/grabber"
	"github.com/pamungkaski/camar/handler"
	"github.com/pamungkaski/camar/listener"
	"github.com/pamungkaski/camar/notifier"
	"github.com/pamungkaski/camar/recorder"
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
	runningPort := os.Getenv("RUNNING_PORT")
	bmkgUserIdString := os.Getenv("BMKG_USER_ID")
	mg, err := recorder.NewMongoDB(username, password, host)
	if err != nil {
		log.Fatal(err)
	}

	bmkgID, _ := strconv.ParseInt(bmkgUserIdString, 10, 64)

	config := oauth1.NewConfig(apiKey, apiKeySecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	twitClient := twitter.NewClient(httpClient)

	grab := grabber.NewGrabber("http://dataweb.bmkg.go.id/inatews/gempadirasakan.xml", client.NewClient())

	fcm := notifier.NewAlerter()

	cam := camar.NewDisasterReporter(grab, mg, fcm)

	twit := listener.NewListener(cam, twitClient, bmkgID)

	route := handler.NewRouter(cam)

	go route.Run(":" + runningPort)

	twit.ListenToQuake()
}
