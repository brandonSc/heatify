package route

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"HeatMap",
		"GET",
		"/heatmap",
		HeatMap,
	},
	Route{
		"HeatMapRepo",
		"GET",
		"/heatmap/repo",
		HeatMapRepo,
	},
	Route{
		"HeatMapUser",
		"GET",
		"/heatmap/user",
		HeatMapUser,
	},
}
