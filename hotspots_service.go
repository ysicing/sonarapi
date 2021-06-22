// Get the list of programming languages supported in this instance.
package sonarapi

import (
	"github.com/google/go-querystring/query"
	"k8s.io/klog/v2"
	"net/http"
	"strings"
)

type HotspotsService struct {
	client *Client
}

type HotspotsShowObject struct {
	Key       string `json:"key"`
	Component struct {
		Key       string `json:"key"`
		Qualifier string `json:"qualifier"`
		Name      string `json:"name"`
		LongName  string `json:"longName"`
		Path      string `json:"path"`
	} `json:"component"`
	Project struct {
		Key       string `json:"key"`
		Qualifier string `json:"qualifier"`
		Name      string `json:"name"`
		LongName  string `json:"longName"`
	} `json:"project"`
	Rule struct {
		Key                      string `json:"key"`
		Name                     string `json:"name"`
		SecurityCategory         string `json:"securityCategory"`
		VulnerabilityProbability string `json:"vulnerabilityProbability"`
	} `json:"rule"`
	Status       string `json:"status"`
	Line         int    `json:"line"`
	Hash         string `json:"hash"`
	Message      string `json:"message"`
	Assignee     string `json:"assignee"`
	Author       string `json:"author"`
	CreationDate string `json:"creationDate"`
	UpdateDate   string `json:"updateDate"`
	Changelog    []struct {
		User         string `json:"user"`
		UserName     string `json:"userName"`
		CreationDate string `json:"creationDate"`
		Diffs        []struct {
			Key      string `json:"key"`
			NewValue string `json:"newValue"`
			OldValue string `json:"oldValue"`
		} `json:"diffs"`
		Avatar       string `json:"avatar"`
		IsUserActive bool   `json:"isUserActive"`
	} `json:"changelog"`
	Comment []struct {
		Key       string `json:"key"`
		Login     string `json:"login"`
		HTMLText  string `json:"htmlText"`
		Markdown  string `json:"markdown"`
		CreatedAt string `json:"createdAt"`
	} `json:"comment"`
	Users []struct {
		Login  string `json:"login"`
		Name   string `json:"name"`
		Active bool   `json:"active"`
	} `json:"users"`
	CanChangeStatus bool `json:"canChangeStatus"`
}

type Hotspots struct {
	Hotspot string `json:"hotspot,omitempty"`
}

type HotspotsShowOption struct {
	Hotspot string `url:"hotspot,omitempty"` // Description:"The size of the list to return, 0 for all languages",ExampleValue:"25"
}

// List List supported programming languages
func (s *HotspotsService) Show(opt *HotspotsShowOption) (v *HotspotsShowObject, resp *http.Response, err error) {
	path := s.client.url + "/api/hotspots/show"
	optv, _ := query.Values(opt)
	req, err := http.NewRequest("GET", path, strings.NewReader(optv.Encode()))
	if err != nil {
		klog.Error(err)
		return
	}
	s.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	v = new(HotspotsShowObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
