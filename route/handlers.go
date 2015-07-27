package route

import (
	//"encoding/json"
	"fmt"
	//"github.com/gorilla/mux"
	"html/template"
	"hub.jazz.net/git/schurman93/Git-Monitor/gitutil"
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
		fmt.Fprintf(w, "Error %s", err)
		return
	}

	name := parse_repo_name(repo)

	res, err := gitutil.ParseCommits(repo)
	if err != nil {
		fmt.Println("Error cloning or parsing repository: %s", err)
		fmt.Fprintf(w, "Error %s", err)
		return
	}

	p, err := LoadPage("heatmap")
	if err != nil {
		fmt.Println("Error: %s", err)
		fmt.Fprintf(w, "Error")
		return
	}
	p.Title = "HeatMap of '" + name + "'"
	p.Data = res

	//fmt.Println(res)

	page := template.Must(template.ParseFiles(
		"views/_base.html",
		"views/heatmap.html",
	))
	page.Execute(w, p)
}

func HeatMapUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "TODO")
}

//
// Return a page about heatmaps
//
func HeatMap(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TODO")
}

//
// Get the name of the repository from the full git URL
// (works for jazz and github)
//
func parse_repo_name(url string) string {
	parts := strings.SplitAfter(url, "/")
	str := parts[len(parts)-1]
	str = strings.TrimRight(str, ".git")
	return str
}
