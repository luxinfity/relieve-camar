package main

import (
	"github.com/pamungkaski/camar/datamodel"
	"github.com/pamungkaski/camar/writter"
	"strings"
)

func main() {
	//godotenv.Load()
	//username := os.Getenv("MONGO_USERNAME")
	//password := os.Getenv("MONGO_PASSWORD")
	//host := os.Getenv("MONGO_HOST")
	//apiKey := os.Getenv("API_KEY")
	//apiKeySecret := os.Getenv("API_KEY_SECRET")
	//accessToken := os.Getenv("ACCESS_TOKEN")
	//accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	//
	//mg, err := recorder.NewMongoDB(username, password, host)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//
	//twit := listener.NewListener(apiKey, apiKeySecret, accessToken, accessTokenSecret)
	//
	//usgs := grabber.NewGrabber("https://earthquake.usgs.gov/fdsnws/event/1/query", client.NewClient())
	//
	//cam := camar.NewDisasterReporter(twit, mg, usgs)
	//
	//route := handler.NewRouter(cam)
	//
	//go http.ListenAndServe(":8080", route)
	//
	//cam.ListenTheEarth()
	a := writter.Writer{}
	s := "Prelim M5.7 Earthquake Indian Ocean Triple Junction Oct-01 18:16 UTC, updates https://go.usa.gov/xPBNT"
	spl := strings.Split(s, " ")
	a.CreateAlertMessage(datamodel.GeoJSON{}, spl)
}
