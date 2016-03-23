package gitutil

import (
	"fmt"
	"os/exec"
)

//
// run the 'git logs' command and parse the output into JSON
//
func commits_to_json(repoUrl string) (string, error) {
	var dir = ".clones/" + UrlToDir(repoUrl)
	var script = `git log \
	--pretty=format:'{%n  "commit": "%H",%n  "author": "%an <%ae>",%n  "date": "%at",%n  "message": "%f"%n},' \
	$@ | \
	perl -pe 'BEGIN{print "["}; END{print "]\n"}' | \
	perl -pe 's/},]/}]/' && exit`

	cmd := exec.Command("/bin/bash", "-c", "cd "+dir+" && "+script)
	out, err := cmd.Output()

	if err != nil {
		fmt.Print("an error occured running 'git log' on repo/dir '" + dir + "'. ")
		fmt.Printf("error is: %s\n", err)
		return "error in crunch_stats", err
	}

	return string(out[:]), nil
}
