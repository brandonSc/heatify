package gitutil

import (
	"strings"
)

func UrlToDir(url string) string {
	if strings.Contains(url, "github.com") {
		// regular github URL
		return github_url_to_dir(url)
	} else if strings.Contains(url, "hub.jazz.net") {
		// IBM jazz hub
		return github_url_to_dir(url)
	} else if strings.Contains(url, "github.rtp") {
		// IBM GitLab url
		return "not implemented"
	} else {
		return "repository not supported: " + url
	}
}

func DirToUrl(path string) {
	if strings.Contains(path, "github.com") {
		// regular github URL
	} else if strings.Contains(path, "hub.jazz.net") {
		// IBM jazz hub
	} else if strings.Contains(path, "github.rtp") {
		// IBM GitLab url
	}
}

func github_url_to_dir(url string) string {
	return strings.Replace(url, "/", ".", -1)
}

func jazzhub_url_to_dir(url string) string {
	return strings.Replace(url, "/", ".", -1)
}
