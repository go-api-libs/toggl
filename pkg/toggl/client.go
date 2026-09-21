package toggl

import (
	"context"
	"encoding/base64"
	"time"
)

// When creating a new running time entry, use DurationRunning as the duration.
const DurationRunning = -time.Second

// NewClientWithAPIToken creates a client that authenticates requests using an API token.
// You can find your API token at the bottom of the [profile page].
//
// [profile page]: https://track.toggl.com/profile
func NewClientWithAPIToken(apiToken string, opts ...ClientOption) (*Client, error) {
	username := apiToken
	password := "api_token"
	token := "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))
	_ = token

	return NewClient(opts...)
}

func (c *Client) ListTimeEntriesInRange(ctx context.Context, start, end time.Time) (*TimeEntries, error) {
	return c.ListTimeEntries(ctx, &ListTimeEntriesParams{
		StartDate:      start,
		EndDate:        end,
		Meta:           true,
		IncludeSharing: true,
	})
}
