package main

import (
	"fmt"

	"github.com/pamungkaski/camar/client"
	"github.com/pamungkaski/camar/grabber"
)

func main() {
	usgs := grabber.NewGrabber("https://earthquake.usgs.gov/fdsnws/event/1/query", client.NewClient())

	data, err := usgs.GetEarthquakeData("us2000ha1k")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data.Properties.Title)
}
