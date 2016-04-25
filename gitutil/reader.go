package gitutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
)

//
// return a list of all repositories in the .clones directory
// all '.' characters are replaced with '/' characters
//
func AllTrackedRepos() []string {
	dirs, err := ioutil.ReadDir(CLONES_DIR)
	if err != nil {
		fmt.Printf("error reading directory structure of .clones %s\n", err)
	}
	repos := make([]string, len(dirs))
	for i := 0; i < len(dirs); i++ {
		dir := dirs[i].Name()
		repos[i] = DirToUrl(dir)
	}
	return repos
}

//
// return a single randomly selected repository
// from all repositories tracked in .clones/
// the repo is returned in original URL format
//
func GetRandomRepo() string {
	dirs, err := ioutil.ReadDir(CLONES_DIR)
	if err != nil {
		fmt.Printf("error reading directory structure of .clones %s\n", err)
	}
	r := rand.Intn(len(dirs))
	return DirToUrl(dirs[r].Name())
}

//
// Get a list of commits-per-day,
// from local clone of repo
//
func GetLocalCommits(url string) (string, error) {
	js, err := commits_to_json(url)
	if err != nil {
		log.Printf("Error, gitutil.UpdateRefs: %s\n", err)
		return "error", err
	}
	allCommits := json_to_repoCommits(js, url)
	b, err := json.Marshal(allCommits)
	//fmt.Printf("%s\n", string(b))
	return string(b), nil
}
