package web

import (
	"log"
	"net/http"
	"os"
)

const webAddress = "WEB_ADDRESS"
const defaultWebAddress = "0.0.0.0:28765"

func Run() {
	http.HandleFunc("/ddnsgo", ddnsgoWebhook)
	err := http.ListenAndServe(getWebAddress(), nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// getWebAddress 获得web服务监听地址
func getWebAddress() string {
	if address := os.Getenv(webAddress); address != "" {
		return address
	}
	return defaultWebAddress
}
