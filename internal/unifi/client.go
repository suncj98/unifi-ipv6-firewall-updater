package unifi

import (
	"context"
	"crypto/tls"
	"github.com/paultyng/go-unifi/unifi"
	"log"
	"net/http"
	"net/http/cookiejar"
)

type Config struct {
	Endpoint string `yaml:"endpoint"`
	Site     string `yaml:"site"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Client struct {
	config *Config
	client *unifi.Client
}

func NewClient(config *Config) *Client {
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
	if err := client.SetBaseURL(config.Endpoint); err != nil {
		log.Fatalln(err)
	}
	if err := client.Login(context.Background(), config.Username, config.Password); err != nil {
		log.Fatalln(err)
	}
	return &Client{
		config: config,
		client: client,
	}
}

func (c *Client) UpdateFirewallGroupMembers(ctx context.Context, groupId string, members []string) error {
	group, err := c.client.GetFirewallGroup(ctx, c.config.Site, groupId)
	if err != nil {
		return err
	}
	log.Println("Before updating:", group)

	group.GroupMembers = members
	group, err = c.client.UpdateFirewallGroup(ctx, c.config.Site, group)
	if err != nil {
		return err
	}
	log.Println("After updating:", group)
	return nil
}
