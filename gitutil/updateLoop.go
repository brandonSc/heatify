package gitutil

import (
	"fmt"
	"io/ioutil"
	"time"
)

const (
	INTERVAL   = 360 // update interval in minutes
	CLONES_DIR = ".clones"
)

//
// runs continually forever (should run this in a new goroutine)
// asynchronously dispatches update requests to each tracked repository
//
func RunUpdateLoop() {
	for {
		dirs, err := ioutil.ReadDir(CLONES_DIR)
		if err != nil {
			fmt.Printf("error reading directory structure of .clones %s\n", err)
		}
		for i := 0; i < len(dirs); i++ {
			dir := dirs[i].Name()
			go UpdateRefs(dir)
		}
		time.Sleep(INTERVAL * time.Minute)
	}
}
