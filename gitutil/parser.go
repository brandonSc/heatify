package gitutil

import (
	"fmt"
	//"hub.jazz.net/git/schurman93/Git-Monitor/model"
	"io/ioutil"
	"os/exec"
)

//
// This is the entry point for parsing a given repository URL,
// then generating returning the repository's git commit history (which can be heat-mapped)
// The function does the following:
// - parse the URL into a directory path
// - clone the repository into that directory, if it does not exist
// - convert the git logs into JSON and return that to send back to the browser
//
func ParseCommits(repoUrl string) (string, error) {
	//fmt.Println("parsing commit history")

	isCloned := check_exists(repoUrl)

	if isCloned == false {
		err := clone_repo(repoUrl)
		if err != nil {
			return "error", err
		}
	}

	res, err := commits_to_json(repoUrl)
	if err != nil {
		return "error", err
	}
	/*
		err = delete_repo(repoUrl)
		if err != nil {
			return "error", err
		}
	*/
	return res, nil
}

//
// if reading a directory returns an error,
// we will assume it doesn't exist
//
func check_exists(repoUrl string) bool {
	var dir = ".clones/" + UrlToDir(repoUrl)
	_, err := ioutil.ReadDir(dir)
	if err == nil {
		return true
	} else {
		return false
	}
}

//
// run the 'git logs' command and parse the output into JSON
//
func commits_to_json(repoUrl string) (string, error) {
	var dir = ".clones/" + UrlToDir(repoUrl)
	var script = `git log \
	--pretty=format:'{%n  "commit": "%H",%n  "author": "%an <%ae>",%n  "date": "%at",%n  "message": "%f"%n},' \
	$@ | \
	perl -pe 'BEGIN{print "["}; END{print "]\n"}' | \
	perl -pe 's/},]/}]/'`

	cmd := exec.Command("/bin/bash", "-c", "cd "+dir+" && "+script)
	//fmt.Print("crunching the numbers...")
	out, err := cmd.Output()
	//fmt.Println("done")

	if err != nil {
		fmt.Print("an error occured running 'git log' on repo/dir '" + dir + "'. ")
		fmt.Printf("error is: %s\n", err)
		return "error in crunch_stats", err
	}

	return string(out[:]), nil
}
