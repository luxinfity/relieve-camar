package main

import (
	"os"
	"log"
	"github.com/joho/godotenv"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo"
	"github.com/pamungkaski/camar/datamodel"
	"time"
)

type CamarQuakeData struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Title    string        `json:"title"`
	Location struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"location"`
	Mag   float64 `json:"mag"`
	Depth float64 `json:"depth"`
	Place string  `json:"place"`
	Time  int64   `json:"time"`
}

func main() {
	godotenv.Load()
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	mongoCredential := &mgo.Credential{
		Username:    username,
		Password:    password,
		Source:      "camar",
		ServiceHost: host,
	}
	mg, _ := mgo.Dial(host)

	mg.SetMode(mgo.Monotonic, true)

	if err := mg.Login(mongoCredential); err != nil {
		log.Fatal(err)
	}

	dbEve := mg.DB("camar").C("event")
	_, err := dbEve.RemoveAll(nil)
	if err != nil {
		log.Fatal(err)
	}

	dbEarth := mg.DB("camar").C("earthquake")
	var derth []CamarQuakeData
	if err := dbEarth.Find(nil).All(&derth); err != nil {
		log.Fatal(err)
	}

	for _, quake := range derth {
		var event datamodel.Event

		event.ID = bson.NewObjectId()
		event.Location = quake.Location
		event.Time = quake.Time
		event.TimeArrived = time.Now().Local().Unix()
		event.EventType = "Earthquake"
		event.Source = "BMKG"
		event.EventDetail = datamodel.CamarQuakeData{
			Title: quake.Title,
			Mag: quake.Mag,
			MMI: quake.Place,
			Depth: quake.Mag,
		}

		dbEve.Insert(&event)
	}
}
