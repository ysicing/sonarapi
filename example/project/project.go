// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/ysicing/sonarapi"
	"k8s.io/klog/v2"
	"os"
	"time"
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

func listbranch(name string) {
	s := sonarapi.ProjectBranchesListOption{Project: name}
	v, resp, err := client.ProjectBranches.List(&s)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func createproject(name, mode string) {
	s := sonarapi.ProjectsCreateOption{
		Name:       name,
		Project:    fmt.Sprintf("api-%v-%v", mode, name),
		Visibility: "public",
	}
	v, resp, err := client.Projects.Create(&s)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func sdproject(name string) {
	s := sonarapi.ProjectsSearchOption{}
	if len(name) > 0 {
		s.Q = name
	}
	v, resp, err := client.Projects.Search(&s)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
	for _, c := range v.Components {
		listbranch(c.Key)
		ds := sonarapi.ProjectsDeleteOption{Project: c.Key}
		delresp, err := client.Projects.Delete(&ds)
		if err != nil {
			klog.Error(err)
			continue
		}
		if delresp.StatusCode >= 400 {
			klog.Infof("%v(%v) delete err, code: %v", c.Name, c.Key, delresp.StatusCode)
			continue
		}
		klog.Infof("%v(%v) delete done", c.Name, c.Key)
	}
}

func main() {
	// listbranch("demo")
	createproject(time.Now().Format("20060102150405"), "ci")
	time.Sleep(5 * time.Second)
	createproject(time.Now().Format("20060102150405"), "ci")
	sdproject("2")
}
