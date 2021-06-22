// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/ysicing/sonarapi"
	"os"
)

var client *sonarapi.Client

func init() {
	sonarURL := os.Getenv("SONAR_URL")
	if sonarURL == "" {
		sonarURL = "http://172.16.16.55:9000"
	}
	c, err := sonarapi.NewClient(sonarURL, "admin", "12345678")
	if err != nil {
		panic(err)
	}
	client = c

}

func list() {
	s := sonarapi.WebhooksListOption{}
	v, resp, err := client.Webhooks.List(&s)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func create(name, url, token string) {
	s := sonarapi.WebhooksCreateOption{
		Name: name,
		Url:  url,
	}
	if len(token) != 0 {
		s.Secret = token
	}
	v, resp, err := client.Webhooks.Create(&s)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func delete(key string) {
	s := sonarapi.WebhooksDeleteOption{
		Webhook: key,
	}

	resp, err := client.Webhooks.Delete(&s)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
}

func main() {
	create("hook33", "http://10.0.0.1", "")
	delete("AXotQZ52zz_xwUCokbU7")
	create("hook44", "http://10.0.0.1", "")
	delete("AXotPtnLzz_xwUCokbU6")
	list()
}
