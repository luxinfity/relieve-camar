package client

import (
	"context"
	"io/ioutil"
	"net/http"
)

type Client interface {
	Do(ctx context.Context, req *http.Request) (*http.Response, []byte, error)
}

type httpClient struct {
	client http.Client
}

func NewClient() *httpClient {
	return &httpClient{
		client: http.Client{},
	}
}

func (c *httpClient) Do(ctx context.Context, req *http.Request) (*http.Response, []byte, error) {
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	resp, err := c.client.Do(req)
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()

	if err != nil {
		return nil, nil, err
	}

	var body []byte
	done := make(chan struct{})
	go func() {
		body, err = ioutil.ReadAll(resp.Body)
		close(done)
	}()

	select {
	case <-ctx.Done():
		err = resp.Body.Close()
		<-done
		if err == nil {
			err = ctx.Err()
		}
	case <-done:
	}

	return resp, body, err
}
