package model

//
// Stores the number of commits made to a repo each day by a particular user
//
type UserRepoCommits struct {
	UserName string `json:"username"` // username of the user
	URL      string `json:"url"`      // URL of the repository
	Date     string `json:"date"`     // date of the commits on the repo: YYYY/MM/DD
	LastId   int    `json:"lastId"`   // most recent commit ID (hex) for distinguishing new commits on current date
	Commits  int    `json:"commits"`  // number of commits on this day
}
