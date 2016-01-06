package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const CONFIG_DIR = "config/squads"

//
// This structure holds the number of commits to a repository on particular day
//
type Squad struct {
	Name  string   `json:"name"`  // Name of the squad
	Image string   `json:"image"` // image URL to display for squad
	Repos []string `json:"repos"` // an array of repo URLs
}

//
// Construct a an array of Squads from the json files
// found in config/squad/*
//
func InitSquadsFromJson() []Squad {
	var a []Squad
	names, err := ioutil.ReadDir(CONFIG_DIR)
	if err != nil {
		fmt.Printf("error reading directory structure: %s\n", err)
	}
	for i := range names {
		js, _ := ioutil.ReadFile(CONFIG_DIR + "/" + names[i].Name())

		var squad Squad
		json.Unmarshal([]byte(js), &squad)
		a = append(a, squad)
	}
	return a
}
