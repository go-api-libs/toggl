package toggl_test

import (
	"context"
	"os"
	"time"

	"github.com/go-api-libs/toggl/pkg/toggl"
)

const serverURL = "https://api.track.toggl.com/api/v9"

func example() error {
	ctx := context.Background()
	tkn := os.Getenv("TOGGL_TOKEN")

	c, err := toggl.NewClientWithAPIToken(tkn)
	if err != nil {
		return err
	}

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
	org, err := c.CreateOrganization(ctx, toggl.NewOrganization{
		Name:          "My Example Organization",
		WorkspaceName: "My Default Workspace",
	})
	if err != nil {
		return err
	}

	// List my organizations
	if _, err := c.ListOrganizations(ctx); err != nil {
		return err
	}

	// Get organization data
	if _, err = c.GetOrganization(ctx, org.ID); err != nil {
		return err
	}

	return nil
}
