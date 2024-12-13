// This file provides types for the API.
//
// Code generated by github.com/MarkRosemaker DO NOT EDIT.

package toggl

import (
	"net/url"
	"time"

	"github.com/go-api-libs/types"
	"github.com/google/uuid"
)

type GetMeParams struct {
	// Retrieve user related data (clients, projects, tasks, tags, workspaces, time entries, etc.)
	WithRelatedData bool
}

// UserWithRelated defines a model
type UserWithRelated struct {
	ID                 int         `json:"id"`
	APIToken           string      `json:"api_token"`
	Email              types.Email `json:"email"`
	Fullname           string      `json:"fullname"`
	Timezone           string      `json:"timezone"`
	TogglAccountsID    string      `json:"toggl_accounts_id"`
	DefaultWorkspaceID int         `json:"default_workspace_id"`
	BeginningOfWeek    int         `json:"beginning_of_week"`
	ImageURL           url.URL     `json:"image_url"`
	CreatedAt          time.Time   `json:"created_at"`
	UpdatedAt          time.Time   `json:"updated_at"`
	OpenidEmail        struct{}    `json:"openid_email"`
	OpenidEnabled      bool        `json:"openid_enabled"`
	CountryID          int         `json:"country_id"`
	HasPassword        bool        `json:"has_password"`
	At                 time.Time   `json:"at"`
	IntercomHash       string      `json:"intercom_hash"`
	OauthProviders     []string    `json:"oauth_providers"`
	// A timestamp when the authorization user session object was last updated.
	AuthorizationUpdatedAt time.Time `json:"authorization_updated_at"`
	Tags                   Tags      `json:"tags"`
	// Clients, null if with_related_data was not set to true or if the user does not have any clients
	Clients     Clients     `json:"clients"`
	TimeEntries TimeEntries `json:"time_entries"`
	Options     *Options    `json:"options"`
	// Projects, null if with_related_data was not set to true or if the user does not have any projects
	Projects   Projects   `json:"projects"`
	Workspaces Workspaces `json:"workspaces"`
}

// Tags defines a model
type Tags []Tag

// Tag defines a model
type Tag struct {
	ID          int       `json:"id"`
	WorkspaceID int       `json:"workspace_id"`
	Name        string    `json:"name"`
	At          time.Time `json:"at"`
	CreatorID   int       `json:"creator_id"`
}

// Clients defines a model
type Clients []WorkClient

// WorkClient defines a model
type WorkClient struct {
	// Client ID
	ID int `json:"id"`
	// Workspace ID
	Wid int `json:"wid"`
	// true, if the client is archived
	Archived bool `json:"archived"`
	// Name of the client
	Name string `json:"name"`
	// When was the last update
	At time.Time `json:"at"`
	// The ID of the user who created the client
	CreatorID int `json:"creator_id"`
	// The external ID of the linked entity in the external system (e.g. JIRA/SalesForce)
	IntegrationExtID *string `json:"integration_ext_id"`
	// The external type of the linked entity in the external system (e.g. JIRA/SalesForce)
	IntegrationExtType *string `json:"integration_ext_type"`
	// The provider (e.g. JIRA/SalesForce) that has an entity linked to this Toggl Track entity
	IntegrationProvider *string `json:"integration_provider"`
	Notes               *string `json:"notes"`
	// List of authorization permissions for this client.
	Permissions *string `json:"permissions"`
}

// TimeEntries defines a model
type TimeEntries []TimeEntry

// TimeEntry defines a model
type TimeEntry struct {
	ID              int       `json:"id"`
	WorkspaceID     int       `json:"workspace_id"`
	ProjectID       int       `json:"project_id"`
	TaskID          struct{}  `json:"task_id"`
	Billable        bool      `json:"billable"`
	Start           time.Time `json:"start"`
	Stop            time.Time `json:"stop"`
	Duration        int       `json:"duration"`
	Description     string    `json:"description"`
	Tags            []string  `json:"tags"`
	TagIds          []string  `json:"tag_ids"`
	Duronly         bool      `json:"duronly"`
	At              time.Time `json:"at"`
	ServerDeletedAt struct{}  `json:"server_deleted_at"`
	UserID          int       `json:"user_id"`
	UID             int       `json:"uid"`
	Wid             int       `json:"wid"`
	Pid             int       `json:"pid"`
}

// Options defines a model
type Options map[string]string

// Projects defines a model
type Projects []Project

// Project defines a model
type Project struct {
	// Project ID
	ID          int `json:"id"`
	WorkspaceID int `json:"workspace_id"`
	// Client ID
	ClientID  int    `json:"client_id"`
	Name      string `json:"name"`
	IsPrivate bool   `json:"is_private"`
	// Whether the project is active or archived
	Active bool `json:"active"`
	// Last updated date
	At time.Time `json:"at"`
	// Creation date
	CreatedAt       time.Time `json:"created_at"`
	ServerDeletedAt struct{}  `json:"server_deleted_at"`
	// Color
	Color string `json:"color"`
	// Whether the project is billable, premium feature
	Billable bool     `json:"billable"`
	Template struct{} `json:"template"`
	// Whether estimates are based on task hours, premium feature
	AutoEstimates bool             `json:"auto_estimates"`
	CurrentPeriod *RecurringPeriod `json:"current_period"`
	// End date
	EndDate *string `json:"end_date"`
	// Estimated hours
	EstimatedHours int `json:"estimated_hours"`
	// Estimated seconds
	EstimatedSeconds int      `json:"estimated_seconds"`
	Rate             struct{} `json:"rate"`
	RateLastUpdated  struct{} `json:"rate_last_updated"`
	// Currency, premium feature
	Currency            string   `json:"currency"`
	Recurring           bool     `json:"recurring"`
	TemplateID          struct{} `json:"template_id"`
	RecurringParameters struct{} `json:"recurring_parameters"`
	// Fixed fee, premium feature
	FixedFee float32 `json:"fixed_fee"`
	// Actual hours
	ActualHours int `json:"actual_hours"`
	// Actual seconds
	ActualSeconds int      `json:"actual_seconds"`
	TasksCount    struct{} `json:"tasks_count"`
	CanTrackTime  bool     `json:"can_track_time"`
	StartDate     string   `json:"start_date"`
	Status        string   `json:"status"`
	Wid           int      `json:"wid"`
	// Client ID legacy field
	Cid      int  `json:"cid"`
	IsShared bool `json:"is_shared"`
	Pinned   bool `json:"pinned"`
	// The external ID of the linked entity in the external system (e.g. JIRA/SalesForce)
	IntegrationExtID *string `json:"integration_ext_id"`
	// The external type of the linked entity in the external system (e.g. JIRA/SalesForce)
	IntegrationExtType *string `json:"integration_ext_type"`
	// The provider (e.g. JIRA/SalesForce) that has an entity linked to this Toggl Track entity
	IntegrationProvider *string `json:"integration_provider"`
}

// Workspaces defines a model
type Workspaces []Workspace

// Workspace defines a model
type Workspace struct {
	ID                          int       `json:"id"`
	OrganizationID              int       `json:"organization_id"`
	Name                        string    `json:"name"`
	Premium                     bool      `json:"premium"`
	BusinessWs                  bool      `json:"business_ws"`
	Admin                       bool      `json:"admin"`
	Role                        string    `json:"role"`
	SuspendedAt                 struct{}  `json:"suspended_at"`
	ServerDeletedAt             struct{}  `json:"server_deleted_at"`
	DefaultHourlyRate           struct{}  `json:"default_hourly_rate"`
	RateLastUpdated             struct{}  `json:"rate_last_updated"`
	DefaultCurrency             string    `json:"default_currency"`
	OnlyAdminsMayCreateProjects bool      `json:"only_admins_may_create_projects"`
	OnlyAdminsMayCreateTags     bool      `json:"only_admins_may_create_tags"`
	OnlyAdminsSeeBillableRates  bool      `json:"only_admins_see_billable_rates"`
	OnlyAdminsSeeTeamDashboard  bool      `json:"only_admins_see_team_dashboard"`
	ProjectsBillableByDefault   bool      `json:"projects_billable_by_default"`
	ProjectsPrivateByDefault    bool      `json:"projects_private_by_default"`
	ProjectsEnforceBillable     bool      `json:"projects_enforce_billable"`
	LimitPublicProjectData      bool      `json:"limit_public_project_data"`
	LastModified                time.Time `json:"last_modified"`
	ReportsCollapse             bool      `json:"reports_collapse"`
	Rounding                    int       `json:"rounding"`
	RoundingMinutes             int       `json:"rounding_minutes"`
	APIToken                    uuid.UUID `json:"api_token"`
	At                          time.Time `json:"at"`
	LogoURL                     url.URL   `json:"logo_url"`
	IcalURL                     string    `json:"ical_url"`
	IcalEnabled                 bool      `json:"ical_enabled"`
	CsvUpload                   struct{}  `json:"csv_upload"`
	Subscription                struct{}  `json:"subscription"`
	HideStartEndTimes           bool      `json:"hide_start_end_times"`
	WorkingHoursInMinutes       struct{}  `json:"working_hours_in_minutes"`
}

// RecurringPeriod defines a model
type RecurringPeriod struct{}

// PostWorkspaces2230580TimeEntriesBadRequestJSONResponse defines a model
type PostWorkspaces2230580TimeEntriesBadRequestJSONResponse string

// CreateTimeEntryJSONRequestBody defines a model
type CreateTimeEntryJSONRequestBody struct {
	CreatedWith string     `json:"created_with"`
	Description string     `json:"description"`
	Tags        []struct{} `json:"tags"`
	Billable    bool       `json:"billable"`
	WorkspaceID int        `json:"workspace_id"`
	Start       time.Time  `json:"start"`
}
