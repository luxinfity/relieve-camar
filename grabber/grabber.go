// Package grabber is the detail implementation of ResourceGrabber interface from Package Camar.
// It use for grabbing USGS Earthquake data using eventID from USGS twitter
package grabber

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"net/url"

	"github.com/pamungkaski/camar/client"
	"github.com/pamungkaski/camar/datamodel"
)

// ResourceGrabber is the bussiness logic contract for getting earthquake data.
type ResourceGrabber interface {
	// GetEarthQuakeData is a function to to retrieve Earthquake detailed data.
	GetEarthquakes() ([]datamodel.CamarQuakeData, error)
}

// USGS is the main struct that implement ResourceGrabber interface.
// The main usage of this struct is to get earthquake data.
type USGS struct {
	endpoint string
	api      client.Client
}

// NewGrabber is the function used to initiate USGS client.
// It save the USGS endpoint to get earthquake data.
func NewGrabber(endpoint string, api client.Client) *USGS {
	return &USGS{
		endpoint: endpoint,
		api:      api,
	}
}

// GetEarthquakeData is the main function fo USGS to get and wrap USGS data into camar data.
func (u *USGS) GetEarthquakeData(eventID string) (datamodel.GeoJSON, error) {
	var data datamodel.GeoJSON
	req, err := u.buildUSGSRequest(eventID)
	if err != nil {
		return data, err
	}

	//fmt.Println(req.URL.String())

	_, body, err := u.api.Do(context.Background(), req)
	if err != nil {
		return data, err
	}

	if err = json.Unmarshal(body, &data); err != nil {
		return data, errors.Wrap(err, "failed to retrive data from usgs")
	}

	//fmt.Println(body)
	data.URL = req.URL.String()

	return data, nil
}

// buildUSGSQuery is the function to build USGS endpoint with the detailed query needed.
func (u *USGS) buildUSGSRequest(eventID string) (*http.Request, error) {
	endpoint, err := url.Parse(u.endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create USGSquery")
	}
	endpoint.Scheme = "https"
	endpoint.Host = "earthquake.usgs.gov"

	query := endpoint.Query()
	query.Add("format", "geojson")
	query.Add("eventid", eventID)

	endpoint.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}