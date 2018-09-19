package client

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name string
		want *httpClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_httpClient_Do(t *testing.T) {
	type fields struct {
		client http.Client
	}
	type args struct {
		ctx context.Context
		req *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		want1   []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &httpClient{
				client: tt.fields.client,
			}
			got, got1, err := c.Do(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("httpClient.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpClient.Do() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("httpClient.Do() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
