// This file provides a client with methods as well as functions to interact with the HTTP API.
//
// Code generated by github.com/MarkRosemaker DO NOT EDIT.

package toggl

import (
	"context"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

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
		json.WithMarshalers(json.NewMarshalers(
			json.MarshalFuncV2(jsonutil.URLMarshal),
			json.MarshalFuncV2(jsonutil.DurationMarshalIntSeconds))),
		json.WithUnmarshalers(json.NewUnmarshalers(
			json.UnmarshalFuncV2(jsonutil.URLUnmarshal),
			json.UnmarshalFuncV2(jsonutil.DurationUnmarshalIntSeconds))))
)

// Client conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The HTTP client to use for requests.
	cli *http.Client
	// The authorization header to use.
	authHeader string
}

// NewClient creates a new Client with the given username and password to be used for basic authentication.
func NewClient(username, password string) (*Client, error) {
	if username == "" {
		return nil, errors.New("username is empty")
	}

	if password == "" {
		return nil, errors.New("password is empty")
	}

	return &Client{
		authHeader: "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password)),
		cli:        http.DefaultClient,
	}, nil
}

// Returns details for the current user.
//
//	GET /me
func (c *Client) GetMe(ctx context.Context, params *GetMeParams) (*UserWithRelated, error) {
	return GetMe[UserWithRelated](ctx, c, params)
}

// Returns details for the current user.
// You can define a custom result to unmarshal the response into.
//
//	GET /me
func GetMe[R any](ctx context.Context, c *Client, params *GetMeParams) (*R, error) {
	u := baseURL.JoinPath("/me")

	if params != nil && params.WithRelatedData {
		u.RawQuery = url.Values{"with_related_data": []string{"true"}}.Encode()
	}

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
		// Returns the user.
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

// Creates a new workspace TimeEntry.
//
//	POST /workspaces/{workspace_id}/time_entries
func (c *Client) CreateTimeEntry(ctx context.Context, workspaceID int, reqBody NewTimeEntry) (*TimeEntry, error) {
	return CreateTimeEntry[TimeEntry](ctx, c, workspaceID, reqBody)
}

// Creates a new workspace TimeEntry.
// You can define a custom request body to marshal and a custom result to unmarshal the response into.
//
//	POST /workspaces/{workspace_id}/time_entries
func CreateTimeEntry[R any, B any](ctx context.Context, c *Client, workspaceID int, reqBody B) (*R, error) {
	u := baseURL.JoinPath("workspaces", strconv.Itoa(workspaceID), "time_entries")
	req := (&http.Request{
		Header: http.Header{
			"Authorization": []string{c.authHeader},
			"Content-Type":  []string{"application/json"},
			"User-Agent":    []string{userAgent},
		},
		Host:       u.Host,
		Method:     http.MethodPost,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		URL:        u,
	}).WithContext(ctx)

	var pw *io.PipeWriter
	req.Body, pw = io.Pipe()
	defer req.Body.Close()
	go func() {
		pw.CloseWithError(json.MarshalWrite(pw, reqBody, jsonOpts))
	}()

	rsp, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	switch rsp.StatusCode {
	case http.StatusOK:
		// Returns a time entry
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
	case http.StatusBadRequest:
		// Returned when the user made a bad request
		switch mt, _, _ := strings.Cut(rsp.Header.Get("Content-Type"), ";"); mt {
		case "application/json":
			var errOut APIErrorString
			if err := json.UnmarshalRead(rsp.Body, &errOut, jsonOpts); err != nil {
				return nil, api.WrapDecodingError(rsp, err)
			}

			return nil, api.NewErrCustom(rsp, &errOut)
		default:
			return nil, api.NewErrUnknownContentType(rsp)
		}
	default:
		return nil, api.NewErrUnknownStatusCode(rsp)
	}
}

// Load running time entry for the current user.
//
//	GET /me/time_entries/current
func (c *Client) GetCurrentTimeEntry(ctx context.Context) (*TimeEntry, error) {
	return GetCurrentTimeEntry[TimeEntry](ctx, c)
}

// Load running time entry for the current user.
// You can define a custom result to unmarshal the response into.
//
//	GET /me/time_entries/current
func GetCurrentTimeEntry[R any](ctx context.Context, c *Client) (*R, error) {
	u := baseURL.JoinPath("/me/time_entries/current")
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
		// Returns a time entry
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
	default:
		return nil, api.NewErrUnknownStatusCode(rsp)
	}
}

// Stops a workspace time entry.
//
//	PATCH /workspaces/{workspace_id}/time_entries/{time_entry_id}/stop
func (c *Client) StopTimeEntry(ctx context.Context, workspaceID int, timeEntryID int) (*TimeEntry, error) {
	return StopTimeEntry[TimeEntry](ctx, c, workspaceID, timeEntryID)
}

// Stops a workspace time entry.
// You can define a custom result to unmarshal the response into.
//
//	PATCH /workspaces/{workspace_id}/time_entries/{time_entry_id}/stop
func StopTimeEntry[R any](ctx context.Context, c *Client, workspaceID int, timeEntryID int) (*R, error) {
	u := baseURL.JoinPath("workspaces", strconv.Itoa(workspaceID), "time_entries", strconv.Itoa(timeEntryID), "stop")
	req := (&http.Request{
		Header: http.Header{
			"Authorization": []string{c.authHeader},
			"User-Agent":    []string{userAgent},
		},
		Host:       u.Host,
		Method:     http.MethodPatch,
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
		// Returns a time entry
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
	default:
		return nil, api.NewErrUnknownStatusCode(rsp)
	}
}

// Lists latest time entries.
//
//	GET /me/time_entries
func (c *Client) GetTimeEntries(ctx context.Context, params *GetTimeEntriesParams) (TimeEntries, error) {
	return GetTimeEntries[TimeEntries](ctx, c, params)
}

// Lists latest time entries.
// You can define a custom result to unmarshal the response into.
//
//	GET /me/time_entries
func GetTimeEntries[R any](ctx context.Context, c *Client, params *GetTimeEntriesParams) (R, error) {
	u := baseURL.JoinPath("/me/time_entries")

	if params != nil {
		q := make(url.Values, 6)

		if params.Since != 0 {
			q["since"] = []string{strconv.Itoa(params.Since)}
		}

		if !params.Before.IsZero() {
			q["before"] = []string{params.Before.Format(time.RFC3339)}
		}

		if !params.StartDate.IsZero() {
			q["start_date"] = []string{params.StartDate.Format(time.RFC3339)}
		}

		if !params.EndDate.IsZero() {
			q["end_date"] = []string{params.EndDate.Format(time.RFC3339)}
		}

		if params.Meta {
			q["meta"] = []string{"true"}
		}

		if params.IncludeSharing {
			q["include_sharing"] = []string{"true"}
		}

		u.RawQuery = q.Encode()
	}

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

	var out R
	rsp, err := c.cli.Do(req)
	if err != nil {
		return out, err
	}
	defer rsp.Body.Close()

	switch rsp.StatusCode {
	case http.StatusOK:
		// Returns a list of time entries
		switch mt, _, _ := strings.Cut(rsp.Header.Get("Content-Type"), ";"); mt {
		case "application/json":
			if err := json.UnmarshalRead(rsp.Body, &out, jsonOpts); err != nil {
				return out, api.WrapDecodingError(rsp, err)
			}

			return out, nil
		default:
			return out, api.NewErrUnknownContentType(rsp)
		}
	case http.StatusBadRequest:
		// Returned when the user made a bad request
		switch mt, _, _ := strings.Cut(rsp.Header.Get("Content-Type"), ";"); mt {
		case "application/json":
			var errOut APIErrorString
			if err := json.UnmarshalRead(rsp.Body, &errOut, jsonOpts); err != nil {
				return out, api.WrapDecodingError(rsp, err)
			}

			return out, api.NewErrCustom(rsp, &errOut)
		default:
			return out, api.NewErrUnknownContentType(rsp)
		}
	default:
		return out, api.NewErrUnknownStatusCode(rsp)
	}
}

// Creates a new organization with a single workspace and assigns current user as the organization owner
//
//	POST /organizations
func (c *Client) CreateOrganization(ctx context.Context, reqBody NewOrganization) (*PostOrganizationsOkJSONResponse, error) {
	return CreateOrganization[PostOrganizationsOkJSONResponse](ctx, c, reqBody)
}

// Creates a new organization with a single workspace and assigns current user as the organization owner
// You can define a custom request body to marshal and a custom result to unmarshal the response into.
//
//	POST /organizations
func CreateOrganization[R any, B any](ctx context.Context, c *Client, reqBody B) (*R, error) {
	u := baseURL.JoinPath("/organizations")
	req := (&http.Request{
		Header: http.Header{
			"Authorization": []string{c.authHeader},
			"Content-Type":  []string{"application/json"},
			"User-Agent":    []string{userAgent},
		},
		Host:       u.Host,
		Method:     http.MethodPost,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		URL:        u,
	}).WithContext(ctx)

	var pw *io.PipeWriter
	req.Body, pw = io.Pipe()
	defer req.Body.Close()
	go func() {
		pw.CloseWithError(json.MarshalWrite(pw, reqBody, jsonOpts))
	}()

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
	default:
		return nil, api.NewErrUnknownStatusCode(rsp)
	}
}
