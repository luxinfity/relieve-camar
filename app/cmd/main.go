package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/pamungkaski/camar/recorder"
	"github.com/prometheus/common/log"
	"os"
	"github.com/pamungkaski/camar"
	"github.com/globalsign/mgo/bson"
)

func main() {
	godotenv.Load()
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	fmt.Println(username)
	fmt.Println(password)
	fmt.Println(host)
	mg, err := recorder.NewMongoDB(username, password, host)
	if err != nil {
		log.Fatal(err)
	}

	//usgs := grabber.NewGrabber("https://earthquake.usgs.gov/fdsnws/event/1/query", client.NewClient())
	//
	//data, err := usgs.GetEarthquakeData("us2000ha1k")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//data.BsonID = bson.NewObjectId()
	Coordinate :=[]float64{120.2015625, -7.6078738}
	dev := camar.Device{
		ID: bson.NewObjectId(),
		DeviceID: "samsok",
	}
	dev.Location.Type = "Point"
	dev.Location.Coordinates = Coordinate
	mg.NewDevice(dev)

	//fmt.Println(Coordinate)
	res, err := mg.GetDeviceInRadius(Coordinate, 1.36)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
