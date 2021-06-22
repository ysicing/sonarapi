// Webhooks allow to notify external services when a project analysis is done
package sonarapi

import (
	"github.com/google/go-querystring/query"
	"k8s.io/klog/v2"
	"net/http"
	"strings"
)

type WebhooksService struct {
	client *Client
}

type WebhooksCreateObject struct {
	Webhook *Webhook `json:"webhook,omitempty"`
}

type Webhook struct {
	Key            string    `json:"key,omitempty"`
	Name           string    `json:"name,omitempty"`
	URL            string    `json:"url,omitempty"`
	LatestDelivery *Delivery `json:"latestDelivery,omitempty"`
}

type WebhooksDeliveriesObject struct {
	Deliveries []*Delivery `json:"deliveries,omitempty"`
	Paging     Paging      `json:"paging,omitempty"`
}

type Delivery struct {
	At              string `json:"at,omitempty"`
	CeTaskID        string `json:"ceTaskId,omitempty"`
	ComponentKey    string `json:"componentKey,omitempty"`
	DurationMs      int64  `json:"durationMs,omitempty"`
	HTTPStatus      int64  `json:"httpStatus,omitempty"`
	ID              string `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Payload         string `json:"payload,omitempty"`
	Success         bool   `json:"success,omitempty"`
	URL             string `json:"url,omitempty"`
	ErrorStackTrace string `json:"errorStacktrace,omitempty"`
}

type WebhooksDeliveryObject struct {
	Delivery *Delivery `json:"delivery,omitempty"`
}

type WebhooksListObject struct {
	Webhooks []*Webhook `json:"webhooks,omitempty"`
}

type WebhooksCreateOption struct {
	Name   string `url:"name,omitempty"` // Description:"Name displayed in the administration console of webhooks",ExampleValue:"My Webhook"
	Url    string `url:"url,omitempty"`  // Description:"Server endpoint that will receive the webhook payload, for example 'http://my_server/foo'. If HTTP Basic authentication is used, HTTPS is recommended to avoid man in the middle attacks. Example: 'https://myLogin:myPassword@my_server/foo'",ExampleValue:"https://www.my-webhook-listener.com/sonar"
	Secret string `url:"secret,omitempty"`
}

// Create Create a Webhook.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
func (s *WebhooksService) Create(opt *WebhooksCreateOption) (v *WebhooksCreateObject, resp *http.Response, err error) {
	path := s.client.url + "/api/webhooks/create"
	optv, _ := query.Values(opt)
	req, err := http.NewRequest("POST", path, strings.NewReader(optv.Encode()))
	if err != nil {
		klog.Error(err)
		return
	}
	s.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	v = new(WebhooksCreateObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type WebhooksDeleteOption struct {
	Webhook string `url:"webhook,omitempty"` // Description:"The key of the webhook to be deleted,auto-generated value can be obtained through api/webhooks/create or api/webhooks/list",ExampleValue:"my_project"
}

// Delete Delete a Webhook.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
func (s *WebhooksService) Delete(opt *WebhooksDeleteOption) (resp *http.Response, err error) {
	path := s.client.url + "/api/webhooks/delete"
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

type WebhooksListOption struct {
	//Project string `url:"project,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// List Search for global webhooks or project webhooks. Webhooks are ordered by name.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
func (s *WebhooksService) List(opt *WebhooksListOption) (v *WebhooksListObject, resp *http.Response, err error) {
	path := s.client.url + "/api/webhooks/list"
	optv, _ := query.Values(opt)
	req, err := http.NewRequest("GET", path, strings.NewReader(optv.Encode()))
	if err != nil {
		klog.Error(err)
		return
	}
	s.client.requestExtHeader(req)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	v = new(WebhooksListObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type WebhooksUpdateOption struct {
	Name    string `url:"name,omitempty"`    // Description:"new name of the webhook",ExampleValue:"My Webhook"
	Url     string `url:"url,omitempty"`     // Description:"new url to be called by the webhook",ExampleValue:"https://www.my-webhook-listener.com/sonar"
	Webhook string `url:"webhook,omitempty"` // Description:"The key of the webhook to be updated,auto-generated value can be obtained through api/webhooks/create or api/webhooks/list",ExampleValue:"my_project"
	Secret  string `url:"secret,omitempty"`
}

// Update Update a Webhook.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
func (s *WebhooksService) Update(opt *WebhooksUpdateOption) (resp *http.Response, err error) {
	path := s.client.url + "/api/webhooks/update"
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
