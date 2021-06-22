// List, create, and delete a user's access tokens.
package sonarapi

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"k8s.io/klog/v2"
	"net/http"
	"strings"
)

type UserTokensService struct {
	client *Client
}

type UserTokensGenerateObject struct {
	CreatedAt string `json:"createdAt,omitempty"`
	Login     string `json:"login,omitempty"`
	Name      string `json:"name,omitempty"`
	Token     string `json:"token,omitempty"`
}

type UserTokensSearchObject struct {
	Login      string       `json:"login,omitempty"`
	UserTokens []*UserToken `json:"userTokens,omitempty"`
}

type UserToken struct {
	CreatedAt string `json:"createdAt,omitempty"`
	Name      string `json:"name,omitempty"`
}

type UserTokensGenerateOption struct {
	Login string `url:"login,omitempty"` // Description:"User login. If not set, the token is generated for the authenticated user.",ExampleValue:"g.hopper"
	Name  string `url:"name,omitempty"`  // Description:"Token name",ExampleValue:"Project scan on Travis"
}

// Generate Generate a user access token. <br />Please keep your tokens secret. They enable to authenticate and analyze projects.<br />If the login is set, it requires administration permissions. Otherwise, a token is generated for the authenticated user.
func (s *UserTokensService) Generate(opt *UserTokensGenerateOption) (v *UserTokensGenerateObject, resp *http.Response, err error) {
	path := s.client.url + "/api/user_tokens/generate"
	optv, _ := query.Values(opt)
	req, err := http.NewRequest("POST", path, strings.NewReader(optv.Encode()))
	if err != nil {
		klog.Error(err)
		return
	}
	s.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	v = new(UserTokensGenerateObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type UserTokensRevokeOption struct {
	Login string `url:"login,omitempty"` // Description:"User login",ExampleValue:"g.hopper"
	Name  string `url:"name,omitempty"`  // Description:"Token name",ExampleValue:"Project scan on Travis"
}

// Revoke Revoke a user access token. <br/>If the login is set, it requires administration permissions. Otherwise, a token is generated for the authenticated user.
func (s *UserTokensService) Revoke(opt *UserTokensRevokeOption) (resp *http.Response, err error) {
	path := s.client.url + "/api/user_tokens/revoke"
	optv, _ := query.Values(opt)
	req, err := http.NewRequest("POST", path, strings.NewReader(optv.Encode()))
	if err != nil {
		klog.Error(err)
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

type UserTokensSearchOption struct {
	Login string `url:"login,omitempty"` // Description:"User login",ExampleValue:"g.hopper"
}

// Search List the access tokens of a user.<br>The login must exist and active.<br>If the login is set, it requires administration permissions. Otherwise, a token is generated for the authenticated user.
func (s *UserTokensService) Search(opt *UserTokensSearchOption) (v *UserTokensSearchObject, resp *http.Response, err error) {
	path := s.client.url + "/api/user_tokens/search"
	if len(opt.Login) != 0 {
		path = fmt.Sprintf("%v?login=%v", path, opt.Login)
	}

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		klog.Error(err)
		return
	}
	s.client.requestExtHeader(req)
	v = new(UserTokensSearchObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		klog.Error(err)
		return nil, resp, err
	}
	return
}
