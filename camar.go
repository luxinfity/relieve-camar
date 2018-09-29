// Package camar is main code file for this service.
// This package will contain every single business logic needed for the service
// benCAna Monitoring and AleRting
package camar

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"github.com/pamungkaski/camar/datamodel"
	"github.com/pkg/errors"
)

// DisasterReporter is the business logic contract for camar service.
// The main idea of the interface is to record disaster into database then alert all service's device.
type DisasterReporter interface {
	// ListenTheEarth is a function that Listen to any Earthquake happen.
	// It is the main function of DisasterReporter Interface
	ListenTheEarth(ctx context.Context)
	// RecordDisaster is a function to save Disaster into our database.
	RecordDisaster(ctx context.Context, disaster datamodel.GeoJSON) (datamodel.GeoJSON, error)
	// AlertDisastrousEvent is a function to alert service's device.
	AlertDisastrousEvent(ctx context.Context, disaster datamodel.GeoJSON) error
	// NewDevice is a function to save new device device for alerting purpose.
	NewDevice(ctx context.Context, device Device) (Device, error)
	// UpdateDevice is a function to update device latitude and longitude coordinate.
	UpdateDevice(ctx context.Context, device Device) (Device, error)
}

// ResourceGrabber is the bussiness logic contract for getting earthquake data.
type ResourceGrabber interface {
	// GetEarthQuakeData is a function to to retrieve Earthquake detailed data.
	GetEarthquakeData(eventID string)
}

// Alerting is the bussiness logic contract for alerting service.
// the main idea is to send alert to all device.
type Alerting interface {
	// SendAlert is a function to send Disastrous Event alert to specific Device using the alerting service.
	SendAlert(alert string, device Device) error
}

// Recorder is the business logic contract for saving data.
type Recorder interface {
	// SaveDisaster is a function to save disaster data into database
	SaveDisaster(disaster datamodel.GeoJSON) error
	// SaveDevice is a function to register device on the alerting service.
	NewDevice(device Device) error
	// UpdateDevice is a function to update device latitude and longitude coordinate.
	UpdateDevice(device Device) error
	// GetDeviceInRadius is a function to get all Device data inside the Disastrous Zone Radius.
	GetDeviceInRadius(disasterCoordinate []float64, radius float64) ([]Device, error)
}

type AlertWritter interface {
	// CreateAlertMessage is a function to create alert message based on th disaster event that currently occurs.
	CreateAlertMessage(disaster datamodel.GeoJSON) (string, error)
}



// Device is the struct for each device conected to service.
// It save the DeviceID for alerting purpose.
// In device side, it is an automatics Device Regist on first start.
// The Latitude and Longitude are update-able.
type Device struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	DeviceID string `json:"device_id"`
	Location struct{
		Type string `json:"type"`
		Coordinates []float64    `json:"coordinates"`
	} `json:"location"`
}

// Camar is the main struct of the service.
// Camar implements DisasterReporter interface.
// It contains alerting and recording interface implementation.
type Camar struct {
	alerting  Alerting
	recording Recorder
	writer    AlertWritter
	grabber   ResourceGrabber
}

// ListenTheEarth is a function that Listen to any Earthquake happen.
// It is the main function of DisasterReporter Interface
func (c *Camar) ListenTheEarth(ctx context.Context) {

}

// RecordDisaster is a function to save Disaster into our database.
func (c *Camar) RecordDisaster(ctx context.Context, disaster datamodel.GeoJSON) (datamodel.GeoJSON, error) {
	disaster.BsonID = bson.NewObjectId()

	if err := c.recording.SaveDisaster(disaster); err != nil {
		return datamodel.GeoJSON{}, errors.Wrap(err, "RecordDisaster error on saving disaster")
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
	device.ID = bson.NewObjectId()

	if err := c.recording.NewDevice(device); err != nil {
		return Device{}, errors.Wrap(err, "NewDevice error on creating new device")
	}

	return device, nil
}

// UpdateDevice is a function to update device latitude and logitude coordinate.
func (c *Camar) UpdateDevice(ctx context.Context, device Device) (Device, error) {
	if err := c.recording.UpdateDevice(device); err != nil {
		return Device{}, errors.Wrap(err, "NewDevice error on creating new device")
	}

	return device, nil
}
