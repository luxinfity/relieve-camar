package recorder

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
	"log"

	"github.com/pamungkaski/camar"
	"github.com/pamungkaski/camar/datamodel"
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

func (m *MongoDB) GetEarthquakeList(limit, page int) ([]datamodel.GeoJSON, error) {
	var results []datamodel.GeoJSON

	dbAct := m.session.DB("camar").C("earthquake")
	err := dbAct.Find(nil).Sort("-properties.time").Skip(limit*(page - 1)).Limit(limit).All(&results)
	if err != nil {
		log.Println(err)
		return nil, errors.Wrap(err, "Get List of Recent Earthquake")
	}
	return results, nil
}

// SaveDisaster is a function to save disaster data into database
func (m *MongoDB) SaveDisaster(disaster datamodel.GeoJSON) error {
	dbAct := m.session.DB("camar").C("earthquake")
	err := dbAct.Insert(disaster)
	if err != nil {
		return errors.Wrap(err, "SaveDisaster error")
	}
	return nil
}
// SaveDisaster is a function to save disaster data into database
func (m *MongoDB) SaveInternationalDisaster(disaster datamodel.GeoJSON) error {
	dbAct := m.session.DB("camar").C("earthquake-international")
	err := dbAct.Insert(disaster)
	if err != nil {
		return errors.Wrap(err, "SaveInternationalDisaster error")
	}
	return nil
}

// SaveClient is a function to register client on the alerting service.
func (m *MongoDB) NewDevice(device camar.Device) error {
	dbAct := m.session.DB("camar").C("user")
	err := dbAct.Insert(device)
	if err != nil {
		return errors.Wrap(err, "NewDevice error")
	}
	return nil
}

//
func (m *MongoDB) GetDevice(deviceID string) (camar.Device, error) {
	var dev camar.Device
	query := bson.M{
		"deviceid": deviceID,
	}
	dbAct := m.session.DB("camar").C("user")
	err := dbAct.Find(query).One(&dev)
	if err != nil {
		return dev, errors.Wrap(err, "GetDevice error")
	}
	return dev, nil
}

// UpdateDevice is a function to update device latitude and longitude coordinate.
func (m *MongoDB) UpdateDevice(device camar.Device) error {
	dbAct := m.session.DB("camar").C("user")
	err := dbAct.UpdateId(device.ID, device)
	if err != nil {
		return errors.Wrap(err, "UpdateDevice error")
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
		return results, errors.Wrap(err, "GetDeviceInRadius error")
	}

	return results, nil
}
