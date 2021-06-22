// Get components or children with specified measures.
package sonarapi

import (
	"net/http"
)

type MeasuresService struct {
	client *Client
}

type MeasuresSearchObject struct {
	Measures []*SonarMeasure `json:"measures,omitempty"`
}

type SonarMeasure struct {
	Metric    string     `json:"metric,omitempty"`
	Value     string     `json:"value,omitempty"`
	Periods   []*Period  `json:"periods,omitempty"`
	Component string     `json:"component,omitempty"`
	Histories []*History `json:"history,omitempty"`
	BestValue bool       `json:"bestValue,omitempty"`
}

type Period struct {
	Date      string `json:"date,omitempty"`
	Index     int64  `json:"index,omitempty"`
	Mode      string `json:"mode,omitempty"`
	Parameter string `json:"parameter,omitempty"`
	Value     string `json:"value,omitempty"`
	BestValue bool   `json:"bestValue,omitempty"`
}

type History struct {
	Date  string `json:"date,omitempty"`
	Value string `json:"value,omitempty"`
}

type MeasuresSearchOption struct {
	MetricKeys  string `url:"metricKeys,omitempty"`
	ProjectKeys string `url:"projectKeys,omitempty"`
}

func (s *MeasuresService) Search(opt *MeasuresSearchOption) (v *MeasuresSearchObject, resp *http.Response, err error) {
	//if len(opt.MetricKeys) == 0 {
	//	opt.MetricKeys = "alert_status,bugs,reliability_rating,vulnerabilities,security_rating,security_hotspots_reviewed,security_review_rating,code_smells,sqale_rating,duplicated_lines_density,coverage,ncloc,ncloc_language_distribution,projects"
	//}
	path := s.client.url + "/api/measures/search?metricKeys=alert_status,bugs,reliability_rating,vulnerabilities,security_rating,security_hotspots_reviewed,security_review_rating,code_smells,sqale_rating,duplicated_lines_density,coverage,ncloc,ncloc_language_distribution,projects&projectKeys=" + opt.ProjectKeys
	// optv, _ := query.Values(opt)
	//req, err := http.NewRequest("GET", path, strings.NewReader(optv.Encode()))
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	s.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/json")
	v = new(MeasuresSearchObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
