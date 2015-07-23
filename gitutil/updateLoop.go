package gitutil

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"
)

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

func UpdateRefs(path string) {
	//fmt.Println("updating refs for " + path + "...")
	cmd := exec.Command("/bin/bash", "-c", "cd "+path+" && git fetch --prune")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("error fetching refs for %s. error is %s\n", path, err)
	}
	//fmt.Println("done updating refs for " + path)
}
