package model

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"hub.jazz.net/git/schurman93/Git-Monitor/cadb"
)

// set this to `gitmonitor-users-dev` for development,
// or `gitmonitor-users` for production only on the RTP VM
const USERS_DB = "gitmonitor-users-dev"

//
// Stores the number of commits made to a repo each day by a particular user
//
type UserCommits struct {
	User    string    `json:"user"`           // username or email of the user
	URL     string    `json:"url"`            // URL of the repository
	Date    time.Time `json:"date"`           // date of the commits on the repo: YYYY/MM/DD
	Commits int       `json:"commits"`        // number of commits on this day
	Id      string    `json:"_id,omitempty"`  // record ID in cloudant
	Rev     string    `json:"_rev,omitempty"` // revision ID in cloudant
}

//
// send an array of `UserCommits` in JSON to cloudant
//
func DbSendUserCommitsArray(ucs []UserCommits) {
	js, err := json.Marshal(ucs)
	if err != nil {
		fmt.Printf("error, model.UserCommits.DbSendUserCommitsArray: %s\n", err)
		return
	}
	res, err := cadb.Post(USERS_DB, `{"docs":`+string(js)+`}`, "_bulk_docs")
	if err != nil {
		fmt.Printf("error, model.UserCommits.DbSendUserCommitsArray: %s. Response is: %s\n", err, res)
	}
	//fmt.Println("cloudant: ", res)
}

//
// Get a list of all db records created under the URL
// this is the function that is called to get the data for the Repository HeatMap
//
func DbRetrieveAllUserCommits(url string) ([]UserCommits, error) {
	js := `{
		"selector": {
			"_id": {
				"$gt": 0
			}, 
			"url": "` + url + `" 
		}
	}`
	res, err := cadb.Post(USERS_DB, js, "_find")
	if err != nil {
		fmt.Printf("error, model.UserCommits.DbRetrieveAll: %s\n", err)
		return nil, err
	}
	buf := new(bytes.Buffer)
	fmt.Printf("RESPONSE BUF: %s", buf)
	if buf == nil {
		return nil, errors.New("model.DbRetrieveAllUserCommits: Response is null")
	}
	buf.ReadFrom(res.Body)
	return json_to_userCommits_array(buf.String())
}

//
// all user commits from all repos
// @param user git author
//
func FindUserCommits(user string) ([]UserCommits, error) {
	js := `{
		"selector": {
			"_id": {
				"$gt": 0
			}, 
			"user": "` + user + `" 
		}
	}`
	res, err := cadb.Post(USERS_DB, js, "_find")
	if err != nil {
		fmt.Printf("error, model.UserCommits.FindUserCommits: %s\n", err)
		return nil, err
	}
	buf := new(bytes.Buffer)
	fmt.Printf("RESPONSE BUF: %s", buf)
	if buf == nil {
		return nil, errors.New("model.FindUserCommits: Response is null")
	}
	buf.ReadFrom(res.Body)
	return json_to_userCommits_array(buf.String())
}

//
// all user commits from all repos, from all user aliaes
// @param users an array of git aliases (e.g. from a heatify user profile)
//
func FindMultiUserCommits(users []string) ([]UserCommits, error) {
	byt1, _ := json.Marshal(users)
	usersjs := string(byt1)
	js := `{
		"selector": {
			"_id": {
				"$gt": 0
			}, 
			"user": {
				"$in":` + usersjs + `
			}
		}
	}`
	res, err := cadb.Post(USERS_DB, js, "_find")
	if err != nil {
		fmt.Printf("error, model.UserCommits.FindUserCommits: %s\n", err)
		return nil, err
	}
	buf := new(bytes.Buffer)
	fmt.Printf("RESPONSE BUF: %s", buf)
	if buf == nil {
		return nil, errors.New("model.FindUserCommits: Response is null")
	}
	buf.ReadFrom(res.Body)
	return json_to_userCommits_array(buf.String())
}

//
// all user commits from all repos, from all user aliaes
// @param users an array of git aliases (e.g. from a heatify user profile)
//
func FindMultiUserCommitsOnMultiRepo(users []string, repos []string) ([]UserCommits, error) {
	byt1, _ := json.Marshal(users)
	usersjs := string(byt1)
	byt2, _ := json.Marshal(repos)
	urlsjs := string(byt2)
	js := `{
		"selector": {
			"_id": {
				"$gt": 0
			}, 
			"user": {
				"$in":` + usersjs + `
			},
			"url": {
				"$in": ` + urlsjs + `
			}
		}
	}`
	res, err := cadb.Post(USERS_DB, js, "_find")
	if err != nil {
		fmt.Printf("error, model.UserCommits.FindUserCommits: %s\n", err)
		return nil, err
	}
	buf := new(bytes.Buffer)
	fmt.Printf("RESPONSE BUF: %s", buf)
	if buf == nil {
		return nil, errors.New("model.FindUserCommits: Response is null")
	}
	buf.ReadFrom(res.Body)
	return json_to_userCommits_array(buf.String())
}

//
// all user comiits on a specific repo
//
func FindUserCommitsOnRepo(user string, repo string) ([]UserCommits, error) {
	js := `{
		"selector": {
			"_id": {
				"$gt": 0
			}, 
			"user": "` + user + `",
			"url": "` + repo + `"
		}
	}`
	res, err := cadb.Post(USERS_DB, js, "_find")
	if err != nil {
		fmt.Printf("error, model.UserCommits.FindUserCommitsOnRepo: %s\n", err)
		return nil, err
	}
	buf := new(bytes.Buffer)
	fmt.Printf("RESPONSE BUF: %s", buf)
	if buf == nil {
		return nil, errors.New("model.FindUserCommitsOnRepo: Response is null")
	}
	buf.ReadFrom(res.Body)
	return json_to_userCommits_array(buf.String())
}

//
// all user commits on all provided repo (e.g. squad)
//
func FindUserCommitsOnMultiRepo(user string, repos []string) ([]UserCommits, error) {
	byt, _ := json.Marshal(repos)
	urlsjs := string(byt)
	js := `{
		"selector": {
			"_id": {
				"$gt": 0
			}, 
			"user": "` + user + `",
			"url": {
				"$in": ` + urlsjs + `
			}
		}
	}`
	fmt.Println(js)
	res, err := cadb.Post(USERS_DB, js, "_find")
	if err != nil {
		fmt.Printf("error, model.FindUserCommitsOnMultiRepo: %s\n", err)
		return nil, err
	}
	buf := new(bytes.Buffer)
	fmt.Printf("RESPONSE BUF: %s", buf)
	if buf == nil {
		return nil, errors.New("model.FindUserCommitsOnMultiRepo: Response is null")
	}
	buf.ReadFrom(res.Body)
	return json_to_userCommits_array(buf.String())
}

//
// all commits on all provided repos (e.g. squad), excluding provided users
// basically, this is for the community heatmap displaed on the squad page
//
func FindCommunityCommitsOnMultiRepo(users []string, repos []string) ([]UserCommits, error) {
	byt1, _ := json.Marshal(users)
	byt1 = bytes.Replace(byt1, []byte("\\u003c"), []byte("<"), -1)
	byt1 = bytes.Replace(byt1, []byte("\\u003e"), []byte(">"), -1)
	usersjs := string(byt1)
	byt2, _ := json.Marshal(repos)
	urlsjs := string(byt2)
	js := `{
		"selector": {
			"_id": {
				"$gt": 0
			}, 
			"user": {
				"$nin":` + usersjs + `
			},
			"url": {
				"$in": ` + urlsjs + `
			}
		}
	}`
	res, err := cadb.Post(USERS_DB, js, "_find")
	if err != nil {
		fmt.Printf("error, model.UserCommits.FindCommunityCommitsOnMultiRepo: %s\n", err)
		return nil, err
	}
	buf := new(bytes.Buffer)
	fmt.Printf("RESPONSE BUF: %s", buf)
	if buf == nil {
		return nil, errors.New("model.FindUserCommitsOnMultiRepo: Response is null")
	}
	buf.ReadFrom(res.Body)
	return json_to_userCommits_array(buf.String())
}

//
// private function for turning json response into an array in go
//
func json_to_userCommits_array(js string) ([]UserCommits, error) {
	var f interface{}
	var a []UserCommits
	err := json.Unmarshal([]byte(js), &f)
	if err != nil {
		return nil, err
	}
	m := f.(map[string]interface{})
	docs := m["docs"].([]interface{})
	for i := range docs {
		//c := &UserCommits{}
		c := docs[i].(map[string]interface{})
		day, _ := time.Parse(time.RFC3339, c["date"].(string))
		a = append(a, UserCommits{
			c["user"].(string),
			c["url"].(string),
			day,
			int(c["commits"].(float64)),
			c["_id"].(string),
			c["_rev"].(string),
		})
	}
	return a, nil
}
