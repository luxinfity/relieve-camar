package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/pamungkaski/camar"
	"github.com/pamungkaski/camar/client"
	"github.com/pamungkaski/camar/grabber"
	"github.com/pamungkaski/camar/listener"
	"github.com/pamungkaski/camar/recorder"
	"github.com/pamungkaski/camar/alerter"
	"github.com/pamungkaski/camar/writter"
	"github.com/dghubble/go-twitter/twitter"
	"fmt"
	"strings"
	"context"
	"net/http"
	"github.com/pkg/errors"
)

func getEarthquakeEventID(link string) (string, error) {
	resp, err := http.Get(link)
	if err != nil {
		return "", errors.Wrap(err, "get earthquake ID error")
	}

	// Your magic function. The Request in the Response is the last URL the
	// client tried to access.
	finalURL := resp.Request.URL.String()
	split := strings.Split(finalURL, "/")

	return split[len(split)-1], nil
}

func main() {
	godotenv.Load()
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	apiKey := os.Getenv("API_KEY")
	apiKeySecret := os.Getenv("API_KEY_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	//runningPort := os.Getenv("RUNNING_PORT")

	mg, err := recorder.NewMongoDB(username, password, host)
	if err != nil {
		log.Fatal(err)
	}

	twit := listener.NewListener(apiKey, apiKeySecret, accessToken, accessTokenSecret)

	usgs := grabber.NewGrabber("https://earthquake.usgs.gov/fdsnws/event/1/query", client.NewClient())

	fcm := alerter.NewAlerter()

	c := camar.NewDisasterReporter(twit, mg, usgs, 94119095, &writter.Writer{}, fcm)
	usr := &twitter.UserTimelineParams{
		UserID: 94119095,
		Count: 200,
	}

	dts, _, _ := twit.Timelines.UserTimeline(usr)
	for _, tweet := range dts {
		text := tweet.Text
		splitted := strings.Split(text, " ")
		textLength := len(splitted)

		// Get
		id, err := getEarthquakeEventID(splitted[textLength-1])
		if err != nil {
			fmt.Println(err)
		}

		data, err := usgs.GetEarthquakeData(id)
		if err != nil {
			fmt.Println(err)
		}

		country, err := usgs.GetEarthquakeCountry(data)
		if err != nil {
			fmt.Println(err)
		}

		if country.CountryName == "Indonesia" {
			data, err = c.RecordDisaster(context.Background(), data)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(data.Properties.Title)
		} else {
			data, err = c.RecordInternationalDisaster(context.Background(), data)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(data.Properties.Title)
		}
	}
}
