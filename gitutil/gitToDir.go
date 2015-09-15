package gitutil

import (
	"strings"
)

const (
	GITHUB_URL_IDENTIFIER = "github.com"
	GITHUB_DIR_IDENTIFIER = "github/com"
	JAZZ_URL_IDENTIFIER   = "hub.jazz.net"
	JAZZ_DIR_IDENTIFIER   = "hub/jazz/net"
)

func UrlToDir(url string) string {
	if strings.Contains(url, GITHUB_URL_IDENTIFIER) {
		// regular github URL
		return github_url_to_dir(url)
	} else if strings.Contains(url, JAZZ_URL_IDENTIFIER) {
		// IBM jazz hub
		return github_url_to_dir(url)
	} else if strings.Contains(url, "github.rtp") {
		// IBM GitLab url
		return "not implemented"
	} else {
		return "repository type not supported: " + url
	}
}

func DirToUrl(path string) string {
	if strings.Contains(path, GITHUB_DIR_IDENTIFIER) {
		// standard github URL
		return github_dir_to_url(path)
	} else if strings.Contains(path, JAZZ_DIR_IDENTIFIER) {
		// IBM jazz hub
		return jazzhub_dir_to_url(path)
	} else if strings.Contains(path, "github/rtp") {
		// IBM GitLab url
		return "not implemented"
	} else {
		return "repository type not supported: " + path
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
	url = strings.Replace(url, "hub/jazz/net.git", "hub.jazz.net/git", -1)
	url = strings.Replace(url, "/com", ".com", -1)
	return url
}
