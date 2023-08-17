package main

import (
	"context"
	"crypto/tls"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/paultyng/go-unifi/unifi"
	"log"
	"net/http"
	"net/http/cookiejar"
)

type UnifiClient struct {
	client *unifi.Client
}

func NewUnifiClient(endpoint, username, password string) *UnifiClient {
	jar, _ := cookiejar.New(nil)
	httpClient := &http.Client{
		Jar: jar,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	client := &unifi.Client{}
	if err := client.SetHTTPClient(httpClient); err != nil {
		log.Fatalln(err)
	}
	if err := client.SetBaseURL(endpoint); err != nil {
		log.Fatalln(err)
	}
	if err := client.Login(context.Background(), username, password); err != nil {
		log.Fatalln(err)
	}
	return &UnifiClient{
		client: client,
	}
}

func (c *UnifiClient) UpdateFirewallGroupMembers(ctx context.Context, site, groupId string, members []string) error {
	group, err := c.client.GetFirewallGroup(ctx, site, groupId)
	if err != nil {
		return err
	}
	log.Println("Before updating:", group)

	oldMemberSet := mapset.NewSet[string](group.GroupMembers...)
	newMemberSet := mapset.NewSet[string](members...)
	if oldMemberSet.Equal(newMemberSet) {
		log.Println("Group members not modified, skip updating.")
		return nil
	}
	group.GroupMembers = newMemberSet.ToSlice()
	group, err = c.client.UpdateFirewallGroup(ctx, site, group)
	if err != nil {
		return err
	}
	log.Println("After updating:", group)
	return nil
}
