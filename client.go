package sonarapi

import (
	"net/http"
	"net/url"
)

type Client struct {
	baseURL                        *url.URL
	username, password, token, url string
	authType                       authType
	httpClient                     *http.Client
	Hotspots                       *HotspotsService
	Measures                       *MeasuresService
	ProjectBranches                *ProjectBranchesService
	Projects                       *ProjectsService
	System                         *SystemService
	UserTokens                     *UserTokensService
	Webhooks                       *WebhooksService
}

func NewClient(endpoint, username, password string) (*Client, error) {
	c := &Client{username: username, password: password, authType: basicAuth, httpClient: http.DefaultClient}
	if endpoint == "" {
		endpoint = defaultBaseURL
	}
	c.SetBaseURL(endpoint)
	c.Hotspots = &HotspotsService{client: c}
	c.Measures = &MeasuresService{client: c}
	c.ProjectBranches = &ProjectBranchesService{client: c}
	c.Projects = &ProjectsService{client: c}
	c.System = &SystemService{client: c}
	c.UserTokens = &UserTokensService{client: c}
	c.Webhooks = &WebhooksService{client: c}
	return c, nil
}

func NewClientByToken(endpoint, token string) (*Client, error) {
	c, err := NewClient(endpoint, "", "")
	c.token = token
	c.authType = privateToken
	return c, err
}
