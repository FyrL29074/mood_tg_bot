package api

import (
	"net/http"
	"time"
)

var client *http.Client

type RetryTransport struct {
	MaxRetries int
	WaitTime   time.Duration
}

func (t *RetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	defaultTransport := http.DefaultTransport

	var resp *http.Response
	var err error

	for range t.MaxRetries {
		resp, err := defaultTransport.RoundTrip(req)

		if err != nil && resp.StatusCode < 400 {
			return resp, nil
		}

		time.Sleep(t.WaitTime)
	}

	return resp, err
}

func initApi() {
	client = &http.Client{
		Timeout: 60 * time.Second,
		Transport: &RetryTransport{
			MaxRetries: 3,
			WaitTime:   2 * time.Second,
		},
	}
}
