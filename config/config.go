package config

import (
	"unifi-ipv6-firewall-updater/dns"
	"unifi-ipv6-firewall-updater/unifi"
)

type Config struct {
	Unifi   unifi.Config `yaml:"unifi"`
	Dns     Dns          `yaml:"dns"`
	Webhook Webhook      `yaml:"webhook"`
}

type Dns struct {
	Resolver dns.Config `yaml:"resolver"`
	Groups   []struct {
		Site  string   `yaml:"site"`
		Id    string   `yaml:"id"`
		Hosts []string `yaml:"hosts"`
	} `yaml:"groups"`
}

type Webhook struct {
	Enabled bool `yaml:"enabled"`
}
