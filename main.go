package main

import (
	"unifi-ipv6-firewall-updater/job"
	"unifi-ipv6-firewall-updater/web"
)

func main() {
	go web.Run()
	go job.RunTimer()
	select {}
}
