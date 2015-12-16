package route

import (
	"encoding/json"
	"fmt"
	"html/template"
	"hub.jazz.net/git/schurman93/Git-Monitor/gitutil"
	"hub.jazz.net/git/schurman93/Git-Monitor/model"
	"net/http"
	"net/url"
	"strings"
)

const viewDir = "../views/"

//
// serve the default landling page for the route `/`
//
func Index(w http.ResponseWriter, r *http.Request) {
	var index = template.Must(template.ParseFiles(
		"views/_base.html",
		"views/index.html",
	))
	index.Execute(w, nil)
}

//
// non-blocking `/heatmap` handler runs in a go-routine
// ** cloning more than one repo at once will cause errors **
// ** avoid using this for now **
//
func HeatMapAsync(w http.ResponseWriter, r *http.Request) {
	go HeatMapRepo(w, r)
}

//
// Return the page `/heatmap/repo` page
// which generates a heatmap for a single repository
//
func HeatMapRepo(w http.ResponseWriter, r *http.Request) {
	repo, err := url.QueryUnescape(r.URL.Query().Get("url"))
	repo = strings.TrimPrefix(repo, "http://")
	repo = strings.TrimPrefix(repo, "https://")
	if err != nil {
		fmt.Println("Error decoding repository from URL: %s", err)
		ShowError(w, ERROR_REPO_URL)
		return
	}

	name, repo := parse_repo(repo)
	if name == "" {
		name = "Git-Monitor"
	}

	var data string
	exists := gitutil.CheckExists(repo)
	if exists {
		dbCommits := model.DbRetrieveAllRepoCommits(repo)
		b, err := json.Marshal(dbCommits)
		if err != nil {
			ShowError(w, ERROR_CLONE_REPO)
			return
		}
		data = string(b)
	} else {
		res, err := gitutil.CloneRepo(repo)
		if err != nil {
			fmt.Println("Error cloning or parsing repository: %s", err)
			ShowError(w, ERROR_CLONE_REPO)
			return
		}
		data = res
	}

	p, err := LoadPage("heatmap")
	if err != nil {
		fmt.Println("Error: %s", err)
		ShowError(w, ERROR_PARSE_PAGE)
		return
	}
	p.Title = name
	p.Data = data
	p.Repo = repo

	page := template.Must(template.ParseFiles(
		"views/_base.html",
		"views/heatmap.html",
	))
	page.Execute(w, p)
}

//
// show the HeatMap search page for Users
//
func HeatMapUser(w http.ResponseWriter, r *http.Request) {
	p, err := LoadPage("users")
	if err != nil {
		fmt.Println("Error: %s", err)
		fmt.Fprintf(w, "An error occured while rendering a page")
		return
	}
	p.Title = "Users"

	page := template.Must(template.ParseFiles(
		"views/_base.html",
		"views/users.html",
	))
	page.Execute(w, p)
}

//
// Return a page about heatmaps
//
func HeatMap(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TODO")
}

//
// Return a list of all repositories
// The user could pick from this list to view the HeatMap
// ** Private repositories must not be displayed **
//
func AllRepos(w http.ResponseWriter, r *http.Request) {
	p, err := LoadPage("repolist")
	if err != nil {
		fmt.Println("Error: %s", err)
		fmt.Fprintf(w, "An error occured while rendering a page")
		return
	}
	p.Title = "About"
	p.Extra = gitutil.AllTrackedRepos()

	page := template.Must(template.ParseFiles(
		"views/_base.html",
		"views/repolist.html",
	))
	page.Execute(w, p)
}

//
// Get the name of the repository from the full git URL
// (works for jazz and github)
// if random is provided, then the original url is replaced
//
func parse_repo(url string) (string, string) {
	if url == "random" {
		url = gitutil.GetRandomRepo()
	}
	parts := strings.SplitAfter(url, "/")
	str := parts[len(parts)-1]
	str = strings.TrimSuffix(str, ".git")
	return str, url
}

//
// show an error page
//
func ShowError(w http.ResponseWriter, reason string) {
	p, err := LoadPage("error")
	if err != nil {
		fmt.Println("Error: %s", err)
		fmt.Fprintf(w, "An error occured while rendering a page")
		return
	}
	p.Title = "Error :("
	p.Data = reason

	page := template.Must(template.ParseFiles(
		"views/_base.html",
		"views/error.html",
	))
	page.Execute(w, p)
}

//
// show the 'about' page /about
//
func ShowAbout(w http.ResponseWriter, r *http.Request) {
	p, err := LoadPage("about")
	if err != nil {
		fmt.Println("Error: %s", err)
		fmt.Fprintf(w, "An error occured while rendering a page")
		return
	}
	p.Title = "About"

	page := template.Must(template.ParseFiles(
		"views/_base.html",
		"views/about.html",
	))
	page.Execute(w, p)
}
