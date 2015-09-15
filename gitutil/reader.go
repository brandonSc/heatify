package gitutil

import (
	"fmt"
	"io/ioutil"
	"math/rand"
)

//
// return a list of all repositories in the .clones directory
// all '.' characters are replaced with '/' characters
//
func AllTrackedRepos() []string {
	dirs, err := ioutil.ReadDir(".clones")
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
	dirs, err := ioutil.ReadDir(".clones")
	if err != nil {
		fmt.Printf("error reading directory structure of .clones %s\n", err)
	}
	r := rand.Intn(len(dirs))
	return DirToUrl(dirs[r].Name())
}
