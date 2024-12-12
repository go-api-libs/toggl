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
	WithRelatedData string
}

// User defines a model
type User struct {
	ID                     int                     `json:"id"`
	APIToken               string                  `json:"api_token"`
	Email                  types.Email             `json:"email"`
	Fullname               string                  `json:"fullname"`
	Timezone               string                  `json:"timezone"`
	TogglAccountsID        string                  `json:"toggl_accounts_id"`
	DefaultWorkspaceID     int                     `json:"default_workspace_id"`
	BeginningOfWeek        int                     `json:"beginning_of_week"`
	ImageURL               url.URL                 `json:"image_url"`
	CreatedAt              time.Time               `json:"created_at"`
	UpdatedAt              time.Time               `json:"updated_at"`
	OpenidEmail            UserOpenidEmail         `json:"openid_email"`
	OpenidEnabled          bool                    `json:"openid_enabled"`
	CountryID              UserCountryID           `json:"country_id"`
	HasPassword            bool                    `json:"has_password"`
	At                     time.Time               `json:"at"`
	IntercomHash           string                  `json:"intercom_hash"`
	OauthProviders         []string                `json:"oauth_providers"`
	AuthorizationUpdatedAt time.Time               `json:"authorization_updated_at"`
	Tags                   *[]UserTagsItems        `json:"tags"`
	Clients                *[]UserClientsItems     `json:"clients"`
	TimeEntries            *[]UserTimeEntriesItems `json:"time_entries"`
	Projects               *[]UserProjectsItems    `json:"projects"`
	Workspaces             *[]UserWorkspacesItems  `json:"workspaces"`
}

// UserCountryID defines a model
type UserCountryID struct{}

// UserOpenidEmail defines a model
type UserOpenidEmail struct{}

// UserProjectsItems defines a model
type UserProjectsItems struct {
	ID                  int                                  `json:"id"`
	WorkspaceID         int                                  `json:"workspace_id"`
	ClientID            int                                  `json:"client_id"`
	Name                string                               `json:"name"`
	IsPrivate           bool                                 `json:"is_private"`
	Active              bool                                 `json:"active"`
	At                  time.Time                            `json:"at"`
	CreatedAt           time.Time                            `json:"created_at"`
	ServerDeletedAt     UserProjectsItemsServerDeletedAt     `json:"server_deleted_at"`
	Color               string                               `json:"color"`
	Billable            bool                                 `json:"billable"`
	Template            UserProjectsItemsTemplate            `json:"template"`
	AutoEstimates       UserProjectsItemsAutoEstimates       `json:"auto_estimates"`
	EstimatedHours      UserProjectsItemsEstimatedHours      `json:"estimated_hours"`
	EstimatedSeconds    UserProjectsItemsEstimatedSeconds    `json:"estimated_seconds"`
	Rate                UserProjectsItemsRate                `json:"rate"`
	RateLastUpdated     UserProjectsItemsRateLastUpdated     `json:"rate_last_updated"`
	Currency            UserProjectsItemsCurrency            `json:"currency"`
	Recurring           bool                                 `json:"recurring"`
	TemplateID          UserProjectsItemsTemplateID          `json:"template_id"`
	RecurringParameters UserProjectsItemsRecurringParameters `json:"recurring_parameters"`
	FixedFee            UserProjectsItemsFixedFee            `json:"fixed_fee"`
	ActualHours         int                                  `json:"actual_hours"`
	ActualSeconds       int                                  `json:"actual_seconds"`
	TasksCount          UserProjectsItemsTasksCount          `json:"tasks_count"`
	CanTrackTime        bool                                 `json:"can_track_time"`
	StartDate           string                               `json:"start_date"`
	Status              string                               `json:"status"`
	Wid                 int                                  `json:"wid"`
	Cid                 int                                  `json:"cid"`
	IsShared            bool                                 `json:"is_shared"`
	Pinned              bool                                 `json:"pinned"`
}

// UserProjectsItemsTemplate defines a model
type UserProjectsItemsTemplate struct{}

// UserProjectsItemsFixedFee defines a model
type UserProjectsItemsFixedFee struct{}

// UserProjectsItemsRecurringParameters defines a model
type UserProjectsItemsRecurringParameters struct{}

// UserProjectsItemsTasksCount defines a model
type UserProjectsItemsTasksCount struct{}

// UserProjectsItemsEstimatedHours defines a model
type UserProjectsItemsEstimatedHours struct{}

// UserProjectsItemsEstimatedSeconds defines a model
type UserProjectsItemsEstimatedSeconds struct{}

// UserProjectsItemsRate defines a model
type UserProjectsItemsRate struct{}

// UserProjectsItemsServerDeletedAt defines a model
type UserProjectsItemsServerDeletedAt struct{}

// UserProjectsItemsAutoEstimates defines a model
type UserProjectsItemsAutoEstimates struct{}

// UserProjectsItemsCurrency defines a model
type UserProjectsItemsCurrency struct{}

// UserProjectsItemsTemplateID defines a model
type UserProjectsItemsTemplateID struct{}

// UserProjectsItemsRateLastUpdated defines a model
type UserProjectsItemsRateLastUpdated struct{}

// UserTagsItems defines a model
type UserTagsItems struct {
	ID          int       `json:"id"`
	WorkspaceID int       `json:"workspace_id"`
	Name        string    `json:"name"`
	At          time.Time `json:"at"`
	CreatorID   int       `json:"creator_id"`
}

// UserWorkspacesItems defines a model
type UserWorkspacesItems struct {
	ID                          int                                      `json:"id"`
	OrganizationID              int                                      `json:"organization_id"`
	Name                        string                                   `json:"name"`
	Premium                     bool                                     `json:"premium"`
	BusinessWs                  bool                                     `json:"business_ws"`
	Admin                       bool                                     `json:"admin"`
	Role                        string                                   `json:"role"`
	SuspendedAt                 UserWorkspacesItemsSuspendedAt           `json:"suspended_at"`
	ServerDeletedAt             UserWorkspacesItemsServerDeletedAt       `json:"server_deleted_at"`
	DefaultHourlyRate           UserWorkspacesItemsDefaultHourlyRate     `json:"default_hourly_rate"`
	RateLastUpdated             UserWorkspacesItemsRateLastUpdated       `json:"rate_last_updated"`
	DefaultCurrency             string                                   `json:"default_currency"`
	OnlyAdminsMayCreateProjects bool                                     `json:"only_admins_may_create_projects"`
	OnlyAdminsMayCreateTags     bool                                     `json:"only_admins_may_create_tags"`
	OnlyAdminsSeeBillableRates  bool                                     `json:"only_admins_see_billable_rates"`
	OnlyAdminsSeeTeamDashboard  bool                                     `json:"only_admins_see_team_dashboard"`
	ProjectsBillableByDefault   bool                                     `json:"projects_billable_by_default"`
	ProjectsPrivateByDefault    bool                                     `json:"projects_private_by_default"`
	ProjectsEnforceBillable     bool                                     `json:"projects_enforce_billable"`
	LimitPublicProjectData      bool                                     `json:"limit_public_project_data"`
	LastModified                time.Time                                `json:"last_modified"`
	ReportsCollapse             bool                                     `json:"reports_collapse"`
	Rounding                    int                                      `json:"rounding"`
	RoundingMinutes             int                                      `json:"rounding_minutes"`
	APIToken                    uuid.UUID                                `json:"api_token"`
	At                          time.Time                                `json:"at"`
	LogoURL                     url.URL                                  `json:"logo_url"`
	IcalURL                     string                                   `json:"ical_url"`
	IcalEnabled                 bool                                     `json:"ical_enabled"`
	CsvUpload                   UserWorkspacesItemsCsvUpload             `json:"csv_upload"`
	Subscription                UserWorkspacesItemsSubscription          `json:"subscription"`
	HideStartEndTimes           bool                                     `json:"hide_start_end_times"`
	WorkingHoursInMinutes       UserWorkspacesItemsWorkingHoursInMinutes `json:"working_hours_in_minutes"`
}

// UserWorkspacesItemsServerDeletedAt defines a model
type UserWorkspacesItemsServerDeletedAt struct{}

// UserWorkspacesItemsDefaultHourlyRate defines a model
type UserWorkspacesItemsDefaultHourlyRate struct{}

// UserWorkspacesItemsRateLastUpdated defines a model
type UserWorkspacesItemsRateLastUpdated struct{}

// UserWorkspacesItemsSubscription defines a model
type UserWorkspacesItemsSubscription struct{}

// UserWorkspacesItemsSuspendedAt defines a model
type UserWorkspacesItemsSuspendedAt struct{}

// UserWorkspacesItemsWorkingHoursInMinutes defines a model
type UserWorkspacesItemsWorkingHoursInMinutes struct{}

// UserWorkspacesItemsCsvUpload defines a model
type UserWorkspacesItemsCsvUpload struct{}

// UserClientsItems defines a model
type UserClientsItems struct {
	ID        int       `json:"id"`
	Wid       int       `json:"wid"`
	Archived  bool      `json:"archived"`
	Name      string    `json:"name"`
	At        time.Time `json:"at"`
	CreatorID int       `json:"creator_id"`
}

// UserTimeEntriesItems defines a model
type UserTimeEntriesItems struct {
	ID              int                                 `json:"id"`
	WorkspaceID     int                                 `json:"workspace_id"`
	ProjectID       int                                 `json:"project_id"`
	TaskID          UserTimeEntriesItemsTaskID          `json:"task_id"`
	Billable        bool                                `json:"billable"`
	Start           time.Time                           `json:"start"`
	Stop            time.Time                           `json:"stop"`
	Duration        int                                 `json:"duration"`
	Description     string                              `json:"description"`
	Tags            []UserTimeEntriesItemsTagsItems     `json:"tags"`
	TagIds          []UserTimeEntriesItemsTagIdsItems   `json:"tag_ids"`
	Duronly         bool                                `json:"duronly"`
	At              time.Time                           `json:"at"`
	ServerDeletedAt UserTimeEntriesItemsServerDeletedAt `json:"server_deleted_at"`
	UserID          int                                 `json:"user_id"`
	UID             int                                 `json:"uid"`
	Wid             int                                 `json:"wid"`
	Pid             int                                 `json:"pid"`
}

// UserTimeEntriesItemsTagsItems defines a model
type UserTimeEntriesItemsTagsItems struct{}

// UserTimeEntriesItemsServerDeletedAt defines a model
type UserTimeEntriesItemsServerDeletedAt struct{}

// UserTimeEntriesItemsTagIdsItems defines a model
type UserTimeEntriesItemsTagIdsItems struct{}

// UserTimeEntriesItemsTaskID defines a model
type UserTimeEntriesItemsTaskID struct{}
