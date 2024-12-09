// This file provides a client with methods as well as functions to interact with the HTTP API.
//
// Code generated by github.com/MarkRosemaker DO NOT EDIT.

package toggl

import (
	"net/http"
	"net/url"

	"github.com/go-json-experiment/json"
)

var (
	baseURL = &url.URL{
		Host:   "",
		Path:   "TODO",
		Scheme: "",
	}

	jsonOpts = json.JoinOptions(
		json.RejectUnknownMembers(true))
)

// Client conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The HTTP client to use for requests.
	cli *http.Client
}

// NewClient creates a new Client.
func NewClient() (*Client, error) {
	return &Client{cli: http.DefaultClient}, nil
}
