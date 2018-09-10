package main

import (
	"fmt"

	"github.com/pamungkaski/camar/grabber"
)

func main() {
	usgs := grabber.NewGrabber("https://earthquake.usgs.gov/fdsnws/event/1/query")

	data, _ := usgs.GetEarthquakeData("us2000ha1k")

	fmt.Println(data.Properties.Title)
}
