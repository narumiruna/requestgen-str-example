package api

import (
	"github.com/c9s/requestgen"
)

//go:generate requestgen -method GET -url "/" -type RootRequest -responseType RootResponse
type RootRequest struct {
	client requestgen.AuthenticatedAPIClient
}

func (c *Client) NewRootRequest() *RootRequest {
	return &RootRequest{
		client: c,
	}
}

type RootResponse string

func (r *RootResponse) UnmarshalJSON(data []byte) error {
	*r = RootResponse(string(data))
	return nil
}
