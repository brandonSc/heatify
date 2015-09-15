package gitutil

import (
	"fmt"
	"os/exec"
)

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
