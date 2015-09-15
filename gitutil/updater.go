package gitutil

import (
	"fmt"
	"os/exec"
)

//
// run `git fetch` on a given repository specified by the path parameter
//
func UpdateRefs(path string) {
	//fmt.Println("updating refs for " + path + "...")
	cmd := exec.Command("/bin/bash", "-c", "cd "+path+" && git pull origin master")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("error fetching refs for %s. error is %s\n", path, err)
	}
	//fmt.Println("done updating refs for " + path)
	//crunch_stats
}
