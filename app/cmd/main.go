package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/pamungkaski/camar/client"
	"github.com/pamungkaski/camar/datamodel"
	"github.com/prometheus/common/log"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "http://dataweb.bmkg.go.id/inatews/gempadirasakan.xml", nil)
	if err != nil {
		log.Fatal(err)
	}
	clie := client.NewClient()

	_, body, err := clie.Do(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	data := datamodel.BMKGQuakes{}

	if err = xml.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	for _, q := range data.Gempa {
		fmt.Println(q.Dirasakan)
	}
}
