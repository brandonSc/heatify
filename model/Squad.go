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
	Name  string `json:"name"`  // URL of the repository
	Repos string `json:"repos"` // number of commits on this day
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
		js, _ := ioutil.ReadFile(names[i].Name())

		var squad Squad
		json.Unmarshal([]byte(js), &squad)
		a = append(a, squad)
	}
	return a
}
