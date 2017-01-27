package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/SchweizerischeBundesbahnen/openshift-monitoring/hub/server"
)

var uiAddr = flag.String("UI_ADDR", "localhost:8080", "http service endpoint")
var hubAddr = flag.String("RPC_ADDR", "localhost:2600", "go hub rcp2 address")
var masterApiUrls = flag.String("MASTER_API_URLS", "https://master1:8443,https://master2:8443", "addresses of master api's")
var deamonPublicUrl = flag.String("DEAMON_PUBLIC_URL", "http://deamon.yourroute.com", "external address of the deamon service (route)")

func main() {
	flag.Parse()
	log.Println("hub waiting for deamons on ", *hubAddr)
	log.Println("ui server waiting for websocket on ", *uiAddr)
	log.Println("master api urls are ", *masterApiUrls)
	log.Println("deamons public url is ", *deamonPublicUrl)

	// Start hub rcp server
	hub := server.NewHub(*hubAddr, *masterApiUrls, *deamonPublicUrl)
	go hub.Serve()

	// Start websocket server for ui
	http.HandleFunc("/ui", func(w http.ResponseWriter, r *http.Request) {
		server.OnUISocket(hub, w, r)
	})

	log.Fatal(http.ListenAndServe(*uiAddr, nil))
}

