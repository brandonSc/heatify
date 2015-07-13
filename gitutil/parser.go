package gitutil

import (
	"fmt"
	"hub.jazz.net/git/schurman93/Git-Monitor/model"
	"os/exec"
	"strings"
)

//
// Public - parse the results of git clone and git logs into structs
//
func ParseCommits(repoUrl string) []model.Commit {
	fmt.Println("parsing commit history")

	clone_repo(repoUrl)
	crunch_stats(repoUrl)
	delete_repo(repoUrl)

	return nil
}

//
// run the 'git logs' command and parse the output into JSON
//
func crunch_stats(repoUrl string) []model.Commit {
	var dir = ".clones/" + strings.Replace(repoUrl, "/", ".", -1)
	var script = `git log \
	--pretty=format:'{%n  "commit": "%H",%n  "author": "%an <%ae>",%n  "date": "%ad",%n  "message": "%f"%n},' \
	$@ | \
	perl -pe 'BEGIN{print "["}; END{print "]\n"}' | \
	perl -pe 's/},]/}]/'`

	cmd := exec.Command("/bin/bash", "-c", "cd "+dir+" && "+script)
	fmt.Print("crunching the numbers...")
	out, err := cmd.Output()
	fmt.Println("done")

	if err != nil {
		fmt.Print("an error occured running 'git log' on repo/dir '" + dir + "'. ")
		fmt.Printf("error is: %s\n", err)
		return nil
	}

	fmt.Printf("result from log is:\n%s\n", out)
	return nil
}

//
// Clone the git repository into the temporary .clones/ directory
//
func clone_repo(repoUrl string) {
	var arg0 = "git"
	var arg1 = "clone"
	var arg2 = "http://" + repoUrl
	var arg3 = ".clones/" + strings.Replace(repoUrl, "/", ".", -1)

	cmd := exec.Command(arg0, arg1, arg2, arg3)
	err := cmd.Start()

	if err != nil {
		fmt.Print("an error occured attempting to execute 'git clone' ")
		fmt.Print("with parameters: " + arg2 + ", " + arg3 + ": ")
		fmt.Printf("error is: %s\n", err)
		return
	}

	fmt.Printf("waiting for clone to finish...")
	err = cmd.Wait()
	fmt.Printf("done\n")

	if err != nil {
		fmt.Printf("an error occured during the execution of 'git clone': %s\n", err)
		return
	}

	fmt.Println("git clone completed successfully")
}

//
// Delete the git repo from .clones/
//
func delete_repo(repoUrl string) {
	var arg0 = "rm"
	var arg1 = "-rf"
	var arg2 = ".clones/" + strings.Replace(repoUrl, "/", ".", -1)

	cmd := exec.Command(arg0, arg1, arg2)
	err := cmd.Run()

	if err != nil {
		fmt.Print("an error occured during the deletion 'rm -rf' of: " + arg2 + ". ")
		fmt.Printf("error is: %s\n", err)
	}

	fmt.Println("local git repo deleted successfully")
}
