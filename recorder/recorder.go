package recorder

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pamungkaski/camar/datamodel"
	"github.com/pkg/errors"
	"strings"
)

// Recorder is the business logic contract for saving data.
type Recorder interface {
	// SaveDisaster is a function to save disaster data into database
	SaveDisaster(disaster datamodel.CamarQuakeData) error
	// SaveDisaster is a function to save disaster data into database
	SaveInternationalDisaster(disaster datamodel.CamarQuakeData) error
	//
	GetEarthquakeList(limit, page int) ([]datamodel.CamarQuakeData, int, error)
	//
	GetEarthquake(ID string) (datamodel.CamarQuakeData, error)
	// SaveDevice is a function to register device on the alerting service.
	NewDevice(device datamodel.Device) error
	//
	GetDevice(deviceID string) (datamodel.Device, error)
	// UpdateDevice is a function to update device latitude and longitude coordinate.
	UpdateDevice(device datamodel.Device) error
	// GetDeviceInRadius is a function to get all Device data inside the Disastrous Zone Radius.
	GetDeviceInRadius(disasterCoordinate []float64, radius float64) ([]datamodel.Device, error)
	//
	GetAllDevice() ([]datamodel.Device, error)
	//
	NewEvent(event datamodel.Event) (error)
	// GetEvent
	GetEvent(eventID string) (datamodel.Event, error)
	// UpdateEvent is a function to update event latitude and longitude coordinate.
	UpdateEvent(event datamodel.Event) (error)
	//
	DeleteEvent(event datamodel.Event) (error)
	//
	GetAllEvent(limit, page int, eventType string) ([]datamodel.Event, int,error)
}

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

func (m *MongoDB) GetEarthquakeList(limit, page int) ([]datamodel.CamarQuakeData, int, error) {
	var results []datamodel.CamarQuakeData

	dbAct := m.session.DB("camar").C("earthquake")
	err := dbAct.Find(nil).Sort("-time").Skip(limit * (page - 1)).Limit(limit).All(&results)
	if err != nil {
		log.Println(err)
		return nil, 0, errors.Wrap(err, "Get List of Recent Earthquake")
	}

	c, err := dbAct.Find(nil).Count()
	if err != nil {
		log.Println(err)
		return nil, 0, errors.Wrap(err, "Get List of Recent Earthquake")
	}

	return results, c, nil
}

func (m *MongoDB) GetEarthquake(ID string) (datamodel.CamarQuakeData, error) {
	quake := datamodel.CamarQuakeData{}
	id := bson.ObjectIdHex(ID)
	query := bson.M{
		"_id": id,
	}
	dbAct := m.session.DB("camar").C("earthquake")
	err := dbAct.Find(query).One(&quake)
	if err != nil {
		return quake, errors.Wrap(err, "Get Earthquake error")
	}
	return quake, nil
}

// SaveDisaster is a function to save disaster data into database
func (m *MongoDB) SaveDisaster(disaster datamodel.CamarQuakeData) error {
	dbAct := m.session.DB("camar").C("earthquake")
	err := dbAct.Insert(disaster)
	if err != nil {
		return errors.Wrap(err, "SaveDisaster error")
	}
	return nil
}

// SaveDisaster is a function to save disaster data into database
func (m *MongoDB) SaveInternationalDisaster(disaster datamodel.CamarQuakeData) error {
	dbAct := m.session.DB("camar").C("earthquake-international")
	err := dbAct.Insert(disaster)
	if err != nil {
		return errors.Wrap(err, "SaveInternationalDisaster error")
	}
	return nil
}

// SaveClient is a function to register client on the alerting service.
func (m *MongoDB) NewDevice(device datamodel.Device) error {
	dbAct := m.session.DB("camar").C("user")
	err := dbAct.Insert(device)
	if err != nil {
		return errors.Wrap(err, "NewDevice error")
	}
	return nil
}

//
func (m *MongoDB) GetDevice(deviceID string) (datamodel.Device, error) {
	var dev datamodel.Device
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

//
func (m *MongoDB) GetAllDevice() ([]datamodel.Device, error) {
	var dev []datamodel.Device
	dbAct := m.session.DB("camar").C("user")
	err := dbAct.Find(nil).All(&dev)
	if err != nil {
		return dev, errors.Wrap(err, "GetAllDevice error")
	}
	return dev, nil
}

// UpdateDevice is a function to update device latitude and longitude coordinate.
func (m *MongoDB) UpdateDevice(device datamodel.Device) error {
	dbAct := m.session.DB("camar").C("user")
	err := dbAct.UpdateId(device.ID, device)
	if err != nil {
		return errors.Wrap(err, "UpdateDevice error")
	}
	return nil
}

// GetDeviceInRadius is a function to get all Device data inside the Disastrous Zone Radius.
func (m *MongoDB) GetDeviceInRadius(disasterCoordinate []float64, radius float64) ([]datamodel.Device, error) {
	var results []datamodel.Device
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


func (m *MongoDB) NewEvent(event datamodel.Event) (error) {
	dbAct := m.session.DB("camar").C("event")
	err := dbAct.Insert(&event)
	if err != nil {
		return errors.Wrap(err, "NewEvent error")
	}
	return nil
}

// GetEvent
func (m *MongoDB) GetEvent(eventID string) (datamodel.Event, error) {
	var eve datamodel.Event
	query := bson.M{
		"_id": bson.ObjectIdHex(eventID),
	}
	dbAct := m.session.DB("camar").C("event")
	err := dbAct.Find(query).One(&eve)
	if err != nil {
		return eve, errors.Wrap(err, "GetEvent error")
	}
	return eve, nil
}
// UpdateEvent is a function to update event latitude and longitude coordinate.
func (m *MongoDB) UpdateEvent(event datamodel.Event) (error) {
	dbAct := m.session.DB("camar").C("event")
	err := dbAct.UpdateId(event.ID, event)
	if err != nil {
		return errors.Wrap(err, "UpdateEvent error")
	}
	return nil
}
//
func (m *MongoDB) DeleteEvent(event datamodel.Event) (error) {
	dbAct := m.session.DB("camar").C("event")
	err := dbAct.RemoveId(event.ID)
	if err != nil {
		return errors.Wrap(err, "DeleteEvent error")
	}
	return nil
}
//
func (m *MongoDB) GetAllEvent(limit, page int, eventType string) ([]datamodel.Event, int, error) {
	var results []datamodel.Event
	var query bson.M

	if eventType != "" {
		query = bson.M{
			"eventtype": strings.Title(eventType),
		}
	}
	dbAct := m.session.DB("camar").C("event")
	err := dbAct.Find(query).Sort("-time").Skip(limit * (page - 1)).Limit(limit).All(&results)
	if err != nil {
		log.Println(err)
		return nil, 0,errors.Wrap(err, "Get List of Recent Event")
	}

	c, err := dbAct.Find(nil).Count()
	if err != nil {
		log.Println(err)
		return nil, 0,errors.Wrap(err, "Get List of Recent Event")
	}

	return results, c, nil
}