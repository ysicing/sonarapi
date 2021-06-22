package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"os"

	"github.com/ysicing/sonarapi"
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

func searchtoken() {
	s := sonarapi.UserTokensSearchOption{}
	v, resp, err := client.UserTokens.Search(&s)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func generate(name string) {
	s := sonarapi.UserTokensGenerateOption{
		Name: name,
	}
	v, resp, err := client.UserTokens.Generate(&s)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func revoke(name string) {
	s := sonarapi.UserTokensRevokeOption{
		Name: name,
	}
	resp, err := client.UserTokens.Revoke(&s)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
}

func main() {
	generate("111111")
	generate("111112")
	searchtoken()
	revoke("111112")
}
