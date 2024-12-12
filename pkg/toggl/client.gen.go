// This file provides a client with methods as well as functions to interact with the HTTP API.
//
// Code generated by github.com/MarkRosemaker DO NOT EDIT.

package toggl

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/MarkRosemaker/jsonutil"
	"github.com/go-api-libs/api"
	"github.com/go-json-experiment/json"
)

const (
	userAgent = "TogglGoAPILibrary/9 (https://github.com/go-api-libs/toggl)"
)

var (
	baseURL = &url.URL{
		Host:   "api.track.toggl.com",
		Path:   "/api/v9",
		Scheme: "https",
	}

	jsonOpts = json.JoinOptions(
		json.RejectUnknownMembers(true),
		json.WithMarshalers(json.MarshalFuncV2(jsonutil.URLMarshal)),
		json.WithUnmarshalers(json.UnmarshalFuncV2(jsonutil.URLUnmarshal)))
)

// Client conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The HTTP client to use for requests.
	cli *http.Client
	// The authorization header to use.
	authHeader string
}

// NewClient creates a new Client.
func NewClient() (*Client, error) {
	return &Client{cli: http.DefaultClient}, nil
}

// Returns details for the current user.
//
//	GET /me
func (c *Client) GetMe(ctx context.Context) (*User, error) {
	return GetMe[User](ctx, c)
}

// Returns details for the current user.
// You can define a custom result to unmarshal the response into.
//
//	GET /me
func GetMe[R any](ctx context.Context, c *Client) (*R, error) {
	u := baseURL.JoinPath("/me")
	req := (&http.Request{
		Header: http.Header{
			"Authorization": []string{c.authHeader},
			"User-Agent":    []string{userAgent},
		},
		Host:       u.Host,
		Method:     http.MethodGet,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		URL:        u,
	}).WithContext(ctx)

	rsp, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	switch rsp.StatusCode {
	case http.StatusOK:
		// TODO
		switch mt, _, _ := strings.Cut(rsp.Header.Get("Content-Type"), ";"); mt {
		case "application/json":
			var out R
			if err := json.UnmarshalRead(rsp.Body, &out, jsonOpts); err != nil {
				return nil, api.WrapDecodingError(rsp, err)
			}

			return &out, nil
		default:
			return nil, api.NewErrUnknownContentType(rsp)
		}
	case http.StatusUnauthorized:
		// User is unauthorized to use the API
		return nil, api.NewErrStatusCode(rsp)
	default:
		return nil, api.NewErrUnknownStatusCode(rsp)
	}
}
