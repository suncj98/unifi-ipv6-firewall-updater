package web

import (
	"context"
	"log"
	"net/http"
	"unifi-ipv6-firewall-updater/config"
	"unifi-ipv6-firewall-updater/unifi"
)

func ddnsgoWebhook(w http.ResponseWriter, r *http.Request) {
	conf, err := config.GetConfigCached()
	if err != nil {
		log.Println("读取配置错误")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !conf.Webhook.Enabled {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	query := r.URL.Query()
	if query.Get("result") != "成功" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	site := query.Get("site")
	groupId := query.Get("groupId")
	ips := query["ip"]
	unifiClient := unifi.NewClient(conf.Unifi)
	err = unifiClient.UpdateFirewallGroupMembers(context.Background(), site, groupId, ips)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
