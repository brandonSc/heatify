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
		log.Printf("error, gitutil.UpdateRefs: error fetching refs for %s. error is %s\n", path, err)
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
	allCommits := json_to_gostruct(js, url)

	// calculate the latest commits
	dbCommits := model.DbRetrieveAllRepoCommits(url)
	newCommits := filter_changeset(allCommits, dbCommits)

	fmt.Println("LEN of changeset: %d\n", len(newCommits))

	// send to cloudant database
	model.DbSendRepoCommitsArray(newCommits)
}

//
// Convert a JSON string to an array of RepoCommits structs
//
func json_to_gostruct(js string, url string) []model.RepoCommits {
	// a list of RepoCommits - indication commits-per-day
	var rcList []model.RepoCommits
	// a map used to build the list of RepoCommits
	rcMap := make(map[int64]model.RepoCommits)

	var f []interface{}
	json.Unmarshal([]byte(js), &f)

	for i := range f {
		m := f[i].(map[string]interface{})
		// parse the date from the JSON
		secs, err := strconv.ParseInt(m["date"].(string), 10, 64)
		if err != nil {
			fmt.Println("Error, gitutil.json_to_gostruct: %s\n")
		}
		cDate := time.Unix(secs, 0)
		// remove the time on the date so that only DD/MM/YYYY info remain
		nDate := time.Date(cDate.Year(), cDate.Month(), cDate.Day(), 0, 0, 0, 0, time.UTC)
		// concatenate the dates ..
		if rcMap[nDate.Unix()].URL == "" {
			rcMap[nDate.Unix()] = model.RepoCommits{
				url,
				nDate,
				"",
				1,
			}
		} else {
			rc := rcMap[nDate.Unix()]
			rc.Commits += 1
			rcMap[nDate.Unix()] = rc
		}
	}

	for _, j := range rcMap {
		//fmt.Printf("\tOn %s -> %d commits made\n", j.Date, j.Commits)
		rcList = append(rcList, j)
	}

	return rcList
}

//
// compare local RepoCommits to the ones stored on cloudant,
// to determine which RepoCommits are new
// -- runs in O(n^2) where dbCommits ~= allCommits = n
//
func filter_changeset(localCommits []model.RepoCommits, dbCommits []model.RepoCommits) []model.RepoCommits {
	var rc []model.RepoCommits

	for i := range localCommits {
		found := false
		for j := range dbCommits {
			if dbCommits[j].Date == localCommits[i].Date {
				found = true
				// found a set of commits in cloudant that matches the local set
				// now need to determine if the cloudant set needs to be updated
				if dbCommits[j].Commits < localCommits[i].Commits {
					rc = append(rc, localCommits[i])
				}
			}
		}
		if !found {
			// didn't find a matching commit set in cloudant
			// need to add it to the changeset
			rc = append(rc, localCommits[i])
		}
	}

	return rc
}

//
// returns true if the dates without the time (hours, minutes, etc) are equal
//
func equal_dates(date1 time.Time, date2 time.Time) bool {
	return date1.Year() == date2.Year() &&
		date1.Month() == date2.Month() &&
		date1.Day() == date2.Day()
}

//
// returns true if the `date1` without the time
// (hours, minutes, etc) is less than `date2`
//
func before_date(date1 time.Time, date2 time.Time) bool {
	return date1.Year() < date2.Year() ||
		date1.Month() < date2.Month() ||
		date1.Day() < date2.Day()
}
