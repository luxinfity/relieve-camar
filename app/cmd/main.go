package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/pamungkaski/camar"
	"github.com/pamungkaski/camar/client"
	"github.com/pamungkaski/camar/grabber"
	"github.com/pamungkaski/camar/notifier"
	"github.com/pamungkaski/camar/recorder"
	"context"
)

func main() {
	godotenv.Load()
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	authDB := os.Getenv("MONGO_AUTH_DB")
	mg, err := recorder.NewMongoDB(username, password, host, authDB)
	if err != nil {
		log.Fatal(err)
	}

	grab := grabber.NewGrabber("http://dataweb.bmkg.go.id/inatews/gempadirasakan.xml", client.NewClient())

	fcm := notifier.NewAlerter()

	cam := camar.NewDisasterReporter(grab, mg, fcm)
	latest, err := grab.GetEarthquakes()
	if err != nil {
		log.Fatal(err)
	}

	for _, event := range latest {
		if err := cam.RecordDisaster(context.Background(), event); err != nil {
			log.Fatal(err)
		}
	}
}
