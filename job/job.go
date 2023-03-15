package job

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"
	"unifi-ipv6-firewall-updater/config"
	"unifi-ipv6-firewall-updater/dns"
	"unifi-ipv6-firewall-updater/unifi"
)

const dnsInterval = "DNS_INTERVAL"
const firstDelay = 10 * time.Second

func RunTimer() {
	interval := getInterval()
	if interval.Seconds() < 1 {
		return
	}
	time.Sleep(firstDelay)
	for {
		RunOnce()
		time.Sleep(interval)
	}
}

func RunOnce() {
	conf, err := config.GetConfigCached()
	if err != nil {
		log.Println("读取配置错误")
		return
	}
	dnsResolver := dns.NewResolver(conf.Dns.Resolver)
	unifiClient := unifi.NewClient(conf.Unifi)
	for _, group := range conf.Dns.Groups {
		ips := dnsResolver.LookupIPv6s(context.Background(), group.Hosts)
		if len(ips) == 0 {
			log.Println("IPv6 address not found", group.Hosts)
			continue
		}
		err = unifiClient.UpdateFirewallGroupMembers(context.Background(), group.Site, group.Id, ips)
		if err != nil {
			log.Println(err)
		}
	}
}

// getInterval 获得定时更新间隔时间
func getInterval() time.Duration {
	intervalStr := os.Getenv(dnsInterval)
	if intervalStr == "" {
		return 0
	}
	interval, err := strconv.Atoi(intervalStr)
	if err != nil {
		log.Fatalln("Environment variable", dnsInterval, "is invalid")
	}
	return time.Duration(interval) * time.Second
}
