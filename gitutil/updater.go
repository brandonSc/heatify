package gitutil

import (
	"encoding/json"
	"fmt"
	"hub.jazz.net/git/schurman93/Git-Monitor/model"
	"log"
	"os/exec"
	"strconv"
	"time"
)

//
// run `git fetch` on a given repository specified by the path parameter
//
func UpdateRefs(path string) {
	// run a `git pull` to get the latest remote data
	cmd := exec.Command("/bin/bash", "-c", "cd "+CLONES_DIR+"/"+path+" && git pull origin master")
	err := cmd.Run()
	if err != nil {
		log.Printf("error fetching refs for %s. error is %s\n", path, err)
		return
	}
	// get the local git commits in JSON
	url := DirToUrl(path)
	js, err := commits_to_json(url)
	if err != nil {
		log.Printf("Error, gitutil.UpdateRefs: %s\n", err)
		return
	}
	// parse the JSON into go structs
	allCommits := json_to_gostruct(js)
	fmt.Printf("LEN: %d\n", len(allCommits))

	// calculate the latest commits
	//dbCommits := model.DbRetrieveAllRepoCommits(url)
	//newCommits := filter_changeset(allCommits, dbCommits)

	//fmt.Printf("%s\n", newCommits)

	// send to cloudant database

}

func json_to_gostruct(js string) []model.RepoCommits {
	var rc []model.RepoCommits

	//datePos := time.Now()
	var f []interface{}
	json.Unmarshal([]byte(js), &f)

	for i := range f {
		m := f[i].(map[string]interface{})
		secs, err := strconv.ParseInt(m["date"].(string), 10, 64)
		if err != nil {
			fmt.Println("Error, gitutil.json_to_gostruct: %s\n")
		}
		cDate := time.Unix(secs, 0)
		fmt.Printf("date: %s\n", cDate)
	}

	return rc
}

func filter_changeset(allCommits []model.RepoCommits, dbCommits []model.RepoCommits) []model.RepoCommits {
	var rc []model.RepoCommits

	// TODO

	return rc
}
