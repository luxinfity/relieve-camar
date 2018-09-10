package camar

import (
	"context"
	"reflect"
	"testing"
)

func TestCamar_ListenTheEarth(t *testing.T) {
	type fields struct {
		alerting  Alerting
		recording Recorder
		writer    AlertWritter
		grabber   ResourceGrabber
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Camar{
				alerting:  tt.fields.alerting,
				recording: tt.fields.recording,
				writer:    tt.fields.writer,
				grabber:   tt.fields.grabber,
			}
			c.ListenTheEarth(tt.args.ctx)
		})
	}
}

func TestCamar_RecordDisaster(t *testing.T) {
	type fields struct {
		alerting  Alerting
		recording Recorder
		writer    AlertWritter
		grabber   ResourceGrabber
	}
	type args struct {
		ctx      context.Context
		disaster Disaster
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Disaster
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Camar{
				alerting:  tt.fields.alerting,
				recording: tt.fields.recording,
				writer:    tt.fields.writer,
				grabber:   tt.fields.grabber,
			}
			got, err := c.RecordDisaster(tt.args.ctx, tt.args.disaster)
			if (err != nil) != tt.wantErr {
				t.Errorf("Camar.RecordDisaster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Camar.RecordDisaster() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCamar_AlertDisastrousEvent(t *testing.T) {
	type fields struct {
		alerting  Alerting
		recording Recorder
		writer    AlertWritter
		grabber   ResourceGrabber
	}
	type args struct {
		ctx      context.Context
		disaster Disaster
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
			c := &Camar{
				alerting:  tt.fields.alerting,
				recording: tt.fields.recording,
				writer:    tt.fields.writer,
				grabber:   tt.fields.grabber,
			}
			if err := c.AlertDisastrousEvent(tt.args.ctx, tt.args.disaster); (err != nil) != tt.wantErr {
				t.Errorf("Camar.AlertDisastrousEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCamar_NewClient(t *testing.T) {
	type fields struct {
		alerting  Alerting
		recording Recorder
		writer    AlertWritter
		grabber   ResourceGrabber
	}
	type args struct {
		ctx    context.Context
		client Client
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Camar{
				alerting:  tt.fields.alerting,
				recording: tt.fields.recording,
				writer:    tt.fields.writer,
				grabber:   tt.fields.grabber,
			}
			got, err := c.NewClient(tt.args.ctx, tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("Camar.NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Camar.NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCamar_UpdateClient(t *testing.T) {
	type fields struct {
		alerting  Alerting
		recording Recorder
		writer    AlertWritter
		grabber   ResourceGrabber
	}
	type args struct {
		ctx    context.Context
		client Client
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Camar{
				alerting:  tt.fields.alerting,
				recording: tt.fields.recording,
				writer:    tt.fields.writer,
				grabber:   tt.fields.grabber,
			}
			got, err := c.UpdateClient(tt.args.ctx, tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("Camar.UpdateClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Camar.UpdateClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
