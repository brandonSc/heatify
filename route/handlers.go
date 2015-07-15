package route

import (
	//"encoding/json"
	"fmt"
	//"github.com/gorilla/mux"
	"html/template"
	"hub.jazz.net/git/schurman93/Git-Monitor/gitutil"
	"net/http"
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
	res := gitutil.ParseCommits("hub.jazz.net/git/schurman93/metrics-service")
	p, err := LoadPage("heatmap")
	if err != nil {
		fmt.Println("Error: %s", err)
		fmt.Fprintf(w, "Error")
		return
	}
	p.Title = "HeatMap of 'metrics-service'"
	p.Data = res
	fmt.Println(res)
	t, _ := template.ParseFiles("views/heatmap.html")
	t.Execute(w, p)
}
