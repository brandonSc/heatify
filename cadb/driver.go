package cadb

//
// This is a custom Cloudant DBaaS driver for interfacing with the HTTP REST API
// The credentials and URL are taken from the environment variable VCAP_SERVICES
//

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var client *http.Client
var host string
var username string
var password string

func Init() {
	client = &http.Client{}
	b := []byte(os.Getenv("VCAP_SERVICES"))
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		m := f.(map[string]interface{})
		cred := m["credentials"].(map[string]interface{})
		host = cred["host"].(string)
		username = cred["username"].(string)
		password = cred["password"].(string)
	}
}

func Get(collection string, argument string) {
	req := build_request("GET", collection, argument)
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Printf("error: %s\n", err2)
	} else {
		fmt.Printf("res: %s\n", res)
	}
}

func Post(collection string, json string, argument string) {
}

func Delete() {
}

func build_request(method string, collection string, argument string) *http.Request {
	url := "https://" + host + "/" + collection
	if argument != "" {
		url += "/" + argument
	}
	req, err1 := http.NewRequest(method, url, nil)
	if err1 != nil {
		fmt.Printf("error: %s\n", err1)
		return nil
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(username, password)
	return req
}
