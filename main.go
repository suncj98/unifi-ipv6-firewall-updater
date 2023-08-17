package main

import (
	"context"
	"log"
	"net/http"
	"os"
)

var (
	address  = "0.0.0.0:28765"
	endpoint = ""
	username = ""
	password = ""
)

func init() {
	if val := os.Getenv("ADDRESS"); val != "" {
		address = val
	}

	if val := os.Getenv("UNIFI_ENDPOINT"); val != "" {
		endpoint = val
	} else {
		log.Fatalln("Env UNIFI_ENDPOINT is required.")
	}

	if val := os.Getenv("UNIFI_USERNAME"); val != "" {
		username = val
	} else {
		log.Fatalln("Env UNIFI_USERNAME is required.")
	}

	if val := os.Getenv("UNIFI_PASSWORD"); val != "" {
		password = val
	} else {
		log.Fatalln("Env UNIFI_PASSWORD is required.")
	}
}

func main() {
	http.HandleFunc("/ddnsgo", ddnsgoWebhook)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func ddnsgoWebhook(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if query.Get("result") != "成功" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	site := query.Get("site")
	groupId := query.Get("groupId")
	ips := query["ip"]
	unifiClient := NewUnifiClient(endpoint, username, password)
	err := unifiClient.UpdateFirewallGroupMembers(context.Background(), site, groupId, ips)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
