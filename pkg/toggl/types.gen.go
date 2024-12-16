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

type GetTimeEntriesParams struct {
	// Get entries modified since this date using UNIX timestamp, including deleted ones.
	Since int
	// Get entries with start time before the given date.
	Before time.Time
	// Get entries with start time from the given time. To be used with end_date.
	StartDate time.Time
	// Get entries with start time until the given time. To be used with start_date.
	EndDate time.Time
	// Should the response contain meta data such as `client_name`, `project_name`, `project_color`, `project_active`, `project_billable`, `user_name`, and `user_avatar_url`.
	Meta bool
	// Include sharing details in the response
	IncludeSharing bool
}

// UserWithRelated defines a model
type UserWithRelated struct {
	ID                 int         `json:"id,omitzero"`
	APIToken           string      `json:"api_token,omitzero"`
	Email              types.Email `json:"email,omitzero"`
	Fullname           string      `json:"fullname,omitzero"`
	Timezone           string      `json:"timezone,omitzero"`
	TogglAccountsID    string      `json:"toggl_accounts_id,omitzero"`
	DefaultWorkspaceID int         `json:"default_workspace_id,omitzero"`
	BeginningOfWeek    int         `json:"beginning_of_week,omitzero"`
	ImageURL           url.URL     `json:"image_url,omitzero"`
	CreatedAt          time.Time   `json:"created_at,omitzero"`
	UpdatedAt          time.Time   `json:"updated_at,omitzero"`
	OpenidEmail        struct{}    `json:"openid_email"`
	OpenidEnabled      bool        `json:"openid_enabled,omitzero"`
	CountryID          int         `json:"country_id,omitzero"`
	HasPassword        bool        `json:"has_password,omitzero"`
	At                 time.Time   `json:"at,omitzero"`
	IntercomHash       string      `json:"intercom_hash,omitzero"`
	OauthProviders     []string    `json:"oauth_providers,omitempty"`
	// A timestamp when the authorization user session object was last updated.
	AuthorizationUpdatedAt time.Time `json:"authorization_updated_at,omitzero"`
	Tags                   Tags      `json:"tags,omitempty"`
	// Clients, null if with_related_data was not set to true or if the user does not have any clients
	Clients     Clients     `json:"clients,omitempty"`
	TimeEntries TimeEntries `json:"time_entries,omitempty"`
	Options     *Options    `json:"options,omitempty"`
	// Projects, null if with_related_data was not set to true or if the user does not have any projects
	Projects   Projects   `json:"projects,omitempty"`
	Workspaces Workspaces `json:"workspaces,omitempty"`
}

// Tags defines a model
type Tags []Tag

// Tag defines a model
type Tag struct {
	ID          int       `json:"id,omitzero"`
	WorkspaceID int       `json:"workspace_id,omitzero"`
	Name        string    `json:"name,omitzero"`
	At          time.Time `json:"at,omitzero"`
	CreatorID   int       `json:"creator_id,omitzero"`
}

// Clients defines a model
type Clients []WorkClient

// WorkClient defines a model
type WorkClient struct {
	// Client ID
	ID int `json:"id,omitzero"`
	// Workspace ID
	Wid int `json:"wid,omitzero"`
	// true, if the client is archived
	Archived bool `json:"archived,omitzero"`
	// Name of the client
	Name string `json:"name,omitzero"`
	// When was the last update
	At time.Time `json:"at,omitzero"`
	// The ID of the user who created the client
	CreatorID int `json:"creator_id,omitzero"`
	// The external ID of the linked entity in the external system (e.g. JIRA/SalesForce)
	IntegrationExtID *string `json:"integration_ext_id,omitempty"`
	// The external type of the linked entity in the external system (e.g. JIRA/SalesForce)
	IntegrationExtType *string `json:"integration_ext_type,omitempty"`
	// The provider (e.g. JIRA/SalesForce) that has an entity linked to this Toggl Track entity
	IntegrationProvider *string `json:"integration_provider,omitempty"`
	Notes               *string `json:"notes,omitempty"`
	// List of authorization permissions for this client.
	Permissions *string `json:"permissions,omitempty"`
}

// TimeEntries defines a model
type TimeEntries []TimeEntry

// TimeEntry defines a model
type TimeEntry struct {
	ID              int           `json:"id,omitzero"`
	WorkspaceID     int           `json:"workspace_id,omitzero"`
	ProjectID       int           `json:"project_id,omitzero"`
	TaskID          struct{}      `json:"task_id"`
	Billable        bool          `json:"billable,omitzero"`
	Start           time.Time     `json:"start,omitzero"`
	Stop            time.Time     `json:"stop,omitzero"`
	Duration        time.Duration `json:"duration,omitzero"`
	Description     string        `json:"description,omitzero"`
	Tags            []string      `json:"tags,omitempty"`
	TagIds          []string      `json:"tag_ids,omitempty"`
	Duronly         bool          `json:"duronly,omitzero"`
	At              time.Time     `json:"at,omitzero"`
	ServerDeletedAt time.Time     `json:"server_deleted_at,omitzero"`
	UserID          int           `json:"user_id,omitzero"`
	UID             int           `json:"uid,omitzero"`
	Wid             int           `json:"wid,omitzero"`
	Pid             int           `json:"pid,omitzero"`
	ClientName      *string       `json:"client_name,omitempty"`
	ProjectName     *string       `json:"project_name,omitempty"`
	ProjectColor    *string       `json:"project_color,omitempty"`
	ProjectActive   *bool         `json:"project_active,omitempty"`
	ProjectBillable *bool         `json:"project_billable,omitempty"`
	UserName        *string       `json:"user_name,omitempty"`
	UserAvatarURL   *url.URL      `json:"user_avatar_url,omitempty"`
}

// Options defines a model
type Options map[string]string

// Projects defines a model
type Projects []Project

// Project defines a model
type Project struct {
	// Project ID
	ID          int `json:"id,omitzero"`
	WorkspaceID int `json:"workspace_id,omitzero"`
	// Client ID
	ClientID  int    `json:"client_id,omitzero"`
	Name      string `json:"name,omitzero"`
	IsPrivate bool   `json:"is_private,omitzero"`
	// Whether the project is active or archived
	Active bool `json:"active,omitzero"`
	// Last updated date
	At time.Time `json:"at,omitzero"`
	// Creation date
	CreatedAt       time.Time `json:"created_at,omitzero"`
	ServerDeletedAt struct{}  `json:"server_deleted_at"`
	// Color
	Color string `json:"color,omitzero"`
	// Whether the project is billable, premium feature
	Billable bool     `json:"billable,omitzero"`
	Template struct{} `json:"template"`
	// Whether estimates are based on task hours, premium feature
	AutoEstimates bool             `json:"auto_estimates,omitzero"`
	CurrentPeriod *RecurringPeriod `json:"current_period,omitempty"`
	// End date
	EndDate *string `json:"end_date,omitempty"`
	// Estimated hours
	EstimatedHours int `json:"estimated_hours,omitzero"`
	// Estimated seconds
	EstimatedSeconds int      `json:"estimated_seconds,omitzero"`
	Rate             struct{} `json:"rate"`
	RateLastUpdated  struct{} `json:"rate_last_updated"`
	// Currency, premium feature
	Currency            string   `json:"currency,omitzero"`
	Recurring           bool     `json:"recurring,omitzero"`
	TemplateID          struct{} `json:"template_id"`
	RecurringParameters struct{} `json:"recurring_parameters"`
	// Fixed fee, premium feature
	FixedFee float32 `json:"fixed_fee"`
	// Actual hours
	ActualHours int `json:"actual_hours,omitzero"`
	// Actual seconds
	ActualSeconds int      `json:"actual_seconds,omitzero"`
	TasksCount    struct{} `json:"tasks_count"`
	CanTrackTime  bool     `json:"can_track_time,omitzero"`
	StartDate     string   `json:"start_date,omitzero"`
	Status        string   `json:"status,omitzero"`
	Wid           int      `json:"wid,omitzero"`
	// Client ID legacy field
	Cid      int  `json:"cid,omitzero"`
	IsShared bool `json:"is_shared,omitzero"`
	Pinned   bool `json:"pinned,omitzero"`
	// The external ID of the linked entity in the external system (e.g. JIRA/SalesForce)
	IntegrationExtID *string `json:"integration_ext_id,omitempty"`
	// The external type of the linked entity in the external system (e.g. JIRA/SalesForce)
	IntegrationExtType *string `json:"integration_ext_type,omitempty"`
	// The provider (e.g. JIRA/SalesForce) that has an entity linked to this Toggl Track entity
	IntegrationProvider *string `json:"integration_provider,omitempty"`
}

// Workspaces defines a model
type Workspaces []Workspace

// Workspace defines a model
type Workspace struct {
	ID                          int       `json:"id,omitzero"`
	OrganizationID              int       `json:"organization_id,omitzero"`
	Name                        string    `json:"name,omitzero"`
	Premium                     bool      `json:"premium,omitzero"`
	BusinessWs                  bool      `json:"business_ws,omitzero"`
	Admin                       bool      `json:"admin,omitzero"`
	Role                        string    `json:"role,omitzero"`
	SuspendedAt                 struct{}  `json:"suspended_at"`
	ServerDeletedAt             struct{}  `json:"server_deleted_at"`
	DefaultHourlyRate           struct{}  `json:"default_hourly_rate"`
	RateLastUpdated             struct{}  `json:"rate_last_updated"`
	DefaultCurrency             string    `json:"default_currency,omitzero"`
	OnlyAdminsMayCreateProjects bool      `json:"only_admins_may_create_projects,omitzero"`
	OnlyAdminsMayCreateTags     bool      `json:"only_admins_may_create_tags,omitzero"`
	OnlyAdminsSeeBillableRates  bool      `json:"only_admins_see_billable_rates,omitzero"`
	OnlyAdminsSeeTeamDashboard  bool      `json:"only_admins_see_team_dashboard,omitzero"`
	ProjectsBillableByDefault   bool      `json:"projects_billable_by_default,omitzero"`
	ProjectsPrivateByDefault    bool      `json:"projects_private_by_default,omitzero"`
	ProjectsEnforceBillable     bool      `json:"projects_enforce_billable,omitzero"`
	LimitPublicProjectData      bool      `json:"limit_public_project_data,omitzero"`
	LastModified                time.Time `json:"last_modified,omitzero"`
	ReportsCollapse             bool      `json:"reports_collapse,omitzero"`
	Rounding                    int       `json:"rounding,omitzero"`
	RoundingMinutes             int       `json:"rounding_minutes,omitzero"`
	APIToken                    uuid.UUID `json:"api_token,omitzero"`
	At                          time.Time `json:"at,omitzero"`
	LogoURL                     url.URL   `json:"logo_url,omitzero"`
	IcalURL                     string    `json:"ical_url,omitzero"`
	IcalEnabled                 bool      `json:"ical_enabled,omitzero"`
	CsvUpload                   struct{}  `json:"csv_upload"`
	Subscription                struct{}  `json:"subscription"`
	HideStartEndTimes           bool      `json:"hide_start_end_times,omitzero"`
	WorkingHoursInMinutes       struct{}  `json:"working_hours_in_minutes"`
}

// RecurringPeriod defines a model
type RecurringPeriod struct{}

// APIErrorString defines a model
type APIErrorString string

// NewTimeEntry defines a model
type NewTimeEntry struct {
	// Whether the time entry is marked as billable, optional, default false
	Billable bool `json:"billable,omitzero"`
	// Must be provided when creating a time entry and should identify the service/application used to create it
	CreatedWith string `json:"created_with,omitzero"`
	// Time entry description, optional
	Description string `json:"description,omitzero"`
	// Time entry duration. For running entries should be negative, preferable -1
	Duration time.Duration `json:"duration,omitzero"`
	// Deprecated: Used to create a time entry with a duration but without a stop time. This parameter can be ignored.
	Duronly       *bool          `json:"duronly,omitempty"`
	EventMetadata *EventMetadata `json:"event_metadata,omitempty"`
	// Project ID, legacy field
	Pid *int `json:"pid,omitempty"`
	// Project ID, optional
	ProjectID *int `json:"project_id,omitempty"`
	// List of user IDs to share this time entry with
	SharedWithUserIds []int `json:"shared_with_user_ids,omitempty"`
	// Start time, required for creation.
	Start time.Time `json:"start,omitzero"`
	// If provided during creation, the date part will take precedence over the date part of "start". Format: 2006-11-07
	StartDate *string `json:"start_date,omitempty"`
	/*
	   Stop time in UTC, can be omitted if it's still running or created with "duration".
	   If "stop" and "duration" are provided, values must be consistent (start + duration == stop)
	*/
	Stop *string `json:"stop,omitempty"`
	// Used when updating an existing time entry
	TagAction *TagAction `json:"tag_action,omitempty"`
	// IDs of tags to add/remove
	TagIds []int `json:"tag_ids,omitempty"`
	// Names of tags to add/remove. If name does not exist as tag, one will be created automatically
	Tags []string `json:"tags,omitempty"`
	// Task ID, optional
	TaskID *int `json:"task_id,omitempty"`
	// Task ID, legacy field
	Tid *int `json:"tid,omitempty"`
	// Time Entry creator ID, legacy field
	UID *int `json:"uid,omitempty"`
	// Time Entry creator ID, if omitted will use the requester user ID
	UserID *int `json:"user_id,omitempty"`
	// Workspace ID, legacy field
	Wid *int `json:"wid,omitempty"`
	// Workspace ID
	WorkspaceID int `json:"workspace_id,omitzero"`
}

// EventMetadata defines a model
type EventMetadata struct{}

// Used when updating an existing time entry
type TagAction string

const (
	TagActionAdd    TagAction = "add"
	TagActionDelete TagAction = "delete"
)

// Organization defines a model
type Organization struct {
	ID            int     `json:"id,omitzero"`
	Name          string  `json:"name,omitzero"`
	Permissions   *string `json:"permissions,omitempty"`
	WorkspaceID   int     `json:"workspace_id,omitzero"`
	WorkspaceName string  `json:"workspace_name,omitzero"`
}

// NewOrganization defines a model
type NewOrganization struct {
	// Name of the organization
	Name string `json:"name,omitzero"`
	// Name of the workspace
	WorkspaceName string `json:"workspace_name,omitzero"`
}

// ListMeOrganizationsOkJSONResponse defines a model
type ListMeOrganizationsOkJSONResponse []ListMeOrganizationsOkJSONResponseItems

// ListMeOrganizationsOkJSONResponseItems defines a model
type ListMeOrganizationsOkJSONResponseItems struct {
	ID                      int                                             `json:"id,omitzero"`
	Name                    string                                          `json:"name,omitzero"`
	PricingPlanID           int                                             `json:"pricing_plan_id,omitzero"`
	CreatedAt               time.Time                                       `json:"created_at,omitzero"`
	At                      time.Time                                       `json:"at,omitzero"`
	ServerDeletedAt         struct{}                                        `json:"server_deleted_at"`
	IsMultiWorkspaceEnabled bool                                            `json:"is_multi_workspace_enabled,omitzero"`
	SuspendedAt             struct{}                                        `json:"suspended_at"`
	UserCount               int                                             `json:"user_count,omitzero"`
	TrialInfo               ListMeOrganizationsOkJSONResponseItemsTrialInfo `json:"trial_info"`
	IsUnified               bool                                            `json:"is_unified,omitzero"`
	MaxWorkspaces           int                                             `json:"max_workspaces,omitzero"`
	Permissions             struct{}                                        `json:"permissions"`
	Admin                   bool                                            `json:"admin,omitzero"`
	Owner                   bool                                            `json:"owner,omitzero"`
	PricingPlanName         string                                          `json:"pricing_plan_name,omitzero"`
	PricingPlanEnterprise   bool                                            `json:"pricing_plan_enterprise,omitzero"`
}

// ListMeOrganizationsOkJSONResponseItemsTrialInfo defines a model
type ListMeOrganizationsOkJSONResponseItemsTrialInfo struct {
	Trial             bool     `json:"trial,omitzero"`
	TrialAvailable    bool     `json:"trial_available,omitzero"`
	TrialEndDate      struct{} `json:"trial_end_date"`
	NextPaymentDate   struct{} `json:"next_payment_date"`
	LastPricingPlanID struct{} `json:"last_pricing_plan_id"`
	CanHaveTrial      bool     `json:"can_have_trial,omitzero"`
	TrialPlanID       struct{} `json:"trial_plan_id"`
}
