package gitutil

import (
	"fmt"
	//"hub.jazz.net/git/schurman93/Git-Monitor/model"
	"io/ioutil"
	"os/exec"
	"strings"
)

//
// Public - parse the results of git clone and git logs into structs
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

	res, err := crunch_stats(repoUrl)
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
// Clone the git repository into the temporary .clones/ directory
//
func clone_repo(repoUrl string) error {
	var arg0 = "git"
	var arg1 = "clone"
	//var arg2 = "--bare"
	var arg2 = "http://" + repoUrl
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

//
// Delete the git repo from .clones/
//
func delete_repo(repoUrl string) error {
	var arg0 = "rm"
	var arg1 = "-rf"
	var arg2 = ".clones/" + UrlToDir(repoUrl)

	cmd := exec.Command(arg0, arg1, arg2)
	err := cmd.Run()

	if err != nil {
		fmt.Print("an error occured during the deletion 'rm -rf' of: " + arg2 + ". ")
		fmt.Printf("error is: %s\n", err)
		return err
	}

	//fmt.Println("local git repo deleted successfully")
	return nil
}
