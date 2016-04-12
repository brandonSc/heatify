package route

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

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
	if model.IsUserAlias(user) {
		users = model.GetGitUsers(user)
		commits, err := model.FindMultiUserCommits(users)
	} else {
		commits, err := model.FindUserCommits(user)
	}
	if err != nil {
		fmt.Fprintf(w, `{"error":"problem downloading commits from the cloud"}`)
		fmt.Printf("Error loading commits from cloudant for user=" + user)
		return
	}
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
	w.Header().Set("Content-Type", "application/json")
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
	commits, err := model.FindUserCommitsOnRepo(user, repo)
	if err != nil {
		fmt.Fprintf(w, `{"error":"problem downloading commits from the cloud"}`)
		fmt.Printf("Error loading commits from cloudant for user=" + user)
		return
	}
	js, err := json.Marshal(commits)
	if err != nil {
		fmt.Printf("GetCommitsByUser: error marshaling user commits to json: %s\n", err)
		return
	}
	fmt.Fprintf(w, string(js))
}

//
// GET /api/commits/squad
// send back json array of all users commits on the provided array of repos (e.g. squad)
// @param user is the git commit author provided in a url parameter
// @param squad
//
func GetCommitsByUserOnSquad(w http.ResponseWriter, r *http.Request) {
	squad, err := url.QueryUnescape(r.URL.Query().Get("squad"))
	if err != nil || squad == "" {
		fmt.Fprintf(w, `{"error":"'squad' parameters is required."}`)
		return
	}
	user, err := url.QueryUnescape(r.URL.Query().Get("user"))
	if err != nil || user == "" {
		fmt.Fprintf(w, `{"error":"'user' parameters is required."}`)
		return
	}
	s, err := model.InitSquadFromJson(squad + ".json")
	if err != nil {
		fmt.Fprintf(w, `{"error":"Squad not found."}`)
		return
	}
	commits, err := model.FindUserCommitsOnMultiRepo(user, s.Repos)
	js, err := json.Marshal(commits)
	if err != nil {
		fmt.Printf("GetCommitsByUser: error marshaling user commits to json: %s\n", err)
		return
	}
	fmt.Fprintf(w, string(js))
}

//
// GET /api/commits/squad/community
// send back json array of all users commits on the provided array of repos (e.g. squad)
// EXCLUDING the list of users found in the squad
// this gives back a list commits on a squad, provded by the outside community.
// @param user is the git commit author provided in a url parameter
// @param squad
//
func GetCommitsByCommunityOnSquad(w http.ResponseWriter, r *http.Request) {
	squad, err := url.QueryUnescape(r.URL.Query().Get("squad"))
	if err != nil || squad == "" {
		fmt.Fprintf(w, `{"error":"'squad' parameters is required."}`)
		return
	}
	s, err := model.InitSquadFromJson(squad + ".json")
	if err != nil {
		fmt.Fprintf(w, `{"error":"Squad not found."}`)
		return
	}
	commits, err := model.FindCommunityCommitsOnMultiRepo(s.Members, s.Repos)
	js, err := json.Marshal(commits)
	if err != nil {
		fmt.Printf("GetCommitsByUser: error marshaling user commits to json: %s\n", err)
		return
	}
	fmt.Fprintf(w, string(js))
}

//
// GET /api/squad/users
// send back a list of squad members, with their rank in the squad
// (i.e. total number of commits to all repositoreis)
// @param squad is the name of the Squad
//
/*
func GetUsersAndRankForSquad(w http.ResponseWriter, r *http.Request) {
	squad, err := url.QueryUnescape(r.URL.Query().Get("squad"))
	if err != nil || squad == "" {
		fmt.Fprintf(w, `{"error":"'squad' parameters is required."}`)
		return
	}
	s, err := model.InitSquadFromJson(squad + ".json")
	squad, err := url.QueryUnescape(r.URL.Query().Get("squad"))
	if err != nil || squad == "" {
		fmt.Fprintf(w, `{"error":"'squad' parameters is required."}`)
		return
	}
	s, err := model.InitSquadFromJson(squad + ".json")
	if err != nil {
		fmt.Fprintf(w, `{"error":"Squad not found."}`)
		return
	}
	if err != nil {
		fmt.Fprintf(w, `{"error":"Squad not found."}`)
		return
	}

	users, err := find_users_by_rank(s)
	if err != nil {
		fmt.Fprintf(w, `{"error":"problem looking up users by rank in database"}`)
		return
	}

	fmt.Fprintf(w, users)
}

func GetSquadCommunityCommits(w http.ResponseWriter, r *http.Request) {
}

// private func for looking up squad members by rank in cloudant
func find_users_by_rank(squad model.Squad) (string, error) {

}
*/
