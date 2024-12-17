package main

import (
	"context"
	"net/http"
	"os"
	"strconv"
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
	_ = c

	const testOrgID = 9011051
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		serverURL+"/organizations/"+strconv.Itoa(testOrgID),
		nil,
	)
	if err != nil {
		return err
	}

	req.SetBasicAuth(tkn, "api_token")

	if _, err := http.DefaultClient.Do(req); err != nil {
		return err
	}

	return nil
}

func maskSecrets(i *cassette.Interaction) error {
	tkn := os.Getenv("TOGGL_TOKEN")
	i.Response.Body = strings.ReplaceAll(i.Response.Body,
		tkn, strings.Repeat("*", len(tkn)))

	return nil
}

func runAll(ctx context.Context, c *toggl.Client) error {
	// Return Current User
	me, err := c.GetMe(ctx, &toggl.GetMeParams{WithRelatedData: true})
	if err != nil {
		return err
	}

	// Create a time entry
	if _, err := c.CreateTimeEntry(ctx, me.DefaultWorkspaceID, toggl.NewTimeEntry{
		WorkspaceID: me.DefaultWorkspaceID,
		Start:       time.Now().Add(-10 * time.Minute),
		CreatedWith: "github.com/go-api-libs/toggl",
		Description: "Hello Toggl",
		Duration:    toggl.DurationRunning,
	}); err != nil {
		return err
	}

	// Get current time entry
	cur, err := c.GetCurrentTimeEntry(ctx)
	if err != nil {
		return err
	}

	// Stop an existing time entry
	if _, err := c.StopTimeEntry(ctx, me.DefaultWorkspaceID, cur.ID); err != nil {
		return err
	}

	// List time entries
	now := time.Now()
	if _, err := c.ListTimeEntries(ctx, &toggl.ListTimeEntriesParams{
		StartDate: now,
		EndDate:   now,
	}); err != nil {
		return err
	}

	// Create a new organization
	if _, err := c.CreateOrganization(ctx, toggl.NewOrganization{
		Name:          "My Example Organization",
		WorkspaceName: "My Default Workspace",
	}); err != nil {
		return err
	}

	// List my organizations
	if _, err := c.ListOrganizations(ctx); err != nil {
		return err
	}

	return nil
}
