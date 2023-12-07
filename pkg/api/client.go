package api

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/c9s/requestgen"
)

const (
	defaultHTTPTimeout = time.Second * 15
	defaultBaseURL     = "http://0.0.0.0:8080"
)

type Client struct {
	requestgen.BaseAPIClient
}

func NewClient() *Client {
	u, err := url.Parse(defaultBaseURL)
	if err != nil {
		panic(err)
	}

	return &Client{
		BaseAPIClient: requestgen.BaseAPIClient{
			BaseURL: u,
			HttpClient: &http.Client{
				Timeout: defaultHTTPTimeout,
			},
		},
	}
}

func (c *Client) NewAuthenticatedRequest(ctx context.Context, method, refURL string, params url.Values, payload interface{}) (*http.Request, error) {
	return c.NewRequest(ctx, method, refURL, params, payload)
}
