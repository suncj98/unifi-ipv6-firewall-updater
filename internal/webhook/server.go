package webhook

import (
	"context"
	"log"
	"net/http"
	"unifi-ipv6-firewall-updater/internal/conf"
	"unifi-ipv6-firewall-updater/internal/unifi"
)

var unifiConfig *unifi.Config

func Run(uc *unifi.Config, sc *conf.HttpServerConfig) {
	unifiConfig = uc
	http.HandleFunc("/ddnsgo", ddnsgoWebhook)
	err := http.ListenAndServe(sc.Address, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func ddnsgoWebhook(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	groupId := query.Get("groupId")
	ips := query["ip"]
	if query.Get("result") != "成功" {
		return
	}
	unifiClient := unifi.NewClient(unifiConfig)
	err := unifiClient.UpdateFirewallGroupMembers(context.Background(), groupId, ips)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
}
