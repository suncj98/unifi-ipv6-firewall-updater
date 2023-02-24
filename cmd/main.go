package main

import (
	"flag"
	"unifi-ipv6-firewall-updater/internal/conf"
	"unifi-ipv6-firewall-updater/internal/job"
	"unifi-ipv6-firewall-updater/internal/webhook"
	"unifi-ipv6-firewall-updater/pkg/config"
)

var (
	flagConf string
)

func init() {
	flag.StringVar(&flagConf, "conf", "./configs", "config path, eg: --conf=./configs")
}

func main() {
	flag.Parse()
	var c conf.Bootstrap
	config.Init(flagConf, &c)

	if c.Dns.Enabled {
		go job.Run(c.Unifi, c.Dns)
	}
	if c.Webhook.Enabled {
		go webhook.Run(c.Unifi, c.Webhook.Server)
	}
	select {}
}
