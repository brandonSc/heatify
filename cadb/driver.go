package cadb

//
// This is a custom Cloudant DBaaS driver for interfacing with the HTTP REST API
// The credentials and URL are taken from the environment variable VCAP_SERVICES
//

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

//
// `client` can be used to send complex requests
// it is needed to do things like setting the Content-Type
// and setting the authentication parameters
//
var client *http.Client

//
// the host address, username and password for our Cloudant DBaaS
// they are read from VCAP_SERVICES in the environment
//
var host string
var username string
var password string

//
// Set the connection and credentials from VCAP_SERVICES
//
func Init() {
	client = &http.Client{}
	b := []byte(os.Getenv("VCAP_SERVICES"))
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Printf("error, cadb.Init: %s\n", err)
	} else {
		m := f.(map[string]interface{})
		cred := m["credentials"].(map[string]interface{})
		host = cred["host"].(string)
		username = cred["username"].(string)
		password = cred["password"].(string)
	}
}

//
// Issue a GET request to Cloudant
// on the given `collection` to the provided `apiEndpoint`
// e.g. the apiEndpoint `_all_docs` can retrieve all IDs of the `collection` in a GET request
//
func Get(collection string, apiEndpoint string) (*http.Response, error) {
	req := build_request("GET", collection, "", apiEndpoint)
	res, err := client.Do(req)
	return res, err
}

//
// Submit a POST to Cloudant with the `json` body to the `apiEndpoint`
//
func Post(collection string, json string, apiEndpoint string) (*http.Response, error) {
	req := build_request("POST", collection, json, apiEndpoint)
	res, err := client.Do(req)
	return res, err
}

//
// Delete a collection with the provided apiEndpoints
//
func Delete(collection string, json string, apiEndpoint string) (*http.Response, error) {
	req := build_request("DELETE", collection, json, apiEndpoint)
	res, err := client.Do(req)
	return res, err
}

//
// PUT request to Cloudant
//
func Put(collection string, json string, apiEndpoint string) (*http.Response, error) {
	req := build_request("PUT", collection, json, apiEndpoint)
	res, err := client.Do(req)
	return res, err
}

//
// form an http request to issue a REST call to Cloudant
//
func build_request(method string, collection string, jsonBody string, apiEndpoint string) *http.Request {
	url := "https://" + host + "/" + collection
	if apiEndpoint != "" {
		url += "/" + apiEndpoint
	}
	var jsonStr = []byte(jsonBody)
	req, err1 := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	if err1 != nil {
		fmt.Printf("error, cadb.build_request: %s\n", err1)
		return nil
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(username, password)
	return req
}
