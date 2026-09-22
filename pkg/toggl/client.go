package toggl

import (
	"context"
	"os"
	"time"
)

// When creating a new running time entry, use DurationRunning as the duration.
const DurationRunning = -time.Second

const (
	keyToken    = "TOGGL_TOKEN"
	keyUserName = "TOGGL_API_USERNAME"
	keyPassword = "TOGGL_API_PASSWORD"
)

func init() {
	// You can find your API token at the bottom of the [profile page].
	//
	// [profile page]: https://track.toggl.com/profile
	token := os.Getenv("TOGGL_TOKEN")
	if token == "" || os.Getenv(keyUserName) != "" || os.Getenv(keyPassword) != "" {
		return
	}

	if err := os.Setenv(keyUserName, token); err != nil {
		panic(err)
	}

	if err := os.Setenv(keyPassword, "api_token"); err != nil {
		panic(err)
	}
}

func (c *Client) ListTimeEntriesInRange(ctx context.Context, start, end time.Time) (TimeEntries, error) {
	return c.ListTimeEntries(ctx, &ListTimeEntriesParams{
		StartDate:      start,
		EndDate:        end,
		Meta:           true,
		IncludeSharing: true,
	})
}
