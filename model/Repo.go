package model

//
// Indicates the number of commits to a repository on particular day
//
type RepoCommits struct {
	URL     string `json:"url"`     // URL of the repository
	Date    string `json:"date"`    // date of the commits on the repo: YYYY/MM/DD
	LastId  int    `json:"lastId"`  // most recent commit ID (hex) for distinguishing new commits on current date
	Commits int    `json:"commits"` // number of commits on this day
}
