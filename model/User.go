package model

type UserCommits struct {
	Username string   `json:"username"`        // username (e.g. Slack name)
	Image    string   `json:"image,omitempty"` // URL for profile pic
	Aliases  []string `json:"aliases"`         // an array of Git account aliases
}

func IsUserAlias(user string) bool {
	return strings.Contains(user, "<") && strings.Contains(user, ">")
}
