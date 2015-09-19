package gitutil

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	"hub.jazz.net/git/schurman93/Git-Monitor/model"
)

//
// run `git fetch` on a given repository specified by the path parameter
//
func UpdateRefs(path string) {
	// run a `git pull`
	cmd := exec.Command("/bin/bash", "-c", "cd "+CLONES_DIR+"/"+path+" && git pull origin master")
	err := cmd.Run()
	if err != nil {
		log.Printf("error fetching refs for %s. error is %s\n", path, err)
		return
	}
	// get the git commits in JSON
	url := DirToUrl(path)
	js, err := commits_to_json(url)
	if err != nil {
		log.Printf("Error, gitutil.UpdateRefs: %s\n", err)
		return
	}
	// parse the JSON into go structs
	rc := json_to_gostruct(js)
	fmt.Printf("LEN: %d\n", len(rc))

	// send to cloudant database
}

func json_to_gostruct(js string) []model.RepoCommits {
	var rc []model.RepoCommits

	var f []interface{}
	json.Unmarshal([]byte(js), &f)

	for i := range f {
		fmt.Printf("f: %s\n", f[i])
	}

	return rc
}
