// Package camar is main code file for this service.
// This package will contain every single business logic needed for the service
// benCAna Monitoring and AleRting
package camar

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"firebase.google.com/go/messaging"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"

	"github.com/pamungkaski/camar/datamodel"
	"sync"
)

// DisasterReporter is the business logic contract for camar service.
// The main idea of the interface is to record disaster into database then alert all service's device.
type DisasterReporter interface {
	// ListenTheEarth is a function that Listen to any Earthquake happen.
	// It is the main function of DisasterReporter Interface
	ListenTheEarth()
	// RecordDisaster is a function to save Disaster into our database.
	RecordDisaster(ctx context.Context, disaster datamodel.GeoJSON) (datamodel.GeoJSON, error)
	// RecordDisaster is a function to save Disaster into our database.
	RecordInternationalDisaster(ctx context.Context, disaster datamodel.GeoJSON) (datamodel.GeoJSON, error)
	//
	GetEarthquakeList(ctx context.Context, limit, page int) ([]datamodel.EarthquakeDataSnapshoot, error)
	// AlertDisastrousEvent is a function to alert service's device.
	AlertDisastrousEvent(ctx context.Context, disaster datamodel.GeoJSON) error
	// NewDevice is a function to save new device device for alerting purpose.
	NewDevice(ctx context.Context, device Device) (Device, error)
	// GetDevice
	GetDevice(ctx context.Context, deviceID string) (Device, error)
	// UpdateDevice is a function to update device latitude and longitude coordinate.
	UpdateDevice(ctx context.Context, device Device) (Device, error)
}

// ResourceGrabber is the bussiness logic contract for getting earthquake data.
type ResourceGrabber interface {
	// GetEarthQuakeData is a function to to retrieve Earthquake detailed data.
	GetEarthquakeData(eventID string) (datamodel.GeoJSON, error)
	//
	GetEarthquakeCountry(data datamodel.GeoJSON) (datamodel.CountryData, error)
}

// Recorder is the business logic contract for saving data.
type Recorder interface {
	// SaveDisaster is a function to save disaster data into database
	SaveDisaster(disaster datamodel.GeoJSON) error
	// SaveDisaster is a function to save disaster data into database
	SaveInternationalDisaster(disaster datamodel.GeoJSON) error
	//
	GetEarthquakeList(limit, page int) ([]datamodel.GeoJSON, error)
	// SaveDevice is a function to register device on the alerting service.
	NewDevice(device Device) error
	//
	GetDevice(deviceID string) (Device, error)
	// UpdateDevice is a function to update device latitude and longitude coordinate.
	UpdateDevice(device Device) error
	// GetDeviceInRadius is a function to get all Device data inside the Disastrous Zone Radius.
	GetDeviceInRadius(disasterCoordinate []float64, radius float64) ([]Device, error)
}

// Alerting is the business logic contract for alerting service.
// the main idea is to send alert to all device.
type Alerting interface {
	// SendAlert is a function to send Disastrous Event alert to specific Device using the alerting service.
	SendAlert(alert messaging.Message, errc chan []error, wg *sync.WaitGroup)
}

// AlerWritter is the business logic contract for alert message writter.
type AlertWritter interface {
	// CreateAlertMessage is a function to create alert message based on th disaster event that currently occurs.
	CreateAlertMessage(disaster datamodel.GeoJSON, alerts []string) (messaging.Message, error)
}

// Device is the struct for each device conected to service.
// It save the DeviceID for alerting purpose.
// In device side, it is an automatics Device Regist on first start.
// The Latitude and Longitude are update-able.
type Device struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	DeviceID string        `json:"device_id"`
	Token    string        `json:"token"`
	Location struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"location"`
}

// Camar is the main struct of the service.
// Camar implements DisasterReporter interface.
// It contains alerting and recording interface implementation.
type Camar struct {
	listener      *twitter.Client
	grabber       ResourceGrabber
	recording     Recorder
	alerting      Alerting
	writer        AlertWritter
	usgsTwitterID int64
}

// NewDisasterReporter is a function that creates an instance of DisasterReporter.
func NewDisasterReporter(client *twitter.Client, recorder Recorder, grabber ResourceGrabber, twitterID int64, writer AlertWritter, alerter Alerting) DisasterReporter {
	return &Camar{
		listener:      client,
		recording:     recorder,
		grabber:       grabber,
		usgsTwitterID: twitterID,
		writer:writer,
		alerting:alerter,
	}
}

// ListenTheEarth is a function that Listen to any Earthquake happen.
// It is the main function of DisasterReporter Interface
func (c *Camar) ListenTheEarth() {
	fmt.Println("Starting Stream...")

	// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		if tweet.User.ID == c.usgsTwitterID {
			// Get Shortened link of the event
			text := tweet.Text
			splitted := strings.Split(text, " ")
			textLength := len(splitted)

			// Get
			id, err := c.getEarthquakeEventID(splitted[textLength-1])
			if err != nil {
				fmt.Println(err)
			}

			data, err := c.grabber.GetEarthquakeData(id)
			if err != nil {
				fmt.Println(err)
			}

			country, err := c.grabber.GetEarthquakeCountry(data)
			if err != nil {
				fmt.Println(err)
			}

			if country.CountryName == "Indonesia" {
				//data, err = c.RecordDisaster(context.Background(), data)
				//if err != nil {
				//	fmt.Println(err)
				//}

				//fmt.Println(data.Properties.Title)

				if err := c.AlertDisastrousEvent(context.Background(), data); err != nil {
					fmt.Println(err)
				}
			} else {
				//data, err = c.RecordInternationalDisaster(context.Background(), data)
				//if err != nil {
				//	fmt.Println(err)
				//}
			}
		}
	}

	// FILTER
	params := &twitter.StreamFilterParams{
		Follow:        []string{"94119095"},
		StallWarnings: twitter.Bool(true),
	}
	stream, err := c.listener.Streams.Filter(params)
	if err != nil {
		log.Fatal(err)
	}

	// Receive messages until stopped or stream quits
	go demux.HandleChan(stream.Messages)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	stream.Stop()
}

// getEarthquakeEventID is a function that will get the Earthquake event id from link that is shared.
func (c *Camar) getEarthquakeEventID(link string) (string, error) {
	resp, err := http.Get(link)
	if err != nil {
		return "", errors.Wrap(err, "get earthquake ID error")
	}

	// Your magic function. The Request in the Response is the last URL the
	// client tried to access.
	finalURL := resp.Request.URL.String()
	split := strings.Split(finalURL, "/")

	return split[len(split)-1], nil
}

func (c *Camar) GetEarthquakeList(ctx context.Context, limit, page int) ([]datamodel.EarthquakeDataSnapshoot, error) {
	var list []datamodel.GeoJSON
	var snapList []datamodel.EarthquakeDataSnapshoot

	list, err := c.recording.GetEarthquakeList(limit, page)
	if err != nil {
		return nil, err
	}

	for _, data := range list {
		snap := datamodel.EarthquakeDataSnapshoot{
			Title: data.Properties.Title,
			Location: data.Geometry,
			Mag: data.Properties.Mag,
			Depth: data.Geometry.Coordinates[2],
			Place: data.Properties.Place,
			Time: data.Properties.Time,
			URL: data.URL,
			Tsunami:data.Properties.Tsunami,
		}
		snapList =  append(snapList, snap)
	}

	return snapList, nil
}

// RecordDisaster is a function to save Disaster into our database.
func (c *Camar) RecordDisaster(ctx context.Context, disaster datamodel.GeoJSON) (datamodel.GeoJSON, error) {
	disaster.BsonID = bson.NewObjectId()

	if err := c.recording.SaveDisaster(disaster); err != nil {
		return datamodel.GeoJSON{}, err
	}

	return disaster, nil
}

func (c *Camar) RecordInternationalDisaster(ctx context.Context, disaster datamodel.GeoJSON) (datamodel.GeoJSON, error) {
	disaster.BsonID = bson.NewObjectId()

	if err := c.recording.SaveInternationalDisaster(disaster); err != nil {
		return datamodel.GeoJSON{}, err
	}

	return disaster, nil
}

// AlertDisastrousEvent is a function to alert service's device.
func (c *Camar) AlertDisastrousEvent(ctx context.Context, disaster datamodel.GeoJSON) error {
	var errs []error
	var wg sync.WaitGroup
	errc := make(chan []error)

	alertMessage, err := c.writer.CreateAlertMessage(disaster, []string{})
	if err != nil {
		return errors.Wrap(err, "AlertDisastrousEvent error on creating alert message")
	}

	devices, err := c.recording.GetDeviceInRadius([]float64{disaster.Geometry.Coordinates[0], disaster.Geometry.Coordinates[1]}, 1.36)
	if err != nil {
		return errors.Wrap(err, "AlertDisastrousEvent error on getting device in radius")
	}

	length := len(devices)
	wg.Add(length)

	for _, device := range devices {
		alertMessage.Token = device.Token
		go c.alerting.SendAlert(alertMessage, errc, &wg)
	}

	for i := 0; i < length; i++ {
		errsx := <-errc
		lenErrsx := len(errsx)
		if lenErrsx > 0 {
			errs = append(errs, errsx...)
		}
	}

	wg.Wait()

	if len(errs) > 0 {
		return errs[0]
	}

	return nil
}

// NewDevice is a function to save new device device for alerting purpose.
func (c *Camar) NewDevice(ctx context.Context, device Device) (Device, error) {
	dev, err := c.recording.GetDevice(device.DeviceID)
	if err != nil {
		device.ID = bson.NewObjectId()
		if err := c.recording.NewDevice(device); err != nil {
			return Device{}, err
		}

		return device, nil
	}

	return dev, nil
}

// UpdateDevice is a function to update device latitude and logitude coordinate.
func (c *Camar) GetDevice(ctx context.Context, deviceid string) (Device, error) {
	device, err := c.recording.GetDevice(deviceid)
	if err != nil {
		return Device{}, err
	}

	return device, nil
}

// UpdateDevice is a function to update device latitude and logitude coordinate.
func (c *Camar) UpdateDevice(ctx context.Context, device Device) (Device, error) {
	dev, err := c.recording.GetDevice(device.DeviceID)
	if err != nil {
		return dev, err
	}
	dev.Location = device.Location
	if err := c.recording.UpdateDevice(dev); err != nil {
		return Device{}, err
	}

	return dev, nil
}
