package main

import (
	"log"
	"net/http"
	"os"
	//for extracting service credentials from VCAP_SERVICES
	//"github.com/cloudfoundry-community/go-cfenv"
	"hub.jazz.net/git/schurman93/Git-Monitor/gitutil"
	"hub.jazz.net/git/schurman93/Git-Monitor/route"
)

const (
	DEFAULT_PORT = "8080"
	DEFAULT_HOST = ""
)

func main() {
	// setup host and port for server
	var port string
	if port = os.Getenv("VCAP_APP_PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}

	var host string
	if host = os.Getenv("VCAP_APP_HOST"); len(host) == 0 {
		host = DEFAULT_HOST
	}

	// grab the router and request handlers
	router := route.NewRouter()

	// run the git update loop in the background
	go gitutil.RunUpdateLoop()

	// launch the server
	log.Printf("Starting app on %+v:%+v\n", host, port)
	log.Fatal(http.ListenAndServe(host+":"+port, router))
}
