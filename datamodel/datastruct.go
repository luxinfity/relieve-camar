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

type Event struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Location struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"location"`
	Time  int64   `json:"time"`
	TimeArrived int64 `json:"time_arrived"`
	EventType string `json:"event_type"`
	Source string `json:"source"`
	EventDetail interface{} `json:"event_detail"`
}

type CamarQuakeData struct {
	Title    string        `json:"title"`
	Mag   float64 `json:"mag"`
	Depth float64 `json:"depth"`
	MMI string  `json:"mmi"`
}

type BMKGQuakes struct {
	XMLName xml.Name `xml:"Infogempa"`
	Text    string   `xml:",chardata"`
	Gempa   []BMKGQuakeData `xml:"Gempa"`
}

type BMKGQuakeData struct {
	Text    string `xml:",chardata"`
	Tanggal string `xml:"Tanggal"`
	Point   struct {
		Text        string `xml:",chardata"`
		Coordinates string `xml:"coordinates"`
	} `xml:"point"`
	Posisi     string `xml:"Posisi"`
	Magnitude  string `xml:"Magnitude"`
	Kedalaman  string `xml:"Kedalaman"`
	Symbol     string `xml:"_symbol"`
	Keterangan string `xml:"Keterangan"`
	Dirasakan  string `xml:"Dirasakan"`
}
