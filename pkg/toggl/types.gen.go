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
	CountryID          struct{}    `json:"country_id"`
	HasPassword        bool        `json:"has_password"`
	At                 time.Time   `json:"at"`
	IntercomHash       string      `json:"intercom_hash"`
	OauthProviders     []string    `json:"oauth_providers"`
	// AuthorizationUpdatedAt timestamp when the authorization user session object was last updated.
	AuthorizationUpdatedAt time.Time   `json:"authorization_updated_at"`
	Tags                   Tags        `json:"tags"`
	Clients                Clients     `json:"clients"`
	TimeEntries            TimeEntries `json:"time_entries"`
	Projects               Projects    `json:"projects"`
	Workspaces             Workspaces  `json:"workspaces"`
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
	ID        int       `json:"id"`
	Wid       int       `json:"wid"`
	Archived  bool      `json:"archived"`
	Name      string    `json:"name"`
	At        time.Time `json:"at"`
	CreatorID int       `json:"creator_id"`
}

// TimeEntries defines a model
type TimeEntries []TimeEntry

// TimeEntry defines a model
type TimeEntry struct {
	ID              int        `json:"id"`
	WorkspaceID     int        `json:"workspace_id"`
	ProjectID       int        `json:"project_id"`
	TaskID          struct{}   `json:"task_id"`
	Billable        bool       `json:"billable"`
	Start           time.Time  `json:"start"`
	Stop            time.Time  `json:"stop"`
	Duration        int        `json:"duration"`
	Description     string     `json:"description"`
	Tags            []struct{} `json:"tags"`
	TagIds          []struct{} `json:"tag_ids"`
	Duronly         bool       `json:"duronly"`
	At              time.Time  `json:"at"`
	ServerDeletedAt struct{}   `json:"server_deleted_at"`
	UserID          int        `json:"user_id"`
	UID             int        `json:"uid"`
	Wid             int        `json:"wid"`
	Pid             int        `json:"pid"`
}

// Projects defines a model
type Projects []Project

// Project defines a model
type Project struct {
	ID                  int       `json:"id"`
	WorkspaceID         int       `json:"workspace_id"`
	ClientID            int       `json:"client_id"`
	Name                string    `json:"name"`
	IsPrivate           bool      `json:"is_private"`
	Active              bool      `json:"active"`
	At                  time.Time `json:"at"`
	CreatedAt           time.Time `json:"created_at"`
	ServerDeletedAt     struct{}  `json:"server_deleted_at"`
	Color               string    `json:"color"`
	Billable            bool      `json:"billable"`
	Template            struct{}  `json:"template"`
	AutoEstimates       struct{}  `json:"auto_estimates"`
	EstimatedHours      struct{}  `json:"estimated_hours"`
	EstimatedSeconds    struct{}  `json:"estimated_seconds"`
	Rate                struct{}  `json:"rate"`
	RateLastUpdated     struct{}  `json:"rate_last_updated"`
	Currency            struct{}  `json:"currency"`
	Recurring           bool      `json:"recurring"`
	TemplateID          struct{}  `json:"template_id"`
	RecurringParameters struct{}  `json:"recurring_parameters"`
	FixedFee            struct{}  `json:"fixed_fee"`
	ActualHours         int       `json:"actual_hours"`
	ActualSeconds       int       `json:"actual_seconds"`
	TasksCount          struct{}  `json:"tasks_count"`
	CanTrackTime        bool      `json:"can_track_time"`
	StartDate           string    `json:"start_date"`
	Status              string    `json:"status"`
	Wid                 int       `json:"wid"`
	Cid                 int       `json:"cid"`
	IsShared            bool      `json:"is_shared"`
	Pinned              bool      `json:"pinned"`
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
