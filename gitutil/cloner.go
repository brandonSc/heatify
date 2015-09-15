package gitutil

import (
	"fmt"
	"os/exec"
)

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
