// Package camar is main code file for this service.
// This package will contain every single business logic needed for the service
// benCAna Monitoring and AleRting
package camar

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"

	"github.com/pamungkaski/camar/datamodel"
	"github.com/dghubble/go-twitter/twitter"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"log"
	"strings"
	"net/http"
)

// DisasterReporter is the business logic contract for camar service.
// The main idea of the interface is to record disaster into database then alert all service's device.
type DisasterReporter interface {
	// ListenTheEarth is a function that Listen to any Earthquake happen.
	// It is the main function of DisasterReporter Interface
	ListenTheEarth()
	// RecordDisaster is a function to save Disaster into our database.
	RecordDisaster(ctx context.Context, disaster datamodel.GeoJSON) (datamodel.GeoJSON, error)
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
}

// Recorder is the business logic contract for saving data.
type Recorder interface {
	// SaveDisaster is a function to save disaster data into database
	SaveDisaster(disaster datamodel.GeoJSON) error
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
	SendAlert(alert string, device Device) error
}

// AlerWritter is the business logic contract for alert message writter.
type AlertWritter interface {
	// CreateAlertMessage is a function to create alert message based on th disaster event that currently occurs.
	CreateAlertMessage(disaster datamodel.GeoJSON) (string, error)
}

// Device is the struct for each device conected to service.
// It save the DeviceID for alerting purpose.
// In device side, it is an automatics Device Regist on first start.
// The Latitude and Longitude are update-able.
type Device struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	DeviceID string        `json:"device_id"`
	Location struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"location"`
}

// Camar is the main struct of the service.
// Camar implements DisasterReporter interface.
// It contains alerting and recording interface implementation.
type Camar struct {
	listener *twitter.Client
	grabber   ResourceGrabber
	recording Recorder
	alerting  Alerting
	writer    AlertWritter
}
// NewDisasterReporter is a function that creates an instance of DisasterReporter.
func NewDisasterReporter(client *twitter.Client, recorder Recorder, grabber ResourceGrabber) DisasterReporter {
	return &Camar{
		listener: client,
		recording: recorder,
		grabber: grabber,
	}
}

// ListenTheEarth is a function that Listen to any Earthquake happen.
// It is the main function of DisasterReporter Interface
func (c *Camar) ListenTheEarth() {
	fmt.Println("Starting Stream...")

	// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		text := tweet.Text
		splitted := strings.Split(text, " ")
		textLength := len(splitted)

		id, err := c.getEarthquakeEventID(splitted[textLength - 1])
		if err != nil {
			fmt.Println(err)
		}

		data, err := c.grabber.GetEarthquakeData(id)
		if err != nil {
			fmt.Println(err)
		}

		data, err = c.RecordDisaster(context.Background(), data)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(data.Properties.Title)
	}

	// FILTER
	params := &twitter.StreamFilterParams{
		Follow:[]string{"94119095"},
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
func (c *Camar) getEarthquakeEventID(link string) (string, error){
	resp, err := http.Get(link)
	if err != nil {
		return "", errors.Wrap(err, "get earthquake ID error")
	}

	// Your magic function. The Request in the Response is the last URL the
	// client tried to access.
	finalURL := resp.Request.URL.String()
	split := strings.Split(finalURL, "/")

	return split[len(split) - 1], nil
}

// RecordDisaster is a function to save Disaster into our database.
func (c *Camar) RecordDisaster(ctx context.Context, disaster datamodel.GeoJSON) (datamodel.GeoJSON, error) {
	disaster.BsonID = bson.NewObjectId()

	if err := c.recording.SaveDisaster(disaster); err != nil {
		return datamodel.GeoJSON{}, err
	}

	return disaster, nil
}

// AlertDisastrousEvent is a function to alert service's device.
func (c *Camar) AlertDisastrousEvent(ctx context.Context, disaster datamodel.GeoJSON) error {
	alertMessage, err := c.writer.CreateAlertMessage(disaster)
	if err != nil {
		return errors.Wrap(err, "AlertDisastrousEvent error on creating alert message")
	}

	devices, err := c.recording.GetDeviceInRadius(disaster.Geometry.Coordinates, 150)
	if err != nil {
		return errors.Wrap(err, "AlertDisastrousEvent error on getting device in radius")
	}

	for _, device := range devices {
		if err = c.alerting.SendAlert(alertMessage, device); err != nil {
			return errors.Wrap(err, "AlertDisastrousEvent error on sending alert message")
		}
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
