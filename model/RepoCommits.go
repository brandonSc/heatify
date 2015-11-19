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
	Commits int       `json:"commits"` // number of commits on this day
	//LastId  string    `json:"lastId"`  // most recent commit ID (hex) for distinguishing new commits on current date
}

//
// Create a record for this entity in the cloudant database
//
func (rc RepoCommits) DbCreate() {
	json := fmt.Sprintf(`{
		"URL": "%s",
		"Date": "%s",
		"Commits": %d
	}`, rc.URL, rc.Date, rc.Commits)
	_, err := cadb.Post("gitmonitor-repos", json, "")
	if err != nil {
		fmt.Printf("error, model.RepoCommits.DbCreate: %s\n", err)
	}
}

//
// send an array of `RepoCommits` in JSON to cloudant
//
func DbSendRepoCommitsArray(rcs []RepoCommits) {
	js, err := json.Marshal(rcs)
	if err != nil {
		fmt.Printf("error, model.RepoCommits.DbSendRepoCommitsArray: %s\n", err)
		return
	}
	res, err := cadb.Post("gitmonitor-repos", `{"docs":`+string(js)+`}`, "_bulk_docs")
	if err != nil {
		fmt.Printf("error, model.RepoCommits.DbSendRepoCommitsArray: %s\n", err)
	} else {
		fmt.Printf("model.RepoCommits.DbSendRepoCommitsArray: %s\n", res)
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
			"url": "` + url + `" 
		}
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
		day, _ := time.Parse(time.RFC3339, c["date"].(string))
		a = append(a, RepoCommits{
			c["url"].(string),
			day,
			int(c["commits"].(float64)),
		})
	}
	return a
}
