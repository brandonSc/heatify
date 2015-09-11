package gitutil

import (
	"fmt"
	"os/exec"
	"strings"
)

//
// run the 'git logs' command and parse the output into JSON
//
func crunch_stats(repoUrl string) (string, error) {
	var dir = ".clones/" + strings.Replace(repoUrl, "/", ".", -1)
	var script = `git log \
	--pretty=format:'{%n  "commit": "%H",%n  "author": "%an <%ae>",%n  "date": "%ad",%n  "message": "%f"%n},' \
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
