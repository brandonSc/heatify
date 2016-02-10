package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"hub.jazz.net/git/schurman93/Git-Monitor/cadb"
)

// set this to `gitmonitor-users-dev` for development,
// or `gitmonitor-users` for production only on the RTP VM
const USERS_DB = "gitmonitor-users-dev"

//
// Stores the number of commits made to a repo each day by a particular user
//
type UserCommits struct {
	User    string    `json:"user"`           // username or email of the user
	URL     string    `json:"url"`            // URL of the repository
	Date    time.Time `json:"date"`           // date of the commits on the repo: YYYY/MM/DD
	Commits int       `json:"commits"`        // number of commits on this day
	Id      string    `json:"_id,omitempty"`  // record ID in cloudant
	Rev     string    `json:"_rev,omitempty"` // revision ID in cloudant
}

//
// send an array of `UserCommits` in JSON to cloudant
//
func DbSendUserCommitsArray(ucs []UserCommits) {
	js, err := json.Marshal(ucs)
	if err != nil {
		fmt.Printf("error, model.UserCommits.DbSendUserCommitsArray: %s\n", err)
		return
	}
	res, err := cadb.Post(USERS_DB, `{"docs":`+string(js)+`}`, "_bulk_docs")
	if err != nil {
		fmt.Printf("error, model.UserCommits.DbSendUserCommitsArray: %s. Response is: %s\n", err, res)
	}
	//fmt.Println("cloudant: ", res)
}

//
// Get a list of all db records created under the URL
// this is the function that is called to get the data for the Repository HeatMap
//
func DbRetrieveAllUserCommits(url string) []UserCommits {
	js := `{
		"selector": {
			"_id": {
				"$gt": 0
			}, 
			"url": "` + url + `" 
		}
	}`
	res, err := cadb.Post(USERS_DB, js, "_find")
	if err != nil {
		fmt.Printf("error, model.UserCommits.DbRetrieveAll: %s\n", err)
		return nil
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	return json_to_userCommits_array(buf.String())
}

//
// private function for turning json response into an array in go
//
func json_to_userCommits_array(js string) []UserCommits {
	var f interface{}
	json.Unmarshal([]byte(js), &f)
	m := f.(map[string]interface{})
	docs := m["docs"].([]interface{})
	var a []UserCommits
	for i := range docs {
		//c := &UserCommits{}
		c := docs[i].(map[string]interface{})
		day, _ := time.Parse(time.RFC3339, c["date"].(string))
		a = append(a, UserCommits{
			c["user"].(string),
			c["url"].(string),
			day,
			int(c["commits"].(float64)),
			c["_id"].(string),
			c["_rev"].(string),
		})
	}
	return a
}
