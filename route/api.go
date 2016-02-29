package route

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"hub.jazz.net/git/schurman93/Git-Monitor/model"
)

//
// send back json array of all users commits on all traked repositories
// @param user is the git commit author provided in a url parameter
//
func GetCommitsByUser(w http.ResponseWriter, r *http.Request) {
	user, err := url.QueryUnescape(r.URL.Query().Get("user"))
	if err != nil {
		fmt.Fprintf(w, `{"error":"'user' parameters is required."}`)
		return
	}
	commits := model.FindCommitsByUser(user)
	js, err := json.Marshal(commits)
	if err != nil {
		fmt.Printf("GetCommitsByUser: error marshaling user commits to json: %s\n", err)
		return
	}
	fmt.Fprintf(w, string(js))
}

//
// send back json array of all users commits on the provided repository
// @param user is the git commit author provided in a url parameter
// @param repo is the git url to look under
//
func GetCommitsByUserOnRepo(w http.ResponseWriter, r *http.Request) {
	user, err := url.QueryUnescape(r.URL.Query().Get("user"))
	if err != nil || user == "" {
		fmt.Fprintf(w, `{"error":"'user' parameters is required."}`)
		return
	}
	repo, err := url.QueryUnescape(r.URL.Query().Get("repo"))
	if err != nil || repo == "" {
		fmt.Fprintf(w, `{"error":"'repo' parameters is required."}`)
		return
	}
	commits := model.FindCommitsByUserOnRepo(user, repo)
	js, err := json.Marshal(commits)
	if err != nil {
		fmt.Printf("GetCommitsByUser: error marshaling user commits to json: %s\n", err)
		return
	}
	fmt.Fprintf(w, string(js))
}

//
// send back json array of all users commits on the provided array of repos (e.g. squad)
// @param user is the git commit author provided in a url parameter
// @param repo is the git url to look under
//
func GetCommitsByUserOnMultiRepo(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL)
	if err != nil {
		panic(err)
	} else {
		m, _ := url.ParseQuery(u.RawQuery)
		fmt.Println(m)
	}
	user, err := url.QueryUnescape(r.URL.Query().Get("user"))
	if err != nil || user == "" {
		fmt.Fprintf(w, `{"error":"'user' parameters is required."}`)
		return
	}
	repo, err := url.QueryUnescape(r.URL.Query().Get("repo"))
	if err != nil || repo == "" {
		fmt.Fprintf(w, `{"error":"'repo' parameters is required."}`)
		return
	}
	commits := model.FindCommitsByUserOnRepo(user, repo)
	js, err := json.Marshal(commits)
	if err != nil {
		fmt.Printf("GetCommitsByUser: error marshaling user commits to json: %s\n", err)
		return
	}
	fmt.Fprintf(w, string(js))
}
