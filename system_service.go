// Get system details, and perform some management actions, such as restarting, and initiating a database migration (as part of a system upgrade).
package sonarapi

import (
	"net/http"
)

type SystemService struct {
	client *Client
}
type LogLevel string

const (
	LogLevelInfo  LogLevel = "INFO"
	LogLevelDebug LogLevel = "DEBUG"
	LogLevelTrace LogLevel = "TRACE"
)

type Cause struct {
	Message   string `json:"message,omitempty"`
	StartedAt string `json:"startedAt,omitempty"`
	State     string `json:"state,omitempty"`
}

type SystemHealthObject struct {
	Causes []*Cause `json:"causes,omitempty"`
	Health string   `json:"health,omitempty"`
	Nodes  []*Node  `json:"nodes,omitempty"`
}

type Node struct {
	Causes    []*Cause `json:"causes,omitempty"`
	Health    string   `json:"health,omitempty"`
	Host      string   `json:"host,omitempty"`
	Name      string   `json:"name,omitempty"`
	Port      int64    `json:"port,omitempty"`
	StartedAt string   `json:"startedAt,omitempty"`
	Type      string   `json:"type,omitempty"`
}

type SystemStatusObject struct {
	ID      string `json:"id,omitempty"`
	Status  string `json:"status,omitempty"`
	Version string `json:"version,omitempty"`
}

type SystemUpgradesObject struct {
	UpdateCenterRefresh string     `json:"updateCenterRefresh,omitempty"`
	Upgrades            []*Upgrade `json:"upgrades,omitempty"`
}

type Incompatible struct {
	Category         string `json:"category,omitempty"`
	Description      string `json:"description,omitempty"`
	EditionBundled   bool   `json:"editionBundled,omitempty"`
	Key              string `json:"key,omitempty"`
	License          string `json:"license,omitempty"`
	Name             string `json:"name,omitempty"`
	OrganizationName string `json:"organizationName,omitempty"`
	OrganizationURL  string `json:"organizationUrl,omitempty"`
}

type UpgradePlugins struct {
	Incompatible  []*Incompatible  `json:"incompatible,omitempty"`
	RequireUpdate []*RequireUpdate `json:"requireUpdate,omitempty"`
}

type RequireUpdate struct {
	Category              string `json:"category,omitempty"`
	Description           string `json:"description,omitempty"`
	EditionBundled        bool   `json:"editionBundled,omitempty"`
	Key                   string `json:"key,omitempty"`
	License               string `json:"license,omitempty"`
	Name                  string `json:"name,omitempty"`
	OrganizationName      string `json:"organizationName,omitempty"`
	OrganizationURL       string `json:"organizationUrl,omitempty"`
	TermsAndConditionsURL string `json:"termsAndConditionsUrl,omitempty"`
	Version               string `json:"version,omitempty"`
}

type Upgrade struct {
	ChangeLogURL string          `json:"changeLogUrl,omitempty"`
	Description  string          `json:"description,omitempty"`
	DownloadURL  string          `json:"downloadUrl,omitempty"`
	Plugins      *UpgradePlugins `json:"plugins,omitempty"`
	ReleaseDate  string          `json:"releaseDate,omitempty"`
	Version      string          `json:"version,omitempty"`
}

type SystemChangeLogLevelOption struct {
	Level LogLevel `url:"level,omitempty"` // Description:"The new level. Be cautious: DEBUG, and even more TRACE, may have performance impacts.",ExampleValue:""
}

// Health Provide health status of SonarQube.<p>Require 'Administer System' permission or authentication with passcode</p><p>  <ul> <li>GREEN: SonarQube is fully operational</li> <li>YELLOW: SonarQube is usable, but it needs attention in order to be fully operational</li> <li>RED: SonarQube is not operational</li> </ul></p>
func (s *SystemService) Health() (v *SystemHealthObject, resp *http.Response, err error) {
	path := s.client.url + "/api/system/health"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	v = new(SystemHealthObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

// Status Get state information about SonarQube.<p>status: the running status <ul> <li>STARTING: SonarQube Web Server is up and serving some Web Services (eg. api/system/status) but initialization is still ongoing</li> <li>UP: SonarQube instance is up and running</li> <li>DOWN: SonarQube instance is up but not running because migration has failed (refer to WS /api/system/migrate_db for details) or some other reason (check logs).</li> <li>RESTARTING: SonarQube instance is still up but a restart has been requested (refer to WS /api/system/restart for details).</li> <li>DB_MIGRATION_NEEDED: database migration is required. DB migration can be started using WS /api/system/migrate_db.</li> <li>DB_MIGRATION_RUNNING: DB migration is running (refer to WS /api/system/migrate_db for details)</li> </ul></p>
func (s *SystemService) Status() (v *SystemStatusObject, resp *http.Response, err error) {
	path := s.client.url + "/api/system/status"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	v = new(SystemStatusObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
