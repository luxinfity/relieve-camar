// Package camar is main code file for this service.
// This package will contain every single business logic needed for the service
// benCAna Monitoring and AleRting
package camar

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"github.com/pamungkaski/camar/datamodel"
	"github.com/pamungkaski/camar/grabber"
	"github.com/pamungkaski/camar/notifier"
	"github.com/pamungkaski/camar/recorder"
	"github.com/pkg/errors"
)

// DisasterReporter is the business logic contract for camar service.
// The main idea of the interface is to record disaster into database then alert all service's device.
type DisasterReporter interface {
	// ListenTheEarth is a function that Listen to any Earthquake happen.
	// It is the main function of DisasterReporter Interface
	ListenTheEarth() error
	// RecordDisaster is a function to save Disaster into our database.
	RecordDisaster(ctx context.Context, disaster datamodel.CamarQuakeData) error
	//
	GetEarthquakeList(ctx context.Context, limit, page int) ([]datamodel.CamarQuakeData, error)
	// AlertDisastrousEvent is a function to alert service's device.
	GetEarthquake(ctx context.Context, ID string) (datamodel.CamarQuakeData, error)
	AlertDisastrousEvent(ctx context.Context, disaster datamodel.CamarQuakeData) error
	// NewDevice is a function to save new device device for alerting purpose.
	NewDevice(ctx context.Context, device datamodel.Device) (datamodel.Device, error)
	// GetDevice
	GetDevice(ctx context.Context, deviceID string) (datamodel.Device, error)
	// UpdateDevice is a function to update device latitude and longitude coordinate.
	UpdateDevice(ctx context.Context, device datamodel.Device) (datamodel.Device, error)
	//
	GetAllDevice(ctx context.Context) ([]datamodel.Device, error)
}

// Camar is the main struct of the service.
// Camar implements DisasterReporter interface.
// It contains alerting and recorder interface implementation.
type Camar struct {
	grabber  grabber.ResourceGrabber
	recorder recorder.Recorder
	notifer  notifier.Notifier
}

//NewDisasterReporter is a function that creates an instance of DisasterReporter.
func NewDisasterReporter(grabber grabber.ResourceGrabber, recorder recorder.Recorder, notifier notifier.Notifier) DisasterReporter {
	return &Camar{
		recorder:     recorder,
		grabber:       grabber,
		notifer: notifier,
	}
}

// ListenTheEarth is a function that Listen to any Earthquake happen.
// It is the main function of DisasterReporter Interface
func (c *Camar) ListenTheEarth() error {
	latest, err := c.grabber.GetEarthquakes()
	if err != nil {
		return err
	}
	quakes, err := c.recorder.GetEarthquakeList(1, 1)

	if !c.verifyQuake(latest[0], quakes[0]) {
		if err := c.RecordDisaster(context.Background(), latest[0]); err != nil {
			return err
		}
		if err := c.AlertDisastrousEvent(context.Background(), latest[0]); err != nil {
			return err
		}
	}

	return nil
}

func (c *Camar) verifyQuake(first, second datamodel.CamarQuakeData) bool {
	if first.Mag != second.Mag {
		return false
	}
	if first.Time == second.Time {
		return false
	}
	if first.Location.Coordinates[0] == first.Location.Coordinates[0] {
		return false
	}

	return first.Location.Coordinates[1] == first.Location.Coordinates[1]
}

func (c *Camar) GetEarthquakeList(ctx context.Context, limit, page int) ([]datamodel.CamarQuakeData, error) {
	var snapList []datamodel.CamarQuakeData

	snapList, err := c.recorder.GetEarthquakeList(limit, page)
	if err != nil {
		return nil, err
	}

	return snapList, nil
}

// RecordDisaster is a function to save Disaster into our database.
func (c *Camar) RecordDisaster(ctx context.Context, disaster datamodel.CamarQuakeData) error {
	return c.recorder.SaveDisaster(disaster)
}
// GetEarthquake return CamarQuakeData with specified ID.
func (c *Camar) GetEarthquake(ctx context.Context, ID string) (datamodel.CamarQuakeData, error) {
	return c.recorder.GetEarthquake(ID)
}
// AlertDisastrousEvent is a function to alert service's device.
func (c *Camar) AlertDisastrousEvent(ctx context.Context, disaster datamodel.CamarQuakeData) error {
	var errs []error
	errc := make(chan []error)

	devices, err := c.recorder.GetDeviceInRadius([]float64{disaster.Location.Coordinates[0], disaster.Location.Coordinates[1]}, 1.36)
	if err != nil {
		return errors.Wrap(err, "AlertDisastrousEvent error on getting device in radius")
	}

	length := len(devices)

	for _, device := range devices {
		go c.notifer.SendAlert(disaster, device.Token, errc)
	}

	for i := 0; i < length; i++ {
		errsx := <-errc
		lenErrsx := len(errsx)
		if lenErrsx > 0 {
			errs = append(errs, errsx...)
		}
	}

	if len(errs) > 0 {
		return errs[0]
	}

	return nil
}

// NewDevice is a function to save new device device for alerting purpose.
func (c *Camar) NewDevice(ctx context.Context, device datamodel.Device) (datamodel.Device, error) {
	dev, err := c.recorder.GetDevice(device.DeviceID)
	if err != nil {
		device.ID = bson.NewObjectId()
		if err := c.recorder.NewDevice(device); err != nil {
			return datamodel.Device{}, err
		}

		return device, nil
	}

	return dev, nil
}

// UpdateDevice is a function to update device latitude and logitude coordinate.
func (c *Camar) GetDevice(ctx context.Context, deviceid string) (datamodel.Device, error) {
	device, err := c.recorder.GetDevice(deviceid)
	if err != nil {
		return datamodel.Device{}, err
	}

	return device, nil
}

// UpdateDevice is a function to update device latitude and logitude coordinate.
func (c *Camar) GetAllDevice(ctx context.Context) ([]datamodel.Device, error) {
	devices, err := c.recorder.GetAllDevice()
	if err != nil {
		return nil, err
	}

	return devices, nil
}

// UpdateDevice is a function to update device latitude and logitude coordinate.
func (c *Camar) UpdateDevice(ctx context.Context, device datamodel.Device) (datamodel.Device, error) {
	dev, err := c.recorder.GetDevice(device.DeviceID)
	if err != nil {
		return dev, err
	}
	dev.Location = device.Location
	if err := c.recorder.UpdateDevice(dev); err != nil {
		return datamodel.Device{}, err
	}

	return dev, nil
}
