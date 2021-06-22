// Manage branch (only available when the Branch plugin is installed)
package sonarapi

import (
	"net/http"
)

type ProjectBranchesService struct {
	client *Client
}

type ProjectBranchesListObject struct {
	Branches []*Branch `json:"branches,omitempty"`
}

type Branch struct {
	AnalysisDate      string  `json:"analysisDate,omitempty"`
	IsMain            bool    `json:"isMain,omitempty"`
	MergeBranch       string  `json:"mergeBranch,omitempty"`
	Name              string  `json:"name,omitempty"`
	Status            *Status `json:"status,omitempty"`
	Type              string  `json:"type,omitempty"`
	ExcludedFromPurge bool    `json:"excludedFromPurge,omitempty"`
}

type Status struct {
	Bugs              int64  `json:"bugs,omitempty"`
	CodeSmells        int64  `json:"codeSmells,omitempty"`
	QualityGateStatus string `json:"qualityGateStatus,omitempty"`
	Vulnerabilities   int64  `json:"vulnerabilities,omitempty"`
}

type ProjectBranchesListOption struct {
	Project string `url:"project,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// List List the branches of a project.<br/>Requires 'Browse' or 'Execute analysis' rights on the specified project.
func (s *ProjectBranchesService) List(opt *ProjectBranchesListOption) (v *ProjectBranchesListObject, resp *http.Response, err error) {
	path := s.client.url + "/api/project_branches/list?project=" + opt.Project
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	s.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(ProjectBranchesListObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
