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
	Route{
		"About",
		"GET",
		"/about",
		ShowAbout,
	},
	Route{
		"AllRepos",
		"GET",
		"/heatmap/repo/all",
		AllRepos,
	},
	Route{
		"AllSquads",
		"GET",
		"/heatmap/squads/all",
		AllSquads,
	},
	Route{
		"HeatMapSquad",
		"GET",
		"/heatmap/squads",
		HeatMapSquad,
	},
	Route{
		"ApiUserCommits",
		"GET",
		"/api/commits/user",
		GetCommitsByUser,
	},
	Route{
		"ApiUserCommitsRepo",
		"GET",
		"/api/commits/repo/user",
		GetCommitsByUserOnRepo,
	},
	Route{
		"ApiUserCommitsSquad",
		"GET",
		"/api/commits/squad/user",
		GetCommitsByUserOnSquad,
	},
	Route{
		"ApiCommunityCommitsSquad",
		"GET",
		"/api/commits/squad/community",
		GetCommitsByCommunityOnSquad,
	},
}
