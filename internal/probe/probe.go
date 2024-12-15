package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

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

	if _, err := c.CreateTimeEntry(ctx, me.DefaultWorkspaceID, toggl.NewTimeEntry{}); err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, serverURL+fmt.Sprintf("/workspaces/%d/time_entries", me.DefaultWorkspaceID), strings.NewReader(fmt.Sprintf(`{"created_with":"API example code","description":"Hello Toggl","tags":[],"billable":false,"workspace_id":%d,"start":"2016-06-08T11:02:53.000Z"}`, me.DefaultWorkspaceID)))
	if err != nil {
		return err
	}

	req.SetBasicAuth(tkn, "api_token")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("got %s", resp.Status)
	}

	return nil
}

func maskSecrets(i *cassette.Interaction) error {
	tkn := os.Getenv("TOGGL_TOKEN")
	i.Response.Body = strings.ReplaceAll(i.Response.Body,
		tkn, strings.Repeat("*", len(tkn)))

	return nil
}
