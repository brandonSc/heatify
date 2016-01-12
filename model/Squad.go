package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

const CONFIG_DIR = "config/squads"

//
// This structure holds the number of commits to a repository on particular day
//
type Squad struct {
	Name    string   `json:"name"`    // Name of the squad
	Image   string   `json:"image"`   // image URL to display for squad
	Repos   []string `json:"repos"`   // an array of repo URLs
	Members []string `json:"members"` // an array of member's Slack usernames
}

//
// Construct a an array of Squads from the json files
// found in config/squad/*
//
func InitSquadsFromJson() []Squad {
	var a []Squad
	files, err := ioutil.ReadDir(CONFIG_DIR)
	if err != nil {
		fmt.Printf("error reading directory structure: %s\n", err)
	}
	for i := range files {
		squad := InitSquadFromJson(files[i].Name())
		a = append(a, squad)
	}
	return a
}

//
// initialize a single squad
// using the JSON config file name
//
func InitSquadFromJson(name string) Squad {
	js, _ := ioutil.ReadFile(CONFIG_DIR + "/" + name)
	var squad Squad
	json.Unmarshal([]byte(js), &squad)
	return squad
}

//
// Get all the RepoCommtis from Cloudant
// for the Squad with the @param name
//
func GetSquadRepoCommits(name string) []RepoCommits {
	var a []RepoCommits
	file := strings.ToLower(name) + ".json"
	squad := InitSquadFromJson(file)
	for i := range squad.Repos {
		commits := DbRetrieveAllRepoCommits(squad.Repos[i])
		for j := range commits {
			a = append(a, commits[j])
		}
	}
	return a
}
