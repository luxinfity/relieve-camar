package grabber

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/pamungkaski/camar/client"
	"github.com/pamungkaski/camar/datamodel"
)

func TestUSGS_GetEarthquakeData(t *testing.T) {
	type fields struct {
		endpoint string
		api      client.Client
	}
	type args struct {
		eventID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    datamodel.GeoJSON
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &USGS{
				endpoint: tt.fields.endpoint,
				api:      tt.fields.api,
			}
			got, err := u.GetEarthquakeData(tt.args.eventID)
			if (err != nil) != tt.wantErr {
				t.Errorf("USGS.GetEarthquakeData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("USGS.GetEarthquakeData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGrabber(t *testing.T) {
	type args struct {
		endpoint string
		api      client.Client
	}
	tests := []struct {
		name string
		args args
		want *USGS
	}{
		// TODO: Add test cases.
		{
			"Normal",
			args{
				"localhost",
				client.NewClientMock(),
			},
			&USGS{
				"localhost",
				client.NewClientMock(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGrabber(tt.args.endpoint, tt.args.api); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGrabber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUSGS_buildUSGSRequest(t *testing.T) {
	type fields struct {
		endpoint string
		api      client.Client
	}
	type args struct {
		eventID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &USGS{
				endpoint: tt.fields.endpoint,
				api:      tt.fields.api,
			}
			got, err := u.buildUSGSRequest(tt.args.eventID)
			if (err != nil) != tt.wantErr {
				t.Errorf("USGS.buildUSGSRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("USGS.buildUSGSRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

