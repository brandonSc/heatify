package route

import (
	//"encoding/json"
	//"fmt"
	//"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var index = template.Must(template.ParseFiles(
		"templates/_base.html",
		"templates/index.html",
	))
	index.Execute(w, nil)
}

func Static(w http.ResponseWriter, r *http.Request) {
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	/*
		todos := Todos{
			Todo{Name: "Write presentation"},
			Todo{Name: "Host meetup"},
		}

		if err := json.NewEncoder(w).Encode(todos); err != nil {
			panic(err)
		}
	*/
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	/*
		vars := mux.Vars(r)
		todoId := vars["todoId"]
		fmt.Fprintln(w, "Todo show:", todoId)
	*/
}
