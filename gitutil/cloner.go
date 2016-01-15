package gitutil

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

//
// This is the entry point for parsing a given repository URL,
// then generating returning the repository's git commit history (which can be heat-mapped)
// The function does the following:
// - parse the URL into a directory path
// - clone the repository into that directory, if it does not exist
// - convert the git logs into JSON and return that to send back to the browser
//
func CloneRepo(repoUrl string) (string, error) {
	//fmt.Println("parsing commit history")

	isCloned := CheckExists(repoUrl)

	if isCloned == false {
		err := clone_repo(repoUrl)
		if err != nil {
			return "error", err
		}
	}

	res, err := GetLocalCommits(repoUrl)
	if err != nil {
		return "error", err
	}

	return res, nil
}

//
// if reading a directory returns an error,
// we will assume it doesn't exist
//
func CheckExists(repoUrl string) bool {
	var dir = ".clones/" + UrlToDir(repoUrl)
	_, err := ioutil.ReadDir(dir)
	if err == nil {
		return true
	} else {
		return false
	}
}

//
// Clone the git repository into the temporary .clones/ directory
//
func clone_repo(repoUrl string) error {
	var arg0 = "git"
	var arg1 = "clone"
	var arg2 = ""
	if strings.Contains(repoUrl, GITHUBIBM_SSH_IDENTIFIER) {
		fmt.Println("GITHUB SSH!")
		arg2 = repoUrl
	} else {
		arg2 = "http://" + repoUrl
	}
	var arg3 = ".clones/" + UrlToDir(repoUrl)

	cmd := exec.Command(arg0, arg1, arg2, arg3)
	err := cmd.Start()
	//stdin, _ := cmd.StdinPipe()

	if err != nil {
		fmt.Print("an error occured attempting to execute 'git clone' ")
		fmt.Print("with parameters: " + arg2 + ", " + arg3 + ": ")
		fmt.Printf("error is: %s\n", err)
		return err
	}

	//fmt.Printf("waiting for clone to finish...")
	err = cmd.Wait()
	//fmt.Printf("done\n")

	if err != nil {
		fmt.Printf("an error occured during the execution of 'git clone': %s\n", err)
		return err
	}

	//fmt.Println("git clone completed successfully")
	return nil
}
