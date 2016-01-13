package gitutil

import (
	"strings"
)

const (
	GITHUB_URL_IDENTIFIER    = "github.com"
	GITHUB_DIR_IDENTIFIER    = "github.com"
	JAZZ_URL_IDENTIFIER      = "hub.jazz.net"
	JAZZ_DIR_IDENTIFIER      = "hub.jazz.net"
	GITLAB_URL_IDENTIFIER    = "github.rtp"
	GITLAB_DIR_IDENTIFIER    = "github.rtp"
	GITHUBIBM_URL_IDENTIFIER = "github.ibm.com"
	GITHUBIBM_DIR_IDENTIFIER = "github.ibm.com"
)

// should return errors here instead ..
func UrlToDir(url string) string {
	if strings.Contains(url, GITHUB_URL_IDENTIFIER) {
		// regular github URL
		return github_url_to_dir(url)
	} else if strings.Contains(url, JAZZ_URL_IDENTIFIER) {
		// IBM jazz hub
		return github_url_to_dir(url)
	} else if strings.Contains(url, GITLAB_URL_IDENTIFIER) {
		// IBM GitLab url
		return gitlab_url_to_dir(url)
	} else if strings.Contains(url, GITHUBIBM_URL_IDENTIFIER) {
		// IBM private github
		return githubIBM_url_to_dir(url)
	} else {
		return "repository type not supported: " + url
	}
}

func DirToUrl(path string) string {
	if strings.Contains(path, ".clones") {
		path = strings.Replace(path, ".clones", "", -1)
	}
	if strings.Contains(path, GITHUB_DIR_IDENTIFIER) {
		// standard github URL
		return github_dir_to_url(path)
	} else if strings.Contains(path, JAZZ_DIR_IDENTIFIER) {
		// IBM jazz hub
		return jazzhub_dir_to_url(path)
	} else if strings.Contains(path, GITLAB_DIR_IDENTIFIER) {
		// IBM GitLab url
		return gitlab_dir_to_url(path)
	} else if strings.Contains(path, GITHUBIBM_DIR_IDENTIFIER) {
		// IBM private Github
		return githubIBM_dir_to_url(path)
	} else {
		return "directory type not supported: " + path
	}
}

func github_url_to_dir(url string) string {
	return strings.Replace(url, "/", ".", -1)
}

func github_dir_to_url(path string) string {
	url := strings.Replace(path, ".", "/", -1)
	url = strings.Replace(url, "/git", ".git", -1)
	url = strings.Replace(url, "/com", ".com", -1)
	return url
}

func jazzhub_url_to_dir(url string) string {
	return strings.Replace(url, "/", ".", -1)
}

func jazzhub_dir_to_url(path string) string {
	url := strings.Replace(path, ".", "/", -1)
	url = strings.Replace(url, "hub/jazz/net", "hub.jazz.net", -1)
	url = strings.Replace(url, "/com", ".com", -1)
	return url
}

//
// e.g.: consumes:
// github.rtp.raleigh.ibm.com/project-alchemy/executive-dashboard.git
// e.g.: produces;
// github.rtp.raleigh.ibm.com.project-alchemy.executive-dashboard.git
//
func gitlab_url_to_dir(url string) string {
	return strings.Replace(url, "/", ".", -1)
}

//
// e.g.: consumes:
// github.rtp.raleigh.ibm.com.project-alchemy.executive-dashboard.git
// e.g.: produces;
// github.rtp.raleigh.ibm.com/project-alchemy/executive-dashboard.git
//
func gitlab_dir_to_url(path string) string {
	url := strings.Replace(path, ".", "/", -1)
	// changes to: github/rtp/raleigh/ibm/com/project-alchemy/executive-dashboard/git
	url = strings.Replace(url, "github/rtp/raleigh/ibm/com", "github.rtp.raleigh.ibm.com", -1)
	// changes to: github.rtp.raleigh.ibm.com/project-alchemy/executive-dashboard/git
	url = strings.Replace(url, "/git", ".git", -1)
	// changes to: github.rtp.raleigh.ibm.com/project-alchemy/executive-dashboard.git
	return url
}

//
// e.g.: consumes:
// github.ibm.com/alchemy-dashboard/executive-dashboard-ui.git
// e.g.: produces:
// github.ibm.com.alchemy-dashboard.executive-dashboard-ui.git
//
func githubIBM_url_to_dir(url string) string {
	return strings.Replace(url, "/", ".", -1)
}

//
// e.g.: consumes:
// github.ibm.com.alchemy-dashboard.executive-dashboard-ui.git
// e.g.: produces:
// github.ibm.com/alchemy-dashboard/executive-dashboard-ui.git
//
func githubIBM_dir_to_url(path string) string {
	url := strings.Replace(path, ".", "/", -1)
	// changes to: github/ibm/com/alchemy-dashboard/executive-dashboard-ui/git
	url = strings.Replace(url, "github/ibm/com", "github.ibm.com", -1)
	// changes to: github.ibm.com/alchemy-dashboard/executive-dashboard-ui/git
	url = strings.Replace(url, "/git", ".git", -1)
	// changes to: github.ibm.com/alchemy-dashboard/executive-dashboard-ui.git
	return url
}
