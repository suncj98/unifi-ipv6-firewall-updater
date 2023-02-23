package conf

import (
	"unifi-ipv6-firewall-updater/internal/dns"
	"unifi-ipv6-firewall-updater/internal/unifi"
)

type Bootstrap struct {
	Unifi *unifi.Config `yaml:"unifi"`
	Dns   *Dns          `yaml:"dns"`
}

type Dns struct {
	Enabled  bool        `yaml:"enabled"`
	Cron     string      `yaml:"cron"`
	Resolver *dns.Config `yaml:"resolver"`
	Groups   []*Group    `yaml:"groups"`
}

type Group struct {
	Id    string   `yaml:"id"`
	Hosts []string `yaml:"hosts"`
}
