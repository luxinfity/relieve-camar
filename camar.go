// Package camar is main code file for this service.
// This package will contain every single business logic needed for the service
// benCAna Monitoring and AleRting
package camar

import (
	"context"

	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

// DisasterReporter is the business logic contract for camar service.
// The main idea of the interface is to record disaster into database then alert all service's client.
type DisasterReporter interface {
	// RecordDisaster is a function to save Disaster into our database.
	RecordDisaster(ctx context.Context, disaster Disaster) (Disaster, error)
	// AlertDisastrousEvent is a function to alert service's client.
	AlertDisastrousEvent(ctx context.Context, disaster Disaster) error
	// NewClient is a function to save new client device for alerting purpose.
	NewClient(ctx context.Context, client Client) (Client, error)
	// UpdateClient is a function to update client latitude and logitude coordinate.
	UpdateClient(ctx context.Context, client Client) (Client, error)
}

// Alerting is the bussines logic contract for alerting service.
// the main idea is to send alert to all client.
type Alerting interface {
	// SendAlert is a function to send Disastrous Event alert to specific Client using the alerting service.
	SendAlert(alert string, client Client) error
}

// Recording is the business logic contract for saving data.
type Recording interface {
	// SaveDisaster is a function to save disaster data into database
	SaveDisaster(disaster Disaster) error
	// SaveClient is a function to register client on the alerting service.
	NewClient(client Client) error
	// UpdateClient is a function to update client latitude and longitude coordinate.
	UpdateClient(client Client) error
	// GetClientInRadius is a function to get all Client data inside the Disastrous Zone Radius.
	GetClientInRadius(coordinate Coordinate, radius float64) ([]Client, error)
}

type AlertWritter interface {
	// CreateAlertMessage is a function to create alert message based on th disaster event that currently occurs.
	CreateAlertMessage(disaster Disaster) (string, error)
}

// Coordinate is the struct to save exact location of data on earth.
type Coordinate struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// Disaster is the struct for each single disastrous event recorded for the service.
// It contains disaster details.
type Disaster struct {
	ID                  bson.ObjectId `bson:"_id" json:"id"`
	Type                string        `json:"type"`
	Coordinate          Coordinate    `json:"coordinate"`
	DangerousZoneRadius float64       `json:"dangerous_zone_radius"`
	Detail              interface{}   `json:"detail"`
}

// Client is the struct for each device conected to service.
// It save the ClientID for alerting purpose.
// In client side, it is an automatics Device Regist on first start.
// The Latitude and Longitude are update-able.
type Client struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Coordinate Coordinate    `json:"coordinate"`
}

// Camar is the main struct of the service.
// Camar implements DisasterReporter interface.
// It contains alerting and recording interface implementation.
type Camar struct {
	alerting  Alerting
	recording Recording
	writer    AlertWritter
}

// RecordDisaster is a function to save Disaster into our database.
func (c *Camar) RecordDisaster(ctx context.Context, disaster Disaster) (Disaster, error) {
	disaster.ID = bson.NewObjectId()

	if err := c.recording.SaveDisaster(disaster); err != nil {
		return Disaster{}, errors.Wrap(err, "RecordDisaster error on saving disaster")
	}

	return disaster, nil
}

// AlertDisastrousEvent is a function to alert service's client.
func (c *Camar) AlertDisastrousEvent(ctx context.Context, disaster Disaster) error {
	alertMessage, err := c.writer.CreateAlertMessage(disaster)
	if err != nil {
		return errors.Wrap(err, "AlertDisastrousEvent error on creating alert message")
	}

	clients, err := c.recording.GetClientInRadius(disaster.Coordinate, disaster.DangerousZoneRadius)
	if err != nil {
		return errors.Wrap(err, "AlertDisastrousEvent error on getting client in radius")
	}

	for _, client := range clients {
		if err = c.alerting.SendAlert(alertMessage, client); err != nil {
			return errors.Wrap(err, "AlertDisastrousEvent error on sending alert message")
		}
	}

	return nil
}

// NewClient is a function to save new client device for alerting purpose.
func (c *Camar) NewClient(ctx context.Context, client Client) (Client, error) {
	client.ID = bson.NewObjectId()

	if err := c.recording.NewClient(client); err != nil {
		return Client{}, errors.Wrap(err, "NewClient error on creating new client")
	}

	return client, nil
}

// UpdateClient is a function to update client latitude and logitude coordinate.
func (c *Camar) UpdateClient(ctx context.Context, client Client) (Client, error) {
	if err := c.recording.UpdateClient(client); err != nil {
		return Client{}, errors.Wrap(err, "NewClient error on creating new client")
	}

	return client, nil
}
