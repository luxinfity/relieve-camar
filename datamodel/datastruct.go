package datamodel

import (
	"encoding/xml"
	"github.com/globalsign/mgo/bson"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
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

type CamarQuakeData struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Title    string        `json:"title"`
	Location struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"location"`
	Mag     float64 `json:"mag"`
	Place   string  `json:"place"`
	Time    int64   `json:"time"`
}

type BMKGQuakes struct {
	XMLName xml.Name        `xml:"Infogempa"`
	Text    string          `xml:",chardata"`
	Gempa   []BMKGQuakeData `xml:"gempa"`
}

type BMKGQuakeData struct {
	Text    string `xml:",chardata"`
	Tanggal string `xml:"Tanggal"`
	Jam     string `xml:"Jam"`
	Point   struct {
		Type        string `xml:",chardata"`
		Coordinates string `xml:"coordinates"`
	} `xml:"point"`
	Lintang   string `xml:"Lintang"`
	Bujur     string `xml:"Bujur"`
	Magnitude string `xml:"Magnitude"`
	Kedalaman string `xml:"Kedalaman"`
	Symbol    string `xml:"_symbol"`
	Wilayah   string `xml:"Wilayah"`
}