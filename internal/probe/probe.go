package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-api-libs/toggl/pkg/toggl"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
)

const serverURL = "https://api.track.toggl.com/api/v9"

// probe calls the API server to check what we can do
func probe() error {
	ctx := context.Background()
	tkn := os.Getenv("TOGGL_TOKEN")

	c, err := toggl.NewClientWithAPIToken(tkn)
	if err != nil {
		return err
	}

	me, err := c.GetMe(ctx, &toggl.GetMeParams{WithRelatedData: true})
	if err != nil {
		return err
	}

	if _, err := c.GetCurrentTimeEntry(ctx); err != nil {
		return err
	}

	return nil

	const timeEntryID = 3730303299

	req, err := http.NewRequest(http.MethodPatch, serverURL+fmt.Sprintf("/workspaces/%d/time_entries/%d/stop", me.DefaultWorkspaceID, timeEntryID), nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(tkn, "api_token")

	if _, err := http.DefaultClient.Do(req); err != nil {
		return err
	}

	return nil

	if _, err := c.CreateTimeEntry(ctx, me.DefaultWorkspaceID, toggl.NewTimeEntry{
		WorkspaceID: me.DefaultWorkspaceID,
		Start:       time.Now().Add(-10 * time.Minute),
		CreatedWith: "github.com/go-api-libs/toggl",
		Description: "Hello Toggl",
		Duration:    toggl.DurationRunning,
	}); err != nil {
		return err
	}

	// strings.NewReader(fmt.Sprintf(`{"workspace_id":%d}`, me.DefaultWorkspaceID))

	return nil
}

func maskSecrets(i *cassette.Interaction) error {
	tkn := os.Getenv("TOGGL_TOKEN")
	i.Response.Body = strings.ReplaceAll(i.Response.Body,
		tkn, strings.Repeat("*", len(tkn)))

	return nil
}
