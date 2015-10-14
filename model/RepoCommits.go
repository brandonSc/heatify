package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hub.jazz.net/git/schurman93/Git-Monitor/cadb"
	//	"reflect"
	"time"
)

//
// This structure holds the number of commits to a repository on particular day
//
type RepoCommits struct {
	URL     string    `json:"url"`     // URL of the repository
	Date    time.Time `json:"date"`    // date of the commits on the repo: YYYY/MM/DD (TODO confirm this)
	LastId  string    `json:"lastId"`  // most recent commit ID (hex) for distinguishing new commits on current date
	Commits int       `json:"commits"` // number of commits on this day
}

//
// Create a record for this entity in the cloudant database
//
func (rc RepoCommits) DbCreate() {
	json := fmt.Sprintf(`{
		"URL": "%s",
		"Date": "%s",
		"LastId": "%s",
		"Commits": %d
	}`, rc.URL, rc.Date, rc.LastId, rc.Commits)
	_, err := cadb.Post("gitmonitor-repos", json, "")
	if err != nil {
		fmt.Printf("error, model.RepoCommits.DbCreate: %s\n", err)
	}
}

//
// Get a list of all db records created under the URL
// this is the function that is called to get the data for the Repository HeatMap
//
func DbRetrieveAllRepoCommits(url string) []RepoCommits {
	fmt.Println("(TODO) retrieving for URL: " + url)
	js := `{
		"selector": {
			"_id": {
				"$gt": 0
			}, 
			"URL": "` + url + `" 
		},
		"sort": [
			{
				"_id": "asc"
			}
		]
	}`
	res, err := cadb.Post("gitmonitor-repos", js, "_find")
	if err != nil {
		fmt.Printf("error, model.RepoCommits.DbRetrieveAll: %s\n", err)
		return nil
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	return json_to_array(buf.String())
}

//
// private function for turning json response into an array in go
//
func json_to_array(js string) []RepoCommits {
	var f interface{}
	json.Unmarshal([]byte(js), &f)
	m := f.(map[string]interface{})
	docs := m["docs"].([]interface{})
	var a []RepoCommits
	for i := range docs {
		//c := &RepoCommits{}
		c := docs[i].(map[string]interface{})
		a = append(a, RepoCommits{
			c["URL"].(string),
			time.Now(),
			//c["date"].(time.Time),
			c["LastId"].(string),
			c["Commits"].(int),
		})
	}
	return a
}
