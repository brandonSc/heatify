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
// non-blocking `/heatmap` handler
// ** cloning more than one repo at once will cause errors **
// ** avoid using this for now **
//
func HeatMapAsync(w http.ResponseWriter, r *http.Request) {
	go HeatMap(w, r)
}

//
// handle the request for `/heatmap`
//
func HeatMap(w http.ResponseWriter, r *http.Request) {
	repo, err := url.QueryUnescape(r.URL.Query().Get("repo"))
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

	t, _ := template.ParseFiles("views/heatmap.html")
	t.Execute(w, p)
}

//
// Generate a heatmap for a specific repo `/heatmap/{repo}`
//
/*
func GenHeatMap(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	repo := vars["repo"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
*/

func parse_repo_name(url string) string {
	parts := strings.SplitAfter(url, "/")
	return parts[len(parts)-1]
}
