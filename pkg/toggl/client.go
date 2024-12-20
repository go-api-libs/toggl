package toggl

import (
	"context"
	"time"
)

// When creating a new running time entry, use DurationRunning as the duration.
const DurationRunning = -time.Second

// NewClientWithAPIToken creates a client that authenticates requests using an API token.
// You can find your API token at the bottom of the [profile page].
//
// [profile page]: https://track.toggl.com/profile
func NewClientWithAPIToken(apiToken string) (*Client, error) {
	return NewClient(apiToken, "api_token")
}

func (c *Client) ListTimeEntriesInRange(ctx context.Context, start, end time.Time) (TimeEntries, error) {
	return c.ListTimeEntries(ctx, &ListTimeEntriesParams{
		StartDate:      start,
		EndDate:        end,
		Meta:           true,
		IncludeSharing: true,
	})
}
