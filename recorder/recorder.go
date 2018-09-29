package recorder

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/pamungkaski/camar/datamodel"
	"github.com/pkg/errors"
	"log"
	"github.com/pamungkaski/camar"
	"github.com/globalsign/mgo/bson"
)

type MongoDB struct {
	session *mgo.Session
}

func NewMongoDB(username, password, host string) (*MongoDB, error) {
	mongoCredential := &mgo.Credential{
		Username:    username,
		Password:    password,
		Source:      "camar",
		ServiceHost: host,
	}
	mg, err := mgo.Dial(host)
	if err != nil {
		return nil, errors.Wrap(err, "MongoDB Dial error")
	}

	mg.SetMode(mgo.Monotonic, true)

	err = mg.Login(mongoCredential)
	if err != nil {
		return nil, errors.Wrap(err, "MongoDB login error")
	}

	return &MongoDB{
		session: mg,
	}, nil
}

func (m *MongoDB) GetAllEarthquakeData() {
	var results []interface{}

	dbAct := m.session.DB("camar").C("earthquake")
	err := dbAct.Find(nil).All(&results)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Results All :", results)
	}
}

// SaveDisaster is a function to save disaster data into database
func (m *MongoDB) SaveDisaster(disaster datamodel.GeoJSON) error {
	dbAct := m.session.DB("camar").C("earthquake")
	err := dbAct.Insert(disaster)
	if err != nil {
		return err
	}
	return nil
}

// SaveClient is a function to register client on the alerting service.
func (m *MongoDB) NewDevice(device camar.Device) error {
	dbAct := m.session.DB("camar").C("user")
	err := dbAct.Insert(device)
	if err != nil {
		return err
	}
	return nil
}

// UpdateDevice is a function to update device latitude and longitude coordinate.
func (m *MongoDB) UpdateDevice(device camar.Device) error {
	dbAct := m.session.DB("camar").C("user")
	err := dbAct.UpdateId(device.ID, device)
	if err != nil {
		return err
	}
	return nil
}

// GetDeviceInRadius is a function to get all Device data inside the Disastrous Zone Radius.
func (m *MongoDB) GetDeviceInRadius(disasterCoordinate []float64, radius float64) ([]camar.Device, error) {
	var results []camar.Device
	var center []interface{}
	center = append(center, disasterCoordinate)
	center = append(center, radius)
	query := bson.M{
		"location": bson.M{
			"$geoWithin": bson.M{
				"$center": center,
			},
		},
	}

	dbAct := m.session.DB("camar").C("user")
	err := dbAct.Find(query).All(&results)
	if err != nil {
		return results, err
	}

	return results, nil
}

