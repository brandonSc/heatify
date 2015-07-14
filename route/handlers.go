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

func Index(w http.ResponseWriter, r *http.Request) {
	var index = template.Must(template.ParseFiles(
		"templates/_base.html",
		"templates/index.html",
	))
	index.Execute(w, nil)
}

func HeatMap(w http.ResponseWriter, r *http.Request) {
	var res = gitutil.ParseCommits("hub.jazz.net/git/schurman93/metrics-service")
	p, err := LoadPage("heatmap")
	if err != nil {
		fmt.Println("Error: %s", err)
		fmt.Fprintf(w, "Error")
		return
	}
	p.Data = res
	fmt.Println(res)
	t, _ := template.ParseFiles("views/heatmap.html")
	t.Execute(w, p)
}

/*
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
*/
