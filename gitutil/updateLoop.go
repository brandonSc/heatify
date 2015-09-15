package gitutil

import (
	"fmt"
	"io/ioutil"
	"time"
)

const INTERVAL = 10 // update interval in minutes

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
		time.Sleep(INTERVAL * time.Minute)
	}
}
