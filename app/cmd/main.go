package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/pamungkaski/camar"
	"github.com/pamungkaski/camar/handler"
	"github.com/pamungkaski/camar/recorder"
	"github.com/prometheus/common/log"
	"net/http"
	"os"
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

	cam := camar.NewDisasterReporter(mg)

	route := handler.NewRouter(cam)

	log.Fatal(http.ListenAndServe(":8080", route))
}
