package job

import (
	"context"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"unifi-ipv6-firewall-updater/internal/conf"
	"unifi-ipv6-firewall-updater/internal/dns"
	"unifi-ipv6-firewall-updater/internal/unifi"
)

func Run(unifiConfig *unifi.Config, dnsConfig *conf.Dns) {
	cr := cron.New(cron.WithChain(cron.SkipIfStillRunning(
		cron.VerbosePrintfLogger(
			log.New(os.Stdout, "cron: ", log.LstdFlags),
		),
	)))
	_, err := cr.AddFunc(dnsConfig.Cron, func() {
		unifiClient := unifi.NewClient(unifiConfig)
		dnsResolver := dns.NewResolver(dnsConfig.Resolver)
		for _, group := range dnsConfig.Groups {
			ips := dnsResolver.LookupIPv6s(context.Background(), group.Hosts)
			if len(ips) == 0 {
				log.Println("IPv6 address not found", group.Hosts)
				continue
			}
			err := unifiClient.UpdateFirewallGroupMembers(context.Background(), group.Id, ips)
			if err != nil {
				log.Println(err)
			}
		}
	})
	if err != nil {
		log.Fatalln("Init cron error.")
	}
	cr.Start()
}
