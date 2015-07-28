package gitutil

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"
)

//
// runs continually forever (should run this in a new goroutine)
// asynchronously dispatches update requests to each tracked repository
//
func RunUpdateLoop() {
	for {
		dirs, err := ioutil.ReadDir(".clones")
		if err != nil {
			fmt.Printf("error reading directory structure of .clones %s\n", err)
		}
		for i := 0; i < len(dirs); i++ {
			dir := dirs[i].Name()
			go UpdateRefs(".clones/" + dir)
		}
		time.Sleep(5 * time.Minute)
	}
}

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
}
