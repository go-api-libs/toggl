package toggl

import "time"

// When creating a new running time entry, use DurationRunning as the duration.
const DurationRunning = -time.Second

// NewClientWithAPIToken creates a client that authenticates requests using an API token.
// You can find your API token at the bottom of the [profile page].
//
// [profile page]: https://track.toggl.com/profile
func NewClientWithAPIToken(apiToken string) (*Client, error) {
	return NewClient(apiToken, "api_token")
}
