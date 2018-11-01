package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/pamungkaski/camar"
	"github.com/pamungkaski/camar/client"
	"github.com/pamungkaski/camar/grabber"
	"github.com/pamungkaski/camar/handler"
	"github.com/pamungkaski/camar/listener"
	"github.com/pamungkaski/camar/recorder"
	"github.com/pamungkaski/camar/alerter"
	"github.com/pamungkaski/camar/writter"
	"strconv"
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
	usgsUserIdString := os.Getenv("USGS_USER_ID")
	mg, err := recorder.NewMongoDB(username, password, host)
	if err != nil {
		log.Fatal(err)
	}

	usgsUserId,_ := strconv.ParseInt(usgsUserIdString, 10, 64)

	twit := listener.NewListener(apiKey, apiKeySecret, accessToken, accessTokenSecret)

	usgs := grabber.NewGrabber("https://earthquake.usgs.gov/fdsnws/event/1/query", client.NewClient())

	fcm := alerter.NewAlerter()

	cam := camar.NewDisasterReporter(twit, mg, usgs, usgsUserId, &writter.Writer{}, fcm)

	route := handler.NewRouter(cam)

	go route.Run(":" +runningPort)

	cam.ListenTheEarth()
}
