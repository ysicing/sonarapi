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

func show(name string) {
	s := sonarapi.HotspotsShowOption{Hotspot: name}
	v, resp, err := client.Hotspots.Show(&s)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func main() {
	show("demo")
}
