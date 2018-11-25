// Package grabber is the detail implementation of ResourceGrabber interface from Package Camar.
// It use for grabbing USGS Earthquake data using eventID from USGS twitter
package grabber

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/pamungkaski/camar/client"
	"github.com/pamungkaski/camar/datamodel"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ResourceGrabber is the bussiness logic contract for getting earthquake data.
type ResourceGrabber interface {
	// GetEarthQuakeData is a function to to retrieve Earthquake detailed data.
	GetEarthquakes() ([]datamodel.CamarQuakeData, error)
}

// USGS is the main struct that implement ResourceGrabber interface.
// The main usage of this struct is to get earthquake data.
type BMKG struct {
	endpoint string
	api      client.Client
}

// NewGrabber is the function used to initiate USGS client.
// It save the USGS endpoint to get earthquake data.
func NewGrabber(endpoint string, api client.Client) ResourceGrabber {
	return &BMKG{
		endpoint: endpoint,
		api:      api,
	}
}

func (b *BMKG) GetEarthquakes() ([]datamodel.CamarQuakeData, error) {
	var data datamodel.BMKGQuakes
	req, err := http.NewRequest(http.MethodGet, b.endpoint, nil)
	if err != nil {
		return nil, err
	}

	_, body, err := b.api.Do(context.Background(), req)
	if err != nil {
		return nil, err
	}

	if err := xml.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return b.typecastBMKGQuakeToCamar(data), nil
}

func (b *BMKG) typecastBMKGQuakeToCamar(quakes datamodel.BMKGQuakes) []datamodel.CamarQuakeData {
	var data []datamodel.CamarQuakeData
	for _, gempa := range quakes.Gempa {
		var quake datamodel.CamarQuakeData
		mag, _ := strconv.ParseFloat(gempa.Magnitude, 64)
		dep, _ := strconv.ParseFloat(strings.Split(gempa.Kedalaman, " ")[0], 64)
		coors := strings.Split(gempa.Point.Coordinates, " ")
		latitude, _ := strconv.ParseFloat(strings.Split(coors[0], ",")[0], 64)
		Longitude, _ := strconv.ParseFloat(coors[1], 64)
		wkt, _ := time.Parse("2/1/2006-15:04:05", strings.Split(gempa.Tanggal, " ")[0])

		quake.Title = fmt.Sprintf("Gempa Mag:%.1f, %s, %s pada kedalaman %s dapat dirasakan di %s", mag, wkt.Format("2/1/2006-15:04:05"), gempa.Keterangan, gempa.Kedalaman, gempa.Dirasakan)
		quake.Time = wkt.Unix()
		quake.Location.Type = "Point"
		quake.Location.Coordinates = append(quake.Location.Coordinates, Longitude)
		quake.Location.Coordinates = append(quake.Location.Coordinates, latitude)
		quake.Mag = mag
		quake.Depth = dep
		quake.Place = gempa.Dirasakan

		data = append(data, quake)
	}

	return data
}
