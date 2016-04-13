package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

//
// a User is defined by their username,
// and a collection of git account aliases they go by
// (e.g. git alias: 'Brandon Schurman <schurman@ca.ibm.com>')
// optionally a profile pic can be added.
//
type User struct {
	Username string   `json:"username"`        // username (e.g. Slack name)
	Image    string   `json:"image,omitempty"` // URL for profile pic
	Aliases  []string `json:"aliases"`         // an array of Git account aliases
}

const USERS_DIR = "config/users"

func IsUserAlias(user string) bool {
	return strings.Contains(user, "<") && strings.Contains(user, ">")
}

//
// Construct a an array of Users from the json files
// found in config/users/*
//
func InitUsersFromJson() []User {
	var a []User
	files, err := ioutil.ReadDir(USERS_DIR)
	if err != nil {
		fmt.Printf("error reading users directory structure: %s\n", err)
		return nil
	}
	for i := range files {
		user, _ := InitUserFromJson(files[i].Name())
		a = append(a, user)
	}
	return a
}

//
// initialize a single User
// using the JSON config file name
//
func InitUserFromJson(name string) (User, error) {
	js, err := ioutil.ReadFile(USERS_DIR + "/" + name)
	if err != nil {
		return User{}, err
	}
	var user User
	json.Unmarshal([]byte(js), &user)
	return user, nil
}

//
// Get all the UserCommits from Cloudant
// for the Squad with the @param name
//
func GetAllUserCommits(name string) ([]UserCommits, error) {
	user, err := InitUserFromJson(name + ".json")
	if err != nil {
		return nil, err
	}
	commits, err := FindUserCommits(user.Username)
	return commits, nil
}
