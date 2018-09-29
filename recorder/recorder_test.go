package recorder

import (
	"reflect"
	"testing"

	"github.com/globalsign/mgo"
	"github.com/pamungkaski/camar"
	"github.com/pamungkaski/camar/datamodel"
)

func TestNewMongoDB(t *testing.T) {
	type args struct {
		username string
		password string
		host     string
	}
	tests := []struct {
		name    string
		args    args
		want    *MongoDB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMongoDB(tt.args.username, tt.args.password, tt.args.host)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMongoDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMongoDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMongoDB_GetAllEarthquakeData(t *testing.T) {
	type fields struct {
		session *mgo.Session
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MongoDB{
				session: tt.fields.session,
			}
			m.GetAllEarthquakeData()
		})
	}
}

func TestMongoDB_SaveDisaster(t *testing.T) {
	type fields struct {
		session *mgo.Session
	}
	type args struct {
		disaster datamodel.GeoJSON
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MongoDB{
				session: tt.fields.session,
			}
			if err := m.SaveDisaster(tt.args.disaster); (err != nil) != tt.wantErr {
				t.Errorf("MongoDB.SaveDisaster() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMongoDB_NewDevice(t *testing.T) {
	type fields struct {
		session *mgo.Session
	}
	type args struct {
		device camar.Device
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MongoDB{
				session: tt.fields.session,
			}
			if err := m.NewDevice(tt.args.device); (err != nil) != tt.wantErr {
				t.Errorf("MongoDB.NewDevice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMongoDB_UpdateDevice(t *testing.T) {
	type fields struct {
		session *mgo.Session
	}
	type args struct {
		device camar.Device
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MongoDB{
				session: tt.fields.session,
			}
			if err := m.UpdateDevice(tt.args.device); (err != nil) != tt.wantErr {
				t.Errorf("MongoDB.UpdateDevice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMongoDB_GetDeviceInRadius(t *testing.T) {
	type fields struct {
		session *mgo.Session
	}
	type args struct {
		disasterCoordinate []float64
		radius             float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []camar.Device
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MongoDB{
				session: tt.fields.session,
			}
			got, err := m.GetDeviceInRadius(tt.args.disasterCoordinate, tt.args.radius)
			if (err != nil) != tt.wantErr {
				t.Errorf("MongoDB.GetDeviceInRadius() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MongoDB.GetDeviceInRadius() = %v, want %v", got, tt.want)
			}
		})
	}
}
