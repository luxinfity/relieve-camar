// Package grabber is the detail implementation of ResourceGrabber interface from Package Camar.
// It use for grabbing USGS Earthquake data using eventID from USGS twitter
package grabber

import (
	"net/url"
	"net/http"
	"encoding/json"

	"github.com/pkg/errors"
)

// USGS is the main struct that implement ResourceGrabber interface.
// The main usage of this struct is to get earthquake data.
type USGS struct {
	endpoint string
	api http.Client
}

// USGSEarthquakeData is the main struct to wrap data from USGS endpoint.
type USGSEarthquakeData struct {
	Type string `json:"type"`
	EventID    string                   `json:"id"`
	Geometry   USGSGeometry             `json:"geometry"`
	Properties USGSEarthquakeProperties `json:"properties"`
}

// USGSGeometry is the struct that holds geo location of USGS event.
type USGSGeometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

// USGSEarthquakeProperties is the struct that holds detailed information of earthquake from USGS endpoint.
type USGSEarthquakeProperties struct {
	Title string `json:"title"`
	Magnitude float64 `json:"mag"`
	Time      float64 `json:"time"`
	Tsunami   int     `json:"tsunami"`
	Alert     string  `json:"alert"`
}

// NewGrabber is the function used to initiate USGS client.
// It save the USGS endpoint to get earthquake data.
func NewGrabber(endpoint string) *USGS {
	return &USGS{
		endpoint: endpoint,
		api: http.Client{},
	}
}

// GetEarthquakeData is the main function fo USGS to get and wrap USGS data into camar data.
func (u *USGS)GetEarthquakeData(eventID string) (USGSEarthquakeData, error){
	var data USGSEarthquakeData
	endpoint, err := u.buildUSGSQuery(eventID)
	if err != nil {
		return data, err
	}

	//fmt.Println(endpoint.String())

	data, err = u.hitUSGSendpoint(endpoint)
	if err != nil {
		return data, err
	}

	return data, nil
}

// buildUSGSQuery is the function to build USGS endpoint with the detailed query needed.
func (u *USGS)buildUSGSQuery(eventID string) (*url.URL, error) {
	endpoint, err := url.Parse(u.endpoint)
	if err != nil {
		return nil, errors.Wrap(err,"failed to create USGSquery")
	}
	endpoint.Scheme = "https"
	endpoint.Host = "earthquake.usgs.gov"

	query := endpoint.Query()
	query.Add("format", "geojson")
	query.Add("eventid", eventID)

	endpoint.RawQuery = query.Encode()

	return endpoint, nil
}

func (u *USGS)hitUSGSendpoint(endpoint *url.URL) (USGSEarthquakeData, error) {
	var data USGSEarthquakeData

	res, err := u.api.Get(endpoint.String())
	if err != nil {
		return data, errors.Wrap(err,"failed to retrive data from usgs")
	}

	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		return data, errors.Wrap(err,"failed to retrive data from usgs")
	}

	defer res.Body.Close()

	return data, nil
}
