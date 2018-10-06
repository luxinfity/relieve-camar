package main

import (
	"github.com/pamungkaski/camar/grabber"
	"fmt"
	"github.com/pamungkaski/camar/client"
)

func main() {
	usgs := grabber.NewGrabber("https://earthquake.usgs.gov/fdsnws/event/1/query", client.NewClient())

	data, err := usgs.GetEarthquakeData("us2000ha1k")
	if err != nil {
		fmt.Println(err)
	}

	country, err := usgs.GetEarthquakeCountry(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data.Geometry.Coordinates)
	fmt.Println(country.CountryName)
}
