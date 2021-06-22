// Manage project existence.
package sonarapi

import (
	"github.com/google/go-querystring/query"
	"net/http"
	"strings"
)

type ProjectsService struct {
	client *Client
}

const (
	ProjectVisibilityPublic  = "public"
	ProjectVisibilityPrivate = "private"
)

type ProjectsBulkUpdateKeyObject struct {
	Keys []*Key `json:"keys,omitempty"`
}

type Key struct {
	Duplicate bool   `json:"duplicate,omitempty"`
	Key       string `json:"key,omitempty"`
	NewKey    string `json:"newKey,omitempty"`
}

type Project struct {
	CreationDate string `json:"creationDate,omitempty"`
	Key          string `json:"key,omitempty"`
	Name         string `json:"name,omitempty"`
	Qualifier    string `json:"qualifier,omitempty"`
	UUID         string `json:"uuid,omitempty"`
	Visibility   string `json:"visibility,omitempty"`
}

type Component struct {
	AnalysisDate     string          `json:"analysisDate,omitempty"`
	Description      string          `json:"description,omitempty"`
	Enabled          bool            `json:"enabled,omitempty"`
	ID               string          `json:"id,omitempty"`
	Key              string          `json:"key,omitempty"`
	Language         string          `json:"language,omitempty"`
	LastAnalysisDate string          `json:"lastAnalysisDate,omitempty"`
	LeakPeriodDate   string          `json:"leakPeriodDate,omitempty"`
	LongName         string          `json:"longName,omitempty"`
	Measures         []*SonarMeasure `json:"measures,omitempty"`
	Name             string          `json:"name,omitempty"`
	Organization     string          `json:"organization,omitempty"`
	Path             string          `json:"path,omitempty"`
	Project          string          `json:"project,omitempty"`
	Qualifier        string          `json:"qualifier,omitempty"`
	Tags             []string        `json:"tags,omitempty"`
	UUID             string          `json:"uuid,omitempty"`
	Version          string          `json:"version,omitempty"`
	Revision         string          `json:"revision,omitempty"`
	Visibility       string          `json:"visibility,omitempty"`
}

type ComponentsSearchObject struct {
	Components []*Component `json:"components,omitempty"`
	Paging     *Paging      `json:"paging,omitempty"`
}

type ProjectsCreateObject struct {
	Project *Project `json:"project,omitempty"`
}

type ProjectsCreateOption struct {
	Name       string `url:"name,omitempty"`       // Description:"Name of the project",ExampleValue:"SonarQube"
	Project    string `url:"project,omitempty"`    // Description:"Key of the project",ExampleValue:"my_project"
	Visibility string `url:"visibility,omitempty"` // Description:"Whether the created project should be visible to everyone, or only specific user/groups.<br/>If no visibility is specified, the default project visibility of the organization will be used.",ExampleValue:""
}

// Create Create a project.<br/>Requires 'Create Projects' permission
func (s *ProjectsService) Create(opt *ProjectsCreateOption) (v *ProjectsCreateObject, resp *http.Response, err error) {
	path := s.client.url + "/api/projects/create"
	optv, _ := query.Values(opt)
	req, err := http.NewRequest("POST", path, strings.NewReader(optv.Encode()))
	if err != nil {
		return
	}
	s.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	v = new(ProjectsCreateObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type ProjectsDeleteOption struct {
	Project   string `url:"project,omitempty"`   // Description:"Project key",ExampleValue:"my_project"
	ProjectId string `url:"projectId,omitempty"` // Description:"Project ID",ExampleValue:"ce4c03d6-430f-40a9-b777-ad877c00aa4d"
}

// Delete Delete a project.<br> Requires 'Administer System' permission or 'Administer' permission on the project.
func (s *ProjectsService) Delete(opt *ProjectsDeleteOption) (resp *http.Response, err error) {
	path := s.client.url + "/api/projects/delete"
	optv, _ := query.Values(opt)
	req, err := http.NewRequest("POST", path, strings.NewReader(optv.Encode()))
	if err != nil {
		return
	}
	s.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type ProjectsSearchOption struct {
	AnalyzedBefore    string `url:"analyzedBefore,omitempty"`    // Description:"Filter the projects for which last analysis is older than the given date (exclusive).<br> Either a date (server timezone) or datetime can be provided.",ExampleValue:"2017-10-19 or 2017-10-19T13:00:00+0200"
	OnProvisionedOnly string `url:"onProvisionedOnly,omitempty"` // Description:"Filter the projects that are provisioned",ExampleValue:""
	P                 string `url:"p,omitempty"`                 // Description:"1-based page number",ExampleValue:"42"
	ProjectIds        string `url:"projectIds,omitempty"`        // Description:"Comma-separated list of project ids",ExampleValue:"AU-Tpxb--iU5OvuD2FLy,AU-TpxcA-iU5OvuD2FLz"
	Projects          string `url:"projects,omitempty"`          // Description:"Comma-separated list of project keys",ExampleValue:"my_project,another_project"
	Ps                string `url:"ps,omitempty"`                // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	Q                 string `url:"q,omitempty"`                 // Description:"Limit search to: <ul><li>component names that contain the supplied string</li><li>component keys that contain the supplied string</li></ul>",ExampleValue:"sonar"
	Qualifiers        string `url:"qualifiers,omitempty"`        // Description:"Comma-separated list of component qualifiers. Filter the results with the specified qualifiers",ExampleValue:""
}
type ProjectSearchObject ComponentsSearchObject

// Search Search for projects or views to administrate them.<br>Requires 'System Administrator' permission
func (s *ProjectsService) Search(opt *ProjectsSearchOption) (v *ProjectSearchObject, resp *http.Response, err error) {
	path := s.client.url + "/api/projects/search"
	optv, _ := query.Values(opt)
	req, err := http.NewRequest("GET", path, strings.NewReader(optv.Encode()))
	if err != nil {
		return
	}
	s.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	v = new(ProjectSearchObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
